---
subcategory: "ObjectWireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_wagprofile"
description: |-
  Configure wireless access gateway (WAG) profiles used for tunnels on AP.
---

# fortimanager_object_wirelesscontroller_wagprofile
Configure wireless access gateway (WAG) profiles used for tunnels on AP.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `dhcp_ip_addr` - IP address of the monitoring DHCP request packet sent through the tunnel
* `name` - Tunnel profile name.
* `ping_interval` - Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
* `ping_number` - Number of the tunnel mointoring echo packets (1 - 65535, default = 5).
* `return_packet_timeout` - Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
* `tunnel_type` - Tunnel type. Valid values: `gre`, `l2tpv3`.

* `wag_ip` - IP Address of the wireless access gateway.
* `wag_port` - UDP port of the wireless access gateway.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController WagProfile can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_wagprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.