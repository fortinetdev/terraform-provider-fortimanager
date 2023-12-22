---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_virtualwanlink"
description: |-
  Configure redundant internet connections using SD-WAN (formerly virtual WAN link).
---

# fortimanager_wantemp_system_virtualwanlink
Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`health_check`: `fortimanager_wantemp_system_virtualwanlink_healthcheck`
`members`: `fortimanager_wantemp_system_virtualwanlink_members`
`neighbor`: `fortimanager_wantemp_system_virtualwanlink_neighbor`
`service`: `fortimanager_wantemp_system_virtualwanlink_service`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

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
* `status` - Enable/disable SD-WAN. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `health_check` block supports:

* `_dynamic_server` - _Dynamic-Server.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `ha_priority` - HA election priority (1 - 50).
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `internet_service_id` - Internet-Service-Id.
* `interval` - Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).
* `members` - Member sequence number list.
* `name` - Status check or health check name.
* `packet_size` - Packet size of a twamp test session,
* `password` - Twamp controller password in authentication mode
* `port` - Port number used to communicate with the server over the selected protocol.
* `probe_packets` - Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.

* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server. Valid values: `ping`, `tcp-echo`, `udp-echo`, `http`, `twamp`, `ping6`, `dns`.

* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `security_mode` - Twamp controller security mode. Valid values: `none`, `authentication`.

* `server` - IP address or FQDN name of the server.
* `sla` - Sla. The structure of `sla` block is documented below.
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `disable`, `enable`.

* `update_static_route` - Enable/disable updating the static route. Valid values: `disable`, `enable`.


The `sla` block supports:

* `id` - SLA ID.
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`.

* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).

The `members` block supports:

* `_dynamic_member` - _Dynamic-Member.
* `comment` - Comments.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `gateway6` - IPv6 gateway.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `interface` - Interface name.
* `priority` - Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.
* `seq_num` - Sequence number(1-255).
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable`, `enable`.

* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.

The `neighbor` block supports:

* `health_check` - SD-WAN health-check name.
* `ip` - IP address of neighbor.
* `member` - Member sequence number.
* `role` - Role of neighbor. Valid values: `primary`, `secondary`, `standalone`.

* `sla_id` - SLA ID.

The `service` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

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
* `gateway` - Enable/disable SD-WAN service gateway. Valid values: `disable`, `enable`.

* `groups` - User groups.
* `health_check` - Health check.
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `id` - Priority rule ID (1 - 4000).
* `input_device` - Source interface name.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `disable`, `enable`.

* `internet_service` - Enable/disable use of Internet service for application-based load balancing. Valid values: `disable`, `enable`.

* `internet_service_ctrl` - Control-based Internet Service ID list.
* `internet_service_ctrl_group` - Control-based Internet Service group list.
* `internet_service_app_ctrl` - Application control based Internet Service ID list.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list.
* `internet_service_custom` - Custom Internet service name list.
* `internet_service_custom_group` - Custom Internet Service group list.
* `internet_service_group` - Internet Service group list.
* `internet_service_id` - Internet service ID list.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `link_cost_factor` - Link cost factor. Valid values: `latency`, `jitter`, `packet-loss`, `inbandwidth`, `outbandwidth`, `bibandwidth`, `custom-profile-1`.

* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `member` - Member sequence number.
* `mode` - Control how the priority rule sets the priority of interfaces in the SD-WAN. Valid values: `auto`, `manual`, `priority`, `sla`, `load-balance`.

* `name` - Priority rule name.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `priority_members` - Member sequence number list.
* `protocol` - Protocol number.
* `quality_link` - Quality grade.
* `role` - Service role to work with neighbor. Valid values: `primary`, `secondary`, `standalone`.

* `route_tag` - IPv4 route map route-tag.
* `sla` - Sla. The structure of `sla` block is documented below.
* `sla_compare_method` - Method to compare SLA value for sla and load balance mode. Valid values: `order`, `number`.

* `src` - Source address name.
* `src_negate` - Enable/disable negation of source address match. Valid values: `disable`, `enable`.

* `src6` - Source address6 name.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `disable`, `enable`.

* `start_port` - Start destination port number.
* `status` - Enable/disable SD-WAN service. Valid values: `disable`, `enable`.

* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `users` - User name.

The `sla` block supports:

* `health_check` - Virtual WAN Link health-check.
* `id` - SLA ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wantemp SystemVirtualWanLink can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_virtualwanlink.labelname WantempSystemVirtualWanLink
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
