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

* `domain_controller` - Domain controller setting.
* `ems_device_owner` - Enable/disable SSH public-key authentication with device owner (default = disable). Valid values: `disable`, `enable`.

* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `fsso_guest` - Enable/disable user fsso-guest authentication (default = disable). Valid values: `disable`, `enable`.

* `kerberos_keytab` - Kerberos keytab setting.
* `method` - Authentication methods (default = basic). Valid values: `ntlm`, `basic`, `digest`, `form`, `negotiate`, `fsso`, `rsso`, `ssh-publickey`.

* `name` - Authentication scheme name.
* `negotiate_ntlm` - Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `disable`, `enable`.

* `require_tfa` - Enable/disable two-factor authentication (default = disable). Valid values: `disable`, `enable`.

* `saml_server` - SAML configuration.
* `saml_timeout` - SAML authentication timeout in seconds.
* `ssh_ca` - SSH CA name.
* `user_database` - Authentication server to contain user information; "local" (default) or "123" (for LDAP).


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
