---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wtpprofile_radio3"
description: |-
  Configuration options for radio 3.
---

# fortimanager_object_wirelesscontroller_wtpprofile_radio3
Configuration options for radio 3.

~> This resource is a sub resource for variable `radio_3` of resource `fortimanager_object_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wtp_profile` - Wtp Profile.

* `n80211d` - Enable/disable 802.11d countryie(default = enable). Valid values: `disable`, `enable`.

* `airtime_fairness` - Enable/disable airtime fairness (default = disable). Valid values: `disable`, `enable`.

* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `disable`, `enable`.

* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `disable`, `enable`.

* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `disable`, `enable`.

* `arrp_profile` - Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `disable`, `enable`.

* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - The target of automatic transmit power adjustment in dBm. (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on. Valid values: `802.11b`, `802.11a`, `802.11g`, `802.11n`, `802.11ac`, `802.11n-5G`, `802.11ax-5G`, `802.11ax`, `802.11ac-2G`, `802.11g-only`, `802.11n-only`, `802.11n,g-only`, `802.11ac-only`, `802.11ac,n-only`, `802.11n-5G-only`, `802.11ax-5G-only`, `802.11ax,ac-only`, `802.11ax,ac,n-only`, `802.11ax-only`, `802.11ax,n-only`, `802.11ax,n,g-only`.

* `band_5g_type` - WiFi 5G band type. Valid values: `5g-full`, `5g-high`, `5g-low`.

* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `disable`, `enable`.

* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in msec (the actual range of beacon interval depends on the AP platform type, default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, 0 means disable. default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto). Valid values: `auto`, `static`.

* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `disable`, `enable`.

* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel` - Selected list of wireless radio channels.
* `channel_bonding` - Channel bandwidth: 160,80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `80MHz`, `40MHz`, `20MHz`, `160MHz`.

* `channel_utilization` - Enable/disable measuring channel utilization. Valid values: `disable`, `enable`.

* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `disable`, `enable`.

* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `disable`, `enable`.

* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable`, `enable`.

* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low`, `medium`, `high`.

* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `iperf_protocol` - Iperf test protocol (default = "UDP"). Valid values: `udp`, `tcp`.

* `iperf_server_port` - Iperf service port number.
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `disable`, `enable`.

* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mimo_mode` - Configure radio MIMO mode (default = default). Valid values: `default`, `1x1`, `2x2`, `3x3`, `4x4`, `8x8`.

* `mode` - Mode of radio 3. Radio 3 can be disabled, configured as an access point, a rogue AP monitor, or a sniffer. Valid values: `disabled`, `ap`, `monitor`, `sniffer`.

* `optional_antenna` - Optional antenna used on FAP (default = none). Valid values: `none`, `FANT-04ABGN-0606-O-N`, `FANT-04ABGN-1414-P-N`, `FANT-04ABGN-8065-P-N`, `FANT-04ABGN-0606-O-R`, `FANT-04ABGN-0606-P-R`, `FANT-10ACAX-1213-D-N`, `FANT-08ABGN-1213-D-R`.

* `power_level` - Radio power level as a percentage of the maximum transmit power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm`, `percentage`.

* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim`, `ac-vo`, `no-obss-scan`, `no-11b-rate`, `client-rate-follow`.

* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts`, `ctsonly`, `disable`.

* `radio_id` - Radio-Id.
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable). Valid values: `disable`, `enable`.

* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_password` - Passphrase for WiFi network connection.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal"). Valid values: `open`, `wpa-personal`, `wpa-enterprise`.

* `sam_server` - SAM test server IP address or domain name.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP"). Valid values: `ip`, `fqdn`.

* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING"). Valid values: `ping`, `iperf`.

* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `disable`, `enable`.

* `spectrum_analysis` - Enable/disable spectrum analysis to find interference that would negatively impact wireless performance. Valid values: `disable`, `enable`, `scan-only`.

* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable`, `power-save`, `aggr-limit`, `retry-limit`, `send-bar`.

* `vap_all` - Enable/disable the automatic inheritance of all Virtual Access Points (VAPs) (default = enable). Valid values: `disable`, `enable`, `tunnel`, `bridge`, `manual`.

* `vap1` - Virtual Access Point (VAP) for wlan ID 1
* `vap2` - Virtual Access Point (VAP) for wlan ID 2
* `vap3` - Virtual Access Point (VAP) for wlan ID 3
* `vap4` - Virtual Access Point (VAP) for wlan ID 4
* `vap5` - Virtual Access Point (VAP) for wlan ID 5
* `vap6` - Virtual Access Point (VAP) for wlan ID 6
* `vap7` - Virtual Access Point (VAP) for wlan ID 7
* `vap8` - Virtual Access Point (VAP) for wlan ID 8
* `vaps` - Manually selected list of Virtual Access Points (VAPs).
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWirelessController WtpProfileRadio3 can be imported using any of these accepted formats:
```
Set import_options = ["wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wtpprofile_radio3.labelname ObjectWirelessControllerWtpProfileRadio3
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
