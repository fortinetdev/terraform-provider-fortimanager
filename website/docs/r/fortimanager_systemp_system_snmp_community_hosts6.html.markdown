---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_snmp_community_hosts6"
description: |-
  Configure IPv6 SNMP managers.
---

# fortimanager_systemp_system_snmp_community_hosts6
Configure IPv6 SNMP managers.

~> This resource is a sub resource for variable `hosts6` of resource `fortimanager_systemp_system_snmp_community`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.
* `community` - Community.

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `disable`, `enable`.

* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.

* `fosid` - Host6 entry ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ipv6` - SNMP manager IPv6 address prefix.
* `source_ipv6` - Source IPv6 address for SNMP traps.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Systemp SystemSnmpCommunityHosts6 can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE", "community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_snmp_community_hosts6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
