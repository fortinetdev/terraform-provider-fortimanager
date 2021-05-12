---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_fortiview_setting"
description: |-
  FortiView settings.
---

# fortimanager_system_fortiview_setting
FortiView settings.

## Argument Reference


The following arguments are supported:


* `not_scanned_apps` - Include/Exclude 'Not.Scanned' applications in FortiView. Set as 'exclude' if you want to filter out never scanned applications. exclude - Exclude 'Not.Scanned' applications in FortiView. include - Include 'Not.Scanned' applications in FortiView. Valid values: `exclude`, `include`.

* `resolve_ip` - Enable or disable resolving IP address to hostname in FortiView.  disable - Disable resolving IP address to hostname. enable - Enable resolving IP address to hostname. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FortiviewSetting can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_fortiview_setting.labelname SystemFortiviewSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

