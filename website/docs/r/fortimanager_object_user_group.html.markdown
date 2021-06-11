---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_group"
description: |-
  Configure user groups.
---

# fortimanager_object_user_group
Configure user groups.

## Example Usage

```hcl
resource "fortimanager_object_user_group" "trname" {
  name     = "terraform-tefv-group"
  password = "specify"
  user_id  = "email"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_concurrent_override` - Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `disable`, `enable`.

* `auth_concurrent_value` - Maximum number of concurrent authenticated connections per user (0 - 100).
* `authtimeout` - Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
* `company` - Set the action for the company guest user field. Valid values: `optional`, `mandatory`, `disabled`.

* `email` - Enable/disable the guest user email address field. Valid values: `disable`, `enable`.

* `expire` - Time in seconds before guest user accounts expire. (1 - 31536000 sec)
* `expire_type` - Determine when the expiration countdown begins. Valid values: `immediately`, `first-successful-login`.

* `group_type` - Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall`, `directory-service`, `fsso-service`, `guest`, `rsso`.

* `guest` - Guest. The structure of `guest` block is documented below.
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `fosid` - Id.
* `match` - Match. The structure of `match` block is documented below.
* `max_accounts` - Maximum number of guest accounts that can be created for this group (0 means unlimited).
* `member` - Names of users, peers, LDAP severs, or RADIUS servers to add to the user group.
* `mobile_phone` - Enable/disable the guest user mobile phone number field. Valid values: `disable`, `enable`.

* `multiple_guest_add` - Enable/disable addition of multiple guests. Valid values: `disable`, `enable`.

* `name` - Group name.
* `password` - Guest user password type. Valid values: `auto-generate`, `specify`, `disable`.

* `sms_custom_server` - SMS server.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard`, `custom`.

* `sponsor` - Set the action for the sponsor guest user field. Valid values: `optional`, `mandatory`, `disabled`.

* `sso_attribute_value` - Name of the RADIUS user group that this local user group represents.
* `user_id` - Guest user ID type. Valid values: `email`, `auto-generate`, `specify`.

* `user_name` - Enable/disable the guest user name entry. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `guest` block supports:

* `comment` - Comment.
* `company` - Set the action for the company guest user field.
* `email` - Email.
* `expiration` - Expire time.
* `id` - Guest ID.
* `mobile_phone` - Mobile phone.
* `name` - Guest name.
* `password` - Guest password.
* `sponsor` - Set the action for the sponsor guest user field.
* `user_id` - Guest ID.

The `match` block supports:

* `_gui_meta` - _Gui_Meta.
* `group_name` - Name of matching user or group on remote authentication server.
* `id` - ID.
* `server_name` - Name of remote auth server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Group can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
