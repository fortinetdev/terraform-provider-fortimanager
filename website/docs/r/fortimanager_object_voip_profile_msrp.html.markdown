---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_voip_profile_msrp"
description: |-
  MSRP.
---

# fortimanager_object_voip_profile_msrp
MSRP.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `log_violations` - Enable/disable logging of MSRP violations. Valid values: `disable`, `enable`.

* `max_msg_size` - Maximum allowable MSRP message size (1-65535).
* `max_msg_size_action` - Action for violation of max-msg-size. Valid values: `pass`, `block`, `reset`, `monitor`.

* `status` - Enable/disable MSRP. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectVoip ProfileMsrp can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_voip_profile_msrp.labelname ObjectVoipProfileMsrp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
