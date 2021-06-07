---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_multicastpolicy6"
description: |-
  Configure IPv6 multicast NAT policies.
---

# fortimanager_packages_firewall_multicastpolicy6
Configure IPv6 multicast NAT policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Accept or deny traffic matching the policy. Valid values: `deny`, `accept`.

* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `disable`, `enable`.

* `dstaddr` - IPv6 destination address name.
* `dstintf` - IPv6 destination interface name.
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
* `fosid` - Policy ID (0 - 4294967294).
* `logtraffic` - Enable/disable logging traffic accepted by this policy. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `srcaddr` - IPv6 source address name.
* `srcintf` - IPv6 source interface name.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Packages FirewallMulticastPolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_multicastpolicy6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
