// Code generated by monitor-code-gen. DO NOT EDIT.

package processes

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/processes"

var groupSet = map[string]bool{}

const (
	diskOctetsRead     = "disk_octets.read"
	diskOctetsWrite    = "disk_octets.write"
	forkRate           = "fork_rate"
	ioOctetsRx         = "io_octets.rx"
	ioOctetsTx         = "io_octets.tx"
	ioOpsRead          = "io_ops.read"
	ioOpsWrite         = "io_ops.write"
	psCode             = "ps_code"
	psCountProcesses   = "ps_count.processes"
	psCountThreads     = "ps_count.threads"
	psCputimeSyst      = "ps_cputime.syst"
	psCputimeUser      = "ps_cputime.user"
	psData             = "ps_data"
	psPagefaultsMajflt = "ps_pagefaults.majflt"
	psPagefaultsMinflt = "ps_pagefaults.minflt"
	psRss              = "ps_rss"
	psStacksize        = "ps_stacksize"
	psStateBlocked     = "ps_state.blocked"
	psStatePaging      = "ps_state.paging"
	psStateRunning     = "ps_state.running"
	psStateSleeping    = "ps_state.sleeping"
	psStateStopped     = "ps_state.stopped"
	psStateZombies     = "ps_state.zombies"
	psVM               = "ps_vm"
)

var metricSet = map[string]monitors.MetricInfo{
	diskOctetsRead:     {Type: datapoint.Counter},
	diskOctetsWrite:    {Type: datapoint.Counter},
	forkRate:           {Type: datapoint.Counter},
	ioOctetsRx:         {Type: datapoint.Counter},
	ioOctetsTx:         {Type: datapoint.Counter},
	ioOpsRead:          {Type: datapoint.Counter},
	ioOpsWrite:         {Type: datapoint.Counter},
	psCode:             {Type: datapoint.Gauge},
	psCountProcesses:   {Type: datapoint.Gauge},
	psCountThreads:     {Type: datapoint.Gauge},
	psCputimeSyst:      {Type: datapoint.Counter},
	psCputimeUser:      {Type: datapoint.Counter},
	psData:             {Type: datapoint.Gauge},
	psPagefaultsMajflt: {Type: datapoint.Counter},
	psPagefaultsMinflt: {Type: datapoint.Counter},
	psRss:              {Type: datapoint.Gauge},
	psStacksize:        {Type: datapoint.Gauge},
	psStateBlocked:     {Type: datapoint.Gauge},
	psStatePaging:      {Type: datapoint.Gauge},
	psStateRunning:     {Type: datapoint.Gauge},
	psStateSleeping:    {Type: datapoint.Gauge},
	psStateStopped:     {Type: datapoint.Gauge},
	psStateZombies:     {Type: datapoint.Gauge},
	psVM:               {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "collectd/processes",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         true,
}
