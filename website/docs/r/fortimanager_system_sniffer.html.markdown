---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_sniffer"
description: |-
  Interface sniffer.
---

# fortimanager_system_sniffer
Interface sniffer.

## Argument Reference


The following arguments are supported:


* `host` - Hosts to filter for in sniffer traffic
* `fosid` - Sniffer ID.
* `interface` - Interface.
* `ipv6` - Enable/disable sniffing IPv6 packets. disable - Disable sniffer for IPv6 packets. enable - Enable sniffer for IPv6 packets. Valid values: `disable`, `enable`.

* `max_packet_count` - Maximum packet count (1000000, default = 4000).
* `non_ip` - Enable/disable sniffing non-IP packets. disable - Disable sniffer for non-IP packets. enable - Enable sniffer for non-IP packets. Valid values: `disable`, `enable`.

* `port` - Ports to sniff (Format examples: 10, 100-200).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `vlan` - List of VLANs to sniff.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Sniffer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_sniffer.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

