---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_identitybasedroute"
description: |-
  Configure identity based routing.
---

# fortimanager_object_firewall_identitybasedroute
Configure identity based routing.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`rule`: `fortimanager_object_firewall_identitybasedroute_rule`



## Example Usage

```hcl
resource "fortimanager_object_firewall_identitybasedroute" "trname" {
  comments = "terraform-comments"
  name     = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comments` - Comments.
* `name` - Name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `device` - Outgoing interface for the rule.
* `gateway` - IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
* `groups` - Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space.
* `id` - Rule ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall IdentityBasedRoute can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_identitybasedroute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
