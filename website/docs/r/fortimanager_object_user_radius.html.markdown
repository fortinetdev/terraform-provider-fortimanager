---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_radius"
description: |-
  Configure RADIUS server entries.
---

# fortimanager_object_user_radius
Configure RADIUS server entries.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `accounting_server`: `fortimanager_object_user_radius_accountingserver`
>- `dynamic_mapping`: `fortimanager_object_user_radius_dynamic_mapping`



## Example Usage

```hcl
resource "fortimanager_object_user_radius" "labelname" {
  acct_all_servers                            = "disable"
  acct_interim_interval                       = 0
  all_usergroup                               = "disable"
  auth_type                                   = "auto"
  class                                       = []
  h3c_compatibility                           = "disable"
  interface_select_method                     = "auto"
  name                                        = "fdsafds"
  nas_ip                                      = "0.0.0.0"
  password_encoding                           = "auto"
  password_renewal                            = "enable"
  radius_coa                                  = "disable"
  radius_port                                 = 0
  rsso                                        = "disable"
  rsso_context_timeout                        = 0
  rsso_ep_one_ip_only                         = "disable"
  rsso_log_flags                              = []
  rsso_log_period                             = 0
  rsso_radius_server_port                     = 0
  rsso_secret                                 = []
  secondary_secret                            = ["tesssssss"]
  secret                                      = ["tesssssss"]
  server                                      = "2.2.2.2"
  sso_attribute_value_override                = "enable"
  switch_controller_acct_fast_framedip_detect = 2
  switch_controller_service_type              = []
  tertiary_secret                             = ["tesssssss"]
  timeout                                     = 5
  use_management_vdom                         = "disable"
  username_case_sensitive                     = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `account_key_cert_field` - Define subject identity field in certificate for user access right checking. Valid values: `othername`, `rfc822name`, `dnsname`.

* `account_key_processing` - Account key processing operation. The FortiGate will keep either the whole domain or strip the domain from the subject identity. Valid values: `same`, `strip`.

* `accounting_server` - Accounting-Server. The structure of `accounting_server` block is documented below.
* `acct_all_servers` - Enable/disable sending of accounting messages to all configured servers (default = disable). Valid values: `disable`, `enable`.

* `acct_interim_interval` - Time in seconds between each accounting interim update message.
* `all_usergroup` - Enable/disable automatically including this RADIUS server in all user groups. Valid values: `disable`, `enable`.

* `auth_type` - Authentication methods/protocols permitted for this RADIUS server. Valid values: `pap`, `chap`, `ms_chap`, `ms_chap_v2`, `auto`.

* `ca_cert` - CA of server to trust under TLS.
* `call_station_id_type` - Calling & Called station identifier type configuration (default = legacy), this option is not available for 802.1x authentication. Valid values: `legacy`, `IP`, `MAC`.

* `class` - Class attribute name(s).
* `client_cert` - Client certificate to use under TLS.
* `delimiter` - Configure delimiter to be used for separating profile group names in the SSO attribute (default = plus character "+"). Valid values: `plus`, `comma`.

* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `group_override_attr_type` - RADIUS attribute type to override user group information. Valid values: `filter-Id`, `class`.

* `h3c_compatibility` - Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `mac_case` - MAC authentication case (default = lowercase). Valid values: `uppercase`, `lowercase`.

* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `name` - RADIUS server entry name.
* `nas_id` - Custom NAS identifier.
* `nas_id_type` - NAS identifier type configuration (default = legacy). Valid values: `legacy`, `custom`, `hostname`.

* `nas_ip` - IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.
* `password_encoding` - Password encoding. Valid values: `ISO-8859-1`, `auto`.

* `password_renewal` - Enable/disable password renewal. Valid values: `disable`, `enable`.

* `radius_coa` - Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated. Valid values: `disable`, `enable`.

* `radius_port` - RADIUS service port number.
* `rsso` - Enable/disable RADIUS based single sign on feature. Valid values: `disable`, `enable`.

* `rsso_context_timeout` - Time in seconds before the logged out user is removed from the "user context list" of logged on users.
* `rsso_endpoint_attribute` - RADIUS attributes used to extract the user end point identifer from the RADIUS Start record. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `rsso_endpoint_block_attribute` - RADIUS attributes used to block a user. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `rsso_ep_one_ip_only` - Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages. Valid values: `disable`, `enable`.

* `rsso_flush_ip_session` - Enable/disable flushing user IP sessions on RADIUS accounting Stop messages. Valid values: `disable`, `enable`.

* `rsso_log_flags` - Events to log. Valid values: `none`, `protocol-error`, `profile-missing`, `context-missing`, `accounting-stop-missed`, `accounting-event`, `radiusd-other`, `endpoint-block`.

* `rsso_log_period` - Time interval in seconds that group event log messages will be generated for dynamic profile events.
* `rsso_radius_response` - Enable/disable sending RADIUS response packets after receiving Start and Stop records. Valid values: `disable`, `enable`.

* `rsso_radius_server_port` - UDP port to listen on for RADIUS Start and Stop records.
* `rsso_secret` - RADIUS secret used by the RADIUS accounting server.
* `rsso_validate_request_secret` - Enable/disable validating the RADIUS request shared secret in the Start or End record. Valid values: `disable`, `enable`.

* `secondary_secret` - Secret key to access the secondary server.
* `secondary_server` - {&lt;name_str|ip_str&gt;} secondary RADIUS CN domain name or IP.
* `secret` - Pre-shared secret key used to access the primary RADIUS server.
* `server` - Primary RADIUS server CN domain name or IP address.
* `server_identity_check` - Enable/disable RADIUS server identity check (verify server domain name/IP address against the server certificate). Valid values: `disable`, `enable`.

* `source_ip` - Source IP address for communications to the RADIUS server.
* `source_ip_interface` - Source interface for communication with the RADIUS server.
* `sso_attribute` - RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `sso_attribute_key` - Key prefix for SSO group value in the SSO attribute.
* `sso_attribute_value_override` - Enable/disable override old attribute value with new value for the same endpoint. Valid values: `disable`, `enable`.

* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `switch_controller_acct_fast_framedip_detect` - Switch controller accounting message Framed-IP detection from DHCP snooping (seconds, default=2).
* `switch_controller_nas_ip_dynamic` - Enable/Disable switch-controller nas-ip dynamic to dynamically set nas-ip. Valid values: `disable`, `enable`.

* `switch_controller_service_type` - RADIUS service type. Valid values: `login`, `framed`, `callback-login`, `callback-framed`, `outbound`, `administrative`, `nas-prompt`, `authenticate-only`, `callback-nas-prompt`, `call-check`, `callback-administrative`.

* `tertiary_secret` - Secret key to access the tertiary server.
* `tertiary_server` - {&lt;name_str|ip_str&gt;} tertiary RADIUS CN domain name or IP.
* `timeout` - Time in seconds between re-sending authentication requests.
* `tls_min_proto_version` - Minimum supported protocol version for TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `transport_protocol` - Transport protocol to be used (default = udp). Valid values: `udp`, `tcp`, `tls`.

* `use_management_vdom` - Enable/disable using management VDOM to send requests. Valid values: `disable`, `enable`.

* `username_case_sensitive` - Enable/disable case sensitive user names. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `accounting_server` block supports:

* `id` - ID (0 - 4294967295).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `port` - RADIUS accounting port number.
* `secret` - Secret key.
* `server` - {&lt;name_str|ip_str&gt;} Server CN domain name or IP.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `status` - Status. Valid values: `disable`, `enable`.


The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `account_key_cert_field` - Define subject identity field in certificate for user access right checking. Valid values: `othername`, `rfc822name`, `dnsname`.

* `account_key_processing` - Account key processing operation. The FortiGate will keep either the whole domain or strip the domain from the subject identity. Valid values: `same`, `strip`.

* `accounting_server` - Accounting-Server. The structure of `accounting_server` block is documented below.
* `acct_all_servers` - Enable/disable sending of accounting messages to all configured servers (default = disable). Valid values: `disable`, `enable`.

* `acct_interim_interval` - Time in seconds between each accounting interim update message.
* `all_usergroup` - Enable/disable automatically including this RADIUS server in all user groups. Valid values: `disable`, `enable`.

* `auth_type` - Authentication methods/protocols permitted for this RADIUS server. Valid values: `pap`, `chap`, `ms_chap`, `ms_chap_v2`, `auto`.

* `ca_cert` - CA of server to trust under TLS.
* `call_station_id_type` - Calling & Called station identifier type configuration (default = legacy), this option is not available for 802.1x authentication. Valid values: `legacy`, `IP`, `MAC`.

* `class` - Class attribute name(s).
* `client_cert` - Client certificate to use under TLS.
* `delimiter` - Delimiter. Valid values: `plus`, `comma`.

* `dp_carrier_endpoint_attribute` - Dp-Carrier-Endpoint-Attribute. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Vendor-Specific`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `dp_carrier_endpoint_block_attribute` - Dp-Carrier-Endpoint-Block-Attribute. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Vendor-Specific`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `dp_context_timeout` - Dp-Context-Timeout.
* `dp_flush_ip_session` - Dp-Flush-Ip-Session. Valid values: `disable`, `enable`.

* `dp_hold_time` - Dp-Hold-Time.
* `dp_http_header` - Dp-Http-Header.
* `dp_http_header_fallback` - Dp-Http-Header-Fallback. Valid values: `ip-header-address`, `default-profile`.

* `dp_http_header_status` - Dp-Http-Header-Status. Valid values: `disable`, `enable`.

* `dp_http_header_suppress` - Dp-Http-Header-Suppress. Valid values: `disable`, `enable`.

* `dp_log_dyn_flags` - Dp-Log-Dyn_Flags. Valid values: `none`, `protocol-error`, `profile-missing`, `context-missing`, `accounting-stop-missed`, `accounting-event`, `radiusd-other`, `endpoint-block`.

* `dp_log_period` - Dp-Log-Period.
* `dp_mem_percent` - Dp-Mem-Percent.
* `dp_profile_attribute` - Dp-Profile-Attribute. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Vendor-Specific`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `dp_profile_attribute_key` - Dp-Profile-Attribute-Key.
* `dp_radius_response` - Dp-Radius-Response. Valid values: `disable`, `enable`.

* `dp_radius_server_port` - Dp-Radius-Server-Port.
* `dp_secret` - Dp-Secret.
* `dp_validate_request_secret` - Dp-Validate-Request-Secret. Valid values: `disable`, `enable`.

* `dynamic_profile` - Dynamic-Profile. Valid values: `disable`, `enable`.

* `endpoint_translation` - Endpoint-Translation. Valid values: `disable`, `enable`.

* `ep_carrier_endpoint_convert_hex` - Ep-Carrier-Endpoint-Convert-Hex. Valid values: `disable`, `enable`.

* `ep_carrier_endpoint_header` - Ep-Carrier-Endpoint-Header.
* `ep_carrier_endpoint_header_suppress` - Ep-Carrier-Endpoint-Header-Suppress. Valid values: `disable`, `enable`.

* `ep_carrier_endpoint_prefix` - Ep-Carrier-Endpoint-Prefix. Valid values: `disable`, `enable`.

* `ep_carrier_endpoint_prefix_range_max` - Ep-Carrier-Endpoint-Prefix-Range-Max.
* `ep_carrier_endpoint_prefix_range_min` - Ep-Carrier-Endpoint-Prefix-Range-Min.
* `ep_carrier_endpoint_prefix_string` - Ep-Carrier-Endpoint-Prefix-String.
* `ep_carrier_endpoint_source` - Ep-Carrier-Endpoint-Source. Valid values: `http-header`, `cookie`.

* `ep_ip_header` - Ep-Ip-Header.
* `ep_ip_header_suppress` - Ep-Ip-Header-Suppress. Valid values: `disable`, `enable`.

* `ep_missing_header_fallback` - Ep-Missing-Header-Fallback. Valid values: `session-ip`, `policy-profile`.

* `ep_profile_query_type` - Ep-Profile-Query-Type. Valid values: `session-ip`, `extract-ip`, `extract-carrier-endpoint`.

* `group_override_attr_type` - Group-Override-Attr-Type. Valid values: `filter-Id`, `class`.

* `h3c_compatibility` - Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `mac_case` - MAC authentication case (default = lowercase). Valid values: `uppercase`, `lowercase`.

* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.

* `nas_id` - Custom NAS identifier.
* `nas_id_type` - NAS identifier type configuration (default = legacy). Valid values: `legacy`, `custom`, `hostname`.

* `nas_ip` - IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.
* `password_encoding` - Password encoding. Valid values: `ISO-8859-1`, `auto`.

* `password_renewal` - Enable/disable password renewal. Valid values: `disable`, `enable`.

* `radius_coa` - Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated. Valid values: `disable`, `enable`.

* `radius_port` - RADIUS service port number.
* `rsso` - Enable/disable RADIUS based single sign on feature. Valid values: `disable`, `enable`.

* `rsso_context_timeout` - Time in seconds before the logged out user is removed from the "user context list" of logged on users.
* `rsso_endpoint_attribute` - RADIUS attributes used to extract the user end point identifer from the RADIUS Start record. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `rsso_endpoint_block_attribute` - RADIUS attributes used to block a user. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `rsso_ep_one_ip_only` - Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages. Valid values: `disable`, `enable`.

* `rsso_flush_ip_session` - Enable/disable flushing user IP sessions on RADIUS accounting Stop messages. Valid values: `disable`, `enable`.

* `rsso_log_flags` - Events to log. Valid values: `none`, `protocol-error`, `profile-missing`, `context-missing`, `accounting-stop-missed`, `accounting-event`, `radiusd-other`, `endpoint-block`.

* `rsso_log_period` - Time interval in seconds that group event log messages will be generated for dynamic profile events.
* `rsso_radius_response` - Enable/disable sending RADIUS response packets after receiving Start and Stop records. Valid values: `disable`, `enable`.

* `rsso_radius_server_port` - UDP port to listen on for RADIUS Start and Stop records.
* `rsso_secret` - RADIUS secret used by the RADIUS accounting server.
* `rsso_validate_request_secret` - Enable/disable validating the RADIUS request shared secret in the Start or End record. Valid values: `disable`, `enable`.

* `secondary_secret` - Secret key to access the secondary server.
* `secondary_server` - {&lt;name_str|ip_str&gt;} secondary RADIUS CN domain name or IP.
* `secret` - Pre-shared secret key used to access the primary RADIUS server.
* `server` - Primary RADIUS server CN domain name or IP address.
* `server_identity_check` - Enable/disable RADIUS server identity check (verify server domain name/IP address against the server certificate). Valid values: `disable`, `enable`.

* `source_ip` - Source IP address for communications to the RADIUS server.
* `source_ip_interface` - Source interface for communication with the RADIUS server.
* `sso_attribute` - RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record. Valid values: `User-Name`, `User-Password`, `CHAP-Password`, `NAS-IP-Address`, `NAS-Port`, `Service-Type`, `Framed-Protocol`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Framed-Routing`, `Filter-Id`, `Framed-MTU`, `Framed-Compression`, `Login-IP-Host`, `Login-Service`, `Login-TCP-Port`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `State`, `Class`, `Session-Timeout`, `Idle-Timeout`, `Termination-Action`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Link`, `Framed-AppleTalk-Network`, `Framed-AppleTalk-Zone`, `Acct-Status-Type`, `Acct-Delay-Time`, `Acct-Input-Octets`, `Acct-Output-Octets`, `Acct-Session-Id`, `Acct-Authentic`, `Acct-Session-Time`, `Acct-Input-Packets`, `Acct-Output-Packets`, `Acct-Terminate-Cause`, `Acct-Multi-Session-Id`, `Acct-Link-Count`, `CHAP-Challenge`, `NAS-Port-Type`, `Port-Limit`, `Login-LAT-Port`.

* `sso_attribute_key` - Key prefix for SSO group value in the SSO attribute.
* `sso_attribute_value_override` - Enable/disable override old attribute value with new value for the same endpoint. Valid values: `disable`, `enable`.

* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `switch_controller_acct_fast_framedip_detect` - Switch-Controller-Acct-Fast-Framedip-Detect.
* `switch_controller_nas_ip_dynamic` - Enable/Disable switch-controller nas-ip dynamic to dynamically set nas-ip. Valid values: `disable`, `enable`.

* `switch_controller_service_type` - Switch-Controller-Service-Type. Valid values: `login`, `framed`, `callback-login`, `callback-framed`, `outbound`, `administrative`, `nas-prompt`, `authenticate-only`, `callback-nas-prompt`, `call-check`, `callback-administrative`.

* `tertiary_secret` - Secret key to access the tertiary server.
* `tertiary_server` - {&lt;name_str|ip_str&gt;} tertiary RADIUS CN domain name or IP.
* `timeout` - Time in seconds between re-sending authentication requests.
* `tls_min_proto_version` - Minimum supported protocol version for TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`.

* `transport_protocol` - Transport protocol to be used (default = udp). Valid values: `udp`, `tcp`, `tls`.

* `use_group_for_profile` - Use-Group-For-Profile. Valid values: `disable`, `enable`.

* `use_management_vdom` - Enable/disable using management VDOM to send requests. Valid values: `disable`, `enable`.

* `username_case_sensitive` - Enable/disable case sensitive user names. Valid values: `disable`, `enable`.


The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `accounting_server` block supports:

* `id` - ID (0 - 4294967295).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `port` - RADIUS accounting port number.
* `secret` - Secret key.
* `server` - {&lt;name_str|ip_str&gt;} Server CN domain name or IP.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Radius can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_radius.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
