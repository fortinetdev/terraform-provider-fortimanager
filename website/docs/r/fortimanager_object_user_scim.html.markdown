---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_scim"
description: |-
  Configure SCIM client entries.
---

# fortimanager_object_user_scim
Configure SCIM client entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_method` - TLS client authentication methods (default = bearer token). Valid values: `token`, `base`.

* `base_url` - Server URL to receive SCIM create, read, update, delete (CRUD) requests.
* `cascade` - Enable/disable to follow SCIM users/groups changes in IDP. Valid values: `disable`, `enable`.

* `certificate` - Certificate for client verification during TLS handshake.
* `client_authentication_method` - Client-Authentication-Method. Valid values: `token`, `base`.

* `client_identity_check` - Enable/disable client identity check. Valid values: `disable`, `enable`.

* `client_secret_token` - Client-Secret-Token.
* `fosid` - SCIM client ID.
* `name` - SCIM client name.
* `secret` - Secret for token verification or base authentication.
* `status` - Enable/disable System for Cross-domain Identity Management (SCIM). Valid values: `disable`, `enable`.

* `token_certificate` - Certificate for token verification.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Scim can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_scim.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
