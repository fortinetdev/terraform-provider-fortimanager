---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_sshfilter_profile"
description: |-
  SSH filter profile.
---

# fortimanager_object_sshfilter_profile
SSH filter profile.

## Example Usage

```hcl
resource "fortimanager_object_sshfilter_profile" "trname" {
  block               = ["shell"]
  default_command_log = "enable"
  log                 = ["exec", "sftp"]
  name                = "terr-ssh-filter-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `block` - SSH blocking options. Valid values: `x11`, `shell`, `exec`, `port-forward`, `tun-forward`, `sftp`, `unknown`, `scp`.

* `default_command_log` - Enable/disable logging unmatched shell commands. Valid values: `disable`, `enable`.

* `file_filter` - File-Filter. The structure of `file_filter` block is documented below.
* `log` - SSH logging options. Valid values: `x11`, `shell`, `exec`, `port-forward`, `tun-forward`, `sftp`, `unknown`, `scp`.

* `name` - SSH filter profile name.
* `shell_commands` - Shell-Commands. The structure of `shell_commands` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `file_filter` block supports:

* `entries` - Entries. The structure of `entries` block is documented below.
* `log` - Enable/disable file filter logging. Valid values: `disable`, `enable`.

* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `disable`, `enable`.

* `status` - Enable/disable file filter. Valid values: `disable`, `enable`.


The `entries` block supports:

* `action` - Action taken for matched file. Valid values: `log`, `block`.

* `comment` - Comment.
* `direction` - Match files transmitted in the session's originating or reply direction. Valid values: `any`, `incoming`, `outgoing`.

* `file_type` - Select file type.
* `filter` - Add a file filter.
* `password_protected` - Match password-protected files. Valid values: `any`, `yes`.

* `protocol` - Protocols to apply with. Valid values: `ssh`.


The `shell_commands` block supports:

* `action` - Action to take for URL filter matches. Valid values: `block`, `allow`.

* `alert` - Enable/disable alert. Valid values: `disable`, `enable`.

* `id` - Id.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `pattern` - SSH shell command pattern.
* `severity` - Log severity. Valid values: `low`, `medium`, `high`, `critical`.

* `type` - Matching type. Valid values: `regex`, `simple`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSshFilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_sshfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
