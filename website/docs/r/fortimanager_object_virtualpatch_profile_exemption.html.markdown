---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_virtualpatch_profile_exemption"
description: |-
  Exempt devices or rules.
---

# fortimanager_object_virtualpatch_profile_exemption
Exempt devices or rules.

~> This resource is a sub resource for variable `exemption` of resource `fortimanager_object_virtualpatch_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `device` - Device MAC addresses.
* `fosid` - IDs.
* `rule` - Patch signature rule IDs.
* `status` - Enable/disable exemption. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVirtualPatch ProfileExemption can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_virtualpatch_profile_exemption.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
