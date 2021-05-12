---
subcategory: "ObjectGtp"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_gtp_messagefilterv0v1"
description: |-
  Message filter for GTPv0/v1 messages.
---

# fortimanager_object_gtp_messagefilterv0v1
Message filter for GTPv0/v1 messages.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `create_mbms` - GTPv1 create MBMS context (req 100, resp 101). Valid values: `allow`, `deny`.

* `create_pdp` - Create PDP context (req 16, resp 17). Valid values: `allow`, `deny`.

* `data_record` - Data record transfer (req 240, resp 241). Valid values: `allow`, `deny`.

* `delete_aa_pdp` - GTPv0 delete AA PDP context (req 24, resp 25). Valid values: `allow`, `deny`.

* `delete_mbms` - GTPv1 delete MBMS context (req 104, resp 105). Valid values: `allow`, `deny`.

* `delete_pdp` - Delete PDP context (req 20, resp 21). Valid values: `allow`, `deny`.

* `echo` - Echo (req 1, resp 2). Valid values: `allow`, `deny`.

* `end_marker` - GTPv1 End marker (254). Valid values: `allow`, `deny`.

* `error_indication` - Error indication (26). Valid values: `allow`, `deny`.

* `failure_report` - Failure report (req 34, resp 35). Valid values: `allow`, `deny`.

* `fwd_relocation` - GTPv1 forward relocation (req 53, resp 54, complete 55, complete ack 59). Valid values: `allow`, `deny`.

* `fwd_srns_context` - GTPv1 forward SRNS (context 58, context ack 60). Valid values: `allow`, `deny`.

* `gtp_pdu` - PDU (255). Valid values: `allow`, `deny`.

* `identification` - Identification (req 48, resp 49). Valid values: `allow`, `deny`.

* `mbms_de_registration` - GTPv1 MBMS de-registration (req 114, resp 115). Valid values: `allow`, `deny`.

* `mbms_notification` - GTPv1 MBMS notification (req 96, resp 97, reject req 98. reject resp 99). Valid values: `allow`, `deny`.

* `mbms_registration` - GTPv1 MBMS registration (req 112, resp 113). Valid values: `allow`, `deny`.

* `mbms_session_start` - GTPv1 MBMS session start (req 116, resp 117). Valid values: `allow`, `deny`.

* `mbms_session_stop` - GTPv1 MBMS session stop (req 118, resp 119). Valid values: `allow`, `deny`.

* `mbms_session_update` - GTPv1 MBMS session update (req 120, resp 121). Valid values: `allow`, `deny`.

* `ms_info_change_notif` - GTPv1 MS info change notification (req 128, resp 129). Valid values: `allow`, `deny`.

* `name` - Message filter name.
* `node_alive` - Node alive (req 4, resp 5). Valid values: `allow`, `deny`.

* `note_ms_present` - Note MS GPRS present (req 36, resp 37). Valid values: `allow`, `deny`.

* `pdu_notification` - PDU notification (req 27, resp 28, reject req 29, reject resp 30). Valid values: `allow`, `deny`.

* `ran_info` - GTPv1 RAN information relay (70). Valid values: `allow`, `deny`.

* `redirection` - Redirection (req 6, resp 7). Valid values: `allow`, `deny`.

* `relocation_cancel` - GTPv1 relocation cancel (req 56, resp 57). Valid values: `allow`, `deny`.

* `send_route` - Send routing information for GPRS (req 32, resp 33). Valid values: `allow`, `deny`.

* `sgsn_context` - SGSN context (req 50, resp 51, ack 52). Valid values: `allow`, `deny`.

* `support_extension` - GTPv1 supported extension headers notify (31). Valid values: `allow`, `deny`.

* `unknown_message` - Allow or Deny unknown messages. Valid values: `allow`, `deny`.

* `unknown_message_white_list` - White list (to allow) of unknown messages.
* `update_mbms` - GTPv1 update MBMS context (req 102, resp 103). Valid values: `allow`, `deny`.

* `update_pdp` - Update PDP context (req 18, resp 19). Valid values: `allow`, `deny`.

* `v0_create_aa_pdp__v1_init_pdp_ctx` - GTPv0 create AA PDP context (req 22, resp 23); Or GTPv1 initiate PDP context (req 22, resp 23). Valid values: `deny`, `allow`.

* `version_not_support` - Version not supported (3). Valid values: `allow`, `deny`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectGtp MessageFilterV0V1 can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_gtp_messagefilterv0v1.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
