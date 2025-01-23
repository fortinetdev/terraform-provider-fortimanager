---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_pblock_firewall_policy"
description: |-
  Configuring policy for a policy block.
---

# fortimanager_packages_pblock_firewall_policy
Configuring policy for a policy block.

~> This resource is a sub resource for variable `firewall_policy` of resource `fortimanager_packages_pblock`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pblock` - Pblock.

* `_policy_block` - Assigned policy block.  When this attribute is set, the policy represent a policy block, and all other attributes are ignored. This attribute is not available when configuring policy inside a policy block.
* `action` - Policy action (accept/deny/ipsec). Valid values: `deny`, `accept`, `ipsec`, `ssl-vpn`, `redirect`, `isolate`.

* `anti_replay` - Enable/disable anti-replay check. Valid values: `disable`, `enable`.

* `app_category` - Application category ID list.
* `app_group` - Application group names.
* `application` - Application ID list.
* `application_list` - Name of an existing Application list.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_path` - Enable/disable authentication-based routing. Valid values: `disable`, `enable`.

* `auth_redirect_addr` - HTTP-to-HTTPS redirect address for firewall authentication.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `disable`, `enable`.

* `av_profile` - Name of an existing Antivirus profile.
* `best_route` - Best-Route. Valid values: `disable`, `enable`.

* `block_notification` - Enable/disable block notification. Valid values: `disable`, `enable`.

* `captive_portal_exempt` - Enable to exempt some users from the captive portal. Valid values: `disable`, `enable`.

* `casb_profile` - Name of an existing CASB profile.
* `capture_packet` - Enable/disable capture packets. Valid values: `disable`, `enable`.

* `cgn_eif` - Enable/Disable CGN endpoint independent filtering. Valid values: `disable`, `enable`.

* `cgn_eim` - Enable/Disable CGN endpoint independent mapping Valid values: `disable`, `enable`.

* `cgn_log_server_grp` - NP log server group name
* `cgn_resource_quota` - resource quota
* `cgn_session_quota` - session quota
* `cgn_sw_eif_ctrl` - Enable/disable software endpoint independent filtering control. Valid values: `disable`, `enable`.

* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `custom_log_fields` - Custom fields to append to log messages for this policy.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake. Valid values: `disable`, `enable`.

* `diameter_filter_profile` - Name of an existing Diameter filter profile.
* `diffserv_copy` - Enable to copy packet's DiffServ values from session's original direction to its reply direction. Valid values: `disable`, `enable`.

* `devices` - Names of devices or device groups that can be matched by the policy.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `disclaimer` - Enable/disable user authentication disclaimer. Valid values: `disable`, `enable`, `user`, `domain`, `policy`.

* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dscp_match` - Enable DSCP check. Valid values: `disable`, `enable`.

* `dscp_negate` - Enable negated DSCP match. Valid values: `disable`, `enable`.

* `dscp_value` - DSCP value.
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `disable`, `enable`.

* `dstaddr` - Destination IPv4 address and address group names.
* `dstaddr_negate` - When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstaddr6` - Destination IPv6 address name and address group names.
* `dstaddr6_negate` - When enabled dstaddr6 specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstintf` - Outgoing (egress) interface.
* `dynamic_shaping` - Enable/disable dynamic RADIUS defined traffic shaping. Valid values: `disable`, `enable`.

* `eif_check` - Enable/Disable check endpoint-independent-filtering pinhole. Valid values: `disable`, `enable`.

* `eif_learn` - Enable/Disable learning of end-point-independent filtering pinhole. Valid values: `disable`, `enable`.

* `email_collect` - Enable/disable email collection. Valid values: `disable`, `enable`.

* `emailfilter_profile` - Name of an existing email filter profile.
* `fec` - Enable/disable Forward Error Correction on traffic matching this policy on a FEC device. Valid values: `disable`, `enable`.

* `file_filter_profile` - Name of an existing file-filter profile.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all`, `check-new`.

* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `disable`, `enable`.

* `fsso` - Enable/disable Fortinet Single Sign-On. Valid values: `disable`, `enable`.

* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `fsso_groups` - Names of FSSO groups.
* `geoip_anycast` - Enable/disable recognition of anycast IP addresses using the geography IP database. Valid values: `disable`, `enable`.

* `geoip_match` - Match geography address based either on its physical location or registered location. Valid values: `physical-location`, `registered-location`.

* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `groups` - Names of user groups that can authenticate with this policy.
* `gtp_profile` - GTP profile.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `disable`, `enable`.

* `icap_profile` - Name of an existing ICAP profile.
* `identity_based_route` - Name of identity-based routing rule.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `disable`, `enable`.

* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy`, `flow`.

* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_id` - Internet Service ID.
* `internet_service_name` - Internet Service name.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_id` - Internet Service source ID.
* `internet_service_src_name` - Internet Service source name.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service6` - Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service6_custom` - Custom IPv6 Internet Service name.
* `internet_service6_custom_group` - Custom Internet Service6 group name.
* `internet_service6_group` - Internet Service group name.
* `internet_service6_name` - IPv6 Internet Service name.
* `internet_service6_negate` - When enabled internet-service6 specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service6_src_custom` - Custom IPv6 Internet Service source name.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name.
* `internet_service6_src_group` - Internet Service6 source group name.
* `internet_service6_src_name` - IPv6 Internet Service source name.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `ip_version_type` - IP version of the policy.
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `log_http_transaction` - Enable/disable HTTP transaction log. Valid values: `disable`, `enable`, `all`, `utm`.

* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated. Valid values: `disable`, `enable`.

* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `enable`, `all`, `utm`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.

* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP. Valid values: `disable`, `enable`.

* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP. Valid values: `disable`, `enable`.

* `mms_profile` - Name of an existing MMS profile.
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.

* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.

* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic. Valid values: `disable`, `enable`.

* `natip` - Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic. Valid values: `disable`, `enable`.

* `network_service_dynamic` - Dynamic Network Service name.
* `network_service_src_dynamic` - Dynamic Network Service source name.
* `np_acceleration` - Enable/disable UTM Network Processor acceleration. Valid values: `disable`, `enable`.

* `ntlm` - Enable/disable NTLM authentication. Valid values: `disable`, `enable`.

* `ntlm_enabled_browsers` - HTTP-User-Agent value of supported browsers.
* `ntlm_guest` - Enable/disable NTLM guest user access. Valid values: `disable`, `enable`.

* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `disable`, `enable`.

* `passive_wan_health_measurement` - Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled. Valid values: `disable`, `enable`.

* `pcp_inbound` - Enable/disable PCP inbound DNAT. Valid values: `disable`, `enable`.

* `pcp_outbound` - Enable/disable PCP outbound SNAT. Valid values: `disable`, `enable`.

* `pcp_poolname` - PCP pool names.
* `per_ip_shaper` - Per-IP traffic shaper.
* `permit_any_host` - Accept UDP packets from any host. Valid values: `disable`, `enable`.

* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host. Valid values: `disable`, `enable`.

* `policy_behaviour_type` - Behaviour of the policy.
* `pfcp_profile` - PFCP profile.
* `policy_expiry` - Enable/disable policy expiry. Valid values: `disable`, `enable`.

* `policy_expiry_date` - Policy expiry date (YYYY-MM-DD HH:MM:SS).
* `policy_expiry_date_utc` - Policy expiry date and time, in epoch format.
* `policy_offload` - Enable/Disable hardware session setup for CGNAT. Valid values: `disable`, `enable`.

* `poolname` - IP Pool names.
* `poolname6` - IPv6 pool names.
* `port_preserve` - Enable/disable preservation of the original source port from source NAT if it has not been used. Valid values: `disable`, `enable`.

* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `radius_ip_auth_bypass` - Enable IP authentication bypass. The bypassed IP address must be received from RADIUS server. Valid values: `disable`, `enable`.

* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server. Valid values: `disable`, `enable`.

* `redirect_url` - URL users are directed to after seeing and accepting the disclaimer or authenticating.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `reputation_direction` - Direction of the initial traffic for reputation to take effect. Valid values: `source`, `destination`.

* `reputation_direction6` - Direction of the initial traffic for IPv6 reputation to take effect. Valid values: `source`, `destination`.

* `reputation_minimum` - Minimum Reputation to take action.
* `reputation_minimum6` - IPv6 Minimum Reputation to take action.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `disable`, `enable`.

* `rtp_addr` - Address names if this is an RTP NAT policy.
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT. Valid values: `disable`, `enable`.

* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.

* `schedule` - Schedule name.
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity. Valid values: `disable`, `enable`.

* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `sgt` - Security group tags.
* `sgt_check` - Enable/disable security group tags (SGT) check. Valid values: `disable`, `enable`.

* `src_vendor_mac` - Vendor MAC source ID.
* `srcaddr` - Source IPv4 address and address group names.
* `srcaddr_negate` - When enabled srcaddr/srcaddr6 specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcaddr6` - Source IPv6 address name and address group names.
* `srcaddr6_negate` - When enabled srcaddr6 specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcintf` - Incoming (ingress) interface.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `disable`, `enable`.

* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring). Valid values: `disable`, `enable`.

* `ssl_mirror_intf` - SSL mirror interface name.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag. Valid values: `all`, `data-only`, `disable`.

* `tcp_timeout_pid` - TCP timeout profile ID
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire. Valid values: `disable`, `enable`.

* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `url_category` - URL category ID list.
* `udp_timeout_pid` - UDP timeout profile ID
* `users` - Names of individual users that can authenticate with this policy.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `virtual_patch_profile` - Name of an existing virtual-patch profile.
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `waf_profile` - Name of an existing Web application firewall profile.
* `wanopt` - Enable/disable WAN optimization. Valid values: `disable`, `enable`.

* `wanopt_detection` - WAN optimization auto-detection mode. Valid values: `active`, `passive`, `off`.

* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect server. Valid values: `default`, `transparent`, `non-transparent`.

* `wanopt_peer` - WAN optimization peer.
* `wanopt_profile` - WAN optimization profile.
* `wccp` - Enable/disable forwarding traffic matching this policy to a configured WCCP server. Valid values: `disable`, `enable`.

* `webcache` - Enable/disable web cache. Valid values: `disable`, `enable`.

* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable`, `ssl-server`, `any`, `enable`.

* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Webproxy forward server name.
* `webproxy_profile` - Webproxy profile name.
* `ztna_device_ownership` - Enable/disable zero trust device ownership. Valid values: `disable`, `enable`.

* `wsso` - Enable/disable WiFi Single Sign On (WSSO). Valid values: `disable`, `enable`.

* `ztna_ems_tag` - Source ztna-ems-tag names.
* `ztna_ems_tag_secondary` - Source ztna-ems-tag-secondary names.
* `ztna_geo_tag` - Source ztna-geo-tag names.
* `ztna_policy_redirect` - Redirect ZTNA traffic to matching Access-Proxy proxy-policy. Valid values: `disable`, `enable`.

* `ztna_status` - Enable/disable zero trust access. Valid values: `disable`, `enable`.

* `ztna_tags_match_logic` - ZTNA tag matching logic. Valid values: `or`, `and`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages PblockFirewallPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pblock=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_pblock_firewall_policy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
