---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_alertconsole"
description: |-
  Alert console.
---

# fortimanager_system_alertconsole
Alert console.

## Argument Reference


The following arguments are supported:


* `period` - Alert console keeps alerts for this period. 1 - 1 day. 2 - 2 days. 3 - 3 days. 4 - 4 days. 5 - 5 days. 6 - 6 days. 7 - 7 days. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`.

* `severity_level` - Alert console keeps alerts of this and higher severity. debug - Debug level. information - Information level. notify - Notify level. warning - Warning level. error - Error level. critical - Critical level. alert - Alert level. emergency - Emergency level. Valid values: `debug`, `information`, `notify`, `warning`, `error`, `critical`, `alert`, `emergency`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AlertConsole can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_alertconsole.labelname SystemAlertConsole
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

