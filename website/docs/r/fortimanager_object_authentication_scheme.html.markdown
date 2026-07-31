---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_authentication_scheme"
description: |-
  Configure Authentication Schemes.
---

# fortimanager_object_authentication_scheme
Configure Authentication Schemes.

## Example Usage

```hcl
resource "fortimanager_object_authentication_scheme" "trname" {
  fsso_guest  = "enable"
  method      = ["basic"]
  name        = "terr-scheme"
  require_tfa = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cert_http_header` - Enable/disable authentication with user certificate in Client-Cert HTTP header (default = disable). Valid values: `disable`, `enable`.

* `digest_algo` - Digest Authentication Algorithms. Valid values: `md5`, `sha-256`.

* `digest_rfc2069` - Enable/disable support for the deprecated RFC2069 Digest Client (no cnonce field, default = disable). Valid values: `disable`, `enable`.

* `domain_controller` - Domain controller setting.
* `external_idp` - External identity provider configuration.
* `ems_device_owner` - Enable/disable SSH public-key authentication with device owner (default = disable). Valid values: `disable`, `enable`.

* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `fsso_guest` - Enable/disable user fsso-guest authentication (default = disable). Valid values: `disable`, `enable`.

* `group_attr_type` - Group attribute type used to match SCIM groups (default = display-name). Valid values: `display-name`, `external-id`.

* `kerberos_keytab` - Kerberos keytab setting.
* `method` - Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.

* `name` - Authentication scheme name.
* `negotiate_ntlm` - Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `disable`, `enable`.

* `require_tfa` - Enable/disable two-factor authentication (default = disable). Valid values: `disable`, `enable`.

* `saml_server` - SAML configuration.
* `saml_timeout` - SAML authentication timeout in seconds.
* `ssh_ca` - SSH CA name.
* `user_cert` - Enable/disable authentication with user certificate (default = disable). Valid values: `disable`, `enable`.

* `user_database` - Authentication server to contain user information; "local" (default) or "123" (for LDAP).
* `auth_user_header` - Auth-User-Header.
* `bearer_format` - Bearer-Format. Valid values: `standard`, `raw`.

* `bearer_header` - Bearer-Header.
* `bearer_type` - Bearer-Type. Valid values: `access-token`.

* `captcha` - Captcha. Valid values: `disable`, `enable`.

* `captcha_secret_key` - Captcha-Secret-Key.
* `captcha_site_key` - Captcha-Site-Key.
* `captcha_vendor` - Captcha-Vendor. Valid values: `google-recaptcha-v2-checkbox`, `google-recaptcha-v2-invisible`, `google-recaptcha-v3`, `cloudflare-turnstile`.

* `oidc_server` - Oidc-Server.
* `oidc_timeout` - Oidc-Timeout.
* `search_all_ldap_databases` - Search-All-Ldap-Databases. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectAuthentication Scheme can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_authentication_scheme.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
