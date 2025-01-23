---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvmdb_device_vdom"
description: |-
  Device VDOM table.
---

# fortimanager_dvmdb_device_vdom
Device VDOM table.

~> This resource is a sub resource for variable `vdom` of resource `fortimanager_dvmdb_device`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `comments` - Comments.
* `metafields` - Meta Fields.
* `name` - Name.
* `opmode` - Opmode. Valid values: `nat`, `transparent`.

* `rtm_prof_id` - Rtm_Prof_Id.
* `status` - Status.
* `vdom_type` - Vdom_Type. Valid values: `traffic`, `admin`.

* `vpn_id` - Vpn_Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dvmdb DeviceVdom can be imported using any of these accepted formats:
```
Set import_options = ["device=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvmdb_device_vdom.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
