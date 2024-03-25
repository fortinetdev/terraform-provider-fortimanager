---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_accesscontrollist"
description: |-
  Configure WiFi bridge access control list.
---

# fortimanager_object_wirelesscontroller_accesscontrollist
Configure WiFi bridge access control list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `layer3_ipv4_rules`: `fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv4rules`
>- `layer3_ipv6_rules`: `fortimanager_object_wirelesscontroller_accesscontrollist_layer3ipv6rules`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_accesscontrollist" "trname" {
  name = "terr-accesscontrollist"
  layer3_ipv4_rules {
    action  = "deny"
    rule_id = 23
    srcport = 12
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Description.
* `layer3_ipv4_rules` - Layer3-Ipv4-Rules. The structure of `layer3_ipv4_rules` block is documented below.
* `layer3_ipv6_rules` - Layer3-Ipv6-Rules. The structure of `layer3_ipv6_rules` block is documented below.
* `name` - AP access control list name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `layer3_ipv4_rules` block supports:

* `action` - Policy action (allow | deny). Valid values: `allow`, `deny`.

* `comment` - Description.
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/&lt;network mask | mask length&gt;], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).

The `layer3_ipv6_rules` block supports:

* `action` - Policy action (allow | deny). Valid values: `allow`, `deny`.

* `comment` - Description.
* `dstaddr` - Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `srcport` - Source port (0 - 65535, default = 0, meaning any).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController AccessControlList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_accesscontrollist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
