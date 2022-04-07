---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_tcptimeoutprofile"
description: |-
  Configure TCP timeout profile.
---

# fortimanager_object_system_npu_tcptimeoutprofile
Configure TCP timeout profile.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `close_wait` - Set close-wait timeout(seconds)
* `fin_wait` - Set fin-wait timeout(seconds)
* `fosid` - Timeout profile ID (5-47)
* `syn_sent` - Set syn-sent timeout(seconds)
* `syn_wait` - Set syn-wait timeout(seconds)
* `tcp_idle` - Set TCP establish timeout(seconds)
* `time_wait` - Set time-wait timeout(seconds)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSystem NpuTcpTimeoutProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_tcptimeoutprofile.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
