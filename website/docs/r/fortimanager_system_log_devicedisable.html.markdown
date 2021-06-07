---
subcategory: "System Log"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_devicedisable"
description: |-
  Disable client device logging.
---

# fortimanager_system_log_devicedisable
Disable client device logging.

## Example Usage

```hcl
resource "fortimanager_system_log_devicedisable" "trname" {
  device = "terr-FGT"
  fosid  = "1"
}
```

## Argument Reference


The following arguments are supported:


* `TTL` - Time to Live
* `device` - Device to be disabled logging
* `fosid` - ID of device logging disable entry.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogDeviceDisable can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_devicedisable.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

