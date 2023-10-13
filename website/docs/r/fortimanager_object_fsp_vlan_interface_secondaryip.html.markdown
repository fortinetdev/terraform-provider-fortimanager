---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_interface_secondaryip"
description: |-
  Second IP address of interface.
---

# fortimanager_object_fsp_vlan_interface_secondaryip
Second IP address of interface.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.

* `allowaccess` - Management access settings for the secondary IP address. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `auto-ipsec`, `radius-acct`, `probe-response`, `capwap`, `dnp`, `ftm`, `fabric`.

* `detectprotocol` - Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.

* `detectserver` - Gateway's ping server for this IP.
* `gwdetect` - Enable/disable detect gateway alive for first. Valid values: `disable`, `enable`.

* `ha_priority` - HA election priority for the PING server.
* `fosid` - ID.
* `ip` - Secondary IP address of the interface.
* `secip_relay_ip` - DHCP relay IP address.
* `ping_serv_status` - Ping-Serv-Status.
* `seq` - Seq.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFsp VlanInterfaceSecondaryip can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_interface_secondaryip.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
