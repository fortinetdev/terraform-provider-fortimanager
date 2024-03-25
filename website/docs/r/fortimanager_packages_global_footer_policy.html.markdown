---
subcategory: "Global Footer/Header Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_global_footer_policy"
description: |-
  Configure IPv4 policies.
---

# fortimanager_packages_global_footer_policy
Configure IPv4 policies.

## Example Usage

```hcl
resource "fortimanager_packages_global_footer_policy" "labelname" {
  action                 = "deny"
  anti_replay            = "enable"
  application_charts     = []
  block_notification     = "disable"
  captive_portal_exempt  = "disable"
  cgn_resource_quota     = 0
  cgn_session_quota      = 0
  delay_tcp_npu_session  = "disable"
  diffserv_forward       = "disable"
  diffserv_reverse       = "disable"
  dsri                   = "disable"
  dstaddr                = ["gall"]
  dstaddr_negate         = "disable"
  dstintf                = ["any"]
  dynamic_profile_access = []
  email_collect          = "disable"
  geoip_anycast          = "disable"
  geoip_match            = "physical-location"
  internet_service       = "disable"
  internet_service_src   = "disable"
  logtraffic             = "all"
  logtraffic_start       = "disable"
  match_vip              = "disable"
  match_vip_only         = "disable"
  name                   = "s"
  natip = [
    "0.0.0.0",
    "0.0.0.0",
  ]
  np_acceleration         = "enable"
  permit_any_host         = "disable"
  pkg                     = "default"
  policyid                = 1074741825
  profile_type            = "single"
  radius_mac_auth_bypass  = "disable"
  reputation_minimum      = 0
  rtp_nat                 = "disable"
  schedule                = "galways"
  schedule_timeout        = "disable"
  send_deny_packet        = "disable"
  service                 = ["gALL"]
  service_negate          = "disable"
  session_ttl             = "0"
  srcaddr                 = ["gall"]
  srcaddr_negate          = "disable"
  srcintf                 = ["any"]
  status                  = "enable"
  tcp_mss_receiver        = 0
  tcp_mss_sender          = 0
  tcp_session_without_syn = "disable"
  tos                     = "0x00"
  tos_mask                = "0x00"
  tos_negate              = "disable"
  vlan_cos_fwd            = 255
  vlan_cos_rev            = 255
  wccp                    = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `_policy_block` - Assigned policy block.  When this attribute is set, the policy represent a policy block, and all other attributes are ignored. This attribute is not available when configuring policy inside a policy block.
* `access_proxy` - Access-Proxy.
* `action` - Policy action (allow/deny/ipsec). Valid values: `deny`, `accept`, `ipsec`, `ssl-vpn`.

* `active_auth_method` - Active-Auth-Method. Valid values: `ntlm`, `basic`, `digest`, `form`.

* `anti_replay` - Enable/disable anti-replay check. Valid values: `disable`, `enable`.

* `app_category` - Application category ID list.
* `app_group` - Application group names.
* `application` - Application ID list.
* `application_charts` - Application-Charts. Valid values: `top10-app`, `top10-p2p-user`, `top10-media-user`.

* `application_list` - Name of an existing Application list.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_method` - Auth-Method. Valid values: `basic`, `digest`, `ntlm`, `fsae`, `form`, `fsso`, `rsso`.

* `auth_path` - Enable/disable authentication-based routing. Valid values: `disable`, `enable`.

* `auth_portal` - Auth-Portal. Valid values: `disable`, `enable`.

* `auth_redirect_addr` - HTTP-to-HTTPS redirect address for firewall authentication.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `disable`, `enable`.

* `av_profile` - Name of an existing Antivirus profile.
* `bandwidth` - Bandwidth. Valid values: `disable`, `enable`.

* `best_route` - Best-Route. Valid values: `disable`, `enable`.

* `block_notification` - Enable/disable block notification. Valid values: `disable`, `enable`.

* `captive_portal_exempt` - Enable to exempt some users from the captive portal. Valid values: `disable`, `enable`.

* `capture_packet` - Enable/disable capture packets. Valid values: `disable`, `enable`.

* `casb_profile` - Name of an existing CASB profile.
* `casi_profile` - Casi-Profile.
* `central_nat` - Central-Nat. Valid values: `disable`, `enable`.

* `cgn_eif` - Enable/Disable CGN endpoint independent filtering. Valid values: `disable`, `enable`.

* `cgn_eim` - Enable/Disable CGN endpoint independent mapping Valid values: `disable`, `enable`.

* `cgn_log_server_grp` - NP log server group name
* `cgn_resource_quota` - resource quota
* `cgn_session_quota` - session quota
* `cifs_profile` - Name of an existing CIFS profile.
* `client_reputation` - Client-Reputation. Valid values: `disable`, `enable`.

* `client_reputation_mode` - Client-Reputation-Mode. Valid values: `learning`, `monitoring`.

* `comments` - Comment.
* `custom_log_fields` - Custom fields to append to log messages for this policy.
* `decrypted_traffic_mirror` - Decrypted-Traffic-Mirror.
* `deep_inspection_options` - Deep-Inspection-Options.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake. Valid values: `disable`, `enable`.

* `delay_tcp_npu_sessoin` - Delay-Tcp-Npu-Sessoin. Valid values: `disable`, `enable`.

* `detect_https_in_http_request` - Detect-Https-In-Http-Request. Valid values: `disable`, `enable`.

* `device_detection_portal` - Device-Detection-Portal. Valid values: `disable`, `enable`.

* `device_ownership` - Device-Ownership. Valid values: `disable`, `enable`.

* `devices` - Devices.
* `diameter_filter_profile` - Name of an existing Diameter filter profile.
* `diffserv_copy` - Enable to copy packet's DiffServ values from session's original direction to its reply direction. Valid values: `disable`, `enable`.

* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `disclaimer` - Enable/disable user authentication disclaimer. Valid values: `disable`, `enable`.

* `dlp_profile` - Dlp-Profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dponly` - Dponly. Valid values: `disable`, `enable`.

* `dscp_match` - Dscp-Match. Valid values: `disable`, `enable`.

* `dscp_negate` - Dscp-Negate. Valid values: `disable`, `enable`.

* `dscp_value` - Dscp-Value.
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `disable`, `enable`.

* `dstaddr` - Destination address and address group names.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstaddr6` - Dstaddr6.
* `dstaddr6_negate` - When enabled dstaddr6 specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstintf` - Outgoing (egress) interface.
* `dynamic_bypass` - Dynamic-Bypass. Valid values: `disable`, `enable`.

* `dynamic_profile` - Dynamic-Profile. Valid values: `disable`, `enable`.

* `dynamic_profile_access` - Dynamic-Profile-Access. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `im`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`, `ssh`.

* `dynamic_profile_fallthrough` - Dynamic-Profile-Fallthrough. Valid values: `disable`, `enable`.

* `dynamic_profile_group` - Dynamic-Profile-Group.
* `dynamic_shaping` - Dynamic-Shaping. Valid values: `disable`, `enable`.

* `email_collect` - Enable/disable email collection. Valid values: `disable`, `enable`.

* `email_collection_portal` - Email-Collection-Portal. Valid values: `disable`, `enable`.

* `emailfilter_profile` - Name of an existing email filter profile.
* `endpoint_check` - Endpoint-Check. Valid values: `disable`, `enable`.

* `endpoint_compliance` - Endpoint-Compliance. Valid values: `disable`, `enable`.

* `endpoint_keepalive_interface` - Endpoint-Keepalive-Interface.
* `endpoint_profile` - Endpoint-Profile.
* `extended_log` - Extended-Log. Valid values: `disable`, `enable`.

* `failed_connection` - Failed-Connection. Valid values: `disable`, `enable`.

* `fall_through_unauthenticated` - Fall-Through-Unauthenticated. Valid values: `disable`, `enable`.

* `fec` - Enable/disable Forward Error Correction on traffic matching this policy on a FEC device. Valid values: `disable`, `enable`.

* `file_filter_profile` - File-Filter-Profile.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all`, `check-new`.

* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `disable`, `enable`.

* `force_proxy` - Force-Proxy. Valid values: `disable`, `enable`.

* `forticlient_compliance_devices` - Forticlient-Compliance-Devices. Valid values: `windows-pc`, `mac`, `iphone-ipad`, `android`.

* `forticlient_compliance_enforcement_portal` - Forticlient-Compliance-Enforcement-Portal. Valid values: `disable`, `enable`.

* `fsae` - Fsae. Valid values: `disable`, `enable`.

* `fsae_server_for_ntlm` - Fsae-Server-For-Ntlm.
* `fsso` - Enable/disable Fortinet Single Sign-On. Valid values: `disable`, `enable`.

* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `fsso_groups` - Names of FSSO groups.
* `geo_location` - Geo-Location. Valid values: `disable`, `enable`.

* `geoip_anycast` - Enable/disable recognition of anycast IP addresses using the geography IP database. Valid values: `disable`, `enable`.

* `geoip_match` - Geoip-Match. Valid values: `physical-location`, `registered-location`.

* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `groups` - Names of user groups that can authenticate with this policy.
* `gtp_profile` - GTP profile.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `disable`, `enable`.

* `http_tunnel_auth` - Http-Tunnel-Auth. Valid values: `disable`, `enable`.

* `ia_profile` - Ia-Profile.
* `icap_profile` - Name of an existing ICAP profile.
* `identity_based` - Identity-Based. Valid values: `disable`, `enable`.

* `identity_based_route` - Name of identity-based routing rule.
* `identity_from` - Identity-From. Valid values: `auth`, `device`.

* `implicit_proxy_detection` - Implicit-Proxy-Detection. Valid values: `disable`, `enable`.

* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `disable`, `enable`.

* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy`, `flow`.

* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_id` - Internet Service ID.
* `internet_service_name` - Internet-Service-Name.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_id` - Internet Service source ID.
* `internet_service_src_name` - Internet-Service-Src-Name.
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

* `ip_based` - Ip-Based. Valid values: `disable`, `enable`.

* `ip_version_type` - IP version of the policy.
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `isolator_profile` - Isolator-Profile.
* `isolator_server` - Isolator-Server.
* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `learning_mode` - Learning-Mode. Valid values: `disable`, `enable`.

* `log_http_transaction` - Log-Http-Transaction. Valid values: `disable`, `enable`.

* `log_unmatched_traffic` - Log-Unmatched-Traffic. Valid values: `disable`, `enable`.

* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `enable`, `all`, `utm`.

* `logtraffic_app` - Logtraffic-App. Valid values: `disable`, `enable`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.

* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP. Valid values: `disable`, `enable`.

* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP. Valid values: `disable`, `enable`.

* `max_session_per_user` - Max-Session-Per-User.
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

* `pass_through` - Pass-Through. Valid values: `disable`, `enable`.

* `passive_wan_health_measurement` - Passive-Wan-Health-Measurement. Valid values: `disable`, `enable`.

* `pcp_inbound` - Enable/disable PCP inbound DNAT. Valid values: `disable`, `enable`.

* `pcp_outbound` - Enable/disable PCP outbound SNAT. Valid values: `disable`, `enable`.

* `pcp_poolname` - PCP pool names.
* `per_ip_shaper` - Per-IP traffic shaper.
* `permit_any_host` - Accept UDP packets from any host. Valid values: `disable`, `enable`.

* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host. Valid values: `disable`, `enable`.

* `pfcp_profile` - PFCP profile.
* `policy_behaviour_type` - Behaviour of the policy.
* `policy_expiry` - Policy-Expiry. Valid values: `disable`, `enable`.

* `policy_expiry_date` - Policy-Expiry-Date.
* `policy_expiry_date_utc` - Policy expiry date and time, in epoch format.
* `policy_offload` - Enable/Disable hardware session setup for CGNAT. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `poolname` - IP Pool names.
* `poolname6` - Poolname6.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server. Valid values: `disable`, `enable`.

* `redirect_profile` - Redirect-Profile.
* `redirect_url` - URL users are directed to after seeing and accepting the disclaimer or authenticating.
* `replacemsg_group` - Replacemsg-Group.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `reputation_direction` - Direction of the initial traffic for reputation to take effect. Valid values: `source`, `destination`.

* `reputation_direction6` - Direction of the initial traffic for IPv6 reputation to take effect. Valid values: `source`, `destination`.

* `reputation_minimum` - Minimum Reputation to take action.
* `reputation_minimum6` - IPv6 Minimum Reputation to take action.
* `require_tfa` - Require-Tfa. Valid values: `disable`, `enable`.

* `reverse_cache` - Reverse-Cache. Valid values: `disable`, `enable`.

* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `disable`, `enable`.

* `rtp_addr` - Address names if this is an RTP NAT policy.
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT. Valid values: `disable`, `enable`.

* `scan_botnet_connections` - Scan-Botnet-Connections. Valid values: `disable`, `block`, `monitor`.

* `schedule` - Schedule name.
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity. Valid values: `disable`, `enable`.

* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `sessions` - Sessions. Valid values: `disable`, `enable`.

* `sgt` - Security group tags.
* `sgt_check` - Enable/disable security group tags (SGT) check. Valid values: `disable`, `enable`.

* `spamfilter_profile` - Spamfilter-Profile.
* `src_vendor_mac` - Src-Vendor-Mac.
* `srcaddr` - Source address and address group names.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcaddr6` - Srcaddr6.
* `srcaddr6_negate` - When enabled srcaddr6 specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcintf` - Incoming (ingress) interface.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_check` - Ssh-Policy-Check. Valid values: `disable`, `enable`.

* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `disable`, `enable`.

* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring). Valid values: `disable`, `enable`.

* `ssl_mirror_intf` - SSL mirror interface name.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `sslvpn_auth` - Sslvpn-Auth. Valid values: `any`, `local`, `radius`, `ldap`, `tacacs+`.

* `sslvpn_ccert` - Sslvpn-Ccert. Valid values: `disable`, `enable`.

* `sslvpn_cipher` - Sslvpn-Cipher. Valid values: `any`, `high`, `medium`.

* `sso_auth_method` - Sso-Auth-Method. Valid values: `fsso`, `rsso`.

* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `tags` - Tags.
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_reset` - Tcp-Reset. Valid values: `disable`, `enable`.

* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag. Valid values: `all`, `data-only`, `disable`.

* `tcp_timeout_pid` - TCP timeout profile ID
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire. Valid values: `disable`, `enable`.

* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `transaction_based` - Transaction-Based. Valid values: `disable`, `enable`.

* `transparent` - Transparent. Valid values: `disable`, `enable`.

* `type` - Type. Valid values: `explicit-web`, `transparent`, `explicit-ftp`, `ssh-tunnel`, `ssh`, `wanopt`, `access-proxy`.

* `udp_timeout_pid` - UDP timeout profile ID
* `url_category` - URL category ID list.
* `users` - Names of individual users that can authenticate with this policy.
* `utm_inspection_mode` - Utm-Inspection-Mode. Valid values: `proxy`, `flow`.

* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `uuid_idx` - Uuid-Idx.
* `vendor_mac` - Vendor-Mac.
* `videofilter_profile` - Videofilter-Profile.
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

* `web_auth_cookie` - Web-Auth-Cookie. Valid values: `disable`, `enable`.

* `webcache` - Enable/disable web cache. Valid values: `disable`, `enable`.

* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable`, `ssl-server`, `any`, `enable`.

* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Webproxy forward server name.
* `webproxy_profile` - Webproxy profile name.
* `wsso` - Enable/disable WiFi Single Sign On (WSSO). Valid values: `disable`, `enable`.

* `ztna_device_ownership` - Enable/disable zero trust device ownership. Valid values: `disable`, `enable`.

* `ztna_ems_tag` - Ztna-Ems-Tag.
* `ztna_ems_tag_secondary` - Source ztna-ems-tag-secondary names.
* `ztna_geo_tag` - Ztna-Geo-Tag.
* `ztna_policy_redirect` - Redirect ZTNA traffic to matching Access-Proxy proxy-policy. Valid values: `disable`, `enable`.

* `ztna_status` - Ztna-Status. Valid values: `disable`, `enable`.

* `ztna_tags_match_logic` - Ztna-Tags-Match-Logic. Valid values: `or`, `and`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages GlobalFooterPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_global_footer_policy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

