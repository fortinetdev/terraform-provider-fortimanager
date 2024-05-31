---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extenderprofile_wifi_radio1"
description: |-
  Radio-1 config for Wi-Fi 2.4GHz
---

# fortimanager_object_extensioncontroller_extenderprofile_wifi_radio1
Radio-1 config for Wi-Fi 2.4GHz

~> This resource is a sub resource for variable `radio_1` of resource `fortimanager_object_extensioncontroller_extenderprofile_wifi`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `n80211d` - Enable/disable Wi-Fi 802.11d. Valid values: `disable`, `enable`.

* `band` - Wi-Fi band selection 2.4GHz / 5GHz. Valid values: `2.4GHz`.

* `bandwidth` - Wi-Fi channel bandwidth. Valid values: `auto`, `20MHz`, `40MHz`, `80MHz`.

* `beacon_interval` - Wi-Fi beacon interval in miliseconds (100 - 3500, default = 100).
* `bss_color` - Wi-Fi 802.11AX BSS color value (0 - 63, 0 = disable, default = 0).
* `bss_color_mode` - Wi-Fi 802.11AX BSS color mode. Valid values: `auto`, `static`.

* `channel` - Wi-Fi channels. Valid values: `CH1`, `CH2`, `CH3`, `CH4`, `CH5`, `CH6`, `CH7`, `CH8`, `CH9`, `CH10`, `CH11`.

* `extension_channel` - Wi-Fi extension channel. Valid values: `auto`, `higher`, `lower`.

* `guard_interval` - Wi-Fi guard interval. Valid values: `auto`, `400ns`, `800ns`.

* `lan_ext_vap` - Wi-Fi LAN-Extention VAP. Select only one VAP.
* `local_vaps` - Wi-Fi local VAP. Select up to three VAPs.
* `max_clients` - Maximum number of Wi-Fi radio clients (0 - 512, 0 = unlimited, default = 0).
* `mode` - Wi-Fi radio mode AP(LAN mode) / Client(WAN mode). Valid values: `AP`, `Client`.

* `operating_standard` - Wi-Fi operating standard. Valid values: `auto`, `11A-N-AC-AX`, `11A-N-AC`, `11A-N`, `11A`, `11N-AC-AX`, `11AC-AX`, `11AC`, `11N-AC`, `11B-G-N-AX`, `11B-G-N`, `11B-G`, `11B`, `11G-N-AX`, `11N-AX`, `11AX`, `11G-N`, `11N`, `11G`.

* `power_level` - Wi-Fi power level in percent (0 - 100, 0 = auto, default = 100).
* `radio_id` - Radio ID.
* `status` - Enable/disable Wi-Fi radio. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectExtensionController ExtenderProfileWifiRadio1 can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extenderprofile_wifi_radio1.labelname ObjectExtensionControllerExtenderProfileWifiRadio1
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
