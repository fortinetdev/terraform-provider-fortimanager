---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan"
description: |-
  Configure redundant Internet connections with multiple outbound links and health-check profiles.
---

# fortimanager_wantemp_system_sdwan
Configure redundant Internet connections with multiple outbound links and health-check profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `duplication`: `fortimanager_wantemp_system_sdwan_duplication`
>- `health_check`: `fortimanager_wantemp_system_sdwan_healthcheck`
>- `members`: `fortimanager_wantemp_system_sdwan_members`
>- `neighbor`: `fortimanager_wantemp_system_sdwan_neighbor`
>- `service`: `fortimanager_wantemp_system_sdwan_service`
>- `zone`: `fortimanager_wantemp_system_sdwan_zone`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

* `app_perf_log_period` - Time interval in seconds that applicationperformance logs are generated (0 - 3600, default = 0).
* `duplication` - Duplication. The structure of `duplication` block is documented below.
* `duplication_max_num` - Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
* `fail_alert_interfaces` - Physical interfaces that will be alerted.
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `disable`, `enable`.

* `health_check` - Health-Check. The structure of `health_check` block is documented below.
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.

* `members` - Members. The structure of `members` block is documented below.
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `neighbor_hold_boot_time` - Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
* `neighbor_hold_down` - Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `disable`, `enable`.

* `neighbor_hold_down_time` - Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
* `service` - Service. The structure of `service` block is documented below.
* `speedtest_bypass_routing` - Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable`, `enable`.

* `status` - Enable/disable SD-WAN. Valid values: `disable`, `enable`.

* `zone` - Zone. The structure of `zone` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `duplication` block supports:

* `dstaddr` - Destination address or address group names.
* `dstaddr6` - Destination address6 or address6 group names.
* `dstintf` - Outgoing (egress) interfaces or zones.
* `id` - Duplication rule ID (1 - 255).
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated. Valid values: `disable`, `enable`.

* `packet_duplication` - Configure packet duplication method. Valid values: `disable`, `force`, `on-demand`.

* `service` - Service and service group name.
* `service_id` - SD-WAN service rule ID list.
* `sla_match_service` - Enable/disable packet duplication matching health-check SLAs in service rule. Valid values: `disable`, `enable`.

* `srcaddr` - Source address or address group names.
* `srcaddr6` - Source address6 or address6 group names.
* `srcintf` - Incoming (ingress) interfaces or zones.

The `health_check` block supports:

* `_dynamic_server` - _Dynamic-Server.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `class_id` - Traffic class ID.
* `detect_mode` - The mode determining how to detect the server. Valid values: `active`, `passive`, `prefer-passive`.

* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `dns_match_ip` - Response IP expected from DNS server if the protocol is DNS.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `embed_measured_health` - Enable/disable embedding measured health information. Valid values: `disable`, `enable`.

* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `ftp_file` - Full path and file name on the FTP server to download for FTP health-check to probe.
* `ftp_mode` - FTP mode. Valid values: `passive`, `port`.

* `ha_priority` - HA election priority (1 - 50).
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `interval` - Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).
* `members` - Member sequence number list.
* `mos_codec` - Codec to use for MOS calculation (default = g711). Valid values: `g711`, `g722`, `g729`.

* `name` - Status check or health check name.
* `packet_size` - Packet size of a twamp test session,
* `password` - Twamp controller password in authentication mode
* `port` - Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, twamp: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21).
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `probe_packets` - Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.

* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `twamp`, `ping6`, `dns`, `tcp-connect`, `ftp`.

* `quality_measured_method` - Method to measure the quality of tcp-connect. Valid values: `half-close`, `half-open`.

* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.

* `server` - IP address or FQDN name of the server.
* `sla` - Sla. The structure of `sla` block is documented below.
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_id_redistribute` - Select the ID from the SLA sub-table. The selected SLA's priority value will be distributed into the routing table (0 - 32, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 addressused in the health-check packet to server.
* `system_dns` - Enable/disable system DNS as the probe server. Valid values: `disable`, `enable`.

* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `disable`, `enable`.

* `update_static_route` - Enable/disable updating the static route. Valid values: `disable`, `enable`.

* `user` - The user name to access probe server.
* `vrf` - Virtual Routing Forwarding ID.

The `sla` block supports:

* `id` - SLA ID.
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`.

* `mos_threshold` - Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
* `priority_in_sla` - Value to be distributed into routing table when in-sla (0 - 65535, default = 0).
* `priority_out_sla` - Value to be distributed into routing table when out-sla (0 - 65535, default = 0).

The `members` block supports:

* `_dynamic_member` - _Dynamic-Member.
* `comment` - Comments.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `gateway6` - IPv6 gateway.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `interface` - Interface name.
* `preferred_source` - Preferred source of route for this member.
* `priority` - Priority of the interface (0 - 65535). Used for SD-WAN rules or priority rules.
* `priority6` - Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.
* `seq_num` - Sequence number(1-512).
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.

* `transport_group` - Measured transport group (0 - 255).
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.
* `zone` - Zone name.

The `neighbor` block supports:

* `health_check` - SD-WAN health-check name.
* `ip` - IP/IPv6 address of neighbor.
* `member` - Member sequence number.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `mode` - What metric to select the neighbor. Valid values: `sla`, `speedtest`.

* `role` - Role of neighbor. Valid values: `primary`, `secondary`, `standalone`.

* `service_id` - SD-WAN service ID to work with the neighbor.
* `sla_id` - SLA ID.

The `service` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `agent_exclusive` - Set/unset the service as agent use exclusively. Valid values: `disable`, `enable`.

* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `default` - Enable/disable use of SD-WAN as default service. Valid values: `disable`, `enable`.

* `dscp_forward` - Enable/disable forward traffic DSCP tag. Valid values: `disable`, `enable`.

* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag. Valid values: `disable`, `enable`.

* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `dst` - Destination address name.
* `dst_negate` - Enable/disable negation of destination address match. Valid values: `disable`, `enable`.

* `dst6` - Destination address6 name.
* `end_port` - End destination port number.
* `end_src_port` - End source port number.
* `gateway` - Enable/disable SD-WAN service gateway. Valid values: `disable`, `enable`.

* `groups` - User groups.
* `hash_mode` - Hash algorithm for selected priority members for load balance mode. Valid values: `round-robin`, `source-ip-based`, `source-dest-ip-based`, `inbandwidth`, `outbandwidth`, `bibandwidth`.

* `health_check` - Health check list.
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `id` - SD-WAN rule ID (1 - 4000).
* `input_device` - Source interface name.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `disable`, `enable`.

* `input_zone` - Source input-zone name.
* `internet_service` - Enable/disable use of Internet service for application-based load balancing. Valid values: `disable`, `enable`.

* `internet_service_app_ctrl` - Application control based Internet Service ID list.
* `internet_service_app_ctrl_category` - IDs of one or more application control categories.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list.
* `internet_service_custom` - Custom Internet service name list.
* `internet_service_custom_group` - Custom Internet Service group list.
* `internet_service_group` - Internet Service group list.
* `internet_service_name` - Internet service name list.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `link_cost_factor` - Link cost factor. Valid values: `latency`, `jitter`, `packet-loss`, `inbandwidth`, `outbandwidth`, `bibandwidth`, `custom-profile-1`.

* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `load_balance` - Enable/disable load-balance. Valid values: `disable`, `enable`.

* `minimum_sla_meet_members` - Minimum number of members which meet SLA.
* `mode` - Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN. Valid values: `auto`, `manual`, `priority`, `sla`, `load-balance`.

* `name` - SD-WAN rule name.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `passive_measurement` - Enable/disable passive measurement based on the service criteria. Valid values: `disable`, `enable`.

* `priority_members` - Member sequence number list.
* `priority_zone` - Priority zone name list.
* `protocol` - Protocol number.
* `quality_link` - Quality grade.
* `role` - Service role to work with neighbor. Valid values: `primary`, `secondary`, `standalone`.

* `shortcut` - Enable/disable shortcut for this service. Valid values: `disable`, `enable`.

* `shortcut_priority` - High priority of ADVPN shortcut for this service. Valid values: `disable`, `enable`, `auto`.

* `shortcut_stickiness` - Enable/disable shortcut-stickiness of ADVPN. Valid values: `disable`, `enable`.

* `route_tag` - IPv4 route map route-tag.
* `sla` - Sla. The structure of `sla` block is documented below.
* `sla_compare_method` - Method to compare SLA value for SLA mode. Valid values: `order`, `number`.

* `sla_stickiness` - Enable/disable SLA stickiness (default = disable). Valid values: `disable`, `enable`.

* `src` - Source address name.
* `src_negate` - Enable/disable negation of source address match. Valid values: `disable`, `enable`.

* `src6` - Source address6 name.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `disable`, `enable`.

* `start_port` - Start destination port number.
* `start_src_port` - Start source port number.
* `status` - Enable/disable SD-WAN service. Valid values: `disable`, `enable`.

* `tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `zone`, `cfg-order`, `fib-best-match`.

* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `use_shortcut_sla` - Enable/disable use of ADVPN shortcut for quality comparison. Valid values: `disable`, `enable`.

* `users` - User name.
* `zone_mode` - Enable/disable zone mode. Valid values: `disable`, `enable`.


The `sla` block supports:

* `health_check` - SD-WAN health-check.
* `id` - SLA ID.

The `zone` block supports:

* `advpn_health_check` - Health check for ADVPN local overlay link quality.
* `advpn_select` - Enable/disable selection of ADVPN based on SDWAN information. Valid values: `disable`, `enable`.

* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `cfg-order`, `fib-best-match`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wantemp SystemSdwan can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_sdwan.labelname WantempSystemSdwan
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
