---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetservicecustom"
description: |-
  Configure custom Internet Services.
---

# fortimanager_object_firewall_internetservicecustom
Configure custom Internet Services.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `disable_entry`: `fortimanager_object_firewall_internetservicecustom_disableentry`
>- `entry`: `fortimanager_object_firewall_internetservicecustom_entry`



## Example Usage

```hcl
resource "fortimanager_object_firewall_internetservicecustom" "trname" {
  comment = "terraform-comment"
  name    = "terraform-tefv"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `disable_entry` - Disable-Entry. The structure of `disable_entry` block is documented below.
* `entry` - Entry. The structure of `entry` block is documented below.
* `master_service_id` - Internet Service ID in the Internet Service database.
* `fosid` - Internet Service ID.
* `name` - Internet Service name.
* `reputation` - Reputation level of the custom Internet Service.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `disable_entry` block supports:

* `id` - Disable entry ID.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `port` - Integer value for the TCP/IP port (0 - 65535).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `ip_range` block supports:

* `end_ip` - End IP address.
* `id` - Disable entry range ID.
* `start_ip` - Start IP address.

The `entry` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.

* `dst` - Destination address or address group name.
* `dst6` - Destination address6 or address6 group name.
* `id` - Entry ID(1-255).
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `port_range` block supports:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall InternetServiceCustom can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetservicecustom.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
