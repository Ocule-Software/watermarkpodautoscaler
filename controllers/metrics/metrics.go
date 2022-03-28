package metrics

import (
	datadoghqv1alpha1 "github.com/DataDog/watermarkpodautoscaler/api/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
)

type gaugeMetric struct {
	name string
	prom *prometheus.GaugeVec
}

func promDo(m *gaugeMetric, fn func()) {
	if config.prometheus.enabled && m != nil {
		fn()
	}
}

func statsdDo(fn func() error) {
	if config.statsd.enabled && config.statsd.client != nil {
		if err := fn(); err != nil {
			config.logger.Error(err, "Error sending statsd metric")
		}
	}
}

func labelsToTags(labels prometheus.Labels) []string {
	tags := make([]string, 0)
	for k, v := range labels {
		tags = append(tags, k+":"+v)
	}
	return tags
}

// Delete deletes the given prometheus metric
func Delete(m *gaugeMetric, labels prometheus.Labels) {
	promDo(m, func() {
		m.prom.Delete(labels)
	})
}

// Set sets the value of the gauge metric
func Set(m *gaugeMetric, labels prometheus.Labels, val float64) {
	promDo(m, func() {
		m.prom.With(labels).Set(val)
	})

	statsdDo(func() error {
		return config.statsd.client.Gauge(m.name, val, labelsToTags(labels), 1)
	})
}

// CleanupAssociatedMetrics cleans up all prometheus metrics associated with the given WatermarkPodAutoscaler
func CleanupAssociatedMetrics(wpa *datadoghqv1alpha1.WatermarkPodAutoscaler, onlyMetricsSpecific bool) {
	if config.prometheus.enabled {
		cleanupAssociatedPrometheusMetrics(wpa, onlyMetricsSpecific)
	}
}
