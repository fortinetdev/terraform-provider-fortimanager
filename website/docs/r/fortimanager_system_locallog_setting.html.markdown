---
subcategory: "System LocalLog"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_locallog_setting"
description: |-
  Settings for locallog logging.
---

# fortimanager_system_locallog_setting
Settings for locallog logging.

## Example Usage

```hcl
resource "fortimanager_system_locallog_setting" "trname" {
  log_interval_dev_no_logging = "300"
  log_interval_disk_full      = "200"
}
```

## Argument Reference


The following arguments are supported:


* `log_daemon_crash` - Send a logmsg when a daemon crashes. enable/disable disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `log_interval_adom_perf_stats` - Interval in minute for logging the event of adom perf stats.
* `log_interval_dev_no_logging` - Interval in minute for logging the event of no logs received from a device.
* `log_interval_disk_full` - Interval in minute for logging the event of disk full.
* `log_interval_gbday_exceeded` - Interval in minute for logging the event of the GB/Day license exceeded.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogSetting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_locallog_setting.labelname SystemLocallogSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

