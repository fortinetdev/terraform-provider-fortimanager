---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_sslsshprofile_dot"
description: |-
  Configure DNS over TLS options.
---

# fortimanager_object_firewall_sslsshprofile_dot
Configure DNS over TLS options.

~> This resource is a sub resource for variable `dot` of resource `fortimanager_object_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_firewall_sslsshprofile_dot" "trname" {
  ssl_ssh_profile         = fortimanager_object_firewall_sslsshprofile.trname.name
  cert_validation_failure = "ignore"
  cert_validation_timeout = "ignore"
  client_certificate      = "bypass"
  depends_on              = [fortimanager_object_firewall_sslsshprofile.trname]
}

resource "fortimanager_object_firewall_sslsshprofile" "trname" {
  name = "terr-sslsshprofile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `cert_validation_failure` - Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.

* `cert_validation_timeout` - Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.

* `client_certificate` - Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.

* `expired_server_cert` - Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.

* `min_allowed_ssl_version` - Min-Allowed-Ssl-Version. Valid values: `ssl-3.0`, `tls-1.0`, `tls-1.1`, `tls-1.2`, `tls-1.3`.

* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `quic` - Enable/disable QUIC inspection (default = disable). Valid values: `disable`, `enable`.

* `revoked_server_cert` - Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.

* `sni_server_cert_check` - Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_ssl_cipher` - Action based on the SSL cipher used being unsupported. Valid values: `block`, `allow`.

* `unsupported_ssl_negotiation` - Action based on the SSL negotiation used being unsupported. Valid values: `block`, `allow`.

* `unsupported_ssl_version` - Action based on the SSL version used being unsupported. Valid values: `block`, `allow`, `inspect`.

* `untrusted_server_cert` - Action based on server certificate is not issued by a trusted CA. Valid values: `allow`, `block`, `ignore`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall SslSshProfileDot can be imported using any of these accepted formats:
```
Set import_options = ["ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_sslsshprofile_dot.labelname ObjectFirewallSslSshProfileDot
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
