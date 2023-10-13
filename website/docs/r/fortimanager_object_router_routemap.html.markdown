---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_router_routemap"
description: |-
  Configure route maps.
---

# fortimanager_object_router_routemap
Configure route maps.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comments` - Optional comments.
* `name` - Name.
* `rule` - Rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Action. Valid values: `permit`, `deny`.

* `id` - Rule ID.
* `match_as_path` - Match BGP AS path list.
* `match_community` - Match BGP community list.
* `match_community_exact` - Enable/disable exact matching of communities. Valid values: `disable`, `enable`.

* `match_extcommunity` - Match BGP extended community list.
* `match_extcommunity_exact` - Enable/disable exact matching of extended communities. Valid values: `disable`, `enable`.

* `match_flags` - Match-Flags.
* `match_interface` - Match interface configuration.
* `match_ip_address` - Match IP address permitted by access-list or prefix-list.
* `match_ip_nexthop` - Match next hop IP address passed by access-list or prefix-list.
* `match_ip6_address` - Match IPv6 address permitted by access-list6 or prefix-list6.
* `match_ip6_nexthop` - Match next hop IPv6 address passed by access-list6 or prefix-list6.
* `match_metric` - Match metric for redistribute routes.
* `match_origin` - Match BGP origin code. Valid values: `none`, `egp`, `igp`, `incomplete`.

* `match_route_type` - Match route type. Valid values: `1`, `2`, `none`, `external-type1`, `external-type2`.

* `match_tag` - Match tag.
* `match_vrf` - Match VRF ID.
* `set_aggregator_as` - BGP aggregator AS.
* `set_aggregator_ip` - BGP aggregator IP.
* `set_aspath` - Prepend BGP AS path attribute.
* `set_aspath_action` - Specify preferred action of set-aspath. Valid values: `prepend`, `replace`.

* `set_atomic_aggregate` - Enable/disable BGP atomic aggregate attribute. Valid values: `disable`, `enable`.

* `set_community` - BGP community attribute.
* `set_community_additive` - Enable/disable adding set-community to existing community. Valid values: `disable`, `enable`.

* `set_community_delete` - Delete communities matching community list.
* `set_dampening_max_suppress` - Maximum duration to suppress a route (1 - 255 min, 0 = unset).
* `set_dampening_reachability_half_life` - Reachability half-life time for the penalty (1 - 45 min, 0 = unset).
* `set_dampening_reuse` - Value to start reusing a route (1 - 20000, 0 = unset).
* `set_dampening_suppress` - Value to start suppressing a route (1 - 20000, 0 = unset).
* `set_dampening_unreachability_half_life` - Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)
* `set_extcommunity_rt` - Route Target extended community.
* `set_extcommunity_soo` - Site-of-Origin extended community.
* `set_flags` - Set-Flags.
* `set_ip_nexthop` - IP address of next hop.
* `set_ip_prefsrc` - IP address of preferred source.
* `set_ip6_nexthop` - IPv6 global address of next hop.
* `set_ip6_nexthop_local` - IPv6 local address of next hop.
* `set_local_preference` - BGP local preference path attribute.
* `set_metric` - Metric value.
* `set_metric_type` - Metric type. Valid values: `1`, `2`, `none`, `external-type1`, `external-type2`.

* `set_origin` - BGP origin code. Valid values: `none`, `egp`, `igp`, `incomplete`.

* `set_originator_id` - BGP originator ID attribute.
* `set_priority` - Priority for routing table.
* `set_route_tag` - Route tag for routing table.
* `set_tag` - Tag value.
* `set_vpnv4_nexthop` - IP address of VPNv4 next-hop.
* `set_weight` - BGP weight for routing table.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectRouter RouteMap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_router_routemap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
