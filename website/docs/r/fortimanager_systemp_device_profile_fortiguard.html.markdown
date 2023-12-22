---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_device_profile_fortiguard"
description: |-
  Systemp DeviceProfileFortiguard
---

# fortimanager_systemp_device_profile_fortiguard
Systemp DeviceProfileFortiguard

## Example Usage

```hcl
resource "fortimanager_systemp_device_profile_fortiguard" "trname" {
  devprof               = "default"
  auto_firmware_upgrade = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `auto_firmware_upgrade` - Auto-Firmware-Upgrade. Valid values: `disable`, `enable`.

* `auto_firmware_upgrade_day` - Auto-Firmware-Upgrade-Day. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `auto_firmware_upgrade_delay` - Auto-Firmware-Upgrade-Delay.
* `auto_firmware_upgrade_end_hour` - Auto-Firmware-Upgrade-End-Hour.
* `auto_firmware_upgrade_start_hour` - Auto-Firmware-Upgrade-Start-Hour.
* `target` - Target. Valid values: `none`, `direct`, `this-fmg`.

* `target_ip` - Target-Ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp DeviceProfileFortiguard can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_device_profile_fortiguard.labelname SystempDeviceProfileFortiguard
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
