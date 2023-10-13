---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_user_adom"
description: |-
  Admin domain.
---

# fortimanager_system_admin_user_adom
Admin domain.

## Argument Reference


The following arguments are supported:

* `user` - User.

* `adom_name` - Admin domain names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{adom_name}}.

## Import

System AdminUserAdom can be imported using any of these accepted formats:
```
Set import_options = ["user=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_user_adom.labelname {{adom_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

