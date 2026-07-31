---
subcategory: "Object CASB"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_useractivity_match_tenantsessionextraction"
description: |-
  CASB user activity tenant session extraction.
---

# fortimanager_object_casb_useractivity_match_tenantsessionextraction
CASB user activity tenant session extraction.

~> This resource is a sub resource for variable `tenant_session_extraction` of resource `fortimanager_object_casb_useractivity_match`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fortimanager_object_casb_useractivity_match_tenantsessionextraction_filters`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `user_activity` - User Activity.
* `match` - Match.

* `filters` - Filters. The structure of `filters` block is documented below.
* `jq` - CASB user activity session extraction jq script.
* `session_match` - CASB user activity session match name.
* `session_source` - Enable/disable CASB session extraction source flag. Valid values: `disable`, `enable`.

* `status` - Enable/disable CASB session extraction. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `body_type` - CASB content extraction filter body type. Valid values: `json`, `form`.

* `cookie_name` - CASB content extraction filter cookie name.
* `direction` - CASB content extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB content extraction filter header name.
* `id` - CASB content extraction filter ID.
* `place` - CASB content extraction filter place type. Valid values: `header`, `path`, `body`, `cookie`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectCasb UserActivityMatchTenantSessionExtraction can be imported using any of these accepted formats:
```
Set import_options = ["user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_useractivity_match_tenantsessionextraction.labelname ObjectCasbUserActivityMatchTenantSessionExtraction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
