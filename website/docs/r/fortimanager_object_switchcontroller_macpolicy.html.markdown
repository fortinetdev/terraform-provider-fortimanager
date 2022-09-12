---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_macpolicy"
description: |-
  Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.
---

# fortimanager_object_switchcontroller_macpolicy
Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this mac-policy is applied. Valid values: `disable`, `enable`.

* `fmgcount` - Enable/disable packet count on the NAC device. Valid values: `disable`, `enable`.

* `description` - Description for the MAC policy.
* `drop` - Drop. Valid values: `disable`, `enable`.

* `name` - MAC policy name.
* `traffic_policy` - Traffic policy to be applied when using this MAC policy.
* `vlan` - Ingress traffic VLAN assignment for the MAC address matching this MAC policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController MacPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_macpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
