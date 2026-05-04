---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_certificate_hsmlocal"
description: |-
  Local certificates whose keys are stored on HSM.
---

# fortimanager_object_vpn_certificate_hsmlocal
Local certificates whose keys are stored on HSM.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `api_version` - API version for communicating with HSM. Valid values: `unknown`, `gch-default`.

* `certificate` - PEM format certificate.
* `comments` - Comment.
* `gch_cloud_service_name` - Cloud service config name to generate access token.
* `gch_cryptokey` - Google Cloud HSM cryptokey.
* `gch_cryptokey_algorithm` - Google Cloud HSM cryptokey algorithm. Valid values: `rsa-sign-pkcs1-2048-sha256`, `rsa-sign-pkcs1-3072-sha256`, `rsa-sign-pkcs1-4096-sha256`, `rsa-sign-pkcs1-4096-sha512`, `rsa-sign-pss-2048-sha256`, `rsa-sign-pss-3072-sha256`, `rsa-sign-pss-4096-sha256`, `rsa-sign-pss-4096-sha512`, `ec-sign-p256-sha256`, `ec-sign-p384-sha384`, `ec-sign-secp256k1-sha256`.

* `gch_cryptokey_version` - Google Cloud HSM cryptokey version.
* `gch_keyring` - Google Cloud HSM keyring.
* `gch_location` - Google Cloud HSM location.
* `gch_project` - Google Cloud HSM project ID.
* `gch_url` - Gch-Url.
* `name` - Name.
* `range` - Either a global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.

* `source` - Certificate source type. Valid values: `factory`, `user`, `bundle`.

* `tmp_cert_file` - Temporary certificate file.
* `vendor` - HSM vendor. Valid values: `unknown`, `gch`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn CertificateHsmLocal can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_certificate_hsmlocal.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
