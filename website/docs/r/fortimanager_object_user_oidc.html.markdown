---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_oidc"
description: |-
  ObjectUser Oidc
---

# fortimanager_object_user_oidc
ObjectUser Oidc

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_method` - Auth-Method. Valid values: `client_secret_basic`, `client_secret_post`, `private_key_jwt`.

* `auth_type` - Auth-Type. Valid values: `client-secret`, `private-key`.

* `authorization_url` - Authorization-Url.
* `client_id` - Client-Id.
* `client_secret` - Client-Secret.
* `clock_tolerance` - Clock-Tolerance.
* `discovery_url` - Discovery-Url.
* `display_name` - Display-Name.
* `domain_hint` - Domain-Hint.
* `group_attr_name` - Group-Attr-Name.
* `icon_url` - Icon-Url.
* `issuer` - Issuer.
* `jwks_uri` - Jwks-Uri.
* `ldap_server` - Ldap-Server.
* `name` - Name.
* `private_key` - Private-Key.
* `token_url` - Token-Url.
* `type` - Type. Valid values: `discovery`, `manual`.

* `user_attr_name` - User-Attr-Name. Valid values: `email`, `sub`, `preferred_username`.

* `user_regex` - User-Regex.
* `verify_cert` - Verify-Cert. Valid values: `disable`, `enable`.

* `verify_issuer` - Verify-Issuer. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Oidc can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_oidc.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
