---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_devicedisable"
description: |-
  Disable client device logging.
---

# fortimanager_system_log_devicedisable
Disable client device logging.

## Argument Reference


The following arguments are supported:


* `TTL` - Time to Live
* `device` - Device to be disabled logging
* `fosid` - ID of device logging disable entry.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogDeviceDisable can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_devicedisable.labelname SystemLogDeviceDisable
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

