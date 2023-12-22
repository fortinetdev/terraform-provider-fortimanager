---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_isfnpqueues"
description: |-
  Configure queues of switch port connected to NP6 XAUI on ingress path.
---

# fortimanager_object_system_npu_isfnpqueues
Configure queues of switch port connected to NP6 XAUI on ingress path.

~> This resource is a sub resource for variable `isf_np_queues` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_isfnpqueues" "trname" {
  cos0 = "cos-profile1"
  cos1 = "cos-profile2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cos0` - CoS profile name for CoS 0.
* `cos1` - CoS profile name for CoS 1.
* `cos2` - CoS profile name for CoS 2.
* `cos3` - CoS profile name for CoS 3.
* `cos4` - CoS profile name for CoS 4.
* `cos5` - CoS profile name for CoS 5.
* `cos6` - CoS profile name for CoS 6.
* `cos7` - CoS profile name for CoS 7.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuIsfNpQueues can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_isfnpqueues.labelname ObjectSystemNpuIsfNpQueues
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
