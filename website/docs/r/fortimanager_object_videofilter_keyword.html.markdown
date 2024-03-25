---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_keyword"
description: |-
  Configure video filter keywords.
---

# fortimanager_object_videofilter_keyword
Configure video filter keywords.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `word`: `fortimanager_object_videofilter_keyword_word`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `fosid` - ID.
* `match` - Keyword matching logic. Valid values: `or`, `and`.

* `name` - Name.
* `word` - Word. The structure of `word` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `word` block supports:

* `comment` - Comment.
* `name` - Name.
* `pattern_type` - Pattern type. Valid values: `wildcard`, `regex`.

* `status` - Enable(consider)/disable(ignore) this keyword. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVideofilter Keyword can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_keyword.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
