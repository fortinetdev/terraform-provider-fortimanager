---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_certificate_ca"
description: |-
  CA certificate.
---

# fortimanager_object_vpn_certificate_ca
CA certificate.

## Example Usage

```hcl
resource "fortimanager_object_vpn_certificate_ca" "trname" {
  auto_update_days         = 5
  auto_update_days_warning = 1
  ca                       = "-----BEGIN CERTIFICATE-----\nMIIDADCCAeigAwIBAgIgRDM5NzdDMTU4MjdDQTk4Njc2RkU4RjZGQzIzMDU3MDUw\nDQYJKoZIhvcNAQEFBQAwKzEWMBQGA1UEChMNRm9ydGluZXQgTHRkLjERMA8GA1UE\nAxMIRm9ydGluZXQwHhcNMjEwMzE4MjIwOTUyWhcNMzEwMzIzMjIwOTUyWjArMRYw\nFAYDVQQKEw1Gb3J0aW5ldCBMdGQuMREwDwYDVQQDEwhGb3J0aW5ldDCCASIwDQYJ\nKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKJuKUPF7vAFBoESVSkVSGhyvoEp0cFs\nVtsS0iOTytJAyozSfNccthQcZTtfS4LDv7y0jKwvlnjqeCULVD9j65emaJQGCLgq\n+cfxLOc7n3Un4JEEyuAjdaGPcrxMuvaZKiP/H5w9E/eY3OUuqCZJRkPjyJNFn8IW\newPMiSlA1LrKcuwqLVrjJb0vhaqiJc+X4/LFgKU5D8PrtKaCotpJ5rwjQ3G4bl4/\nL/BUrKiBnREQifRbS3bVQ05Zg2pqIm+xmOlpAV0yjR2yhwhd5KuCb53bZcaBtSLf\n3Q2bCHEWZ2XQjMbXArPnt50N+oVk5kSENskl+47Km72gHDwhP1B63JcCAwEAAaMQ\nMA4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEAf1uOXvng/SYM28Aw\n8yy1sTM2lmxnFMhYID3d2rXmsAgMZ3vVX3clmq4vYmrRNxhq6Ny8crrjllzLwRY0\nhcHqP7ZBR+hJeYDaROFMFn6jeOcbmPWeRnEI0QUHkUCM5o6xnjqoP0iQmmcjPt4w\npgPH3AnGSl309xtqb8/FwUJuU0zppYzwDt/auP2ORkXcfWeXAhrgFnRpDKyEo7nN\n/x/BtT3uPovNoJvnRngvLpFF670wK8ZGY781h/e7TIdFgpv512bJ/zk/ZAJsCKSy\nnltfV9zv7+y0TsZ4BF/GDicCSAVtll1IyDp6nrUXFroAqcYnLCATfQqy8Xkxt8wS\n9LpnyA==\n-----END CERTIFICATE-----"
  name                     = "terr-vpn-cer-ca"
  range                    = "vdom"
  source                   = "user"
  source_ip                = "0.0.0.0"
  ssl_inspection_trusted   = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_private_key` - _Private_Key.
* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `ca` - CA certificate as a PEM file.
* `ca_identifier` - CA identifier of the SCEP server.
* `last_updated` - Time at which CA was last updated.
* `name` - Name.
* `obsolete` - Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.

* `range` - Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.

* `scep_url` - URL of the SCEP server.
* `source` - CA certificate source type. Valid values: `factory`, `user`, `bundle`, `fortiguard`.

* `source_ip` - Source IP address for communications to the SCEP server.
* `trusted` - Enable/disable as a trusted CA. Valid values: `disable`, `enable`.

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
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
