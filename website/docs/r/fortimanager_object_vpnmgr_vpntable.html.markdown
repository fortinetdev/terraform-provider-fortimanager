---
subcategory: "ObjectVpnmgr"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpnmgr_vpntable"
description: |-
  ObjectVpnmgr Vpntable
---

# fortimanager_object_vpnmgr_vpntable
ObjectVpnmgr Vpntable

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `authmethod` - Authmethod. Valid values: `psk`, `rsa-signature`, `signature`.

* `auto_zone_policy` - Auto-Zone-Policy. Valid values: `disable`, `enable`.

* `certificate` - Certificate.
* `description` - Description.
* `dpd` - Dpd. Valid values: `disable`, `enable`, `on-idle`, `on-demand`.

* `dpd_retrycount` - Dpd-Retrycount.
* `dpd_retryinterval` - Dpd-Retryinterval.
* `fcc_enforcement` - Fcc-Enforcement. Valid values: `disable`, `enable`.

* `hub2spoke_zone` - Hub2Spoke-Zone.
* `ike_version` - Ike-Version. Valid values: `1`, `2`.

* `ike1dhgroup` - Ike1Dhgroup. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`, `32`.

* `ike1dpd` - Ike1Dpd. Valid values: `disable`, `enable`.

* `ike1keylifesec` - Ike1Keylifesec.
* `ike1localid` - Ike1Localid.
* `ike1mode` - Ike1Mode. Valid values: `main`, `aggressive`.

* `ike1natkeepalive` - Ike1Natkeepalive.
* `ike1nattraversal` - Ike1Nattraversal. Valid values: `disable`, `enable`, `forced`.

* `ike1proposal` - Ike1Proposal. Valid values: `des-md5`, `des-sha1`, `3des-md5`, `3des-sha1`, `aes128-md5`, `aes128-sha1`, `aes192-md5`, `aes192-sha1`, `aes256-md5`, `aes256-sha1`, `des-sha256`, `3des-sha256`, `aes128-sha256`, `aes192-sha256`, `aes256-sha256`, `des-sha384`, `des-sha512`, `3des-sha384`, `3des-sha512`, `aes128-sha384`, `aes128-sha512`, `aes192-sha384`, `aes192-sha512`, `aes256-sha384`, `aes256-sha512`, `aria128-md5`, `aria128-sha1`, `aria128-sha256`, `aria128-sha384`, `aria128-sha512`, `aria192-md5`, `aria192-sha1`, `aria192-sha256`, `aria192-sha384`, `aria192-sha512`, `aria256-md5`, `aria256-sha1`, `aria256-sha256`, `aria256-sha384`, `aria256-sha512`, `seed-md5`, `seed-sha1`, `seed-sha256`, `seed-sha384`, `seed-sha512`, `aes128gcm-prfsha1`, `aes128gcm-prfsha256`, `aes128gcm-prfsha384`, `aes128gcm-prfsha512`, `aes256gcm-prfsha1`, `aes256gcm-prfsha256`, `aes256gcm-prfsha384`, `aes256gcm-prfsha512`, `chacha20poly1305-prfsha1`, `chacha20poly1305-prfsha256`, `chacha20poly1305-prfsha384`, `chacha20poly1305-prfsha512`.

* `ike2autonego` - Ike2Autonego. Valid values: `disable`, `enable`.

* `ike2dhgroup` - Ike2Dhgroup. Valid values: `1`, `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `27`, `28`, `29`, `30`, `31`, `32`.

* `ike2keepalive` - Ike2Keepalive. Valid values: `disable`, `enable`.

* `ike2keylifekbs` - Ike2Keylifekbs.
* `ike2keylifesec` - Ike2Keylifesec.
* `ike2keylifetype` - Ike2Keylifetype. Valid values: `seconds`, `kbs`, `both`.

* `ike2proposal` - Ike2Proposal. Valid values: `null-md5`, `null-sha1`, `des-null`, `3des-null`, `des-md5`, `des-sha1`, `3des-md5`, `3des-sha1`, `aes128-md5`, `aes128-sha1`, `aes192-md5`, `aes192-sha1`, `aes256-md5`, `aes256-sha1`, `aes128-null`, `aes192-null`, `aes256-null`, `null-sha256`, `des-sha256`, `3des-sha256`, `aes128-sha256`, `aes192-sha256`, `aes256-sha256`, `des-sha384`, `des-sha512`, `3des-sha384`, `3des-sha512`, `aes128-sha384`, `aes128-sha512`, `aes192-sha384`, `aes192-sha512`, `aes256-sha384`, `aes256-sha512`, `null-sha384`, `null-sha512`, `aria128-null`, `aria128-md5`, `aria128-sha1`, `aria128-sha256`, `aria128-sha384`, `aria128-sha512`, `aria192-null`, `aria192-md5`, `aria192-sha1`, `aria192-sha256`, `aria192-sha384`, `aria192-sha512`, `aria256-null`, `aria256-md5`, `aria256-sha1`, `aria256-sha256`, `aria256-sha384`, `aria256-sha512`, `seed-null`, `seed-md5`, `seed-sha1`, `seed-sha256`, `seed-sha384`, `seed-sha512`, `aes128gcm`, `aes256gcm`, `chacha20poly1305`.

* `inter_vdom` - Inter-Vdom. Valid values: `disable`, `enable`.

* `intf_mode` - Intf-Mode. Valid values: `off`, `on`.

* `localid_type` - Localid-Type. Valid values: `auto`, `fqdn`, `user-fqdn`, `keyid`, `address`, `asn1dn`.

* `name` - Name.
* `negotiate_timeout` - Negotiate-Timeout.
* `network_id` - Network-Id.
* `network_overlay` - Network-Overlay. Valid values: `disable`, `enable`.

* `npu_offload` - Npu-Offload. Valid values: `disable`, `enable`.

* `pfs` - Pfs. Valid values: `disable`, `enable`.

* `psk_auto_generate` - Psk-Auto-Generate. Valid values: `disable`, `enable`.

* `psksecret` - Psksecret.
* `replay` - Replay. Valid values: `disable`, `enable`.

* `rsa_certificate` - Rsa-Certificate.
* `spoke2hub_zone` - Spoke2Hub-Zone.
* `topology` - Topology. Valid values: `meshed`, `star`, `dialup`.

* `vpn_zone` - Vpn-Zone.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpnmgr Vpntable can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpnmgr_vpntable.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
