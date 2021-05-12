---
subcategory: "ObjectVpn"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_certificate_ca"
description: |-
  CA certificate.
---

# fortimanager_object_vpn_certificate_ca
CA certificate.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_private_key` - _Private_Key.
* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `ca` - CA certificate as a PEM file.
* `last_updated` - Time at which CA was last updated.
* `name` - Name.
* `range` - Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.

* `scep_url` - URL of the SCEP server.
* `source` - CA certificate source type. Valid values: `factory`, `user`, `bundle`, `fortiguard`.

* `source_ip` - Source IP address for communications to the SCEP server.
* `ssl_inspection_trusted` - Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn CertificateCa can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_certificate_ca.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
