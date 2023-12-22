---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address6_subnetsegment"
description: |-
  IPv6 subnet segments.
---

# fortimanager_object_firewall_address6_subnetsegment
IPv6 subnet segments.

~> This resource is a sub resource for variable `subnet_segment` of resource `fortimanager_object_firewall_address6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_address6_subnetsegment" "trname" {
  address6   = fortimanager_object_firewall_address6.trname3.name
  name       = "terr-subnetsegment"
  type       = "any"
  depends_on = [fortimanager_object_firewall_address6.trname3]
}

resource "fortimanager_object_firewall_address6" "trname3" {
  name = "terr-firewall-address6"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `address6` - Address6.

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any`, `specific`.

* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Address6SubnetSegment can be imported using any of these accepted formats:
```
Set import_options = ["address6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address6_subnetsegment.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
