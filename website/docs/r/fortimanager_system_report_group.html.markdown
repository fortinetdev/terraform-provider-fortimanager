---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_report_group"
description: |-
  Report group.
---

# fortimanager_system_report_group
Report group.

## Argument Reference


The following arguments are supported:


* `fmgadom` - Admin domain name.
* `case_insensitive` - Case insensitive. disable - Disable the case insensitive match. enable - Enable the case insensitive match. Valid values: `disable`, `enable`.

* `chart_alternative` - Chart-Alternative. The structure of `chart_alternative` block is documented below.
* `group_by` - Group-By. The structure of `group_by` block is documented below.
* `group_id` - Group ID.
* `report_like` - Report pattern.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `chart_alternative` block supports:

* `chart_name` - Chart name.
* `chart_replace` - Chart replacement.

The `group_by` block supports:

* `var_expression` - Variable expression.
* `var_name` - Variable name.
* `var_type` - Variable type. integer - Integer. string - String. enum - Enum. ip - IP. Valid values: `integer`, `string`, `enum`, `ip`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{group_id}}.

## Import

System ReportGroup can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_report_group.labelname {{group_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

