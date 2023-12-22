---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_portcpumap"
description: |-
  Configure NPU interface to CPU core mapping.
---

# fortimanager_object_system_npu_portcpumap
Configure NPU interface to CPU core mapping.

~> This resource is a sub resource for variable `port_cpu_map` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_portcpumap" "trname" {
  interface = "port1"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cpu_core` - The CPU core to map to an interface.
* `interface` - The interface to map to a CPU core.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{interface}}.

## Import

ObjectSystem NpuPortCpuMap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_portcpumap.labelname {{interface}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
