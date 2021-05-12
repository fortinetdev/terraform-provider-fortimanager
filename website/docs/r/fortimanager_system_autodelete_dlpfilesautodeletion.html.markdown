---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_autodelete_dlpfilesautodeletion"
description: |-
  Automatic deletion policy for DLP archives.
---

# fortimanager_system_autodelete_dlpfilesautodeletion
Automatic deletion policy for DLP archives.

## Argument Reference


The following arguments are supported:


* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoDeleteDlpFilesAutoDeletion can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_autodelete_dlpfilesautodeletion.labelname SystemAutoDeleteDlpFilesAutoDeletion
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

