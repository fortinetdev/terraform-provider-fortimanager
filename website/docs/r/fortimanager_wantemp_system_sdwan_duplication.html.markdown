---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_duplication"
description: |-
  Create SD-WAN duplication rule.
---

# fortimanager_wantemp_system_sdwan_duplication
Create SD-WAN duplication rule.

~> This resource is a sub resource for variable `duplication` of resource `fortimanager_wantemp_system_sdwan`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_wantemp_system_sdwan_duplication" "trname" {
  dstaddr    = "all"
  dstintf    = "port4"
  fosid      = 1
  wanprof    = fortimanager_wan_template.trname.name
  depends_on = [fortimanager_wan_template.trname]
}

resource "fortimanager_wan_template" "trname" {
  name = "terr-template"
  adom = "root"
  type = "wanprof"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

* `dstaddr` - Destination address or address group names.
* `dstaddr6` - Destination address6 or address6 group names.
* `dstintf` - Outgoing (egress) interfaces or zones.
* `fosid` - Duplication rule ID (1 - 255).
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated. Valid values: `disable`, `enable`.

* `packet_duplication` - Configure packet duplication method. Valid values: `disable`, `force`, `on-demand`.

* `service` - Service and service group name.
* `service_id` - SD-WAN service rule ID list.
* `sla_match_service` - Enable/disable packet duplication matching health-check SLAs in service rule. Valid values: `disable`, `enable`.

* `srcaddr` - Source address or address group names.
* `srcaddr6` - Source address6 or address6 group names.
* `srcintf` - Incoming (ingress) interfaces or zones.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Wantemp SystemSdwanDuplication can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_sdwan_duplication.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
