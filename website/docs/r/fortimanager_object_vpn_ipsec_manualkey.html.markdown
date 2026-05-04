---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ipsec_manualkey"
description: |-
  Configure IPsec manual keys.
---

# fortimanager_object_vpn_ipsec_manualkey
Configure IPsec manual keys.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `authentication` - Authentication algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `md5`, `sha1`, `sha256`, `sha384`, `sha512`.

* `authkey` - Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enckey` - Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `encryption` - Encryption algorithm. Must be the same for both ends of the tunnel. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`, `aria128`, `aria192`, `aria256`, `seed`.

* `interface` - Name of the physical, aggregate, or VLAN interface.
* `local_gw` - Local gateway.
* `localspi` - Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `name` - IPsec tunnel name.
* `npu_offload` - Enable/disable NPU offloading. Valid values: `disable`, `enable`.

* `remote_gw` - Peer gateway.
* `remotespi` - Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn IpsecManualkey can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ipsec_manualkey.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
