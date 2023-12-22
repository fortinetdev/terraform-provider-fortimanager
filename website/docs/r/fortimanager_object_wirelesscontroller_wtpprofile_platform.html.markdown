---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wtpprofile_platform"
description: |-
  WTP, FortiAP, or AP platform.
---

# fortimanager_object_wirelesscontroller_wtpprofile_platform
WTP, FortiAP, or AP platform.

~> This resource is a sub resource for variable `platform` of resource `fortimanager_object_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_wtpprofile_platform" "trname" {
  _local_platform_str = "platfor"
  depends_on          = [fortimanager_object_wirelesscontroller_wtpprofile.trname6]
  wtp_profile         = fortimanager_object_wirelesscontroller_wtpprofile.trname6.name
}

resource "fortimanager_object_wirelesscontroller_wtpprofile" "trname6" {
  name = "terr-wtpprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wtp_profile` - Wtp Profile.

* `_local_platform_str` - _Local_Platform_Str.
* `ddscan` - Enable/disable use of one radio for dedicated dual-band scanning to detect RF characterization and wireless threat management. Valid values: `disable`, `enable`.

* `mode` - Configure operation mode of 5G radios (default = single-5G). Valid values: `dual-5G`, `single-5G`.

* `type` - WTP, FortiAP or AP platform type. There are built-in WTP profiles for all supported FortiAP models. You can select a built-in profile and customize it or create a new profile. Valid values: `30B-50B`, `60B`, `80CM-81CM`, `220A`, `220B`, `210B`, `60C`, `222B`, `112B`, `320B`, `11C`, `14C`, `223B`, `28C`, `320C`, `221C`, `25D`, `222C`, `224D`, `214B`, `21D`, `24D`, `112D`, `223C`, `321C`, `C220C`, `C225C`, `S321C`, `S323C`, `FWF`, `S311C`, `S313C`, `AP-11N`, `S322C`, `S321CR`, `S322CR`, `S323CR`, `S421E`, `S422E`, `S423E`, `421E`, `423E`, `C221E`, `C226E`, `C23JD`, `C24JE`, `C21D`, `U421E`, `U423E`, `221E`, `222E`, `223E`, `S221E`, `S223E`, `U221EV`, `U223EV`, `U321EV`, `U323EV`, `224E`, `U422EV`, `U24JEV`, `321E`, `U431F`, `U433F`, `231E`, `431F`, `433F`, `231F`, `432F`, `234F`, `23JF`, `U231F`, `831F`, `U234F`, `U432F`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWirelessController WtpProfilePlatform can be imported using any of these accepted formats:
```
Set import_options = ["wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wtpprofile_platform.labelname ObjectWirelessControllerWtpProfilePlatform
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
