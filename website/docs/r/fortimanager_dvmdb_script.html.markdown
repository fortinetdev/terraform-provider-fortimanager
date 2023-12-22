---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvmdb_script"
description: |-
  Script table.
---

# fortimanager_dvmdb_script
Script table.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`execute`: `fortimanager_dvmdb_script_execute`
`script_schedule`: `fortimanager_dvmdb_script_script_schedule`



## Example Usage

```hcl
resource "fortimanager_dvmdb_script" "trname" {
  content = "terraform-tefv"
  name    = "terraform-tefv"
  target  = "device_database"
  type    = "cli"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `content` - The full content of the script result log.
* `desc` - Desc.
* `filter_build` - The value will be ignored in add/set/update requests if filter_ostype is not set. It has no effect if target is "adom_database".
* `filter_device` - Name or id of an existing device in the database. It has no effect if target is "adom_database".
* `filter_hostname` - The value has no effect if target is "adom_database".
* `filter_ostype` - The value has no effect if target is "adom_database". Valid values: `unknown`, `fos`.

* `filter_osver` - The value will be ignored in add/set/update requests if filter_ostype is not set. It has no effect if target is "adom_database". Valid values: `unknown`, `4.00`, `5.00`, `6.00`.

* `filter_platform` - The value will be ignored in add/set/update requests if filter_ostype is not set. It has no effect if target is "adom_database".
* `filter_serial` - The value has no effect if target is "adom_database".
* `modification_time` - It is a read-only attribute indicating the time when the script was created or modified. The value will be ignored in add/set/update requests.
* `name` - Name.
* `script_schedule` - Script_Schedule. The structure of `script_schedule` block is documented below.
* `target` - Target. Valid values: `device_database`, `remote_device`, `adom_database`.

* `type` - Type. Valid values: `cli`, `tcl`, `cligrp`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `script_schedule` block supports:

* `datetime` - Indicates the date and time of the schedule. It should follow the following format for each scheduling type:<ul><li>onetime: "YYYY-MM-DD hh:mm:ss"</li><li>daily: "hh:mm"</li><li>weekly: "hh:mm"</li><li>monthly: "DD hh:mm"</li></ul>
* `day_of_week` - Day_Of_Week. Valid values: `unknown`, `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.

* `device` - Name or id of an existing device in the database.
* `name` - Name.
* `run_on_db` - Indicates if the scheduled script should be executed on device database. It should always be disable for tcl scripts. Valid values: `disable`, `enable`.

* `type` - Type. Valid values: `auto`, `onetime`, `daily`, `weekly`, `monthly`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dvmdb Script can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvmdb_script.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
