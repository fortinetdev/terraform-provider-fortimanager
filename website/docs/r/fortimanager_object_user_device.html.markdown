---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_device"
description: |-
  Configure devices.
---

# fortimanager_object_user_device
Configure devices. Applies to `FortiManager <= 6.4 and Controlled FortiOS <= 6.2`.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `alias` - Device alias.
* `avatar` - Image file for avatar (maximum 4K base64 encoded).
* `category` - Family.
* `comment` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `mac` - Device MAC address.
* `master_device` - Master device (optional).
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type.
* `user` - User name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `avatar` - Image file for avatar (maximum 4K base64 encoded).
* `category` - Family.
* `comment` - Comment.
* `family` - Family.
* `hardware_vendor` - Hardware-Vendor.
* `hardware_version` - Hardware-Version.
* `mac` - Device MAC address.
* `master_device` - Master device (optional).
* `os` - Os.
* `software_version` - Software-Version.
* `tags` - Tags.
* `type` - Type.
* `user` - User name.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{alias}}.

## Import

ObjectUser Device can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_device.labelname {{alias}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
