---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_certificate_ocspserver"
description: |-
  OCSP server configuration.
---

# fortimanager_object_vpn_certificate_ocspserver
OCSP server configuration.

## Example Usage

```hcl
resource "fortimanager_object_vpn_certificate_ocspserver" "trname" {
  cert           = "ACCVRAIZ1"
  name           = "terr-vpn-cer-ocsp-server"
  source_ip      = "192.168.1.1"
  unavail_action = "revoke"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cert` - OCSP server certificate.
* `name` - OCSP server entry name.
* `secondary_cert` - Secondary OCSP server certificate.
* `secondary_url` - Secondary OCSP server URL.
* `source_ip` - Source IP address for communications to the OCSP server.
* `unavail_action` - Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.

* `url` - OCSP server URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn CertificateOcspServer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_certificate_ocspserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
