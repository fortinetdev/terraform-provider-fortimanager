---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_vap_vlanname_move"
description: |-
  Table for mapping VLAN name to VLAN ID.
---

# fortimanager_object_wirelesscontroller_vap_vlanname_move
Table for mapping VLAN name to VLAN ID.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_vap_vlanname_move" "trname" {
  vlan_name  = fortimanager_object_wirelesscontroller_vap_vlanname.trname2.name
  target     = fortimanager_object_wirelesscontroller_vap_vlanname.trname.name
  option     = "before"
  vap        = fortimanager_object_wirelesscontroller_vap.trname.name
  depends_on = [fortimanager_object_wirelesscontroller_vap_vlanname.trname2, fortimanager_object_wirelesscontroller_vap_vlanname.trname]
}

resource "fortimanager_object_wirelesscontroller_vap_vlanname" "trname2" {
  name       = "terr-vlanname2"
  vap        = fortimanager_object_wirelesscontroller_vap.trname.name
  depends_on = [fortimanager_object_wirelesscontroller_vap.trname]
}

resource "fortimanager_object_wirelesscontroller_vap_vlanname" "trname" {
  name       = "terr-vlanname1"
  vap        = fortimanager_object_wirelesscontroller_vap.trname.name
  depends_on = [fortimanager_object_wirelesscontroller_vap.trname]
}

resource "fortimanager_object_wirelesscontroller_vap" "trname" {
  name = "terr-wictl-vap5"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vap` - Vap.
* `vlan_name` - Vlan Name.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the vlan name changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of vlan names.
