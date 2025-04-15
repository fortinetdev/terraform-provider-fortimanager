---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_profile_saasapplication_accessrule"
description: |-
  CASB profile access rule.
---

# fortimanager_object_casb_profile_saasapplication_accessrule
CASB profile access rule.

~> This resource is a sub resource for variable `access_rule` of resource `fortimanager_object_casb_profile_saasapplication`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `attribute_filter`: `fortimanager_object_casb_profile_saasapplication_accessrule_attributefilter`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.
* `saas_application` - Saas Application.

* `action` - CASB access rule action. Valid values: `block`, `bypass`, `monitor`.

* `attribute_filter` - Attribute-Filter. The structure of `attribute_filter` block is documented below.
* `bypass` - CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.

* `name` - CASB access rule activity name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `attribute_filter` block supports:

* `action` - CASB access rule tenant control action. Valid values: `block`, `monitor`, `bypass`.

* `attribute_match` - CASB access rule tenant match.
* `id` - CASB tenant control ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb ProfileSaasApplicationAccessRule can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE", "saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_profile_saasapplication_accessrule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
