---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extensioncontroller_extenderprofile"
description: |-
  FortiExtender extender profile configuration.
---

# fortimanager_object_extensioncontroller_extenderprofile
FortiExtender extender profile configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cellular`: `fortimanager_object_extensioncontroller_extenderprofile_cellular`
>- `lan_extension`: `fortimanager_object_extensioncontroller_extenderprofile_lanextension`
>- `wifi`: `fortimanager_object_extensioncontroller_extenderprofile_wifi`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_is_factory_setting` - _Is_Factory_Setting. Valid values: `disable`, `enable`, `ext`.

* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`.

* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `cellular` - Cellular. The structure of `cellular` block is documented below.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `disable`, `enable`.

* `extension` - Extension option. Valid values: `wan-extension`, `lan-extension`.

* `fosid` - ID.
* `lan_extension` - Lan-Extension. The structure of `lan_extension` block is documented below.
* `login_password` - Set the managed extender's administrator password.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `no`, `yes`, `default`.

* `model` - Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.

* `name` - FortiExtender profile name.
* `wifi` - Wifi. The structure of `wifi` block is documented below.

The `cellular` block supports:

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
* `phone_number` - Receiver phone number. Format: [+][country code][area code][local phone number]. For example, +16501234567.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.


The `lan_extension` block supports:

* `backhaul` - Backhaul. The structure of `backhaul` block is documented below.
* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `downlinks` - Downlinks. The structure of `downlinks` block is documented below.
* `ipsec_tunnel` - IPsec tunnel name.
* `link_loadbalance` - LAN extension link load balance strategy. Valid values: `activebackup`, `loadbalance`.


The `backhaul` block supports:

* `name` - FortiExtender LAN extension backhaul name.
* `port` - FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.

* `role` - FortiExtender uplink port. Valid values: `primary`, `secondary`.

* `weight` - WRR weight parameter.

The `downlinks` block supports:

* `name` - FortiExtender LAN extension downlink config entry name.
* `port` - FortiExtender LAN extension downlink port. Valid values: `port1`, `port2`, `port3`, `port4`, `port5`, `lan1`, `lan2`.

* `pvid` - FortiExtender LAN extension downlink PVID.
* `type` - FortiExtender LAN extension downlink type [port/vap]. Valid values: `port`, `vap`.

* `vap` - FortiExtender LAN extension downlink vap.

The `wifi` block supports:

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
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtensionController ExtenderProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extensioncontroller_extenderprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
