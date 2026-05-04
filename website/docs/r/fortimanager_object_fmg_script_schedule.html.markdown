---
subcategory: "Object FMG"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_script_schedule"
description: |-
  ObjectFmg ScriptSchedule
---

# fortimanager_object_fmg_script_schedule
ObjectFmg ScriptSchedule

~> This resource is a sub resource for variable `schedule` of resource `fortimanager_object_fmg_script`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `script` - Script.

* `datetime` - Datetime.
* `day_of_week` - Day-Of-Week.
* `device` - Device.
* `run_on_db` - Run-On-Db.
* `timestamp` - Timestamp.
* `type` - Type. Valid values: `auto`, `onetime`, `daily`, `weekly`, `monthly`.

* `user` - User.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{device}}.

## Import

ObjectFmg ScriptSchedule can be imported using any of these accepted formats:
```
Set import_options = ["script=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_script_schedule.labelname {{device}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
