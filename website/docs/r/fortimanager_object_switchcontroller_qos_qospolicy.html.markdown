---
subcategory: "Object Switch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_qos_qospolicy"
description: |-
  Configure FortiSwitch QoS policy.
---

# fortimanager_object_switchcontroller_qos_qospolicy
Configure FortiSwitch QoS policy.

## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_qos_qospolicy" "trname" {
  default_cos  = 2
  name         = "terr-switch-controller-qos-policy"
  queue_policy = "default"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `default_cos` - Default cos queue for untagged packets.
* `name` - QoS policy name.
* `queue_policy` - QoS egress queue policy.
* `trust_dot1p_map` - QoS trust 802.1p map.
* `trust_ip_dscp_map` - QoS trust ip dscp map.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController QosQosPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_qos_qospolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
