---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_oftp"
description: |-
  OFTP certificates and keys.
---

# fortimanager_system_certificate_oftp
OFTP certificates and keys.

## Argument Reference


The following arguments are supported:


* `certificate` - PEM format certificate.
* `comment` - OFTP certificate comment.
* `local` - Choose from a local certificates.
* `mode` - Mode of certificates used by oftpd. default - Default mode. custom - Use custom certificate. local - Use a local certificate. Valid values: `default`, `custom`, `local`.

* `password` - Password for encrypted 'private-key', unset for non-encrypted.
* `private_key` - PEM format private key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System CertificateOftp can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_oftp.labelname SystemCertificateOftp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

