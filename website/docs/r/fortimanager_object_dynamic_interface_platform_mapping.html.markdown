---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dynamic_interface_platform_mapping"
description: |-
  ObjectDynamic InterfacePlatformMapping
---

# fortimanager_object_dynamic_interface_platform_mapping
ObjectDynamic InterfacePlatformMapping

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `interface` - Interface.

* `egress_shaping_profile` - Egress-Shaping-Profile.
* `ingress_shaping_profile` - Ingress-Shaping-Profile.
* `intf_zone` - Intf-Zone.
* `intrazone_deny` - Intrazone-Deny. Valid values: `disable`, `enable`.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDynamic InterfacePlatformMapping can be imported using any of these accepted formats:
```
Set import_options = ["interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dynamic_interface_platform_mapping.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
