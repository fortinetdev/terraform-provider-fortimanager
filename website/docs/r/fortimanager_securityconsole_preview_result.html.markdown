---
subcategory: "Securityconsole"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_securityconsole_preview_result"
description: |-
  Retrieve the result of previous install/preview command.
---

# fortimanager_securityconsole_preview_result
Retrieve the result of previous install/preview command.

## Argument Reference


The following arguments are supported:


* `fmgadom` - Source ADOM name.
* `device` - Device.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Securityconsole PreviewResult can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_securityconsole_preview_result.labelname SecurityconsolePreviewResult
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
