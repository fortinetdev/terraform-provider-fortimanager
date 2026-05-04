---
subcategory: "Object VPN"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_kmipserver"
description: |-
  KMIP server entry configuration.
---

# fortimanager_object_vpn_kmipserver
KMIP server entry configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_list`: `fortimanager_object_vpn_kmipserver_serverlist`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `name` - KMIP server entry name.
* `password` - Password to use for connectivity to the KMIP server.
* `server_identity_check` - Enable/disable KMIP server identity check (verify server FQDN/IP address against the server certificate). Valid values: `disable`, `enable`.

* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `source_ip` - FortiGate IP address to be used for communication with the KMIP server.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `username` - User name to use for connectivity to the KMIP server.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `cert` - Client certificate to use for connectivity to the KMIP server.
* `id` - ID
* `port` - KMIP server port.
* `server` - KMIP server FQDN or IP address.
* `status` - Enable/disable KMIP server. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn KmipServer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_kmipserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
