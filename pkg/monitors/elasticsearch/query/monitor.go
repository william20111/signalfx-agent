package query

import (
	"context"
	"encoding/json"
	"time"

	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/core/common/httpclient"
	"github.com/signalfx/signalfx-agent/pkg/core/config"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
	"github.com/signalfx/signalfx-agent/pkg/monitors/types"
	"github.com/signalfx/signalfx-agent/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// Config for this monitor
type Config struct {
	config.MonitorConfig  `yaml:",inline" acceptsEndpoints:"true"`
	httpclient.HTTPConfig `yaml:",inline"`

	Host string `yaml:"host" validate:"required"`
	Port string `yaml:"port" validate:"required"`

	MetricQueries []*metricQuery `yaml:"metricQueries" validate:"required"`
}

// Conceptually the data field in search queries define  bucket aggregations
// that determine "buckets" (some group of elasticsearch documents that match
// the search query) and metric aggregations that compute metrics from each "bucket"
type metricQuery struct {
	// Index that's being queried
	Index string `yaml:"index" validate:"required"`
	// Lucene search query which the elasticsearch documents should match
	LuceneSearchQuery string `yaml:"luceneSearchQuery" validate:"required"`
	// Query Interval
	QueryInterval time.Duration `yaml:"queryInterval" default:"30s"`
	// HTTP Post data field for aggregations
	HTTPPostData string `yaml:"httpPostData" validate:"required"`
}

// Monitor for ES queries
type Monitor struct {
	Output types.FilteringOutput
	cancel context.CancelFunc
	ctx    context.Context
	logger *utils.ThrottledLogger
}

func init() {
	monitors.Register(&monitorMetadata, func() interface{} { return &Monitor{} }, &Config{})
}

type metricQueryInfo struct {
	interval float64
	aggsMeta map[string]*AggregationMeta
}

// Configure monitor
func (m *Monitor) Configure(c *Config) error {
	m.logger = utils.NewThrottledLogger(log.WithFields(log.Fields{"monitorType": monitorType}), 20*time.Second)

	httpClient, err := c.HTTPConfig.Build()
	if err != nil {
		return err
	}
	esClient := NewESQueryClient(c.Host, c.Port, c.HTTPConfig.Scheme(), httpClient)
	m.ctx, m.cancel = context.WithCancel(context.Background())

	metricQueryInfos := make([]*metricQueryInfo, len(c.MetricQueries))

	for i, mq := range c.MetricQueries {
		var reqBody HTTPPostData
		err = json.Unmarshal([]byte(mq.HTTPPostData), &reqBody)
		if err != nil {
			return err
		}

		aggsMeta, err := reqBody.getAggregationsMeta()
		if err != nil {
			return err
		}

		metricQueryInfos[i] = &metricQueryInfo{
			interval: mq.QueryInterval.Seconds(),
			aggsMeta: aggsMeta,
		}
	}

	// Since metric queries can have their own interval, set them running on separate go routines
	for i, mq := range c.MetricQueries {
		index := i
		query := mq
		utils.RunOnInterval(m.ctx, func() {
			body, err := esClient.makeHTTPRequestFromConfig(query)

			if err != nil {
				log.Errorf("Failed to make HTTP request: %s", err)
				return
			}

			var resBody HTTPResponse
			if err := json.Unmarshal(body, &resBody); err != nil {
				log.Errorf("Error processing HTTP response: %s", err)
				return
			}

			dps := collectDatapoints(resBody, metricQueryInfos[index].aggsMeta, map[string]string{
				"index": c.MetricQueries[index].Index,
			})

			m.sendDatapoints(dps)
		}, time.Duration(metricQueryInfos[i].interval)*time.Second)
	}
	return nil
}

func (m *Monitor) sendDatapoints(dps []*datapoint.Datapoint) {
	// This is the filtering in place trick from https://github.com/golang/go/wiki/SliceTricks#filter-in-place
	n := 0
	for i := range dps {
		if dps[i] == nil {
			continue
		}
		dps[n] = dps[i]
		n++
	}
	m.Output.SendDatapoints(dps[:n]...)
}

// Shutdown stops the metric sync
func (m *Monitor) Shutdown() {
	if m.cancel != nil {
		m.cancel()
	}
}
