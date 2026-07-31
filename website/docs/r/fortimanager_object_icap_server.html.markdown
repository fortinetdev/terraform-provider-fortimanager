---
subcategory: "Object ICAP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_server"
description: |-
  Configure ICAP servers.
---

# fortimanager_object_icap_server
Configure ICAP servers.

## Example Usage

```hcl
resource "fortimanager_object_icap_server" "trname" {
  ip_address      = "192.168.1.1"
  ip_version      = "4"
  max_connections = 100
  name            = "terr-icap-server"
  port            = 1344
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `addr_type` - Address type of the remote ICAP server: IPv4, IPv6 or FQDN. Valid values: `fqdn`, `ip4`, `ip6`.

* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `fqdn` - ICAP remote server Fully Qualified Domain Name (FQDN).
* `healthcheck` - Enable/disable ICAP remote server health checking. Attempts to connect to the remote ICAP server to verify that the server is operating normally. Valid values: `disable`, `enable`.

* `healthcheck_service` - ICAP Service name to use for health checks.
* `ip_address` - IPv4 address of the ICAP server.
* `ip_version` - IP version. Valid values: `4`, `6`.

* `ip6_address` - IPv6 address of the ICAP server.
* `max_connections` - Maximum number of concurrent connections to ICAP server. Must not be less than wad-worker-count.
* `name` - Server name.
* `port` - ICAP server port.
* `secure` - Enable/disable secure connection to ICAP server. Valid values: `disable`, `enable`.

* `ssl_cert` - CA certificate name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIcap Server can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_server.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
