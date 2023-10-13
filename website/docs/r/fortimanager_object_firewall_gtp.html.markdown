---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp"
description: |-
  Configure GTP.
---

# fortimanager_object_firewall_gtp
Configure GTP.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `addr_notify` - overbilling notify address
* `apn` - Apn. The structure of `apn` block is documented below.
* `apn_filter` - apn filter Valid values: `disable`, `enable`.

* `authorized_ggsns` - Authorized GGSN group
* `authorized_ggsns6` - Authorized GGSN/PGW IPv6 group.
* `authorized_sgsns` - Authorized SGSN group
* `authorized_sgsns6` - Authorized SGSN/SGW IPv6 group.
* `comment` - Comment.
* `context_id` - Overbilling context.
* `control_plane_message_rate_limit` - control plane message rate limit
* `default_apn_action` - default apn action Valid values: `allow`, `deny`.

* `default_imsi_action` - default imsi action Valid values: `allow`, `deny`.

* `default_ip_action` - default action for encapsulated IP traffic Valid values: `allow`, `deny`.

* `default_noip_action` - default action for encapsulated non-IP traffic Valid values: `allow`, `deny`.

* `default_policy_action` - default advanced policy action Valid values: `allow`, `deny`.

* `denied_log` - log denied Valid values: `disable`, `enable`.

* `echo_request_interval` - echo request interval (in seconds)
* `extension_log` - log in extension format Valid values: `disable`, `enable`.

* `forwarded_log` - log forwarded Valid values: `disable`, `enable`.

* `global_tunnel_limit` - Global tunnel limit.
* `gtp_in_gtp` - gtp in gtp Valid values: `allow`, `deny`.

* `gtpu_denied_log` - Enable/disable logging of denied GTP-U packets. Valid values: `disable`, `enable`.

* `gtpu_forwarded_log` - Enable/disable logging of forwarded GTP-U packets. Valid values: `disable`, `enable`.

* `gtpu_log_freq` - Logging of frequency of GTP-U packets.
* `half_close_timeout` - Half-close tunnel timeout (in seconds).
* `half_open_timeout` - Half-open tunnel timeout (in seconds).
* `handover_group` - Handover SGSN group
* `handover_group6` - Handover SGSN/SGW IPv6 group.
* `ie_allow_list_v0v1` - IE allow list.
* `ie_allow_list_v2` - IE allow list.
* `ie_remove_policy` - Ie-Remove-Policy. The structure of `ie_remove_policy` block is documented below.
* `ie_remover` - IE removal policy. Valid values: `disable`, `enable`.

* `ie_validation` - Ie-Validation. The structure of `ie_validation` block is documented below.
* `ie_white_list_v0v1` - IE white list.
* `ie_white_list_v2` - IE white list.
* `imsi` - Imsi. The structure of `imsi` block is documented below.
* `imsi_filter` - imsi filter Valid values: `disable`, `enable`.

* `interface_notify` - overbilling interface
* `invalid_reserved_field` - Invalid reserved field in GTP header Valid values: `allow`, `deny`.

* `invalid_sgsns_to_log` - Invalid SGSN group to be logged
* `invalid_sgsns6_to_log` - Invalid SGSN IPv6 group to be logged.
* `ip_filter` - IP filter for encapsulted traffic Valid values: `disable`, `enable`.

* `ip_policy` - Ip-Policy. The structure of `ip_policy` block is documented below.
* `log_freq` - Logging of frequency of GTP-C packets.
* `log_gtpu_limit` - the user data log limit (0-512 bytes)
* `log_imsi_prefix` - IMSI prefix for selective logging.
* `log_msisdn_prefix` - the msisdn prefix for selective logging
* `max_message_length` - max message length
* `message_filter_v0v1` - Message filter.
* `message_filter_v2` - Message filter.
* `message_rate_limit` - Message-Rate-Limit. The structure of `message_rate_limit` block is documented below.
* `message_rate_limit_v0` - Message-Rate-Limit-V0. The structure of `message_rate_limit_v0` block is documented below.
* `message_rate_limit_v1` - Message-Rate-Limit-V1. The structure of `message_rate_limit_v1` block is documented below.
* `message_rate_limit_v2` - Message-Rate-Limit-V2. The structure of `message_rate_limit_v2` block is documented below.
* `min_message_length` - min message length
* `miss_must_ie` - Missing mandatory information element Valid values: `allow`, `deny`.

* `monitor_mode` - GTP monitor mode. Valid values: `disable`, `enable`, `vdom`.

* `name` - Profile name.
* `noip_filter` - non-IP filter for encapsulted traffic Valid values: `disable`, `enable`.

* `noip_policy` - Noip-Policy. The structure of `noip_policy` block is documented below.
* `out_of_state_ie` - Out of state information element. Valid values: `allow`, `deny`.

* `out_of_state_message` - Out of state GTP message Valid values: `allow`, `deny`.

* `per_apn_shaper` - Per-Apn-Shaper. The structure of `per_apn_shaper` block is documented below.
* `policy` - Policy. The structure of `policy` block is documented below.
* `policy_filter` - Advanced policy filter Valid values: `disable`, `enable`.

* `policy_v2` - Policy-V2. The structure of `policy_v2` block is documented below.
* `port_notify` - overbilling notify port
* `rat_timeout_profile` - RAT timeout profile.
* `rate_limit_mode` - GTP rate limit mode. Valid values: `per-profile`, `per-stream`, `per-apn`.

* `rate_limited_log` - log rate limited Valid values: `disable`, `enable`.

* `rate_sampling_interval` - rate sampling interval (1-3600 seconds)
* `remove_if_echo_expires` - remove if echo response expires Valid values: `disable`, `enable`.

* `remove_if_recovery_differ` - remove upon different Recovery IE Valid values: `disable`, `enable`.

* `reserved_ie` - reserved information element Valid values: `allow`, `deny`.

* `send_delete_when_timeout` - send DELETE request to path endpoints when GTPv0/v1 tunnel timeout. Valid values: `disable`, `enable`.

* `send_delete_when_timeout_v2` - send DELETE request to path endpoints when GTPv2 tunnel timeout. Valid values: `disable`, `enable`.

* `spoof_src_addr` - Spoofed source address for Mobile Station. Valid values: `allow`, `deny`.

* `state_invalid_log` - log state invalid Valid values: `disable`, `enable`.

* `sub_second_interval` - Sub-second interval (0.1, 0.25, or 0.5 sec, default = 0.5). Valid values: `0.1`, `0.25`, `0.5`.

* `sub_second_sampling` - Enable/disable sub-second sampling. Valid values: `disable`, `enable`.

* `traffic_count_log` - log tunnel traffic counter Valid values: `disable`, `enable`.

* `tunnel_limit` - tunnel limit
* `tunnel_limit_log` - tunnel limit Valid values: `disable`, `enable`.

* `tunnel_timeout` - Established tunnel timeout (in seconds).
* `unknown_version_action` - action for unknown gtp version Valid values: `allow`, `deny`.

* `user_plane_message_rate_limit` - user plane message rate limit
* `warning_threshold` - Warning threshold for rate limiting (0 - 99 percent).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `apn` block supports:

* `action` - Action. Valid values: `allow`, `deny`.

* `apnmember` - APN member.
* `id` - ID.
* `selection_mode` - APN selection mode. Valid values: `ms`, `net`, `vrf`.


The `ie_remove_policy` block supports:

* `id` - ID.
* `remove_ies` - GTP IEs to be removed. Valid values: `apn-restriction`, `rat-type`, `rai`, `uli`, `imei`.

* `sgsn_addr` - SGSN address name.
* `sgsn_addr6` - SGSN IPv6 address name.

The `ie_validation` block supports:

* `apn_restriction` - Validate APN restriction. Valid values: `disable`, `enable`.

* `charging_id` - Validate charging ID. Valid values: `disable`, `enable`.

* `charging_gateway_addr` - Validate charging gateway address. Valid values: `disable`, `enable`.

* `end_user_addr` - Validate end user address. Valid values: `disable`, `enable`.

* `gsn_addr` - Validate GSN address. Valid values: `disable`, `enable`.

* `imei` - Validate IMEI(SV). Valid values: `disable`, `enable`.

* `imsi` - Validate IMSI. Valid values: `disable`, `enable`.

* `mm_context` - Validate MM context. Valid values: `disable`, `enable`.

* `ms_tzone` - Validate MS time zone. Valid values: `disable`, `enable`.

* `ms_validated` - Validate MS validated. Valid values: `disable`, `enable`.

* `msisdn` - Validate MSISDN. Valid values: `disable`, `enable`.

* `nsapi` - Validate NSAPI. Valid values: `disable`, `enable`.

* `pdp_context` - Validate PDP context. Valid values: `disable`, `enable`.

* `qos_profile` - Validate Quality of Service(QoS) profile. Valid values: `disable`, `enable`.

* `rai` - Validate RAI. Valid values: `disable`, `enable`.

* `rat_type` - Validate RAT type. Valid values: `disable`, `enable`.

* `reordering_required` - Validate re-ordering required. Valid values: `disable`, `enable`.

* `selection_mode` - Validate selection mode. Valid values: `disable`, `enable`.

* `uli` - Validate user location information. Valid values: `disable`, `enable`.


The `imsi` block supports:

* `action` - Action. Valid values: `allow`, `deny`.

* `apnmember` - APN member.
* `id` - ID.
* `mcc_mnc` - MCC MNC.
* `msisdn_prefix` - MSISDN prefix.
* `selection_mode` - APN selection mode. Valid values: `ms`, `net`, `vrf`.


The `ip_policy` block supports:

* `action` - Action. Valid values: `allow`, `deny`.

* `dstaddr` - Destination address name.
* `dstaddr6` - Destination IPv6 address name.
* `id` - ID.
* `srcaddr` - Source address name.
* `srcaddr6` - Source IPv6 address name.

The `message_rate_limit` block supports:

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

The `message_rate_limit_v0` block supports:

* `create_pdp_request` - Rate limit (packets/s) for create PDP context request.
* `delete_pdp_request` - Rate limit (packets/s) for delete PDP context request.
* `echo_request` - Rate limit (packets/s) for echo request.

The `message_rate_limit_v1` block supports:

* `create_pdp_request` - Rate limit (packets/s) for create PDP context request.
* `delete_pdp_request` - Rate limit (packets/s) for delete PDP context request.
* `echo_request` - Rate limit (packets/s) for echo request.

The `message_rate_limit_v2` block supports:

* `create_session_request` - Rate limit (packets/s) for create session request.
* `delete_session_request` - Rate limit (packets/s) for delete session request.
* `echo_request` - Rate limit (packets/s) for echo request.

The `noip_policy` block supports:

* `action` - Action. Valid values: `allow`, `deny`.

* `end` - End of protocol range (0 - 255).
* `id` - ID.
* `start` - Start of protocol range (0 - 255).
* `type` - Protocol field type. Valid values: `etsi`, `ietf`.


The `per_apn_shaper` block supports:

* `apn` - APN name.
* `id` - ID.
* `rate_limit` - Rate limit (packets/s) for create PDP context request.
* `version` - GTP version number: 0 or 1.

The `policy` block supports:

* `action` - Action. Valid values: `allow`, `deny`.

* `apn_sel_mode` - APN selection mode. Valid values: `ms`, `net`, `vrf`.

* `apnmember` - APN member.
* `id` - ID.
* `imei` - IMEI pattern.
* `imsi` - IMSI prefix.
* `imsi_prefix` - IMSI prefix.
* `max_apn_restriction` - Maximum APN restriction value. Valid values: `all`, `public-1`, `public-2`, `private-1`, `private-2`.

* `messages` - GTP messages. Valid values: `create-req`, `create-res`, `update-req`, `update-res`.

* `msisdn` - MSISDN prefix.
* `msisdn_prefix` - MSISDN prefix.
* `rai` - RAI pattern.
* `rat_type` - RAT Type. Valid values: `any`, `utran`, `geran`, `wlan`, `gan`, `hspa`, `eutran`, `virtual`, `nbiot`.

* `uli` - ULI pattern.

The `policy_v2` block supports:

* `action` - Action. Valid values: `deny`, `allow`.

* `apn_sel_mode` - APN selection mode. Valid values: `ms`, `net`, `vrf`.

* `apnmember` - APN member.
* `id` - ID.
* `imsi_prefix` - IMSI prefix.
* `max_apn_restriction` - Maximum APN restriction value. Valid values: `all`, `public-1`, `public-2`, `private-1`, `private-2`.

* `mei` - MEI pattern.
* `messages` - GTP messages. Valid values: `create-ses-req`, `create-ses-res`, `modify-bearer-req`, `modify-bearer-res`.

* `msisdn_prefix` - MSISDN prefix.
* `rat_type` - RAT Type. Valid values: `any`, `utran`, `geran`, `wlan`, `gan`, `hspa`, `eutran`, `virtual`, `nbiot`, `ltem`, `nr`.

* `uli` - GTPv2 ULI patterns (in order of CGI SAI RAI TAI ECGI LAI).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Gtp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
