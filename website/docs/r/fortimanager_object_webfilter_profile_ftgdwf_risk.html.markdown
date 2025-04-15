---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_ftgdwf_risk"
description: |-
  FortiGuard risk level settings.
---

# fortimanager_object_webfilter_profile_ftgdwf_risk
FortiGuard risk level settings.

~> This resource is a sub resource for variable `risk` of resource `fortimanager_object_webfilter_profile_ftgdwf`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action to take for matches. Valid values: `block`, `monitor`.

* `fosid` - ID number.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `risk_level` - Risk level to be examined.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectWebfilter ProfileFtgdWfRisk can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_ftgdwf_risk.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
