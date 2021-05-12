---
subcategory: "Dvm"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvm_cmd_discover_device"
description: |-
  Probe a remote device and retrieve its device information and system status.
---

# fortimanager_dvm_cmd_discover_device
Probe a remote device and retrieve its device information and system status.

## Argument Reference


The following arguments are supported:


* `device` - Device. The structure of `device` block is documented below.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `device` block supports:

* `adm_pass` - Adm_Pass.
* `adm_usr` - Adm_Usr.
* `ip` - Ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dvm CmdDiscoverDevice can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvm_cmd_discover_device.labelname DvmCmdDiscoverDevice
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
