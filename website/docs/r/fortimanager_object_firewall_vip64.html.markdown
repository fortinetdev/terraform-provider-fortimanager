---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vip64"
description: |-
  Configure IPv6 to IPv4 virtual IPs.
---

# fortimanager_object_firewall_vip64
Configure IPv6 to IPv4 virtual IPs.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`dynamic_mapping`: `fortimanager_object_firewall_vip64_dynamic_mapping`
`realservers`: `fortimanager_object_firewall_vip64_realservers`



## Example Usage

```hcl
resource "fortimanager_object_firewall_vip64" "trname" {
  arp_reply   = "enable"
  color       = 2
  comment     = "This is a Terraform example"
  extip       = "2001:192:168:1::1"
  extport     = "0"
  fosid       = 1
  ldb_method  = "static"
  mappedport  = "0"
  name        = "terr-firewall-vip64"
  portforward = "disable"
  protocol    = "tcp"
  server_type = "http"
  type        = "server-load-balance"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `arp_reply` - Enable ARP reply. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `extip` - Start-external-IP [-end-external-IP].
* `extport` - External service port.
* `fosid` - Custom defined id.
* `ldb_method` - Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.

* `mappedip` - Start-mapped-IP [-end-mapped-IP].
* `mappedport` - Mapped service port.
* `monitor` - Health monitors.
* `name` - VIP64 name.
* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.

* `protocol` - Mapped port protocol. Valid values: `tcp`, `udp`.

* `realservers` - Realservers. The structure of `realservers` block is documented below.
* `server_type` - Server type. Valid values: `http`, `tcp`, `udp`, `ip`.

* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x).
* `type` - VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `arp_reply` - Enable ARP reply. Valid values: `disable`, `enable`.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `extip` - Start-external-IP [-end-external-IP].
* `extport` - External service port.
* `id` - Custom defined id.
* `ldb_method` - Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.

* `mappedip` - Start-mapped-IP [-end-mapped-IP].
* `mappedport` - Mapped service port.
* `monitor` - Health monitors.
* `portforward` - Enable port forwarding. Valid values: `disable`, `enable`.

* `protocol` - Mapped port protocol. Valid values: `tcp`, `udp`.

* `server_type` - Server type. Valid values: `http`, `tcp`, `udp`, `ip`.

* `src_filter` - Source IP6 filter (x:x:x:x:x:x:x:x/x).
* `type` - VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `realservers` block supports:

* `client_ip` - Restrict server to a client IP in this range.
* `healthcheck` - Per server health check. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Hold down interval.
* `id` - Real server ID.
* `ip` - Mapped server IP.
* `max_connections` - Maximum number of connections allowed to server.
* `monitor` - Health monitors.
* `port` - Mapped server port.
* `status` - Server administrative status. Valid values: `active`, `standby`, `disable`.

* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall Vip64 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vip64.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
