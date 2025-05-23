---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extenderprofile_lanextension"
description: |-
  FortiExtender lan extension configuration.
---

# fortimanager_object_extensioncontroller_extenderprofile_lanextension
FortiExtender lan extension configuration.

~> This resource is a sub resource for variable `lan_extension` of resource `fortimanager_object_extensioncontroller_extenderprofile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `backhaul`: `fortimanager_object_extensioncontroller_extenderprofile_lanextension_backhaul`
>- `downlinks`: `fortimanager_object_extensioncontroller_extenderprofile_lanextension_downlinks`
>- `traffic_split_services`: `fortimanager_object_extensioncontroller_extenderprofile_lanextension_trafficsplitservices`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `backhaul` - Backhaul. The structure of `backhaul` block is documented below.
* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `downlinks` - Downlinks. The structure of `downlinks` block is documented below.
* `ipsec_tunnel` - IPsec tunnel name.
* `link_loadbalance` - LAN extension link load balance strategy. Valid values: `activebackup`, `loadbalance`.

* `traffic_split_services` - Traffic-Split-Services. The structure of `traffic_split_services` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `backhaul` block supports:

* `name` - FortiExtender LAN extension backhaul name.
* `port` - FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.

* `role` - FortiExtender uplink port. Valid values: `primary`, `secondary`.

* `weight` - WRR weight parameter.

The `downlinks` block supports:

* `name` - FortiExtender LAN extension downlink config entry name.
* `port` - FortiExtender LAN extension downlink port. Valid values: `port1`, `port2`, `port3`, `port4`, `port5`, `lan1`, `lan2`.

* `pvid` - FortiExtender LAN extension downlink PVID.
* `type` - FortiExtender LAN extension downlink type [port/vap]. Valid values: `port`, `vap`.

* `vap` - FortiExtender LAN extension downlink vap.

The `traffic_split_services` block supports:

* `address` - Address selection.
* `name` - FortiExtender LAN extension tunnel split entry name.
* `service` - Service selection.
* `vsdb` - Select vsdb [enable/disable]. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectExtensionController ExtenderProfileLanExtension can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extenderprofile_lanextension.labelname ObjectExtensionControllerExtenderProfileLanExtension
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
