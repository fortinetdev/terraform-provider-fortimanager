---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_healthcheck"
description: |-
  SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.
---

# fortimanager_wantemp_system_sdwan_healthcheck
SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

~> This resource is a sub resource for variable `health_check` of resource `fortimanager_wantemp_system_sdwan`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `sla`: `fortimanager_wantemp_system_sdwan_healthcheck_sla`



## Example Usage

```hcl
resource "fortimanager_wantemp_system_sdwan_healthcheck" "trname" {
  addr_mode  = "ipv4"
  name       = "healthcheck"
  wanprof    = fortimanager_wan_template.trname.name
  depends_on = [fortimanager_wan_template.trname]
}

resource "fortimanager_wan_template" "trname" {
  name = "terr21"
  adom = "root"
  type = "wanprof"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

* `_dynamic_server` - _Dynamic-Server.
* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `class_id` - Traffic class ID.
* `detect_mode` - The mode determining how to detect the server. Valid values: `active`, `passive`, `prefer-passive`.

* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `dns_match_ip` - Response IP expected from DNS server if the protocol is DNS.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `embed_measured_health` - Enable/disable embedding measured health information. Valid values: `disable`, `enable`.

* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `fortiguard` - Enable/disable use of FortiGuard predefined server. Valid values: `disable`, `enable`.

* `fortiguard_name` - Predefined health-check target name.
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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `sla` block supports:

* `id` - SLA ID.
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`.

* `mos_threshold` - Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
* `priority_in_sla` - Value to be distributed into routing table when in-sla (0 - 65535, default = 0).
* `priority_out_sla` - Value to be distributed into routing table when out-sla (0 - 65535, default = 0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wantemp SystemSdwanHealthCheck can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_sdwan_healthcheck.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
