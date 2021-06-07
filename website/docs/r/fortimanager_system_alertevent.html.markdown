---
subcategory: "System Alert"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_alertevent"
description: |-
  Alert events.
---

# fortimanager_system_alertevent
Alert events.

## Example Usage

```hcl
resource "fortimanager_system_alertevent" "trname" {
  enable_generic_text    = ["enable"]
  enable_severity_filter = ["enable"]
  event_time_period      = 1
  name                   = "tsysalert1"
  num_events             = 1
  severity_filter        = "high"
  severity_level_comp    = ["<="]
  severity_level_logs    = ["information", "notify", "warning", "alert", "emergency"]
}
```

## Argument Reference


The following arguments are supported:


* `alert_destination` - Alert-Destination. The structure of `alert_destination` block is documented below.
* `enable_generic_text` - Enable/disable generic text match. enable - Enable setting. disable - Disable setting. Valid values: `enable`, `disable`.

* `enable_severity_filter` - Enable/disable alert severity filter. enable - Enable setting. disable - Disable setting. Valid values: `enable`, `disable`.

* `event_time_period` - Time period (hours). 0.5 - 30 minutes. 1 - 1 hour. 3 - 3 hours. 6 - 6 hours. 12 - 12 hours. 24 - 1 day. 72 - 3 days. 168 - 1 week. Valid values: `0.5`, `1`, `3`, `6`, `12`, `24`, `72`, `168`.

* `generic_text` - Text that must be contained in a log to trigger alert.
* `name` - Alert name.
* `num_events` - Minimum number of events required within time period. 1 - 1 event. 5 - 5 events. 10 - 10 events. 50 - 50 events. 100 - 100 events. Valid values: `1`, `5`, `10`, `50`, `100`.

* `severity_filter` - Required log severity to trigger alert. high - High level alert. medium-high - Medium-high level alert. medium - Medium level alert. medium-low - Medium-low level alert. low - Low level alert. Valid values: `high`, `medium-high`, `medium`, `medium-low`, `low`.

* `severity_level_comp` - Log severity threshold comparison criterion. Valid values: `>=`, `=`, `<=`.

* `severity_level_logs` - Log severity threshold level. no-check - Do not check severity level for this log type. information - Information level. notify - Notify level. warning - Warning level. error - Error level. critical - Critical level. alert - Alert level. emergency - Emergency level. Valid values: `no-check`, `information`, `notify`, `warning`, `error`, `critical`, `alert`, `emergency`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `alert_destination` block supports:

* `from` - Sender email address to use in alert emails.
* `smtp_name` - SMTP server name.
* `snmp_name` - SNMP trap name.
* `syslog_name` - Syslog server name.
* `to` - Recipient email address to use in alert emails.
* `type` - Destination type. mail - Send email alert. snmp - Send SNMP trap. syslog - Send syslog message. Valid values: `mail`, `snmp`, `syslog`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AlertEvent can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_alertevent.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

