---
subcategory: "System LocalLog"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_locallog_disk_setting"
description: |-
  Settings for local disk logging.
---

# fortimanager_system_locallog_disk_setting
Settings for local disk logging.

## Example Usage

```hcl
resource "fortimanager_system_locallog_disk_setting" "trname" {
  diskfull      = "overwrite"
  roll_day      = ["sunday"]
  roll_schedule = "weekly"
  server_type   = "FTP"
  severity      = "emergency"
  status        = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `diskfull` - Policy to apply when disk is full. overwrite - Overwrite oldest log when disk is full. nolog - Stop logging when disk is full. Valid values: `overwrite`, `nolog`.

* `log_disk_full_percentage` - Consider log disk as full at this usage percentage.
* `log_disk_quota` - Quota for controlling local log size.
* `max_log_file_num` - Maximum number of log files before rolling.
* `max_log_file_size` - Maximum log file size before rolling.
* `roll_day` - Days of week to roll logs. sunday - Sunday. monday - Monday. tuesday - Tuesday. wednesday - Wednesday. thursday - Thursday. friday - Friday. saturday - Saturday. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `roll_schedule` - Frequency to check log file for rolling. none - Not scheduled. daily - Every day. weekly - Every week. Valid values: `none`, `daily`, `weekly`.

* `roll_time` - Time to roll logs (hh:mm).
* `server_type` - Server type. FTP - Upload via FTP. SFTP - Upload via SFTP. SCP - Upload via SCP. Valid values: `FTP`, `SFTP`, `SCP`.

* `severity` - Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Enable/disable local disk log. disable - Do not log to local disk. enable - Log to local disk. Valid values: `disable`, `enable`.

* `upload` - Upload log file when rolling. disable - Disable uploading when rolling log file. enable - Enable uploading when rolling log file. Valid values: `disable`, `enable`.

* `upload_delete_files` - Delete log files after uploading (default = enable). disable - Do not delete log files after uploading. enable - Delete log files after uploading. Valid values: `disable`, `enable`.

* `upload_time` - Time to upload logs (hh:mm).
* `uploaddir` - Log file upload remote directory.
* `uploadip` - IP address of log uploading server.
* `uploadpass` - Password of user account in upload server.
* `uploadport` - Server port (0 = default protocol port).
* `uploadsched` - Scheduled upload (disable = upload when rolling). disable - Upload when rolling. enable - Scheduled upload. Valid values: `disable`, `enable`.

* `uploadtype` - Types of log files that need to be uploaded. event - Upload event log. Valid values: `event`.

* `uploaduser` - User account in upload server.
* `uploadzip` - Compress upload logs. disable - Upload log files as plain text. enable - Upload log files compressed. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogDiskSetting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_locallog_disk_setting.labelname SystemLocallogDiskSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

