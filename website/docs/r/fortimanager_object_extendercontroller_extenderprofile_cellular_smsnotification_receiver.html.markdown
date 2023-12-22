---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver"
description: |-
  SMS notification receiver list.
---

# fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver
SMS notification receiver list.

~> This resource is a sub resource for variable `receiver` of resource `fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver" "trname" {
  extender_profile = fortimanager_object_extendercontroller_extenderprofile.trname.name
  alert            = ["system-reboot"]
  name             = "receiver-name"
  phone_number     = "+16501234567"
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

* `alert` - Alert multi-options. Valid values: `system-reboot`, `data-exhausted`, `session-disconnect`, `low-signal-strength`, `mode-switch`, `os-image-fallback`, `fgt-backup-mode-switch`.

* `name` - FortiExtender SMS notification receiver name.
* `phone_number` - Receiver phone number.  Format: [+][country code][area code][local phone number].  For example: +16501234567.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectExtenderController ExtenderProfileCellularSmsNotificationReceiver can be imported using any of these accepted formats:
```
Set import_options = ["extender_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_extendercontroller_extenderprofile_cellular_smsnotification_receiver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
