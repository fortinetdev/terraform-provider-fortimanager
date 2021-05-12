---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_serveroverridestatus"
description: |-
  Configure strict/loose server override.
---

# fortimanager_fmupdate_serveroverridestatus
Configure strict/loose server override.

## Argument Reference


The following arguments are supported:


* `mode` - Server override mode (default = loose). strict - Access override server only. loose - Allow access other servers. Valid values: `strict`, `loose`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate ServerOverrideStatus can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_serveroverridestatus.labelname FmupdateServerOverrideStatus
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

