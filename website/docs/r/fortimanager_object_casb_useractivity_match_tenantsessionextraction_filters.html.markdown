---
subcategory: "Object CASB"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_useractivity_match_tenantsessionextraction_filters"
description: |-
  CASB user activity session extraction filters.
---

# fortimanager_object_casb_useractivity_match_tenantsessionextraction_filters
CASB user activity session extraction filters.

~> This resource is a sub resource for variable `filters` of resource `fortimanager_object_casb_useractivity_match_tenantsessionextraction`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `user_activity` - User Activity.
* `match` - Match.

* `body_type` - CASB content extraction filter body type. Valid values: `json`, `form`.

* `cookie_name` - CASB content extraction filter cookie name.
* `direction` - CASB content extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB content extraction filter header name.
* `fosid` - CASB content extraction filter ID.
* `place` - CASB content extraction filter place type. Valid values: `header`, `path`, `body`, `cookie`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectCasb UserActivityMatchTenantSessionExtractionFilters can be imported using any of these accepted formats:
```
Set import_options = ["user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_useractivity_match_tenantsessionextraction_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
