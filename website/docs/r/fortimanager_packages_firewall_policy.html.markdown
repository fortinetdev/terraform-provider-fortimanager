---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_policy"
description: |-
  Configure IPv4 policies.
---

# fortimanager_packages_firewall_policy
Configure IPv4 policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_policy" "labelname" {
  action                = "deny"
  anti_replay           = "enable"
  block_notification    = "disable"
  captive_portal_exempt = "disable"
  delay_tcp_npu_session = "disable"
  diffserv_forward      = "disable"
  diffserv_reverse      = "disable"
  dsri                  = "disable"
  dstaddr               = ["all"]
  dstaddr_negate        = "disable"
  dstintf               = ["any"]
  email_collect         = "disable"
  geoip_anycast         = "disable"
  geoip_match           = "physical-location"
  internet_service      = "disable"
  internet_service_src  = "disable"
  logtraffic            = "all"
  logtraffic_start      = "disable"
  match_vip             = "disable"
  match_vip_only        = "disable"
  name                  = "s13"
  natip = [
    "1.0.0.1",
    "255.255.255.255",
  ]
  np_acceleration         = "enable"
  ntlm_enabled_browsers   = []
  permit_any_host         = "disable"
  pkg                     = "default"
  policyid                = 146
  profile_type            = "single"
  radius_mac_auth_bypass  = "disable"
  reputation_minimum      = 0
  rtp_nat                 = "disable"
  schedule                = "always"
  schedule_timeout        = "disable"
  send_deny_packet        = "disable"
  service                 = ["ALL"]
  service_negate          = "disable"
  session_ttl             = "0"
  srcaddr                 = ["all"]
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

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Policy action (allow/deny/ipsec). Valid values: `deny`, `accept`, `ipsec`, `ssl-vpn`.

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

* `capture_packet` - Enable/disable capture packets. Valid values: `disable`, `enable`.

* `cgn_eif` - Enable/Disable CGN endpoint independent filtering. Valid values: `disable`, `enable`.

* `cgn_eim` - Enable/Disable CGN endpoint independent mapping Valid values: `disable`, `enable`.

* `cgn_log_server_grp` - NP log server group name
* `cgn_resource_quota` - resource quota
* `cgn_session_quota` - session quota
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `custom_log_fields` - Custom fields to append to log messages for this policy.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake. Valid values: `disable`, `enable`.

* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `disclaimer` - Enable/disable user authentication disclaimer. Valid values: `disable`, `enable`.

* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `disable`, `enable`.

* `dstaddr` - Destination address and address group names.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstaddr6` - Destination IPv6 address name and address group names.
* `dstintf` - Outgoing (egress) interface.
* `dynamic_shaping` - Enable/disable dynamic RADIUS defined traffic shaping. Valid values: `disable`, `enable`.

* `email_collect` - Enable/disable email collection. Valid values: `disable`, `enable`.

* `emailfilter_profile` - Name of an existing email filter profile.
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
* `internet_service_name` - Internet Service name.
* `internet_service_id` - Internet Service ID.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_name` - Internet Service source name.
* `internet_service_src_id` - Internet Service source ID.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `ippool` - Enable to use IP Pools for source NAT. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `enable`, `all`, `utm`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.

* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP. Valid values: `disable`, `enable`.

* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP. Valid values: `disable`, `enable`.

* `mms_profile` - Name of an existing MMS profile.
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic. Valid values: `disable`, `enable`.

* `natip` - Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic. Valid values: `disable`, `enable`.

* `np_acceleration` - Enable/disable UTM Network Processor acceleration. Valid values: `disable`, `enable`.

* `ntlm` - Enable/disable NTLM authentication. Valid values: `disable`, `enable`.

* `ntlm_enabled_browsers` - HTTP-User-Agent value of supported browsers.
* `ntlm_guest` - Enable/disable NTLM guest user access. Valid values: `disable`, `enable`.

* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `disable`, `enable`.

* `passive_wan_health_measurement` - Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled. Valid values: `disable`, `enable`.

* `per_ip_shaper` - Per-IP traffic shaper.
* `permit_any_host` - Accept UDP packets from any host. Valid values: `disable`, `enable`.

* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host. Valid values: `disable`, `enable`.

* `policy_offload` - Enable/Disable hardware session setup for CGNAT. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `poolname` - IP Pool names.
* `poolname6` - IPv6 pool names.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server. Valid values: `disable`, `enable`.

* `redirect_url` - URL users are directed to after seeing and accepting the disclaimer or authenticating.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `reputation_direction` - Direction of the initial traffic for reputation to take effect. Valid values: `source`, `destination`.

* `reputation_minimum` - Minimum Reputation to take action.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `disable`, `enable`.

* `rtp_addr` - Address names if this is an RTP NAT policy.
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT. Valid values: `disable`, `enable`.

* `schedule` - Schedule name.
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity. Valid values: `disable`, `enable`.

* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `src_vendor_mac` - Vendor MAC source ID.
* `srcaddr` - Source address and address group names.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcaddr6` - Source IPv6 address name and address group names.
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

* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire. Valid values: `disable`, `enable`.

* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `url_category` - URL category ID list.
* `users` - Names of individual users that can authenticate with this policy.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile.
* `vpn_dst_node` - Vpn_Dst_Node. The structure of `vpn_dst_node` block is documented below.
* `vpn_src_node` - Vpn_Src_Node. The structure of `vpn_src_node` block is documented below.
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
* `ztna_ems_tag` - Source ztna-ems-tag names.
* `ztna_geo_tag` - Source ztna-geo-tag names.
* `ztna_status` - Enable/disable zero trust access. Valid values: `disable`, `enable`.

* `wsso` - Enable/disable WiFi Single Sign On (WSSO). Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `vpn_dst_node` block supports:

* `host` - Host.
* `seq` - Seq.
* `subnet` - Subnet.

The `vpn_src_node` block supports:

* `host` - Host.
* `seq` - Seq.
* `subnet` - Subnet.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_policy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
