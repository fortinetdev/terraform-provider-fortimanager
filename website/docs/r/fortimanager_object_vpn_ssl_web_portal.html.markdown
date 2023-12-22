---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_portal"
description: |-
  Portal.
---

# fortimanager_object_vpn_ssl_web_portal
Portal.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`bookmark_group`: `fortimanager_object_vpn_ssl_web_portal_bookmarkgroup`
`landing_page`: `fortimanager_object_vpn_ssl_web_portal_landingpage`
`mac_addr_check_rule`: `fortimanager_object_vpn_ssl_web_portal_macaddrcheckrule`
`os_check_list`: `fortimanager_object_vpn_ssl_web_portal_oschecklist`
`split_dns`: `fortimanager_object_vpn_ssl_web_portal_splitdns`



## Example Usage

```hcl
resource "fortimanager_object_vpn_ssl_web_portal" "trname" {
  allow_user_access                   = ["citrix", "ftp", "ping", "portforward", "rdp", "sftp", "smb", "ssh", "telnet", "vnc", "web"]
  customize_forticlient_download_url  = "disable"
  exclusive_routing                   = "disable"
  forticlient_download                = "enable"
  forticlient_download_method         = "direct"
  ipv6_exclusive_routing              = "disable"
  ipv6_service_restriction            = "disable"
  ipv6_split_tunneling_routing_negate = "disable"
  ipv6_tunnel_mode                    = "disable"
  limit_user_logins                   = "disable"
  name                                = "terr-vpn-ssl-web-portal"
  service_restriction                 = "disable"
  skip_check_for_browser              = "enable"
  split_tunneling_routing_negate      = "disable"
  tunnel_mode                         = "disable"
  use_sdwan                           = "disable"
  web_mode                            = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `allow_user_access` - Allow user access to SSL-VPN applications. Valid values: `web`, `ftp`, `telnet`, `smb`, `vnc`, `rdp`, `ssh`, `ping`, `citrix`, `portforward`, `sftp`.

* `auto_connect` - Enable/disable automatic connect by client when system is up. Valid values: `disable`, `enable`.

* `bookmark_group` - Bookmark-Group. The structure of `bookmark_group` block is documented below.
* `client_src_range` - Allow client to add source range for the tunnel traffic. Valid values: `disable`, `enable`.

* `clipboard` - Enable to support RDP/VPC clipboard functionality. Valid values: `disable`, `enable`.

* `custom_lang` - Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files.
* `customize_forticlient_download_url` - Enable support of customized download URL for FortiClient. Valid values: `disable`, `enable`.

* `default_protocol` - Application type that is set by default. Valid values: `web`, `ftp`, `telnet`, `smb`, `vnc`, `rdp`, `ssh`, `sftp`.

* `default_window_height` - Screen height (range from 0 - 65535, default = 768).
* `default_window_width` - Screen width (range from 0 - 65535, default = 1024).
* `dhcp_ip_overlap` - Configure overlapping DHCP IP allocation assignment. Valid values: `use-old`, `use-new`.

* `dhcp_ra_giaddr` - Relay agent gateway IP address to use in the giaddr field of DHCP requests.
* `dhcp6_ra_linkaddr` - Relay agent IPv6 link address to use in DHCP6 requests.
* `display_bookmark` - Enable to display the web portal bookmark widget. Valid values: `disable`, `enable`.

* `display_connection_tools` - Enable to display the web portal connection tools widget. Valid values: `disable`, `enable`.

* `display_history` - Enable to display the web portal user login history widget. Valid values: `disable`, `enable`.

* `display_status` - Enable to display the web portal status widget. Valid values: `disable`, `enable`.

* `dns_server1` - IPv4 DNS server 1.
* `dns_server2` - IPv4 DNS server 2.
* `dns_suffix` - DNS suffix.
* `exclusive_routing` - Enable/disable all traffic go through tunnel only. Valid values: `disable`, `enable`.

* `focus_bookmark` - Enable to prioritize the placement of the bookmark section over the quick-connection section in the SSL-VPN application. Valid values: `disable`, `enable`.

* `forticlient_download` - Enable/disable download option for FortiClient. Valid values: `disable`, `enable`.

* `forticlient_download_method` - FortiClient download method. Valid values: `direct`, `ssl-vpn`.

* `heading` - Web portal heading message.
* `hide_sso_credential` - Enable to prevent SSO credential being sent to client. Valid values: `disable`, `enable`.

* `host_check` - Type of host checking performed on endpoints. Valid values: `none`, `av`, `fw`, `av-fw`, `custom`.

* `host_check_interval` - Periodic host check interval. Value of 0 means disabled and host checking only happens when the endpoint connects.
* `host_check_policy` - One or more policies to require the endpoint to have specific security software.
* `ip_mode` - Method by which users of this SSL-VPN tunnel obtain IP addresses. Valid values: `range`, `user-group`.

* `ip_pools` - IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_exclusive_routing` - Enable/disable all IPv6 traffic go through tunnel only. Valid values: `disable`, `enable`.

* `ipv6_pools` - IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients.
* `ipv6_service_restriction` - Enable/disable IPv6 tunnel service restriction. Valid values: `disable`, `enable`.

* `ipv6_split_tunneling` - Enable/disable IPv6 split tunneling. Valid values: `disable`, `enable`.

* `ipv6_split_tunneling_routing_address` - IPv6 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.
* `ipv6_split_tunneling_routing_negate` - Enable to negate IPv6 split tunneling routing address. Valid values: `disable`, `enable`.

* `ipv6_tunnel_mode` - Enable/disable IPv6 SSL-VPN tunnel mode. Valid values: `disable`, `enable`.

* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `keep_alive` - Enable/disable automatic reconnect for FortiClient connections. Valid values: `disable`, `enable`.

* `landing_page` - Landing-Page. The structure of `landing_page` block is documented below.
* `landing_page_mode` - Enable/disable SSL-VPN landing page mode. Valid values: `disable`, `enable`.

* `limit_user_logins` - Enable to limit each user to one SSL-VPN session at a time. Valid values: `disable`, `enable`.

* `mac_addr_action` - Client MAC address action. Valid values: `deny`, `allow`.

* `mac_addr_check` - Enable/disable MAC address host checking. Valid values: `disable`, `enable`.

* `mac_addr_check_rule` - Mac-Addr-Check-Rule. The structure of `mac_addr_check_rule` block is documented below.
* `macos_forticlient_download_url` - Download URL for Mac FortiClient.
* `name` - Portal name.
* `os_check` - Enable to let the FortiGate decide action based on client OS. Valid values: `disable`, `enable`.

* `os_check_list` - Os-Check-List. The structure of `os_check_list` block is documented below.
* `prefer_ipv6_dns` - prefer to query IPv6 dns first if enabled. Valid values: `disable`, `enable`.

* `redir_url` - Client login redirect URL.
* `rewrite_ip_uri_ui` - Rewrite contents for URI contains IP and "/ui/". (default = disable) Valid values: `disable`, `enable`.

* `save_password` - Enable/disable FortiClient saving the user's password. Valid values: `disable`, `enable`.

* `service_restriction` - Enable/disable tunnel service restriction. Valid values: `disable`, `enable`.

* `skip_check_for_unsupported_browser` - Enable to skip host check if browser does not support it. Valid values: `disable`, `enable`.

* `skip_check_for_browser` - Enable to skip host check for browser support. Valid values: `disable`, `enable`.

* `skip_check_for_unsupported_os` - Enable to skip host check if client OS does not support it. Valid values: `disable`, `enable`.

* `smb_max_version` - SMB maximum client protocol version. Valid values: `smbv1`, `smbv2`, `smbv3`.

* `smb_min_version` - SMB minimum client protocol version. Valid values: `smbv1`, `smbv2`, `smbv3`.

* `smb_ntlmv1_auth` - Enable support of NTLMv1 for Samba authentication. Valid values: `disable`, `enable`.

* `smbv1` - Smbv1. Valid values: `disable`, `enable`.

* `split_dns` - Split-Dns. The structure of `split_dns` block is documented below.
* `split_tunneling` - Enable/disable IPv4 split tunneling. Valid values: `disable`, `enable`.

* `split_tunneling_routing_address` - IPv4 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access.
* `split_tunneling_routing_negate` - Enable to negate split tunneling routing address. Valid values: `disable`, `enable`.

* `theme` - Web portal color scheme. Valid values: `gray`, `blue`, `orange`, `crimson`, `steelblue`, `darkgrey`, `green`, `melongene`, `red`, `mariner`, `neutrino`.

* `transform_backward_slashes` - Transform backward slashes to forward slashes in URLs. Valid values: `disable`, `enable`.

* `tunnel_mode` - Enable/disable IPv4 SSL-VPN tunnel mode. Valid values: `disable`, `enable`.

* `use_sdwan` - Use SD-WAN rules to get output interface. Valid values: `disable`, `enable`.

* `user_bookmark` - Enable to allow web portal users to create their own bookmarks. Valid values: `disable`, `enable`.

* `user_group_bookmark` - Enable to allow web portal users to create bookmarks for all users in the same user group. Valid values: `disable`, `enable`.

* `web_mode` - Enable/disable SSL VPN web mode. Valid values: `disable`, `enable`.

* `windows_forticlient_download_url` - Download URL for Windows FortiClient.
* `wins_server1` - IPv4 WINS server 1.
* `wins_server2` - IPv4 WINS server 1.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `bookmark_group` block supports:

* `bookmarks` - Bookmarks. The structure of `bookmarks` block is documented below.
* `name` - Bookmark group name.

The `bookmarks` block supports:

* `additional_params` - Additional parameters.
* `apptype` - Application type. Valid values: `web`, `telnet`, `ssh`, `ftp`, `smb`, `vnc`, `rdp`, `citrix`, `rdpnative`, `portforward`, `sftp`.

* `color_depth` - Color depth per pixel. Valid values: `8`, `16`, `32`.

* `description` - Description.
* `domain` - Login domain.
* `folder` - Network shared file folder parameter.
* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `height` - Screen height (range from 480 - 65535, default = 768).
* `host` - Host name/IP parameter.
* `keyboard_layout` - Keyboard layout. Valid values: `ar`, `da`, `de`, `de-ch`, `en-gb`, `en-uk`, `en-us`, `es`, `fi`, `fr`, `fr-be`, `fr-ca`, `fr-ch`, `hr`, `hu`, `it`, `ja`, `lt`, `lv`, `mk`, `no`, `pl`, `pt`, `pt-br`, `ru`, `sl`, `sv`, `tk`, `tr`, `fr-ca-m`, `wg`, `ar-101`, `ar-102`, `ar-102-azerty`, `can-mul`, `cz`, `cz-qwerty`, `cz-pr`, `nl`, `de-ibm`, `en-uk-ext`, `en-us-dvorak`, `es-var`, `fi-sami`, `hu-101`, `it-142`, `ko`, `lt-ibm`, `lt-std`, `lav-std`, `lav-leg`, `mk-std`, `no-sami`, `pol-214`, `pol-pr`, `pt-br-abnt2`, `ru-mne`, `ru-t`, `sv-sami`, `tuk`, `tur-f`, `tur-q`, `zh-sym-sg-us`, `zh-sym-us`, `zh-tr-hk`, `zh-tr-mo`, `zh-tr-us`.

* `listening_port` - Listening port (0 - 65535).
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `logon_password` - Logon password.
* `logon_user` - Logon user.
* `name` - Bookmark name.
* `port` - Remote port.
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `preconnection_id` - The numeric ID of the RDP source (0-2147483648).
* `remote_port` - Remote port (0 - 65535).
* `restricted_admin` - Enable/disable restricted admin mode for RDP. Valid values: `disable`, `enable`.

* `security` - Security mode for RDP connection. Valid values: `rdp`, `nla`, `tls`, `any`.

* `send_preconnection_id` - Enable/disable sending of preconnection ID. Valid values: `disable`, `enable`.

* `server_layout` - Server side keyboard layout. Valid values: `en-us-qwerty`, `de-de-qwertz`, `fr-fr-azerty`, `it-it-qwerty`, `sv-se-qwerty`, `failsafe`, `en-gb-qwerty`, `es-es-qwerty`, `fr-ch-qwertz`, `ja-jp-qwerty`, `pt-br-qwerty`, `tr-tr-qwerty`, `fr-ca-qwerty`.

* `show_status_window` - Enable/disable showing of status window. Valid values: `disable`, `enable`.

* `sso` - Single Sign-On. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server. Valid values: `disable`, `enable`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - URL parameter.
* `vnc_keyboard_layout` - Keyboard layout. Valid values: `da`, `de`, `de-ch`, `en-uk`, `es`, `fi`, `fr`, `fr-be`, `it`, `no`, `pt`, `sv`, `nl`, `en-uk-ext`, `it-142`, `pt-br-abnt2`, `default`, `fr-ca-mul`, `gd`, `us-intl`.

* `width` - Screen width (range from 640 - 65535, default = 1024).

The `form_data` block supports:

* `name` - Name.
* `value` - Value.

The `landing_page` block supports:

* `form_data` - Form-Data. The structure of `form_data` block is documented below.
* `logout_url` - Landing page log out URL.
* `sso` - Single sign-on. Valid values: `disable`, `static`, `auto`.

* `sso_credential` - Single sign-on credentials. Valid values: `sslvpn-login`, `alternative`.

* `sso_password` - SSO password.
* `sso_username` - SSO user name.
* `url` - Landing page URL.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.

The `mac_addr_check_rule` block supports:

* `mac_addr_list` - Client MAC address list.
* `mac_addr_mask` - Client MAC address mask.
* `name` - Client MAC address check rule name.

The `os_check_list` block supports:

* `action` - OS check options. Valid values: `allow`, `check-up-to-date`, `deny`.

* `latest_patch_level` - Latest OS patch level.
* `name` - Name.
* `tolerance` - OS patch level tolerance.

The `split_dns` block supports:

* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `domains` - Split DNS domains used for SSL-VPN clients separated by comma(,).
* `id` - ID.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn SslWebPortal can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_portal.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
