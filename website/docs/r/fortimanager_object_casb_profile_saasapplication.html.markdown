---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_profile_saasapplication"
description: |-
  CASB profile SaaS application.
---

# fortimanager_object_casb_profile_saasapplication
CASB profile SaaS application.

~> This resource is a sub resource for variable `saas_application` of resource `fortimanager_object_casb_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `access_rule`: `fortimanager_object_casb_profile_saasapplication_accessrule`
>- `advanced_tenant_control`: `fortimanager_object_casb_profile_saasapplication_advancedtenantcontrol`
>- `custom_control`: `fortimanager_object_casb_profile_saasapplication_customcontrol`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `access_rule` - Access-Rule. The structure of `access_rule` block is documented below.
* `advanced_tenant_control` - Advanced-Tenant-Control. The structure of `advanced_tenant_control` block is documented below.
* `custom_control` - Custom-Control. The structure of `custom_control` block is documented below.
* `domain_control` - Enable/disable domain control. Valid values: `disable`, `enable`.

* `domain_control_domains` - CASB profile domain control domains.
* `log` - Enable/disable log settings. Valid values: `disable`, `enable`.

* `name` - CASB profile SaaS application name.
* `safe_search` - Enable/disable safe search. Valid values: `disable`, `enable`.

* `safe_search_control` - CASB profile safe search control.
* `status` - Enable/disable setting. Valid values: `disable`, `enable`.

* `tenant_control` - Enable/disable tenant control. Valid values: `disable`, `enable`.

* `tenant_control_tenants` - CASB profile tenant control tenants.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `access_rule` block supports:

* `action` - CASB access rule action. Valid values: `block`, `bypass`, `monitor`.

* `attribute_filter` - Attribute-Filter. The structure of `attribute_filter` block is documented below.
* `bypass` - CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.

* `name` - CASB access rule activity name.

The `attribute_filter` block supports:

* `action` - CASB access rule tenant control action. Valid values: `block`, `monitor`, `bypass`.

* `attribute_match` - CASB access rule tenant match.
* `id` - CASB tenant control ID.

The `advanced_tenant_control` block supports:

* `attribute` - Attribute. The structure of `attribute` block is documented below.
* `name` - CASB advanced tenant control name.

The `attribute` block supports:

* `input` - CASB extend user input value.
* `name` - CASB extend user input name.

The `custom_control` block supports:

* `attribute_filter` - Attribute-Filter. The structure of `attribute_filter` block is documented below.
* `name` - CASB custom control user activity name.
* `option` - Option. The structure of `option` block is documented below.

The `attribute_filter` block supports:

* `action` - CASB access rule tenant control action. Valid values: `block`, `monitor`, `bypass`.

* `attribute_match` - CASB access rule tenant match.
* `id` - CASB tenant control ID.

The `option` block supports:

* `name` - CASB custom control option name.
* `user_input` - CASB custom control user input.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb ProfileSaasApplication can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_profile_saasapplication.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
