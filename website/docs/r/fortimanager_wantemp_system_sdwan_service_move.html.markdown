---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_wantemp_system_sdwan_service_move"
description: |-
  Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.
---

# fortimanager_wantemp_system_sdwan_service_move
Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

## Example Usage

```hcl
resource "fortimanager_wantemp_system_sdwan_service_move" "trname" {
  wanprof    = fortimanager_wan_template.trname.name
  service    = fortimanager_wantemp_system_sdwan_service.trname2.fosid
  target     = fortimanager_wantemp_system_sdwan_service.trname.fosid
  option     = "after"
  depends_on = [fortimanager_wantemp_system_sdwan_service.trname2, fortimanager_wantemp_system_sdwan_service.trname]
}

resource "fortimanager_wantemp_system_sdwan_service" "trname2" {
  wanprof          = fortimanager_wan_template.trname.name
  fosid            = 2
  priority_members = 3
  dst              = "all"
  depends_on       = [fortimanager_wantemp_system_sdwan_service.trname2]
}

resource "fortimanager_wantemp_system_sdwan_service" "trname" {
  wanprof          = fortimanager_wan_template.trname.name
  fosid            = 1
  priority_members = 3
  dst              = "all"
  depends_on       = [fortimanager_wantemp_system_sdwan_service.trname]
}

resource "fortimanager_wantemp_system_sdwan_members" "trname2" {
  wanprof    = fortimanager_wan_template.trname.name
  cost       = 2
  interface  = "port2"
  seq_num    = 3
  depends_on = [fortimanager_wan_template.trname]
}

resource "fortimanager_wantemp_system_sdwan_members" "trname" {
  wanprof    = fortimanager_wan_template.trname.name
  cost       = 1
  interface  = "port7"
  seq_num    = 2
  depends_on = [fortimanager_wan_template.trname]
}

resource "fortimanager_wan_template" "trname" {
  name = "terr4"
  adom = "root"
  type = "wanprof"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `wanprof` - Wanprof.
* `service` - Service.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the service changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of services.
