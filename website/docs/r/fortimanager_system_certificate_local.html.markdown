---
subcategory: "System Certificate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_local"
description: |-
  Local keys and certificates.
---

# fortimanager_system_certificate_local
Local keys and certificates.

## Argument Reference


The following arguments are supported:


* `certificate` - Certificate.
* `comment` - Local certificate comment.
* `csr` - Certificate Signing Request.
* `name` - Name of local certificate.
* `password` - Local certificate password.
* `private_key` - Local certificate private-key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateLocal can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_local.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

