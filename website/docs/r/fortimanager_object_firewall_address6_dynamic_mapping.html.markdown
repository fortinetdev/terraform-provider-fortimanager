---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_address6_dynamic_mapping"
description: |-
  Configure IPv6 firewall addresses.
---

# fortimanager_object_firewall_address6_dynamic_mapping
Configure IPv6 firewall addresses.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_firewall_address6`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `subnet_segment`: `fortimanager_object_firewall_address6_dynamic_mapping_subnetsegment`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `address6` - Address6.

* `_image_base64` - _Image-Base64.
* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `country` - Country.
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Fabric-Object. Valid values: `disable`, `enable`.

* `filter` - Match criteria filter.
* `fqdn` - Fully qualified domain name.
* `global_object` - Global-Object.
* `host` - Host Address.
* `host_type` - Host type. Valid values: `any`, `specific`.

* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `macaddr` - Macaddr.
* `obj_id` - Object ID for NSX.
* `route_tag` - route-tag address.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `all`, `private`, `public`.

* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `start_mac` - First MAC address in the range.
* `subnet_segment` - Subnet-Segment. The structure of `subnet_segment` block is documented below.
* `tags` - Tags.
* `template` - IPv6 address template.
* `tenant` - Tenant.
* `type` - Type of IPv6 address object (default = ipprefix). Valid values: `ipprefix`, `iprange`, `nsx`, `dynamic`, `fqdn`, `template`, `mac`, `geography`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable the visibility of the object in the GUI. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `subnet_segment` block supports:

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any`, `specific`.

* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectFirewall Address6DynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["address6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_address6_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
