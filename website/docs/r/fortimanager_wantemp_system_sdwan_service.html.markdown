---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_service"
description: |-
  Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.
---

# fortimanager_wantemp_system_sdwan_service
Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

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
* `hash_mode` - Hash algorithm for selected priority members for load balance mode. Valid values: `round-robin`, `source-ip-based`, `source-dest-ip-based`, `inbandwidth`, `outbandwidth`, `bibandwidth`.

* `health_check` - Health check list.
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `fosid` - SD-WAN rule ID (1 - 4000).
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

* `route_tag` - IPv4 route map route-tag.
* `sla` - Sla. The structure of `sla` block is documented below.
* `sla_compare_method` - Method to compare SLA value for SLA mode. Valid values: `order`, `number`.

* `src` - Source address name.
* `src_negate` - Enable/disable negation of source address match. Valid values: `disable`, `enable`.

* `src6` - Source address6 name.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `disable`, `enable`.

* `start_port` - Start destination port number.
* `status` - Enable/disable SD-WAN service. Valid values: `disable`, `enable`.

* `tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `zone`, `cfg-order`, `fib-best-match`.

* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `use_shortcut_sla` - Enable/disable use of ADVPN shortcut for quality comparison. Valid values: `disable`, `enable`.

* `users` - User name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `sla` block supports:

* `health_check` - SD-WAN health-check.
* `id` - SLA ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wantemp SystemSdwanService can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=mywanprof"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_sdwan_service.labelname WantempSystemSdwanService
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
