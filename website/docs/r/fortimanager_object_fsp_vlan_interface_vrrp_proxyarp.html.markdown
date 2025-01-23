---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_interface_vrrp_proxyarp"
description: |-
  VRRP Proxy ARP configuration.
---

# fortimanager_object_fsp_vlan_interface_vrrp_proxyarp
VRRP Proxy ARP configuration.

~> This resource is a sub resource for variable `proxy_arp` of resource `fortimanager_object_fsp_vlan_interface_vrrp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.
* `vrrp` - Vrrp.

* `fosid` - ID.
* `ip` - Set IP addresses of proxy ARP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFsp VlanInterfaceVrrpProxyArp can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE", "vrrp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_interface_vrrp_proxyarp.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
