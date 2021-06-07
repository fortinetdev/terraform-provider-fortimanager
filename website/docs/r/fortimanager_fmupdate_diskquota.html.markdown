---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_diskquota"
description: |-
  Configure disk space available for use by the Upgrade Manager.
---

# fortimanager_fmupdate_diskquota
Configure disk space available for use by the Upgrade Manager.

## Argument Reference


The following arguments are supported:


* `value` - Configure the size of the Upgrade Manager disk quota, in megabytes.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate DiskQuota can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_diskquota.labelname FmupdateDiskQuota
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

