---
subcategory: "System Route"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_route6"
description: |-
  Routing table configuration.
---

# fortimanager_system_route6
Routing table configuration.

## Example Usage

```hcl
resource "fortimanager_system_route6" "trname" {
  device  = "port1"
  dst     = ["2001::/96"]
  gateway = "2001::2"
  prio    = 23
}
```

## Argument Reference


The following arguments are supported:


* `device` - Gateway out interface.
* `dst` - Destination IP and mask for this route.
* `gateway` - Gateway IP for this route.
* `prio` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{prio}}.

## Import

System Route6 can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_route6.labelname {{prio}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

