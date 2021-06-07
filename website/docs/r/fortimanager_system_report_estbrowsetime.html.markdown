---
subcategory: "System Report"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_report_estbrowsetime"
description: |-
  Report estimated browse time settings
---

# fortimanager_system_report_estbrowsetime
Report estimated browse time settings

## Argument Reference


The following arguments are supported:


* `max_read_time` - Read time threshold for each page view.
* `status` - Estimate browse time status. disable - Disable estimating browse time. enable - Enable estimating browse time. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ReportEstBrowseTime can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_report_estbrowsetime.labelname SystemReportEstBrowseTime
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

