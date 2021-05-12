---
subcategory: "ObjectVpn"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_certificate_remote"
description: |-
  Remote certificate as a PEM file.
---

# fortimanager_object_vpn_certificate_remote
Remote certificate as a PEM file.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Name.
* `range` - Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.

* `remote` - Remote certificate.
* `source` - Remote certificate source type. Valid values: `factory`, `user`, `bundle`, `fortiguard`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn CertificateRemote can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_certificate_remote.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
