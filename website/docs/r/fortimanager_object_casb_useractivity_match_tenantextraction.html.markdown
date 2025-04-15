---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_useractivity_match_tenantextraction"
description: |-
  CASB user activity tenant extraction.
---

# fortimanager_object_casb_useractivity_match_tenantextraction
CASB user activity tenant extraction.

~> This resource is a sub resource for variable `tenant_extraction` of resource `fortimanager_object_casb_useractivity_match`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fortimanager_object_casb_useractivity_match_tenantextraction_filters`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `user_activity` - User Activity.
* `match` - Match.

* `filters` - Filters. The structure of `filters` block is documented below.
* `jq` - CASB user activity tenant extraction jq script.
* `status` - Enable/disable CASB tenant extraction. Valid values: `disable`, `enable`.

* `type` - CASB user activity tenant extraction type. Valid values: `json-query`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `body_type` - CASB tenant extraction filter body type. Valid values: `json`.

* `direction` - CASB tenant extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB tenant extraction filter header name.
* `id` - CASB tenant extraction filter ID.
* `place` - CASB tenant extraction filter place type. Valid values: `path`, `header`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectCasb UserActivityMatchTenantExtraction can be imported using any of these accepted formats:
```
Set import_options = ["user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_useractivity_match_tenantextraction.labelname ObjectCasbUserActivityMatchTenantExtraction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
