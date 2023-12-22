---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_interface_ipv6_ip6extraaddr"
description: |-
  Extra IPv6 address prefixes of interface.
---

# fortimanager_object_fsp_vlan_interface_ipv6_ip6extraaddr
Extra IPv6 address prefixes of interface.

~> This resource is a sub resource for variable `ip6_extra_addr` of resource `fortimanager_object_fsp_vlan_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.

* `prefix` - IPv6 address prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{prefix}}.

## Import

ObjectFsp VlanInterfaceIpv6Ip6ExtraAddr can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_interface_ipv6_ip6extraaddr.labelname {{prefix}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
