---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvmdb_group"
description: |-
  Device group table.
---

# fortimanager_dvmdb_group
Device group table.

## Example Usage

```hcl
resource "fortimanager_dvmdb_group" "trname" {
  desc    = "terraform-tefv"
  name    = "terraform-tefv"
  os_type = "fos"
  type    = "normal"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `desc` - Desc.
* `metafields` - Default metafields: none.
* `name` - Name.
* `os_type` - Os_Type. Valid values: `unknown`, `fos`, `fsw`, `foc`, `fml`, `faz`, `fwb`, `fch`, `fct`, `log`, `fmg`, `fsa`, `fdd`, `fac`, `fpx`, `fna`, `ffw`, `fsr`, `fad`, `fdc`.

* `type` - Type. Valid values: `normal`, `default`, `auto`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dvmdb Group can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvmdb_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
