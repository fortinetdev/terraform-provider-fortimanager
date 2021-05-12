---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_policy46"
description: |-
  Configure IPv4 to IPv6 policies.
---

# fortimanager_packages_firewall_policy46
Configure IPv4 to IPv6 policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Accept or deny traffic matching the policy. Valid values: `deny`, `accept`.

* `comments` - Comment.
* `dstaddr` - Destination address objects.
* `dstintf` - Destination interface name.
* `fixedport` - Enable/disable fixed port for this policy. Valid values: `disable`, `enable`.

* `ippool` - Enable/disable use of IP Pools for source NAT. Valid values: `disable`, `enable`.

* `logtraffic` - Enable/disable traffic logging for this policy. Valid values: `disable`, `enable`.

* `logtraffic_start` - Record logs when a session starts and ends. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `per_ip_shaper` - Per IP traffic shaper.
* `permit_any_host` - Enable/disable allowing any host. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `poolname` - IP Pool names.
* `schedule` - Schedule name.
* `service` - Service name.
* `srcaddr` - Source address objects.
* `srcintf` - Source interface name.
* `status` - Enable/disable this policy. Valid values: `disable`, `enable`.

* `tcp_mss_receiver` - TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
* `tcp_mss_sender` - TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallPolicy46 can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_policy46.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.