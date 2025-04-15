---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_casb_profile_saasapplication_customcontrol_attributefilter"
description: |-
  CASB attribute filter.
---

# fortimanager_object_casb_profile_saasapplication_customcontrol_attributefilter
CASB attribute filter.

~> This resource is a sub resource for variable `attribute_filter` of resource `fortimanager_object_casb_profile_saasapplication_customcontrol`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.
* `saas_application` - Saas Application.
* `custom_control` - Custom Control.

* `action` - CASB access rule tenant control action. Valid values: `block`, `monitor`, `bypass`.

* `attribute_match` - CASB access rule tenant match.
* `fosid` - CASB tenant control ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectCasb ProfileSaasApplicationCustomControlAttributeFilter can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE", "saas_application=YOUR_VALUE", "custom_control=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_casb_profile_saasapplication_customcontrol_attributefilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
