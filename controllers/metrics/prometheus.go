package metrics

import (
	datadoghqv1alpha1 "github.com/DataDog/watermarkpodautoscaler/api/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
	sigmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

func registerPrometheus() {
	Value.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "value",
			Help:      "Gauge of the value used for autoscaling",
		},
		[]string{
			WpaNamePromLabel,
			MetricNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	Highwm.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "high_watermak",
			Help:      "Gauge for the high watermark of a given WPA",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
			MetricNamePromLabel,
		})
	HighwmV2.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "high_watermark",
			Help:      "Gauge for the high watermark of a given WPA",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
			MetricNamePromLabel,
		})
	TransitionCountdown.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "transition_countdown",
			Help:      "Gauge indicating the time in seconds before scaling is authorized",
		},
		[]string{
			WpaNamePromLabel,
			TransitionPromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	Lowwm.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "low_watermak",
			Help:      "Gauge for the low watermark of a given WPA",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
			MetricNamePromLabel,
		})
	LowwmV2.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "low_watermark",
			Help:      "Gauge for the low watermark of a given WPA",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
			MetricNamePromLabel,
		})
	ReplicaProposal.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "replicas_scaling_proposal",
			Help:      "Gauge for the number of replicas the WPA will suggest to scale to",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
			MetricNamePromLabel,
		})
	ReplicaEffective.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "replicas_scaling_effective",
			Help:      "Gauge for the number of replicas the WPA will instruct to scale to",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	RestrictedScaling.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "restricted_scaling",
			Help:      "Gauge indicating whether the metric is within the watermarks bounds",
		},
		[]string{
			WpaNamePromLabel,
			ReasonPromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	ReplicaMin.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "min_replicas",
			Help:      "Gauge for the minReplicas value of a given WPA",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	ReplicaMax.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "max_replicas",
			Help:      "Gauge for the maxReplicas value of a given WPA",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	DryRun.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "dry_run",
			Help:      "Gauge reflecting the WPA dry-run status",
		},
		[]string{
			WpaNamePromLabel,
			ResourceNamespacePromLabel,
			ResourceNamePromLabel,
			ResourceKindPromLabel,
		})
	LabelsInfo.prom = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: config.namespace,
			Subsystem: Subsystem,
			Name:      "labels_info",
			Help:      "Info metric for additional labels to associate to metrics as tags",
		},
		append(ExtraPromLabels, WpaNamePromLabel, ResourceNamespacePromLabel),
	)

	sigmetrics.Registry.MustRegister(Value.prom)
	sigmetrics.Registry.MustRegister(Highwm.prom)
	sigmetrics.Registry.MustRegister(HighwmV2.prom)
	sigmetrics.Registry.MustRegister(Lowwm.prom)
	sigmetrics.Registry.MustRegister(LowwmV2.prom)
	sigmetrics.Registry.MustRegister(ReplicaProposal.prom)
	sigmetrics.Registry.MustRegister(ReplicaEffective.prom)
	sigmetrics.Registry.MustRegister(RestrictedScaling.prom)
	sigmetrics.Registry.MustRegister(TransitionCountdown.prom)
	sigmetrics.Registry.MustRegister(ReplicaMin.prom)
	sigmetrics.Registry.MustRegister(ReplicaMax.prom)
	sigmetrics.Registry.MustRegister(DryRun.prom)
	sigmetrics.Registry.MustRegister(LabelsInfo.prom)
}

func cleanupAssociatedPrometheusMetrics(wpa *datadoghqv1alpha1.WatermarkPodAutoscaler, onlyMetricsSpecific bool) {
	promLabelsForWpa := prometheus.Labels{
		WpaNamePromLabel:           wpa.Name,
		ResourceNamespacePromLabel: wpa.Namespace,
		ResourceNamePromLabel:      wpa.Spec.ScaleTargetRef.Name,
		ResourceKindPromLabel:      wpa.Spec.ScaleTargetRef.Kind,
	}

	if !onlyMetricsSpecific {
		ReplicaEffective.prom.Delete(promLabelsForWpa)
		ReplicaMin.prom.Delete(promLabelsForWpa)
		ReplicaMax.prom.Delete(promLabelsForWpa)

		for _, reason := range reasonValues {
			promLabelsForWpa[ReasonPromLabel] = reason
			RestrictedScaling.prom.Delete(promLabelsForWpa)
		}
		delete(promLabelsForWpa, ReasonPromLabel)

		promLabelsForWpa[TransitionPromLabel] = "downscale"
		TransitionCountdown.prom.Delete(promLabelsForWpa)
		promLabelsForWpa[TransitionPromLabel] = "upscale"
		TransitionCountdown.prom.Delete(promLabelsForWpa)
		delete(promLabelsForWpa, TransitionPromLabel)

		promLabelsInfo := prometheus.Labels{WpaNamePromLabel: wpa.Name, ResourceNamespacePromLabel: wpa.Namespace}
		for _, eLabel := range ExtraPromLabels {
			eLabelValue := wpa.Labels[eLabel]
			promLabelsInfo[eLabel] = eLabelValue
		}
		LabelsInfo.prom.Delete(promLabelsInfo)
		DryRun.prom.Delete(promLabelsForWpa)
	}

	for _, metricSpec := range wpa.Spec.Metrics {
		if metricSpec.Type == datadoghqv1alpha1.ResourceMetricSourceType {
			promLabelsForWpa[MetricNamePromLabel] = string(metricSpec.Resource.Name)
		} else {
			promLabelsForWpa[MetricNamePromLabel] = metricSpec.External.MetricName
		}

		Lowwm.prom.Delete(promLabelsForWpa)
		LowwmV2.prom.Delete(promLabelsForWpa)
		ReplicaProposal.prom.Delete(promLabelsForWpa)
		Highwm.prom.Delete(promLabelsForWpa)
		HighwmV2.prom.Delete(promLabelsForWpa)
		Value.prom.Delete(promLabelsForWpa)
	}
}
