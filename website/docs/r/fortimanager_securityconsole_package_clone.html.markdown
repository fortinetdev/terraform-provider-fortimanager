---
subcategory: "Security Console"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_package_clone"
description: |-
  Clone a policy package within the same ADOM.
---

# fortimanager_securityconsole_package_clone
Clone a policy package within the same ADOM.

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `dst_name` - Name of the new policy package.
* `dst_parent` - Path to the folder for the target package. If the package is to be placed in root, leave this field blank.
* `pkg` - Source package path and name.
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

Securityconsole PackageClone can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_package_clone.labelname SecurityconsolePackageClone
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
