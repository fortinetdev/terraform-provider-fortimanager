---
subcategory: "Object DLP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_exactdatamatch_columns"
description: |-
  DLP exact-data-match column types.
---

# fortimanager_object_dlp_exactdatamatch_columns
DLP exact-data-match column types.

~> This resource is a sub resource for variable `columns` of resource `fortimanager_object_dlp_exactdatamatch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `exact_data_match` - Exact Data Match.

* `index` - Column index.
* `optional` - Enable/disable optional match. Valid values: `disable`, `enable`.

* `type` - Data-type for this column.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

ObjectDlp ExactDataMatchColumns can be imported using any of these accepted formats:
```
Set import_options = ["exact_data_match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_exactdatamatch_columns.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
