---
subcategory: "System Global"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ha_peer"
description: |-
  Peer.
---

# fortimanager_system_ha_peer
Peer.

~> This resource is a sub resource for variable `peer` of resource `fortimanager_system_ha`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_ha_peer" "trname" {
  fosid         = "1"
  ip            = "11.11.11.5"
  serial_number = "FMG-VM0000000003"
  status        = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `fosid` - Id.
* `ip` - IP address of peer.
* `ip6` - IP address (V6) of peer.
* `serial_number` - Serial number of peer.
* `status` - Peer admin status. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaPeer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_ha_peer.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

