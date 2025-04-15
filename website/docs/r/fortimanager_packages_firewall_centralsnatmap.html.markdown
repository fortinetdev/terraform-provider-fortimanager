---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_centralsnatmap"
description: |-
  Configure central SNAT policies.
---

# fortimanager_packages_firewall_centralsnatmap
Configure central SNAT policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_centralsnatmap" "labelname" {
  pkg       = "default"
  nat       = "enable"
  nat_port  = "0"
  orig_port = "0"
  policyid  = 1
  protocol  = 33
  status    = "enable"
  dst_addr  = ["all"]
  dstintf   = ["port3"]
  orig_addr = ["all"]
  srcintf   = ["port1"]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `comments` - Comment.
* `dst_addr` - Destination address name from available addresses.
* `dst_addr6` - IPv6 Destination address.
* `dst_port` - Destination port or port range (1 to 65535, 0 means any port).
* `dstintf` - Destination interface name from available interfaces.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools.
* `nat_ippool6` - IPv6 pools to be used for source NAT.
* `nat_port` - Translated port or port range (0 to 65535).
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `orig_addr` - Original address.
* `orig_addr6` - IPv6 Original address.
* `orig_port` - Original TCP port (0 to 65535).
* `policyid` - Policy ID.
* `port_preserve` - Enable/disable preservation of the original source port from source NAT if it has not been used. Valid values: `disable`, `enable`.

* `port_random` - Enable/disable random source port selection for source NAT. Valid values: `disable`, `enable`.

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
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_centralsnatmap.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
