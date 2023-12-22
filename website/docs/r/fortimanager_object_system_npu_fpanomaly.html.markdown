---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_fpanomaly"
description: |-
  NP6Lite anomaly protection (packet drop or send trap to host).
---

# fortimanager_object_system_npu_fpanomaly
NP6Lite anomaly protection (packet drop or send trap to host).

~> This resource is a sub resource for variable `fp_anomaly` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_fpanomaly" "trname" {
  icmp_csum_err  = "drop"
  esp_minlen_err = "drop"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `capwap_minlen_err` - Invalid IPv4 capwap min length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `esp_minlen_err` - Invalid IPv4 ESP short packet anomalies. Valid values: `drop`, `trap-to-host`.

* `gre_csum_err` - Invalid IPv4 GRE checksum anomalies. Valid values: `drop`, `trap-to-host`.

* `gtpu_plen_err` - Invalid IPv4 gtpu packet length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `icmp_csum_err` - Invalid IPv4 ICMP packet checksum anomalies. Valid values: `drop`, `trap-to-host`.

* `icmp_frag` - Layer 3 fragmented packets that could be part of layer 4 ICMP anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `icmp_land` - ICMP land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `icmp_minlen_err` - Invalid IPv4 ICMP short packet anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv4_csum_err` - Invalid IPv4 packet checksum anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv4_ihl_err` - Invalid IPv4 header length anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv4_land` - Land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_len_err` - Invalid IPv4 packet length anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv4_opt_err` - Invalid IPv4 option parsing anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv4_optlsrr` - Loose source record route option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_optrr` - Record route option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_optsecurity` - Security option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_optssrr` - Strict source record route option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_optstream` - Stream option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_opttimestamp` - Timestamp option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_proto_err` - Invalid layer 4 protocol anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_ttlzero_err` - Invalid IPv4 TTL field zero anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv4_unknopt` - Unknown option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv4_ver_err` - Invalid IPv4 header version anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv6_daddr_err` - Destination address as unspecified or loopback address anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_exthdr_len_err` - Invalid IPv6 packet chain extension header total length anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv6_exthdr_order_err` - Invalid IPv6 packet extension header ordering anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv6_ihl_err` - Invalid IPv6 packet length anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv6_land` - Land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_optendpid` - End point identification anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_opthomeaddr` - Home address option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_optinvld` - Invalid option anomalies.Invalid option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_optjumbo` - Jumbo options anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_optnsap` - Network service access point address option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_optralert` - Router alert option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_opttunnel` - Tunnel encapsulation limit option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_plen_zero` - Invalid IPv6 packet payload length zero anomalies. Valid values: `drop`, `trap-to-host`.

* `ipv6_proto_err` - Layer 4 invalid protocol anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_saddr_err` - Source address as multicast anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_unknopt` - Unknown option anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `ipv6_ver_err` - Invalid IPv6 packet version anomalies. Valid values: `drop`, `trap-to-host`.

* `nvgre_minlen_err` - Invalid IPv4 nvgre min length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `sctp_clen_err` - Invalid IPv4 SCTP length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `sctp_crc_err` - Invalid IPv4 SCTP CRC error anomalies. Valid values: `drop`, `trap-to-host`.

* `sctp_l4len_err` - Invalid IPv4 SCTP L4 packet length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `tcp_csum_err` - Invalid IPv4 TCP packet checksum anomalies. Valid values: `drop`, `trap-to-host`.

* `tcp_fin_noack` - TCP SYN flood with FIN flag set without ACK setting anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `tcp_fin_only` - TCP SYN flood with only FIN flag set anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `tcp_hlen_err` - Invalid IPv4 TCP header length anomalies. Valid values: `drop`, `trap-to-host`.

* `tcp_hlenvsl4len_err` - Invalid IPv4 tcp header vs packet length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `tcp_land` - TCP land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `tcp_no_flag` - TCP SYN flood with no flag set anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `tcp_plen_err` - Invalid IPv4 TCP packet length anomalies. Valid values: `drop`, `trap-to-host`.

* `tcp_syn_data` - TCP SYN flood packets with data anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `tcp_syn_fin` - TCP SYN flood SYN/FIN flag set anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `tcp_winnuke` - TCP WinNuke anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `udp_csum_err` - Invalid IPv4 UDP packet checksum anomalies. Valid values: `drop`, `trap-to-host`.

* `udp_hlen_err` - Invalid IPv4 UDP packet header length anomalies. Valid values: `drop`, `trap-to-host`.

* `udp_land` - UDP land anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

* `udp_len_err` - Invalid IPv4 UDP packet length anomalies. Valid values: `drop`, `trap-to-host`.

* `udp_plen_err` - Invalid IPv4 UDP packet minimum length anomalies. Valid values: `drop`, `trap-to-host`.

* `udplite_cover_err` - Invalid IPv4 UDP-Lite packet coverage anomalies. Valid values: `drop`, `trap-to-host`.

* `udplite_csum_err` - Invalid IPv4 UDP-Lite packet checksum anomalies. Valid values: `drop`, `trap-to-host`.

* `uesp_minlen_err` - Invalid IPv4 UESP min length check error anomalies. Valid values: `drop`, `trap-to-host`.

* `unknproto_minlen_err` - Invalid IPv4 L4 unknown protocol short packet anomalies. Valid values: `drop`, `trap-to-host`.

* `vxlan_minlen_err` - Invalid IPv4 vxlan min length check error anomalies. Valid values: `drop`, `trap-to-host`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuFpAnomaly can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_fpanomaly.labelname ObjectSystemNpuFpAnomaly
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
