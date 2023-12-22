---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_dswqueuedtsprofile"
description: |-
  Configure NPU DSW Queue DTS profile.
---

# fortimanager_object_system_npu_dswqueuedtsprofile
Configure NPU DSW Queue DTS profile.

~> This resource is a sub resource for variable `dsw_queue_dts_profile` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_dswqueuedtsprofile" "trname" {
  name  = "terr-dswqueuedtsprofile"
  oport = "EIF0"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `iport` - Set NPU DSW DTS in port. Valid values: `EIF0`, `EIF1`, `EIF2`, `EIF3`, `EIF4`, `EIF5`, `EIF6`, `EIF7`, `HTX0`, `HTX1`, `SSE0`, `SSE1`, `SSE2`, `SSE3`, `RLT`, `DFR`, `IPSECI`, `IPSECO`, `IPTI`, `IPTO`, `VEP0`, `VEP2`, `VEP4`, `VEP6`, `IVS`, `L2TI1`, `L2TO`, `L2TI0`, `PLE`, `SPATH`, `QTM`.

* `name` - Name.
* `oport` - Set NPU DSW DTS out port. Valid values: `EIF0`, `EIF1`, `EIF2`, `EIF3`, `EIF4`, `EIF5`, `EIF6`, `EIF7`, `HRX`, `SSE0`, `SSE1`, `SSE2`, `SSE3`, `RLT`, `DFR`, `IPSECI`, `IPSECO`, `IPTI`, `IPTO`, `VEP0`, `VEP2`, `VEP4`, `VEP6`, `IVS`, `L2TI1`, `L2TO`, `L2TI0`, `PLE`, `SYNK`, `NSS`, `TSK`, `QTM`.

* `profile_id` - Set NPU DSW DTS profile id.
* `queue_select` - Set NPU DSW DTS queue id select(0 - reset to default).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem NpuDswQueueDtsProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_dswqueuedtsprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
