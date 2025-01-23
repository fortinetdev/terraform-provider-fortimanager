---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_exec_fgfm_reclaimdevtunnel"
description: |-
  Reclaim management tunnel to device.  Request without device name specified will reclaim tunnels for all managed devices.
---

# fortimanager_exec_fgfm_reclaimdevtunnel
Reclaim management tunnel to device.  Request without device name specified will reclaim tunnels for all managed devices.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `flags` - Flags. Valid values: `force`.

* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Exec FgfmReclaimDevTunnel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_exec_fgfm_reclaimdevtunnel.labelname ExecFgfmReclaimDevTunnel
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
