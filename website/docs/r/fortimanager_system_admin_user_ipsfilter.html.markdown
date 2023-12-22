---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_user_ipsfilter"
description: |-
  IPS filter.
---

# fortimanager_system_admin_user_ipsfilter
IPS filter.

~> This resource is a sub resource for variable `ips_filter` of resource `fortimanager_system_admin_user`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `user` - User.

* `ips_filter_name` - IPS filter name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ips_filter_name}}.

## Import

System AdminUserIpsFilter can be imported using any of these accepted formats:
```
Set import_options = ["user=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_user_ipsfilter.labelname {{ips_filter_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

