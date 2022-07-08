---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_ha_monitoredinterfaces"
description: |-
  Monitored interfaces.
---

# fortimanager_system_ha_monitoredinterfaces
Monitored interfaces.

## Argument Reference


The following arguments are supported:


* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System HaMonitoredInterfaces can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_ha_monitoredinterfaces.labelname SystemHaMonitoredInterfaces
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

