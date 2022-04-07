---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_hyperscalepolicy46"
description: |-
  Configure IPv4 to IPv6 policies.
---

# fortimanager_packages_firewall_hyperscalepolicy46
Configure IPv4 to IPv6 policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Accept or deny traffic matching the policy. Valid values: `deny`, `accept`.

* `cgn_log_server_grp` - NP log server group name
* `comments` - Comment.
* `dstaddr` - Destination address objects.
* `dstintf` - Destination interface name.
* `name` - Policy name.
* `policy_offload` - Enable/disable offloading policy configuration to CP processors. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `service` - Service name.
* `srcaddr` - Source address objects.
* `srcintf` - Source interface name.
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `tcp_timeout_pid` - TCP timeout profile ID
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `udp_timeout_pid` - UDP timeout profile ID
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallHyperscalePolicy46 can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_hyperscalepolicy46.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
