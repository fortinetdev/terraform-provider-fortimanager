---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_topology"
description: |-
  Logging topology settings.
---

# fortimanager_system_log_topology
Logging topology settings.

## Argument Reference


The following arguments are supported:


* `max_depth` - Maximum descend levels below this device.
* `max_depth_share` - Maximum descend levels below this device to share with upstream.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogTopology can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_topology.labelname SystemLogTopology
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

