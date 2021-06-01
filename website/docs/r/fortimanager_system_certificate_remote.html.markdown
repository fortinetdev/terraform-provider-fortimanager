---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_remote"
description: |-
  Remote certificate.
---

# fortimanager_system_certificate_remote
Remote certificate.

## Argument Reference


The following arguments are supported:


* `cert` - Remote certificate.
* `comment` - Remote certificate comment.
* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateRemote can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_remote.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

