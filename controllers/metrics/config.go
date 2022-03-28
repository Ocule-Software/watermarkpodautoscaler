package metrics

import (
	"os"
	"strings"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/go-logr/logr"
)

const (
	// Subsystem concatenated to beginning of metric name
	Subsystem = "controller"

	// WpaNamePromLabel is the label name for the WatermarkPodAutoscaler CRD object name
	WpaNamePromLabel = "wpa_name"

	// WpaNamespacePromLabel is the label name for the WatermarkPodAutoscaler CRD object namaespace
	WpaNamespacePromLabel = "wpa_namespace"

	// ResourceNamePromLabel is the label name for the WatermarkPodAutoscaler target resource name
	ResourceNamePromLabel = "resource_name"

	// ResourceKindPromLabel is the label name for the WatermarkPodAutoscaler target resource kind
	ResourceKindPromLabel = "resource_kind"

	// ResourceNamespacePromLabel is the label name for the WatermarkPodAutoscaler target resource namespace
	ResourceNamespacePromLabel = "resource_namespace"

	// MetricNamePromLabel is the label name for the WatermarkPodAutoscaler metric being checked
	MetricNamePromLabel = "metric_name"

	// ReasonPromLabel is the label name for the reason the WatermarkPodAutoscaler isn't being scaled
	ReasonPromLabel = "reason"

	// TransitionPromLabel is the label name for the transition / cooldown the WatermarkPodAutoscaler is currently in
	TransitionPromLabel = "transition"

	// DownscaleCappingPromLabelVal is the label value the ReasonPromLabel
	DownscaleCappingPromLabelVal = "downscale_capping"

	// UpscaleCappingPromLabelVal is the label value the ReasonPromLabel
	UpscaleCappingPromLabelVal = "upscale_capping"

	// WithinBoundsPromLabelVal is the label value the ReasonPromLabel
	WithinBoundsPromLabelVal = "within_bounds"
)

// reasonValues contains the 3 possible values of the 'reason' label
var reasonValues = []string{DownscaleCappingPromLabelVal, UpscaleCappingPromLabelVal, WithinBoundsPromLabelVal}

// ExtraPromLabels labels to add to an info metric and join on (with WpaNamePromLabel) in the Datadog prometheus check
var ExtraPromLabels = strings.Fields(os.Getenv("DD_LABELS_AS_TAGS"))

// OptsFunc is a function that can be passed to Init to modify the metricsConfig
type OptsFunc = func(*metricsConfig)

type statsdConfig struct {
	enabled bool
	host    string
	client  *statsd.Client
}

type prometheusConfig struct {
	enabled bool
}

type metricsConfig struct {
	// StatsD
	statsd *statsdConfig

	// Prometheus
	prometheus *prometheusConfig

	namespace string
	logger    logr.Logger
}

var config = &metricsConfig{
	statsd: &statsdConfig{
		enabled: false,
		host:    "127.0.0.1:8125",
	},
	prometheus: &prometheusConfig{
		enabled: false,
	},
	logger: logr.Discard(),
}

// Init initializes the metricsConfig with the given OptsFunc
func Init(opts ...OptsFunc) {
	for _, opt := range opts {
		opt(config)
	}
}

// WithStatsD enables StatsD in the metricsConfig with the given statsd host
func WithStatsD(host string) OptsFunc {
	return func(m *metricsConfig) {
		client, err := statsd.New(host, statsd.WithNamespace(m.namespace+"."+Subsystem))
		if err != nil {
			// TODO: Add namespace and default tag system
			m.logger.Error(err, "Failed to create statsd client, proceeding with it disabled...")
			m.statsd = &statsdConfig{
				enabled: false,
				host:    host,
				client:  nil,
			}
		} else {
			m.statsd = &statsdConfig{
				enabled: true,
				host:    host,
				client:  client,
			}
		}
	}
}

// WithPrometheus enables Prometheus in the metricsConfig
func WithPrometheus() OptsFunc {
	return func(m *metricsConfig) {
		registerPrometheus()
		m.prometheus = &prometheusConfig{
			enabled: true,
		}
	}
}

// WithNamespace sets the metrics namespace / prefix
func WithNamespace(namespace string) OptsFunc {
	return func(m *metricsConfig) {
		m.namespace = namespace
	}
}

// WithLogger sets the logger to be used by the metrics handlers for errors and warnings
func WithLogger(logger logr.Logger) OptsFunc {
	return func(m *metricsConfig) {
		m.logger = logger
	}
}

// * For more details on all metrics check the docs/metrics.md file
var (
	// Value gauge metric
	Value = &gaugeMetric{
		name: "value",
		prom: nil,
	}

	// Highwm gauge metric
	Highwm = &gaugeMetric{
		name: "high_watermak",
		prom: nil,
	}

	// HighwmV2 gauge metric
	HighwmV2 = &gaugeMetric{
		name: "high_watermark",
		prom: nil,
	}

	// TransitionCountdown gauge metric
	TransitionCountdown = &gaugeMetric{
		name: "transition_countdown",
		prom: nil,
	}

	// Lowwm gauge metric
	Lowwm = &gaugeMetric{
		name: "low_watermak",
		prom: nil,
	}

	// LowwmV2 gauge metric
	LowwmV2 = &gaugeMetric{
		name: "low_watermark",
		prom: nil,
	}

	// ReplicaProposal gauge metric
	ReplicaProposal = &gaugeMetric{
		name: "replicas_scaling_proposal",
		prom: nil,
	}

	// ReplicaEffective gauge metric
	ReplicaEffective = &gaugeMetric{
		name: "replicas_scaling_effective",
		prom: nil,
	}

	// RestrictedScaling gauge metric
	RestrictedScaling = &gaugeMetric{
		name: "restricted_scaling",
		prom: nil,
	}

	// ReplicaMin gauge metric
	ReplicaMin = &gaugeMetric{
		name: "min_replicas",
		prom: nil}

	// ReplicaMax gauge metric
	ReplicaMax = &gaugeMetric{
		name: "max_replicas",
		prom: nil}

	// DryRun gauge metric
	DryRun = &gaugeMetric{
		name: "dry_run",
		prom: nil}

	// LabelsInfo gauge metric
	LabelsInfo = &gaugeMetric{
		name: "labels_info",
		prom: nil,
	}
)
