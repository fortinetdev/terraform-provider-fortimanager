---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_identitybasedroute_rule"
description: |-
  Rule.
---

# fortimanager_object_firewall_identitybasedroute_rule
Rule.

~> This resource is a sub resource for variable `rule` of resource `fortimanager_object_firewall_identitybasedroute`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_identitybasedroute_rule" "trname" {
  depends_on           = [fortimanager_object_firewall_identitybasedroute.trname]
  identity_based_route = fortimanager_object_firewall_identitybasedroute.trname.name
  fosid                = 2
  gateway              = "34.8.24.1"
}

resource "fortimanager_object_firewall_identitybasedroute" "trname" {
  name = "terr-identitybasedroute"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `identity_based_route` - Identity Based Route.

* `device` - Outgoing interface for the rule.
* `gateway` - IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
* `groups` - Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space.
* `fosid` - Rule ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall IdentityBasedRouteRule can be imported using any of these accepted formats:
```
Set import_options = ["identity_based_route=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_identitybasedroute_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
