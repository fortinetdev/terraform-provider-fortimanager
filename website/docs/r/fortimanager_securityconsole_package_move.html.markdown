---
subcategory: "Security Console"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_package_move"
description: |-
  Move and/or rename a policy package within the same ADOM.
---

# fortimanager_securityconsole_package_move
Move and/or rename a policy package within the same ADOM.

## Example Usage

```hcl
resource "fortimanager_securityconsole_package_move" "trname" {
  fmgadom  = "root"
  dst_name = "terr-secconsole-pkg-move"
  pkg      = "terr-pkg"
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `dst_name` - Name of the new policy package. If omitted from request, original name will be kept or a new name will be generated in case of duplication.
* `dst_parent` - Path to the folder for the target package. If the package is to be placed in root, leave this field blank.
* `pkg` - Source package path and name.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole PackageMove can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_package_move.labelname SecurityconsolePackageMove
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
