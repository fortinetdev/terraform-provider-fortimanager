---
subcategory: "Object Telemetry Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_telemetrycontroller_profile"
description: |-
  Configure FortiTelemetry profiles.
---

# fortimanager_object_telemetrycontroller_profile
Configure FortiTelemetry profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `application`: `fortimanager_object_telemetrycontroller_profile_application`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `application` - Application. The structure of `application` block is documented below.
* `comment` - Comment.
* `name` - Name of the profile.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `application` block supports:

* `app_name` - Application name.
* `app_throughput` - Application throughput in megabytes (0 - 10,000, default = 2 MB).
* `atdt_threshold` - Threshold of application total downloading time in milliseconds (0 - 10,000,000, default = 20,000 ms).
* `dns_time_threshold` - Threshold of DNS resolution time in milliseconds (0 - 10,000,000, default = 20 ms).
* `experience_score_threshold` - Threshold of experience score (0 - 10, default = 8).
* `failure_rate_threshold` - Threshold of failure rate (0 - 100, default = 5 percentage).
* `id` - ID.
* `interval` - Time in milliseconds to check the application (1000 - 86,400 * 1000, default = 300 * 1000 ms).
* `jitter_threshold` - Threshold of jitter in milliseconds (0 - 10,000,000, default = 20 ms).
* `latency_threshold` - Threshold of latency in milliseconds (0 - 10,000,000, default = 20 ms).
* `monitor` - Enable/disable monitoring of the application. Valid values: `disable`, `enable`.

* `packet_loss_threshold` - Threshold of packet loss (0 - 100, default = 5 percentage).
* `sla` - Sla. The structure of `sla` block is documented below.
* `tcp_rtt_threshold` - Threshold of TCP round-trip time in milliseconds (0 - 10,000,000, default = 20 ms).
* `tls_time_threshold` - Threshold of TLS handshake time in milliseconds (0 - 10,000,000, default = 20 ms).
* `ttfb_threshold` - Threshold of time to first byte in milliseconds (0 - 10,000,000, default = 20 ms).

The `sla` block supports:

* `app_throughput_threshold` - Threshold of application throughput in megabytes (0 - 10,000, default = 2 MB).
* `atdt_threshold` - Threshold of application total downloading time in milliseconds (0 - 10,000,000, default = 3,000 ms).
* `dns_time_threshold` - Threshold of 95th percentile of DNS resolution time in milliseconds (0 - 10,000,000, default = 100 ms).
* `experience_score_threshold` - Threshold of experience score (0 - 10, default = 6).
* `failure_rate_threshold` - Threshold of failure rate (0 - 100, default = 5 percentage).
* `jitter_threshold` - Threshold of jitter in milliseconds (0 - 10,000,000, default = 50 ms).
* `latency_threshold` - Threshold of latency in milliseconds (0 - 10,000,000, default = 100 ms).
* `packet_loss_threshold` - Threshold of packet loss (0 - 100, default = 5 percentage).
* `sla_factor` - Criteria on which metric to SLA threshold list. Valid values: `latency`, `jitter`, `packet-loss`, `experience-score`, `failure-rate`, `ttfb`, `atdt`, `tcp-rtt`, `dns-time`, `tls-time`, `app-throughput`.

* `tcp_rtt_threshold` - Threshold of TCP round-trip time in milliseconds (0 - 10,000,000, default = 100 ms).
* `tls_time_threshold` - Threshold of 95th percentile of TLS handshake time in milliseconds (0 - 10,000,000, default = 200 ms).
* `ttfb_threshold` - Threshold of time to first byte in milliseconds (0 - 10,000,000, default = 200 ms).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectTelemetryController Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_telemetrycontroller_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
