---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_sslsshprofile"
description: |-
  Configure SSL/SSH protocol options.
---

# fortimanager_object_firewall_sslsshprofile
Configure SSL/SSH protocol options.

## Example Usage

```hcl
resource "fortimanager_object_firewall_sslsshprofile" "trname" {
  comment         = "terraform-comment1"
  mapi_over_https = "disable"
  name            = "terraform-tefv"
  use_ssl_server  = "disable"
  whitelist       = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `allowlist` - Enable/disable exempting servers by FortiGuard allowlist. Valid values: `disable`, `enable`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `block_blocklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `block_blacklisted_certificates` - Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
 (`ver Controlled FortiOS <= 6.4`)
* `caname` - CA certificate used by SSL Inspection.
* `comment` - Optional comments.
* `dot` - Dot. The structure of `dot` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `ftps` - Ftps. The structure of `ftps` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `https` - Https. The structure of `https` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `imaps` - Imaps. The structure of `imaps` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `mapi_over_https` - Enable/disable inspection of MAPI over HTTPS. Valid values: `disable`, `enable`.

* `name` - Name.
* `pop3s` - Pop3S. The structure of `pop3s` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `rpc_over_https` - Enable/disable inspection of RPC over HTTPS. Valid values: `disable`, `enable`.

* `server_cert` - Certificate used by SSL Inspection to replace server certificate.
* `server_cert_mode` - Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.

* `smtps` - Smtps. The structure of `smtps` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `ssh` - Ssh. The structure of `ssh` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `ssl` - Ssl. The structure of `ssl` block is documented below. (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`)
* `ssl_anomalies_log` - Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.

* `ssl_exempt` - Ssl-Exempt. The structure of `ssl_exempt` block is documented below.
* `ssl_exemptions_log` - Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.

* `ssl_negotiation_log` - Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
 (`ver Controlled FortiOS >= 6.4`)
* `ssl_server` - Ssl-Server. The structure of `ssl_server` block is documented below.
* `supported_alpn` - Configure ALPN option. Valid values: `none`, `http1-1`, `http2`, `all`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `untrusted_caname` - Untrusted CA certificate used by SSL Inspection.
* `use_ssl_server` - Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.

* `whitelist` - Enable/disable exempting servers by FortiGuard whitelist. Valid values: `disable`, `enable`.
 (`ver Controlled FortiOS <= 6.4`)
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dot` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`):

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `block`, `allow`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `block`, `allow`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `ftps` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `https` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `cert_probe_failure` - Action based on certificate probe failure. Valid values: `block`, `allow`.
 (`ver FortiManager >= 7.0 and Controlled FortiOS >= 7.0`)
* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `certificate-inspection`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `imaps` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `pop3s` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `smtps` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `ssh` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `inspect_all` - Level of SSL inspection. Valid values: `disable`, `deep-inspection`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `ssh_algorithm` - Relative strength of encryption algorithms accepted during negotiation. Valid values: `compatible`, `high-encryption`.

* `ssh_tun_policy_check` - Enable/disable SSH tunnel policy check. Valid values: `disable`, `enable`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_version` - Action based on SSH version being unsupported. Valid values: `block`, `bypass`.


The `ssl` block supports (`ver FortiManager >= 7.0 and Controlled FortiOS >= 6.4`):

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `inspect_all` - Level of SSL inspection. Valid values: `disable`, `certificate-inspection`, `deep-inspection`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.


The `ssl_exempt` block supports:

* `address` - IPv4 address object.
* `address6` - IPv6 address object.
* `fortiguard_category` - FortiGuard category ID.
* `id` - ID number.
* `regex` - Exempt servers by regular expression.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category. Valid values: `fortiguard-category`, `address`, `address6`, `wildcard-fqdn`, `regex`.

* `wildcard_fqdn` - Exempt servers by wildcard FQDN.

The `ssl_server` block supports:

* `ftps_client_certificate` - Action based on received client certificate during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver Controlled FortiOS >= 6.4`)
* `https_client_certificate` - Action based on received client certificate during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver Controlled FortiOS >= 6.4`)
* `ftps_client_cert_request` - Action based on client certificate request during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `https_client_cert_request` - Action based on client certificate request during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `id` - SSL server ID.
* `imaps_client_certificate` - Action based on received client certificate during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver Controlled FortiOS >= 6.4`)
* `imaps_client_cert_request` - Action based on client certificate request during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `ip` - IPv4 address of the SSL server.
* `pop3s_client_certificate` - Action based on received client certificate during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver Controlled FortiOS >= 6.4`)
* `smtps_client_certificate` - Action based on received client certificate during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver Controlled FortiOS >= 6.4`)
* `ssl_other_client_certificate` - Action based on received client certificate during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver Controlled FortiOS >= 6.4`)
* `pop3s_client_cert_request` - Action based on client certificate request during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `smtps_client_cert_request` - Action based on client certificate request during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)
* `ssl_other_client_cert_request` - Action based on client certificate request during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.
 (`ver FortiManager <= 6.4 and Controlled FortiOS <= 6.2`)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall SslSshProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_sslsshprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
