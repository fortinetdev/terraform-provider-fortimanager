---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_group_member"
description: |-
  Group members.
---

# fortimanager_system_admin_group_member
Group members.

~> This resource is a sub resource for variable `member` of resource `fortimanager_system_admin_group`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_admin_group_member" "trname" {
  group      = fortimanager_system_admin_group.trname.name
  name       = fortimanager_system_admin_radius.trname.name
  depends_on = [fortimanager_system_admin_group.trname, fortimanager_system_admin_radius.trname]
}

resource "fortimanager_system_admin_group" "trname" {
  name = "terraform-group"
}

resource "fortimanager_system_admin_radius" "trname" {
  name   = "tradius2"
  server = "ALL"
}
```

## Argument Reference


The following arguments are supported:

* `group` - Group.

* `name` - Group member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminGroupMember can be imported using any of these accepted formats:
```
Set import_options = ["group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_group_member.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

