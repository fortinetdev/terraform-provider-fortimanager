---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fmg_fabric_authorization_template_platforms"
description: |-
  ObjectFmg FabricAuthorizationTemplatePlatforms
---

# fortimanager_object_fmg_fabric_authorization_template_platforms
ObjectFmg FabricAuthorizationTemplatePlatforms

~> This resource is a sub resource for variable `platforms` of resource `fortimanager_object_fmg_fabric_authorization_template`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `template` - Template.

* `fmgcount` - Count.
* `extension_type` - Extension-Type. Valid values: `wan-extension`, `lan-extension`.

* `fortilink` - Fortilink.
* `prefix` - Prefix.
* `type` - Type. Valid values: `ap`, `extender`, `switch`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{prefix}}.

## Import

ObjectFmg FabricAuthorizationTemplatePlatforms can be imported using any of these accepted formats:
```
Set import_options = ["template=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fmg_fabric_authorization_template_platforms.labelname {{prefix}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
