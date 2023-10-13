---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_profile_saasapplication"
description: |-
  CASB profile SaaS application.
---

# fortimanager_object_casb_profile_saasapplication
CASB profile SaaS application.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `access_rule` - Access-Rule. The structure of `access_rule` block is documented below.
* `custom_control` - Custom-Control. The structure of `custom_control` block is documented below.
* `domain_control` - Enable/disable domain control. Valid values: `disable`, `enable`.

* `domain_control_domains` - CASB profile domain control domains.
* `log` - Enable/disable log settings. Valid values: `disable`, `enable`.

* `name` - CASB profile SaaS application name.
* `safe_search` - Enable/disable safe search. Valid values: `disable`, `enable`.

* `safe_search_control` - CASB profile safe search control.
* `tenant_control` - Enable/disable tenant control. Valid values: `disable`, `enable`.

* `tenant_control_tenants` - CASB profile tenant control tenants.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `access_rule` block supports:

* `action` - CASB access rule action. Valid values: `block`, `bypass`, `monitor`.

* `bypass` - CASB bypass options. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.

* `name` - CASB access rule activity name.

The `custom_control` block supports:

* `name` - CASB custom control user activity name.
* `option` - Option. The structure of `option` block is documented below.

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
