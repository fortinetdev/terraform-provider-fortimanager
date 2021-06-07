---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_serveraccesspriorities"
description: |-
  Configure priorities for FortiGate units accessing antivirus updates and web filtering services.
---

# fortimanager_fmupdate_serveraccesspriorities
Configure priorities for FortiGate units accessing antivirus updates and web filtering services.

## Argument Reference


The following arguments are supported:


* `access_public` - Enable/disable FortiGates to Access Public FortiGuard Servers when Private Servers are Unavailable (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `av_ips` - Enable/disable Antivirus and IPS Update Service for Private Server(default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `private_server` - Private-Server. The structure of `private_server` block is documented below.
* `web_spam` - Enable/disable Web Filter and Email Filter Update Service for Private Server (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `private_server` block supports:

* `id` - Private server ID (1 - 10).
* `ip` - IPv4 address of the FortiManager unit or private server.
* `ip6` - IPv6 address of the FortiManager unit or private server.
* `time_zone` - Time zone of the private server (-24 = local time zone, default = -24).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate ServerAccessPriorities can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_serveraccesspriorities.labelname FmupdateServerAccessPriorities
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

