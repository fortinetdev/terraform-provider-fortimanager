---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vip6_quic"
description: |-
  QUIC setting.
---

# fortimanager_object_firewall_vip6_quic
QUIC setting.

~> This resource is a sub resource for variable `quic` of resource `fortimanager_object_firewall_vip6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vip6` - Vip6.

* `ack_delay_exponent` - <i>Support meta variable</i> ACK delay exponent (1 - 20, default = 3).
* `active_connection_id_limit` - <i>Support meta variable</i> Active connection ID limit (1 - 8, default = 2).
* `active_migration` - Enable/disable active migration (default = disable). Valid values: `disable`, `enable`.

* `grease_quic_bit` - Enable/disable grease QUIC bit (default = enable). Valid values: `disable`, `enable`.

* `max_ack_delay` - <i>Support meta variable</i> Maximum ACK delay in milliseconds (1 - 16383, default = 25).
* `max_datagram_frame_size` - <i>Support meta variable</i> Maximum datagram frame size in bytes (1 - 1500, default = 1500).
* `max_idle_timeout` - <i>Support meta variable</i> Maximum idle timeout milliseconds (1 - 60000, default = 30000).
* `max_udp_payload_size` - <i>Support meta variable</i> Maximum UDP payload size in bytes (1200 - 1500, default = 1500).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectFirewall Vip6Quic can be imported using any of these accepted formats:
```
Set import_options = ["vip6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vip6_quic.labelname ObjectFirewallVip6Quic
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
