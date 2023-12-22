---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_hascheduledcheck"
description: |-
  Scheduled HA integrity check.
---

# fortimanager_system_hascheduledcheck
Scheduled HA integrity check.

## Example Usage

```hcl
resource "fortimanager_system_hascheduledcheck" "trname" {
  status    = "enable"
  week_days = ["sunday"]
  time      = "12:12:12"
}
```

## Argument Reference


The following arguments are supported:


* `status` - Enable/disable schedule backup. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `time` - Time to backup.
* `week_days` - Week days to backup. monday - Monday. tuesday - Tuesday. wednesday - Wednesday. thursday - Thursday. friday - Friday. saturday - Saturday. sunday - Sunday. Valid values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System HaScheduledCheck can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_hascheduledcheck.labelname SystemHaScheduledCheck
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

