---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_icon_iconlist"
description: |-
  Icon list.
---

# fortimanager_object_wirelesscontroller_hotspot20_icon_iconlist
Icon list.

~> This resource is a sub resource for variable `icon_list` of resource `fortimanager_object_wirelesscontroller_hotspot20_icon`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_icon_iconlist" "trname" {
  name       = "terr-iconlist"
  type       = "gif"
  icon       = fortimanager_object_wirelesscontroller_hotspot20_icon.trname.name
  depends_on = [fortimanager_object_wirelesscontroller_hotspot20_icon.trname]
}

resource "fortimanager_object_wirelesscontroller_hotspot20_icon" "trname" {
  name = "terr-icon"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `icon` - Icon.

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

ObjectWirelessController Hotspot20IconIconList can be imported using any of these accepted formats:
```
Set import_options = ["icon=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_icon_iconlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
