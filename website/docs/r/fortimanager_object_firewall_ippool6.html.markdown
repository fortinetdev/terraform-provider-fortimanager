---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_ippool6"
description: |-
  Configure IPv6 IP pools.
---

# fortimanager_object_firewall_ippool6
Configure IPv6 IP pools.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`dynamic_mapping`: `fortimanager_object_firewall_ippool6_dynamic_mapping`



## Example Usage

```hcl
resource "fortimanager_object_firewall_ippool6" "trname" {
  comments = "terraform-comment"
  endip    = "2001::101"
  name     = "terraform-tefv"
  startip  = "2001::0"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `name` - IPv6 IP pool name.
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `startip` - First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `startip` - First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Ippool6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_ippool6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
