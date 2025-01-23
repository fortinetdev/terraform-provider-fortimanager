---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_fsp_vlan_dynamic_mapping_dhcpserver_reservedaddress"
description: |-
  Options for the DHCP server to assign IP settings to specific MAC addresses.
---

# fortimanager_object_fsp_vlan_dynamic_mapping_dhcpserver_reservedaddress
Options for the DHCP server to assign IP settings to specific MAC addresses.

~> This resource is a sub resource for variable `reserved_address` of resource `fortimanager_object_fsp_vlan_dynamic_mapping_dhcpserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vlan` - Vlan.
* `dynamic_mapping_name` - Dynamic Mapping Name.
* `dynamic_mapping_vdom` - Dynamic Mapping Vdom.

* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.

* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `description` - Description.
* `fosid` - ID.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFsp VlanDynamicMappingDhcpServerReservedAddress can be imported using any of these accepted formats:
```
Set import_options = ["vlan=YOUR_VALUE", "dynamic_mapping_name=YOUR_VALUE", "dynamic_mapping_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_fsp_vlan_dynamic_mapping_dhcpserver_reservedaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
