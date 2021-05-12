---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fdssetting_pushoverridetoclient"
description: |-
  Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.
---

# fortimanager_fmupdate_fdssetting_pushoverridetoclient
Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.

## Argument Reference


The following arguments are supported:


* `announce_ip` - Announce-Ip. The structure of `announce_ip` block is documented below.
* `status` - Enable/disable push updates (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `announce_ip` block supports:

* `id` - ID of the announce IP address (1 - 10).
* `ip` - Announce IPv4 address.
* `port` - Announce IP port (1 - 65535, default = 8890).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSettingPushOverrideToClient can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fdssetting_pushoverridetoclient.labelname FmupdateFdsSettingPushOverrideToClient
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

