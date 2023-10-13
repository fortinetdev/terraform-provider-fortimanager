---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_passwordpolicy"
description: |-
  Configure user password policy.
---

# fortimanager_object_user_passwordpolicy
Configure user password policy.

## Example Usage

```hcl
resource "fortimanager_object_user_passwordpolicy" "trname" {
  expire_days              = 180
  expired_password_renewal = "disable"
  name                     = "terr-user-pwdpolicy"
  warn_days                = 15
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `expire_days` - Time in days before the user's password expires.
* `expire_status` - Enable/disable password expiration. Valid values: `disable`, `enable`.

* `expired_password_renewal` - Enable/disable renewal of a password that already is expired. Valid values: `disable`, `enable`.

* `min_change_characters` - Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).
* `min_lower_case_letter` - Minimum number of lowercase characters in password (0 - 128, default = 0).
* `min_non_alphanumeric` - Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
* `min_number` - Minimum number of numeric characters in password (0 - 128, default = 0).
* `min_upper_case_letter` - Minimum number of uppercase characters in password (0 - 128, default = 0).
* `minimum_length` - Minimum password length (8 - 128, default = 8).
* `name` - Password policy name.
* `reuse_password` - Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides. Valid values: `disable`, `enable`.

* `warn_days` - Time in days before a password expiration warning message is displayed to the user upon login.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser PasswordPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_passwordpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
