---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_videofilter_keyword_word"
description: |-
  List of keywords.
---

# fortimanager_object_videofilter_keyword_word
List of keywords.

~> This resource is a sub resource for variable `word` of resource `fortimanager_object_videofilter_keyword`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `keyword` - Keyword.

* `comment` - Comment.
* `name` - Name.
* `pattern_type` - Pattern type. Valid values: `wildcard`, `regex`.

* `status` - Enable(consider)/disable(ignore) this keyword. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVideofilter KeywordWord can be imported using any of these accepted formats:
```
Set import_options = ["keyword=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_videofilter_keyword_word.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
