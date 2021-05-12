---
subcategory: "ObjectFirewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address6template"
description: |-
  Configure IPv6 address templates.
---

# fortimanager_object_firewall_address6template
Configure IPv6 address templates.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_image_base64` - _Image-Base64.
* `ip6` - IPv6 address prefix.
* `name` - IPv6 address template name.
* `subnet_segment` - Subnet-Segment. The structure of `subnet_segment` block is documented below.
* `subnet_segment_count` - Number of IPv6 subnet segments.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `subnet_segment` block supports:

* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value. Valid values: `disable`, `enable`.

* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `values` - Values. The structure of `values` block is documented below.

The `values` block supports:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Address6Template can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address6template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
