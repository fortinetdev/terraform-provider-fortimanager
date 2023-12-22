---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey_move"
description: |-
  List of multiple PSK entries.
---

# fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey_move
List of multiple PSK entries.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey_move" "trname" {
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  mpsk_group   = fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname.name
  mpsk_key     = fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey.trname2.name
  target       = fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey.trname.name
  option       = "after"
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey.trname2, fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey" "trname2" {
  name         = "terr-mpskkey2"
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  mpsk_group   = fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname.name
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup_mpskkey" "trname" {
  name         = "terr-mpskkey1"
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  mpsk_group   = fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname.name
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile_mpskgroup" "trname" {
  name         = "terr-mpskgroup"
  mpsk_profile = fortimanager_object_wirelesscontroller_mpskprofile.trname.name
  depends_on   = [fortimanager_object_wirelesscontroller_mpskprofile.trname]
}

resource "fortimanager_object_wirelesscontroller_mpskprofile" "trname" {
  name = "terr-wictl-mpsk-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `mpsk_profile` - Mpsk Profile.
* `mpsk_group` - Mpsk Group.
* `mpsk_key` - Mpsk Key.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the mpsk key changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of mpsk keys.
