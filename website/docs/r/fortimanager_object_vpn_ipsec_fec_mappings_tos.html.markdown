---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ipsec_fec_mappings_tos"
description: |-
  FEC redundancy mapping table for specific type of service (TOS).
---

# fortimanager_object_vpn_ipsec_fec_mappings_tos
FEC redundancy mapping table for specific type of service (TOS).

~> This resource is a sub resource for variable `tos` of resource `fortimanager_object_vpn_ipsec_fec_mappings`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `fec` - Fec.
* `mappings` - Mappings.

* `base` - Number of base FEC packets (1 - 40).
* `redundant` - Number of redundant FEC packets (0 - 20).
* `seqno` - Sequence number (1 - 8).
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seqno}}.

## Import

ObjectVpn IpsecFecMappingsTos can be imported using any of these accepted formats:
```
Set import_options = ["fec=YOUR_VALUE", "mappings=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ipsec_fec_mappings_tos.labelname {{seqno}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
