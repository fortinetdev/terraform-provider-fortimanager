---
subcategory: "ObjectGtp"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_gtp_messagefilterv2"
description: |-
  Message filter for GTPv2 messages.
---

# fortimanager_object_gtp_messagefilterv2
Message filter for GTPv2 messages.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `bearer_resource_cmd_fail` - Bearer resource (command 68, failure indication 69). Valid values: `allow`, `deny`.

* `change_notification` - Change notification (req 38, resp 39). Valid values: `allow`, `deny`.

* `create_bearer` - Create bearer (req 95, resp 96). Valid values: `allow`, `deny`.

* `create_session` - Create session (req 32, resp 33). Valid values: `allow`, `deny`.

* `delete_bearer_cmd_fail` - Delete bearer (command 66, failure indication 67). Valid values: `allow`, `deny`.

* `delete_bearer_req_resp` - Delete bearer (req 99, resp 100). Valid values: `allow`, `deny`.

* `delete_pdn_connection_set` - Delete PDN connection set (req 101, resp 102). Valid values: `allow`, `deny`.

* `delete_session` - Delete session (req 36, resp 37). Valid values: `allow`, `deny`.

* `echo` - Echo (req 1, resp 2). Valid values: `allow`, `deny`.

* `modify_bearer_cmd_fail` - Modify bearer (command 64 , failure indication 65). Valid values: `allow`, `deny`.

* `modify_bearer_req_resp` - Modify bearer (req 34, resp 35). Valid values: `allow`, `deny`.

* `name` - Message filter name.
* `resume` - Resume (notify 164 , ack 165). Valid values: `allow`, `deny`.

* `suspend` - Suspend (notify 162, ack 163). Valid values: `allow`, `deny`.

* `trace_session` - Trace session (activation 71, deactivation 72). Valid values: `allow`, `deny`.

* `unknown_message` - Allow or Deny unknown messages. Valid values: `allow`, `deny`.

* `unknown_message_white_list` - White list (to allow) of unknown messages.
* `update_bearer` - Update bearer (req 97, resp 98). Valid values: `allow`, `deny`.

* `update_pdn_connection_set` - Update PDN connection set (req 200, resp 201). Valid values: `allow`, `deny`.

* `version_not_support` - Version not supported (3). Valid values: `allow`, `deny`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectGtp MessageFilterV2 can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_gtp_messagefilterv2.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
