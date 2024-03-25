---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_filefilter_profile"
description: |-
  Configure file-filter profiles.
---

# fortimanager_object_filefilter_profile
Configure file-filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rules`: `fortimanager_object_filefilter_profile_rules`



## Example Usage

```hcl
resource "fortimanager_object_filefilter_profile" "trname" {
  comment               = "This is a Terraform example"
  extended_log          = "disable"
  feature_set           = "flow"
  log                   = "enable"
  name                  = "terr-file-filter-profile"
  scan_archive_contents = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `extended_log` - Enable/disable file-filter extended logging. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `log` - Enable/disable file-filter logging. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `replacemsg_group` - Replacement message group
* `rules` - Rules. The structure of `rules` block is documented below.
* `scan_archive_contents` - Enable/disable archive contents scan. (Not for CIFS) Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rules` block supports:

* `action` - Action taken for matched file. Valid values: `log-only`, `block`.

* `comment` - Comment.
* `direction` - Traffic direction. (HTTP, FTP, SSH, CIFS only) Valid values: `any`, `incoming`, `outgoing`.

* `file_type` - Select file type.
* `name` - File-filter rule name.
* `password_protected` - Match password-protected files. Valid values: `any`, `yes`.

* `protocol` - Protocols to apply rule to. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `mapi`, `cifs`, `ssh`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFileFilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_filefilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
