---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod"
description: |-
  EAP Methods.
---

# fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod
EAP Methods.

~> This resource is a sub resource for variable `eap_method` of resource `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`auth_param`: `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod" "trname" {
  index          = 2
  method         = "eap-md5"
  depends_on     = [fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist.trname2]
  anqp_nai_realm = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname2.name
  nai_list       = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist.trname2.name
}

resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist" "trname2" {
  name           = "terr-nailist"
  depends_on     = [fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname2]
  anqp_nai_realm = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname2.name
}

resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm" "trname2" {
  name = "terr-anqpnairealm"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `anqp_nai_realm` - Anqp Nai Realm.
* `nai_list` - Nai List.

* `auth_param` - Auth-Param. The structure of `auth_param` block is documented below.
* `index` - EAP method index.
* `method` - EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `auth_param` block supports:

* `id` - ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.

* `index` - Param index.
* `val` - Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

ObjectWirelessController Hotspot20AnqpNaiRealmNaiListEapMethod can be imported using any of these accepted formats:
```
Set import_options = ["anqp_nai_realm=YOUR_VALUE", "nai_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
