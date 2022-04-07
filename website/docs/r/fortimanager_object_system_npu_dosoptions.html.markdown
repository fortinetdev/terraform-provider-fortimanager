---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_dosoptions"
description: |-
  NPU DoS configurations.
---

# fortimanager_object_system_npu_dosoptions
NPU DoS configurations.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `npu_dos_meter_mode` - Set DoS meter npu offloading mode. Valid values: `local`, `global`.

* `npu_dos_synproxy_mode` - Set NPU DoS SYNPROXY mode. Valid values: `synack2ack`, `pass-synack`.

* `npu_dos_tpe_mode` - Enable/Disable inserting DoS meter id to session table. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuDosOptions can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_dosoptions.labelname ObjectSystemNpuDosOptions
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
