---
subcategory: "Security Console"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_pblock_clone"
description: |-
  Securityconsole PblockClone
---

# fortimanager_securityconsole_pblock_clone
Securityconsole PblockClone

## Example Usage

```hcl
resource "fortimanager_securityconsole_pblock_clone" "trname" {
  fmgadom  = "root"
  dst_name = "terr-secconsole-pbloc-clone"
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `dst_name` - Dst_Name.
* `pblock` - Pblock.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole PblockClone can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_pblock_clone.labelname SecurityconsolePblockClone
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
