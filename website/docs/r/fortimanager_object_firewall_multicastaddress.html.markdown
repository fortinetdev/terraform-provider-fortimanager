---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_multicastaddress"
description: |-
  Configure multicast addresses.
---

# fortimanager_object_firewall_multicastaddress
Configure multicast addresses.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`tagging`: `fortimanager_object_firewall_multicastaddress_tagging`



## Example Usage

```hcl
resource "fortimanager_object_firewall_multicastaddress" "trname" {
  color = 0
  name  = "sdf"
  subnet = [
    "255.255.111.0",
    "255.255.255.255",
  ]
  type = "broadcastmask"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Comment.
* `end_ip` - Final IPv4 address (inclusive) in the range for the address.
* `name` - Multicast address name.
* `start_ip` - First IPv4 address (inclusive) in the range for the address.
* `subnet` - Broadcast address and subnet.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.

* `visibility` - Enable/disable visibility of the multicast address on the GUI. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall MulticastAddress can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_multicastaddress.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
