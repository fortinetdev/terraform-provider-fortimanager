---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_local_dynamic_mapping"
description: |-
  Configure local users.
---

# fortimanager_object_user_local_dynamic_mapping
Configure local users.

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_user_local`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `local` - Local.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `auth_concurrent_override` - Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `disable`, `enable`.

* `auth_concurrent_value` - Maximum number of concurrent logins permitted from the same user.
* `authtimeout` - Time in minutes before the authentication timeout for a user is reached.
* `email_to` - Two-factor recipient's email address.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `fortitoken` - Two-factor recipient's FortiToken serial number.
* `history0` - History0.
* `history1` - History1.
* `history10` - History10.
* `history11` - History11.
* `history12` - History12.
* `history13` - History13.
* `history14` - History14.
* `history15` - History15.
* `history16` - History16.
* `history17` - History17.
* `history18` - History18.
* `history19` - History19.
* `history2` - History2.
* `history3` - History3.
* `history4` - History4.
* `history5` - History5.
* `history6` - History6.
* `history7` - History7.
* `history8` - History8.
* `history9` - History9.
* `fosid` - Id.
* `ldap_server` - Name of LDAP server with which the user must authenticate.
* `passwd` - User's password.
* `passwd_policy` - Password policy to apply to this user, as defined in config user password-policy.
* `passwd_time` - Time of the last password update.
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `qkd_profile` - Quantum Key Distribution (QKD) profile.
* `radius_server` - Name of RADIUS server with which the user must authenticate.
* `saml_server` - Name of SAML server with which the user must authenticate.
* `sms_custom_server` - Two-factor recipient's SMS server.
* `sms_phone` - Two-factor recipient's mobile phone number.
* `sms_provider` - Sms-Provider.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.

* `status` - Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `disable`, `enable`.

* `tacacs_server` - Name of TACACS+ server with which the user must authenticate.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken`, `email`, `sms`, `fortitoken-cloud`.

* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.

* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.

* `type` - Authentication method. Valid values: `password`, `radius`, `tacacs+`, `ldap`, `saml`.

* `username_case_insensitivity` - Username-Case-Insensitivity. Valid values: `disable`, `enable`.

* `username_case_sensitivity` - Username-Case-Sensitivity. Valid values: `disable`, `enable`.

* `username_sensitivity` - Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `workstation` - Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectUser LocalDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["local=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_local_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
