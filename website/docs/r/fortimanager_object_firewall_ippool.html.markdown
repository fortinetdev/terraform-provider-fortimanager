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

* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.

* `associated_interface` - Associated interface name.
* `block_size` - Number of addresses in a block (64 to 4096, default = 128).
* `comments` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `name` - IP pool name.
* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.

* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`, `cgn-resource-allocation`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.

* `associated_interface` - Associated interface name.
* `block_size` - Number of addresses in a block (64 to 4096, default = 128).
* `cgn_block_size` - Cgn-Block-Size.
* `cgn_client_endip` - Cgn-Client-Endip.
* `cgn_client_startip` - Cgn-Client-Startip.
* `cgn_fixedalloc` - Cgn-Fixedalloc. Valid values: `disable`, `enable`.

* `cgn_overload` - Cgn-Overload. Valid values: `disable`, `enable`.

* `cgn_port_end` - Cgn-Port-End.
* `cgn_port_start` - Cgn-Port-Start.
* `cgn_spa` - Cgn-Spa. Valid values: `disable`, `enable`.

* `comments` - Comment.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endport` - Endport.
* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.

* `port_per_user` - Port-Per-User.
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - Startport.
* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`, `cgn-resource-allocation`.

* `utilization_alarm_clear` - Utilization-Alarm-Clear.
* `utilization_alarm_raise` - Utilization-Alarm-Raise.

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
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
