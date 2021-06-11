---
subcategory: "Security Console"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_install_device"
description: |-
  Securityconsole InstallDevice
---

# fortimanager_securityconsole_install_device
Securityconsole InstallDevice

## Example Usage

```hcl
resource "fortimanager_securityconsole_install_device" "trname" {
  fmgadom          = "root"
  dev_rev_comments = "terraform-comment"
  flags            = ["none"]
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `dev_rev_comments` - Dev_Rev_Comments.
* `flags` - preview - Generate preview cache only. auto_lock_ws - Automatically lock and unlock workspace when performing security console task. Valid values: `none`, `preview`, `auto_lock_ws`.

* `scope` - Scope. The structure of `scope` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole InstallDevice can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_install_device.labelname SecurityconsoleInstallDevice
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
