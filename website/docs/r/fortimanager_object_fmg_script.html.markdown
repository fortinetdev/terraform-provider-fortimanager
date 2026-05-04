---
subcategory: "Object FMG"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_script"
description: |-
  ObjectFmg Script
---

# fortimanager_object_fmg_script
ObjectFmg Script

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `schedule`: `fortimanager_object_fmg_script_schedule`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `content` - Content.
* `desc` - Desc.
* `filter_build` - Filter_Build.
* `filter_device` - Filter_Device.
* `filter_hostname` - Filter_Hostname.
* `filter_ostype` - Filter_Ostype.
* `filter_osver` - Filter_Osver.
* `filter_platform` - Filter_Platform.
* `filter_serial` - Filter_Serial.
* `member` - Member.
* `name` - Name.
* `schedule` - Schedule. The structure of `schedule` block is documented below.
* `target` - Target. Valid values: `devdb`, `remote`, `adomdb`.

* `type` - Type. Valid values: `cli`, `tcl`, `cligrp`, `tclgrp`, `jinja`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `schedule` block supports:

* `datetime` - Datetime.
* `day_of_week` - Day-Of-Week.
* `device` - Device.
* `run_on_db` - Run-On-Db.
* `timestamp` - Timestamp.
* `type` - Type. Valid values: `auto`, `onetime`, `daily`, `weekly`, `monthly`.

* `user` - User.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFmg Script can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_script.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
