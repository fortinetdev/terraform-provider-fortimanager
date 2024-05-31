---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extenderprofile_wifi"
description: |-
  FortiExtender wifi configuration.
---

# fortimanager_object_extensioncontroller_extenderprofile_wifi
FortiExtender wifi configuration.

~> This resource is a sub resource for variable `wifi` of resource `fortimanager_object_extensioncontroller_extenderprofile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `radio_1`: `fortimanager_object_extensioncontroller_extenderprofile_wifi_radio1`
>- `radio_2`: `fortimanager_object_extensioncontroller_extenderprofile_wifi_radio2`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.

* `dfs` - Wi-Fi 5G Radio DFS channel enable/disable. Valid values: `disable`, `enable`.

* `country` - Country in which this FEX will operate (default = NA). Valid values: `AL`, `DZ`, `AR`, `AM`, `AU`, `AT`, `AZ`, `BH`, `BD`, `BY`, `BE`, `BZ`, `BO`, `BA`, `BR`, `BN`, `BG`, `CA`, `CL`, `CN`, `CO`, `CR`, `HR`, `CY`, `CZ`, `DK`, `DO`, `EC`, `EG`, `SV`, `EE`, `FI`, `FR`, `GE`, `DE`, `GR`, `GT`, `HN`, `HK`, `HU`, `IS`, `IN`, `ID`, `IE`, `IL`, `IT`, `JM`, `JP`, `JO`, `KZ`, `KE`, `KR`, `KW`, `LV`, `LB`, `LI`, `LT`, `LU`, `MO`, `MK`, `MY`, `MT`, `MX`, `MC`, `MA`, `NP`, `NL`, `AN`, `NZ`, `NO`, `OM`, `PK`, `PA`, `PG`, `PE`, `PH`, `PL`, `PT`, `PR`, `QA`, `RO`, `RU`, `SA`, `SG`, `SK`, `SI`, `ZA`, `ES`, `LK`, `SE`, `CH`, `TW`, `TH`, `TT`, `TN`, `TR`, `AE`, `UA`, `GB`, `US`, `PS`, `UY`, `UZ`, `VE`, `VN`, `YE`, `ZW`, `NA`, `BS`, `VC`, `KH`, `MV`, `AF`, `NG`, `TZ`, `ZM`, `SN`, `CI`, `GH`, `CM`, `MW`, `AO`, `GA`, `ML`, `BJ`, `MG`, `TD`, `BW`, `LY`, `RW`, `MZ`, `GM`, `LS`, `MU`, `CG`, `UG`, `BF`, `SL`, `SO`, `CD`, `NE`, `CF`, `SZ`, `TG`, `LR`, `MR`, `DJ`, `RE`, `RS`, `ME`, `IQ`, `MD`, `KY`, `BB`, `BM`, `TC`, `VI`, `PM`, `MF`, `GD`, `IM`, `FO`, `GI`, `GL`, `TM`, `MN`, `VU`, `FJ`, `LA`, `GU`, `WF`, `MH`, `BT`, `FM`, `PF`, `NI`, `PY`, `HT`, `GY`, `AW`, `KN`, `GF`, `AS`, `MP`, `PW`, `MM`, `LC`, `GP`, `ET`, `SR`, `CX`, `DM`, `MQ`, `YT`, `BL`, `--`.

* `radio_1` - Radio-1. The structure of `radio_1` block is documented below.
* `radio_2` - Radio-2. The structure of `radio_2` block is documented below.

The `radio_1` block supports:

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


The `radio_2` block supports:

* `n80211d` - Enable/disable Wi-Fi 802.11d. Valid values: `disable`, `enable`.

* `band` - Wi-Fi band selection 2.4GHz / 5GHz. Valid values: `5GHz`.

* `bandwidth` - Wi-Fi channel bandwidth. Valid values: `auto`, `20MHz`, `40MHz`, `80MHz`.

* `beacon_interval` - Wi-Fi beacon interval in miliseconds (100 - 3500, default = 100).
* `bss_color` - Wi-Fi 802.11AX BSS color value (0 - 63, 0 = disable, default = 0).
* `bss_color_mode` - Wi-Fi 802.11AX BSS color mode. Valid values: `auto`, `static`.

* `channel` - Wi-Fi channels. Valid values: `CH36`, `CH40`, `CH44`, `CH48`, `CH52`, `CH56`, `CH60`, `CH64`, `CH100`, `CH104`, `CH108`, `CH112`, `CH116`, `CH120`, `CH124`, `CH128`, `CH132`, `CH136`, `CH140`, `CH144`, `CH149`, `CH153`, `CH157`, `CH161`, `CH165`.

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

ObjectExtensionController ExtenderProfileWifi can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extenderprofile_wifi.labelname ObjectExtensionControllerExtenderProfileWifi
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
