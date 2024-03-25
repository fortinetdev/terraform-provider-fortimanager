---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_sctpfilter_profile"
description: |-
  Configure SCTP filter profiles.
---

# fortimanager_object_sctpfilter_profile
Configure SCTP filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ppid_filters`: `fortimanager_object_sctpfilter_profile_ppidfilters`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `name` - Profile name.
* `ppid_filters` - Ppid-Filters. The structure of `ppid_filters` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ppid_filters` block supports:

* `action` - Action taken when PPID is matched. Valid values: `pass`, `reset`, `replace`.

* `comment` - Comment.
* `id` - ID.
* `ppid` - Payload protocol identifier.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSctpFilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_sctpfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
