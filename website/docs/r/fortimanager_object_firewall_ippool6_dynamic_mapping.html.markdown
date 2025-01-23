---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_ippool6_dynamic_mapping"
description: |-
  Configure IPv6 IP pools.
---

# fortimanager_object_firewall_ippool6_dynamic_mapping
Configure IPv6 IP pools.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_firewall_ippool6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ippool6` - Ippool6.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `external_prefix` - <i>Support meta variable</i> External NPTv6 prefix length (32 - 64).
* `internal_prefix` - <i>Support meta variable</i> Internal NPTv6 prefix length (32 - 64).
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `startip` - First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `type` - Configure IPv6 pool type (overload or NPTv6). Valid values: `overload`, `nptv6`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectFirewall Ippool6DynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["ippool6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_ippool6_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
