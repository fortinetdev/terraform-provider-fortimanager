---
subcategory: "System Certificate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_certificate_ssh"
description: |-
  SSH certificates and keys.
---

# fortimanager_system_certificate_ssh
SSH certificates and keys.

## Argument Reference


The following arguments are supported:


* `certificate` - SSH certificate.
* `comment` - SSH certificate comment.
* `name` - Name of SSH certificate.
* `private_key` - SSH private-key


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CertificateSsh can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_certificate_ssh.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

