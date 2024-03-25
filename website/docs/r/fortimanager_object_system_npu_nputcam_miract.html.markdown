---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_nputcam_miract"
description: |-
  Mirror action of TCAM.
---

# fortimanager_object_system_npu_nputcam_miract
Mirror action of TCAM.

~> This resource is a sub resource for variable `mir_act` of resource `fortimanager_object_system_npu_nputcam`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `npu_tcam` - Npu Tcam.

* `vlif` - tcam mirror action vlif.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuNpuTcamMirAct can be imported using any of these accepted formats:
```
Set import_options = ["npu_tcam=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_nputcam_miract.labelname ObjectSystemNpuNpuTcamMirAct
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
