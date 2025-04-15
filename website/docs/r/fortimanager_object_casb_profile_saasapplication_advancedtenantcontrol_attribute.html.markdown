---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_profile_saasapplication_advancedtenantcontrol_attribute"
description: |-
  CASB advanced tenant control attribute.
---

# fortimanager_object_casb_profile_saasapplication_advancedtenantcontrol_attribute
CASB advanced tenant control attribute.

~> This resource is a sub resource for variable `attribute` of resource `fortimanager_object_casb_profile_saasapplication_advancedtenantcontrol`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.
* `saas_application` - Saas Application.
* `advanced_tenant_control` - Advanced Tenant Control.

* `input` - CASB extend user input value.
* `name` - CASB extend user input name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCasb ProfileSaasApplicationAdvancedTenantControlAttribute can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE", "saas_application=YOUR_VALUE", "advanced_tenant_control=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_profile_saasapplication_advancedtenantcontrol_attribute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
