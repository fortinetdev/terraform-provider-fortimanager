---
subcategory: "Systemp"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_template_interface"
description: |-
  Systemp SystemTemplateInterface
---

# fortimanager_systemp_system_template_interface
Systemp SystemTemplateInterface

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ip_range`: `fortimanager_systemp_system_template_interface_iprange`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `action` - Action. Valid values: `add-aggregate`, `add-loopback`, `add-vlan`, `add-zone`, `conf-intf`, `conf-dhcp-server`, `conf-monitor-bandwd`, `conf-vap-ssid`.

* `allowaccess` - Allowaccess. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `dnp`, `ftm`, `fabric`, `speed-test`.

* `dhcp_id` - <i>Support meta variable</i>
* `gateway` - <i>Support meta variable</i>
* `interface` - <i>Support meta variable</i>
* `interface_members` - <i>Support meta variable</i>
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ipmask` - <i>Support meta variable</i>
* `model` - Model.
* `monitor_bandwidth` - Monitor-Bandwidth. Valid values: `disable`, `enable`.

* `name` - <i>Support meta variable</i>
* `netmask` - <i>Support meta variable</i>
* `role` - Role. Valid values: `lan`, `wan`, `dmz`, `undefined`.

* `seq` - Seq.
* `vdom` - <i>Support meta variable</i>
* `vlan_id` - <i>Support meta variable</i>
* `wifi_key` - Wifi-Key.
* `wifi_security` - Wifi-Security. Valid values: `None`, `wep64`, `wep128`, `WPA_PSK`, `WPA_RADIUS`, `WPA`, `WPA2`, `WPA2_AUTO`, `open`, `wpa-personal`, `wpa-enterprise`, `captive-portal`, `wpa-only-personal`, `wpa-only-enterprise`, `wpa2-only-personal`, `wpa2-only-enterprise`, `wpa-personal+captive-portal`, `wpa-only-personal+captive-portal`, `wpa2-only-personal+captive-portal`, `osen`, `wpa3-enterprise`, `sae`, `sae-transition`, `owe`, `wpa3-sae`, `wpa3-sae-transition`, `wpa3-only-enterprise`, `wpa3-enterprise-transition`.

* `wifi_ssid` - <i>Support meta variable</i>
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `end_ip` - End-Ip.
* `id` - Id.
* `start_ip` - Start-Ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq}}.

## Import

Systemp SystemTemplateInterface can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_template_interface.labelname {{seq}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
