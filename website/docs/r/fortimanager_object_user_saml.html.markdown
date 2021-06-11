---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_saml"
description: |-
  SAML server entry configuration.
---

# fortimanager_object_user_saml
SAML server entry configuration.

## Example Usage

```hcl
resource "fortimanager_object_user_saml" "trname" {
  name      = "terr-user-saml"
  user_name = "admin"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cert` - Certificate to sign SAML messages.
* `entity_id` - SP entity ID.
* `group_name` - Group name in assertion statement.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `name` - SAML server entry name.
* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `user_name` - User name in assertion statement.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Saml can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_saml.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
