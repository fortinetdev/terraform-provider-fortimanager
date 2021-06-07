---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_analyzer_virusreport"
description: |-
  Send virus detection notification to FortiGuard.
---

# fortimanager_fmupdate_analyzer_virusreport
Send virus detection notification to FortiGuard.

## Argument Reference


The following arguments are supported:


* `status` - Enable/disable sending virus detection notification to FortiGuard (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate AnalyzerVirusreport can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_analyzer_virusreport.labelname FmupdateAnalyzerVirusreport
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

