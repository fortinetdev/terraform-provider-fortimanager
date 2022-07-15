---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ha_monitoredips"
description: |-
  Monitored IP addresses.
---

# fortimanager_system_ha_monitoredips
Monitored IP addresses.

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

