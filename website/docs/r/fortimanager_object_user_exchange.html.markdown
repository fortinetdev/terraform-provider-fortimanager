---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_exchange"
description: |-
  Configure MS Exchange server entries.
---

# fortimanager_object_user_exchange
Configure MS Exchange server entries.

## Example Usage

```hcl
resource "fortimanager_object_user_exchange" "labelname" {
  auth_level            = "privacy"
  auth_type             = "ntlm"
  auto_discover_kdc     = "enable"
  connect_protocol      = "rpc-over-https"
  domain_name           = "sien.com"
  http_auth_type        = "basic"
  ip                    = "22.22.2.2"
  name                  = "ss2"
  password              = ["fdafd"]
  ssl_min_proto_version = "default"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `addr_type` - Addr-Type. Valid values: `ipv4`, `ipv6`.

* `auth_level` - Authentication security level used for the RPC protocol layer. Valid values: `low`, `medium`, `normal`, `high`, `connect`, `call`, `packet`, `integrity`, `privacy`.

* `auth_type` - Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.

* `auto_discover_kdc` - Enable/disable automatic discovery of KDC IP addresses. Valid values: `disable`, `enable`.

* `connect_protocol` - Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.

* `domain_name` - MS Exchange server fully qualified domain name.
* `http_auth_type` - Authentication security type used for the HTTP transport. Valid values: `ntlm`, `basic`.

* `ip` - Server IPv4 address.
* `ip6` - Ip6.
* `kdc_ip` - KDC IPv4 addresses for Kerberos authentication.
* `name` - MS Exchange server entry name.
* `password` - Password for the specified username.
* `server_name` - MS Exchange server hostname.
* `ssl_min_proto_version` - Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`.

* `username` - User name used to sign in to the server. Must have proper permissions for service.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Exchange can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_exchange.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
