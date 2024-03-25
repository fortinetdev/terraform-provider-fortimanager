---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_npqueues"
description: |-
  Configure queue assignment on NP7.
---

# fortimanager_object_system_npu_npqueues
Configure queue assignment on NP7.

~> This resource is a sub resource for variable `np_queues` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ethernet_type`: `fortimanager_object_system_npu_npqueues_ethernettype`
>- `ip_protocol`: `fortimanager_object_system_npu_npqueues_ipprotocol`
>- `ip_service`: `fortimanager_object_system_npu_npqueues_ipservice`
>- `profile`: `fortimanager_object_system_npu_npqueues_profile`
>- `scheduler`: `fortimanager_object_system_npu_npqueues_scheduler`



## Example Usage

```hcl
resource "fortimanager_object_system_npu_npqueues" "trname" {
  ethernet_type {
    name  = "type"
    queue = 10
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ethernet_type` - Ethernet-Type. The structure of `ethernet_type` block is documented below.
* `ip_protocol` - Ip-Protocol. The structure of `ip_protocol` block is documented below.
* `ip_service` - Ip-Service. The structure of `ip_service` block is documented below.
* `profile` - Profile. The structure of `profile` block is documented below.
* `scheduler` - Scheduler. The structure of `scheduler` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ethernet_type` block supports:

* `name` - Ethernet Type Name.
* `queue` - Queue Number.
* `type` - Ethernet Type.
* `weight` - Class Weight.

The `ip_protocol` block supports:

* `name` - IP Protocol Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `weight` - Class Weight.

The `ip_service` block supports:

* `dport` - Destination Port.
* `name` - IP Service Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `sport` - Source Port.
* `weight` - Class Weight.

The `profile` block supports:

* `cos0` - Queue number of CoS 0. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos1` - Queue number of CoS 1. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos2` - Queue number of CoS 2. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos3` - Queue number of CoS 3. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos4` - Queue number of CoS 4. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos5` - Queue number of CoS 5. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos6` - Queue number of CoS 6. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos7` - Queue number of CoS 7. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp0` - Queue number of DSCP 0. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp1` - Queue number of DSCP 1. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp10` - Queue number of DSCP 10. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp11` - Queue number of DSCP 11. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp12` - Queue number of DSCP 12. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp13` - Queue number of DSCP 13. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp14` - Queue number of DSCP 14. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp15` - Queue number of DSCP 15. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp16` - Queue number of DSCP 16. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp17` - Queue number of DSCP 17. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp18` - Queue number of DSCP 18. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp19` - Queue number of DSCP 19. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp2` - Queue number of DSCP 2. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp20` - Queue number of DSCP 20. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp21` - Queue number of DSCP 21. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp22` - Queue number of DSCP 22. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp23` - Queue number of DSCP 23. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp24` - Queue number of DSCP 24. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp25` - Queue number of DSCP 25. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp26` - Queue number of DSCP 26. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp27` - Queue number of DSCP 27. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp28` - Queue number of DSCP 28. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp29` - Queue number of DSCP 29. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp3` - Queue number of DSCP 3. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp30` - Queue number of DSCP 30. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp31` - Queue number of DSCP 31. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp32` - Queue number of DSCP 32. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp33` - Queue number of DSCP 33. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp34` - Queue number of DSCP 34. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp35` - Queue number of DSCP 35. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp36` - Queue number of DSCP 36. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp37` - Queue number of DSCP 37. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp38` - Queue number of DSCP 38. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp39` - Queue number of DSCP 39. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp4` - Queue number of DSCP 4. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp40` - Queue number of DSCP 40. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp41` - Queue number of DSCP 41. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp42` - Queue number of DSCP 42. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp43` - Queue number of DSCP 43. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp44` - Queue number of DSCP 44. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp45` - Queue number of DSCP 45. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp46` - Queue number of DSCP 46. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp47` - Queue number of DSCP 47. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp48` - Queue number of DSCP 48. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp49` - Queue number of DSCP 49. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp5` - Queue number of DSCP 5. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp50` - Queue number of DSCP 50. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp51` - Queue number of DSCP 51. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp52` - Queue number of DSCP 52. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp53` - Queue number of DSCP 53. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp54` - Queue number of DSCP 54. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp55` - Queue number of DSCP 55. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp56` - Queue number of DSCP 56. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp57` - Queue number of DSCP 57. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp58` - Queue number of DSCP 58. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp59` - Queue number of DSCP 59. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp6` - Queue number of DSCP 6. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp60` - Queue number of DSCP 60. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp61` - Queue number of DSCP 61. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp62` - Queue number of DSCP 62. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp63` - Queue number of DSCP 63. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp7` - Queue number of DSCP 7. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp8` - Queue number of DSCP 8. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp9` - Queue number of DSCP 9. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `id` - Profile ID.
* `type` - Profile type. Valid values: `cos`, `dscp`.

* `weight` - Class weight.

The `scheduler` block supports:

* `mode` - Scheduler Mode. Valid values: `none`, `priority`, `round-robin`.

* `name` - Scheduler Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuNpQueues can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_npqueues.labelname ObjectSystemNpuNpQueues
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
