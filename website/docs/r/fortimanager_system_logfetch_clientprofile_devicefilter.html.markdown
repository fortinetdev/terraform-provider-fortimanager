---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_logfetch_clientprofile_devicefilter"
description: |-
  List of device filter.
---

# fortimanager_system_logfetch_clientprofile_devicefilter
List of device filter.

## Argument Reference


The following arguments are supported:

* `client_profile` - Client Profile.

* `fmgadom` - Adom name.
* `device` - Device name or Serial number.
* `fosid` - Add or edit a device filter.
* `vdom` - Vdom filters.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogFetchClientProfileDeviceFilter can be imported using any of these accepted formats:
```
Set import_options = ["client_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_logfetch_clientprofile_devicefilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

