---
subcategory: "System SQL"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_sql_tsindexfield"
description: |-
  List of SQL text search index fields.
---

# fortimanager_system_sql_tsindexfield
List of SQL text search index fields.

~> This resource is a sub resource for variable `ts_index_field` of resource `fortimanager_system_sql`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_sql_tsindexfield" "trname" {
  category = "FGT-app-ctrl"
  value    = "user,group,srcip,dstip,dstport,service,app,action,hostname"
}
```

## Argument Reference


The following arguments are supported:


* `category` - Category of text search index fields.
* `value` - Fields of text search index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{category}}.

## Import

System SqlTsIndexField can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_sql_tsindexfield.labelname {{category}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

