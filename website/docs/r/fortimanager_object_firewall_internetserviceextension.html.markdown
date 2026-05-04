---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetserviceextension"
description: |-
  Configure Internet Services Extension.
---

# fortimanager_object_firewall_internetserviceextension
Configure Internet Services Extension.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `disable_entry`: `fortimanager_object_firewall_internetserviceextension_disableentry`
>- `entry`: `fortimanager_object_firewall_internetserviceextension_entry`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `disable_entry` - Disable-Entry. The structure of `disable_entry` block is documented below.
* `entry` - Entry. The structure of `entry` block is documented below.
* `fosid` - Internet Service ID in the Internet Service database.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `disable_entry` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `id` - Disable entry ID.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ip6_range` - Ip6-Range. The structure of `ip6_range` block is documented below.
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `ip_range` block supports:

* `end_ip` - End IPv4 address.
* `id` - Disable entry range ID.
* `start_ip` - Start IPv4 address.

The `ip6_range` block supports:

* `end_ip6` - End IPv6 address.
* `id` - Disable entry range ID.
* `start_ip6` - Start IPv6 address.

The `port_range` block supports:

* `end_port` - Ending TCP/UDP/SCTP destination port (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (0 to 65535).

The `entry` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `dst` - Destination address or address group name.
* `dst6` - Destination address6 or address6 group name.
* `id` - Entry ID(1-255).
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `port_range` block supports:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall InternetServiceExtension can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetserviceextension.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
