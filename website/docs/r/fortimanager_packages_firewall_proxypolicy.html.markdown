---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_proxypolicy"
description: |-
  Configure proxy policies.
---

# fortimanager_packages_firewall_proxypolicy
Configure proxy policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_proxypolicy" "labelname" {
  action              = "deny"
  disclaimer          = "disable"
  dstaddr             = ["all"]
  dstaddr_negate      = "disable"
  dstintf             = ["any"]
  http_tunnel_auth    = "disable"
  internet_service    = "disable"
  logtraffic          = "all"
  logtraffic_start    = "disable"
  name                = "11"
  pkg                 = "default"
  policyid            = 1
  profile_type        = "single"
  proxy               = "explicit-web"
  schedule            = "always"
  service             = ["webproxy"]
  service_negate      = "disable"
  session_ttl         = 0
  srcaddr             = ["all"]
  srcaddr_negate      = "disable"
  ssh_policy_redirect = "disable"
  status              = "enable"
  webcache            = "disable"
  webcache_https      = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `access_proxy` - Access Proxy.
* `access_proxy6` - IPv6 access proxy.
* `action` - Accept or deny traffic matching the policy parameters. Valid values: `accept`, `deny`, `redirect`.

* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `block_notification` - Enable/disable block notification. Valid values: `disable`, `enable`.

* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Optional comments.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `device_ownership` - When enabled, the ownership enforcement will be done at policy level. Valid values: `disable`, `enable`.

* `disclaimer` - Web proxy disclaimer setting: by domain, policy, or user. Valid values: `disable`, `domain`, `policy`, `user`.

* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dstaddr` - Destination address objects.
* `dstaddr_negate` - When enabled, destination addresses match against any address EXCEPT the specified destination addresses. Valid values: `disable`, `enable`.

* `dstaddr6` - IPv6 destination address objects.
* `dstintf` - Destination interface names.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `global_label` - Global web-based manager visible label.
* `groups` - Names of group objects.
* `http_tunnel_auth` - Enable/disable HTTP tunnel authentication. Valid values: `disable`, `enable`.

* `icap_profile` - Name of an existing ICAP profile.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_name` - Internet Service name.
* `internet_service_id` - Internet Service ID.
* `internet_service_negate` - When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service. Valid values: `disable`, `enable`.

* `ips_sensor` - Name of an existing IPS sensor.
* `label` - VDOM-specific GUI visible label.
* `logtraffic` - Enable/disable logging traffic through the policy. Valid values: `disable`, `all`, `utm`.

* `logtraffic_start` - Enable/disable policy log traffic start. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `mms_profile` - Name of an existing MMS profile.
* `policyid` - Policy ID.
* `poolname` - Name of IP pool object.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.

* `proxy` - Type of explicit proxy. Valid values: `explicit-web`, `transparent-web`, `ftp`, `wanopt`, `ssh`, `ssh-tunnel`.

* `redirect_url` - Redirect URL for further explicit web proxy processing.
* `replacemsg_override_group` - Authentication replacement message override group.
* `scan_botnet_connections` - Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.

* `schedule` - Name of schedule object.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `service` - Name of service objects.
* `service_negate` - When enabled, services match against any service EXCEPT the specified destination services. Valid values: `disable`, `enable`.

* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `srcaddr` - Source address objects.
* `srcaddr_negate` - When enabled, source addresses match against any address EXCEPT the specified source addresses. Valid values: `disable`, `enable`.

* `srcaddr6` - IPv6 source address objects.
* `srcintf` - Source interface names.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `disable`, `enable`.

* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable/disable the active status of the policy. Valid values: `disable`, `enable`.

* `transparent` - Enable to use the IP address of the client to connect to the server. Valid values: `disable`, `enable`.

* `users` - Names of user objects.
* `utm_status` - Enable the use of UTM profiles/sensors/lists. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `voip_profile` - Name of an existing VoIP profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webcache` - Enable/disable web caching. Valid values: `disable`, `enable`.

* `webcache_https` - Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile). Valid values: `disable`, `enable`.

* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Name of web proxy profile.
* `ztna_ems_tag` - ZTNA EMS Tag names.
* `ztna_tags_match_logic` - ZTNA tag matching logic. Valid values: `or`, `and`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallProxyPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_proxypolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
