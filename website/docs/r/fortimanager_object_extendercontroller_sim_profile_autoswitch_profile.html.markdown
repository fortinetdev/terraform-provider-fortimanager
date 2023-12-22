---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_sim_profile_autoswitch_profile"
description: |-
  ObjectExtenderController SimProfileAutoSwitchProfile
---

# fortimanager_object_extendercontroller_sim_profile_autoswitch_profile
ObjectExtenderController SimProfileAutoSwitchProfile

~> This resource is a sub resource for variable `auto_switch_profile` of resource `fortimanager_object_extendercontroller_sim_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_extendercontroller_sim_profile_autoswitch_profile" "trname" {
  sim_profile = fortimanager_object_extendercontroller_sim_profile.trname.name
  dataplan    = "disable"
  disconnect  = "enable"
  depends_on  = [fortimanager_object_extendercontroller_sim_profile.trname]
}

resource "fortimanager_object_extendercontroller_sim_profile" "trname" {
  name = "terr-ex-ctl-sim-profile"
  auto_switch_profile {
    dataplan = "disable"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sim_profile` - Sim Profile.

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
* `id` - an identifier for the resource.

## Import

ObjectExtenderController SimProfileAutoSwitchProfile can be imported using any of these accepted formats:
```
Set import_options = ["sim_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_sim_profile_autoswitch_profile.labelname ObjectExtenderControllerSimProfileAutoSwitchProfile
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
