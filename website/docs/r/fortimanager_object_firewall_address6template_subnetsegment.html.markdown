---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address6template_subnetsegment"
description: |-
  IPv6 subnet segments.
---

# fortimanager_object_firewall_address6template_subnetsegment
IPv6 subnet segments.

~> This resource is a sub resource for variable `subnet_segment` of resource `fortimanager_object_firewall_address6template`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `values`: `fortimanager_object_firewall_address6template_subnetsegment_values`



## Example Usage

```hcl
resource "fortimanager_object_firewall_address6template_subnetsegment" "trname" {
  address6_template = fortimanager_object_firewall_address6template.trname.name
  bits              = 3
  exclusive         = "disable"
  fosid             = 1
  name              = "terr-subnetsegment"
  values {
    name  = "value"
    value = "0b101"
  }
  depends_on = [fortimanager_object_firewall_address6template.trname]
}

resource "fortimanager_object_firewall_address6template" "trname" {
  name = "terr-address6template"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `address6_template` - Address6 Template.

* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value. Valid values: `disable`, `enable`.

* `fosid` - Subnet segment ID.
* `name` - Subnet segment name.
* `values` - Values. The structure of `values` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `values` block supports:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall Address6TemplateSubnetSegment can be imported using any of these accepted formats:
```
Set import_options = ["address6_template=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address6template_subnetsegment.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
