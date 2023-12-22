---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_profile_fortiguardcategory_filters_move"
description: |-
  Move VideoFilter FortiGuard category.
---

# fortimanager_object_videofilter_profile_fortiguardcategory_filters_move
Move VideoFilter FortiGuard category.

## Example Usage

```hcl
resource "fortimanager_object_videofilter_profile_fortiguardcategory_filters_move" "trname" {
  profile = fortimanager_object_videofilter_profile.trname.name
  filters = fortimanager_object_videofilter_profile_fortiguardcategory_filters.trname.fosid
  target  = fortimanager_object_videofilter_profile_fortiguardcategory_filters.trname2.fosid
  option  = "before"
}

resource "fortimanager_object_videofilter_profile_fortiguardcategory_filters" "trname2" {
  profile    = fortimanager_object_videofilter_profile.trname.name
  fosid      = 2
  depends_on = [fortimanager_object_videofilter_profile.trname]
}

resource "fortimanager_object_videofilter_profile_fortiguardcategory_filters" "trname" {
  profile    = fortimanager_object_videofilter_profile.trname.name
  fosid      = 1
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
* `filters` - Filters.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the filters changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of filter.
