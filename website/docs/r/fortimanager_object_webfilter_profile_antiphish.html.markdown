---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_profile_antiphish"
description: |-
  AntiPhishing profile.
---

# fortimanager_object_webfilter_profile_antiphish
AntiPhishing profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `authentication` - Authentication methods. Valid values: `domain-controller`, `ldap`.

* `check_basic_auth` - Enable/disable checking of HTTP Basic Auth field for known credentials. Valid values: `disable`, `enable`.

* `check_uri` - Enable/disable checking of GET URI parameters for known credentials. Valid values: `disable`, `enable`.

* `check_username_only` - Enable/disable acting only on valid username credentials. Action will be taken for valid usernames regardless of password validity. Valid values: `disable`, `enable`.

* `custom_patterns` - Custom-Patterns. The structure of `custom_patterns` block is documented below.
* `default_action` - Action to be taken when there is no matching rule. Valid values: `log`, `block`, `exempt`.

* `domain_controller` - Domain for which to verify received credentials against.
* `inspection_entries` - Inspection-Entries. The structure of `inspection_entries` block is documented below.
* `ldap` - LDAP server for which to verify received credentials against.
* `max_body_len` - Maximum size of a POST body to check for credentials.
* `status` - Toggle AntiPhishing functionality. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_patterns` block supports:

* `category` - Category that the pattern matches. Valid values: `username`, `password`.

* `pattern` - Target pattern.
* `type` - Pattern will be treated either as a regex pattern or literal string. Valid values: `regex`, `literal`.


The `inspection_entries` block supports:

* `action` - Action to be taken upon an AntiPhishing match. Valid values: `log`, `block`, `exempt`.

* `fortiguard_category` - FortiGuard category to match.
* `name` - Inspection target name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectWebfilter ProfileAntiphish can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_profile_antiphish.labelname ObjectWebfilterProfileAntiphish
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
