---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_backgroundssescan"
description: |-
  Configure driver background scan for SSE.
---

# fortimanager_object_system_npu_backgroundssescan
Configure driver background scan for SSE.

~> This resource is a sub resource for variable `background_sse_scan` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `scan` - Enable/disable background SSE scan by driver thread(default enabled). Valid values: `disable`, `enable`.

* `scan_stale` - Configure scanning of active or stale sessions (default = 0 = active sessions).
* `scan_vt` - Select version/type to scan: bit-0: 44; bit-1: 46; bit-2: 64; bit-3: 66 (default = 0xF).
* `stats_qual_access` - Statistics update access qualification in seconds (0 - INT_MAX, default = 180).
* `stats_qual_duration` - Statistics update duration qualification in seconds (0 - INT_MAX, default = 300).
* `stats_update_interval` - Stats update interval(&gt;=5*60 seconds, default 5*60 seconds).
* `udp_keepalive_interval` - UDP keepalive interval(&gt;=90 seconds, default 90 seconds).
* `udp_qual_access` - UDP keepalive access qualification in seconds (0 - INT_MAX, default = 30).
* `udp_qual_duration` - UDP keepalive duration qualification in seconds (0 - INT_MAX, default = 90).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuBackgroundSseScan can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_backgroundssescan.labelname ObjectSystemNpuBackgroundSseScan
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
