---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_router_prefixlist_rule"
description: |-
  IPv4 prefix list rule.
---

# fortimanager_object_router_prefixlist_rule
IPv4 prefix list rule.

~> This resource is a sub resource for variable `rule` of resource `fortimanager_object_router_prefixlist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_router_prefixlist_rule" "trname" {
  prefix_list = fortimanager_object_router_prefixlist.trname.name
  ge          = 12
  fosid       = 23
  le          = 20
  depends_on  = [fortimanager_object_router_prefixlist.trname]
}

resource "fortimanager_object_router_prefixlist" "trname" {
  name = "terr-router-prefixlist"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `prefix_list` - Prefix List.

* `action` - Permit or deny this IP address and netmask prefix. Valid values: `permit`, `deny`.

* `flags` - Flags.
* `ge` - Minimum prefix length to be matched (0 - 32).
* `fosid` - Rule ID.
* `le` - Maximum prefix length to be matched (0 - 32).
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectRouter PrefixListRule can be imported using any of these accepted formats:
```
Set import_options = ["prefix_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_router_prefixlist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
