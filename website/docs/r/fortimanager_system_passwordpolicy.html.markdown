---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_passwordpolicy"
description: |-
  Password policy.
---

# fortimanager_system_passwordpolicy
Password policy.

## Example Usage

```hcl
resource "fortimanager_system_passwordpolicy" "trname" {
  must_contain = ["terraform"]
  status       = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `change_4_characters` - Enable/disable changing at least 4 characters for new password. disable - Disable changing at least 4 characters for new password. enable - Enable changing at least 4 characters for new password. Valid values: `disable`, `enable`.

* `expire` - Number of days after which admin users' password will expire (0 - 3650, 0 = never expire).
* `minimum_length` - Minimum password length.
* `must_contain` - Password character requirements. upper-case-letter - Require password to contain upper case letter. lower-case-letter - Require password to contain lower case letter. number - Require password to contain number. non-alphanumeric - Require password to contain non-alphanumeric characters. Valid values: `upper-case-letter`, `lower-case-letter`, `number`, `non-alphanumeric`.

* `status` - Enable/disable password policy. disable - Disable password policy. enable - Enable password policy. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System PasswordPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_passwordpolicy.labelname SystemPasswordPolicy
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

