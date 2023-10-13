---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_sslsshprofile_imaps"
description: |-
  Configure IMAPS options.
---

# fortimanager_object_firewall_sslsshprofile_imaps
Configure IMAPS options.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `allow_invalid_server_cert` - When enabled, allows SSL sessions whose server certificate validation failed. Valid values: `disable`, `enable`.

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `min_allowed_ssl_version` - Min-Allowed-Ssl-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `client_cert_request` - Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.

* `invalid_server_cert` - Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `disable`, `enable`, `strict`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.

* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `block`, `allow`, `inspect`.

* `unsupported_ssl` - Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.

* `untrusted_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.

* `untrusted_server_cert` - Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall SslSshProfileImaps can be imported using any of these accepted formats:
```
Set import_options = ["ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_sslsshprofile_imaps.labelname ObjectFirewallSslSshProfileImaps
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
