---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_hpe"
description: |-
  Host protection engine configuration.
---

# fortimanager_object_system_npu_hpe
Host protection engine configuration.

~> This resource is a sub resource for variable `hpe` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_hpe" "trname" {
  all_protocol = "400000"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `all_protocol` - Maximum packet rate of each host queue except high priority traffic(1K - 40M pps, default = 10M pps), set 0 to disable.
* `arp_max` - Maximum ARP packet rate (1K - 40M pps, default = 40K pps).
* `enable_queue_shaper` - Enable/Disable NPU host protection engine (HPE) queue shaper. Valid values: `disable`, `enable`.

* `enable_shaper` - Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper. Valid values: `disable`, `enable`.

* `esp_max` - Maximum ESP packet rate (1K - 40M pps, default = 40K pps).
* `exception_code` - Maximum exception code rate of traffic(1K - 32M pps, default = 1M pps).
* `fragment_with_sess` - Maximum fragment with session rate of traffic(1K - 32M pps, default = 1M pps).
* `fragment_without_session` - Maximum fragment without session rate of traffic(1K - 32M pps, default = 1M pps).
* `high_priority` - Maximum packet rate for TCAM high priority traffic (1K - 40M pps, default = 10M pps),set 0 to disable.
* `icmp_max` - Maximum ICMP packet rate (1K - 40M pps, default = 40K pps).
* `ip_frag_max` - Maximum fragmented IP packet rate (1K - 40M pps, default = 40K pps).
* `ip_others_max` - Maximum IP packet rate for other packets (packet types that cannot be set with other options) (1K - 1G pps, default = 40K pps).
* `l2_others_max` - Maximum L2 packet rate for L2 packets that are not ARP packets (1K - 40M pps, default = 40K pps).
* `pri_type_max` - Maximum overflow rate of priority type traffic(1K - 40M pps, default = 40K pps). Includes L2: HA, 802.3ad LACP, heartbeats. L3: OSPF. L4_TCP: BGP. L4_UDP: IKE, SLBC, BFD.
* `queue_shaper_max` - Maximum per queue byte rate of traffic(1K - 32M pps, default = 1M pps).
* `sctp_max` - Maximum SCTP packet rate (1K - 40M pps, default = 40K pps).
* `tcp_max` - Maximum TCP packet rate (1K - 40M pps, default = 600K pps).
* `tcpfin_rst_max` - Maximum TCP carries FIN or RST flags packet rate (1K - 40M pps, default = 600K pps).
* `tcpsyn_ack_max` - Maximum TCP carries SYN and ACK flags packet rate (1K - 40M pps, default = 600K pps).
* `tcpsyn_max` - Maximum TCP SYN packet rate (1K - 40M pps, default = 600K pps).
* `udp_max` - Maximum UDP packet rate (1K - 40M pps, default = 600K pps).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuHpe can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_hpe.labelname ObjectSystemNpuHpe
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
