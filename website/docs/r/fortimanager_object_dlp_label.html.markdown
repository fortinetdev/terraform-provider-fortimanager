---
subcategory: "Object DLP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dlp_label"
description: |-
  Configure labels used by DLP blocking.
---

# fortimanager_object_dlp_label
Configure labels used by DLP blocking.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fortimanager_object_dlp_label_entries`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `connector` - Name of SDN connector.
* `entries` - Entries. The structure of `entries` block is documented below.
* `mpip_type` - MPIP label type. Valid values: `local`, `remote`.

* `name` - Name of table containing the label.
* `type` - Label type. Valid values: `mpip`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `fortidata_label_name` - Name of FortiData label
* `guid` - MPIP label guid.
* `id` - ID.
* `mpip_label_name` - Name of MPIP label.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDlp Label can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dlp_label.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
