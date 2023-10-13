---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_voip_profile_sip"
description: |-
  SIP.
---

# fortimanager_object_voip_profile_sip
SIP.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `profile` - Profile.

* `ack_rate` - ACK request rate limit (per second, per policy).
* `ack_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `block_ack` - Enable/disable block ACK requests. Valid values: `disable`, `enable`.

* `block_bye` - Enable/disable block BYE requests. Valid values: `disable`, `enable`.

* `block_cancel` - Enable/disable block CANCEL requests. Valid values: `disable`, `enable`.

* `block_geo_red_options` - Enable/disable block OPTIONS requests, but OPTIONS requests still notify for redundancy. Valid values: `disable`, `enable`.

* `block_info` - Enable/disable block INFO requests. Valid values: `disable`, `enable`.

* `block_invite` - Enable/disable block INVITE requests. Valid values: `disable`, `enable`.

* `block_long_lines` - Enable/disable block requests with headers exceeding max-line-length. Valid values: `disable`, `enable`.

* `block_message` - Enable/disable block MESSAGE requests. Valid values: `disable`, `enable`.

* `block_notify` - Enable/disable block NOTIFY requests. Valid values: `disable`, `enable`.

* `block_options` - Enable/disable block OPTIONS requests and no OPTIONS as notifying message for redundancy either. Valid values: `disable`, `enable`.

* `block_prack` - Enable/disable block prack requests. Valid values: `disable`, `enable`.

* `block_publish` - Enable/disable block PUBLISH requests. Valid values: `disable`, `enable`.

* `block_refer` - Enable/disable block REFER requests. Valid values: `disable`, `enable`.

* `block_register` - Enable/disable block REGISTER requests. Valid values: `disable`, `enable`.

* `block_subscribe` - Enable/disable block SUBSCRIBE requests. Valid values: `disable`, `enable`.

* `block_unknown` - Block unrecognized SIP requests (enabled by default). Valid values: `disable`, `enable`.

* `block_update` - Enable/disable block UPDATE requests. Valid values: `disable`, `enable`.

* `bye_rate` - BYE request rate limit (per second, per policy).
* `bye_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `call_id_regex` - Validate PCRE regular expression for Call-Id header value.
* `call_keepalive` - Continue tracking calls with no RTP for this many minutes.
* `cancel_rate` - CANCEL request rate limit (per second, per policy).
* `cancel_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `contact_fixup` - Fixup contact anyway even if contact's IP:port doesn't match session's IP:port. Valid values: `disable`, `enable`.

* `content_type_regex` - Validate PCRE regular expression for Content-Type header value.
* `hnt_restrict_source_ip` - Enable/disable restrict RTP source IP to be the same as SIP source IP when HNT is enabled. Valid values: `disable`, `enable`.

* `hosted_nat_traversal` - Hosted NAT Traversal (HNT). Valid values: `disable`, `enable`.

* `info_rate` - INFO request rate limit (per second, per policy).
* `info_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `invite_rate` - INVITE request rate limit (per second, per policy).
* `invite_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `ips_rtp` - Enable/disable allow IPS on RTP. Valid values: `disable`, `enable`.

* `log_call_summary` - Enable/disable logging of SIP call summary. Valid values: `disable`, `enable`.

* `log_violations` - Enable/disable logging of SIP violations. Valid values: `disable`, `enable`.

* `malformed_header_allow` - Action for malformed Allow header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_call_id` - Action for malformed Call-ID header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_contact` - Action for malformed Contact header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_content_length` - Action for malformed Content-Length header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_content_type` - Action for malformed Content-Type header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_cseq` - Action for malformed CSeq header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_expires` - Action for malformed Expires header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_from` - Action for malformed From header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_max_forwards` - Action for malformed Max-Forwards header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_no_proxy_require` - Action for malformed SIP messages without Proxy-Require header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_no_require` - Action for malformed SIP messages without Require header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_p_asserted_identity` - Action for malformed P-Asserted-Identity header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_rack` - Action for malformed RAck header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_record_route` - Action for malformed Record-Route header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_route` - Action for malformed Route header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_rseq` - Action for malformed RSeq header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_a` - Action for malformed SDP a line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_b` - Action for malformed SDP b line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_c` - Action for malformed SDP c line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_i` - Action for malformed SDP i line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_k` - Action for malformed SDP k line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_m` - Action for malformed SDP m line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_o` - Action for malformed SDP o line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_r` - Action for malformed SDP r line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_s` - Action for malformed SDP s line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_t` - Action for malformed SDP t line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_v` - Action for malformed SDP v line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_sdp_z` - Action for malformed SDP z line. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_to` - Action for malformed To header. Valid values: `pass`, `discard`, `respond`.

* `malformed_header_via` - Action for malformed VIA header. Valid values: `pass`, `discard`, `respond`.

* `malformed_request_line` - Action for malformed request line. Valid values: `pass`, `discard`, `respond`.

* `max_body_length` - Maximum SIP message body length (0 meaning no limit).
* `max_dialogs` - Maximum number of concurrent calls/dialogs (per policy).
* `max_idle_dialogs` - Maximum number established but idle dialogs to retain (per policy).
* `max_line_length` - Maximum SIP header line length (78-4096).
* `message_rate` - MESSAGE request rate limit (per second, per policy).
* `message_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `nat_port_range` - RTP NAT port range.
* `nat_trace` - Enable/disable preservation of original IP in SDP i line. Valid values: `disable`, `enable`.

* `no_sdp_fixup` - Enable/disable no SDP fix-up. Valid values: `disable`, `enable`.

* `notify_rate` - NOTIFY request rate limit (per second, per policy).
* `notify_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `open_contact_pinhole` - Enable/disable open pinhole for non-REGISTER Contact port. Valid values: `disable`, `enable`.

* `open_record_route_pinhole` - Enable/disable open pinhole for Record-Route port. Valid values: `disable`, `enable`.

* `open_register_pinhole` - Enable/disable open pinhole for REGISTER Contact port. Valid values: `disable`, `enable`.

* `open_via_pinhole` - Enable/disable open pinhole for Via port. Valid values: `disable`, `enable`.

* `options_rate` - OPTIONS request rate limit (per second, per policy).
* `options_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `prack_rate` - PRACK request rate limit (per second, per policy).
* `prack_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `preserve_override` - Override i line to preserve original IPS (default: append). Valid values: `disable`, `enable`.

* `provisional_invite_expiry_time` - Expiry time for provisional INVITE (10 - 3600 sec).
* `publish_rate` - PUBLISH request rate limit (per second, per policy).
* `publish_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `refer_rate` - REFER request rate limit (per second, per policy).
* `refer_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `register_contact_trace` - Enable/disable trace original IP/port within the contact header of REGISTER requests. Valid values: `disable`, `enable`.

* `register_rate` - REGISTER request rate limit (per second, per policy).
* `register_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `rfc2543_branch` - Enable/disable support via branch compliant with RFC 2543. Valid values: `disable`, `enable`.

* `rtp` - Enable/disable create pinholes for RTP traffic to traverse firewall. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.

* `ssl_auth_client` - Require a client certificate and authenticate it with the peer/peergrp.
* `ssl_auth_server` - Authenticate the server's certificate with the peer/peergrp.
* `ssl_client_certificate` - Name of Certificate to offer to server if requested.
* `ssl_client_renegotiation` - Allow/block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.

* `ssl_max_version` - Highest SSL/TLS version to negotiate. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest SSL/TLS version to negotiate. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - SSL/TLS mode for encryption & decryption of traffic. Valid values: `off`, `full`.

* `ssl_pfs` - SSL Perfect Forward Secrecy. Valid values: `require`, `deny`, `allow`.

* `ssl_send_empty_frags` - Send empty fragments to avoid attack on CBC IV (SSL 3.0 & TLS 1.0 only). Valid values: `disable`, `enable`.

* `ssl_server_certificate` - Name of Certificate return to the client in every SSL connection.
* `status` - Enable/disable SIP. Valid values: `disable`, `enable`.

* `strict_register` - Enable/disable only allow the registrar to connect. Valid values: `disable`, `enable`.

* `subscribe_rate` - SUBSCRIBE request rate limit (per second, per policy).
* `subscribe_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.

* `unknown_header` - Action for unknown SIP header. Valid values: `pass`, `discard`, `respond`.

* `update_rate` - UPDATE request rate limit (per second, per policy).
* `update_rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectVoip ProfileSip can be imported using any of these accepted formats:
```
Set import_options = ["profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_voip_profile_sip.labelname ObjectVoipProfileSip
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
