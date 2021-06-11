---
subcategory: "System Log"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_alert"
description: |-
  Log based alert settings.
---

# fortimanager_system_log_alert
Log based alert settings.

## Example Usage

```hcl
resource "fortimanager_system_log_alert" "trname" {
  max_alert_count = "500"
}
```

## Argument Reference


The following arguments are supported:


* `max_alert_count` - Maximum number of alerts supported.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogAlert can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_alert.labelname SystemLogAlert
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

