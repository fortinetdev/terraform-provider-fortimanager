---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_priorityprotocol"
description: |-
  Configure NPU priority protocol.
---

# fortimanager_object_system_npu_priorityprotocol
Configure NPU priority protocol.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `bfd` - Enable/disable NPU BFD priority protocol. Valid values: `disable`, `enable`.

* `bgp` - Enable/disable NPU BGP priority protocol. Valid values: `disable`, `enable`.

* `slbc` - Enable/disable NPU SLBC priority protocol. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuPriorityProtocol can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_priorityprotocol.labelname ObjectSystemNpuPriorityProtocol
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
