---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address_dynamic_mapping"
description: |-
  Configure IPv4 addresses.
---

# fortimanager_object_firewall_address_dynamic_mapping
Configure IPv4 addresses.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_firewall_address`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `address` - Address.

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `allow_routing` - Enable/disable use of this address in the static route configuration. Valid values: `disable`, `enable`.

* `associated_interface` - Network interface associated with address.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transition`, `infected`, `transient`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `dirty` - To be deleted address. Valid values: `dirty`, `clean`.

* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Fabric-Object. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `fsso_group` - FSSO group(s).
* `global_object` - Global-Object.
* `hw_model` - Dynamic address matching hardware model.
* `hw_vendor` - Dynamic address matching hardware vendor.
* `interface` - Name of interface whose IP address is to be used.
* `macaddr` - Macaddr.
* `node_ip_only` - Node-Ip-Only. Valid values: `disable`, `enable`.

* `obj_id` - Object ID for NSX.
* `obj_tag` - Obj-Tag.
* `obj_type` - Obj-Type. Valid values: `ip`, `mac`.

* `organization` - Organization domain name (Syntax: organization/domain).
* `os` - Dynamic address matching operating system.
* `pattern_end` - Pattern-End.
* `pattern_start` - Pattern-Start.
* `policy_group` - Policy group name.
* `route_tag` - route-tag address.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.

* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - First MAC address in the range.
* `sub_type` - Sub-type of address. Valid values: `sdn`, `clearpass-spt`, `fsso`, `ems-tag`.

* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `sw_version` - Dynamic address matching software version.
* `tag_detection_level` - Tag detection level of dynamic address object.
* `tag_type` - Tag type of dynamic address object.
* `tags` - Tags.
* `tenant` - Tenant.
* `type` - Type of address. Valid values: `ipmask`, `iprange`, `fqdn`, `wildcard`, `geography`, `url`, `wildcard-fqdn`, `nsx`, `aws`, `dynamic`, `interface-subnet`, `mac`.

* `url` - Url.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `disable`, `enable`.

* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectFirewall AddressDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["address=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
