---
subcategory: "System AutoDelete"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_autodelete"
description: |-
  Automatic deletion policy for logs, reports, archived, and quarantined files.
---

# fortimanager_system_autodelete
Automatic deletion policy for logs, reports, archived, and quarantined files.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`dlp_files_auto_deletion`: `fortimanager_system_autodelete_dlpfilesautodeletion`
`log_auto_deletion`: `fortimanager_system_autodelete_logautodeletion`
`quarantine_files_auto_deletion`: `fortimanager_system_autodelete_quarantinefilesautodeletion`
`report_auto_deletion`: `fortimanager_system_autodelete_reportautodeletion`



## Example Usage

```hcl
resource "fortimanager_system_autodelete" "trname" {
  status_fake = 1
}
```

## Argument Reference


The following arguments are supported:


* `dlp_files_auto_deletion` - Dlp-Files-Auto-Deletion. The structure of `dlp_files_auto_deletion` block is documented below.
* `log_auto_deletion` - Log-Auto-Deletion. The structure of `log_auto_deletion` block is documented below.
* `quarantine_files_auto_deletion` - Quarantine-Files-Auto-Deletion. The structure of `quarantine_files_auto_deletion` block is documented below.
* `report_auto_deletion` - Report-Auto-Deletion. The structure of `report_auto_deletion` block is documented below.
* `status_fake` - Fake value for the menu to work.

The `dlp_files_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.

The `log_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.

The `quarantine_files_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.

The `report_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoDelete can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_autodelete.labelname SystemAutoDelete
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

