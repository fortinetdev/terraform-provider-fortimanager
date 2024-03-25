---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_meta"
description: |-
  ObjectSystem Meta
---

# fortimanager_object_system_meta
ObjectSystem Meta

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `sys_meta_fields`: `fortimanager_object_system_meta_sys_meta_fields`



## Example Usage

```hcl
resource "fortimanager_object_system_meta" "trname" {
  name = "terraform-tefv-meta"
  sys_meta_fields {
    fieldlength = "100"
    importance  = "optional"
    name        = "sys_meta_field"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Name.
* `sys_meta_fields` - Sys_Meta_Fields. The structure of `sys_meta_fields` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `sys_meta_fields` block supports:

* `fieldlength` - Fieldlength.
* `importance` - Importance. Valid values: `optional`, `required`.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem Meta can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_meta.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
