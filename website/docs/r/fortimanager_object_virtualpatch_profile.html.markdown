---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_virtualpatch_profile"
description: |-
  Configure virtual-patch profile.
---

# fortimanager_object_virtualpatch_profile
Configure virtual-patch profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `action` - Action (pass/block). Valid values: `pass`, `block`.

* `comment` - Comment.
* `exemption` - Exemption. The structure of `exemption` block is documented below.
* `log` - Enable/disable logging of detection. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `severity` - Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `exemption` block supports:

* `device` - Device MAC addresses.
* `id` - IDs.
* `rule` - Patch signature rule IDs.
* `status` - Enable/disable exemption. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVirtualPatch Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_virtualpatch_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
