---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_sshfilter_profile_shellcommands_move"
description: |-
  SSH command filter.
---

# fortimanager_object_sshfilter_profile_shellcommands_move
SSH command filter.

## Example Usage

```hcl
resource "fortimanager_object_sshfilter_profile_shellcommands_move" "trname" {
  profile        = fortimanager_object_sshfilter_profile.trname.name
  shell_commands = fortimanager_object_sshfilter_profile_shellcommands.trname.fosid
  target         = fortimanager_object_sshfilter_profile_shellcommands.trname2.fosid
  option         = "after"
  depends_on     = [fortimanager_object_sshfilter_profile_shellcommands.trname, fortimanager_object_sshfilter_profile_shellcommands.trname2]
}

resource "fortimanager_object_sshfilter_profile_shellcommands" "trname" {
  profile    = fortimanager_object_sshfilter_profile.trname.name
  fosid      = 1
  depends_on = [fortimanager_object_sshfilter_profile.trname]
}

resource "fortimanager_object_sshfilter_profile_shellcommands" "trname2" {
  profile    = fortimanager_object_sshfilter_profile.trname.name
  fosid      = 2
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
* `shell_commands` - Shell Commands.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the shell commands changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of shell commandss.
