---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_profile_fortiguardcategory_filters"
description: |-
  Configure VideoFilter FortiGuard category.
---

# fortimanager_object_videofilter_profile_fortiguardcategory_filters
Configure VideoFilter FortiGuard category.

~> This resource is a sub resource for variable `filters` of resource `fortimanager_object_videofilter_profile_fortiguardcategory`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_videofilter_profile_fortiguardcategory_filters" "trname" {
  profile    = fortimanager_object_videofilter_profile.trname.name
  action     = "monitor"
  fosid      = 1
  log        = "enable"
  depends_on = [fortimanager_object_videofilter_profile.trname]
}

resource "fortimanager_object_videofilter_profile" "trname" {
  name = "terr-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - VideoFilter action. Valid values: `block`, `bypass`, `monitor`.

* `category_id` - Category ID.
* `fosid` - ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVideofilter ProfileFortiguardCategoryFilters can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_profile_fortiguardcategory_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
