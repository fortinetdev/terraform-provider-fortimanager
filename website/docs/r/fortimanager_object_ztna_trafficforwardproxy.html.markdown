---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_trafficforwardproxy"
description: |-
  Configure ZTNA traffic forward proxy.
---

# fortimanager_object_ztna_trafficforwardproxy
Configure ZTNA traffic forward proxy.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `quic`: `fortimanager_object_ztna_trafficforwardproxy_quic`
>- `ssl_cipher_suites`: `fortimanager_object_ztna_trafficforwardproxy_sslciphersuites`
>- `ssl_server_cipher_suites`: `fortimanager_object_ztna_trafficforwardproxy_sslserverciphersuites`
>- `url_route`: `fortimanager_object_ztna_trafficforwardproxy_urlroute`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `auth_portal` - Enable/disable authentication portal. Valid values: `disable`, `enable`.

* `auth_virtual_host` - Virtual host for authentication portal.
* `client_cert` - Client-Cert. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `empty_cert_action` - Empty-Cert-Action. Valid values: `accept`, `block`, `accept-unmanageable`.

* `h3_support` - H3-Support. Valid values: `disable`, `enable`.

* `host` - Virtual or real host name.
* `interface` - Interface.
* `log_blocked_traffic` - Enable/disable logging of blocked traffic. Valid values: `disable`, `enable`.

* `name` - ZTNA proxy name.
* `port` - Port.
* `quic` - Quic. The structure of `quic` block is documented below.
* `ssl_accept_ffdhe_groups` - Ssl-Accept-Ffdhe-Groups. Valid values: `disable`, `enable`.

* `ssl_algorithm` - Ssl-Algorithm. Valid values: `high`, `low`, `medium`, `custom`.

* `ssl_certificate` - Ssl-Certificate.
* `ssl_cipher_suites` - Ssl-Cipher-Suites. The structure of `ssl_cipher_suites` block is documented below.
* `ssl_client_fallback` - Ssl-Client-Fallback. Valid values: `disable`, `enable`.

* `ssl_client_rekey_count` - Ssl-Client-Rekey-Count.
* `ssl_client_renegotiation` - Ssl-Client-Renegotiation. Valid values: `allow`, `deny`, `secure`.

* `ssl_client_session_state_max` - Ssl-Client-Session-State-Max.
* `ssl_client_session_state_timeout` - Ssl-Client-Session-State-Timeout.
* `ssl_client_session_state_type` - Ssl-Client-Session-State-Type. Valid values: `disable`, `time`, `count`, `both`.

* `ssl_dh_bits` - Ssl-Dh-Bits. Valid values: `768`, `1024`, `1536`, `2048`, `3072`, `4096`.

* `ssl_hpkp` - Ssl-Hpkp. Valid values: `disable`, `enable`, `report-only`.

* `ssl_hpkp_age` - Ssl-Hpkp-Age.
* `ssl_hpkp_backup` - Ssl-Hpkp-Backup.
* `ssl_hpkp_include_subdomains` - Ssl-Hpkp-Include-Subdomains. Valid values: `disable`, `enable`.

* `ssl_hpkp_primary` - Ssl-Hpkp-Primary.
* `ssl_hpkp_report_uri` - Ssl-Hpkp-Report-Uri.
* `ssl_hsts` - Ssl-Hsts. Valid values: `disable`, `enable`.

* `ssl_hsts_age` - Ssl-Hsts-Age.
* `ssl_hsts_include_subdomains` - Ssl-Hsts-Include-Subdomains. Valid values: `disable`, `enable`.

* `ssl_http_location_conversion` - Ssl-Http-Location-Conversion. Valid values: `disable`, `enable`.

* `ssl_http_match_host` - Ssl-Http-Match-Host. Valid values: `disable`, `enable`.

* `ssl_max_version` - Ssl-Max-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_min_version` - Ssl-Min-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `ssl_mode` - Ssl-Mode. Valid values: `half`, `full`.

* `ssl_pfs` - Ssl-Pfs. Valid values: `require`, `deny`, `allow`.

* `ssl_send_empty_frags` - Ssl-Send-Empty-Frags. Valid values: `disable`, `enable`.

* `ssl_server_algorithm` - Ssl-Server-Algorithm. Valid values: `high`, `low`, `medium`, `custom`, `client`.

* `ssl_server_cipher_suites` - Ssl-Server-Cipher-Suites. The structure of `ssl_server_cipher_suites` block is documented below.
* `ssl_server_max_version` - Ssl-Server-Max-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_min_version` - Ssl-Server-Min-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `client`, `tls-1.3`.

* `ssl_server_renegotiation` - Ssl-Server-Renegotiation. Valid values: `disable`, `enable`.

* `ssl_server_session_state_max` - Ssl-Server-Session-State-Max.
* `ssl_server_session_state_timeout` - Ssl-Server-Session-State-Timeout.
* `ssl_server_session_state_type` - Ssl-Server-Session-State-Type. Valid values: `disable`, `time`, `count`, `both`.

* `status` - Status. Valid values: `disable`, `enable`.

* `svr_pool_multiplex` - Svr-Pool-Multiplex. Valid values: `disable`, `enable`.

* `svr_pool_server_max_concurrent_request` - Svr-Pool-Server-Max-Concurrent-Request.
* `svr_pool_server_max_request` - Svr-Pool-Server-Max-Request.
* `svr_pool_ttl` - Svr-Pool-Ttl.
* `user_agent_detect` - User-Agent-Detect. Valid values: `disable`, `enable`.

* `vip` - Virtual IP name.
* `vip6` - Virtual IPv6 name.
* `url_route` - Url-Route. The structure of `url_route` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `quic` block supports:

* `ack_delay_exponent` - Ack-Delay-Exponent.
* `active_connection_id_limit` - Active-Connection-Id-Limit.
* `active_migration` - Active-Migration. Valid values: `disable`, `enable`.

* `grease_quic_bit` - Grease-Quic-Bit. Valid values: `disable`, `enable`.

* `max_ack_delay` - Max-Ack-Delay.
* `max_datagram_frame_size` - Max-Datagram-Frame-Size.
* `max_idle_timeout` - Max-Idle-Timeout.
* `max_udp_payload_size` - Max-Udp-Payload-Size.

The `ssl_cipher_suites` block supports:

* `cipher` - Cipher. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`.

* `priority` - Priority.
* `versions` - Versions. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


The `ssl_server_cipher_suites` block supports:

* `cipher` - Cipher. Valid values: `TLS-RSA-WITH-RC4-128-MD5`, `TLS-RSA-WITH-RC4-128-SHA`, `TLS-RSA-WITH-DES-CBC-SHA`, `TLS-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA`, `TLS-RSA-WITH-AES-256-CBC-SHA`, `TLS-RSA-WITH-AES-128-CBC-SHA256`, `TLS-RSA-WITH-AES-256-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-RSA-WITH-SEED-CBC-SHA`, `TLS-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-RSA-WITH-DES-CBC-SHA`, `TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-DHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-AES-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-RSA-WITH-SEED-CBC-SHA`, `TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-RC4-128-SHA`, `TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA`, `TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256`, `TLS-DHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-DHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA`, `TLS-DHE-DSS-WITH-AES-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-128-GCM-SHA256`, `TLS-DHE-DSS-WITH-AES-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA`, `TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384`, `TLS-RSA-WITH-AES-128-GCM-SHA256`, `TLS-RSA-WITH-AES-256-GCM-SHA384`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA`, `TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256`, `TLS-DHE-DSS-WITH-SEED-CBC-SHA`, `TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256`, `TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384`, `TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256`, `TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384`, `TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA`, `TLS-DHE-DSS-WITH-DES-CBC-SHA`, `TLS-AES-128-GCM-SHA256`, `TLS-AES-256-GCM-SHA384`, `TLS-CHACHA20-POLY1305-SHA256`, `TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA`.

* `priority` - Priority.
* `versions` - Versions. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.


The `url_route` block supports:

* `name` - Name.
* `service_connector` - Service-Connector.
* `url_pattern` - Url-Pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectZtna TrafficForwardProxy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_trafficforwardproxy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
