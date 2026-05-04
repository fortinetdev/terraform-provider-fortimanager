---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_shapingprofile_classes"
description: |-
  ObjectFirewall ShapingProfileClasses
---

# fortimanager_object_firewall_shapingprofile_classes
ObjectFirewall ShapingProfileClasses

~> This resource is a sub resource for variable `classes` of resource `fortimanager_object_firewall_shapingprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `shaping_profile` - Shaping Profile.

* `class_id` - Class-Id.
* `guaranteed_bandwidth` - Guaranteed-Bandwidth.
* `maximum_bandwidth` - Maximum-Bandwidth.
* `name` - Name.
* `priority` - Priority. Valid values: `top`, `critical`, `high`, `medium`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ShapingProfileClasses can be imported using any of these accepted formats:
```
Set import_options = ["shaping_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_shapingprofile_classes.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
