---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_dictionary"
description: |-
  Configure dictionaries used by DLP blocking.
---

# fortimanager_object_dlp_dictionary
Configure dictionaries used by DLP blocking.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `match_type` - Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.

* `name` - Name of table containing the dictionary.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `comment` - Optional comments.
* `id` - ID.
* `ignore_case` - Enable/disable ignore case. Valid values: `disable`, `enable`.

* `pattern` - Pattern to match.
* `repeat` - Enable/disable repeat match. Valid values: `disable`, `enable`.

* `status` - Enable/disable this pattern. Valid values: `disable`, `enable`.

* `type` - Pattern type to match.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDlp Dictionary can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_dictionary.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
