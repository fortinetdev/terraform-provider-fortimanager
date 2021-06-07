---
subcategory: "Object Extender-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_sim_profile"
description: |-
  ObjectExtenderController SimProfile
---

# fortimanager_object_extendercontroller_sim_profile
ObjectExtenderController SimProfile

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auto_switch_profile` - Auto-Switch_Profile. The structure of `auto_switch_profile` block is documented below.
* `conn_status` - Conn-Status.
* `default_sim` - Default-Sim. Valid values: `sim1`, `sim2`, `carrier`, `cost`.

* `description` - Description.
* `gps` - Gps. Valid values: `disable`, `enable`.

* `modem_id` - Modem-Id.
* `name` - Name.
* `preferred_carrier` - Preferred-Carrier.
* `redundant_intf` - Redundant-Intf.
* `redundant_mode` - Redundant-Mode. Valid values: `disable`, `enable`.

* `sim1_pin` - Sim1-Pin. Valid values: `disable`, `enable`.

* `sim1_pin_code` - Sim1-Pin-Code.
* `sim2_pin` - Sim2-Pin. Valid values: `disable`, `enable`.

* `sim2_pin_code` - Sim2-Pin-Code.
* `status` - Status. Valid values: `disable`, `enable`.


The `auto_switch_profile` block supports:

* `dataplan` - Dataplan. Valid values: `disable`, `enable`.

* `disconnect` - Disconnect. Valid values: `disable`, `enable`.

* `disconnect_period` - Disconnect-Period.
* `disconnect_threshold` - Disconnect-Threshold.
* `signal` - Signal. Valid values: `disable`, `enable`.

* `status` - Status. Valid values: `disable`, `enable`.

* `switch_back` - Switch-Back. Valid values: `time`, `timer`.

* `switch_back_time` - Switch-Back-Time.
* `switch_back_timer` - Switch-Back-Timer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtenderController SimProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_sim_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
