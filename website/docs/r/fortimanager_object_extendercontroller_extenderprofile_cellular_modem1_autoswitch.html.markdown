---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_cellular_modem1_autoswitch"
description: |-
  FortiExtender auto switch configuration.
---

# fortimanager_object_extendercontroller_extenderprofile_cellular_modem1_autoswitch
FortiExtender auto switch configuration.

~> This resource is a sub resource for variable `auto_switch` of resource `fortimanager_object_extendercontroller_extenderprofile_cellular_modem1`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_extendercontroller_extenderprofile_cellular_modem1_autoswitch" "trname" {
  dataplan             = "enable"
  disconnect           = "enable"
  disconnect_period    = 700
  disconnect_threshold = 80
  extender_profile     = fortimanager_object_extendercontroller_extenderprofile.trname.name
  depends_on           = [fortimanager_object_extendercontroller_extenderprofile.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile" "trname" {
  name = "terr-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

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

ObjectExtenderController ExtenderProfileCellularModem1AutoSwitch can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_extenderprofile_cellular_modem1_autoswitch.labelname ObjectExtenderControllerExtenderProfileCellularModem1AutoSwitch
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
