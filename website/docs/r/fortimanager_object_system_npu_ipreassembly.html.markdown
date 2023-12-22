---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_ipreassembly"
description: |-
  IP reassebmly engine configuration.
---

# fortimanager_object_system_npu_ipreassembly
IP reassebmly engine configuration.

~> This resource is a sub resource for variable `ip_reassembly` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_ipreassembly" "trname" {
  max_timeout = 200000
  min_timeout = 12000
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `max_timeout` - Maximum timeout value for IP reassembly (5 us - 600,000,000 us).
* `min_timeout` - Minimum timeout value for IP reassembly (5 us - 600,000,000 us).
* `status` - Set IP reassembly processing status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuIpReassembly can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_ipreassembly.labelname ObjectSystemNpuIpReassembly
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
