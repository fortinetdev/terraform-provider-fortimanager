---
subcategory: "Object DLP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_exactdatamatch"
description: |-
  Configure exact-data-match template used by DLP scan.
---

# fortimanager_object_dlp_exactdatamatch
Configure exact-data-match template used by DLP scan.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `columns`: `fortimanager_object_dlp_exactdatamatch_columns`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `columns` - Columns. The structure of `columns` block is documented below.
* `data` - External resource for exact data match.
* `name` - Name of table containing the exact-data-match template.
* `optional` - Number of optional columns need to match.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `columns` block supports:

* `index` - Column index.
* `optional` - Enable/disable optional match. Valid values: `disable`, `enable`.

* `type` - Data-type for this column.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDlp ExactDataMatch can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_exactdatamatch.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
