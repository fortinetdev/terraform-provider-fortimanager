---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_portal_splitdns"
description: |-
  Split DNS for SSL VPN.
---

# fortimanager_object_vpn_ssl_web_portal_splitdns
Split DNS for SSL VPN.

~> This resource is a sub resource for variable `split_dns` of resource `fortimanager_object_vpn_ssl_web_portal`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `portal` - Portal.

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `domains` - Split DNS domains used for SSL-VPN clients separated by comma(,).
* `fosid` - ID.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectVpn SslWebPortalSplitDns can be imported using any of these accepted formats:
```
Set import_options = ["portal=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_portal_splitdns.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
