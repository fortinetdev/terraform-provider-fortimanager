---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_device_profile_fortianalyzer"
description: |-
  Systemp DeviceProfileFortianalyzer
---

# fortimanager_systemp_device_profile_fortianalyzer
Systemp DeviceProfileFortianalyzer

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `managed_sn` - Managed-Sn.
* `target` - Target. Valid values: `none`, `this-fmg`, `managed`, `others`.

* `target_ip` - Target-Ip.
* `target_sn` - Target-Sn.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp DeviceProfileFortianalyzer can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_device_profile_fortianalyzer.labelname SystempDeviceProfileFortianalyzer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
