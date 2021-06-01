---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_ca"
description: |-
  CA certificate.
---

# fortimanager_system_certificate_ca
CA certificate.

## Argument Reference


The following arguments are supported:


* `ca` - CA certificate.
* `comment` - CA certificate comment.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateCa can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_ca.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

