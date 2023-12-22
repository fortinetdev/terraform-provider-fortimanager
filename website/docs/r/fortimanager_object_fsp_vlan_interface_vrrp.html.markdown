---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_interface_vrrp"
description: |-
  VRRP configuration.
---

# fortimanager_object_fsp_vlan_interface_vrrp
VRRP configuration.

~> This resource is a sub resource for variable `vrrp` of resource `fortimanager_object_fsp_vlan_interface`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`proxy_arp`: `fortimanager_object_fsp_vlan_interface_vrrp_proxyarp`



## Example Usage

```hcl
resource "fortimanager_object_fsp_vlan_interface_vrrp" "trname" {
  vlan                 = fortimanager_object_fsp_vlan.trname.name
  adv_interval         = 100
  ignore_default_route = "disable"
  preempt              = "enable"
  vrid                 = 3
  depends_on           = [fortimanager_object_fsp_vlan.trname]
}

resource "fortimanager_object_fsp_vlan" "trname" {
  name   = "terr-fsp-vlan"
  vlanid = 101
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `ignore_default_route` - Enable/disable ignoring of default route when checking destination. Valid values: `disable`, `enable`.

* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `proxy_arp` - Proxy-Arp. The structure of `proxy_arp` block is documented below.
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable this VRRP configuration. Valid values: `disable`, `enable`.

* `version` - VRRP version. Valid values: `2`, `3`.

* `vrdst` - Monitor the route to this destination.
* `vrdst_priority` - Priority of the virtual router when the virtual router destination becomes unreachable (0 - 254).
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip` - IP address of the virtual router.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `proxy_arp` block supports:

* `id` - ID.
* `ip` - Set IP addresses of proxy ARP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vrid}}.

## Import

ObjectFsp VlanInterfaceVrrp can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_interface_vrrp.labelname {{vrid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
