---
subcategory: "System NTP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ntp"
description: |-
  NTP settings.
---

# fortimanager_system_ntp
NTP settings.

## Argument Reference


The following arguments are supported:


* `ntpserver` - Ntpserver. The structure of `ntpserver` block is documented below.
* `status` - Enable/disable NTP. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `sync_interval` - NTP sync interval (min).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ntpserver` block supports:

* `authentication` - Enable/disable MD5 authentication. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `id` - Time server ID.
* `key` - Key for authentication.
* `key_id` - Key ID for authentication.
* `ntpv3` - Enable/disable NTPv3. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `server` - IP address/hostname of NTP Server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ntp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_ntp.labelname SystemNtp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

