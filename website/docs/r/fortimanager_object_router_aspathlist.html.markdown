---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_router_aspathlist"
description: |-
  Configure Autonomous System (AS) path lists.
---

# fortimanager_object_router_aspathlist
Configure Autonomous System (AS) path lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fortimanager_object_router_aspathlist_rule`



## Example Usage

```hcl
resource "fortimanager_object_router_aspathlist" "trname" {
  name = "terr-router-aspathlist"
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

* `name` - AS path list name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute. Valid values: `deny`, `permit`.

* `id` - ID.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectRouter AspathList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_router_aspathlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
