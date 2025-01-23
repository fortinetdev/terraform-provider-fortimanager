---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_datatype"
description: |-
  Configure predefined data type used by DLP blocking.
---

# fortimanager_object_dlp_datatype
Configure predefined data type used by DLP blocking.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `fgd_id` - ID of object in FortiGuard database.
* `look_ahead` - Number of characters to obtain in advance for verification (1 - 255, default = 1).
* `look_back` - Number of characters required to save for verification (1 - 255, default = 1).
* `match_ahead` - Number of characters behind for match-around (1 - 4096, default = 1).
* `match_around` - Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
* `match_back` - Number of characters in front for match-around (1 - 4096, default = 1).
* `name` - Name of table containing the data type.
* `pattern` - Regular expression pattern string without look around.
* `transform` - Template to transform user input to a pattern using capture group from 'pattern'.
* `verify` - Regular expression pattern string used to verify the data type.
* `verify_transformed_pattern` - Enable/disable verification for transformed pattern. Valid values: `disable`, `enable`.

* `verify2` - Extra regular expression pattern string used to verify the data type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDlp DataType can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_datatype.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
