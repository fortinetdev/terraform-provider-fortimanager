---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vip64_realservers"
description: |-
  Real servers.
---

# fortimanager_object_firewall_vip64_realservers
Real servers.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vip64` - Vip64.

* `client_ip` - Restrict server to a client IP in this range.
* `healthcheck` - Per server health check. Valid values: `disable`, `enable`, `vip`.

* `holddown_interval` - Hold down interval.
* `fosid` - Real server ID.
* `ip` - Mapped server IP.
* `max_connections` - Maximum number of connections allowed to server.
* `monitor` - Health monitors.
* `port` - Mapped server port.
* `status` - Server administrative status. Valid values: `active`, `standby`, `disable`.

* `weight` - Weight.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall Vip64Realservers can be imported using any of these accepted formats:
```
Set import_options = ["vip64=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vip64_realservers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
