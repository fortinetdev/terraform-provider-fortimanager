---
subcategory: "ObjectWireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_region"
description: |-
  Configure FortiAP regions (for floor plans and maps).
---

# fortimanager_object_wirelesscontroller_region
Configure FortiAP regions (for floor plans and maps).

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comments` - Comments.
* `grayscale` - Region image grayscale. Valid values: `disable`, `enable`.

* `image_type` - FortiAP region image type (png|jpeg|gif). Valid values: `gif`, `jpeg`, `png`.

* `name` - FortiAP region name.
* `opacity` - Region image opacity (0 - 100).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Region can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_region.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
