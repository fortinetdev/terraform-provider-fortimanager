---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_multilayer"
description: |-
  Configure multilayer mode.
---

# fortimanager_fmupdate_multilayer
Configure multilayer mode.

## Argument Reference


The following arguments are supported:


* `webspam_rating` - Enable/disable URL/Antispam rating service (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate Multilayer can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_multilayer.labelname FmupdateMultilayer
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

