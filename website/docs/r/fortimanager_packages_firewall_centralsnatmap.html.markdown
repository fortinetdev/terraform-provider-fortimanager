---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_centralsnatmap"
description: |-
  Configure IPv4 and IPv6 central SNAT policies.
---

# fortimanager_packages_firewall_centralsnatmap
Configure IPv4 and IPv6 central SNAT policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `comments` - Comment.
* `dst_addr` - IPv4 Destination address.
* `dst_addr6` - IPv6 Destination address.
* `dstintf` - Destination interface name from available interfaces.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools.
* `nat_ippool6` - IPv6 pools to be used for source NAT.
* `nat_port` - Translated port or port range (1 to 65535, 0 means any port).
* `orig_addr` - IPv4 Original address.
* `orig_addr6` - IPv6 Original address.
* `orig_port` - Original TCP port (1 to 65535, 0 means any port).
* `policyid` - Policy ID.
* `protocol` - Integer value for the protocol type (0 - 255).
* `srcintf` - Source interface name from available interfaces.
* `status` - Enable/disable the active status of this policy. Valid values: `disable`, `enable`.

* `type` - IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallCentralSnatMap can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_centralsnatmap.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
