---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_list_entries_parameters_members"
description: |-
  Parameter tuple members.
---

# fortimanager_object_application_list_entries_parameters_members
Parameter tuple members.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `list` - List.
* `entries` - Entries.
* `parameters` - Parameters.

* `fosid` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectApplication ListEntriesParametersMembers can be imported using any of these accepted formats:
```
Set import_options = ["list=YOUR_VALUE", "entries=YOUR_VALUE", "parameters=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_list_entries_parameters_members.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
