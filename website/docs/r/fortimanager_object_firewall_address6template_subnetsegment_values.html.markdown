---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address6template_subnetsegment_values"
description: |-
  Subnet segment values.
---

# fortimanager_object_firewall_address6template_subnetsegment_values
Subnet segment values.

~> This resource is a sub resource for variable `values` of resource `fortimanager_object_firewall_address6template_subnetsegment`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_address6template_subnetsegment_values" "trname" {
  address6_template = fortimanager_object_firewall_address6template.trname2.name
  subnet_segment    = fortimanager_object_firewall_address6template_subnetsegment.trname2.fosid
  name              = "terr-values"
  value             = "0b101"
  depends_on        = [fortimanager_object_firewall_address6template_subnetsegment.trname2]
}

resource "fortimanager_object_firewall_address6template_subnetsegment" "trname2" {
  address6_template = fortimanager_object_firewall_address6template.trname2.name
  bits              = 3
  fosid             = 1
  name              = "terr-subnetsegment"
  depends_on        = [fortimanager_object_firewall_address6template.trname2]
}

resource "fortimanager_object_firewall_address6template" "trname2" {
  name = "terr-address6template"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `address6_template` - Address6 Template.
* `subnet_segment` - Subnet Segment.

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Address6TemplateSubnetSegmentValues can be imported using any of these accepted formats:
```
Set import_options = ["address6_template=YOUR_VALUE", "subnet_segment=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address6template_subnetsegment_values.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
