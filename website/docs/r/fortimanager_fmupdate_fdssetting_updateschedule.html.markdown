---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fdssetting_updateschedule"
description: |-
  Configure the schedule when built-in FortiGuard retrieves antivirus and IPS updates.
---

# fortimanager_fmupdate_fdssetting_updateschedule
Configure the schedule when built-in FortiGuard retrieves antivirus and IPS updates.

## Argument Reference


The following arguments are supported:


* `day` - Configure the day the update will occur, if the freqnecy is weekly (Sunday - Saturday, default = Monday). Sunday - Update every Sunday. Monday - Update every Monday. Tuesday - Update every Tuesday. Wednesday - Update every Wednesday. Thursday - Update every Thursday. Friday - Update every Friday. Saturday - Update every Saturday. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.

* `frequency` - Configure update frequency: every - time interval, daily - once a day, weekly - once a week (default = every). every - Time interval. daily - Every day. weekly - Every week. Valid values: `every`, `daily`, `weekly`.

* `status` - Enable/disable scheduled updates. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `time` - Time interval between updates, or the hour and minute when the update occurs (hh: 0 - 23, mm: 0 - 59 or 60 = random, default = 00:10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSettingUpdateSchedule can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fdssetting_updateschedule.labelname FmupdateFdsSettingUpdateSchedule
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

