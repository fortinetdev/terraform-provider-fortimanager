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
* `expired_password_renewal` - Enable/disable renewal of a password that already is expired. Valid values: `disable`, `enable`.

* `name` - Password policy name.
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
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
