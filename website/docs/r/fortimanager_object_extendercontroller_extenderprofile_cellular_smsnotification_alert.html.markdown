---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_alert"
description: |-
  SMS alert list.
---

# fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_alert
SMS alert list.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `data_exhausted` - Display string when data exhausted.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.
* `low_signal_strength` - Display string when signal strength is low.
* `mode_switch` - Display string when mode is switched.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `session_disconnect` - Display string when session disconnected.
* `system_reboot` - Display string when system rebooted.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectExtenderController ExtenderProfileCellularSmsNotificationAlert can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_alert.labelname ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
