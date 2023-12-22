---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_multicastaddress6"
description: |-
  Configure IPv6 multicast address.
---

# fortimanager_object_firewall_multicastaddress6
Configure IPv6 multicast address.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`tagging`: `fortimanager_object_firewall_multicastaddress6_tagging`



## Example Usage

```hcl
resource "fortimanager_object_firewall_multicastaddress6" "trname" {
  color   = 2
  comment = "terraform-comment"
  ip6     = "2001::10/128"
  name    = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `name` - IPv6 multicast address name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `visibility` - Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall MulticastAddress6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_multicastaddress6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
