---
subcategory: "System Saml"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_saml"
description: |-
  Global settings for SAML authentication.
---

# fortimanager_system_saml
Global settings for SAML authentication.

## Example Usage

```hcl
resource "fortimanager_system_saml" "trname" {
  default_profile     = "Restricted_User"
  login_auto_redirect = "disable"
  role                = "SP"
  status              = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `acs_url` - SP ACS(login) URL.
* `auth_request_signed` - Enable/Disable auth request signed. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `cert` - Certificate name.
* `default_profile` - Default Profile Name.
* `entity_id` - SP entity ID.
* `fabric_idp` - Fabric-Idp. The structure of `fabric_idp` block is documented below.
* `forticloud_sso` - Enable/disable FortiCloud SSO (default = disable). disable - Disable Forticloud SSO. enable - Enabld Forticloud SSO. Valid values: `disable`, `enable`.

* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `login_auto_redirect` - Enable/Disable auto redirect to IDP login page. disable - Disable auto redirect to IDP Login Page. enable - Enable auto redirect to IDP Login Page. Valid values: `disable`, `enable`.

* `role` - SAML role. IDP - IDentiy Provider. SP - Service Provider. FAB-SP - Fabric Service Provider. Valid values: `IDP`, `SP`, `FAB-SP`.

* `server_address` - server address.
* `service_providers` - Service-Providers. The structure of `service_providers` block is documented below.
* `sls_url` - SP SLS(logout) URL.
* `status` - Enable/disable SAML authentication (default = disable). disable - Disable SAML authentication. enable - Enabld SAML authentication. Valid values: `disable`, `enable`.

* `user_auto_create` - Enable/disable user auto creation (default = disable). disable - Disable auto create user. enable - Enable auto create user. Valid values: `disable`, `enable`.

* `want_assertions_signed` - Enable/Disable want assertions signed. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fabric_idp` block supports:

* `dev_id` - IDP Device ID.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_status` - Enable/disable SAML authentication (default = disable). disable - Disable SAML authentication. enable - Enabld SAML authentication. Valid values: `disable`, `enable`.


The `service_providers` block supports:

* `idp_entity_id` - IDP Entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `name` - Name.
* `prefix` - Prefix.
* `sp_adom` - SP adom name.
* `sp_cert` - SP certificate name.
* `sp_entity_id` - SP Entity ID.
* `sp_profile` - SP profile name.
* `sp_single_logout_url` - SP single logout URL.
* `sp_single_sign_on_url` - SP single sign-on URL.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Saml can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_saml.labelname SystemSaml
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

