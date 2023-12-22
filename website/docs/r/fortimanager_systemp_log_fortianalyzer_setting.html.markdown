---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_log_fortianalyzer_setting"
description: |-
  Global FortiAnalyzer settings.
---

# fortimanager_systemp_log_fortianalyzer_setting
Global FortiAnalyzer settings.

## Example Usage

```hcl
resource "fortimanager_systemp_log_fortianalyzer_setting" "trname" {
  devprof             = "default"
  conn_timeout        = 60
  fallback_to_primary = "disable"
  max_log_rate        = 12
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `access_config` - Enable/disable FortiAnalyzer access to configuration and data. Valid values: `disable`, `enable`.

* `alt_server` - Alt-Server.
* `certificate` - Certificate used to communicate with FortiAnalyzer.
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `disable`, `enable`.

* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `default`, `high`, `low`, `disable`, `high-medium`, `low-medium`.

* `fallback_to_primary` - Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `disable`, `enable`.

* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ips_archive` - Enable/disable IPS packet archive logging. Valid values: `disable`, `enable`.

* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `preshared_key` - Preshared-key used for auto-authorization on FortiAnalyzer.
* `priority` - Set log transmission priority. Valid values: `low`, `default`.

* `reliable` - Enable/disable reliable logging to FortiAnalyzer. Valid values: `disable`, `enable`.

* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `upload_day` - Day of week (month) to upload logs.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.

* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.

* `upload_time` - Time to upload logs (hh:mm).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp LogFortianalyzerSetting can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_log_fortianalyzer_setting.labelname SystempLogFortianalyzerSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
