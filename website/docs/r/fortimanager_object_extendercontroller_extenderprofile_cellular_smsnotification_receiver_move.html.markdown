---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver_move"
description: |-
  SMS notification receiver list.
---

# fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver_move
SMS notification receiver list.

## Example Usage

```hcl
resource "fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver_move" "trname" {
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  receiver         = "terraform"
  target           = "terraform2"
  option           = "after"
  depends_on       = [fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver.trname2, fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver" "trname2" {
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  name             = "terraform2"
  depends_on       = [fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver" "trname" {
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  name             = "terraform"
  depends_on       = [fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification" "trname" {
  status           = "disable"
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  depends_on       = [fortimanager_object_extendercontroller_extenderprofile.trname]
}

resource "fortimanager_object_extendercontroller_extenderprofile" "trname" {
  name = "terr-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `extender_profile` - Extender Profile.
* `receiver` - Receiver.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the receiver changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of receivers.
