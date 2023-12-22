---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvm_cmd_add_devlist"
description: |-
  Add multiple devices to the Device Manager database.
---

# fortimanager_dvm_cmd_add_devlist
Add multiple devices to the Device Manager database.

## Example Usage

```hcl
resource "fortimanager_dvm_cmd_add_devlist" "trname" {
  fmgadom = "root"
  flags   = ["create_task"]
  add_dev_list {
    adm_pass     = ["fortinet"]
    adm_usr      = "admin"
    ip           = "10.160.88.186"
    name         = "VM64_host"
    platform_str = "FortiGate-VM64"
    sn           = "FGVMULTM23002193"
  }
}
```

## Argument Reference


The following arguments are supported:


* `add_dev_list` - A list of device objects to be added. Refer to <i>device action</i> attribute to select different type of add operation. The structure of `add_dev_list` block is documented below.
* `fmgadom` - Name or ID of the ADOM where the command is to be executed on.
* `flags` - create_task - Create a new task in task manager database. nonblocking - The API will return immediately in for non-blocking call. This flag will be set automatically when the adding, importing, updating, and deleting a list of devices. Valid values: `none`, `create_task`, `nonblocking`, `log_dev`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `add_dev_list` block supports:

* `adm_pass` - <i>add real and promote device</i>.
* `adm_usr` - <i>add real and promote device</i>.
* `authorizationtemplate` - <i>add model device only</i>. Fabric Authorization Template to auto genreate for the new model device upon creation.
* `desc` - <i>available for all operations</i>.
* `deviceaction` - Specify add device operations, or leave blank to add real device:<ul><li>"add_model" - add a model device.<li>"promote_unreg" - promote an unregistered device to be managed by FortiManager using information from database.</ul>
* `deviceblueprint` - <i>add model device only</i>. Device blueprint to apply to the new model device.
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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dvm CmdAddDevList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvm_cmd_add_devlist.labelname DvmCmdAddDevList
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
