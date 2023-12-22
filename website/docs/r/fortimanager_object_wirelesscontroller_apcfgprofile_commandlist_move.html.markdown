---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_apcfgprofile_commandlist_move"
description: |-
  AP local configuration command list.
---

# fortimanager_object_wirelesscontroller_apcfgprofile_commandlist_move
AP local configuration command list.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_apcfgprofile_commandlist_move" "trname" {
  apcfg_profile = fortimanager_object_wirelesscontroller_apcfgprofile.trname.name
  command_list  = 23
  target        = 24
  option        = "before"
  depends_on    = [fortimanager_object_wirelesscontroller_apcfgprofile_commandlist.trname2, fortimanager_object_wirelesscontroller_apcfgprofile_commandlist.trname]
}

resource "fortimanager_object_wirelesscontroller_apcfgprofile_commandlist" "trname2" {
  apcfg_profile = fortimanager_object_wirelesscontroller_apcfgprofile.trname.name
  fosid         = 24
  depends_on    = [fortimanager_object_wirelesscontroller_apcfgprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_apcfgprofile_commandlist" "trname" {
  apcfg_profile = fortimanager_object_wirelesscontroller_apcfgprofile.trname.name
  fosid         = 23
  depends_on    = [fortimanager_object_wirelesscontroller_apcfgprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_apcfgprofile" "trname" {
  name = "terr-apcfgprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `apcfg_profile` - Apcfg Profile.
* `command_list` - Command List.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the command list changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of command lists.
