---
subcategory: "Object Switch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_customcommand"
description: |-
  Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.
---

# fortimanager_object_switchcontroller_customcommand
Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `command` - String of commands to send to FortiSwitch devices (For example (0x0.00000172e012p-1022 = return key): config switch trunk 0x0.07f258d1c4808p-1022 edit myTrunk 0x0.00000172e012p-1022 set members port1 port2 0x0.0000000000001p-1022 end 0x0.00000172df7ep-1022).
* `command_name` - Command name called by the FortiGate switch controller in the execute command.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{command_name}}.

## Import

ObjectSwitchController CustomCommand can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_customcommand.labelname {{command_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
