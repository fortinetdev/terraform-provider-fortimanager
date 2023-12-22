---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_casbprofile_saasapplication_accessrule"
description: |-
  ObjectFirewall CasbProfileSaasApplicationAccessRule
---

# fortimanager_object_firewall_casbprofile_saasapplication_accessrule
ObjectFirewall CasbProfileSaasApplicationAccessRule

~> This resource is a sub resource for variable `access_rule` of resource `fortimanager_object_firewall_casbprofile_saasapplication`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `casb_profile` - Casb Profile.
* `saas_application` - Saas Application.

* `action` - Action. Valid values: `block`, `monitor`, `bypass`.

* `bypass` - Bypass. Valid values: `av`, `dlp`, `web-filter`, `file-filter`, `video-filter`.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall CasbProfileSaasApplicationAccessRule can be imported using any of these accepted formats:
```
Set import_options = ["casb_profile=YOUR_VALUE", "saas_application=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_casbprofile_saasapplication_accessrule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
