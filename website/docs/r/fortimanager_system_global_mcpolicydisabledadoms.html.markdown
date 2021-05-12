---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_global_mcpolicydisabledadoms"
description: |-
  Multicast policy disabled adoms.
---

# fortimanager_system_global_mcpolicydisabledadoms
Multicast policy disabled adoms.

## Argument Reference


The following arguments are supported:


* `adom_name` - Adom names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System GlobalMcPolicyDisabledAdoms can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_global_mcpolicydisabledadoms.labelname SystemGlobalMcPolicyDisabledAdoms
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

