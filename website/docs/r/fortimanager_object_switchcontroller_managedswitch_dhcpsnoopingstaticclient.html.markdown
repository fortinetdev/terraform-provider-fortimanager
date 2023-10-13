---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_managedswitch_dhcpsnoopingstaticclient"
description: |-
  Configure FortiSwitch DHCP snooping static clients.
---

# fortimanager_object_switchcontroller_managedswitch_dhcpsnoopingstaticclient
Configure FortiSwitch DHCP snooping static clients.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `managed_switch` - Managed Switch.

* `ip` - Client static IP address.
* `mac` - Client MAC address.
* `name` - Client name.
* `port` - Interface name.
* `vlan` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController ManagedSwitchDhcpSnoopingStaticClient can be imported using any of these accepted formats:
```
Set import_options = ["managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_managedswitch_dhcpsnoopingstaticclient.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
