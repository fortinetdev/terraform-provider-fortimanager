---
subcategory: "ObjectSwitch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_qos_dot1pmap"
description: |-
  Configure FortiSwitch QoS 802.1p.
---

# fortimanager_object_switchcontroller_qos_dot1pmap
Configure FortiSwitch QoS 802.1p.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description of the 802.1p name.
* `egress_pri_tagging` - Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.

* `name` - Dot1p map name.
* `priority_0` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_1` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_2` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_3` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_4` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_5` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_6` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_7` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController QosDot1PMap can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_qos_dot1pmap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
