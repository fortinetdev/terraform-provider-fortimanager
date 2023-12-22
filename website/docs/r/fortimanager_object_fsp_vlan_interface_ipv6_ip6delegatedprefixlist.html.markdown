---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_interface_ipv6_ip6delegatedprefixlist"
description: |-
  Advertised IPv6 delegated prefix list.
---

# fortimanager_object_fsp_vlan_interface_ipv6_ip6delegatedprefixlist
Advertised IPv6 delegated prefix list.

~> This resource is a sub resource for variable `ip6_delegated_prefix_list` of resource `fortimanager_object_fsp_vlan_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.

* `autonomous_flag` - Enable/disable the autonomous flag. Valid values: `disable`, `enable`.

* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `onlink_flag` - Enable/disable the onlink flag. Valid values: `disable`, `enable`.

* `prefix_id` - Prefix ID.
* `rdnss` - Recursive DNS server option.
* `rdnss_service` - Recursive DNS service option. Valid values: `delegated`, `default`, `specify`.

* `subnet` - Add subnet ID to routing prefix.
* `upstream_interface` - Name of the interface that provides delegated information.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{prefix_id}}.

## Import

ObjectFsp VlanInterfaceIpv6Ip6DelegatedPrefixList can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_interface_ipv6_ip6delegatedprefixlist.labelname {{prefix_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
