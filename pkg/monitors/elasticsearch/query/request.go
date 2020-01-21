package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Represents HTTP post data for elasticsearch search queries
// This format is generally determined by the format specified here:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html#_structuring_aggregations
// Note that it also depends on the aggregations supported by the
// query monitor
type HTTPPostData struct {
	// Holds a map of aggregation names to the corresponding aggregation
	Aggregations map[string]aggregationInfo `json:"aggs"`
}

// Represents both, bucket and metric aggregations
// See https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html
// for details about various type of bucket and metric aggregations
type aggregationInfo struct {
	// Holds aggregationInfo type and properties specific to the aggregation
	AggregationProps map[string]interface{} `json:"-"`
	// Sub aggregations within the parent aggregation
	SubAggregations map[string]aggregationInfo `json:"aggs,omitempty"`
}

// Required to handle certain fields specially
func (agg *aggregationInfo) UnmarshalJSON(b []byte) error {
	type aggregationRequestX aggregationInfo // prevent recursion
	var temp aggregationRequestX

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	*agg = aggregationInfo(temp)
	agg.AggregationProps = map[string]interface{}{}

	// Capture aggregation type and body as mentioned in docs here:
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html#_structuring_aggregations
	var aggregationBody map[string]map[string]interface{}
	if err := json.Unmarshal(b, &aggregationBody); err != nil {
		return err
	}

	// Based on elasticsearch, the underlying assumption made by this loop is that
	// aggregationBody will always only contain a map between the type of aggregation
	// to the aggregation body
	for k, v := range aggregationBody {
		// These fields have special meaning in the request and this loop does not care about them
		if k == "aggs" || k == "meta" || k == "aggregations" {
			continue
		}

		agg.AggregationProps[k] = v
	}

	return nil
}

// Returns a response for a given elasticsearch query
func (es ESQueryHTTPClient) makeHTTPRequestFromConfig(mq *metricQuery) ([]byte, error) {
	url := fmt.Sprintf("%s://%s:%s/%s/_search?", es.esClient.Scheme, es.esClient.Host, es.esClient.Port, mq.Index)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(mq.HTTPPostData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := es.esClient.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	//fmt.Println("response Body:", string(body))
	return ioutil.ReadAll(resp.Body)
}

type AggregationMeta struct {
	Type string
}

// Returns a map of aggregation names to aggregation types
func (reqBody *HTTPPostData) getAggregationsMeta() (map[string]*AggregationMeta, error) {
	aggsToType := map[string]string{}

	for aggName, agg := range reqBody.Aggregations {
		err := getAggregationsWithType(aggName, agg, aggsToType)

		if err != nil {
			return nil, err
		}
	}

	return getAggregationsMetaHelper(aggsToType), nil
}

func getAggregationsWithType(aggName string, agg aggregationInfo, out map[string]string) error {
	a, err := getAggregationType(agg)

	if err != nil {
		return err
	}

	out[aggName] = a

	for subAggName, subAgg := range agg.SubAggregations {
		if err := getAggregationsWithType(subAggName, subAgg, out); err != nil {
			return err
		}
	}

	return nil
}

func getAggregationType(agg aggregationInfo) (string, error) {
	for k := range agg.AggregationProps {
		return k, nil
	}

	return "", errors.Errorf("Unable to determine type for %v", agg)
}

func getAggregationsMetaHelper(aggsToType map[string]string) map[string]*AggregationMeta {
	out := map[string]*AggregationMeta{}

	for aggName, aggType := range aggsToType {
		out[aggName] = &AggregationMeta{
			Type: aggType,
		}
	}
	return out
}
