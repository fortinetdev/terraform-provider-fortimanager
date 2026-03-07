---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_imageanalyzer_profile"
description: |-
  ObjectImageAnalyzer Profile
---

# fortimanager_object_imageanalyzer_profile
ObjectImageAnalyzer Profile

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `alcohol_block_strictness_level` - Alcohol-Block-Strictness-Level.
* `alcohol_status` - Alcohol-Status. Valid values: `allow`, `deny`, `monitor`.

* `blocked_img_cache` - Blocked-Img-Cache. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `drugs_block_strictness_level` - Drugs-Block-Strictness-Level.
* `drugs_status` - Drugs-Status. Valid values: `allow`, `deny`, `monitor`.

* `extremism_block_strictness_level` - Extremism-Block-Strictness-Level.
* `extremism_status` - Extremism-Status. Valid values: `allow`, `deny`, `monitor`.

* `gambling_block_strictness_level` - Gambling-Block-Strictness-Level.
* `gambling_status` - Gambling-Status. Valid values: `allow`, `deny`, `monitor`.

* `gore_block_strictness_level` - Gore-Block-Strictness-Level.
* `gore_status` - Gore-Status. Valid values: `allow`, `deny`, `monitor`.

* `image_skip_height` - Image-Skip-Height.
* `image_skip_size` - Image-Skip-Size.
* `image_skip_width` - Image-Skip-Width.
* `log_option` - Log-Option. Valid values: `all`, `violation`.

* `name` - Name.
* `ocr_activation_threshold` - Ocr-Activation-Threshold.
* `optical_character_recognition` - Optical-Character-Recognition. Valid values: `disable`, `enable`.

* `porn_block_strictness_level` - Porn-Block-Strictness-Level.
* `porn_status` - Porn-Status. Valid values: `allow`, `deny`, `monitor`.

* `rating_err_action` - Rating-Err-Action. Valid values: `block`, `pass`.

* `replace_image` - Replace-Image.
* `source_url` - Source-Url. Valid values: `disable`, `enable`.

* `swim_underwear_block_strictness_level` - Swim_Underwear-Block-Strictness-Level.
* `swim_underwear_status` - Swim_Underwear-Status. Valid values: `allow`, `deny`, `monitor`.

* `weapons_block_strictness_level` - Weapons-Block-Strictness-Level.
* `weapons_status` - Weapons-Status. Valid values: `allow`, `deny`, `monitor`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectImageAnalyzer Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_imageanalyzer_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
