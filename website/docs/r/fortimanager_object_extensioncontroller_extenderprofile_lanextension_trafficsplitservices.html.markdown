---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extenderprofile_lanextension_trafficsplitservices"
description: |-
  Config FortiExtender traffic split interface for LAN extension.
---

# fortimanager_object_extensioncontroller_extenderprofile_lanextension_trafficsplitservices
Config FortiExtender traffic split interface for LAN extension.

~> This resource is a sub resource for variable `traffic_split_services` of resource `fortimanager_object_extensioncontroller_extenderprofile_lanextension`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `address` - Address selection.
* `name` - FortiExtender LAN extension tunnel split entry name.
* `service` - Service selection.
* `vsdb` - Select vsdb [enable/disable]. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtensionController ExtenderProfileLanExtensionTrafficSplitServices can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extenderprofile_lanextension_trafficsplitservices.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
