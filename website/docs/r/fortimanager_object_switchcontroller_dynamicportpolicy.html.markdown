---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_dynamicportpolicy"
description: |-
  Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.
---

# fortimanager_object_switchcontroller_dynamicportpolicy
Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description for the Dynamic port policy.
* `name` - Dynamic port policy name.
* `policy` - Policy. The structure of `policy` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `policy` block supports:

* `n802_1x` - 802.1x security policy to be applied when using this policy.
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.

* `category` - Category of Dynamic port policy. Valid values: `device`, `interface-tag`.

* `description` - Description for the policy.
* `family` - Match policy based on family.
* `host` - Match policy based on host.
* `hw_vendor` - Match policy based on hardware vendor.
* `interface_tags` - Match policy based on the FortiSwitch interface object tags.
* `lldp_profile` - LLDP profile to be applied when using this policy.
* `mac` - Match policy based on MAC address.
* `name` - Policy name.
* `qos_policy` - QoS policy to be applied when using this policy.
* `status` - Enable/disable policy. Valid values: `disable`, `enable`.

* `type` - Match policy based on type.
* `vlan_policy` - VLAN policy to be applied when using this policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController DynamicPortPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_dynamicportpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
