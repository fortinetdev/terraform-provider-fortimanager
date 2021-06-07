---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetserviceaddition"
description: |-
  Configure Internet Services Addition.
---

# fortimanager_object_firewall_internetserviceaddition
Configure Internet Services Addition.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `entry` - Entry. The structure of `entry` block is documented below.
* `fosid` - Internet Service ID in the Internet Service database.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entry` block supports:

* `id` - Entry ID(1-255).
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `port_range` block supports:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall InternetServiceAddition can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetserviceaddition.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
