---
subcategory: "Global Footer/Header Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_global_footer_policy6"
description: |-
  Configure IPv6 policies.
---

# fortimanager_packages_global_footer_policy6
Configure IPv6 policies.

## Example Usage

```hcl
resource "fortimanager_packages_global_header_policy6" "labelname" {
  pkg      = "default"
  action   = "accept"
  comments = "terraform-comment"
  dstaddr  = "gall"
  dstintf  = "any"
  name     = "terraform-tefv-footer"
  schedule = "galways"
  service  = "gALL"
  srcaddr  = "gall"
  srcintf  = "any"
  status   = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `pkg` - Package.

* `action` - Policy action (allow/deny/ipsec). Valid values: `deny`, `accept`, `ipsec`, `ssl-vpn`.

* `anti_replay` - Enable/disable anti-replay check. Valid values: `disable`, `enable`.

* `app_category` - Application category ID list.
* `app_group` - Application group names.
* `application` - Application ID list.
* `application_charts` - Application-Charts. Valid values: `top10-app`, `top10-p2p-user`, `top10-media-user`.

* `application_list` - Name of an existing Application list.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `disable`, `enable`.

* `av_profile` - Name of an existing Antivirus profile.
* `casi_profile` - Casi-Profile.
* `cgn_log_server_grp` - NP log server group name
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `custom_log_fields` - Log field index numbers to append custom log fields to log messages for this policy.
* `decrypted_traffic_mirror` - Decrypted-Traffic-Mirror.
* `deep_inspection_options` - Deep-Inspection-Options.
* `device_detection_portal` - Device-Detection-Portal. Valid values: `disable`, `enable`.

* `devices` - Devices.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dscp_match` - Dscp-Match. Valid values: `disable`, `enable`.

* `dscp_negate` - Dscp-Negate. Valid values: `disable`, `enable`.

* `dscp_value` - Dscp-Value.
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `disable`, `enable`.

* `dstaddr` - Destination address and address group names.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstintf` - Outgoing (egress) interface.
* `dynamic_profile` - Dynamic-Profile. Valid values: `disable`, `enable`.

* `dynamic_profile_access` - Dynamic-Profile-Access. Valid values: `imap`, `smtp`, `pop3`, `http`, `ftp`, `im`, `nntp`, `imaps`, `smtps`, `pop3s`, `https`, `ftps`.

* `dynamic_profile_group` - Dynamic-Profile-Group.
* `email_collection_portal` - Email-Collection-Portal. Valid values: `disable`, `enable`.

* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - File-Filter-Profile.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all`, `check-new`.

* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `disable`, `enable`.

* `fsae` - Fsae. Valid values: `disable`, `enable`.

* `fsso_groups` - Names of FSSO groups.
* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `groups` - Names of user groups that can authenticate with this policy.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `disable`, `enable`.

* `icap_profile` - Name of an existing ICAP profile.
* `identity_based` - Identity-Based. Valid values: `disable`, `enable`.

* `identity_from` - Identity-From. Valid values: `auth`, `device`.

* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `disable`, `enable`.

* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy`, `flow`.

* `ippool` - Enable to use IP Pools for source NAT. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `enable`, `all`, `utm`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.

* `mms_profile` - Name of an existing MMS profile.
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic. Valid values: `disable`, `enable`.

* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic. Valid values: `disable`, `enable`.

* `np_accelation` - Np-Accelation. Valid values: `disable`, `enable`.

* `np_acceleration` - Enable/disable UTM Network Processor acceleration. Valid values: `disable`, `enable`.

* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `disable`, `enable`.

* `per_ip_shaper` - Per-IP traffic shaper.
* `policy_offload` - Enable/disable offloading policy configuration to CP processors. Valid values: `disable`, `enable`.

* `policyid` - Policy ID (0 - 4294967294).
* `poolname` - IP Pool names.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `replacemsg_group` - Replacemsg-Group.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `disable`, `enable`.

* `schedule` - Schedule name.
* `send_deny_packet` - Enable/disable return of deny-packet. Valid values: `disable`, `enable`.

* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `session_ttl` - Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.
* `spamfilter_profile` - Spamfilter-Profile.
* `srcaddr` - Source address and address group names.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcintf` - Incoming (ingress) interface.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `disable`, `enable`.

* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring). Valid values: `disable`, `enable`.

* `ssl_mirror_intf` - SSL mirror interface name.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `sslvpn_auth` - Sslvpn-Auth. Valid values: `any`, `local`, `radius`, `ldap`, `tacacs+`.

* `sslvpn_ccert` - Sslvpn-Ccert. Valid values: `disable`, `enable`.

* `sslvpn_cipher` - Sslvpn-Cipher. Valid values: `any`, `high`, `medium`.

* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `tags` - Tags.
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag. Valid values: `all`, `data-only`, `disable`.

* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire. Valid values: `disable`, `enable`.

* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `disable`, `enable`.

* `traffic_shaper` - Reverse traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `url_category` - URL category ID list.
* `users` - Names of individual users that can authenticate with this policy.
* `utm_inspection_mode` - Utm-Inspection-Mode. Valid values: `proxy`, `flow`.

* `utm_status` - Enable AV/web/ips protection profile. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webcache` - Enable/disable web cache. Valid values: `disable`, `enable`.

* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable`, `enable`.

* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Webproxy profile name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages GlobalFooterPolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_global_footer_policy6.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

