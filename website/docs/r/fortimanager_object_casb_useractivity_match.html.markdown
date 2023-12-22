---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_useractivity_match"
description: |-
  CASB user activity match rules.
---

# fortimanager_object_casb_useractivity_match
CASB user activity match rules.

~> This resource is a sub resource for variable `match` of resource `fortimanager_object_casb_useractivity`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`rules`: `fortimanager_object_casb_useractivity_match_rules`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `user_activity` - User Activity.

* `fosid` - CASB user activity match rules ID.
* `rules` - Rules. The structure of `rules` block is documented below.
* `strategy` - CASB user activity rules strategy. Valid values: `or`, `and`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rules` block supports:

* `case_sensitive` - CASB user activity match case sensitive. Valid values: `disable`, `enable`.

* `domains` - CASB user activity domain list.
* `header_name` - CASB user activity rule header name.
* `id` - CASB user activity rule ID.
* `match_pattern` - CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB user activity rule match value.
* `methods` - CASB user activity method list.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.

* `type` - CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectCasb UserActivityMatch can be imported using any of these accepted formats:
```
Set import_options = ["user_activity=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_useractivity_match.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
