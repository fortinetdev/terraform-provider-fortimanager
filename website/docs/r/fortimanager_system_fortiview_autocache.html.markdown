---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_fortiview_autocache"
description: |-
  FortiView auto-cache settings.
---

# fortimanager_system_fortiview_autocache
FortiView auto-cache settings.

## Argument Reference


The following arguments are supported:


* `aggressive_fortiview` - Enable/disable auto-cache on fortiview aggressively. disable - Disable the aggressive fortiview auto-cache. enable - Enable the aggressive fortiview auto-cache. Valid values: `disable`, `enable`.

* `interval` - The time interval in hours for fortiview auto-cache.
* `status` - Enable/disable fortiview auto-cache. disable - Disable the fortiview auto-cache. enable - Enable the fortiview auto-cache. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FortiviewAutoCache can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_fortiview_autocache.labelname SystemFortiviewAutoCache
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

