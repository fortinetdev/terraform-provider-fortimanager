---
subcategory: "System SNMP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_snmp_sysinfo"
description: |-
  SNMP configuration.
---

# fortimanager_system_snmp_sysinfo
SNMP configuration.

## Argument Reference


The following arguments are supported:


* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engineID string (maximum 24 characters).
* `location` - System location.
* `status` - Enable/disable SNMP. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `trap_cpu_high_exclude_nice_threshold` - SNMP trap for CPU usage threshold (exclude NICE processes).
* `trap_high_cpu_threshold` - SNMP trap for CPU usage threshold.
* `trap_low_memory_threshold` - SNMP trap for memory usage threshold.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SnmpSysinfo can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_snmp_sysinfo.labelname SystemSnmpSysinfo
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

