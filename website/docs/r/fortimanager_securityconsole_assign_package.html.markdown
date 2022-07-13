---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_assign_package"
description: |-
  Assign or unassign global policy package to ADOM packages.
---

# fortimanager_securityconsole_assign_package
Assign or unassign global policy package to ADOM packages.

## Example Usage

```hcl
resource "fortimanager_securityconsole_assign_package" "trname" {
  flags = ["none"]
  pkg     = "default"
  target {
    adom = "root"
    excluded = "enable"
  }
}
```

## Argument Reference


The following arguments are supported:


* `flags` - cp_all_objs - Assign all objects during global policy assignment. copy_assigned_pkg - For global policy assignment - copy assigned package from ADOM to device. unassign - Remove global policy from ADOM. Valid values: `none`, `cp_all_objs`, `copy_assigned_pkg`, `unassign`.

* `pkg` - Source package path and name.
* `target` - Target. The structure of `target` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.

The `target` block supports:

* `adom` - Destination ADOM.
* `excluded` - disable - Only include the packages listed in the <i>pkg</i> list. enable - Exclude the package listed in the <i>pkg</i> list, and assign to all other packages in the ADOM. Valid values: `disable`, `enable`.

* `pkg` - Destination ADOM policy package path and name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole AssignPackage can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_assign_package.labelname SecurityconsoleAssignPackage
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
