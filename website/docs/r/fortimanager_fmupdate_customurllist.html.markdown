---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_customurllist"
description: |-
  Configure the URL database for rating and filtering.
---

# fortimanager_fmupdate_customurllist
Configure the URL database for rating and filtering.

## Argument Reference


The following arguments are supported:


* `db_selection` - Manage the URL database (default = both). both - Support both custom-url and FortiGuard database. custom-url - Custom imported URL list. fortiguard-db - Fortinet's Fortiguard database. Valid values: `both`, `custom-url`, `fortiguard-db`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate CustomUrlList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_customurllist.labelname FmupdateCustomUrlList
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

