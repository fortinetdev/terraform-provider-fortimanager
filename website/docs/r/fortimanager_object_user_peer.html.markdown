---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_peer"
description: |-
  Configure peer users.
---

# fortimanager_object_user_peer
Configure peer users.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ca` - Name of the CA certificate.
* `cn` - Peer certificate common name.
* `cn_type` - Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.

* `ldap_mode` - Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.

* `ldap_password` - Password for LDAP server bind.
* `ldap_server` - Name of an LDAP server defined under the user ldap command. Performs client access rights check.
* `ldap_username` - Username for LDAP server bind.
* `mandatory_ca_verify` - Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `disable`, `enable`.

* `name` - Peer name.
* `ocsp_override_server` - Online Certificate Status Protocol (OCSP) server for certificate retrieval.
* `passwd` - Peer's password used for two-factor authentication.
* `subject` - Peer certificate name constraints.
* `two_factor` - Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Peer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_peer.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
