---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu"
description: |-
  Configure NPU attributes.
---

# fortimanager_object_system_npu
Configure NPU attributes.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `background_sse_scan` - Background-Sse-Scan. The structure of `background_sse_scan` block is documented below.
* `capwap_offload` - Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `disable`, `enable`.

* `dedicated_management_affinity` - Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `dedicated_management_cpu` - Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `disable`, `enable`.

* `default_qos_type` - Set default QoS type. Valid values: `policing`, `shaping`.

* `dos_options` - Dos-Options. The structure of `dos_options` block is documented below.
* `double_level_mcast_offload` - Enable double level mcast offload. Valid values: `disable`, `enable`.

* `dse_timeout` - DSE timeout in seconds (0-3600, default = 10).
* `dsw_dts_profile` - Dsw-Dts-Profile. The structure of `dsw_dts_profile` block is documented below.
* `dsw_queue_dts_profile` - Dsw-Queue-Dts-Profile. The structure of `dsw_queue_dts_profile` block is documented below.
* `fastpath` - Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.

* `fp_anomaly` - Fp-Anomaly. The structure of `fp_anomaly` block is documented below.
* `gtp_enhanced_cpu_range` - GTP enhanced CPU range option. Valid values: `0`, `1`, `2`.

* `gtp_enhanced_mode` - Enable/disable GTP enhanced mode. Valid values: `disable`, `enable`.

* `gtp_support` - Enable/Disable NP7 GTP support Valid values: `disable`, `enable`.

* `hash_config` - Configure NPU trunk hash. Valid values: `5-tuple`, `src-ip`, `src-dst-ip`.

* `hash_ipv6_sel` - Select which 4bytes of the IPv6 address are used for traffic hash(0~3).
* `hash_tbl_spread` - Enable/disable hash table entry spread (default enabled). Valid values: `disable`, `enable`.

* `host_shortcut_mode` - Set np6 host shortcut mode. Valid values: `bi-directional`, `host-shortcut`.

* `hpe` - Hpe. The structure of `hpe` block is documented below.
* `htab_dedi_queue_nr` - Set the number of dedicate queue for hash table messages.
* `htab_msg_queue` - Set hash table message queue mode. Valid values: `idle`, `data`, `dedicated`.

* `htx_gtse_quota` - Configure HTX GTSE quota. Valid values: `100Mbps`, `200Mbps`, `300Mbps`, `400Mbps`, `500Mbps`, `600Mbps`, `700Mbps`, `800Mbps`, `900Mbps`, `1Gbps`, `2Gbps`, `4Gbps`, `8Gbps`, `10Gbps`.

* `htx_icmp_csum_chk` - Set HTX icmp csum checking mode. Valid values: `pass`, `drop`.

* `hw_ha_scan_interval` - HW HA periodical scan interval in seconds (0-3600, default = 120, 0 to disable).
* `inbound_dscp_copy` - Enable/disable copying the DSCP field from outer IP header to inner IP Header. Valid values: `disable`, `enable`.

* `inbound_dscp_copy_port` - Physical interfaces that support inbound-dscp-copy.
* `ip_reassembly` - Ip-Reassembly. The structure of `ip_reassembly` block is documented below.
* `intf_shaping_offload` - Enable/disable NPU offload when doing interface-based traffic shaping according to the egress-shaping-profile. Valid values: `disable`, `enable`.

* `ip_fragment_offload` - Enable/disable NP7 NPU IP fragment offload. Valid values: `disable`, `enable`.

* `iph_rsvd_re_cksum` - Enable/disable IP checksum re-calculation for packets with iph.reserved bit set. Valid values: `disable`, `enable`.

* `ippool_overload_high` - High threshold for overload ippool port reuse (100%-2000%, default = 200).
* `ippool_overload_low` - Low threshold for overload ippool port reuse (100%-2000%, default = 150).
* `ipsec_dec_subengine_mask` - IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_enc_subengine_mask` - IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_host_dfclr` - Enable/disable DF clearing of NP4lite host IPsec offload. Valid values: `disable`, `enable`.

* `ipsec_inbound_cache` - Enable/disable IPsec inbound cache for anti-replay. Valid values: `disable`, `enable`.

* `ipsec_local_uesp_port` - Ipsec-Local-Uesp-Port.
* `ipsec_mtu_override` - Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.

* `ipsec_ob_np_sel` - IPsec NP selection for OB SA offloading. Valid values: `RR`, `Packet`, `Hash`.

* `ipsec_over_vlink` - Enable/disable IPSEC over vlink. Valid values: `disable`, `enable`.

* `isf_np_queues` - Isf-Np-Queues. The structure of `isf_np_queues` block is documented below.
* `isf_np_rx_tr_distr` - Select ISF NP Rx trunk distribution (PSC) mode. Valid values: `port-flow`, `round-robin`, `randomized`.

* `lag_out_port_select` - Enable/disable LAG outgoing port selection based on incoming traffic port. Valid values: `disable`, `enable`.

* `max_session_timeout` - Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).
* `mcast_session_accounting` - Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `disable`, `session-based`, `tpe-based`.

* `mcast_session_counting` - Mcast-Session-Counting. Valid values: `disable`, `enable`, `session-based`, `tpe-based`.

* `mcast_session_counting6` - Enable/disable traffic accounting for each multicast session6 through TAE counter. Valid values: `disable`, `enable`, `session-based`, `tpe-based`.

* `napi_break_interval` - NAPI break interval (default 0).
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `disable`, `enable`.

* `np_queues` - Np-Queues. The structure of `np_queues` block is documented below.
* `np6_cps_optimization_mode` - Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `disable`, `enable`.

* `pba_eim` - Configure option for PBA(non-overload)/EIM combination. Valid values: `disallow`, `allow`.

* `per_policy_accounting` - Set per-policy accounting. Valid values: `disable`, `enable`.

* `per_session_accounting` - Enable/disable per-session accounting. Valid values: `enable`, `disable`, `enable-by-log`, `all-enable`, `traffic-log-only`.

* `ple_non_syn_tcp_action` - Configure action for the PLE to take on TCP packets that have the SYN field unset. Valid values: `forward`, `drop`.

* `policy_offload_level` - Firewall Policy Offload Level(DISABLE/DOS/FULL). Valid values: `disable`, `dos-offload`, `full-offload`.

* `port_cpu_map` - Port-Cpu-Map. The structure of `port_cpu_map` block is documented below.
* `port_npu_map` - Port-Npu-Map. The structure of `port_npu_map` block is documented below.
* `port_path_option` - Port-Path-Option. The structure of `port_path_option` block is documented below.
* `priority_protocol` - Priority-Protocol. The structure of `priority_protocol` block is documented below.
* `process_icmp_by_host` - Enable/disable process ICMP by host when received from IPsec tunnel and payload size &lt; 119. Valid values: `disable`, `enable`.

* `prp_port_in` - Ingress port configured to allow the PRP trailer not be stripped off when the PRP packets come in.                    All of the traffic originating from these ports will always be sent to the host.
* `prp_port_out` - Egress port configured to allow the PRP trailer not be stripped off when the PRP packets go out.
* `qos_mode` - QoS mode on switch and NP. Valid values: `disable`, `priority`, `round-robin`.

* `qtm_buf_mode` - QTM channel configuration for packet buffer. Valid values: `6ch`, `4ch`.

* `rdp_offload` - Enable/disable rdp offload. Valid values: `disable`, `enable`.

* `recover_np6_link` - Enable/disable internal link failure check and recovery after boot up. Valid values: `disable`, `enable`.

* `rps_mode` - Enable/disable receive packet steering (RPS) optimization mode. Valid values: `disable`, `enable`.

* `session_acct_interval` - Session accounting update interval (1 - 10 sec, default 5 sec).
* `session_denied_offload` - Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.

* `sse_backpressure` - Enable/disable sse backpressure. Valid values: `disable`, `enable`.

* `sse_ha_scan` - Sse-Ha-Scan. The structure of `sse_ha_scan` block is documented below.
* `strip_clear_text_padding` - Enable/disable stripping clear text padding. Valid values: `disable`, `enable`.

* `strip_esp_padding` - Enable/disable stripping ESP padding. Valid values: `disable`, `enable`.

* `sw_eh_hash` - Sw-Eh-Hash. The structure of `sw_eh_hash` block is documented below.
* `sw_np_bandwidth` - Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`, `7G`, `8G`, `9G`.

* `switch_np_hash` - Switch-NP trunk port selection Criteria. Valid values: `src-ip`, `dst-ip`, `src-dst-ip`.

* `tcp_rst_timeout` - TCP RST timeout in seconds (0-3600, default = 5).
* `tcp_timeout_profile` - Tcp-Timeout-Profile. The structure of `tcp_timeout_profile` block is documented below.
* `udp_timeout_profile` - Udp-Timeout-Profile. The structure of `udp_timeout_profile` block is documented below.
* `uesp_offload` - Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `disable`, `enable`.

* `ull_port_mode` - Set ULL port's speed to 10G/25G (default 10G). Valid values: `10G`, `25G`.

* `vlan_lookup_cache` - Enable/disable vlan lookup cache (default enabled). Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `background_sse_scan` block supports:

* `scan` - Enable/disable background SSE scan by driver thread(default enabled). Valid values: `disable`, `enable`.

* `stats_update_interval` - Stats update interval(&gt;=5*60 seconds, default 5*60 seconds).
* `udp_keepalive_interval` - UDP keepalive interval(&gt;=90 seconds, default 90 seconds).

The `dos_options` block supports:

* `npu_dos_meter_mode` - Set DoS meter npu offloading mode. Valid values: `local`, `global`.

* `npu_dos_synproxy_mode` - Set NPU DoS SYNPROXY mode. Valid values: `synack2ack`, `pass-synack`.

* `npu_dos_tpe_mode` - Enable/Disable inserting DoS meter id to session table. Valid values: `disable`, `enable`.


The `dsw_dts_profile` block supports:

* `action` - Set NPU DSW DTS profile action. Valid values: `wait`, `drop`, `drop_tmr_0`, `drop_tmr_1`, `enque`, `enque_0`, `enque_1`.

* `min_limit` - Set NPU DSW DTS profile min-limt.
* `profile_id` - Set NPU DSW DTS profile profile id.
* `step` - Set NPU DSW DTS profile step.

The `dsw_queue_dts_profile` block supports:

* `iport` - Set NPU DSW DTS in port. Valid values: `EIF0`, `EIF1`, `EIF2`, `EIF3`, `EIF4`, `EIF5`, `EIF6`, `EIF7`, `HTX0`, `HTX1`, `SSE0`, `SSE1`, `SSE2`, `SSE3`, `RLT`, `DFR`, `IPSECI`, `IPSECO`, `IPTI`, `IPTO`, `VEP0`, `VEP2`, `VEP4`, `VEP6`, `IVS`, `L2TI1`, `L2TO`, `L2TI0`, `PLE`, `SPATH`, `QTM`.

* `name` - Name.
* `oport` - Set NPU DSW DTS out port. Valid values: `EIF0`, `EIF1`, `EIF2`, `EIF3`, `EIF4`, `EIF5`, `EIF6`, `EIF7`, `HRX`, `SSE0`, `SSE1`, `SSE2`, `SSE3`, `RLT`, `DFR`, `IPSECI`, `IPSECO`, `IPTI`, `IPTO`, `VEP0`, `VEP2`, `VEP4`, `VEP6`, `IVS`, `L2TI1`, `L2TO`, `L2TI0`, `PLE`, `SYNK`, `NSS`, `TSK`, `QTM`.

* `profile_id` - Set NPU DSW DTS profile id.
* `queue_select` - Set NPU DSW DTS queue id select(0 - reset to default).

The `fp_anomaly` block supports:

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


The `hpe` block supports:

* `all_protocol` - Maximum packet rate of each host queue except high priority traffic(1K - 40M pps, default = 10M pps), set 0 to disable.
* `arp_max` - Maximum ARP packet rate (1K - 40M pps, default = 40K pps).
* `enable_shaper` - Enable/Disable NPU Host Protection Engine (HPE) for packet type shaper. Valid values: `disable`, `enable`.

* `esp_max` - Maximum ESP packet rate (1K - 40M pps, default = 40K pps).
* `high_priority` - Maximum packet rate for TCAM high priority traffic (1K - 40M pps, default = 10M pps),set 0 to disable.
* `icmp_max` - Maximum ICMP packet rate (1K - 40M pps, default = 40K pps).
* `ip_frag_max` - Maximum fragmented IP packet rate (1K - 40M pps, default = 40K pps).
* `ip_others_max` - Maximum IP packet rate for other packets (packet types that cannot be set with other options) (1K - 1G pps, default = 40K pps).
* `l2_others_max` - Maximum L2 packet rate for L2 packets that are not ARP packets (1K - 40M pps, default = 40K pps).
* `pri_type_max` - Maximum overflow rate of priority type traffic(1K - 40M pps, default = 40K pps). Includes L2: HA, 802.3ad LACP, heartbeats. L3: OSPF. L4_TCP: BGP. L4_UDP: IKE, SLBC, BFD.
* `sctp_max` - Maximum SCTP packet rate (1K - 40M pps, default = 40K pps).
* `tcp_max` - Maximum TCP packet rate (1K - 40M pps, default = 600K pps).
* `tcpfin_rst_max` - Maximum TCP carries FIN or RST flags packet rate (1K - 40M pps, default = 600K pps).
* `tcpsyn_ack_max` - Maximum TCP carries SYN and ACK flags packet rate (1K - 40M pps, default = 600K pps).
* `tcpsyn_max` - Maximum TCP SYN packet rate (1K - 40M pps, default = 600K pps).
* `udp_max` - Maximum UDP packet rate (1K - 40M pps, default = 600K pps).

The `ip_reassembly` block supports:

* `max_timeout` - Maximum timeout value for IP reassembly (5 us - 600,000,000 us).
* `min_timeout` - Minimum timeout value for IP reassembly (5 us - 600,000,000 us).
* `status` - Set IP reassembly processing status. Valid values: `disable`, `enable`.


The `isf_np_queues` block supports:

* `cos0` - CoS profile name for CoS 0.
* `cos1` - CoS profile name for CoS 1.
* `cos2` - CoS profile name for CoS 2.
* `cos3` - CoS profile name for CoS 3.
* `cos4` - CoS profile name for CoS 4.
* `cos5` - CoS profile name for CoS 5.
* `cos6` - CoS profile name for CoS 6.
* `cos7` - CoS profile name for CoS 7.

The `np_queues` block supports:

* `ethernet_type` - Ethernet-Type. The structure of `ethernet_type` block is documented below.
* `ip_protocol` - Ip-Protocol. The structure of `ip_protocol` block is documented below.
* `ip_service` - Ip-Service. The structure of `ip_service` block is documented below.
* `profile` - Profile. The structure of `profile` block is documented below.
* `scheduler` - Scheduler. The structure of `scheduler` block is documented below.

The `ethernet_type` block supports:

* `name` - Ethernet Type Name.
* `queue` - Queue Number.
* `type` - Ethernet Type.
* `weight` - Class Weight.

The `ip_protocol` block supports:

* `name` - IP Protocol Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `weight` - Class Weight.

The `ip_service` block supports:

* `dport` - Destination Port.
* `name` - IP Service Name.
* `protocol` - IP Protocol.
* `queue` - Queue Number.
* `sport` - Source Port.
* `weight` - Class Weight.

The `profile` block supports:

* `cos0` - Queue number of CoS 0. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos1` - Queue number of CoS 1. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos2` - Queue number of CoS 2. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos3` - Queue number of CoS 3. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos4` - Queue number of CoS 4. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos5` - Queue number of CoS 5. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos6` - Queue number of CoS 6. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `cos7` - Queue number of CoS 7. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp0` - Queue number of DSCP 0. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp1` - Queue number of DSCP 1. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp10` - Queue number of DSCP 10. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp11` - Queue number of DSCP 11. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp12` - Queue number of DSCP 12. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp13` - Queue number of DSCP 13. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp14` - Queue number of DSCP 14. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp15` - Queue number of DSCP 15. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp16` - Queue number of DSCP 16. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp17` - Queue number of DSCP 17. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp18` - Queue number of DSCP 18. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp19` - Queue number of DSCP 19. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp2` - Queue number of DSCP 2. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp20` - Queue number of DSCP 20. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp21` - Queue number of DSCP 21. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp22` - Queue number of DSCP 22. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp23` - Queue number of DSCP 23. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp24` - Queue number of DSCP 24. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp25` - Queue number of DSCP 25. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp26` - Queue number of DSCP 26. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp27` - Queue number of DSCP 27. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp28` - Queue number of DSCP 28. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp29` - Queue number of DSCP 29. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp3` - Queue number of DSCP 3. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp30` - Queue number of DSCP 30. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp31` - Queue number of DSCP 31. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp32` - Queue number of DSCP 32. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp33` - Queue number of DSCP 33. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp34` - Queue number of DSCP 34. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp35` - Queue number of DSCP 35. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp36` - Queue number of DSCP 36. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp37` - Queue number of DSCP 37. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp38` - Queue number of DSCP 38. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp39` - Queue number of DSCP 39. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp4` - Queue number of DSCP 4. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp40` - Queue number of DSCP 40. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp41` - Queue number of DSCP 41. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp42` - Queue number of DSCP 42. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp43` - Queue number of DSCP 43. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp44` - Queue number of DSCP 44. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp45` - Queue number of DSCP 45. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp46` - Queue number of DSCP 46. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp47` - Queue number of DSCP 47. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp48` - Queue number of DSCP 48. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp49` - Queue number of DSCP 49. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp5` - Queue number of DSCP 5. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp50` - Queue number of DSCP 50. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp51` - Queue number of DSCP 51. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp52` - Queue number of DSCP 52. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp53` - Queue number of DSCP 53. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp54` - Queue number of DSCP 54. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp55` - Queue number of DSCP 55. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp56` - Queue number of DSCP 56. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp57` - Queue number of DSCP 57. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp58` - Queue number of DSCP 58. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp59` - Queue number of DSCP 59. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp6` - Queue number of DSCP 6. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp60` - Queue number of DSCP 60. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp61` - Queue number of DSCP 61. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp62` - Queue number of DSCP 62. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp63` - Queue number of DSCP 63. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp7` - Queue number of DSCP 7. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp8` - Queue number of DSCP 8. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `dscp9` - Queue number of DSCP 9. Valid values: `queue0`, `queue1`, `queue2`, `queue3`, `queue4`, `queue5`, `queue6`, `queue7`.

* `id` - Profile ID.
* `type` - Profile type. Valid values: `cos`, `dscp`.

* `weight` - Class weight.

The `scheduler` block supports:

* `mode` - Scheduler Mode. Valid values: `none`, `priority`, `round-robin`.

* `name` - Scheduler Name.

The `port_cpu_map` block supports:

* `cpu_core` - The CPU core to map to an interface.
* `interface` - The interface to map to a CPU core.

The `port_npu_map` block supports:

* `interface` - Set npu interface port to NPU group map.
* `npu_group_index` - Mapping NPU group index.

The `port_path_option` block supports:

* `ports_using_npu` - Set ha/aux ports to handle traffic with NPU (otherise traffic goes to Intel-NIC and then CPU).

The `priority_protocol` block supports:

* `bfd` - Enable/disable NPU BFD priority protocol. Valid values: `disable`, `enable`.

* `bgp` - Enable/disable NPU BGP priority protocol. Valid values: `disable`, `enable`.

* `slbc` - Enable/disable NPU SLBC priority protocol. Valid values: `disable`, `enable`.


The `sse_ha_scan` block supports:

* `gap` - Scanning message gap(0~32767, default 6000)
* `max_session_cnt` - If the session count(in millions) is larger than this, HA scan will be skipped. (0~0xffff, default 0xffff)
* `min_duration` - Scanning filter for minimum duration of the session. (0~3600, default 0)

The `sw_eh_hash` block supports:

* `computation` - Set hashing computation. Valid values: `xor16`, `xor8`, `xor4`, `crc16`.

* `destination_ip_lower_16` - Include/exclude destination IP address lower 16 bits. Valid values: `include`, `exclude`.

* `destination_ip_upper_16` - Include/exclude destination IP address upper 16 bits. Valid values: `include`, `exclude`.

* `destination_port` - Include/exclude destination port if TCP/UDP. Valid values: `include`, `exclude`.

* `ip_protocol` - Include/exclude IP protocol. Valid values: `include`, `exclude`.

* `netmask_length` - Network mask length.
* `source_ip_lower_16` - Include/exclude source IP address lower 16 bits. Valid values: `include`, `exclude`.

* `source_ip_upper_16` - Include/exclude source IP address upper 16 bits. Valid values: `include`, `exclude`.

* `source_port` - Include/exclude source port if TCP/UDP. Valid values: `include`, `exclude`.


The `tcp_timeout_profile` block supports:

* `close_wait` - Set close-wait timeout(seconds)
* `fin_wait` - Set fin-wait timeout(seconds)
* `id` - Timeout profile ID (5-47)
* `syn_sent` - Set syn-sent timeout(seconds)
* `syn_wait` - Set syn-wait timeout(seconds)
* `tcp_idle` - Set TCP establish timeout(seconds)
* `time_wait` - Set time-wait timeout(seconds)

The `udp_timeout_profile` block supports:

* `id` - Timeout profile ID (5-63)
* `udp_idle` - Set UDP idle timeout(seconds)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem Npu can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu.labelname ObjectSystemNpu
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
