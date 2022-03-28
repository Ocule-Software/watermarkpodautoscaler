# Metrics

* **Default Tags**

`{wpa_name: string, wpa_namespace: string, resource_namespace: string}`

* **Metrics**

* `value{metric_name: string, resource_name: string, resource_kind: string}` (Gauge) value used for autoscaling
* `high_watermak{metric_name: string, resource_name: string, resource_kind: string}` (Gauge) High watermark of a given WPA
* `high_watermark{metric_name: string, resource_name: string, resource_kind: string}` (Gauge) High watermark of a given WPA
* `low_watermak{metric_name: string, resource_name: string, resource_kind: string}` (Gauge) Low watermark of a given WPA
* `low_watermark{metric_name: string, resource_name: string, resource_kind: string}` (Gauge) Low watermark of a given WPA
* `transition_countdown{transition: string, resource_name: string, resource_kind: string}` (Gauge) Value indicating time in seconds before scaling is authorized
* `replicas_scaling_proposal{metric_name: string, resource_name: string, resource_kind: string}` (Gauge) Number of replicas the WPA will suggest to scale to
* `replicas_scaling_effective{resource_name: string, resource_kind: string}` (Gauge) Number of replicas the WPA will instruct to scale to
* `restricted_scaling{reason: string, resource_name: string, resource_kind: string}` (Gauge) 0/1 value indicating whether the metric is within the watermarks bounds
* `min_replicas{resource_name: string, resource_kind: string}` (Gauge) WPA minReplicas yaml value
* `max_replicas{resource_name: string, resource_kind: string}` (Gauge) WPA maxReplicas yaml value
* `dry_run{resource_name: string, resource_kind: string}` (Gauge) 0/1 value indicating whether the WPA is in dry run mode
* `labels_info{resource_name: string, resource_kind: string, ...map[string]string}` (Gauge) Additional labels to associate to metrics as tags via env variable e.g `DD_LABELS_AS_TAGS=label1 label2`

## Configuration

* **Arguments**
* `--prometheus`: (Bool=true) Enable prometheus metric publishing via metric port
* `--metrics-addr`: (String=0.0.0.0:8383) The address the metric endpoint binds to
* `--statsd`: (Bool=false) Enable statsd metric publishing
* `--statsd-addr`: (String=127.0.0.1:8125) DogStatsD address, supports unix sockets in format (unix:///path/to/socket)
* `--metrics-namespace`: (String=watermarkpodautoscaler) The metric namespace / prefix to use