---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_youtubechannelfilter"
description: |-
  Configure YouTube channel filter.
---

# fortimanager_object_videofilter_youtubechannelfilter
Configure YouTube channel filter.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `default_action` - YouTube channel filter default action. Valid values: `monitor`, `block`, `allow`.

* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `log` - Eanble/disable logging. Valid values: `disable`, `enable`.

* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - YouTube channel filter action. Valid values: `block`, `bypass`, `monitor`.

* `channel_id` - Channel ID.
* `comment` - Comment.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVideofilter YoutubeChannelFilter can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_youtubechannelfilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
