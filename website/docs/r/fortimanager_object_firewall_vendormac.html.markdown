---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vendormac"
description: |-
  Show vendor and the MAC address they have.
---

# fortimanager_object_firewall_vendormac
Show vendor and the MAC address they have.

## Example Usage

```hcl
resource "fortimanager_object_firewall_vendormac" "trname" {
  fosid      = 23
  mac_number = 12
  name       = "demo"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fosid` - Id.
* `mac_number` - Mac-Number.
* `name` - Name.
* `obsolete` - Obsolete.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall VendorMac can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vendormac.labelname ObjectFirewallVendorMac
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
