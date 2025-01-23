---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_saml_dynamic_mapping"
description: |-
  SAML server entry configuration.
---

# fortimanager_object_user_saml_dynamic_mapping
SAML server entry configuration.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_user_saml`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `saml` - Saml.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `adfs_claim` - Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `disable`, `enable`.

* `auth_url` - URL to verify authentication.
* `cert` - Certificate to sign SAML messages.
* `clock_tolerance` - Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
* `digest_method` - Digest method algorithm (default = sha1). Valid values: `sha1`, `sha256`.

* `entity_id` - SP entity ID.
* `group_claim_type` - Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.

* `group_name` - Group name in assertion statement.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `limit_relaystate` - Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `disable`, `enable`.

* `reauth` - Enable/disable signalling of IDP to force user re-authentication (default = disable). Valid values: `disable`, `enable`.

* `scim_client` - SCIM client name.
* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `user_claim_type` - User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.

* `user_name` - User name in assertion statement.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectUser SamlDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["saml=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_saml_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
