---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp_messageratelimit"
description: |-
  Message rate limiting.
---

# fortimanager_object_firewall_gtp_messageratelimit
Message rate limiting.

~> This resource is a sub resource for variable `message_rate_limit` of resource `fortimanager_object_firewall_gtp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `gtp` - Gtp.

* `create_aa_pdp_request` - Rate limit for create AA PDP context request (packets per second).
* `create_aa_pdp_response` - Rate limit for create AA PDP context response (packets per second).
* `create_mbms_request` - Rate limit for create MBMS context request (packets per second).
* `create_mbms_response` - Rate limit for create MBMS context response (packets per second).
* `create_pdp_request` - Rate limit for create PDP context request (packets per second).
* `create_pdp_response` - Rate limit for create PDP context response (packets per second).
* `delete_aa_pdp_request` - Rate limit for delete AA PDP context request (packets per second).
* `delete_aa_pdp_response` - Rate limit for delete AA PDP context response (packets per second).
* `delete_mbms_request` - Rate limit for delete MBMS context request (packets per second).
* `delete_mbms_response` - Rate limit for delete MBMS context response (packets per second).
* `delete_pdp_request` - Rate limit for delete PDP context request (packets per second).
* `delete_pdp_response` - Rate limit for delete PDP context response (packets per second).
* `echo_reponse` - Rate limit for echo response (packets per second).
* `echo_request` - Rate limit for echo requests (packets per second).
* `error_indication` - Rate limit for error indication (packets per second).
* `failure_report_request` - Rate limit for failure report request (packets per second).
* `failure_report_response` - Rate limit for failure report response (packets per second).
* `fwd_reloc_complete_ack` - Rate limit for forward relocation complete acknowledge (packets per second).
* `fwd_relocation_complete` - Rate limit for forward relocation complete (packets per second).
* `fwd_relocation_request` - Rate limit for forward relocation request (packets per second).
* `fwd_relocation_response` - Rate limit for forward relocation response (packets per second).
* `fwd_srns_context` - Rate limit for forward SRNS context (packets per second).
* `fwd_srns_context_ack` - Rate limit for forward SRNS context acknowledge (packets per second).
* `g_pdu` - Rate limit for G-PDU (packets per second).
* `identification_request` - Rate limit for identification request (packets per second).
* `identification_response` - Rate limit for identification response (packets per second).
* `mbms_de_reg_request` - Rate limit for MBMS de-registration request (packets per second).
* `mbms_de_reg_response` - Rate limit for MBMS de-registration response (packets per second).
* `mbms_notify_rej_request` - Rate limit for MBMS notification reject request (packets per second).
* `mbms_notify_rej_response` - Rate limit for MBMS notification reject response (packets per second).
* `mbms_notify_request` - Rate limit for MBMS notification request (packets per second).
* `mbms_notify_response` - Rate limit for MBMS notification response (packets per second).
* `mbms_reg_request` - Rate limit for MBMS registration request (packets per second).
* `mbms_reg_response` - Rate limit for MBMS registration response (packets per second).
* `mbms_ses_start_request` - Rate limit for MBMS session start request (packets per second).
* `mbms_ses_start_response` - Rate limit for MBMS session start response (packets per second).
* `mbms_ses_stop_request` - Rate limit for MBMS session stop request (packets per second).
* `mbms_ses_stop_response` - Rate limit for MBMS session stop response (packets per second).
* `note_ms_request` - Rate limit for note MS GPRS present request (packets per second).
* `note_ms_response` - Rate limit for note MS GPRS present response (packets per second).
* `pdu_notify_rej_request` - Rate limit for PDU notify reject request (packets per second).
* `pdu_notify_rej_response` - Rate limit for PDU notify reject response (packets per second).
* `pdu_notify_request` - Rate limit for PDU notify request (packets per second).
* `pdu_notify_response` - Rate limit for PDU notify response (packets per second).
* `ran_info` - Rate limit for RAN information relay (packets per second).
* `relocation_cancel_request` - Rate limit for relocation cancel request (packets per second).
* `relocation_cancel_response` - Rate limit for relocation cancel response (packets per second).
* `send_route_request` - Rate limit for send routing information for GPRS request (packets per second).
* `send_route_response` - Rate limit for send routing information for GPRS response (packets per second).
* `sgsn_context_ack` - Rate limit for SGSN context acknowledgement (packets per second).
* `sgsn_context_request` - Rate limit for SGSN context request (packets per second).
* `sgsn_context_response` - Rate limit for SGSN context response (packets per second).
* `support_ext_hdr_notify` - Rate limit for support extension headers notification (packets per second).
* `update_mbms_request` - Rate limit for update MBMS context request (packets per second).
* `update_mbms_response` - Rate limit for update MBMS context response (packets per second).
* `update_pdp_request` - Rate limit for update PDP context request (packets per second).
* `update_pdp_response` - Rate limit for update PDP context response (packets per second).
* `version_not_support` - Rate limit for version not supported (packets per second).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall GtpMessageRateLimit can be imported using any of these accepted formats:
```
Set import_options = ["gtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp_messageratelimit.labelname ObjectFirewallGtpMessageRateLimit
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
