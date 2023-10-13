---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_cellular_modem1"
description: |-
  Configuration options for modem 1.
---

# fortimanager_object_extendercontroller_extenderprofile_cellular_modem1
Configuration options for modem 1.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectExtenderController ExtenderProfileCellularModem1 can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_extenderprofile_cellular_modem1.labelname ObjectExtenderControllerExtenderProfileCellularModem1
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
