---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_filefilter_profile_rules_move"
description: |-
  File filter rules.
---

# fortimanager_object_filefilter_profile_rules_move
File filter rules.

## Example Usage

```hcl
resource "fortimanager_object_filefilter_profile_rules_move" "trname" {
  profile    = fortimanager_object_filefilter_profile.trname.name
  rules      = fortimanager_object_filefilter_profile_rules.trname2.name
  target     = fortimanager_object_filefilter_profile_rules.trname.name
  option     = "after"
  depends_on = [fortimanager_object_filefilter_profile_rules.trname, fortimanager_object_filefilter_profile_rules.trname2]
}

resource "fortimanager_object_filefilter_profile_rules" "trname" {
  profile    = "terr-filefilter-profile"
  name       = "terr-rules1"
  file_type  = ["txt"]
  depends_on = [fortimanager_object_filefilter_profile.trname]
}

resource "fortimanager_object_filefilter_profile_rules" "trname2" {
  profile    = "terr-filefilter-profile"
  name       = "terr-rules2"
  file_type  = ["txt"]
  depends_on = [fortimanager_object_filefilter_profile.trname]
}

resource "fortimanager_object_filefilter_profile" "trname" {
  name = "terr-filefilter-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.
* `rules` - Rules.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the rules changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of rule.
