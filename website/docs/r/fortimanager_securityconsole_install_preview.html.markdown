---
subcategory: "Security Console"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_install_preview"
description: |-
  Generate install preview for a device.
---

# fortimanager_securityconsole_install_preview
Generate install preview for a device.

## Example Usage

```hcl
resource "fortimanager_securityconsole_install_preview" "trname" {
  fmgadom = "root"
  device  = "terr-FGT"
  flags   = ["none"]
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `device` - Device.
* `flags` - Flags. Valid values: `none`, `json`.

* `vdoms` - Vdoms.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole InstallPreview can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_install_preview.labelname SecurityconsoleInstallPreview
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
