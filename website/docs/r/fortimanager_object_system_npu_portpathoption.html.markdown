---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_portpathoption"
description: |-
  Configure port using NPU or Intel-NIC.
---

# fortimanager_object_system_npu_portpathoption
Configure port using NPU or Intel-NIC.

~> This resource is a sub resource for variable `port_path_option` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_portpathoption" "trname" {
  ports_using_npu = [20, 23]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ports_using_npu` - Set ha/aux ports to handle traffic with NPU (otherise traffic goes to Intel-NIC and then CPU).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuPortPathOption can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_portpathoption.labelname ObjectSystemNpuPortPathOption
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
