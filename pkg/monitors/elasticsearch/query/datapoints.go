package query

import (
	"fmt"
	"strconv"

	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// Walks through the response, collecting dimensions and datapoints depending on the
// type of aggregation at each recursive level
func collectDatapoints(resBody HTTPResponse, aggsMeta map[string]*AggregationMeta, sfxDimensions map[string]string) []*datapoint.Datapoint {
	out := make([]*datapoint.Datapoint, 0)
	aggsResult := resBody.Aggregations

	for k, v := range aggsResult {
		// each aggregation at the highest level starts with an empty set of dimensions
		out = append(out, collectDatapointsHelper(k, *v, aggsMeta, sfxDimensions)...)
	}

	return out
}

func collectDatapointsHelper(aggName string, aggRes aggregationResponse,
	aggsMeta map[string]*AggregationMeta, sfxDimensions map[string]string) []*datapoint.Datapoint {

	aggType := aggsMeta[aggName].Type

	sfxDatapoints := make([]*datapoint.Datapoint, 0)

	// The absence of "doc_count" and "buckets" field is a good indicator that
	// the aggregation is a metric aggregation
	if isMetricAggregation(&aggRes) {
		return collectDatapointsFromMetricAggregation(&aggRes, aggName, aggType, sfxDimensions)
	}

	// Send document count as metrics when there are no metric aggregations specified
	// under a bucket aggregation and there aren't sub aggregations as well
	if isTerminalBucketAggregation(&aggRes) {
		return collectDocCountFromTerminalBucketAggregation(&aggRes, aggName, aggType, sfxDimensions)
	}

	// Recursively collect all datapoints from buckets at this level
	for _, b := range aggRes.Buckets {
		// Pick the current bucket's key as a dimension before recursing down to the next level
		sfxDimensionsForBucket := utils.CloneStringMap(sfxDimensions)
		sfxDimensionsForBucket[aggName] = b.Key.(string)

		for k, v := range b.SubAggregations {
			sfxDatapoints = append(sfxDatapoints, collectDatapointsHelper(k, v, aggsMeta, sfxDimensionsForBucket)...)
		}
	}

	// Recursively collect datapoint from sub aggregations
	for k, v := range aggRes.SubAggregations {
		sfxDatapoints = append(sfxDatapoints, collectDatapointsHelper(k, v, aggsMeta, sfxDimensions)...)
	}

	return sfxDatapoints
}

// Collects "doc_count" as a SFx datapoint if a bucket aggregation does not have
// sub metric aggregations
func collectDocCountFromTerminalBucketAggregation(aggRes *aggregationResponse, aggName string,
	aggType string, dims map[string]string) []*datapoint.Datapoint {
	dimsForBucket := utils.CloneStringMap(dims)
	dimsForBucket["bucket_aggregation_name"] = aggName

	return []*datapoint.Datapoint{
		{
			Metric:     fmt.Sprintf("%s.%s", aggType, "doc_count"),
			Dimensions: dims,
			Value:      datapoint.NewIntValue(*aggRes.DocCount),
			MetricType: datapoint.Gauge,
		},
	}
}

// Collects datapoints from supported metric aggregations
func collectDatapointsFromMetricAggregation(aggRes *aggregationResponse, aggName string,
	aggType string, sfxDimensions map[string]string) []*datapoint.Datapoint {

	out := make([]*datapoint.Datapoint, 0)

	// Add metric aggregation name as a dimension
	sfxDimensions["metric_aggregation_name"] = aggName

	switch aggType {
	case "extended_stats":
		out = append(out, getDatapointsFromExtendedStats(aggRes, sfxDimensions)...)
	case "percentiles":
		out = append(out, getDatapointsFromPerciltes(aggRes, sfxDimensions)...)
	default:
		metricName := aggType
		dp, ok := collectDatapoint(metricName, aggRes.Value, sfxDimensions)

		if !ok {
			log.WithFields(log.Fields{"aggregation_type": aggType,
				"aggregation_name": aggName}).Warnf("Invalid value found for stat: %v", aggRes.Value)
			return out
		}

		out = append(out, dp)
	}

	return out
}

// Collect datapoints from "extended_stats" metric aggregation
func getDatapointsFromExtendedStats(aggRes *aggregationResponse, dims map[string]string) []*datapoint.Datapoint {
	aggName := dims["metric_aggregation_name"]
	out := make([]*datapoint.Datapoint, 0)

	for k, v := range aggRes.OtherValues {
		switch k {
		case "std_deviation_bounds":
			m, ok := v.(map[string]interface{})

			if !ok {
				log.WithFields(log.Fields{"extended_stat": k,
					"aggregation_name": aggName}).Warnf("Invalid value found for stat: %v", v)
				continue
			}

			for bk, bv := range m {
				metricName := fmt.Sprintf("%s.%s.%s", "extended_stats", k, bk)
				dp, ok := collectDatapoint(metricName, bv, dims)

				if !ok {
					log.WithFields(log.Fields{"extended_stat": k,
						"aggregation_name": aggName}).Warnf("Invalid value found for stat: %v", bv)
					continue
				}

				out = append(out, dp)
			}
		default:
			metricName := fmt.Sprintf("%s.%s", "extended_stats", k)
			dp, ok := collectDatapoint(metricName, v, dims)

			if !ok {
				log.WithFields(log.Fields{"extended_stat": k,
					"aggregation_name": aggName}).Warnf("Invalid value found for stat: %v", v)
				continue
			}

			out = append(out, dp)
		}
	}

	return out
}

// Collect datapoint from "percentiles" metric aggregation
func getDatapointsFromPerciltes(aggRes *aggregationResponse, dims map[string]string) []*datapoint.Datapoint {
	aggName := dims["metric_aggregation_name"]
	out := make([]*datapoint.Datapoint, 0)

	// Values are always expected to be a map between the percentile and the
	// actual value itself of the metric
	values, ok := aggRes.Values.(map[string]interface{})

	if !ok {
		log.WithFields(log.Fields{"aggregation_name": aggName}).Warnf("No valid values found in percentiles aggregation")
	}

	// Metric name is constituted of the aggregation type "percentiles" and the actual percentile
	// Metric names from this aggregation will look like "percentiles.p99", "percentiles.p50" and
	// the aggregation name used to compute the metric will be a sent in as "metric_aggregation_name"
	// dimension on the datapoint
	for k, v := range values {
		p, err := strconv.ParseFloat(k, 64)

		if err != nil {
			log.WithFields(log.Fields{"aggregation_name": aggName}).Warnf("Invalid percentile found: %s", k)
			continue
		}

		// Remove trailing zeros
		metricName := fmt.Sprintf("%s.p%s", "percetiles", strconv.FormatFloat(p, 'f', -1, 64))
		dp, ok := collectDatapoint(metricName, v, dims)

		if !ok {
			log.WithFields(log.Fields{"percentile": k,
				"aggregation_name": aggName}).Warnf("Invalid value found for percentile: %v", v)
			continue
		}

		out = append(out, dp)
	}

	return out
}

// Returns true if aggregation is a metric aggregation
func isMetricAggregation(aggRes *aggregationResponse) bool {
	return aggRes.DocCount == nil && len(aggRes.Buckets) == 0
}

// Returns true if bucket aggregation is at the deepest level without
// sub metric aggregations
func isTerminalBucketAggregation(aggRes *aggregationResponse) bool {
	return aggRes.DocCount != nil && len(aggRes.Buckets) == 0 && len(aggRes.SubAggregations) == 0
}

// Collects a single datapoint from an interface, returns false if no datapoint can be derived
func collectDatapoint(metricName string, value interface{}, dims map[string]string) (*datapoint.Datapoint, bool) {
	metricValue, ok := value.(float64)

	if !ok {
		return nil, false
	}

	return &datapoint.Datapoint{
		Metric:     metricName,
		Dimensions: dims,
		Value:      datapoint.NewFloatValue(metricValue),
		MetricType: datapoint.Gauge,
	}, true
}
