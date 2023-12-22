---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_sshfilter_profile_shellcommands"
description: |-
  SSH command filter.
---

# fortimanager_object_sshfilter_profile_shellcommands
SSH command filter.

~> This resource is a sub resource for variable `shell_commands` of resource `fortimanager_object_sshfilter_profile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_sshfilter_profile_shellcommands" "trname" {
  profile    = fortimanager_object_sshfilter_profile.trname.name
  fosid      = 1
  log        = "enable"
  alert      = "enable"
  depends_on = [fortimanager_object_sshfilter_profile.trname]
}

resource "fortimanager_object_sshfilter_profile" "trname" {
  name = "terr-sshfilter-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `action` - Action to take for URL filter matches. Valid values: `block`, `allow`.

* `alert` - Enable/disable alert. Valid values: `disable`, `enable`.

* `fosid` - Id.
* `log` - Enable/disable logging. Valid values: `disable`, `enable`.

* `pattern` - SSH shell command pattern.
* `severity` - Log severity. Valid values: `low`, `medium`, `high`, `critical`.

* `type` - Matching type. Valid values: `regex`, `simple`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSshFilter ProfileShellCommands can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_sshfilter_profile_shellcommands.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
