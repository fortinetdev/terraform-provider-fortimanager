---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvm_cmd_update_devlist"
description: |-
  Refresh FGFM connection and system information for a list of devices.
---

# fortimanager_dvm_cmd_update_devlist
Refresh FGFM connection and system information for a list of devices.

## Example Usage

```hcl
resource "fortimanager_dvm_cmd_update_devlist" "trname" {
  fmgadom = "root"
  update_dev_member_list {
    name = "VM64_host"
    vdom = "root"
  }
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Name or ID of the ADOM where the command is to be executed on.
* `flags` - create_task - Create a new task in task manager database. nonblocking - The API will return immediately in for non-blocking call. This flag will be set automatically when the adding, importing, updating, and deleting a list of devices. Valid values: `none`, `create_task`, `nonblocking`, `log_dev`.

* `update_dev_member_list` - Update-Dev-Member-List. The structure of `update_dev_member_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `update_dev_member_list` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dvm CmdUpdateDevList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvm_cmd_update_devlist.labelname DvmCmdUpdateDevList
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
