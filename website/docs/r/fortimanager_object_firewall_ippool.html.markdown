---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_ippool"
description: |-
  Configure IPv4 IP pools.
---

# fortimanager_object_firewall_ippool
Configure IPv4 IP pools.

## Example Usage

```hcl
resource "fortimanager_object_firewall_ippool" "trname" {
  comments = "terraform-comment"
  endip    = "222.222.222.254"
  name     = "terraform-tefv"
  startip  = "222.222.222.0"
  type     = "overload"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.

* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.

* `associated_interface` - Associated interface name.
* `block_size` - Number of addresses in a block (64 to 4096, default = 128).
* `cgn_block_size` - Number of ports in a block(64 to 4096 in unit of 64, default = 128).
* `cgn_client_endip` - Final client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_client_ipv6shift` - IPv6 shift for fixed-allocation.(default 0)
* `cgn_client_startip` - First client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_fixedalloc` - Enable/disable fixed-allocation mode. Valid values: `disable`, `enable`.

* `cgn_overload` - Enable/disable overload mode. Valid values: `disable`, `enable`.

* `cgn_port_end` - Ending public port can be allocated.
* `cgn_port_start` - Starting public port can be allocated.
* `cgn_spa` - Enable/disable single port allocation mode. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endport` - Final port number (inclusive) in the range for the address pool (Default: 65533).
* `name` - IP pool name.
* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.

* `port_per_user` - Number of port for each user (32 to 60416, default = 0, auto).
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - First port number (inclusive) in the range for the address pool (Default: 5117).
* `subnet_broadcast_in_ippool` - Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.

* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`, `cgn-resource-allocation`.

* `utilization_alarm_clear` - Pool utilization alarm clear threshold (40-100).
* `utilization_alarm_raise` - Pool utilization alarm raise threshold (50-100).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.

* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.

* `associated_interface` - Associated interface name.
* `block_size` - Number of addresses in a block (64 to 4096, default = 128).
* `cgn_block_size` - Number of ports in a block(64 to 4096 in unit of 64, default = 128).
* `cgn_client_endip` - Final client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_client_ipv6shift` - Cgn-Client-Ipv6Shift.
* `cgn_client_startip` - First client IPv4 address (inclusive) (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `cgn_fixedalloc` - Enable/disable fixed-allocation mode. Valid values: `disable`, `enable`.

* `cgn_overload` - Enable/disable overload mode. Valid values: `disable`, `enable`.

* `cgn_port_end` - Ending public port can be allocated.
* `cgn_port_start` - Starting public port can be allocated.
* `cgn_spa` - Enable/disable single port allocation mode. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endport` - Endport.
* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.

* `port_per_user` - Port-Per-User.
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - Startport.
* `subnet_broadcast_in_ippool` - Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.

* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`, `cgn-resource-allocation`.

* `utilization_alarm_clear` - Pool utilization alarm clear threshold (40-100).
* `utilization_alarm_raise` - Pool utilization alarm raise threshold (50-100).

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Ippool can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_ippool.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
