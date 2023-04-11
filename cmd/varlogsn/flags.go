package main

import (
	"github.com/urfave/cli/v2"

	"github.com/kakao/varlog/internal/flags"
	"github.com/kakao/varlog/internal/storage"
	"github.com/kakao/varlog/internal/storagenode"
	"github.com/kakao/varlog/pkg/util/units"
)

var (
	flagClusterID     = flags.ClusterID()
	flagStorageNodeID = flags.StorageNodeID()
	flagListen        = flags.FlagDesc{
		Name:    "listen",
		Aliases: []string{"listen-address"},
		Envs:    []string{"LISTEN", "LISTEN_ADDRESS"},
	}
	flagAdvertise = flags.FlagDesc{
		Name:    "advertise",
		Aliases: []string{"advertise-address"},
		Envs:    []string{"ADVERTISE", "ADVERTISE_ADDRESS"},
	}
	flagBallastSize = flags.FlagDesc{
		Name:  "ballast-size",
		Envs:  []string{"BALLAST_SIZE"},
		Usage: "B, KiB, MiB, GiB",
	}
	flagVolumes = flags.FlagDesc{
		Name:    "volumes",
		Aliases: []string{"volume"},
		Envs:    []string{"VOLUMES", "VOLUME"},
	}

	flagMaxLogStreamReplicasCount = &cli.IntFlag{
		Name:  "max-logstream-replicas-count",
		Usage: "The maximum number of log stream replicas in a storage node, infinity if a negative value",
		Value: storagenode.DefaultMaxLogStreamReplicasCount,
	}

	// flags for grpc options.
	flagServerReadBufferSize = flags.FlagDesc{
		Name:  "server-read-buffer-size",
		Envs:  []string{"SERVER_READ_BUFFER_SIZE"},
		Usage: "B, KiB, MiB, GiB",
	}
	flagServerWriteBufferSize = flags.FlagDesc{
		Name:  "server-write-buffer-size",
		Envs:  []string{"SERVER_WRITE_BUFFER_SIZE"},
		Usage: "B, KiB, MiB, GiB",
	}
	flagServerMaxRecvMsgSize = flags.FlagDesc{
		Name:    "server-max-msg-size",
		Aliases: []string{"server-max-message-size"},
		Envs:    []string{"SERVER_MAX_MSG_SIZE", "SERVER_MAX_MESSAGE_SIZE"},
		Usage:   "B, KiB, MiB, GiB",
	}
	flagReplicationClientReadBufferSize = flags.FlagDesc{
		Name:  "replication-client-read-buffer-size",
		Envs:  []string{"REPLICATION_CLIENT_READ_BUFFER_SIZE"},
		Usage: "B, KiB, MiB, GiB",
	}
	flagReplicationClientWriteBufferSize = flags.FlagDesc{
		Name:  "replication-client-write-buffer-size",
		Envs:  []string{"REPLICATION_CLIENT_WRITE_BUFFER_SIZE"},
		Usage: "B, KiB, MiB, GiB",
	}

	// flags for logstream executor options.
	flagLogStreamExecutorSequenceQueueCapacity = flags.FlagDesc{
		Name:    "logstream-executor-sequence-queue-capacity",
		Aliases: []string{"lse-sequence-queue-capacity"},
	}
	flagLogStreamExecutorWriteQueueCapacity = flags.FlagDesc{
		Name:    "logstream-executor-write-queue-capacity",
		Aliases: []string{"lse-write-queue-capacity"},
	}
	flagLogStreamExecutorCommitQueueCapacity = flags.FlagDesc{
		Name:    "logstream-executor-commit-queue-capacity",
		Aliases: []string{"lse-commit-queue-capacity"},
	}
	flagLogStreamExecutorReplicateclientQueueCapacity = flags.FlagDesc{
		Name:    "logstream-executor-replicate-client-queue-capacity",
		Aliases: []string{"lse-replicate-client-queue-capacity"},
	}

	// flags for storage.
	flagStorageDisableWAL = flags.FlagDesc{
		Name: "storage-disable-wal",
		Envs: []string{"STORAGE_DISABLE_WAL"},
	}
	flagStorageNoSync = flags.FlagDesc{
		Name: "storage-no-sync",
		Envs: []string{"STORAGE_NO_SYNC"},
	}
	flagStorageL0CompactionThreshold = flags.FlagDesc{
		Name: "storage-l0-compaction-threshold",
		Envs: []string{"STORAGE_L0_COMPACTION_THRESHOLD"},
	}
	flagStorageL0StopWritesThreshold = flags.FlagDesc{
		Name: "storage-l0-stop-writes-threshold",
		Envs: []string{"STORAGE_L0_STOP_WRITES_THRESHOLD"},
	}
	flagStorageLBaseMaxBytes = flags.FlagDesc{
		Name: "storage-lbase-max-bytes",
		Envs: []string{"STORAGE_LBASE_MAX_BYTES"},
	}
	flagStorageMaxOpenFiles = flags.FlagDesc{
		Name: "storage-max-open-files",
		Envs: []string{"STORAGE_MAX_OPEN_FILES"},
	}
	flagStorageMemTableSize = &cli.StringFlag{
		Name:    "storage-mem-table-size",
		Aliases: []string{"storage-memtable-size"},
		EnvVars: []string{"STORAGE_MEM_TABLE_SIZE", "STORAGE_MEMTABLE_SIZE"},
		Value:   units.ToByteSizeString(storage.DefaultMemTableSize),
	}
	flagStorageMemTableStopWritesThreshold = flags.FlagDesc{
		Name:    "storage-mem-table-stop-writes-threshold",
		Aliases: []string{"storage-memtable-stop-writes-threshold"},
		Envs:    []string{"STORAGE_MEM_TABLE_STOP_WRITES_THRESHOLD", "STORAGE_MEMTABLE_STOP_WRITES_THRESHOLD"},
	}
	flagStorageMaxConcurrentCompaction = flags.FlagDesc{
		Name: "storage-max-concurrent-compaction",
		Envs: []string{"STORAGE_MAX_CONCURRENT_COMPACTION"},
	}
	flagStorageMetricsLogInterval = &cli.DurationFlag{
		Name:    "storage-metrics-log-interval",
		EnvVars: []string{"STORAGE_METRICS_LOG_INTERVAL"},
		Value:   storage.DefaultMetricsLogInterval,
	}
	flagStorageVerbose = flags.FlagDesc{
		Name: "storage-verbose",
		Envs: []string{"STORAGE_VERBOSE"},
	}

	// flags for logging.
	flagLogDir = flags.FlagDesc{
		Name:    "log-dir",
		Aliases: []string{"logdir"},
		Envs:    []string{"LOG_DIR", "LOGDIR"},
	}
	flagLogToStderr = flags.FlagDesc{
		Name: "logtostderr",
		Envs: []string{"LOGTOSTDERR"},
	}
	flagLogFileRetentionDays = flags.FlagDesc{
		Name:    "logfile-retention-days",
		Aliases: []string{"log-file-retention-days"},
		Envs:    []string{"LOGFILE_RETENTION_DAYS", "LOG_FILE_RETENTION_DAYS"},
	}
	flagLogFileCompression = flags.FlagDesc{
		Name:    "logfile-compression",
		Aliases: []string{"log-file-compression"},
		Envs:    []string{"LOGFILE_COMPRESSION", "LOG_FILE_COMPRESSION"},
	}
	flagLogLevel = flags.FlagDesc{
		Name:    "loglevel",
		Aliases: []string{"log-level"},
		Envs:    []string{"LOGLEVEL", "LOG_LEVEL"},
	}

	// flags for telemetry.
	flagExporterType = flags.FlagDesc{
		Name:  "exporter-type",
		Usage: "exporter type: stdout, otlp or noop",
		Envs:  []string{"EXPORTER_TYPE"},
	}
	flagExporterStopTimeout = flags.FlagDesc{
		Name:  "expoter-stop-timeout",
		Usage: "timeout for stopping exporter",
		Envs:  []string{"EXPORTER_STOP_TIMEOUT"},
	}
	flagStdoutExporterPrettyPrint = flags.FlagDesc{
		Name:  "exporter-pretty-print",
		Usage: "pretty print when using stdout exporter",
		Envs:  []string{"EXPORTER_PRETTY_PRINT"},
	}
	flagOTLPExporterInsecure = flags.FlagDesc{
		Name:  "exporter-otlp-insecure",
		Usage: "disable client transport security for the OTLP exporter",
		Envs:  []string{"EXPORTER_OTLP_INSECURE"},
	}
	flagOTLPExporterEndpoint = flags.FlagDesc{
		Name:  "exporter-otlp-endpoint",
		Usage: "the endpoint that exporter connects",
		Envs:  []string{"EXPORTER_OTLP_ENDPOINT"},
	}
)
