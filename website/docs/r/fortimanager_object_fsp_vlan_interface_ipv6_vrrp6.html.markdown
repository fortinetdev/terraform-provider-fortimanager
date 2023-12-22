---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_interface_ipv6_vrrp6"
description: |-
  IPv6 VRRP configuration.
---

# fortimanager_object_fsp_vlan_interface_ipv6_vrrp6
IPv6 VRRP configuration.

~> This resource is a sub resource for variable `vrrp6` of resource `fortimanager_object_fsp_vlan_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_fsp_vlan_interface_ipv6_vrrp6" "trname" {
  vlan         = fortimanager_object_fsp_vlan.trname.name
  accept_mode  = "enable"
  adv_interval = 100
  vrid         = 3
  depends_on   = [fortimanager_object_fsp_vlan.trname]
}

resource "fortimanager_object_fsp_vlan" "trname" {
  name   = "terr-fsp-vlan"
  vlanid = 104
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.

* `accept_mode` - Enable/disable accept mode. Valid values: `disable`, `enable`.

* `adv_interval` - Advertisement interval (1 - 255 seconds).
* `preempt` - Enable/disable preempt mode. Valid values: `disable`, `enable`.

* `priority` - Priority of the virtual router (1 - 255).
* `start_time` - Startup time (1 - 255 seconds).
* `status` - Enable/disable VRRP. Valid values: `disable`, `enable`.

* `vrdst6` - Monitor the route to this destination.
* `vrgrp` - VRRP group ID (1 - 65535).
* `vrid` - Virtual router identifier (1 - 255).
* `vrip6` - IPv6 address of the virtual router.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vrid}}.

## Import

ObjectFsp VlanInterfaceIpv6Vrrp6 can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_interface_ipv6_vrrp6.labelname {{vrid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
