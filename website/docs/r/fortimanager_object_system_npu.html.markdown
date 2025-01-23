---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu"
description: |-
  Configure NPU attributes.
---

# fortimanager_object_system_npu
Configure NPU attributes.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `background_sse_scan`: `fortimanager_object_system_npu_backgroundssescan`
>- `dos_options`: `fortimanager_object_system_npu_dosoptions`
>- `dsw_dts_profile`: `fortimanager_object_system_npu_dswdtsprofile`
>- `dsw_queue_dts_profile`: `fortimanager_object_system_npu_dswqueuedtsprofile`
>- `fp_anomaly`: `fortimanager_object_system_npu_fpanomaly`
>- `hpe`: `fortimanager_object_system_npu_hpe`
>- `icmp_error_rate_ctrl`: `fortimanager_object_system_npu_icmperrorratectrl`
>- `icmp_rate_ctrl`: `fortimanager_object_system_npu_icmpratectrl`
>- `ip_reassembly`: `fortimanager_object_system_npu_ipreassembly`
>- `isf_np_queues`: `fortimanager_object_system_npu_isfnpqueues`
>- `np_queues`: `fortimanager_object_system_npu_npqueues`
>- `npu_tcam`: `fortimanager_object_system_npu_nputcam`
>- `port_cpu_map`: `fortimanager_object_system_npu_portcpumap`
>- `port_npu_map`: `fortimanager_object_system_npu_portnpumap`
>- `port_path_option`: `fortimanager_object_system_npu_portpathoption`
>- `priority_protocol`: `fortimanager_object_system_npu_priorityprotocol`
>- `sse_ha_scan`: `fortimanager_object_system_npu_ssehascan`
>- `sw_eh_hash`: `fortimanager_object_system_npu_swehhash`
>- `sw_tr_hash`: `fortimanager_object_system_npu_swtrhash`
>- `tcp_timeout_profile`: `fortimanager_object_system_npu_tcptimeoutprofile`
>- `udp_timeout_profile`: `fortimanager_object_system_npu_udptimeoutprofile`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `background_sse_scan` - Background-Sse-Scan. The structure of `background_sse_scan` block is documented below.
* `capwap_offload` - Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `disable`, `enable`.

* `dedicated_lacp_queue` - Enable to dedicate one HIF queue for LACP. Valid values: `disable`, `enable`.

* `dedicated_management_affinity` - Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `dedicated_management_cpu` - Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `disable`, `enable`.

* `default_qos_type` - Set default QoS type. Valid values: `policing`, `shaping`.

* `default_tcp_refresh_dir` - Default SSE timeout TCP refresh direction. Valid values: `both`, `outgoing`, `incoming`.

* `default_udp_refresh_dir` - Default SSE timeout UDP refresh direction. Valid values: `both`, `outgoing`, `incoming`.

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
* `icmp_error_rate_ctrl` - Icmp-Error-Rate-Ctrl. The structure of `icmp_error_rate_ctrl` block is documented below.
* `icmp_rate_ctrl` - Icmp-Rate-Ctrl. The structure of `icmp_rate_ctrl` block is documented below.
* `inbound_dscp_copy` - Enable/disable copying the DSCP field from outer IP header to inner IP Header. Valid values: `disable`, `enable`.

* `inbound_dscp_copy_port` - Physical interfaces that support inbound-dscp-copy.
* `ip_reassembly` - Ip-Reassembly. The structure of `ip_reassembly` block is documented below.
* `intf_shaping_offload` - Enable/disable NPU offload when doing interface-based traffic shaping according to the egress-shaping-profile. Valid values: `disable`, `enable`.

* `ip_fragment_offload` - Enable/disable NP7 NPU IP fragment offload. Valid values: `disable`, `enable`.

* `iph_rsvd_re_cksum` - Enable/disable IP checksum re-calculation for packets with iph.reserved bit set. Valid values: `disable`, `enable`.

* `ippool_overload_high` - High threshold for overload ippool port reuse (100%-2000%, default = 200).
* `ippool_overload_low` - Low threshold for overload ippool port reuse (100%-2000%, default = 150).
* `ipsec_sts_timeout` - Set NP7Lite IPsec STS msg timeout. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.

* `ipsec_dec_subengine_mask` - IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_enc_subengine_mask` - IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_host_dfclr` - Enable/disable DF clearing of NP4lite host IPsec offload. Valid values: `disable`, `enable`.

* `ipsec_inbound_cache` - Enable/disable IPsec inbound cache for anti-replay. Valid values: `disable`, `enable`.

* `ipsec_local_uesp_port` - Ipsec-Local-Uesp-Port.
* `ipsec_mtu_override` - Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.

* `ipsec_ob_np_sel` - IPsec NP selection for OB SA offloading. Valid values: `RR`, `Packet`, `Hash`.

* `ipsec_over_vlink` - Enable/disable IPSEC over vlink. Valid values: `disable`, `enable`.

* `ipv4_session_quota` - Enable/Disable NoNAT IPv4 session quota for hyperscale VDOMs. Valid values: `disable`, `enable`.

* `ipv4_session_quota_high` - Configure NoNAT IPv4 session quota high threshold.
* `ipv4_session_quota_low` - Configure NoNAT IPv4 session quota low threshold.
* `ipv6_prefix_session_quota` - Enable/Disable hardware IPv6 /64 prefix session quota for hyperscale VDOMs. Valid values: `disable`, `enable`.

* `ipv6_prefix_session_quota_high` - Configure IPv6 prefix session quota high threshold.
* `ipv6_prefix_session_quota_low` - Configure IPv6 prefix session quota low threshold.
* `ipsec_throughput_msg_frequency` - Set NP7Lite IPsec throughput msg frequency: 0--disable 1--32KB 3--64KB ... 0x3fff--256MB 0x7fff--512MB 0xffff--1GB. Valid values: `disable`, `32KB`, `64KB`, `128KB`, `256KB`, `512KB`, `1MB`, `2MB`, `4MB`, `8MB`, `16MB`, `32MB`, `64MB`, `128MB`, `256MB`, `512MB`, `1GB`.

* `ipt_sts_timeout` - Set NP7Lite IPT STS msg timeout. Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.

* `ipt_throughput_msg_frequency` - Set NP7Lite IPT throughput msg frequency: 0--disable 1--32KB 3--64KB ... 0x3fff--256MB 0x7fff--512MB 0xffff--1GB. Valid values: `disable`, `32KB`, `64KB`, `128KB`, `256KB`, `512KB`, `1MB`, `2MB`, `4MB`, `8MB`, `16MB`, `32MB`, `64MB`, `128MB`, `256MB`, `512MB`, `1GB`.

* `isf_np_queues` - Isf-Np-Queues. The structure of `isf_np_queues` block is documented below.
* `isf_np_rx_tr_distr` - Select ISF NP Rx trunk distribution (PSC) mode. Valid values: `port-flow`, `round-robin`, `randomized`.

* `lag_out_port_select` - Enable/disable LAG outgoing port selection based on incoming traffic port. Valid values: `disable`, `enable`.

* `max_receive_unit` - Set the maximum packet size for receive, larger packets will be silently dropped.
* `max_session_timeout` - Maximum time interval for refreshing NPU-offloaded sessions (10 - 1000 sec, default 40 sec).
* `mcast_session_accounting` - Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `disable`, `session-based`, `tpe-based`.

* `mcast_session_counting` - Mcast-Session-Counting. Valid values: `disable`, `enable`, `session-based`, `tpe-based`.

* `mcast_session_counting6` - Enable/disable traffic accounting for each multicast session6 through TAE counter. Valid values: `disable`, `enable`, `session-based`, `tpe-based`.

* `napi_break_interval` - NAPI break interval (default 0).
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `disable`, `enable`.

* `np_queues` - Np-Queues. The structure of `np_queues` block is documented below.
* `np6_cps_optimization_mode` - Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `disable`, `enable`.

* `npu_group_effective_scope` - npu-group-effective-scope defines under which npu-group cmds such as list/purge will be excecuted. Default scope is for all four HS-ok groups. (0-3, default = 255).
* `npu_tcam` - Npu-Tcam. The structure of `npu_tcam` block is documented below.
* `nss_threads_option` - Configure thread options for the NP7's NSS module. Valid values: `4t-eif`, `4t-noeif`, `2t`.

* `pba_eim` - Configure option for PBA(non-overload)/EIM combination. Valid values: `disallow`, `allow`.

* `pba_port_select_mode` - Port selection mode for PBA IP pool. Valid values: `random`, `direct`.

* `per_policy_accounting` - Set per-policy accounting. Valid values: `disable`, `enable`.

* `per_session_accounting` - Enable/disable per-session accounting. Valid values: `enable`, `disable`, `enable-by-log`, `all-enable`, `traffic-log-only`.

* `ple_non_syn_tcp_action` - Configure action for the PLE to take on TCP packets that have the SYN field unset. Valid values: `forward`, `drop`.

* `policy_offload_level` - Firewall Policy Offload Level(DISABLE/DOS/FULL). Valid values: `disable`, `dos-offload`, `full-offload`.

* `port_cpu_map` - Port-Cpu-Map. The structure of `port_cpu_map` block is documented below.
* `port_npu_map` - Port-Npu-Map. The structure of `port_npu_map` block is documented below.
* `port_path_option` - Port-Path-Option. The structure of `port_path_option` block is documented below.
* `priority_protocol` - Priority-Protocol. The structure of `priority_protocol` block is documented below.
* `prp_session_clear_mode` - PRP session clear mode for excluded ip sessions. Valid values: `blocking`, `non-blocking`, `do-not-clear`.

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

* `shaping_stats` - Enable/disable NP7 traffic shaping statistics (default = disable). Valid values: `disable`, `enable`.

* `spa_port_select_mode` - Port selection mode for SPA IP pool. Valid values: `random`, `direct`.

* `split_ipsec_engines` - Enable/disable Split IPsec Engines. Valid values: `disable`, `enable`.

* `sse_backpressure` - Enable/disable sse backpressure. Valid values: `disable`, `enable`.

* `sse_ha_scan` - Sse-Ha-Scan. The structure of `sse_ha_scan` block is documented below.
* `strip_clear_text_padding` - Enable/disable stripping clear text padding. Valid values: `disable`, `enable`.

* `strip_esp_padding` - Enable/disable stripping ESP padding. Valid values: `disable`, `enable`.

* `sw_eh_hash` - Sw-Eh-Hash. The structure of `sw_eh_hash` block is documented below.
* `sw_np_bandwidth` - Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`, `7G`, `8G`, `9G`.

* `sw_tr_hash` - Sw-Tr-Hash. The structure of `sw_tr_hash` block is documented below.
* `switch_np_hash` - Switch-NP trunk port selection Criteria. Valid values: `src-ip`, `dst-ip`, `src-dst-ip`.

* `tcp_rst_timeout` - TCP RST timeout in seconds (0-3600, default = 5).
* `tunnel_over_vlink` - Enable/disable selection of which NP6 chip the tunnel uses (default = enable). Valid values: `disable`, `enable`.

* `tcp_timeout_profile` - Tcp-Timeout-Profile. The structure of `tcp_timeout_profile` block is documented below.
* `udp_timeout_profile` - Udp-Timeout-Profile. The structure of `udp_timeout_profile` block is documented below.
* `uesp_offload` - Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `disable`, `enable`.

* `ull_port_mode` - Set ULL port's speed to 10G/25G (default 10G). Valid values: `10G`, `25G`.

* `vlan_lookup_cache` - Enable/disable vlan lookup cache (default enabled). Valid values: `disable`, `enable`.

* `vxlan_offload` - Enable/disable offloading vxlan. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `background_sse_scan` block supports:

* `scan` - Enable/disable background SSE scan by driver thread(default enabled). Valid values: `disable`, `enable`.

* `scan_stale` - Configure scanning of active or stale sessions (default = 0 = active sessions).
* `scan_vt` - Select version/type to scan: bit-0: 44; bit-1: 46; bit-2: 64; bit-3: 66 (default = 0xF).
* `stats_qual_access` - Statistics update access qualification in seconds (0 - INT_MAX, default = 180).
* `stats_qual_duration` - Statistics update duration qualification in seconds (0 - INT_MAX, default = 300).
* `stats_update_interval` - Stats update interval(&gt;=5*60 seconds, default 5*60 seconds).
* `udp_keepalive_interval` - UDP keepalive interval(&gt;=90 seconds, default 90 seconds).
* `udp_qual_access` - UDP keepalive access qualification in seconds (0 - INT_MAX, default = 30).
* `udp_qual_duration` - UDP keepalive duration qualification in seconds (0 - INT_MAX, default = 90).

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

* `sctp_csum_err` - Invalid IPv4 SCTP checksum anomalies. Valid values: `allow`, `drop`, `trap-to-host`.

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

The `icmp_error_rate_ctrl` block supports:

* `icmpv4_error_bucket_size` - Bucket size used in the token bucket algorithm for controlling the flow of ICMPv4 error packets (1 - 100, default = 20).
* `icmpv4_error_rate` - Average rate of ICMPv4 error packets that allowed to be generated per second (1 - 100, default = 1).
* `icmpv4_error_rate_limit` - Enable to limit the ICMPv4 error packets generated by this FortiGate. Valid values: `disable`, `enable`.

* `icmpv6_error_bucket_size` - Bucket size used in the token bucket algorithm for controlling the flow of ICMPv6 error packets (1 - 100, default = 20).
* `icmpv6_error_rate` - Average rate of ICMPv6 error packets that allowed to be generated per second (1 - 100, default = 1).
* `icmpv6_error_rate_limit` - Enable to limit the ICMPv6 error packets generated by this FortiGate. Valid values: `disable`, `enable`.


The `icmp_rate_ctrl` block supports:

* `icmp_v4_bucket_size` - Bucket size used in the token bucket algorithm for controlling the flow of ICMPv4 packets (1 - 100, default = 10).
* `icmp_v4_rate` - Average rate of ICMPv4 packets that allowed to be generated per second (1 - 100, default = 1).
* `icmp_v6_bucket_size` - Bucket size used in the token bucket algorithm for controlling the flow of ICMPv6 packets (1 - 100, default = 10).
* `icmp_v6_rate` - Average rate of ICMPv6 packets that allowed to be generated per second (1 - 100, default = 1).

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

The `npu_tcam` block supports:

* `data` - Data. The structure of `data` block is documented below.
* `dbg_dump` - Debug driver dump data/mask pdq.
* `mask` - Mask. The structure of `mask` block is documented below.
* `mir_act` - Mir-Act. The structure of `mir_act` block is documented below.
* `name` - NPU TCAM policies name.
* `oid` - NPU TCAM OID.
* `pri_act` - Pri-Act. The structure of `pri_act` block is documented below.
* `sact` - Sact. The structure of `sact` block is documented below.
* `tact` - Tact. The structure of `tact` block is documented below.
* `type` - TCAM policy type. Valid values: `L2_src_tc`, `L2_tgt_tc`, `L2_src_mir`, `L2_tgt_mir`, `L2_src_act`, `L2_tgt_act`, `IPv4_src_tc`, `IPv4_tgt_tc`, `IPv4_src_mir`, `IPv4_tgt_mir`, `IPv4_src_act`, `IPv4_tgt_act`, `IPv6_src_tc`, `IPv6_tgt_tc`, `IPv6_src_mir`, `IPv6_tgt_mir`, `IPv6_src_act`, `IPv6_tgt_act`.

* `vid` - NPU TCAM VID.

The `data` block supports:

* `df` - tcam data ip flag df. Valid values: `disable`, `enable`.

* `dstip` - tcam data dst ipv4 address.
* `dstipv6` - tcam data dst ipv6 address.
* `dstmac` - tcam data dst macaddr.
* `dstport` - tcam data L4 dst port.
* `ethertype` - tcam data ethertype.
* `ext_tag` - tcam data extension tag. Valid values: `disable`, `enable`.

* `frag_off` - tcam data ip flag fragment offset.
* `gen_buf_cnt` - tcam data gen info buffer count.
* `gen_iv` - tcam data gen info iv. Valid values: `invalid`, `valid`.

* `gen_l3_flags` - tcam data gen info L3 flags.
* `gen_l4_flags` - tcam data gen info L4 flags.
* `gen_pkt_ctrl` - tcam data gen info packet control.
* `gen_pri` - tcam data gen info priority.
* `gen_pri_v` - tcam data gen info priority valid. Valid values: `invalid`, `valid`.

* `gen_tv` - tcam data gen info tv. Valid values: `invalid`, `valid`.

* `ihl` - tcam data ipv4 IHL.
* `ip4_id` - tcam data ipv4 id.
* `ip6_fl` - tcam data ipv6 flow label.
* `ipver` - tcam data ip header version.
* `l4_wd10` - tcam data L4 word10.
* `l4_wd11` - tcam data L4 word11.
* `l4_wd8` - tcam data L4 word8.
* `l4_wd9` - tcam data L4 word9.
* `mf` - tcam data ip flag mf. Valid values: `disable`, `enable`.

* `protocol` - tcam data ip protocol.
* `slink` - tcam data sublink.
* `smac_change` - tcam data source MAC change. Valid values: `disable`, `enable`.

* `sp` - tcam data source port.
* `src_cfi` - tcam data source cfi. Valid values: `disable`, `enable`.

* `src_prio` - tcam data source priority.
* `src_updt` - tcam data source update. Valid values: `disable`, `enable`.

* `srcip` - tcam data src ipv4 address.
* `srcipv6` - tcam data src ipv6 address.
* `srcmac` - tcam data src macaddr.
* `srcport` - tcam data L4 src port.
* `svid` - tcam data source vid.
* `tcp_ack` - tcam data tcp flag ack. Valid values: `disable`, `enable`.

* `tcp_cwr` - tcam data tcp flag cwr. Valid values: `disable`, `enable`.

* `tcp_ece` - tcam data tcp flag ece. Valid values: `disable`, `enable`.

* `tcp_fin` - tcam data tcp flag fin. Valid values: `disable`, `enable`.

* `tcp_push` - tcam data tcp flag push. Valid values: `disable`, `enable`.

* `tcp_rst` - tcam data tcp flag rst. Valid values: `disable`, `enable`.

* `tcp_syn` - tcam data tcp flag syn. Valid values: `disable`, `enable`.

* `tcp_urg` - tcam data tcp flag urg. Valid values: `disable`, `enable`.

* `tgt_cfi` - tcam data target cfi. Valid values: `disable`, `enable`.

* `tgt_prio` - tcam data target priority.
* `tgt_updt` - tcam data target port update. Valid values: `disable`, `enable`.

* `tgt_v` - tcam data target valid. Valid values: `invalid`, `valid`.

* `tos` - tcam data ip tos.
* `tp` - tcam data target port.
* `ttl` - tcam data ip ttl.
* `tvid` - tcam data target vid.
* `vdid` - tcam data vdom id.

The `mask` block supports:

* `df` - tcam mask ip flag df. Valid values: `disable`, `enable`.

* `dstip` - tcam mask dst ipv4 address.
* `dstipv6` - tcam mask dst ipv6 address.
* `dstmac` - tcam mask dst macaddr.
* `dstport` - tcam mask L4 dst port.
* `ethertype` - tcam mask ethertype.
* `ext_tag` - tcam mask extension tag. Valid values: `disable`, `enable`.

* `frag_off` - tcam data ip flag fragment offset.
* `gen_buf_cnt` - tcam mask gen info buffer count.
* `gen_iv` - tcam mask gen info iv. Valid values: `invalid`, `valid`.

* `gen_l3_flags` - tcam mask gen info L3 flags.
* `gen_l4_flags` - tcam mask gen info L4 flags.
* `gen_pkt_ctrl` - tcam mask gen info packet control.
* `gen_pri` - tcam mask gen info priority.
* `gen_pri_v` - tcam mask gen info priority valid. Valid values: `invalid`, `valid`.

* `gen_tv` - tcam mask gen info tv. Valid values: `invalid`, `valid`.

* `ihl` - tcam mask ipv4 IHL.
* `ip4_id` - tcam mask ipv4 id.
* `ip6_fl` - tcam mask ipv6 flow label.
* `ipver` - tcam mask ip header version.
* `l4_wd10` - tcam mask L4 word10.
* `l4_wd11` - tcam mask L4 word11.
* `l4_wd8` - tcam mask L4 word8.
* `l4_wd9` - tcam mask L4 word9.
* `mf` - tcam mask ip flag mf. Valid values: `disable`, `enable`.

* `protocol` - tcam mask ip protocol.
* `slink` - tcam mask sublink.
* `smac_change` - tcam mask source MAC change. Valid values: `disable`, `enable`.

* `sp` - tcam mask source port.
* `src_cfi` - tcam mask source cfi. Valid values: `disable`, `enable`.

* `src_prio` - tcam mask source priority.
* `src_updt` - tcam mask source update. Valid values: `disable`, `enable`.

* `srcip` - tcam mask src ipv4 address.
* `srcipv6` - tcam mask src ipv6 address.
* `srcmac` - tcam mask src macaddr.
* `srcport` - tcam mask L4 src port.
* `svid` - tcam mask source vid.
* `tcp_ack` - tcam mask tcp flag ack. Valid values: `disable`, `enable`.

* `tcp_cwr` - tcam mask tcp flag cwr. Valid values: `disable`, `enable`.

* `tcp_ece` - tcam mask tcp flag ece. Valid values: `disable`, `enable`.

* `tcp_fin` - tcam mask tcp flag fin. Valid values: `disable`, `enable`.

* `tcp_push` - tcam mask tcp flag push. Valid values: `disable`, `enable`.

* `tcp_rst` - tcam mask tcp flag rst. Valid values: `disable`, `enable`.

* `tcp_syn` - tcam mask tcp flag syn. Valid values: `disable`, `enable`.

* `tcp_urg` - tcam mask tcp flag urg. Valid values: `disable`, `enable`.

* `tgt_cfi` - tcam mask target cfi. Valid values: `disable`, `enable`.

* `tgt_prio` - tcam mask target priority.
* `tgt_updt` - tcam mask target port update. Valid values: `disable`, `enable`.

* `tgt_v` - tcam mask target valid. Valid values: `invalid`, `valid`.

* `tos` - tcam mask ip tos.
* `tp` - tcam mask target port.
* `ttl` - tcam mask ip ttl.
* `tvid` - tcam mask target vid.
* `vdid` - tcam mask vdom id.

The `mir_act` block supports:

* `vlif` - tcam mirror action vlif.

The `pri_act` block supports:

* `priority` - tcam priority action priority.
* `weight` - tcam priority action weight.

The `sact` block supports:

* `act` - tcam sact act.
* `act_v` - Enable to set sact act. Valid values: `disable`, `enable`.

* `bmproc` - tcam sact bmproc.
* `bmproc_v` - Enable to set sact bmproc. Valid values: `disable`, `enable`.

* `df_lif` - tcam sact df-lif.
* `df_lif_v` - Enable to set sact df-lif. Valid values: `disable`, `enable`.

* `dfr` - tcam sact dfr.
* `dfr_v` - Enable to set sact dfr. Valid values: `disable`, `enable`.

* `dmac_skip` - tcam sact dmac-skip.
* `dmac_skip_v` - Enable to set sact dmac-skip. Valid values: `disable`, `enable`.

* `dosen` - tcam sact dosen.
* `dosen_v` - Enable to set sact dosen. Valid values: `disable`, `enable`.

* `espff_proc` - tcam sact espff-proc.
* `espff_proc_v` - Enable to set sact espff-proc. Valid values: `disable`, `enable`.

* `etype_pid` - tcam sact etype-pid.
* `etype_pid_v` - Enable to set sact etype-pid. Valid values: `disable`, `enable`.

* `frag_proc` - tcam sact frag-proc.
* `frag_proc_v` - Enable to set sact frag-proc. Valid values: `disable`, `enable`.

* `fwd` - tcam sact fwd.
* `fwd_lif` - tcam sact fwd-lif.
* `fwd_lif_v` - Enable to set sact fwd-lif. Valid values: `disable`, `enable`.

* `fwd_tvid` - tcam sact fwd-tvid.
* `fwd_tvid_v` - Enable to set sact fwd-vid. Valid values: `disable`, `enable`.

* `fwd_v` - Enable to set sact fwd. Valid values: `disable`, `enable`.

* `icpen` - tcam sact icpen.
* `icpen_v` - Enable to set sact icpen. Valid values: `disable`, `enable`.

* `igmp_mld_snp` - tcam sact igmp-mld-snp.
* `igmp_mld_snp_v` - Enable to set sact igmp-mld-snp. Valid values: `disable`, `enable`.

* `learn` - tcam sact learn.
* `learn_v` - Enable to set sact learn. Valid values: `disable`, `enable`.

* `m_srh_ctrl` - tcam sact m-srh-ctrl.
* `m_srh_ctrl_v` - Enable to set sact m-srh-ctrl. Valid values: `disable`, `enable`.

* `mac_id` - tcam sact mac-id.
* `mac_id_v` - Enable to set sact mac-id. Valid values: `disable`, `enable`.

* `mss` - tcam sact mss.
* `mss_v` - Enable to set sact mss. Valid values: `disable`, `enable`.

* `pleen` - tcam sact pleen.
* `pleen_v` - Enable to set sact pleen. Valid values: `disable`, `enable`.

* `prio_pid` - tcam sact prio-pid.
* `prio_pid_v` - Enable to set sact prio-pid. Valid values: `disable`, `enable`.

* `promis` - tcam sact promis.
* `promis_v` - Enable to set sact promis. Valid values: `disable`, `enable`.

* `rfsh` - tcam sact rfsh.
* `rfsh_v` - Enable to set sact rfsh. Valid values: `disable`, `enable`.

* `smac_skip` - tcam sact smac-skip.
* `smac_skip_v` - Enable to set sact smac-skip. Valid values: `disable`, `enable`.

* `tp_smchk_v` - Enable to set sact tp mode. Valid values: `disable`, `enable`.

* `tp_smchk` - tcam sact tp mode.
* `tpe_id` - tcam sact tpe-id.
* `tpe_id_v` - Enable to set sact tpe-id. Valid values: `disable`, `enable`.

* `vdm` - tcam sact vdm.
* `vdm_v` - Enable to set sact vdm. Valid values: `disable`, `enable`.

* `vdom_id` - tcam sact vdom-id.
* `vdom_id_v` - Enable to set sact vdom-id. Valid values: `disable`, `enable`.

* `x_mode` - tcam sact x-mode.
* `x_mode_v` - Enable to set sact x-mode. Valid values: `disable`, `enable`.


The `tact` block supports:

* `act` - tcam tact act.
* `act_v` - Enable to set tact act. Valid values: `disable`, `enable`.

* `fmtuv4_s` - tcam tact fmtuv4-s.
* `fmtuv4_s_v` - Enable to set tact fmtuv4-s. Valid values: `disable`, `enable`.

* `fmtuv6_s` - tcam tact fmtuv6-s.
* `fmtuv6_s_v` - Enable to set tact fmtuv6-s. Valid values: `disable`, `enable`.

* `lnkid` - tcam tact lnkid.
* `lnkid_v` - Enable to set tact lnkid. Valid values: `disable`, `enable`.

* `mac_id` - tcam tact mac-id.
* `mac_id_v` - Enable to set tact mac-id. Valid values: `disable`, `enable`.

* `mss_t` - tcam tact mss.
* `mss_t_v` - Enable to set tact mss. Valid values: `disable`, `enable`.

* `mtuv4` - tcam tact mtuv4.
* `mtuv4_v` - Enable to set tact mtuv4. Valid values: `disable`, `enable`.

* `mtuv6` - tcam tact mtuv6.
* `mtuv6_v` - Enable to set tact mtuv6. Valid values: `disable`, `enable`.

* `slif_act` - tcam tact slif-act.
* `slif_act_v` - Enable to set tact slif-act. Valid values: `disable`, `enable`.

* `sublnkid` - tcam tact sublnkid.
* `sublnkid_v` - Enable to set tact sublnkid. Valid values: `disable`, `enable`.

* `tgtv_act` - tcam tact tgtv-act.
* `tgtv_act_v` - Enable to set tact tgtv-act. Valid values: `disable`, `enable`.

* `tlif_act` - tcam tact tlif-act.
* `tlif_act_v` - Enable to set tact tlif-act. Valid values: `disable`, `enable`.

* `tpeid` - tcam tact tpeid.
* `tpeid_v` - Enable to set tact tpeid. Valid values: `disable`, `enable`.

* `v6fe` - tcam tact v6fe.
* `v6fe_v` - Enable to set tact v6fe. Valid values: `disable`, `enable`.

* `vep_en_v` - Enable to set tact vep-en. Valid values: `disable`, `enable`.

* `vep_slid` - tcam tact vep_slid.
* `vep_slid_v` - Enable to set tact vep-slid. Valid values: `disable`, `enable`.

* `vep_en` - tcam tact vep_en.
* `xlt_lif` - tcam tact xlt-lif.
* `xlt_lif_v` - Enable to set tact xlt-lif. Valid values: `disable`, `enable`.

* `xlt_vid` - tcam tact xlt-vid.
* `xlt_vid_v` - Enable to set tact xlt-vid. Valid values: `disable`, `enable`.


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


The `sw_tr_hash` block supports:

* `draco15` - Enable/disable DRACO15 hashing. Valid values: `disable`, `enable`.

* `tcp_udp_port` - Include/exclude TCP/UDP source and destination port for unicast trunk traffic. Valid values: `include`, `exclude`.


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
