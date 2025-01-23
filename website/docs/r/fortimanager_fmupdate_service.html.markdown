---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_service"
description: |-
  Enable/disable services provided by the built-in FortiGuard.
---

# fortimanager_fmupdate_service
Enable/disable services provided by the built-in FortiGuard.

## Example Usage

```hcl
resource "fortimanager_fmupdate_service" "trname" {
  avips           = "enable"
  query_antispam  = "enable"
  query_antivirus = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `avips` - Enable/disable the built-in FortiGuard to provide FortiGuard antivirus and IPS updates (default = enable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_antispam` - Enable/disable antispam service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_antivirus` - Enable/disable antivirus query service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_filequery` - Enable/disable file query service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_ioc` - Enable/disable the built-in FortiGuard to provide IoC query (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_iot` - Enable/disable file query service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_iot_collection` - Enable/disable IoT Collection Query service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_iot_vulnerability` - Enable/disable IoT Vulnerability Query service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_outbreak_prevention` - Enable/disable  outbreak prevention query service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_webfilter` - Enable/disable Web Filter service (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `webfilter_https_traversal` - Enable/disable Web Filter HTTPS traversal (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate Service can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_service.labelname FmupdateService
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

