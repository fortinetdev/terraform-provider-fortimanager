---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_filefilter_profile_rules"
description: |-
  File filter rules.
---

# fortimanager_object_filefilter_profile_rules
File filter rules.

~> This resource is a sub resource for variable `rules` of resource `fortimanager_object_filefilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

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

ObjectFileFilter ProfileRules can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_filefilter_profile_rules.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
