---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_icon"
description: |-
  Configure OSU provider icon.
---

# fortimanager_object_wirelesscontroller_hotspot20_icon
Configure OSU provider icon.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`icon_list`: `fortimanager_object_wirelesscontroller_hotspot20_icon_iconlist`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_icon" "trname" {
  name = "terr-icon"
  icon_list {
    name = "terr-icon"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `icon_list` - Icon-List. The structure of `icon_list` block is documented below.
* `name` - Icon list ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `icon_list` block supports:

* `file` - Icon file.
* `height` - Icon height.
* `lang` - Language code.
* `name` - Icon name.
* `type` - Icon type. Valid values: `bmp`, `gif`, `jpeg`, `png`, `tiff`.

* `width` - Icon width.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20Icon can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_icon.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
