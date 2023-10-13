---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_hyperscalepolicy"
description: |-
  Configure IPv4 hyperscale-policies.
---

# fortimanager_packages_firewall_hyperscalepolicy
Configure IPv4 hyperscale-policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `action` - Policy action (allow/deny/ipsec). Valid values: `deny`, `accept`.

* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `disable`, `enable`.

* `cgn_eif` - Enable/Disable CGN endpoint independent filtering. Valid values: `disable`, `enable`.

* `cgn_eim` - Enable/Disable CGN endpoint independent mapping Valid values: `disable`, `enable`.

* `cgn_log_server_grp` - NP log server group name
* `cgn_resource_quota` - resource quota
* `cgn_session_quota` - session quota
* `comments` - Comment.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake. Valid values: `disable`, `enable`.

* `dstaddr` - Destination address and address group names.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstaddr6` - Destination IPv6 address name and address group names.
* `dstintf` - Outgoing (egress) interface.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all`, `check-new`.

* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `disable`, `enable`.

* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `policy_offload` - Enable/Disable hardware session setup for CGNAT. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 15000).
* `poolname` - IP Pool names.
* `poolname6` - IPv6 pool names.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `srcaddr` - Source address and address group names.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcaddr6` - Source IPv6 address name and address group names.
* `srcintf` - Incoming (ingress) interface.
* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `tcp_timeout_pid` - TCP timeout profile ID
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `udp_timeout_pid` - UDP timeout profile ID
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallHyperscalePolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_hyperscalepolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
