---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvm_cmd_add_device"
description: |-
  Add a device to the Device Manager database.
---

# fortimanager_dvm_cmd_add_device
Add a device to the Device Manager database.

## Example Usage

```hcl
resource "fortimanager_dvm_cmd_add_device" "trname" {
  fmgadom = "root"
  device {
    ip        = "192.168.52.177"
    mgmt_mode = "fmg"
    name      = "terraform-tefv"
    adm_usr   = "admin"
    adm_pass  = "admin"
  }
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Name or ID of the ADOM where the command is to be executed on.
* `device` - Device. The structure of `device` block is documented below.
* `flags` - create_task - Create a new task in task manager database. nonblocking - The API will return immediately in for non-blocking call. This flag will be set automatically when the adding, importing, updating, and deleting a list of devices. Valid values: `none`, `create_task`, `nonblocking`, `log_dev`.

* `groups` - Groups. The structure of `groups` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `device` block supports:

* `adm_pass` - <i>add real and promote device</i>.
* `adm_usr` - <i>add real and promote device</i>.
* `desc` - <i>available for all operations</i>.
* `deviceaction` - Specify add device operations, or leave blank to add real device:<ul><li>"add_model" - add a model device.<li>"promote_unreg" - promote an unregistered device to be managed by FortiManager using information from database.</ul>
* `fazquota` - <i>available for all operations</i>.
* `ip` - <i>add real device only</i>. Add device will probe with this IP using the log in credential specified.
* `metafields` - <i>add real and model device</i>.
* `mgmt_mode` - <i>add real and model device</i>. Valid values: `unreg`, `fmg`, `faz`, `fmgfaz`.

* `mr` - <i>add model device only</i>.
* `name` - <i>required for all operations</i>. Unique name for the device.
* `os_type` - <i>add model device only</i>. Valid values: `unknown`, `fos`, `fsw`, `foc`, `fml`, `faz`, `fwb`, `fch`, `fct`, `log`, `fmg`, `fsa`, `fdd`, `fac`, `fpx`, `fna`.

* `os_ver` - <i>add model device only</i>. Valid values: `unknown`, `0.0`, `1.0`, `2.0`, `3.0`, `4.0`, `5.0`, `6.0`, `7.0`, `8.0`.

* `patch` - <i>add model device only</i>.
* `platform_str` - <i>add model device only</i>. Required for determine the platform for VM platforms.
* `sn` - <i>add model device only</i>. This attribute will be used to determine the device platform, except for VM platforms, where <i>platform_str</i> is also required.

The `groups` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dvm CmdAddDevice can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvm_cmd_add_device.labelname DvmCmdAddDevice
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
