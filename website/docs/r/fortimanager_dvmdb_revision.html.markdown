---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvmdb_revision"
description: |-
  ADOM revision table.
---

# fortimanager_dvmdb_revision
ADOM revision table.

## Example Usage

```hcl
resource "fortimanager_dvmdb_revision" "trname" {
  created_by   = "admin"
  created_time = 1619027731
  desc         = "This is a Terraform example"
  name         = "terr-revision"
  version      = 1
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `created_by` - Created_By.
* `created_time` - Created_Time.
* `desc` - Desc.
* `locked` - Locked.
* `name` - Name.
* `version` - Version.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{version}}.

## Import

Dvmdb Revision can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvmdb_revision.labelname {{version}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
