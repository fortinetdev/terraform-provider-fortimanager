---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_anqpipaddresstype"
description: |-
  Configure IP address type availability.
---

# fortimanager_object_wirelesscontroller_hotspot20_anqpipaddresstype
Configure IP address type availability.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_anqpipaddresstype" "trname" {
  ipv4_address_type = "not-available"
  ipv6_address_type = "available"
  name              = "terr-wictl-hot20-anqp-ip-addrtype"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ipv4_address_type` - IPv4 address type. Valid values: `not-available`, `not-known`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`.

* `ipv6_address_type` - IPv6 address type. Valid values: `not-available`, `available`, `not-known`.

* `name` - IP type name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20AnqpIpAddressType can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_anqpipaddresstype.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
