---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_serveraccesspriorities_privateserver"
description: |-
  Configure multiple FortiManager units and private servers.
---

# fortimanager_fmupdate_serveraccesspriorities_privateserver
Configure multiple FortiManager units and private servers.

## Argument Reference


The following arguments are supported:


* `fosid` - Private server ID (1 - 10).
* `ip` - IPv4 address of the FortiManager unit or private server.
* `ip6` - IPv6 address of the FortiManager unit or private server.
* `time_zone` - Time zone of the private server (-24 = local time zone, default = -24).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Fmupdate ServerAccessPrioritiesPrivateServer can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_serveraccesspriorities_privateserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

