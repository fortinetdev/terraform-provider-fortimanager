---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_dospolicy6_anomaly"
description: |-
  Anomaly name.
---

# fortimanager_packages_firewall_dospolicy6_anomaly
Anomaly name.

~> This resource is a sub resource for variable `anomaly` of resource `fortimanager_packages_firewall_dospolicy6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.
* `dos_policy6` - Dos Policy6.

* `action` - Action taken when the threshold is reached. Valid values: `pass`, `block`, `proxy`.

* `log` - Enable/disable anomaly logging. Valid values: `disable`, `enable`.

* `name` - Anomaly name.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`, `both`, `interface`.

* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.

* `status` - Enable/disable this anomaly. Valid values: `disable`, `enable`.

* `synproxy_tcp_mss` - Determine TCP maximum segment size (MSS) value for packets replied by syn proxy module. Valid values: `0`, `256`, `512`, `1024`, `1300`, `1360`, `1460`, `1500`.

* `synproxy_tcp_sack` - enable/disable TCP selective acknowledage (SACK) for packets replied by syn proxy module. Valid values: `disable`, `enable`.

* `synproxy_tcp_timestamp` - enable/disable TCP timestamp option for packets replied by syn proxy module. Valid values: `disable`, `enable`.

* `synproxy_tcp_window` - Determine TCP Window size for packets replied by syn proxy module. Valid values: `4096`, `8192`, `16384`, `32768`.

* `synproxy_tcp_windowscale` - Determine TCP window scale option value for packets replied by syn proxy module. Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `13`, `14`.

* `synproxy_tos` - Determine TCP differentiated services code point value (type of service). Valid values: `0`, `10`, `12`, `14`, `18`, `20`, `22`, `26`, `28`, `30`, `34`, `36`, `38`, `40`, `46`, `255`.

* `synproxy_ttl` - Determine Time to live (TTL) value for packets replied by syn proxy module. Valid values: `32`, `64`, `128`, `255`.

* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Threshold(Default).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Packages FirewallDosPolicy6Anomaly can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE", "dos_policy6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_dospolicy6_anomaly.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
