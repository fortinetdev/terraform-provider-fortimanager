---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address"
description: |-
  Configure IPv4 addresses.
---

# fortimanager_object_firewall_address
Configure IPv4 addresses.

## Example Usage

```hcl
resource "fortimanager_object_firewall_address" "trname" {
  color    = 0
  name     = "ssssss"
  obj_type = "ip"
  subnet = [
    "192.168.0.0",
    "255.255.255.255",
  ]
  type = "ipmask"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_image_base64` - _Image-Base64.
* `allow_routing` - Enable/disable use of this address in the static route configuration. Valid values: `disable`, `enable`.

* `associated_interface` - Network interface associated with address.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transition`, `infected`, `transient`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `fsso_group` - FSSO group(s).
* `global_object` - Global Object.
* `interface` - Name of interface whose IP address is to be used.
* `list` - List. The structure of `list` block is documented below.
* `macaddr` - Multiple MAC address ranges.
* `name` - Address name.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes. Valid values: `disable`, `enable`.

* `obj_id` - Object ID for NSX.
* `obj_tag` - Tag of dynamic address object.
* `obj_type` - Object type. Valid values: `ip`, `mac`.

* `organization` - Organization domain name (Syntax: organization/domain).
* `policy_group` - Policy group name.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.

* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - First MAC address in the range.
* `sub_type` - Sub-type of address. Valid values: `sdn`, `clearpass-spt`, `fsso`, `ems-tag`.

* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `tenant` - Tenant.
* `type` - Type of address. Valid values: `ipmask`, `iprange`, `fqdn`, `wildcard`, `geography`, `url`, `wildcard-fqdn`, `nsx`, `aws`, `dynamic`, `interface-subnet`, `mac`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `disable`, `enable`.

* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `allow_routing` - Enable/disable use of this address in the static route configuration. Valid values: `disable`, `enable`.

* `associated_interface` - Network interface associated with address.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transition`, `infected`, `transient`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Fabric-Object. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `fsso_group` - FSSO group(s).
* `global_object` - Global-Object.
* `interface` - Name of interface whose IP address is to be used.
* `macaddr` - Macaddr.
* `node_ip_only` - Node-Ip-Only. Valid values: `disable`, `enable`.

* `obj_id` - Object ID for NSX.
* `obj_tag` - Obj-Tag.
* `obj_type` - Obj-Type. Valid values: `ip`, `mac`.

* `organization` - Organization domain name (Syntax: organization/domain).
* `policy_group` - Policy group name.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.

* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - First MAC address in the range.
* `sub_type` - Sub-type of address. Valid values: `sdn`, `clearpass-spt`, `fsso`, `ems-tag`.

* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `tags` - Tags.
* `tenant` - Tenant.
* `type` - Type of address. Valid values: `ipmask`, `iprange`, `fqdn`, `wildcard`, `geography`, `url`, `wildcard-fqdn`, `nsx`, `aws`, `dynamic`, `interface-subnet`, `mac`.

* `url` - Url.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `disable`, `enable`.

* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `list` block supports:

* `ip` - IP.
* `net_id` - Network ID.
* `obj_id` - Object ID.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Address can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
