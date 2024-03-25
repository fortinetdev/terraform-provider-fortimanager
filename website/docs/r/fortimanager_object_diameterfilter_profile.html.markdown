---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_diameterfilter_profile"
description: |-
  Configure Diameter filter profiles.
---

# fortimanager_object_diameterfilter_profile
Configure Diameter filter profiles.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `cmd_flags_reserve_set` - Action to be taken for messages with cmd flag reserve bits set. Valid values: `block`, `reset`, `monitor`, `allow`.

* `command_code_invalid` - Action to be taken for messages with invalid command code. Valid values: `block`, `reset`, `monitor`, `allow`.

* `command_code_range` - Valid range for command codes (0-16777215).
* `comment` - Comment.
* `log_packet` - Enable/disable packet log for triggered diameter settings. Valid values: `disable`, `enable`.

* `message_length_invalid` - Action to be taken for invalid message length. Valid values: `block`, `reset`, `monitor`, `allow`.

* `missing_request_action` - Action to be taken for answers without corresponding request. Valid values: `block`, `reset`, `monitor`, `allow`.

* `monitor_all_messages` - Enable/disable logging for all User Name and Result Code AVP messages. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `protocol_version_invalid` - Action to be taken for invalid protocol version. Valid values: `block`, `reset`, `monitor`, `allow`.

* `request_error_flag_set` - Action to be taken for request messages with error flag set. Valid values: `block`, `reset`, `monitor`, `allow`.

* `track_requests_answers` - Enable/disable validation that each answer has a corresponding request. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDiameterFilter Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_diameterfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
