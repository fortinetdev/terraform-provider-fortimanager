---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_qos_queuepolicy_cosqueue"
description: |-
  COS queue configuration.
---

# fortimanager_object_switchcontroller_qos_queuepolicy_cosqueue
COS queue configuration.

~> This resource is a sub resource for variable `cos_queue` of resource `fortimanager_object_switchcontroller_qos_queuepolicy`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `queue_policy` - Queue Policy.

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

ObjectSwitchController QosQueuePolicyCosQueue can be imported using any of these accepted formats:
```
Set import_options = ["queue_policy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_qos_queuepolicy_cosqueue.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
