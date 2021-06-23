---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_consolidated_policy"
description: |-
  Configure consolidated IPv4/IPv6 policies.
---

# fortimanager_packages_firewall_consolidated_policy
Configure consolidated IPv4/IPv6 policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Policy action (allow/deny/ipsec). Valid values: `deny`, `accept`, `ipsec`.

* `app_category` - App-Category.
* `app_group` - App-Group.
* `application` - Application.
* `application_list` - Name of an existing Application list.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `disable`, `enable`.

* `av_profile` - Name of an existing Antivirus profile.
* `captive_portal_exempt` - Enable exemption of some users from the captive portal. Valid values: `disable`, `enable`.

* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `dstaddr4` - Destination IPv4 address name and address group names.
* `dstaddr6` - Destination IPv6 address name and address group names.
* `dstintf` - Outgoing (egress) interface.
* `emailfilter_profile` - Name of an existing email filter profile.
* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `disable`, `enable`.

* `fsso_groups` - Names of FSSO groups.
* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `groups` - Names of user groups that can authenticate with this policy.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `disable`, `enable`.

* `icap_profile` - Name of an existing ICAP profile.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `disable`, `enable`.

* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy`, `flow`.

* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_id` - Internet Service ID.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_id` - Internet Service source ID.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `ippool` - Enable to use IP Pools for source NAT. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `disable`, `all`, `utm`.

* `logtraffic_start` - Record logs when a session starts. Valid values: `disable`, `enable`.

* `mms_profile` - Name of an existing MMS profile.
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `disable`, `enable`.

* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `disable`, `enable`.

* `per_ip_shaper` - Per-IP traffic shaper.
* `policyid` - Policy ID (0 - 4294967294).
* `poolname4` - IPv4 pool names.
* `poolname6` - IPv6 pool names.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `schedule` - Schedule name.
* `service` - Service and service group names.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `srcaddr4` - Source IPv4 address name and address group names.
* `srcaddr6` - Source IPv6 address name and address group names.
* `srcintf` - Incoming (ingress) interface.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `disable`, `enable`.

* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy. Valid values: `disable`, `enable`.

* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `url_category` - Url-Category.
* `users` - Names of individual users that can authenticate with this policy.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `voip_profile` - Name of an existing VoIP profile.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `waf_profile` - Name of an existing Web application firewall profile.
* `wanopt` - Enable/disable WAN optimization. Valid values: `disable`, `enable`.

* `wanopt_detection` - WAN optimization auto-detection mode. Valid values: `active`, `passive`, `off`.

* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect to server. Valid values: `default`, `transparent`, `non-transparent`.

* `wanopt_peer` - WAN optimization peer.
* `wanopt_profile` - WAN optimization profile.
* `webcache` - Enable/disable web cache. Valid values: `disable`, `enable`.

* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable`, `enable`.

* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Webproxy forward server name.
* `webproxy_profile` - Webproxy profile name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallConsolidatedPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_consolidated_policy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
