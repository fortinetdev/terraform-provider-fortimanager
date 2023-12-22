---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_log_syslogd_setting"
description: |-
  Global settings for remote syslog server.
---

# fortimanager_systemp_log_syslogd_setting
Global settings for remote syslog server.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`custom_field_name`: `fortimanager_systemp_log_syslogd_setting_customfieldname`



## Example Usage

```hcl
resource "fortimanager_systemp_log_syslogd_setting" "trname" {
  devprof      = "default"
  format       = "rfc5424"
  max_log_rate = 200
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `certificate` - Certificate used to communicate with Syslog server.
* `custom_field_name` - Custom-Field-Name. The structure of `custom_field_name` block is documented below.
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption. Valid values: `high`, `low`, `disable`, `high-medium`.

* `facility` - Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.

* `format` - Log format. Valid values: `default`, `csv`, `cef`, `rfc5424`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `mode` - Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.

* `port` - Server listen port.
* `priority` - Set log transmission priority. Valid values: `low`, `default`.

* `server` - Address of remote syslog server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`.

* `status` - Enable/disable remote syslog logging. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_field_name` block supports:

* `custom` - Field custom name.
* `id` - Entry ID.
* `name` - Field name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp LogSyslogdSetting can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_log_syslogd_setting.labelname SystempLogSyslogdSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
