---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_dictionary_entries"
description: |-
  DLP dictionary entries.
---

# fortimanager_object_dlp_dictionary_entries
DLP dictionary entries.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_dlp_dictionary`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `dictionary` - Dictionary.

* `comment` - Optional comments.
* `fosid` - ID.
* `ignore_case` - Enable/disable ignore case. Valid values: `disable`, `enable`.

* `pattern` - Pattern to match.
* `repeat` - Enable/disable repeat match. Valid values: `disable`, `enable`.

* `status` - Enable/disable this pattern. Valid values: `disable`, `enable`.

* `type` - Pattern type to match.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDlp DictionaryEntries can be imported using any of these accepted formats:
```
Set import_options = ["dictionary=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_dictionary_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
