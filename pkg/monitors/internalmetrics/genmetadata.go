// Code generated by monitor-code-gen. DO NOT EDIT.

package internalmetrics

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "internal-metrics"

var groupSet = map[string]bool{}

const (
	sfxagentActiveMonitors             = "sfxagent.active_monitors"
	sfxagentActiveObservers            = "sfxagent.active_observers"
	sfxagentConfiguredMonitors         = "sfxagent.configured_monitors"
	sfxagentDatapointChannelLen        = "sfxagent.datapoint_channel_len"
	sfxagentDatapointRequestsActive    = "sfxagent.datapoint_requests_active"
	sfxagentDatapointsFailed           = "sfxagent.datapoints_failed"
	sfxagentDatapointsFiltered         = "sfxagent.datapoints_filtered"
	sfxagentDatapointsInFlight         = "sfxagent.datapoints_in_flight"
	sfxagentDatapointsReceived         = "sfxagent.datapoints_received"
	sfxagentDatapointsSent             = "sfxagent.datapoints_sent"
	sfxagentDatapointsWaiting          = "sfxagent.datapoints_waiting"
	sfxagentDimRequestSenders          = "sfxagent.dim_request_senders"
	sfxagentDimUpdatesCompleted        = "sfxagent.dim_updates_completed"
	sfxagentDimUpdatesCurrentlyDelayed = "sfxagent.dim_updates_currently_delayed"
	sfxagentDimUpdatesDropped          = "sfxagent.dim_updates_dropped"
	sfxagentDimUpdatesFailed           = "sfxagent.dim_updates_failed"
	sfxagentDimUpdatesFlappyTotal      = "sfxagent.dim_updates_flappy_total"
	sfxagentDimUpdatesStarted          = "sfxagent.dim_updates_started"
	sfxagentDiscoveredEndpoints        = "sfxagent.discovered_endpoints"
	sfxagentEventsBuffered             = "sfxagent.events_buffered"
	sfxagentEventsSent                 = "sfxagent.events_sent"
	sfxagentGoFrees                    = "sfxagent.go_frees"
	sfxagentGoHeapAlloc                = "sfxagent.go_heap_alloc"
	sfxagentGoHeapIdle                 = "sfxagent.go_heap_idle"
	sfxagentGoHeapInuse                = "sfxagent.go_heap_inuse"
	sfxagentGoHeapReleased             = "sfxagent.go_heap_released"
	sfxagentGoHeapSys                  = "sfxagent.go_heap_sys"
	sfxagentGoMallocs                  = "sfxagent.go_mallocs"
	sfxagentGoNextGc                   = "sfxagent.go_next_gc"
	sfxagentGoNumGc                    = "sfxagent.go_num_gc"
	sfxagentGoNumGoroutine             = "sfxagent.go_num_goroutine"
	sfxagentGoStackInuse               = "sfxagent.go_stack_inuse"
	sfxagentGoTotalAlloc               = "sfxagent.go_total_alloc"
)

var metricSet = map[string]monitors.MetricInfo{
	sfxagentActiveMonitors:             {Type: datapoint.Gauge},
	sfxagentActiveObservers:            {Type: datapoint.Gauge},
	sfxagentConfiguredMonitors:         {Type: datapoint.Gauge},
	sfxagentDatapointChannelLen:        {Type: datapoint.Gauge},
	sfxagentDatapointRequestsActive:    {Type: datapoint.Gauge},
	sfxagentDatapointsFailed:           {Type: datapoint.Counter},
	sfxagentDatapointsFiltered:         {Type: datapoint.Counter},
	sfxagentDatapointsInFlight:         {Type: datapoint.Gauge},
	sfxagentDatapointsReceived:         {Type: datapoint.Counter},
	sfxagentDatapointsSent:             {Type: datapoint.Counter},
	sfxagentDatapointsWaiting:          {Type: datapoint.Gauge},
	sfxagentDimRequestSenders:          {Type: datapoint.Gauge},
	sfxagentDimUpdatesCompleted:        {Type: datapoint.Counter},
	sfxagentDimUpdatesCurrentlyDelayed: {Type: datapoint.Gauge},
	sfxagentDimUpdatesDropped:          {Type: datapoint.Counter},
	sfxagentDimUpdatesFailed:           {Type: datapoint.Counter},
	sfxagentDimUpdatesFlappyTotal:      {Type: datapoint.Counter},
	sfxagentDimUpdatesStarted:          {Type: datapoint.Counter},
	sfxagentDiscoveredEndpoints:        {Type: datapoint.Gauge},
	sfxagentEventsBuffered:             {Type: datapoint.Gauge},
	sfxagentEventsSent:                 {Type: datapoint.Counter},
	sfxagentGoFrees:                    {Type: datapoint.Counter},
	sfxagentGoHeapAlloc:                {Type: datapoint.Gauge},
	sfxagentGoHeapIdle:                 {Type: datapoint.Gauge},
	sfxagentGoHeapInuse:                {Type: datapoint.Gauge},
	sfxagentGoHeapReleased:             {Type: datapoint.Gauge},
	sfxagentGoHeapSys:                  {Type: datapoint.Gauge},
	sfxagentGoMallocs:                  {Type: datapoint.Counter},
	sfxagentGoNextGc:                   {Type: datapoint.Gauge},
	sfxagentGoNumGc:                    {Type: datapoint.Gauge},
	sfxagentGoNumGoroutine:             {Type: datapoint.Gauge},
	sfxagentGoStackInuse:               {Type: datapoint.Gauge},
	sfxagentGoTotalAlloc:               {Type: datapoint.Counter},
}

var defaultMetrics = map[string]bool{
	sfxagentActiveMonitors:             true,
	sfxagentActiveObservers:            true,
	sfxagentConfiguredMonitors:         true,
	sfxagentDatapointChannelLen:        true,
	sfxagentDatapointRequestsActive:    true,
	sfxagentDatapointsFailed:           true,
	sfxagentDatapointsFiltered:         true,
	sfxagentDatapointsInFlight:         true,
	sfxagentDatapointsReceived:         true,
	sfxagentDatapointsSent:             true,
	sfxagentDatapointsWaiting:          true,
	sfxagentDimRequestSenders:          true,
	sfxagentDimUpdatesCompleted:        true,
	sfxagentDimUpdatesCurrentlyDelayed: true,
	sfxagentDimUpdatesDropped:          true,
	sfxagentDimUpdatesFailed:           true,
	sfxagentDimUpdatesFlappyTotal:      true,
	sfxagentDimUpdatesStarted:          true,
	sfxagentDiscoveredEndpoints:        true,
	sfxagentEventsBuffered:             true,
	sfxagentEventsSent:                 true,
	sfxagentGoFrees:                    true,
	sfxagentGoHeapAlloc:                true,
	sfxagentGoHeapIdle:                 true,
	sfxagentGoHeapInuse:                true,
	sfxagentGoHeapReleased:             true,
	sfxagentGoHeapSys:                  true,
	sfxagentGoMallocs:                  true,
	sfxagentGoNextGc:                   true,
	sfxagentGoNumGc:                    true,
	sfxagentGoNumGoroutine:             true,
	sfxagentGoStackInuse:               true,
	sfxagentGoTotalAlloc:               true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "internal-metrics",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         true,
}
