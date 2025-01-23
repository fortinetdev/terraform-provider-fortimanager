---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_cellular"
description: |-
  FortiExtender cellular configuration.
---

# fortimanager_object_extendercontroller_extenderprofile_cellular
FortiExtender cellular configuration.

~> This resource is a sub resource for variable `cellular` of resource `fortimanager_object_extendercontroller_extenderprofile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `controller_report`: `fortimanager_object_extendercontroller_extenderprofile_cellular_controllerreport`
>- `modem1`: `fortimanager_object_extendercontroller_extenderprofile_cellular_modem1`
>- `modem2`: `fortimanager_object_extendercontroller_extenderprofile_cellular_modem2`
>- `sms_notification`: `fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `controller_report` - Controller-Report. The structure of `controller_report` block is documented below.
* `dataplan` - Dataplan names.
* `modem1` - Modem1. The structure of `modem1` block is documented below.
* `modem2` - Modem2. The structure of `modem2` block is documented below.
* `sms_notification` - Sms-Notification. The structure of `sms_notification` block is documented below.

The `controller_report` block supports:

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.


The `modem1` block supports:

* `auto_switch` - Auto-Switch. The structure of `auto_switch` block is documented below.
* `conn_status` - Conn-Status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.

* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.

* `modem_id` - Modem ID.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.

* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.

* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.

* `sim2_pin_code` - SIM #2 PIN password.

The `auto_switch` block supports:

* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.

* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.

* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.

* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `modem2` block supports:

* `auto_switch` - Auto-Switch. The structure of `auto_switch` block is documented below.
* `conn_status` - Conn-Status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.

* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.

* `modem_id` - Modem ID.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.

* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.

* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.

* `sim2_pin_code` - SIM #2 PIN password.

The `auto_switch` block supports:

* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.

* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.

* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.

* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `sms_notification` block supports:

* `alert` - Alert. The structure of `alert` block is documented below.
* `receiver` - Receiver. The structure of `receiver` block is documented below.
* `status` - FortiExtender SMS notification status. Valid values: `disable`, `enable`.


The `alert` block supports:

* `data_exhausted` - Display string when data exhausted.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.
* `low_signal_strength` - Display string when signal strength is low.
* `mode_switch` - Display string when mode is switched.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `session_disconnect` - Display string when session disconnected.
* `system_reboot` - Display string when system rebooted.

The `receiver` block supports:

* `alert` - Alert multi-options. Valid values: `system-reboot`, `data-exhausted`, `session-disconnect`, `low-signal-strength`, `mode-switch`, `os-image-fallback`, `fgt-backup-mode-switch`.

* `name` - FortiExtender SMS notification receiver name.
* `phone_number` - Receiver phone number.  Format: [+][country code][area code][local phone number].  For example: +16501234567.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectExtenderController ExtenderProfileCellular can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_extenderprofile_cellular.labelname ObjectExtenderControllerExtenderProfileCellular
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
