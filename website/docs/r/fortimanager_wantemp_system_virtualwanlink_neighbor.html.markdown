---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_virtualwanlink_neighbor"
description: |-
  Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.
---

# fortimanager_wantemp_system_virtualwanlink_neighbor
Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

* `health_check` - SD-WAN health-check name.
* `ip` - IP address of neighbor.
* `member` - Member sequence number.
* `role` - Role of neighbor. Valid values: `primary`, `secondary`, `standalone`.

* `sla_id` - SLA ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{sla_id}}.

## Import

Wantemp SystemVirtualWanLinkNeighbor can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_virtualwanlink_neighbor.labelname {{sla_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
