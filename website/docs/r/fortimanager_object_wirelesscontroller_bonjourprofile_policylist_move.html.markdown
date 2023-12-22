---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_bonjourprofile_policylist_move"
description: |-
  Bonjour policy list.
---

# fortimanager_object_wirelesscontroller_bonjourprofile_policylist_move
Bonjour policy list.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_bonjourprofile_policylist_move" "trname" {
  bonjour_profile = fortimanager_object_wirelesscontroller_bonjourprofile.trname2.name
  policy_list     = 2
  target          = 3
  option          = "after"
  depends_on      = [fortimanager_object_wirelesscontroller_bonjourprofile_policylist.trname2, fortimanager_object_wirelesscontroller_bonjourprofile_policylist.trname]
}

resource "fortimanager_object_wirelesscontroller_bonjourprofile_policylist" "trname2" {
  bonjour_profile = fortimanager_object_wirelesscontroller_bonjourprofile.trname2.name
  policy_id       = 2
  depends_on      = [fortimanager_object_wirelesscontroller_bonjourprofile.trname2]
}

resource "fortimanager_object_wirelesscontroller_bonjourprofile_policylist" "trname" {
  bonjour_profile = fortimanager_object_wirelesscontroller_bonjourprofile.trname2.name
  policy_id       = 3
  depends_on      = [fortimanager_object_wirelesscontroller_bonjourprofile.trname2]
}

resource "fortimanager_object_wirelesscontroller_bonjourprofile" "trname2" {
  name = "teset2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `bonjour_profile` - Bonjour Profile.
* `policy_list` - Policy List.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policy_id}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the policy list changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of policy lists.
