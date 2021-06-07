---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpnmgr_node"
description: |-
  VPN node for VPN Manager. Must specify vpntable and scope member.
---

# fortimanager_object_vpnmgr_node
VPN node for VPN Manager. Must specify vpntable and scope member.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `add_route` - Add-Route. Valid values: `disable`, `enable`.

* `assign_ip` - Assign-Ip. Valid values: `disable`, `enable`.

* `assign_ip_from` - Assign-Ip-From. Valid values: `range`, `usrgrp`, `dhcp`, `name`.

* `authpasswd` - Authpasswd.
* `authusr` - Authusr.
* `authusrgrp` - Authusrgrp.
* `auto_configuration` - Auto-Configuration. Valid values: `disable`, `enable`.

* `automatic_routing` - Automatic_Routing. Valid values: `disable`, `enable`.

* `banner` - Banner.
* `default_gateway` - Default-Gateway.
* `dhcp_server` - Dhcp-Server. Valid values: `disable`, `enable`.

* `dns_mode` - Dns-Mode. Valid values: `auto`, `manual`.

* `dns_service` - Dns-Service. Valid values: `default`, `specify`, `local`.

* `domain` - Domain.
* `exchange_interface_ip` - Exchange-Interface-Ip. Valid values: `disable`, `enable`.

* `extgw` - Extgw.
* `extgw_hubip` - Extgw_Hubip.
* `extgw_p2_per_net` - Extgw_P2_Per_Net. Valid values: `disable`, `enable`.

* `extgwip` - Extgwip.
* `hub_public_ip` - Hub-Public-Ip.
* `hub_iface` - Hub_Iface.
* `fosid` - Id.
* `iface` - Iface.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ipsec_lease_hold` - Ipsec-Lease-Hold.
* `ipv4_dns_server1` - Ipv4-Dns-Server1.
* `ipv4_dns_server2` - Ipv4-Dns-Server2.
* `ipv4_dns_server3` - Ipv4-Dns-Server3.
* `ipv4_end_ip` - Ipv4-End-Ip.
* `ipv4_exclude_range` - Ipv4-Exclude-Range. The structure of `ipv4_exclude_range` block is documented below.
* `ipv4_netmask` - Ipv4-Netmask.
* `ipv4_split_exclude` - Ipv4-Split-Exclude.
* `ipv4_split_include` - Ipv4-Split-Include.
* `ipv4_start_ip` - Ipv4-Start-Ip.
* `ipv4_wins_server1` - Ipv4-Wins-Server1.
* `ipv4_wins_server2` - Ipv4-Wins-Server2.
* `local_gw` - Local-Gw.
* `localid` - Localid.
* `mode_cfg` - Mode-Cfg. Valid values: `disable`, `enable`.

* `mode_cfg_ip_version` - Mode-Cfg-Ip-Version. Valid values: `4`, `6`.

* `net_device` - Net-Device. Valid values: `disable`, `enable`.

* `peer` - Peer.
* `peergrp` - Peergrp.
* `peerid` - Peerid.
* `peertype` - Peertype. Valid values: `any`, `one`, `dialup`, `peer`, `peergrp`.

* `protected_subnet` - Protected_Subnet. The structure of `protected_subnet` block is documented below.
* `public_ip` - Public-Ip.
* `role` - Role. Valid values: `hub`, `spoke`.

* `route_overlap` - Route-Overlap. Valid values: `use-old`, `use-new`, `allow`.

* `spoke_zone` - Spoke-Zone.
* `summary_addr` - Summary_Addr. The structure of `summary_addr` block is documented below.
* `tunnel_search` - Tunnel-Search. Valid values: `selectors`, `nexthop`.

* `unity_support` - Unity-Support. Valid values: `disable`, `enable`.

* `usrgrp` - Usrgrp.
* `vpn_interface_priority` - Vpn-Interface-Priority.
* `vpn_zone` - Vpn-Zone.
* `vpntable` - Vpntable.
* `xauthtype` - Xauthtype. Valid values: `disable`, `client`, `pap`, `chap`, `auto`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `end_ip` - End-Ip.
* `id` - Id.
* `start_ip` - Start-Ip.

The `ipv4_exclude_range` block supports:

* `end_ip` - End-Ip.
* `id` - Id.
* `start_ip` - Start-Ip.

The `protected_subnet` block supports:

* `addr` - Addr.
* `seq` - Seq.

The `summary_addr` block supports:

* `addr` - Addr.
* `priority` - Priority.
* `seq` - Seq.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVpnmgr Node can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpnmgr_node.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
