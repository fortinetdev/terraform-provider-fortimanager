---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_interface_member"
description: |-
  Physical interfaces that belong to the aggregate or redundant interface.
---

# fortimanager_system_interface_member
Physical interfaces that belong to the aggregate or redundant interface.

~> This resource is a sub resource for variable `member` of resource `fortimanager_system_interface`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `interface` - Interface.

* `interface_name` - Physical interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface_name}}.

## Import

System InterfaceMember can be imported using any of these accepted formats:
```
Set import_options = ["interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_interface_member.labelname {{interface_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

