---
subcategory: "ObjectSystem"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_geoipoverride"
description: |-
  Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
---

# fortimanager_object_system_geoipoverride
Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `country_id` - Two character Country ID code.
* `description` - Description.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ip6_range` - Ip6-Range. The structure of `ip6_range` block is documented below.
* `name` - Location name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `id` - ID of individual entry in the IP range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).

The `ip6_range` block supports:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `id` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem GeoipOverride can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_geoipoverride.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
