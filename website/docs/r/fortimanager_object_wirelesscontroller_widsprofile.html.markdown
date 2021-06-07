---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_widsprofile"
description: |-
  Configure wireless intrusion detection system (WIDS) profiles.
---

# fortimanager_object_wirelesscontroller_widsprofile
Configure wireless intrusion detection system (WIDS) profiles.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ap_auto_suppress` - Enable/disable on-wire rogue AP auto-suppression (default = disable). Valid values: `disable`, `enable`.

* `ap_bgscan_disable_schedules` - Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space.
* `ap_bgscan_duration` - Listening time on a scanning channel (10 - 1000 msec, default = 20).
* `ap_bgscan_idle` - Waiting time for channel inactivity before scanning this channel (0 - 1000 msec, default = 0).
* `ap_bgscan_intv` - Period of time between scanning two channels (1 - 600 sec, default = 1).
* `ap_bgscan_period` - Period of time between background scans (10 - 3600 sec, default = 600).
* `ap_bgscan_report_intv` - Period of time between background scan reports (15 - 600 sec, default = 30).
* `ap_fgscan_report_intv` - Period of time between foreground scan reports (15 - 600 sec, default = 15).
* `ap_scan` - Enable/disable rogue AP detection. Valid values: `disable`, `enable`.

* `ap_scan_passive` - Enable/disable passive scanning. Enable means do not send probe request on any channels (default = disable). Valid values: `disable`, `enable`.

* `ap_scan_threshold` - Minimum signal level/threshold in dBm required for the AP to report detected rogue AP (-95 to -20, default = -90).
* `asleap_attack` - Enable/disable asleap attack detection (default = disable). Valid values: `disable`, `enable`.

* `assoc_flood_thresh` - The threshold value for association frame flooding.
* `assoc_flood_time` - Number of seconds after which a station is considered not connected.
* `assoc_frame_flood` - Enable/disable association frame flooding detection (default = disable). Valid values: `disable`, `enable`.

* `auth_flood_thresh` - The threshold value for authentication frame flooding.
* `auth_flood_time` - Number of seconds after which a station is considered not connected.
* `auth_frame_flood` - Enable/disable authentication frame flooding detection (default = disable). Valid values: `disable`, `enable`.

* `comment` - Comment.
* `deauth_broadcast` - Enable/disable broadcasting de-authentication detection (default = disable). Valid values: `disable`, `enable`.

* `deauth_unknown_src_thresh` - Threshold value per second to deauth unknown src for DoS attack (0: no limit).
* `eapol_fail_flood` - Enable/disable EAPOL-Failure flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_fail_intv` - The detection interval for EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_fail_thresh` - The threshold value for EAPOL-Failure flooding in specified interval.
* `eapol_logoff_flood` - Enable/disable EAPOL-Logoff flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_logoff_intv` - The detection interval for EAPOL-Logoff flooding (1 - 3600 sec).
* `eapol_logoff_thresh` - The threshold value for EAPOL-Logoff flooding in specified interval.
* `eapol_pre_fail_flood` - Enable/disable premature EAPOL-Failure flooding (to STA) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_pre_fail_intv` - The detection interval for premature EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_pre_fail_thresh` - The threshold value for premature EAPOL-Failure flooding in specified interval.
* `eapol_pre_succ_flood` - Enable/disable premature EAPOL-Success flooding (to STA) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_pre_succ_intv` - The detection interval for premature EAPOL-Success flooding (1 - 3600 sec).
* `eapol_pre_succ_thresh` - The threshold value for premature EAPOL-Success flooding in specified interval.
* `eapol_start_flood` - Enable/disable EAPOL-Start flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_start_intv` - The detection interval for EAPOL-Start flooding (1 - 3600 sec).
* `eapol_start_thresh` - The threshold value for EAPOL-Start flooding in specified interval.
* `eapol_succ_flood` - Enable/disable EAPOL-Success flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_succ_intv` - The detection interval for EAPOL-Success flooding (1 - 3600 sec).
* `eapol_succ_thresh` - The threshold value for EAPOL-Success flooding in specified interval.
* `invalid_mac_oui` - Enable/disable invalid MAC OUI detection. Valid values: `disable`, `enable`.

* `long_duration_attack` - Enable/disable long duration attack detection based on user configured threshold (default = disable). Valid values: `disable`, `enable`.

* `long_duration_thresh` - Threshold value for long duration attack detection (1000 - 32767 usec, default = 8200).
* `name` - WIDS profile name.
* `null_ssid_probe_resp` - Enable/disable null SSID probe response detection (default = disable). Valid values: `disable`, `enable`.

* `sensor_mode` - Scan nearby WiFi stations (default = disable). Valid values: `disable`, `foreign`, `both`.

* `spoofed_deauth` - Enable/disable spoofed de-authentication attack detection (default = disable). Valid values: `disable`, `enable`.

* `weak_wep_iv` - Enable/disable weak WEP IV (Initialization Vector) detection (default = disable). Valid values: `disable`, `enable`.

* `wireless_bridge` - Enable/disable wireless bridge detection (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController WidsProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_widsprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
