---
subcategory: "ObjectUser"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_local"
description: |-
  Configure local users.
---

# fortimanager_object_user_local
Configure local users.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_concurrent_override` - Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `disable`, `enable`.

* `auth_concurrent_value` - Maximum number of concurrent logins permitted from the same user.
* `authtimeout` - Time in minutes before the authentication timeout for a user is reached.
* `email_to` - Two-factor recipient's email address.
* `fortitoken` - Two-factor recipient's FortiToken serial number.
* `fosid` - Id.
* `ldap_server` - Name of LDAP server with which the user must authenticate.
* `name` - User name.
* `passwd` - User's password.
* `passwd_policy` - Password policy to apply to this user, as defined in config user password-policy.
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `radius_server` - Name of RADIUS server with which the user must authenticate.
* `sms_custom_server` - Two-factor recipient's SMS server.
* `sms_phone` - Two-factor recipient's mobile phone number.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.

* `status` - Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `disable`, `enable`.

* `tacacs_server` - Name of TACACS+ server with which the user must authenticate.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken`, `email`, `sms`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`.

* `username_case_insensitivity` - Username-Case-Insensitivity. Valid values: `disable`, `enable`.

* `username_case_sensitivity` - Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable`, `enable`.

* `workstation` - Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Local can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_local.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
