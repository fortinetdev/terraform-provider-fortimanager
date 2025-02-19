---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_managedswitch_routeoffloadrouter"
description: |-
  Configure route offload MCLAG IP address.
---

# fortimanager_object_switchcontroller_managedswitch_routeoffloadrouter
Configure route offload MCLAG IP address.

~> This resource is a sub resource for variable `route_offload_router` of resource `fortimanager_object_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `managed_switch` - Managed Switch.

* `router_ip` - Router IP address.
* `vlan_name` - VLAN name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vlan_name}}.

## Import

ObjectSwitchController ManagedSwitchRouteOffloadRouter can be imported using any of these accepted formats:
```
Set import_options = ["managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_managedswitch_routeoffloadrouter.labelname {{vlan_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
