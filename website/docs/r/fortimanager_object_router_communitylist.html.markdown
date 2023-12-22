---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_router_communitylist"
description: |-
  Configure community lists.
---

# fortimanager_object_router_communitylist
Configure community lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`rule`: `fortimanager_object_router_communitylist_rule`



## Example Usage

```hcl
resource "fortimanager_object_router_communitylist" "trname" {
  name = "terr-router-communitylist"
  rule {
    id     = 1
    action = "permit"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Community list name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `type` - Community list type (standard or expanded). Valid values: `standard`, `expanded`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Permit or deny route-based operations, based on the route's COMMUNITY attribute. Valid values: `deny`, `permit`.

* `id` - ID.
* `match` - Community specifications for matching a reserved community.
* `regexp` - Ordered list of COMMUNITY attributes as a regular expression.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectRouter CommunityList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_router_communitylist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
