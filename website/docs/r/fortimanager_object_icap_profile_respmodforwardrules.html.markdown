---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_profile_respmodforwardrules"
description: |-
  ICAP response mode forward rules.
---

# fortimanager_object_icap_profile_respmodforwardrules
ICAP response mode forward rules.

~> This resource is a sub resource for variable `respmod_forward_rules` of resource `fortimanager_object_icap_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `header_group`: `fortimanager_object_icap_profile_respmodforwardrules_headergroup`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action to be taken for ICAP server. Valid values: `bypass`, `forward`.

* `header_group` - Header-Group. The structure of `header_group` block is documented below.
* `host` - Address object for the host.
* `http_resp_status_code` - HTTP response status code.
* `name` - Address name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `header_group` block supports:

* `case_sensitivity` - Enable/disable case sensitivity when matching header. Valid values: `disable`, `enable`.

* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIcap ProfileRespmodForwardRules can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_profile_respmodforwardrules.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
