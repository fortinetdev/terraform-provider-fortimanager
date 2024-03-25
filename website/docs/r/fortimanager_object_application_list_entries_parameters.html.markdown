---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_list_entries_parameters"
description: |-
  Application parameters.
---

# fortimanager_object_application_list_entries_parameters
Application parameters.

~> This resource is a sub resource for variable `parameters` of resource `fortimanager_object_application_list_entries`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `members`: `fortimanager_object_application_list_entries_parameters_members`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `list` - List.
* `entries` - Entries.

* `fosid` - Parameter ID.
* `members` - Members. The structure of `members` block is documented below.
* `value` - Parameter value.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `members` block supports:

* `id` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectApplication ListEntriesParameters can be imported using any of these accepted formats:
```
Set import_options = ["list=YOUR_VALUE", "entries=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_list_entries_parameters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
