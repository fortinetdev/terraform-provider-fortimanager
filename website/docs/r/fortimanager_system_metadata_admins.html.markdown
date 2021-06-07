---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_metadata_admins"
description: |-
  Configure admins.
---

# fortimanager_system_metadata_admins
Configure admins.

## Argument Reference


The following arguments are supported:


* `fieldlength` - Field length. 20 - Field length of 20. 50 - Field length of 50. 255 - Field length of 255. Valid values: `20`, `50`, `255`.

* `fieldname` - Field name.
* `importance` - Field importance. optional - This field is optional. required - This field is required. Valid values: `optional`, `required`.

* `status` - Field status. disabled - This field is disabled. enabled - This field is enabled. Valid values: `disabled`, `enabled`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fieldname}}.

## Import

System MetadataAdmins can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_metadata_admins.labelname {{fieldname}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

