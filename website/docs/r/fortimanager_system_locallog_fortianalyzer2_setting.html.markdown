---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_locallog_fortianalyzer2_setting"
description: |-
  Settings for locallog to fortianalyzer.
---

# fortimanager_system_locallog_fortianalyzer2_setting
Settings for locallog to fortianalyzer.

## Argument Reference


The following arguments are supported:


* `reliable` - Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: `disable`, `enable`.

* `secure_connection` - Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: `disable`, `enable`.

* `server` - Remote FortiAnalyzer server FQDN, hostname, or IP address
* `severity` - Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Log to FortiAnalyzer status. disable - Log to FortiAnalyzer disabled. realtime - Log to FortiAnalyzer in realtime. upload - Log to FortiAnalyzer at schedule time. Valid values: `disable`, `realtime`, `upload`.

* `upload_time` - Time to upload local log files (hh:mm).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogFortianalyzer2Setting can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_locallog_fortianalyzer2_setting.labelname SystemLocallogFortianalyzer2Setting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

