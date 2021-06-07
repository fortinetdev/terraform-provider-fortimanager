---
subcategory: "Object Webfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_content"
description: |-
  Configure Web filter banned word table.
---

# fortimanager_object_webfilter_content
Configure Web filter banned word table.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Block or exempt word when a match is found. Valid values: `exempt`, `block`.

* `lang` - Language of banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`, `cyrillic`.

* `name` - Banned word.
* `pattern_type` - Banned word pattern type: wildcard pattern or Perl regular expression. Valid values: `wildcard`, `regexp`.

* `score` - Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
* `status` - Enable/disable banned word. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWebfilter Content can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_content.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
