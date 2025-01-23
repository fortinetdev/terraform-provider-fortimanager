---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_user_policyblock"
description: |-
  Policy block write access.
---

# fortimanager_system_admin_user_policyblock
Policy block write access.

~> This resource is a sub resource for variable `policy_block` of resource `fortimanager_system_admin_user`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `user` - User.

* `policy_block_name` - Policy block names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policy_block_name}}.

## Import

System AdminUserPolicyBlock can be imported using any of these accepted formats:
```
Set import_options = ["user=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_user_policyblock.labelname {{policy_block_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

