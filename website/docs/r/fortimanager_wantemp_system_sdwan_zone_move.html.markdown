---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_zone_move"
description: |-
  Move SD-WAN zones.
---

# fortimanager_wantemp_system_sdwan_zone_move
Move SD-WAN zones.

## Example Usage

```hcl
resource "fortimanager_wantemp_system_sdwan_zone_move" "trname" {
  wanprof = fortimanager_wan_template.trname2.name
  zone    = fortimanager_wantemp_system_sdwan_zone.trname.name
  target  = fortimanager_wantemp_system_sdwan_zone.trname2.name
  option  = "after"
}

resource "fortimanager_wantemp_system_sdwan_zone" "trname2" {
  name                  = "zone2"
  service_sla_tie_break = "cfg-order"
  wanprof               = fortimanager_wan_template.trname2.name
  depends_on            = [fortimanager_wan_template.trname2]
}

resource "fortimanager_wantemp_system_sdwan_zone" "trname" {
  name                  = "zone"
  service_sla_tie_break = "cfg-order"
  wanprof               = fortimanager_wan_template.trname2.name
  depends_on            = [fortimanager_wan_template.trname2]
}

resource "fortimanager_wan_template" "trname2" {
  name = "terr15"
  adom = "root"
  type = "wanprof"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.
* `zone` - Zone.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the zone changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of zones.
