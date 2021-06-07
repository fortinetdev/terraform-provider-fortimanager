---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_publicnetwork"
description: |-
  Enable/disable access to the public FortiGuard.
---

# fortimanager_fmupdate_publicnetwork
Enable/disable access to the public FortiGuard.

## Argument Reference


The following arguments are supported:


* `status` - Enable/disable public network (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate Publicnetwork can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_publicnetwork.labelname FmupdatePublicnetwork
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

