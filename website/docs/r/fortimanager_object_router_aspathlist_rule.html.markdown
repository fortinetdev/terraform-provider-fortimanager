---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_router_aspathlist_rule"
description: |-
  AS path list rule.
---

# fortimanager_object_router_aspathlist_rule
AS path list rule.

~> This resource is a sub resource for variable `rule` of resource `fortimanager_object_router_aspathlist`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_router_aspathlist_rule" "trname" {
  aspath_list = fortimanager_object_router_aspathlist.trname.name
  action      = "permit"
  fosid       = 1
  depends_on  = [fortimanager_object_router_aspathlist.trname]
}

resource "fortimanager_object_router_aspathlist" "trname" {
  name = "terr-router-aspathlist"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `aspath_list` - Aspath List.

* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute. Valid values: `deny`, `permit`.

* `fosid` - ID.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectRouter AspathListRule can be imported using any of these accepted formats:
```
Set import_options = ["aspath_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_router_aspathlist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
