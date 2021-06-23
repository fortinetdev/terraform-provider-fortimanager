---
subcategory: "No Category"
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
* `extended_utm_log` - Enable/disable detailed UTM log messages. Valid values: `disable`, `enable`.

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
* `remove_blocked_const_length` - Enable/disable MMS replacement of blocked file constant length. Valid values: `disable`, `enable`.

* `replacemsg_group` - Replacement message group.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `notif_msisdn` block supports:

* `msisdn` - Recipient MSISDN.
* `threshold` - Thresholds on which this MSISDN will receive an alert. Valid values: `flood-thresh-1`, `flood-thresh-2`, `flood-thresh-3`, `dupe-thresh-1`, `dupe-thresh-2`, `dupe-thresh-3`.



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
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
