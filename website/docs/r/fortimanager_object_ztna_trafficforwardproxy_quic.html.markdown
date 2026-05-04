---
subcategory: "Object ZTNA"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ztna_trafficforwardproxy_quic"
description: |-
  ObjectZtna TrafficForwardProxyQuic
---

# fortimanager_object_ztna_trafficforwardproxy_quic
ObjectZtna TrafficForwardProxyQuic

~> This resource is a sub resource for variable `quic` of resource `fortimanager_object_ztna_trafficforwardproxy`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `traffic_forward_proxy` - Traffic Forward Proxy.

* `ack_delay_exponent` - Ack-Delay-Exponent.
* `active_connection_id_limit` - Active-Connection-Id-Limit.
* `active_migration` - Active-Migration. Valid values: `disable`, `enable`.

* `grease_quic_bit` - Grease-Quic-Bit. Valid values: `disable`, `enable`.

* `max_ack_delay` - Max-Ack-Delay.
* `max_datagram_frame_size` - Max-Datagram-Frame-Size.
* `max_idle_timeout` - Max-Idle-Timeout.
* `max_udp_payload_size` - Max-Udp-Payload-Size.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectZtna TrafficForwardProxyQuic can be imported using any of these accepted formats:
```
Set import_options = ["traffic_forward_proxy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ztna_trafficforwardproxy_quic.labelname ObjectZtnaTrafficForwardProxyQuic
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
