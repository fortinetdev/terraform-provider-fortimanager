---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_profile_fortiguardcategory"
description: |-
  Configure FortiGuard categories.
---

# fortimanager_object_videofilter_profile_fortiguardcategory
Configure FortiGuard categories.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `filters` - Filters. The structure of `filters` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `action` - VideoFilter action. Valid values: `block`, `bypass`, `monitor`.

* `category_id` - Category ID.
* `id` - ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectVideofilter ProfileFortiguardCategory can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_profile_fortiguardcategory.labelname ObjectVideofilterProfileFortiguardCategory
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
