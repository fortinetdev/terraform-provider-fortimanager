---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fdssetting_pushoverridetoclient_announceip"
description: |-
  Announce IP addresses for the device.
---

# fortimanager_fmupdate_fdssetting_pushoverridetoclient_announceip
Announce IP addresses for the device.

## Argument Reference


The following arguments are supported:


* `fosid` - ID of the announce IP address (1 - 10).
* `ip` - Announce IPv4 address.
* `port` - Announce IP port (1 - 65535, default = 8890).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSettingPushOverrideToClientAnnounceIp can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fdssetting_pushoverridetoclient_announceip.labelname FmupdateFdsSettingPushOverrideToClientAnnounceIp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

