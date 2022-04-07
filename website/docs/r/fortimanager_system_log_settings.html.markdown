---
subcategory: "System Log"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_settings"
description: |-
  Log settings.
---

# fortimanager_system_log_settings
Log settings.

## Example Usage

```hcl
resource "fortimanager_system_log_settings" "trname" {
  dns_resolve_dstip     = "disable"
  ha_auto_migrate       = "disable"
  log_file_archive_name = "basic"
}
```

## Argument Reference


The following arguments are supported:


* `FAC_custom_field1` - Name of custom log field to index.
* `FAZ_custom_field1` - Name of custom log field to index.
* `FCH_custom_field1` - Name of custom log field to index.
* `FCT_custom_field1` - Name of custom log field to index.
* `FDD_custom_field1` - Name of custom log field to index.
* `FGT_custom_field1` - Name of custom log field to index.
* `FMG_custom_field1` - Name of custom log field to index.
* `FML_custom_field1` - Name of custom log field to index.
* `FPX_custom_field1` - Name of custom log field to index.
* `FSA_custom_field1` - Name of custom log field to index.
* `FWB_custom_field1` - Name of custom log field to index.
* `browse_max_logfiles` - Maximum number of log files for each log browse attempt for each Adom.
* `dns_resolve_dstip` - Enable/Disable resolving destination IP by DNS. disable - Disable resolving destination IP by DNS. enable - Enable resolving destination IP by DNS. Valid values: `disable`, `enable`.

* `download_max_logs` - Maximum number of logs for each log download attempt.
* `ha_auto_migrate` - Enabled/Disable automatically merging HA member's logs to HA cluster. disable - Disable automatically merging HA member's logs to HA cluster. enable - Enable automatically merging HA member's logs to HA cluster. Valid values: `disable`, `enable`.

* `import_max_logfiles` - Maximum number of log files for each log import attempt.
* `keep_dev_logs` - Enable/Disable keeping the dev logs after the device has been deleted. disable - Disable keeping the dev logs after the device has been deleted. enable - Enable keeping the dev logs after the device has been deleted. Valid values: `disable`, `enable`.

* `log_file_archive_name` - Log file name format for archiving, such as backup, upload or download. basic - Basic format for log archive file name, e.g. FGT20C0000000001.tlog.1417797247.log. extended - Extended format for log archive file name, e.g. FGT20C0000000001.2014-12-05-08:34:58.tlog.1417797247.log. Valid values: `basic`, `extended`.

* `rolling_analyzer` - Rolling-Analyzer. The structure of `rolling_analyzer` block is documented below.
* `rolling_local` - Rolling-Local. The structure of `rolling_local` block is documented below.
* `rolling_regular` - Rolling-Regular. The structure of `rolling_regular` block is documented below.
* `sync_search_timeout` - Maximum number of seconds for running a log search session in synchronous mode.

The `rolling_analyzer` block supports:

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


The `rolling_local` block supports:

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


The `rolling_regular` block supports:

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

System LogSettings can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_settings.labelname SystemLogSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

