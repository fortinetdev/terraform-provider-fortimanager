---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_neighbor"
description: |-
  Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.
---

# fortimanager_wantemp_system_sdwan_neighbor
Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

~> This resource is a sub resource for variable `neighbor` of resource `fortimanager_wantemp_system_sdwan`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

* `health_check` - SD-WAN health-check name.
* `ip` - IP/IPv6 address of neighbor.
* `member` - Member sequence number.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `mode` - What metric to select the neighbor. Valid values: `sla`, `speedtest`.

* `role` - Role of neighbor. Valid values: `primary`, `secondary`, `standalone`.

* `route_metric` - Route-metric of neighbor. Valid values: `preferable`, `priority`.

* `service_id` - SD-WAN service ID to work with the neighbor.
* `sla_id` - SLA ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip}}.

## Import

Wantemp SystemSdwanNeighbor can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_sdwan_neighbor.labelname {{ip}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
