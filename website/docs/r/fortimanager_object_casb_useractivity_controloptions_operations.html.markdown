---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_useractivity_controloptions_operations"
description: |-
  CASB control option operations.
---

# fortimanager_object_casb_useractivity_controloptions_operations
CASB control option operations.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `user_activity` - User Activity.
* `control_options` - Control Options.

* `action` - CASB operation action. Valid values: `append`, `prepend`, `replace`, `new`, `new-on-not-found`, `delete`.

* `case_sensitive` - CASB operation search case sensitive. Valid values: `disable`, `enable`.

* `direction` - CASB operation direction. Valid values: `request`.

* `header_name` - CASB operation header name to search.
* `name` - CASB control option operation name.
* `search_key` - CASB operation key to search.
* `search_pattern` - CASB operation search pattern. Valid values: `simple`, `substr`, `regexp`.

* `target` - CASB operation target. Valid values: `header`, `path`.

* `value_from_input` - Enable/disable value from user input. Valid values: `disable`, `enable`.

* `values` - CASB operation new values.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb UserActivityControlOptionsOperations can be imported using any of these accepted formats:
```
Set import_options = ["user_activity=YOUR_VALUE", "control_options=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_useractivity_controloptions_operations.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
