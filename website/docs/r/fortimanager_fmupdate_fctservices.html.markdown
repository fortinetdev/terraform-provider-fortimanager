---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fctservices"
description: |-
  Configure FortiGuard to provide services to FortiClient installations.
---

# fortimanager_fmupdate_fctservices
Configure FortiGuard to provide services to FortiClient installations.

## Argument Reference


The following arguments are supported:


* `port` - Configure the port number on which the built-in FortiGuard should provide updates to FortiClient installations (1 - 65535, default = 80).
* `status` - Enable/disable built-in FortiGuard service to FortiClient installations (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FctServices can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fctservices.labelname FmupdateFctServices
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

