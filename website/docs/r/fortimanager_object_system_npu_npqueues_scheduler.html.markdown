---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_npqueues_scheduler"
description: |-
  Configure a NP7 QoS Scheduler.
---

# fortimanager_object_system_npu_npqueues_scheduler
Configure a NP7 QoS Scheduler.

~> This resource is a sub resource for variable `scheduler` of resource `fortimanager_object_system_npu_npqueues`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_npqueues_scheduler" "trname" {
  name = "terr-scheduler"
  mode = "priority"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `mode` - Scheduler Mode. Valid values: `none`, `priority`, `round-robin`.

* `name` - Scheduler Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem NpuNpQueuesScheduler can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_npqueues_scheduler.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
