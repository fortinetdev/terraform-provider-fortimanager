---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ha_monitoredips"
description: |-
  Monitored IP addresses.
---

# fortimanager_system_ha_monitoredips
Monitored IP addresses.

~> This resource is a sub resource for variable `monitored_ips` of resource `fortimanager_system_ha`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_ha_monitoredips" "trname" {
  fosid     = 1
  interface = "port4"
  ip        = "1.2.3.4"
}
```

## Argument Reference


The following arguments are supported:


* `fosid` - Id.
* `interface` - Interface name.
* `ip` - IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System HaMonitoredIps can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_ha_monitoredips.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

