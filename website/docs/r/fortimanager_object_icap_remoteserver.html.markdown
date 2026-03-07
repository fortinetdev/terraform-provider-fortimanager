---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_remoteserver"
description: |-
  ObjectIcap RemoteServer
---

# fortimanager_object_icap_remoteserver
ObjectIcap RemoteServer

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `addr_type` - Addr-Type. Valid values: `fqdn`, `ip4`, `ip6`.

* `fqdn` - Fqdn.
* `healthcheck` - Healthcheck. Valid values: `disable`, `enable`.

* `healthcheck_service` - Healthcheck-Service.
* `ip_address` - Ip-Address.
* `ip6_address` - Ip6-Address.
* `max_connections` - Max-Connections.
* `name` - Name.
* `port` - Port.
* `secure` - Secure. Valid values: `disable`, `enable`.

* `ssl_cert` - Ssl-Cert.
* `validate_server_certificate` - Validate-Server-Certificate. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIcap RemoteServer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_remoteserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
