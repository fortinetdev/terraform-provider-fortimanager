---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_profile"
description: |-
  Configure VideoFilter profile.
---

# fortimanager_object_videofilter_profile
Configure VideoFilter profile.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fortimanager_object_videofilter_profile_filters`
>- `fortiguard_category`: `fortimanager_object_videofilter_profile_fortiguardcategory`



## Example Usage

```hcl
resource "fortimanager_object_videofilter_profile" "trname" {
  name        = "terr-profile"
  comment     = "This is a Terraform example"
  dailymotion = "enable"
  fortiguard_category {
    filters {
      id = 1
    }
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `dailymotion` - Enable/disable Dailymotion video source. Valid values: `disable`, `enable`.

* `default_action` - Video filter default action. Valid values: `block`, `monitor`, `allow`.

* `filters` - Filters. The structure of `filters` block is documented below.
* `fortiguard_category` - Fortiguard-Category. The structure of `fortiguard_category` block is documented below.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `name` - Name.
* `replacemsg_group` - Replacement message group.
* `vimeo` - Enable/disable Vimeo video source. Valid values: `disable`, `enable`.

* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `youtube` - Enable/disable YouTube video source. Valid values: `disable`, `enable`.

* `youtube_channel_filter` - Set YouTube channel filter.
* `youtube_restrict` - Set YouTube-restrict mode. Valid values: `strict`, `none`, `moderate`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `action` - Video filter action. Valid values: `block`, `monitor`, `allow`.

* `category` - FortiGuard category ID.
* `channel` - Channel ID.
* `comment` - Comment.
* `id` - ID.
* `keyword` - Video filter keyword ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `type` - Filter type. Valid values: `category`, `channel`, `title`, `description`.


The `fortiguard_category` block supports:

* `filters` - Filters. The structure of `filters` block is documented below.

The `filters` block supports:

* `action` - VideoFilter action. Valid values: `block`, `bypass`, `monitor`.

* `category_id` - Category ID.
* `id` - ID.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVideofilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
