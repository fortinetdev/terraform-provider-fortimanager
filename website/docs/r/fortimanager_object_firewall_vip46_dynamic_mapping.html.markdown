---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vip46_dynamic_mapping"
description: |-
  Configure IPv4 to IPv6 virtual IPs.
---

# fortimanager_object_firewall_vip46_dynamic_mapping
Configure IPv4 to IPv6 virtual IPs.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_firewall_vip46`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vip46` - Vip46.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `arp_reply` - Enable ARP reply. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `extip` - Start-external-IP [-end-external-IP].
* `extport` - External service port.
* `fosid` - Custom defined id.
* `ldb_method` - Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.

* `mappedip` - Start-mapped-IP [-end mapped-IP].
* `mappedport` - Mapped service port.
* `monitor` - Health monitors.
* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.

* `protocol` - Mapped port protocol. Valid values: `tcp`, `udp`.

* `server_type` - Server type. Valid values: `http`, `tcp`, `udp`, `ip`.

* `src_filter` - Source IP filter (x.x.x.x/x).
* `srcintf_filter` - Interfaces to which the VIP46 applies. Separate the names with spaces.
* `type` - VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectFirewall Vip46DynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["vip46=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vip46_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
