---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_meta_sys_meta_fields"
description: |-
  ObjectSystem MetaSysMetaFields
---

# fortimanager_object_system_meta_sys_meta_fields
ObjectSystem MetaSysMetaFields

~> This resource is a sub resource for variable `sys_meta_fields` of resource `fortimanager_object_system_meta`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_meta_sys_meta_fields" "trname" {
  meta        = fortimanager_object_system_meta.trname.name
  name        = "sys_meta_field"
  importance  = "optional"
  fieldlength = 23
  depends_on  = [fortimanager_object_system_meta.trname]
}

resource "fortimanager_object_system_meta" "trname" {
  name = "terr-system-meta"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `meta` - Meta.

* `fieldlength` - Fieldlength.
* `importance` - Importance. Valid values: `optional`, `required`.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem MetaSysMetaFields can be imported using any of these accepted formats:
```
Set import_options = ["meta=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_meta_sys_meta_fields.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
