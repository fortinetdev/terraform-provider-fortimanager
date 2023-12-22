---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_casbprofile_saasapplication_customcontrol"
description: |-
  ObjectFirewall CasbProfileSaasApplicationCustomControl
---

# fortimanager_object_firewall_casbprofile_saasapplication_customcontrol
ObjectFirewall CasbProfileSaasApplicationCustomControl

~> This resource is a sub resource for variable `custom_control` of resource `fortimanager_object_firewall_casbprofile_saasapplication`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`option`: `fortimanager_object_firewall_casbprofile_saasapplication_customcontrol_option`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `casb_profile` - Casb Profile.
* `saas_application` - Saas Application.

* `name` - Name.
* `option` - Option. The structure of `option` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `option` block supports:

* `name` - Name.
* `user_input` - User-Input.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall CasbProfileSaasApplicationCustomControl can be imported using any of these accepted formats:
```
Set import_options = ["casb_profile=YOUR_VALUE", "saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_casbprofile_saasapplication_customcontrol.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
