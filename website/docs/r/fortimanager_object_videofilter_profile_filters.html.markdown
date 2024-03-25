---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_profile_filters"
description: |-
  YouTube filter entries.
---

# fortimanager_object_videofilter_profile_filters
YouTube filter entries.

~> This resource is a sub resource for variable `filters` of resource `fortimanager_object_videofilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Video filter action. Valid values: `block`, `monitor`, `allow`.

* `category` - FortiGuard category ID.
* `channel` - Channel ID.
* `comment` - Comment.
* `fosid` - ID.
* `keyword` - Video filter keyword ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `type` - Filter type. Valid values: `category`, `channel`, `title`, `description`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVideofilter ProfileFilters can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_profile_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
