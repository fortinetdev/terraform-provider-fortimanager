---
subcategory: "Object Switch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_qos_queuepolicy"
description: |-
  Configure FortiSwitch QoS egress queue policy.
---

# fortimanager_object_switchcontroller_qos_queuepolicy
Configure FortiSwitch QoS egress queue policy.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cos_queue`: `fortimanager_object_switchcontroller_qos_queuepolicy_cosqueue`



## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_qos_queuepolicy" "trname" {
  name     = "terr-switch-controller-qos-queue-policy"
  rate_by  = "kbps"
  schedule = "round-robin"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cos_queue` - Cos-Queue. The structure of `cos_queue` block is documented below.
* `name` - QoS policy name
* `rate_by` - COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.

* `schedule` - COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `cos_queue` block supports:

* `description` - Description of the COS queue.
* `drop_policy` - COS queue drop policy. Valid values: `taildrop`, `weighted-random-early-detection`.

* `ecn` - Enable/disable ECN packet marking to drop eligible packets. Valid values: `disable`, `enable`.

* `max_rate` - Maximum rate (0 - 4294967295 kbps, 0 to disable).
* `max_rate_percent` - Maximum rate (16100036777f link speed).
* `min_rate` - Minimum rate (0 - 4294967295 kbps, 0 to disable).
* `min_rate_percent` - Minimum rate (16100036777f link speed).
* `name` - Cos queue ID.
* `weight` - Weight of weighted round robin scheduling.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController QosQueuePolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_qos_queuepolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
