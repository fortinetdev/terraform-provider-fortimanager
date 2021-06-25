---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_adom_options"
description: |-
  ObjectAdom Options
---

# fortimanager_object_adom_options
ObjectAdom Options

## Example Usage

```hcl
resource "fortimanager_object_adom_options" "trname" {
  assign_excluded         = "disable"
  specify_assign_pkg_list = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `assign_excluded` - Assign_Excluded. Valid values: `disable`, `enable`.

* `specify_assign_pkg_list` - Specify_Assign_Pkg_List. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectAdom Options can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_adom_options.labelname ObjectAdomOptions
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
