---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_webspam_fgdsetting_serveroverride_servlist"
description: |-
  Override server.
---

# fortimanager_fmupdate_webspam_fgdsetting_serveroverride_servlist
Override server.

## Argument Reference


The following arguments are supported:


* `fosid` - Override server ID (1 - 10).
* `ip` - IPv4 address of the override server.
* `ip6` - IPv6 address of the override server.
* `port` - Port number to use when contacting FortiGuard (1 - 65535, default = 443).
* `service_type` - Override service type. fgd - Server override config for fgd fgc - Server override config for fgc fsa - Server override config for fsa Valid values: `fgd`, `fgc`, `fsa`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate WebSpamFgdSettingServerOverrideServlist can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_webspam_fgdsetting_serveroverride_servlist.labelname FmupdateWebSpamFgdSettingServerOverrideServlist
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

