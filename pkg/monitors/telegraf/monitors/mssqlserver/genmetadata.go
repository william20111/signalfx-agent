// Code generated by monitor-code-gen. DO NOT EDIT.

package mssqlserver

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "telegraf/sqlserver"

var groupSet = map[string]bool{}

const (
	sqlserverDatabaseIoReadBytes                      = "sqlserver_database_io.read_bytes"
	sqlserverDatabaseIoReadLatencyMs                  = "sqlserver_database_io.read_latency_ms"
	sqlserverDatabaseIoReads                          = "sqlserver_database_io.reads"
	sqlserverDatabaseIoWriteBytes                     = "sqlserver_database_io.write_bytes"
	sqlserverDatabaseIoWriteLatencyMs                 = "sqlserver_database_io.write_latency_ms"
	sqlserverDatabaseIoWrites                         = "sqlserver_database_io.writes"
	sqlserverMemoryClerksSizeKbBoundTrees             = "sqlserver_memory_clerks.size_kb.bound_trees"
	sqlserverMemoryClerksSizeKbBufferPool             = "sqlserver_memory_clerks.size_kb.buffer_pool"
	sqlserverMemoryClerksSizeKbConnectionPool         = "sqlserver_memory_clerks.size_kb.connection_pool"
	sqlserverMemoryClerksSizeKbGeneral                = "sqlserver_memory_clerks.size_kb.general"
	sqlserverMemoryClerksSizeKbInMemoryOltp           = "sqlserver_memory_clerks.size_kb.in-memory_oltp"
	sqlserverMemoryClerksSizeKbLogPool                = "sqlserver_memory_clerks.size_kb.log_pool"
	sqlserverMemoryClerksSizeKbMemoryclerkSqltrace    = "sqlserver_memory_clerks.size_kb.memoryclerk_sqltrace"
	sqlserverMemoryClerksSizeKbSchemaManagerUserStore = "sqlserver_memory_clerks.size_kb.schema_manager_user_store"
	sqlserverMemoryClerksSizeKbSosNode                = "sqlserver_memory_clerks.size_kb.sos_node"
	sqlserverMemoryClerksSizeKbSQLOptimizer           = "sqlserver_memory_clerks.size_kb.sql_optimizer"
	sqlserverMemoryClerksSizeKbSQLPlans               = "sqlserver_memory_clerks.size_kb.sql_plans"
	sqlserverMemoryClerksSizeKbSQLReservations        = "sqlserver_memory_clerks.size_kb.sql_reservations"
	sqlserverMemoryClerksSizeKbSQLStorageEngine       = "sqlserver_memory_clerks.size_kb.sql_storage_engine"
	sqlserverMemoryClerksSizeKbSystemRowsetStore      = "sqlserver_memory_clerks.size_kb.system_rowset_store"
	sqlserverPerformanceActiveMemoryGrantAmountKb     = "sqlserver_performance.active_memory_grant_amount_kb"
	sqlserverPerformanceActiveTempTables              = "sqlserver_performance.active_temp_tables"
	sqlserverPerformanceBackgroundWriterPagesSec      = "sqlserver_performance.background_writer_pages_sec"
	sqlserverPerformanceBackupRestoreThroughputSec    = "sqlserver_performance.backup_restore_throughput_sec"
	sqlserverPerformanceBatchRequestsSec              = "sqlserver_performance.batch_requests_sec"
	sqlserverPerformanceBlockedTasks                  = "sqlserver_performance.blocked_tasks"
	sqlserverPerformanceBufferCacheHitRatio           = "sqlserver_performance.buffer_cache_hit_ratio"
	sqlserverPerformanceBytesReceivedFromReplicaSec   = "sqlserver_performance.bytes_received_from_replica_sec"
	sqlserverPerformanceBytesSentToReplicaSec         = "sqlserver_performance.bytes_sent_to_replica_sec"
	sqlserverPerformanceBytesSentToTransportSec       = "sqlserver_performance.bytes_sent_to_transport_sec"
	sqlserverPerformanceCheckpointPagesSec            = "sqlserver_performance.checkpoint_pages_sec"
	sqlserverPerformanceCPULimitViolationCount        = "sqlserver_performance.cpu_limit_violation_count"
	sqlserverPerformanceCPUUsagePct                   = "sqlserver_performance.cpu_usage_pct"
	sqlserverPerformanceCPUUsageTime                  = "sqlserver_performance.cpu_usage_time"
	sqlserverPerformanceDataFileSizeKb                = "sqlserver_performance.data_file_size_kb"
	sqlserverPerformanceDiskReadBytesSec              = "sqlserver_performance.disk_read_bytes_sec"
	sqlserverPerformanceDiskReadIoSec                 = "sqlserver_performance.disk_read_io_sec"
	sqlserverPerformanceDiskReadIoThrottledSec        = "sqlserver_performance.disk_read_io_throttled_sec"
	sqlserverPerformanceDiskWriteBytesSec             = "sqlserver_performance.disk_write_bytes_sec"
	sqlserverPerformanceDiskWriteIoSec                = "sqlserver_performance.disk_write_io_sec"
	sqlserverPerformanceDiskWriteIoThrottledSec       = "sqlserver_performance.disk_write_io_throttled_sec"
	sqlserverPerformanceErrorsSec                     = "sqlserver_performance.errors_sec"
	sqlserverPerformanceFlowControlSec                = "sqlserver_performance.flow_control_sec"
	sqlserverPerformanceFlowControlTimeMsSec          = "sqlserver_performance.flow_control_time_ms_sec"
	sqlserverPerformanceForwardedRecordsSec           = "sqlserver_performance.forwarded_records_sec"
	sqlserverPerformanceFreeListStallsSec             = "sqlserver_performance.free_list_stalls_sec"
	sqlserverPerformanceFreeSpaceInTempdbKb           = "sqlserver_performance.free_space_in_tempdb_kb"
	sqlserverPerformanceFullScansSec                  = "sqlserver_performance.full_scans_sec"
	sqlserverPerformanceIndexSearchesSec              = "sqlserver_performance.index_searches_sec"
	sqlserverPerformanceLatchWaitsSec                 = "sqlserver_performance.latch_waits_sec"
	sqlserverPerformanceLazyWritesSec                 = "sqlserver_performance.lazy_writes_sec"
	sqlserverPerformanceLockTimeoutsSec               = "sqlserver_performance.lock_timeouts_sec"
	sqlserverPerformanceLockWaitCount                 = "sqlserver_performance.lock_wait_count"
	sqlserverPerformanceLockWaitTime                  = "sqlserver_performance.lock_wait_time"
	sqlserverPerformanceLockWaitsSec                  = "sqlserver_performance.lock_waits_sec"
	sqlserverPerformanceLogApplyPendingQueue          = "sqlserver_performance.log_apply_pending_queue"
	sqlserverPerformanceLogApplyReadyQueue            = "sqlserver_performance.log_apply_ready_queue"
	sqlserverPerformanceLogBytesFlushedSec            = "sqlserver_performance.log_bytes_flushed_sec"
	sqlserverPerformanceLogBytesReceivedSec           = "sqlserver_performance.log_bytes_received_sec"
	sqlserverPerformanceLogFileSizeKb                 = "sqlserver_performance.log_file_size_kb"
	sqlserverPerformanceLogFileUsedSizeKb             = "sqlserver_performance.log_file_used_size_kb"
	sqlserverPerformanceLogFlushWaitTime              = "sqlserver_performance.log_flush_wait_time"
	sqlserverPerformanceLogFlushesSec                 = "sqlserver_performance.log_flushes_sec"
	sqlserverPerformanceLogSendQueue                  = "sqlserver_performance.log_send_queue"
	sqlserverPerformanceLoginsSec                     = "sqlserver_performance.logins_sec"
	sqlserverPerformanceLogoutsSec                    = "sqlserver_performance.logouts_sec"
	sqlserverPerformanceMemoryBrokerClerkSize         = "sqlserver_performance.memory_broker_clerk_size"
	sqlserverPerformanceMemoryGrantsOutstanding       = "sqlserver_performance.memory_grants_outstanding"
	sqlserverPerformanceMemoryGrantsPending           = "sqlserver_performance.memory_grants_pending"
	sqlserverPerformanceNumberOfDeadlocksSec          = "sqlserver_performance.number_of_deadlocks_sec"
	sqlserverPerformancePageLifeExpectancy            = "sqlserver_performance.page_life_expectancy"
	sqlserverPerformancePageLookupsSec                = "sqlserver_performance.page_lookups_sec"
	sqlserverPerformancePageReadsSec                  = "sqlserver_performance.page_reads_sec"
	sqlserverPerformancePageSplitsSec                 = "sqlserver_performance.page_splits_sec"
	sqlserverPerformancePageWritesSec                 = "sqlserver_performance.page_writes_sec"
	sqlserverPerformancePctLogUsed                    = "sqlserver_performance.pct_log_used"
	sqlserverPerformanceProcessesBlocked              = "sqlserver_performance.processes_blocked"
	sqlserverPerformanceQuery                         = "sqlserver_performance.query"
	sqlserverPerformanceQueuedRequestCount            = "sqlserver_performance.queued_request_count"
	sqlserverPerformanceQueuedRequests                = "sqlserver_performance.queued_requests"
	sqlserverPerformanceReadaheadPagesSec             = "sqlserver_performance.readahead_pages_sec"
	sqlserverPerformanceReceivesFromReplicaSec        = "sqlserver_performance.receives_from_replica_sec"
	sqlserverPerformanceRecoveryQueue                 = "sqlserver_performance.recovery_queue"
	sqlserverPerformanceRedoneBytesSec                = "sqlserver_performance.redone_bytes_sec"
	sqlserverPerformanceReducedMemoryGrantCount       = "sqlserver_performance.reduced_memory_grant_count"
	sqlserverPerformanceRequestCount                  = "sqlserver_performance.request_count"
	sqlserverPerformanceRequestsCompletedSec          = "sqlserver_performance.requests_completed_sec"
	sqlserverPerformanceResentMessagesSec             = "sqlserver_performance.resent_messages_sec"
	sqlserverPerformanceSendsToReplicaSec             = "sqlserver_performance.sends_to_replica_sec"
	sqlserverPerformanceSendsToTransportSec           = "sqlserver_performance.sends_to_transport_sec"
	sqlserverPerformanceSQLCompilationsSec            = "sqlserver_performance.sql_compilations_sec"
	sqlserverPerformanceSQLReCompilationsSec          = "sqlserver_performance.sql_re-compilations_sec"
	sqlserverPerformanceTargetServerMemoryKb          = "sqlserver_performance.target_server_memory_kb"
	sqlserverPerformanceTempTablesCreationRate        = "sqlserver_performance.temp_tables_creation_rate"
	sqlserverPerformanceTempTablesForDestruction      = "sqlserver_performance.temp_tables_for_destruction"
	sqlserverPerformanceTotalServerMemoryKb           = "sqlserver_performance.total_server_memory_kb"
	sqlserverPerformanceTransactionDelay              = "sqlserver_performance.transaction_delay"
	sqlserverPerformanceTransactionsSec               = "sqlserver_performance.transactions_sec"
	sqlserverPerformanceUsedMemoryKb                  = "sqlserver_performance.used_memory_kb"
	sqlserverPerformanceUserConnections               = "sqlserver_performance.user_connections"
	sqlserverPerformanceVersionStoreSizeKb            = "sqlserver_performance.version_store_size_kb"
	sqlserverPerformanceWriteTransactionsSec          = "sqlserver_performance.write_transactions_sec"
	sqlserverPerformanceXtpMemoryUsedKb               = "sqlserver_performance.xtp_memory_used_kb"
	sqlserverServerPropertiesAvailableStorageMb       = "sqlserver_server_properties.available_storage_mb"
	sqlserverServerPropertiesCPUCount                 = "sqlserver_server_properties.cpu_count"
	sqlserverServerPropertiesDbOffline                = "sqlserver_server_properties.db_offline"
	sqlserverServerPropertiesDbOnline                 = "sqlserver_server_properties.db_online"
	sqlserverServerPropertiesDbRecovering             = "sqlserver_server_properties.db_recovering"
	sqlserverServerPropertiesDbRecoverypending        = "sqlserver_server_properties.db_recoverypending"
	sqlserverServerPropertiesDbRestoring              = "sqlserver_server_properties.db_restoring"
	sqlserverServerPropertiesDbSuspect                = "sqlserver_server_properties.db_suspect"
	sqlserverServerPropertiesEngineEdition            = "sqlserver_server_properties.engine_edition"
	sqlserverServerPropertiesServerMemory             = "sqlserver_server_properties.server_memory"
	sqlserverServerPropertiesTotalStorageMb           = "sqlserver_server_properties.total_storage_mb"
	sqlserverServerPropertiesUptime                   = "sqlserver_server_properties.uptime"
	sqlserverWaitstatsMaxWaitTimeMs                   = "sqlserver_waitstats.max_wait_time_ms"
	sqlserverWaitstatsResourceWaitMs                  = "sqlserver_waitstats.resource_wait_ms"
	sqlserverWaitstatsSignalWaitTimeMs                = "sqlserver_waitstats.signal_wait_time_ms"
	sqlserverWaitstatsWaitTimeMs                      = "sqlserver_waitstats.wait_time_ms"
	sqlserverWaitstatsWaitingTasksCount               = "sqlserver_waitstats.waiting_tasks_count"
)

var metricSet = map[string]monitors.MetricInfo{
	sqlserverDatabaseIoReadBytes:                      {Type: datapoint.Gauge},
	sqlserverDatabaseIoReadLatencyMs:                  {Type: datapoint.Gauge},
	sqlserverDatabaseIoReads:                          {Type: datapoint.Gauge},
	sqlserverDatabaseIoWriteBytes:                     {Type: datapoint.Gauge},
	sqlserverDatabaseIoWriteLatencyMs:                 {Type: datapoint.Gauge},
	sqlserverDatabaseIoWrites:                         {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbBoundTrees:             {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbBufferPool:             {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbConnectionPool:         {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbGeneral:                {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbInMemoryOltp:           {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbLogPool:                {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbMemoryclerkSqltrace:    {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSchemaManagerUserStore: {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSosNode:                {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSQLOptimizer:           {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSQLPlans:               {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSQLReservations:        {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSQLStorageEngine:       {Type: datapoint.Gauge},
	sqlserverMemoryClerksSizeKbSystemRowsetStore:      {Type: datapoint.Gauge},
	sqlserverPerformanceActiveMemoryGrantAmountKb:     {Type: datapoint.Gauge},
	sqlserverPerformanceActiveTempTables:              {Type: datapoint.Gauge},
	sqlserverPerformanceBackgroundWriterPagesSec:      {Type: datapoint.Gauge},
	sqlserverPerformanceBackupRestoreThroughputSec:    {Type: datapoint.Gauge},
	sqlserverPerformanceBatchRequestsSec:              {Type: datapoint.Gauge},
	sqlserverPerformanceBlockedTasks:                  {Type: datapoint.Gauge},
	sqlserverPerformanceBufferCacheHitRatio:           {Type: datapoint.Gauge},
	sqlserverPerformanceBytesReceivedFromReplicaSec:   {Type: datapoint.Gauge},
	sqlserverPerformanceBytesSentToReplicaSec:         {Type: datapoint.Gauge},
	sqlserverPerformanceBytesSentToTransportSec:       {Type: datapoint.Gauge},
	sqlserverPerformanceCheckpointPagesSec:            {Type: datapoint.Gauge},
	sqlserverPerformanceCPULimitViolationCount:        {Type: datapoint.Gauge},
	sqlserverPerformanceCPUUsagePct:                   {Type: datapoint.Gauge},
	sqlserverPerformanceCPUUsageTime:                  {Type: datapoint.Gauge},
	sqlserverPerformanceDataFileSizeKb:                {Type: datapoint.Gauge},
	sqlserverPerformanceDiskReadBytesSec:              {Type: datapoint.Gauge},
	sqlserverPerformanceDiskReadIoSec:                 {Type: datapoint.Gauge},
	sqlserverPerformanceDiskReadIoThrottledSec:        {Type: datapoint.Gauge},
	sqlserverPerformanceDiskWriteBytesSec:             {Type: datapoint.Gauge},
	sqlserverPerformanceDiskWriteIoSec:                {Type: datapoint.Gauge},
	sqlserverPerformanceDiskWriteIoThrottledSec:       {Type: datapoint.Gauge},
	sqlserverPerformanceErrorsSec:                     {Type: datapoint.Gauge},
	sqlserverPerformanceFlowControlSec:                {Type: datapoint.Gauge},
	sqlserverPerformanceFlowControlTimeMsSec:          {Type: datapoint.Gauge},
	sqlserverPerformanceForwardedRecordsSec:           {Type: datapoint.Gauge},
	sqlserverPerformanceFreeListStallsSec:             {Type: datapoint.Gauge},
	sqlserverPerformanceFreeSpaceInTempdbKb:           {Type: datapoint.Gauge},
	sqlserverPerformanceFullScansSec:                  {Type: datapoint.Gauge},
	sqlserverPerformanceIndexSearchesSec:              {Type: datapoint.Gauge},
	sqlserverPerformanceLatchWaitsSec:                 {Type: datapoint.Gauge},
	sqlserverPerformanceLazyWritesSec:                 {Type: datapoint.Gauge},
	sqlserverPerformanceLockTimeoutsSec:               {Type: datapoint.Gauge},
	sqlserverPerformanceLockWaitCount:                 {Type: datapoint.Gauge},
	sqlserverPerformanceLockWaitTime:                  {Type: datapoint.Gauge},
	sqlserverPerformanceLockWaitsSec:                  {Type: datapoint.Gauge},
	sqlserverPerformanceLogApplyPendingQueue:          {Type: datapoint.Gauge},
	sqlserverPerformanceLogApplyReadyQueue:            {Type: datapoint.Gauge},
	sqlserverPerformanceLogBytesFlushedSec:            {Type: datapoint.Gauge},
	sqlserverPerformanceLogBytesReceivedSec:           {Type: datapoint.Gauge},
	sqlserverPerformanceLogFileSizeKb:                 {Type: datapoint.Gauge},
	sqlserverPerformanceLogFileUsedSizeKb:             {Type: datapoint.Gauge},
	sqlserverPerformanceLogFlushWaitTime:              {Type: datapoint.Gauge},
	sqlserverPerformanceLogFlushesSec:                 {Type: datapoint.Gauge},
	sqlserverPerformanceLogSendQueue:                  {Type: datapoint.Gauge},
	sqlserverPerformanceLoginsSec:                     {Type: datapoint.Gauge},
	sqlserverPerformanceLogoutsSec:                    {Type: datapoint.Gauge},
	sqlserverPerformanceMemoryBrokerClerkSize:         {Type: datapoint.Gauge},
	sqlserverPerformanceMemoryGrantsOutstanding:       {Type: datapoint.Gauge},
	sqlserverPerformanceMemoryGrantsPending:           {Type: datapoint.Gauge},
	sqlserverPerformanceNumberOfDeadlocksSec:          {Type: datapoint.Gauge},
	sqlserverPerformancePageLifeExpectancy:            {Type: datapoint.Gauge},
	sqlserverPerformancePageLookupsSec:                {Type: datapoint.Gauge},
	sqlserverPerformancePageReadsSec:                  {Type: datapoint.Gauge},
	sqlserverPerformancePageSplitsSec:                 {Type: datapoint.Gauge},
	sqlserverPerformancePageWritesSec:                 {Type: datapoint.Gauge},
	sqlserverPerformancePctLogUsed:                    {Type: datapoint.Gauge},
	sqlserverPerformanceProcessesBlocked:              {Type: datapoint.Gauge},
	sqlserverPerformanceQuery:                         {Type: datapoint.Gauge},
	sqlserverPerformanceQueuedRequestCount:            {Type: datapoint.Gauge},
	sqlserverPerformanceQueuedRequests:                {Type: datapoint.Gauge},
	sqlserverPerformanceReadaheadPagesSec:             {Type: datapoint.Gauge},
	sqlserverPerformanceReceivesFromReplicaSec:        {Type: datapoint.Gauge},
	sqlserverPerformanceRecoveryQueue:                 {Type: datapoint.Gauge},
	sqlserverPerformanceRedoneBytesSec:                {Type: datapoint.Gauge},
	sqlserverPerformanceReducedMemoryGrantCount:       {Type: datapoint.Gauge},
	sqlserverPerformanceRequestCount:                  {Type: datapoint.Gauge},
	sqlserverPerformanceRequestsCompletedSec:          {Type: datapoint.Gauge},
	sqlserverPerformanceResentMessagesSec:             {Type: datapoint.Gauge},
	sqlserverPerformanceSendsToReplicaSec:             {Type: datapoint.Gauge},
	sqlserverPerformanceSendsToTransportSec:           {Type: datapoint.Gauge},
	sqlserverPerformanceSQLCompilationsSec:            {Type: datapoint.Gauge},
	sqlserverPerformanceSQLReCompilationsSec:          {Type: datapoint.Gauge},
	sqlserverPerformanceTargetServerMemoryKb:          {Type: datapoint.Gauge},
	sqlserverPerformanceTempTablesCreationRate:        {Type: datapoint.Gauge},
	sqlserverPerformanceTempTablesForDestruction:      {Type: datapoint.Gauge},
	sqlserverPerformanceTotalServerMemoryKb:           {Type: datapoint.Gauge},
	sqlserverPerformanceTransactionDelay:              {Type: datapoint.Gauge},
	sqlserverPerformanceTransactionsSec:               {Type: datapoint.Gauge},
	sqlserverPerformanceUsedMemoryKb:                  {Type: datapoint.Gauge},
	sqlserverPerformanceUserConnections:               {Type: datapoint.Gauge},
	sqlserverPerformanceVersionStoreSizeKb:            {Type: datapoint.Gauge},
	sqlserverPerformanceWriteTransactionsSec:          {Type: datapoint.Gauge},
	sqlserverPerformanceXtpMemoryUsedKb:               {Type: datapoint.Gauge},
	sqlserverServerPropertiesAvailableStorageMb:       {Type: datapoint.Gauge},
	sqlserverServerPropertiesCPUCount:                 {Type: datapoint.Gauge},
	sqlserverServerPropertiesDbOffline:                {Type: datapoint.Gauge},
	sqlserverServerPropertiesDbOnline:                 {Type: datapoint.Gauge},
	sqlserverServerPropertiesDbRecovering:             {Type: datapoint.Gauge},
	sqlserverServerPropertiesDbRecoverypending:        {Type: datapoint.Gauge},
	sqlserverServerPropertiesDbRestoring:              {Type: datapoint.Gauge},
	sqlserverServerPropertiesDbSuspect:                {Type: datapoint.Gauge},
	sqlserverServerPropertiesEngineEdition:            {Type: datapoint.Gauge},
	sqlserverServerPropertiesServerMemory:             {Type: datapoint.Gauge},
	sqlserverServerPropertiesTotalStorageMb:           {Type: datapoint.Gauge},
	sqlserverServerPropertiesUptime:                   {Type: datapoint.Gauge},
	sqlserverWaitstatsMaxWaitTimeMs:                   {Type: datapoint.Gauge},
	sqlserverWaitstatsResourceWaitMs:                  {Type: datapoint.Gauge},
	sqlserverWaitstatsSignalWaitTimeMs:                {Type: datapoint.Gauge},
	sqlserverWaitstatsWaitTimeMs:                      {Type: datapoint.Gauge},
	sqlserverWaitstatsWaitingTasksCount:               {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	sqlserverDatabaseIoReadBytes:                true,
	sqlserverDatabaseIoReadLatencyMs:            true,
	sqlserverDatabaseIoReads:                    true,
	sqlserverDatabaseIoWriteBytes:               true,
	sqlserverDatabaseIoWriteLatencyMs:           true,
	sqlserverDatabaseIoWrites:                   true,
	sqlserverPerformanceBatchRequestsSec:        true,
	sqlserverPerformanceBufferCacheHitRatio:     true,
	sqlserverPerformanceDiskReadIoThrottledSec:  true,
	sqlserverPerformanceDiskWriteIoThrottledSec: true,
	sqlserverPerformanceLockWaitCount:           true,
	sqlserverPerformanceProcessesBlocked:        true,
	sqlserverPerformanceSQLCompilationsSec:      true,
	sqlserverPerformanceSQLReCompilationsSec:    true,
	sqlserverPerformanceUserConnections:         true,
	sqlserverServerPropertiesDbOffline:          true,
	sqlserverServerPropertiesDbOnline:           true,
	sqlserverServerPropertiesDbRecovering:       true,
	sqlserverServerPropertiesDbRecoverypending:  true,
	sqlserverServerPropertiesDbRestoring:        true,
	sqlserverServerPropertiesDbSuspect:          true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "telegraf/sqlserver",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
