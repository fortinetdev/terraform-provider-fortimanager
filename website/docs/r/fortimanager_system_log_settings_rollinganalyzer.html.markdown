---
subcategory: "System Log"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_settings_rollinganalyzer"
description: |-
  Log rolling policy for Network Analyzer logs.
---

# fortimanager_system_log_settings_rollinganalyzer
Log rolling policy for Network Analyzer logs.

## Example Usage

```hcl
resource "fortimanager_system_log_settings_rollinganalyzer" "trname" {
  days     = ["sun"]
  password = ["fortinet"]
  username = "admin"
}
```

## Argument Reference


The following arguments are supported:


* `days` - Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.

* `del_files` - Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: `disable`, `enable`.

* `directory` - Upload server directory, for Unix server, use absolute
* `file_size` - Roll log files when they reach this size (MB).
* `gzip_format` - Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: `disable`, `enable`.

* `hour` - Log files rolling schedule (hour).
* `ip` - Upload server IP address.
* `ip2` - Upload server IP2 address.
* `ip3` - Upload server IP3 address.
* `log_format` - Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: `native`, `text`, `csv`.

* `min` - Log files rolling schedule (minutes).
* `password` - Upload server login password.
* `password2` - Upload server login password2.
* `password3` - Upload server login password3.
* `port` - Upload server IP1 port number.
* `port2` - Upload server IP2 port number.
* `port3` - Upload server IP3 port number.
* `rolling_upgrade_status` - rolling upgrade status (1|0).
* `server_type` - Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: `ftp`, `sftp`, `scp`.

* `upload` - Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: `disable`, `enable`.

* `upload_hour` - Log files upload schedule (hour).
* `upload_mode` - Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: `backup`, `mirror`.

* `upload_trigger` - Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: `on-roll`, `on-schedule`.

* `username` - Upload server login username.
* `username2` - Upload server login username2.
* `username3` - Upload server login username3.
* `when` - Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: `none`, `daily`, `weekly`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogSettingsRollingAnalyzer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_settings_rollinganalyzer.labelname SystemLogSettingsRollingAnalyzer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

