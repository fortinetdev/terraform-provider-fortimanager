---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam"
description: |-
  EAP auth param.
---

# fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam
EAP auth param.

~> This resource is a sub resource for variable `auth_param` of resource `fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam" "trname" {
  index          = 2
  val            = "eap-md5"
  anqp_nai_realm = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname.name
  nai_list       = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist.trname2.name
  eap_method     = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod.trname.index
  depends_on     = [fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod.trname]
}

resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod" "trname" {
  index          = 2
  anqp_nai_realm = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname.name
  nai_list       = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist.trname2.name
  depends_on     = [fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist.trname2]
}

resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist" "trname2" {
  name           = "terr-nailist"
  anqp_nai_realm = fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname.name
  depends_on     = [fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm.trname]
}

resource "fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm" "trname" {
  name = "terr-anqpnairealm"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `anqp_nai_realm` - Anqp Nai Realm.
* `nai_list` - Nai List.
* `eap_method` - Eap Method.

* `fosid` - ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.

* `index` - Param index.
* `val` - Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

ObjectWirelessController Hotspot20AnqpNaiRealmNaiListEapMethodAuthParam can be imported using any of these accepted formats:
```
Set import_options = ["anqp_nai_realm=YOUR_VALUE", "nai_list=YOUR_VALUE", "eap_method=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_anqpnairealm_nailist_eapmethod_authparam.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
