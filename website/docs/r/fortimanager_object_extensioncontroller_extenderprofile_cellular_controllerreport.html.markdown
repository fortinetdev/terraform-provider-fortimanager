---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extenderprofile_cellular_controllerreport"
description: |-
  FortiExtender controller report configuration.
---

# fortimanager_object_extensioncontroller_extenderprofile_cellular_controllerreport
FortiExtender controller report configuration.

~> This resource is a sub resource for variable `controller_report` of resource `fortimanager_object_extensioncontroller_extenderprofile_cellular`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectExtensionController ExtenderProfileCellularControllerReport can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extenderprofile_cellular_controllerreport.labelname ObjectExtensionControllerExtenderProfileCellularControllerReport
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
