---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_mmsprofile"
description: |-
  Configure MMS profiles.
---

# fortimanager_object_firewall_mmsprofile
Configure MMS profiles.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `avnotificationtable` - AntiVirus notification table ID.
* `bwordtable` - MMS banned word table ID.
* `carrier_endpoint_prefix` - Enable/disable prefixing of end point values. Valid values: `disable`, `enable`.

* `carrier_endpoint_prefix_range_max` - Maximum length of end point value that can be prefixed (1 - 48).
* `carrier_endpoint_prefix_range_min` - Minimum end point length to be prefixed (1 - 48).
* `carrier_endpoint_prefix_string` - String with which to prefix End point values.
* `carrierendpointbwltable` - Carrier end point filter table ID.
* `comment` - Comment.
* `dupe` - Dupe. The structure of `dupe` block is documented below.
* `extended_utm_log` - Enable/disable detailed UTM log messages. Valid values: `disable`, `enable`.

* `flood` - Flood. The structure of `flood` block is documented below.
* `mm1` - MM1 options. Valid values: `avmonitor`, `block`, `oversize`, `quarantine`, `scan`, `avquery`, `bannedword`, `no-content-summary`, `archive-summary`, `archive-full`, `carrier-endpoint-bwl`, `remove-blocked`, `chunkedbypass`, `clientcomfort`, `servercomfort`, `strict-file`, `mms-checksum`.

* `mm1_addr_hdr` - HTTP header field (for MM1) containing user address.
* `mm1_addr_source` - Source for MM1 user address. Valid values: `http-header`, `cookie`.

* `mm1_convert_hex` - Enable/disable converting user address from HEX string for MM1. Valid values: `disable`, `enable`.

* `mm1_outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.

* `mm1_retr_dupe` - Enable/disable duplicate scanning of MM1 retr. Valid values: `disable`, `enable`.

* `mm1_retrieve_scan` - Enable/disable scanning on MM1 retrieve configuration messages. Valid values: `disable`, `enable`.

* `mm1comfortamount` - MM1 comfort amount (0 - 4294967295).
* `mm1comfortinterval` - MM1 comfort interval (0 - 4294967295).
* `mm1oversizelimit` - Maximum file size to scan (1 - 819200 kB).
* `mm3` - MM3 options. Valid values: `avmonitor`, `block`, `oversize`, `quarantine`, `scan`, `avquery`, `bannedword`, `no-content-summary`, `archive-summary`, `archive-full`, `carrier-endpoint-bwl`, `remove-blocked`, `fragmail`, `splice`, `mms-checksum`.

* `mm3_outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.

* `mm3oversizelimit` - Maximum file size to scan (1 - 819200 kB).
* `mm4` - MM4 options. Valid values: `avmonitor`, `block`, `oversize`, `quarantine`, `scan`, `avquery`, `bannedword`, `no-content-summary`, `archive-summary`, `archive-full`, `carrier-endpoint-bwl`, `remove-blocked`, `fragmail`, `splice`, `mms-checksum`.

* `mm4_outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.

* `mm4oversizelimit` - Maximum file size to scan (1 - 819200 kB).
* `mm7` - MM7 options. Valid values: `avmonitor`, `block`, `oversize`, `quarantine`, `scan`, `avquery`, `bannedword`, `no-content-summary`, `archive-summary`, `archive-full`, `carrier-endpoint-bwl`, `remove-blocked`, `chunkedbypass`, `clientcomfort`, `servercomfort`, `strict-file`, `mms-checksum`.

* `mm7_addr_hdr` - HTTP header field (for MM7) containing user address.
* `mm7_addr_source` - Source for MM7 user address. Valid values: `http-header`, `cookie`.

* `mm7_convert_hex` - Enable/disable conversion of user address from HEX string for MM7. Valid values: `disable`, `enable`.

* `mm7_outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.

* `mm7comfortamount` - MM7 comfort amount (0 - 4294967295).
* `mm7comfortinterval` - MM7 comfort interval (0 - 4294967295).
* `mm7oversizelimit` - Maximum file size to scan (1 - 819200 kB).
* `mms_antispam_mass_log` - Enable/disable logging for MMS antispam mass. Valid values: `disable`, `enable`.

* `mms_av_block_log` - Enable/disable logging for MMS antivirus file blocking. Valid values: `disable`, `enable`.

* `mms_av_oversize_log` - Enable/disable logging for MMS antivirus oversize file blocking. Valid values: `disable`, `enable`.

* `mms_av_virus_log` - Enable/disable logging for MMS antivirus scanning. Valid values: `disable`, `enable`.

* `mms_carrier_endpoint_filter_log` - Enable/disable logging for MMS end point filter blocking. Valid values: `disable`, `enable`.

* `mms_checksum_log` - Enable/disable MMS content checksum logging. Valid values: `disable`, `enable`.

* `mms_checksum_table` - MMS content checksum table ID.
* `mms_notification_log` - Enable/disable logging for MMS notification messages. Valid values: `disable`, `enable`.

* `mms_web_content_log` - Enable/disable logging for MMS web content blocking. Valid values: `disable`, `enable`.

* `mmsbwordthreshold` - MMS banned word threshold.
* `name` - Profile name.
* `notif_msisdn` - Notif-Msisdn. The structure of `notif_msisdn` block is documented below.
* `notification` - Notification. The structure of `notification` block is documented below.
* `outbreak_prevention` - Outbreak-Prevention. The structure of `outbreak_prevention` block is documented below.
* `remove_blocked_const_length` - Enable/disable MMS replacement of blocked file constant length. Valid values: `disable`, `enable`.

* `replacemsg_group` - Replacement message group.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dupe` block supports:

* `action1` - Action to take when threshold reached. Valid values: `log`, `archive`, `intercept`, `block`, `archive-first`, `alert-notif`.

* `action2` - Action to take when threshold reached. Valid values: `log`, `archive`, `intercept`, `block`, `archive-first`, `alert-notif`.

* `action3` - Action to take when threshold reached. Valid values: `log`, `archive`, `intercept`, `block`, `archive-first`, `alert-notif`.

* `block_time1` - Duration for which action takes effect (0 - 35791 min).
* `block_time2` - Duration for which action takes effect (0 - 35791 min).
* `block_time3` - Duration action takes effect (0 - 35791 min).
* `limit1` - Maximum number of messages allowed.
* `limit2` - Maximum number of messages allowed.
* `limit3` - Maximum number of messages allowed.
* `protocol` - Protocol.
* `status1` - Enable/disable status1 detection. Valid values: `disable`, `enable`.

* `status2` - Enable/disable status2 detection. Valid values: `disable`, `enable`.

* `status3` - Enable/disable status3 detection. Valid values: `disable`, `enable`.

* `window1` - Window to count messages over (1 - 2880 min).
* `window2` - Window to count messages over (1 - 2880 min).
* `window3` - Window to count messages over (1 - 2880 min).

The `flood` block supports:

* `action1` - Action to take when threshold reached. Valid values: `log`, `archive`, `intercept`, `block`, `archive-first`, `alert-notif`.

* `action2` - Action to take when threshold reached. Valid values: `log`, `archive`, `intercept`, `block`, `archive-first`, `alert-notif`.

* `action3` - Action to take when threshold reached. Valid values: `log`, `archive`, `intercept`, `block`, `archive-first`, `alert-notif`.

* `block_time1` - Duration for which action takes effect (0 - 35791 min).
* `block_time2` - Duration for which action takes effect (0 - 35791 min).
* `block_time3` - Duration action takes effect (0 - 35791 min).
* `limit1` - Maximum number of messages allowed.
* `limit2` - Maximum number of messages allowed.
* `limit3` - Maximum number of messages allowed.
* `protocol` - Protocol.
* `status1` - Enable/disable status1 detection. Valid values: `disable`, `enable`.

* `status2` - Enable/disable status2 detection. Valid values: `disable`, `enable`.

* `status3` - Enable/disable status3 detection. Valid values: `disable`, `enable`.

* `window1` - Window to count messages over (1 - 2880 min).
* `window2` - Window to count messages over (1 - 2880 min).
* `window3` - Window to count messages over (1 - 2880 min).

The `notif_msisdn` block supports:

* `msisdn` - Recipient MSISDN.
* `threshold` - Thresholds on which this MSISDN will receive an alert. Valid values: `flood-thresh-1`, `flood-thresh-2`, `flood-thresh-3`, `dupe-thresh-1`, `dupe-thresh-2`, `dupe-thresh-3`.


The `notification` block supports:

* `alert_int` - Alert notification send interval.
* `alert_int_mode` - Alert notification interval mode. Valid values: `hours`, `minutes`.

* `alert_src_msisdn` - Specify from address for alert messages.
* `alert_status` - Alert notification status. Valid values: `disable`, `enable`.

* `bword_int` - Banned word notification send interval.
* `bword_int_mode` - Banned word notification interval mode. Valid values: `hours`, `minutes`.

* `bword_status` - Banned word notification status. Valid values: `disable`, `enable`.

* `carrier_endpoint_bwl_int` - Carrier end point black/white list notification send interval.
* `carrier_endpoint_bwl_int_mode` - Carrier end point black/white list notification interval mode. Valid values: `hours`, `minutes`.

* `carrier_endpoint_bwl_status` - Carrier end point black/white list notification status. Valid values: `disable`, `enable`.

* `days_allowed` - Weekdays on which notification messages may be sent. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `detect_server` - Enable/disable automatic server address determination. Valid values: `disable`, `enable`.

* `dupe_int` - Duplicate notification send interval.
* `dupe_int_mode` - Duplicate notification interval mode. Valid values: `hours`, `minutes`.

* `dupe_status` - Duplicate notification status. Valid values: `disable`, `enable`.

* `file_block_int` - File block notification send interval.
* `file_block_int_mode` - File block notification interval mode. Valid values: `hours`, `minutes`.

* `file_block_status` - File block notification status. Valid values: `disable`, `enable`.

* `flood_int` - Flood notification send interval.
* `flood_int_mode` - Flood notification interval mode. Valid values: `hours`, `minutes`.

* `flood_status` - Flood notification status. Valid values: `disable`, `enable`.

* `from_in_header` - Enable/disable insertion of from address in HTTP header. Valid values: `disable`, `enable`.

* `mms_checksum_int` - MMS checksum notification send interval.
* `mms_checksum_int_mode` - MMS checksum notification interval mode. Valid values: `hours`, `minutes`.

* `mms_checksum_status` - MMS checksum notification status. Valid values: `disable`, `enable`.

* `mmsc_hostname` - Host name or IP address of the MMSC.
* `mmsc_password` - Password required for authentication with the MMSC.
* `mmsc_port` - Port used on the MMSC for sending MMS messages (1 - 65535).
* `mmsc_url` - URL used on the MMSC for sending MMS messages.
* `mmsc_username` - User name required for authentication with the MMSC.
* `msg_protocol` - Protocol to use for sending notification messages. Valid values: `mm1`, `mm3`, `mm4`, `mm7`.

* `msg_type` - MM7 message type. Valid values: `submit-req`, `deliver-req`.

* `protocol` - Protocol.
* `rate_limit` - Rate limit for sending notification messages (0 - 250).
* `tod_window_duration` - Time of day window duration.
* `tod_window_end` - Obsolete.
* `tod_window_start` - Time of day window start.
* `user_domain` - Domain name to which the user addresses belong.
* `vas_id` - VAS identifier.
* `vasp_id` - VASP identifier.
* `virus_int` - Virus notification send interval.
* `virus_int_mode` - Virus notification interval mode. Valid values: `hours`, `minutes`.

* `virus_status` - Virus notification status. Valid values: `disable`, `enable`.


The `outbreak_prevention` block supports:

* `external_blocklist` - Enable/disable external malware blocklist. Valid values: `disable`, `enable`.

* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall MmsProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_mmsprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
