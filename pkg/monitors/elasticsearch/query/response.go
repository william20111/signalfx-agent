package query

import (
	"encoding/json"
)

// Represents all kinds of supported aggregations. aggregationResponse type can
// be determined in the following manner -
// 1) single bucket aggregationInfo contains only "doc_count" and a map to sub aggregations
// 2) multi bucket aggregationInfo contains "buckets"and a map to sub aggregations
// 3) metric aggregationInfo contains "value" or "values"
// These structs are defined based on inputs from
// https://github.com/elastic/elasticsearch/tree/master/server/src/main/java/org/elasticsearch/search/aggregations
type aggregationResponse struct {
	// Number of documents backing the aggregationInfo
	DocCount *int64 `json:"doc_count,omitempty"`
	// Non nil for single value metric aggregations
	Value interface{} `json:"value,omitempty"`
	// Non nil for multi-value metrics aggregations
	Values interface{} `json:"values,omitempty"`
	// Map of sub aggregations with aggregationInfo names, includes both bucket
	// and metric aggregations
	SubAggregations map[string]aggregationResponse `json:"-"`
	// Non nil for multi-value bucket aggregations
	Buckets []*bucketResponse `json:"buckets,omitempty"`
	// This field only exists to handle cases where buckets are returned as a
	// map from the key to the object itself. In such cases the "buckets" field
	// will be post processed and put into the "Buckets" field of this struct
	BucketsMap map[string]*bucketResponse `json:"buckets,omitempty"`
	// All other key-value pairs
	OtherValues map[string]interface{} `json:"-"`
}

// Represents a slice ("bucket") of ES documents
type bucketResponse struct {
	// Value uniquely identifying a bucket for a given bucket aggregationInfo
	Key interface{} `json:"key,omitempty"`
	// Reported in cases where buckets are grouped by a non string key,
	// for instance time
	KeyAsString string `json:"key_as_string,omitempty"`
	// Number of documents in the bucket
	DocCount *int64 `json:"doc_count,omitempty"`
	// Map of sub aggregations with aggregationInfo names, includes both bucket
	// and metric aggregations
	SubAggregations map[string]aggregationResponse `json:"-"`
}

type HTTPResponse struct {
	Aggregations map[string]*aggregationResponse `json:"aggregations"`
}

// Required to handle certain fields specially
func (agg *aggregationResponse) UnmarshalJSON(b []byte) error {
	type aggregationResponseX aggregationResponse // prevent recursion
	var temp aggregationResponseX

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	*agg = aggregationResponse(temp)
	agg.SubAggregations = map[string]aggregationResponse{}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	t := map[string]interface{}{}
	for k, v := range m {

		switch k {
		// "value" is a special key in single value metric aggregations
		// and will be picked as a field on aggregationResponse
		case "value":
			continue
		// "values" is a special key in multi value metric aggregations
		// and will be picked as a field on aggregationResponse
		case "values":
			continue
		// Some aggregations return buckets as a map of "key" to the bucket
		// object. Handle such cases separately. Cases when buckets are returned
		// as an array where each object has the "key" field, the unmarshaller
		// for "aggregationResponse" will pick it up
		case "buckets":
			_, ok := v.(map[string]interface{})
			if ok {
				buckets := make([]*bucketResponse, 0)
				for bk, bv := range v.(map[string]interface{}) {
					b1, err := json.Marshal(bv)
					if err != nil {
						return err
					}

					var bucket bucketResponse
					if err := json.Unmarshal(b1, &bucket); err != nil {
						return err
					}

					bucket.Key = bk
					buckets = append(buckets, &bucket)
				}
				agg.Buckets = buckets
			}
		default:
			// Absence of these "doc_count" and "buckets" fields is a good indicator that the aggregation
			// is not a bucket aggregation. If this metric aggregation does not happen to have the standard
			// "value" or "values" fields through which values are typically available put the fields into
			// OtherValues field of the struct
			if m["doc_count"] == nil && m["buckets"] == nil && m["value"] == nil && m["values"] == nil {
				t[k] = v
				continue
			}

			_, ok := v.(map[string]interface{})

			if ok {
				b1, err := json.Marshal(v)

				if err != nil {
					return err
				}

				var subAgg aggregationResponse
				if err := json.Unmarshal(b1, &subAgg); err != nil {
					return err
				}

				agg.SubAggregations[k] = subAgg
			}
		}
	}

	agg.OtherValues = t
	return nil
}

// Required to handle certain fields specially
func (bucket *bucketResponse) UnmarshalJSON(b []byte) error {
	type bucketResponseX bucketResponse // prevent recursion
	var temp bucketResponseX

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	*bucket = bucketResponse(temp)
	bucket.SubAggregations = map[string]aggregationResponse{}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	for k, v := range m {
		_, ok := v.(map[string]interface{})
		if ok {
			b1, err := json.Marshal(v)

			if err != nil {
				return err
			}

			var subAgg aggregationResponse
			if err := json.Unmarshal(b1, &subAgg); err != nil {
				return err
			}

			bucket.SubAggregations[k] = subAgg
		}
	}

	return nil
}
