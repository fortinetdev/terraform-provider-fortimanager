---
subcategory: "System Admin"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_group"
description: |-
  User group.
---

# fortimanager_system_admin_group
User group.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `member`: `fortimanager_system_admin_group_member`



## Example Usage

```hcl
resource "fortimanager_system_admin_group" "trname" {
  name = "terraform-group"
}
```

## Argument Reference


The following arguments are supported:


* `member` - Member. The structure of `member` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `member` block supports:

* `name` - Group member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

