---
subcategory: "System Route"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_route"
description: |-
  Routing table configuration.
---

# fortimanager_system_route
Routing table configuration.

## Example Usage

```hcl
resource "fortimanager_system_route" "trname" {
  device = "port1"
  dst = [
    "222.222.22.0",
    "255.255.255.0"
  ]
  gateway = "222.222.222.1"
  seq_num = 7
}
```

## Argument Reference


The following arguments are supported:


* `device` - Gateway out interface.
* `dst` - Destination IP and mask for this route.
* `gateway` - Gateway IP for this route.
* `seq_num` - Entry number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

System Route can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_route.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

