---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_certificate_template"
description: |-
  ObjectCertificate Template
---

# fortimanager_object_certificate_template
ObjectCertificate Template

## Example Usage

```hcl
resource "fortimanager_object_certificate_template" "trname" {
  city         = "Cd"
  country      = "AM"
  curve_name   = "secp256r1"
  digest_type  = "sha1"
  email        = "dfa@efsa.com"
  id_type      = "host-ip"
  key_size     = "2048"
  key_type     = "rsa"
  name         = "ssss"
  organization = "fds"
  organization_unit = [
    "FortinetTestLab",
  ]
  scep_password = [
    "33333",
  ]
  state = "dfa"
  type  = "external"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `city` - City.
* `country` - Country.
* `curve_name` - Curve-Name. Valid values: `secp256r1`, `secp384r1`, `secp521r1`.

* `digest_type` - Digest-Type. Valid values: `sha1`, `sha256`.

* `email` - Email.
* `id_type` - Id-Type. Valid values: `host-ip`, `domain-name`, `email`.

* `key_size` - Key-Size. Valid values: `512`, `1024`, `1536`, `2048`, `4096`.

* `key_type` - Key-Type. Valid values: `rsa`, `ec`.

* `name` - Name.
* `organization` - Organization.
* `organization_unit` - Organization-Unit.
* `scep_password` - Scep-Password.
* `scep_server` - Scep-Server.
* `state` - State.
* `subject_name` - Subject-Name.
* `type` - Type. Valid values: `external`, `local`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCertificate Template can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_certificate_template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
