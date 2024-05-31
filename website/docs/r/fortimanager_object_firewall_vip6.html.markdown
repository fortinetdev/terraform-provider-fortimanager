---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vip6"
description: |-
  Configure virtual IP for IPv6.
---

# fortimanager_object_firewall_vip6
Configure virtual IP for IPv6.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dynamic_mapping`: `fortimanager_object_firewall_vip6_dynamic_mapping`
>- `quic`: `fortimanager_object_firewall_vip6_quic`
>- `realservers`: `fortimanager_object_firewall_vip6_realservers`
>- `ssl_cipher_suites`: `fortimanager_object_firewall_vip6_sslciphersuites`
>- `ssl_server_cipher_suites`: `fortimanager_object_firewall_vip6_sslserverciphersuites`



## Example Usage

```hcl
resource "fortimanager_object_firewall_vip6" "trname" {
  arp_reply                     = "enable"
  color                         = 1
  comment                       = "This is a Terraform example"
  extip                         = "2001:192:168:1::2"
  extport                       = "0"
  http_cookie_age               = 60
  http_cookie_domain_from_host  = "disable"
  http_cookie_share             = "same-ip"
  http_ip_header                = "disable"
  http_multiplex                = "disable"
  http_redirect                 = "disable"
  https_cookie_secure           = "disable"
  ldb_method                    = "static"
  mappedip                      = ["2001:192:168:1::2"]
  mappedport                    = "0"
  max_embryonic_connections     = 1000
  name                          = "terr-firewall-vip6"
  nat_source_vip                = "disable"
  outlook_web_access            = "disable"
  persistence                   = "none"
  portforward                   = "disable"
  protocol                      = "tcp"
  ssl_client_fallback           = "enable"
  ssl_hsts                      = "disable"
  ssl_hsts_age                  = 5184000
  ssl_hsts_include_subdomains   = "disable"
  ssl_http_location_conversion  = "disable"
  ssl_server_algorithm          = "client"
  ssl_server_max_version        = "client"
  ssl_server_min_version        = "client"
  ssl_server_session_state_type = "both"
  type                          = "static-nat"
  weblogic_server               = "disable"
  websphere_server              = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.

* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `embedded_ipv4_address` - Enable/disable embedded IPv4 address. Valid values: `disable`, `enable`.

* `extip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `h2_support` - Enable/disable HTTP2 support (default = enable). Valid values: `disable`, `enable`.

* `h3_support` - Enable/disable HTTP3/QUIC support (default = disable). Valid values: `disable`, `enable`.

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 seconds. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.

* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across virtual servers. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.

* `http_ip_header` - For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header. Valid values: `disable`, `enable`.

* `http_ip_header_name` - For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
* `http_multiplex` - Enable/disable HTTP multiplexing. Valid values: `disable`, `enable`.

* `http_redirect` - Enable/disable redirection of HTTP to HTTPS Valid values: `disable`, `enable`.

* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.

* `fosid` - Custom defined ID.
* `ipv4_mappedip` - Start-mapped-IPv4-address [-end mapped-IPv4-address].
* `ipv4_mappedport` - IPv4 port number range on the destination network to which the external port number range is mapped.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`, `http-host`.

* `mappedip` - Mapped IP address range in the format startIP-endIP.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `name` - Virtual ip6 name.
* `nat_source_vip` - Enable to perform SNAT on traffic from mappedip to the extip for all egress interfaces. Valid values: `disable`, `enable`.

* `nat64` - Enable/disable DNAT64. Valid values: `disable`, `enable`.

* `nat66` - Enable/disable DNAT66. Valid values: `disable`, `enable`.

* `ndp_reply` - Enable/disable this FortiGate unit's ability to respond to NDP requests for this virtual IP address (default = enable). Valid values: `disable`, `enable`.

* `outlook_web_access` - Enable to add the Front-End-Https header for Microsoft Outlook Web Access. Valid values: `disable`, `enable`.

* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`, `ssl-session-id`.

* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.

* `protocol` - Protocol to use when forwarding packets. Valid values: `tcp`, `udp`, `sctp`.

* `quic` - Quic. The structure of `quic` block is documented below.
* `realservers` - Realservers. The structure of `realservers` block is documented below.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP). Valid values: `http`, `https`, `ssl`, `tcp`, `udp`, `ip`, `imaps`, `pop3s`, `smtps`.

* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x). Separate addresses with spaces.
* `src_vip_filter` - Enable/disable use of 'src-filter' to match destinations for the reverse SNAT rule. Valid values: `disable`, `enable`.

* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength. Valid values: `high`, `low`, `medium`, `custom`.

* `ssl_certificate` - The name of the SSL certificate to use for SSL acceleration.
* `ssl_cipher_suites` - Ssl-Cipher-Suites. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507). Valid values: `disable`, `enable`.

* `ssl_client_rekey_count` - Maximum length of data in MB before triggering a client rekey (0 = disable).
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746. Valid values: `deny`, `allow`, `secure`.

* `ssl_client_session_state_max` - Maximum number of client to FortiGate SSL session states to keep.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiGate SSL session state.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.

* `ssl_hpkp` - Enable/disable including HPKP header in response. Valid values: `disable`, `enable`, `report-only`.

* `ssl_hpkp_age` - Number of minutes the web browser should keep HPKP.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.

* `ssl_hsts_age` - Number of seconds the client should honour the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field. Valid values: `disable`, `enable`.

* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion. Valid values: `disable`, `enable`.

* `ssl_max_version` - Highest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full). Valid values: `half`, `full`.

* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions. Valid values: `require`, `deny`, `allow`.

* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems. Valid values: `disable`, `enable`.

* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `low`, `medium`, `custom`, `client`.

* `ssl_server_cipher_suites` - Ssl-Server-Cipher-Suites. The structure of `ssl_server_cipher_suites` block is documented below.
* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `disable`, `enable`.

* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `type` - Configure a static NAT or server load balance VIP. Valid values: `static-nat`, `server-load-balance`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `weblogic_server` - Enable to add an HTTP header to indicate SSL offloading for a WebLogic server. Valid values: `disable`, `enable`.

* `websphere_server` - Enable to add an HTTP header to indicate SSL offloading for a WebSphere server. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.

* `arp_reply` - Enable to respond to ARP requests for this virtual IP address. Enabled by default. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `embedded_ipv4_address` - Enable/disable embedded IPv4 address. Valid values: `disable`, `enable`.

* `extip` - IP address or address range on the external interface that you want to map to an address or address range on the destination network.
* `extport` - Incoming port number range that you want to map to a port number range on the destination network.
* `h2_support` - Enable/disable HTTP2 support (default = enable). Valid values: `disable`, `enable`.

* `h3_support` - Enable/disable HTTP3/QUIC support (default = disable). Valid values: `disable`, `enable`.

* `http_cookie_age` - Time in minutes that client web browsers should keep a cookie. Default is 60 seconds. 0 = no time limit.
* `http_cookie_domain` - Domain that HTTP cookie persistence should apply to.
* `http_cookie_domain_from_host` - Enable/disable use of HTTP cookie domain from host field in HTTP. Valid values: `disable`, `enable`.

* `http_cookie_generation` - Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.
* `http_cookie_path` - Limit HTTP cookie persistence to the specified path.
* `http_cookie_share` - Control sharing of cookies across virtual servers. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing. Valid values: `disable`, `same-ip`.

* `http_ip_header` - For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header. Valid values: `disable`, `enable`.

* `http_ip_header_name` - For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.
* `http_multiplex` - Enable/disable HTTP multiplexing. Valid values: `disable`, `enable`.

* `http_redirect` - Enable/disable redirection of HTTP to HTTPS Valid values: `disable`, `enable`.

* `https_cookie_secure` - Enable/disable verification that inserted HTTPS cookies are secure. Valid values: `disable`, `enable`.

* `id` - Custom defined ID.
* `ipv4_mappedip` - Start-mapped-IPv4-address [-end mapped-IPv4-address].
* `ipv4_mappedport` - IPv4 port number range on the destination network to which the external port number range is mapped.
* `ldb_method` - Method used to distribute sessions to real servers. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`, `http-host`.

* `mappedip` - Mapped IP address range in the format startIP-endIP.
* `mappedport` - Port number range on the destination network to which the external port number range is mapped.
* `max_embryonic_connections` - Maximum number of incomplete connections.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `nat_source_vip` - Nat-Source-Vip. Valid values: `disable`, `enable`.

* `nat64` - Enable/disable DNAT64. Valid values: `disable`, `enable`.

* `nat66` - Enable/disable DNAT66. Valid values: `disable`, `enable`.

* `ndp_reply` - Enable/disable this FortiGate unit's ability to respond to NDP requests for this virtual IP address (default = enable). Valid values: `disable`, `enable`.

* `outlook_web_access` - Enable to add the Front-End-Https header for Microsoft Outlook Web Access. Valid values: `disable`, `enable`.

* `persistence` - Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session. Valid values: `none`, `http-cookie`, `ssl-session-id`.

* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.

* `protocol` - Protocol to use when forwarding packets. Valid values: `tcp`, `udp`, `sctp`.

* `realservers` - Realservers. The structure of `realservers` block is documented below.
* `server_type` - Protocol to be load balanced by the virtual server (also called the server load balance virtual IP). Valid values: `http`, `https`, `ssl`, `tcp`, `udp`, `ip`, `imaps`, `pop3s`, `smtps`.

* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x). Separate addresses with spaces.
* `src_vip_filter` - Enable/disable use of 'src-filter' to match destinations for the reverse SNAT rule. Valid values: `disable`, `enable`.

* `ssl_accept_ffdhe_groups` - Enable/disable FFDHE cipher suite for SSL key exchange. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Permitted encryption algorithms for SSL sessions according to encryption strength. Valid values: `high`, `low`, `medium`, `custom`.

* `ssl_certificate` - The name of the SSL certificate to use for SSL acceleration.
* `ssl_cipher_suites` - Ssl-Cipher-Suites. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_client_fallback` - Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507). Valid values: `disable`, `enable`.

* `ssl_client_rekey_count` - Maximum length of data in MB before triggering a client rekey (0 = disable).
* `ssl_client_renegotiation` - Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746. Valid values: `deny`, `allow`, `secure`.

* `ssl_client_session_state_max` - Maximum number of client to FortiGate SSL session states to keep.
* `ssl_client_session_state_timeout` - Number of minutes to keep client to FortiGate SSL session state.
* `ssl_client_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `ssl_dh_bits` - Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.

* `ssl_hpkp` - Enable/disable including HPKP header in response. Valid values: `disable`, `enable`, `report-only`.

* `ssl_hpkp_age` - Number of minutes the web browser should keep HPKP.
* `ssl_hpkp_backup` - Certificate to generate backup HPKP pin from.
* `ssl_hpkp_include_subdomains` - Indicate that HPKP header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_hpkp_primary` - Certificate to generate primary HPKP pin from.
* `ssl_hpkp_report_uri` - URL to report HPKP violations to.
* `ssl_hsts` - Enable/disable including HSTS header in response. Valid values: `disable`, `enable`.

* `ssl_hsts_age` - Number of seconds the client should honour the HSTS setting.
* `ssl_hsts_include_subdomains` - Indicate that HSTS header applies to all subdomains. Valid values: `disable`, `enable`.

* `ssl_http_location_conversion` - Enable to replace HTTP with HTTPS in the reply's Location HTTP header field. Valid values: `disable`, `enable`.

* `ssl_http_match_host` - Enable/disable HTTP host matching for location conversion. Valid values: `disable`, `enable`.

* `ssl_max_version` - Highest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Lowest SSL/TLS version acceptable from a client. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full). Valid values: `half`, `full`.

* `ssl_pfs` - Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions. Valid values: `require`, `deny`, `allow`.

* `ssl_send_empty_frags` - Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems. Valid values: `disable`, `enable`.

* `ssl_server_algorithm` - Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength. Valid values: `high`, `low`, `medium`, `custom`, `client`.

* `ssl_server_max_version` - Highest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_min_version` - Lowest SSL/TLS version acceptable from a server. Use the client setting by default. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_renegotiation` - Enable/disable secure renegotiation to comply with RFC 5746. Valid values: `disable`, `enable`.

* `ssl_server_session_state_max` - Maximum number of FortiGate to Server SSL session states to keep.
* `ssl_server_session_state_timeout` - Number of minutes to keep FortiGate to Server SSL session state.
* `ssl_server_session_state_type` - How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate. Valid values: `disable`, `time`, `count`, `both`.

* `type` - Configure a static NAT or server load balance VIP. Valid values: `static-nat`, `server-load-balance`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `weblogic_server` - Enable to add an HTTP header to indicate SSL offloading for a WebLogic server. Valid values: `disable`, `enable`.

* `websphere_server` - Enable to add an HTTP header to indicate SSL offloading for a WebSphere server. Valid values: `disable`, `enable`.


The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `realservers` block supports:

* `client_ip` - Only clients in this IP range can connect to this real server.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Time in seconds that the health check monitor continues to monitor an unresponsive server that should be active.
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IP address of the real server.
* `max_connections` - Max number of active connections that can directed to the real server. When reached, sessions are sent to other real servers.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.

The `ssl_cipher_suites` block supports:

* `cipher` - Cipher suite name. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`.

* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


The `quic` block supports:

* `ack_delay_exponent` - <i>Support meta variable</i> ACK delay exponent (1 - 20, default = 3).
* `active_connection_id_limit` - <i>Support meta variable</i> Active connection ID limit (1 - 8, default = 2).
* `active_migration` - Enable/disable active migration (default = disable). Valid values: `disable`, `enable`.

* `grease_quic_bit` - Enable/disable grease QUIC bit (default = enable). Valid values: `disable`, `enable`.

* `max_ack_delay` - <i>Support meta variable</i> Maximum ACK delay in milliseconds (1 - 16383, default = 25).
* `max_datagram_frame_size` - <i>Support meta variable</i> Maximum datagram frame size in bytes (1 - 1500, default = 1500).
* `max_idle_timeout` - <i>Support meta variable</i> Maximum idle timeout milliseconds (1 - 60000, default = 30000).
* `max_udp_payload_size` - <i>Support meta variable</i> Maximum UDP payload size in bytes (1200 - 1500, default = 1500).

The `realservers` block supports:

* `client_ip` - Only clients in this IP range can connect to this real server.
* `healthcheck` - Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Time in seconds that the health check monitor continues to monitor an unresponsive server that should be active.
* `http_host` - HTTP server domain name in HTTP header.
* `id` - Real server ID.
* `ip` - IPv6 address of the real server.
* `max_connections` - Max number of active connections that can directed to the real server. When reached, sessions are sent to other real servers.
* `monitor` - Name of the health check monitor to use when polling to determine a virtual server's connectivity status.
* `port` - Port for communicating with the real server. Required if port forwarding is enabled.
* `status` - Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.

* `translate_host` - Enable/disable translation of hostname/IP from virtual server to real server. Valid values: `disable`, `enable`.

* `weight` - Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.

The `ssl_cipher_suites` block supports:

* `cipher` - Cipher suite name. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`.

* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


The `ssl_server_cipher_suites` block supports:

* `cipher` - Cipher suite name. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`.

* `priority` - SSL/TLS cipher suites priority.
* `versions` - SSL/TLS versions that the cipher suite can be used with. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Vip6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vip6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
