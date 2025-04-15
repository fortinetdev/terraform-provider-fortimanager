---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_useractivity"
description: |-
  Configure CASB user activity.
---

# fortimanager_object_casb_useractivity
Configure CASB user activity.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `control_options`: `fortimanager_object_casb_useractivity_controloptions`
>- `match`: `fortimanager_object_casb_useractivity_match`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `application` - CASB SaaS application name.
* `casb_name` - CASB user activity signature name.
* `category` - CASB user activity category. Valid values: `activity-control`, `tenant-control`, `domain-control`, `safe-search-control`, `other`.

* `control_options` - Control-Options. The structure of `control_options` block is documented below.
* `description` - CASB user activity description.
* `match` - Match. The structure of `match` block is documented below.
* `match_strategy` - CASB user activity match strategy. Valid values: `or`, `and`.

* `name` - CASB user activity name.
* `status` - CASB user activity status. Valid values: `disable`, `enable`.

* `type` - CASB user activity type. Valid values: `built-in`, `customized`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `control_options` block supports:

* `name` - CASB control option name.
* `operations` - Operations. The structure of `operations` block is documented below.
* `status` - CASB control option status. Valid values: `disable`, `enable`.


The `operations` block supports:

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

The `match` block supports:

* `id` - CASB user activity match rules ID.
* `rules` - Rules. The structure of `rules` block is documented below.
* `strategy` - CASB user activity rules strategy. Valid values: `or`, `and`.

* `tenant_extraction` - Tenant-Extraction. The structure of `tenant_extraction` block is documented below.

The `rules` block supports:

* `body_type` - CASB user activity match rule body type. Valid values: `json`.

* `case_sensitive` - CASB user activity match case sensitive. Valid values: `disable`, `enable`.

* `domains` - CASB user activity domain list.
* `header_name` - CASB user activity rule header name.
* `id` - CASB user activity rule ID.
* `jq` - CASB user activity rule match jq script.
* `match_pattern` - CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB user activity rule match value.
* `methods` - CASB user activity method list.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.

* `type` - CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`.


The `tenant_extraction` block supports:

* `filters` - Filters. The structure of `filters` block is documented below.
* `jq` - CASB user activity tenant extraction jq script.
* `status` - Enable/disable CASB tenant extraction. Valid values: `disable`, `enable`.

* `type` - CASB user activity tenant extraction type. Valid values: `json-query`.


The `filters` block supports:

* `body_type` - CASB tenant extraction filter body type. Valid values: `json`.

* `direction` - CASB tenant extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB tenant extraction filter header name.
* `id` - CASB tenant extraction filter ID.
* `place` - CASB tenant extraction filter place type. Valid values: `path`, `header`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb UserActivity can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_useractivity.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
