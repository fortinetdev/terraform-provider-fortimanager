---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_zone"
description: |-
  Configure SD-WAN zones.
---

# fortimanager_wantemp_system_sdwan_zone
Configure SD-WAN zones.

~> This resource is a sub resource for variable `zone` of resource `fortimanager_wantemp_system_sdwan`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_wantemp_system_sdwan_zone" "trname" {
  name                  = "zone"
  service_sla_tie_break = "cfg-order"
  wanprof               = fortimanager_wan_template.trname.name
  depends_on            = [fortimanager_wan_template.trname]
}

resource "fortimanager_wan_template" "trname" {
  name = "template"
  adom = "root"
  type = "wanprof"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.

* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `cfg-order`, `fib-best-match`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wantemp SystemSdwanZone can be imported using any of these accepted formats:
```
Set import_options = ["wanprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_wantemp_system_sdwan_zone.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
