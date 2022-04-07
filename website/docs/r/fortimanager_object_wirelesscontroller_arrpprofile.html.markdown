---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_arrpprofile"
description: |-
  Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.
---

# fortimanager_object_wirelesscontroller_arrpprofile
Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
* `darrp_optimize_schedules` - Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space.
* `include_dfs_channel` - Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable). Valid values: `no`, `disable`, `yes`, `enable`.

* `include_weather_channel` - Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable). Valid values: `no`, `disable`, `yes`, `enable`.

* `monitor_period` - Period in seconds to measure average transmit retries and receive errors (default = 300).
* `name` - WiFi ARRP profile name.
* `override_darrp_optimize` - Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable). Valid values: `disable`, `enable`.

* `selection_period` - Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).
* `threshold_ap` - Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).
* `threshold_channel_load` - Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).
* `threshold_noise_floor` - Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).
* `threshold_rx_errors` - Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).
* `threshold_spectral_rssi` - Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).
* `threshold_tx_retries` - Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).
* `weight_channel_load` - Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).
* `weight_dfs_channel` - Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).
* `weight_managed_ap` - Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).
* `weight_noise_floor` - Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).
* `weight_rogue_ap` - Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).
* `weight_spectral_rssi` - Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).
* `weight_weather_channel` - Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController ArrpProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_arrpprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
