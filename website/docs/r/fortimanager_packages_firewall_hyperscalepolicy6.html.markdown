---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_hyperscalepolicy6"
description: |-
  Configure IPv6 policies.
---

# fortimanager_packages_firewall_hyperscalepolicy6
Configure IPv6 policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `action` - Policy action (allow/deny/ipsec). Valid values: `deny`, `accept`, `ipsec`.

* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `disable`, `enable`.

* `cgn_log_server_grp` - NP log server group name
* `comments` - Comment.
* `dstaddr` - Destination address and address group names.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstintf` - Outgoing (egress) interface.
* `name` - Policy name.
* `policy_offload` - Enable/disable offloading policy configuration to CP processors. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `srcaddr` - Source address and address group names.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcintf` - Incoming (ingress) interface.
* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `tcp_timeout_pid` - TCP timeout profile ID
* `traffic_shaper` - Reverse traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `udp_timeout_pid` - UDP timeout profile ID
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallHyperscalePolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_hyperscalepolicy6.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
