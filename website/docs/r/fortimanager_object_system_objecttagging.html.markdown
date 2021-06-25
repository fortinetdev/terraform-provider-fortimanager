---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_objecttagging"
description: |-
  Configure object tagging.
---

# fortimanager_object_system_objecttagging
Configure object tagging.

## Example Usage

```hcl
resource "fortimanager_object_system_objecttagging" "labelname" {
  address   = "mandatory"
  category  = "ss"
  color     = 0
  device    = "mandatory"
  interface = "mandatory"
  multiple  = "enable"
  tags = [
    "11",
  ]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `address` - Address. Valid values: `optional`, `mandatory`, `disable`.

* `category` - Tag Category.
* `color` - Color of icon on the GUI.
* `device` - Device. Valid values: `optional`, `mandatory`, `disable`.

* `interface` - Interface. Valid values: `optional`, `mandatory`, `disable`.

* `multiple` - Allow multiple tag selection. Valid values: `disable`, `enable`.

* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{category}}.

## Import

ObjectSystem ObjectTagging can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_objecttagging.labelname {{category}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
