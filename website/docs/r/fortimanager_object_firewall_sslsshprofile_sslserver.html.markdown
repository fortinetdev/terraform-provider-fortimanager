---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_sslsshprofile_sslserver"
description: |-
  SSL servers.
---

# fortimanager_object_firewall_sslsshprofile_sslserver
SSL servers.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `ftps_client_certificate` - Action based on received client certificate during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `https_client_certificate` - Action based on received client certificate during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `ftps_client_cert_request` - Action based on client certificate request during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `https_client_cert_request` - Action based on client certificate request during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `fosid` - SSL server ID.
* `imaps_client_certificate` - Action based on received client certificate during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `imaps_client_cert_request` - Action based on client certificate request during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `ip` - IPv4 address of the SSL server.
* `pop3s_client_certificate` - Action based on received client certificate during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.

* `smtps_client_certificate` - Action based on received client certificate during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `ssl_other_client_certificate` - Action based on received client certificate during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.

* `pop3s_client_cert_request` - Action based on client certificate request during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.

* `smtps_client_cert_request` - Action based on client certificate request during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `ssl_other_client_cert_request` - Action based on client certificate request during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall SslSshProfileSslServer can be imported using any of these accepted formats:
```
Set import_options = ["ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_sslsshprofile_sslserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
