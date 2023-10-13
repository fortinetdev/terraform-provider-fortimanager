---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_qos_ipdscpmap_map"
description: |-
  Maps between IP-DSCP value to COS queue.
---

# fortimanager_object_switchcontroller_qos_ipdscpmap_map
Maps between IP-DSCP value to COS queue.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ip_dscp_map` - Ip Dscp Map.

* `cos_queue` - COS queue number.
* `diffserv` - Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.

* `ip_precedence` - IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.

* `name` - Dscp mapping entry name.
* `value` - Raw values of DSCP (0 - 63).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController QosIpDscpMapMap can be imported using any of these accepted formats:
```
Set import_options = ["ip_dscp_map=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_qos_ipdscpmap_map.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
