---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_hyperscalepolicy64"
description: |-
  Configure IPv6 to IPv4 policies.
---

# fortimanager_packages_firewall_hyperscalepolicy64
Configure IPv6 to IPv4 policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Policy action. Valid values: `deny`, `accept`.

* `cgn_eif` - Enable/disable CGN endpoint independent filtering. Valid values: `disable`, `enable`.

* `cgn_eim` - Enable/disable CGN endpoint independent mapping. Valid values: `disable`, `enable`.

* `cgn_log_server_grp` - NP log server group name
* `cgn_resource_quota` - resource quota
* `cgn_session_quota` - session quota
* `comments` - Comment.
* `dstaddr` - Destination address name.
* `dstintf` - Destination interface name.
* `ippool` - Enable/disable policy64 IP pool. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `policy_offload` - Enable/disable hardware session setup for CGNAT. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `poolname` - Policy IP pool names.
* `service` - Service name.
* `srcaddr` - Source address name.
* `srcintf` - Source interface name.
* `status` - Enable/disable policy status. Valid values: `disable`, `enable`.

* `tcp_timeout_pid` - TCP timeout profile ID
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `udp_timeout_pid` - UDP timeout profile ID
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallHyperscalePolicy64 can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_hyperscalepolicy64.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
