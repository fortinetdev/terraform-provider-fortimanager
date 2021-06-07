---
subcategory: "Object Webfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_ftgdlocalcat"
description: |-
  Configure FortiGuard Web Filter local categories.
---

# fortimanager_object_webfilter_ftgdlocalcat
Configure FortiGuard Web Filter local categories.

## Example Usage

```hcl
resource "fortimanager_object_webfilter_ftgdlocalcat" "labelname" {
  desc   = "s2s"
  status = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `desc` - Local category description.
* `fosid` - Local category ID.
* `status` - Enable/disable the local category. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{desc}}.

## Import

ObjectWebfilter FtgdLocalCat can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_ftgdlocalcat.labelname {{desc}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
