---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fdssetting_pushoverride"
description: |-
  Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.
---

# fortimanager_fmupdate_fdssetting_pushoverride
Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.

## Argument Reference


The following arguments are supported:


* `ip` - External or virtual IP address of the NAT device that will forward push messages to the FortiManager unit.
* `port` - Receiving port number on the NAT device (1 - 65535, default = 9443).
* `status` - Enable/disable push updates for clients (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSettingPushOverride can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fdssetting_pushoverride.labelname FmupdateFdsSettingPushOverride
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

