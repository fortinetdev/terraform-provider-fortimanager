---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_apcfgprofile_commandlist"
description: |-
  AP local configuration command list.
---

# fortimanager_object_wirelesscontroller_apcfgprofile_commandlist
AP local configuration command list.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `apcfg_profile` - Apcfg Profile.

* `fosid` - Command ID.
* `name` - AP local configuration command name.
* `passwd_value` - AP local configuration command password value.
* `type` - The command type (default = non-password). Valid values: `non-password`, `password`.

* `value` - AP local configuration command value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWirelessController ApcfgProfileCommandList can be imported using any of these accepted formats:
```
Set import_options = ["apcfg_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_apcfgprofile_commandlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
