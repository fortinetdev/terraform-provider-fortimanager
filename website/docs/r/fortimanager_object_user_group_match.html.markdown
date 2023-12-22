---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_group_match"
description: |-
  Group matches.
---

# fortimanager_object_user_group_match
Group matches.

~> This resource is a sub resource for variable `match` of resource `fortimanager_object_user_group`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `group` - Group.

* `_gui_meta` - _Gui_Meta.
* `group_name` - Name of matching user or group on remote authentication server.
* `fosid` - ID.
* `server_name` - Name of remote auth server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectUser GroupMatch can be imported using any of these accepted formats:
```
Set import_options = ["group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_group_match.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
