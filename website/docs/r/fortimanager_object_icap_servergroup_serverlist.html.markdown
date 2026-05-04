---
subcategory: "Object ICAP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_servergroup_serverlist"
description: |-
  Add ICAP servers to a list to form a server group. Optionally assign weights to each server.
---

# fortimanager_object_icap_servergroup_serverlist
Add ICAP servers to a list to form a server group. Optionally assign weights to each server.

~> This resource is a sub resource for variable `server_list` of resource `fortimanager_object_icap_servergroup`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `server_group` - Server Group.

* `name` - ICAP server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIcap ServerGroupServerList can be imported using any of these accepted formats:
```
Set import_options = ["server_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_servergroup_serverlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
