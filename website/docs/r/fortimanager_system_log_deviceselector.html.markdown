---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_deviceselector"
description: |-
  Accept/reject devices matching specified filter types.
---

# fortimanager_system_log_deviceselector
Accept/reject devices matching specified filter types.

## Argument Reference


The following arguments are supported:


* `action` - Include or exclude the matching devices. include - Include devices matching specified filter type. exclude - Exclude devices matching specified filter type. Valid values: `include`, `exclude`.

* `comment` - Additional comment for the selector.
* `devid` - Device ID. Wildcard matching supported.
* `expire` - Expiration time of the selector.
* `fosid` - ID of device selector entry.
* `srcip` - Source IP or IP range.
* `srcip_mode` - Apply the selector to UDP/514, TCP/514 or any mode. UDP514 - Clients logging through UDP port 514. TCP514 - Clients logging through TCP port 514. any - Clients logging through any mode. Valid values: `UDP514`, `TCP514`, `any`.

* `type` - Type of the selector. unspecified - Filter type unspecified. devid - Filter devices by DeviceID. srcip - Filter devices by source IP. Valid values: `unspecified`, `devid`, `srcip`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogDeviceSelector can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_deviceselector.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

