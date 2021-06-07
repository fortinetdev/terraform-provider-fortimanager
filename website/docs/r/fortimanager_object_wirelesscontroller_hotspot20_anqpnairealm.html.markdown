---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm"
description: |-
  Configure network access identifier (NAI) realm.
---

# fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm
Configure network access identifier (NAI) realm.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `nai_list` - Nai-List. The structure of `nai_list` block is documented below.
* `name` - NAI realm list name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `nai_list` block supports:

* `eap_method` - Eap-Method. The structure of `eap_method` block is documented below.
* `encoding` - Enable/disable format in accordance with IETF RFC 4282. Valid values: `disable`, `enable`.

* `nai_realm` - Configure NAI realms (delimited by a semi-colon character).
* `name` - NAI realm name.

The `eap_method` block supports:

* `auth_param` - Auth-Param. The structure of `auth_param` block is documented below.
* `index` - EAP method index.
* `method` - EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.


The `auth_param` block supports:

* `id` - ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.

* `index` - Param index.
* `val` - Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20AnqpNaiRealm can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
