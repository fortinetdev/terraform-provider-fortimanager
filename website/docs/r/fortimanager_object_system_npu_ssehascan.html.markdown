---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_ssehascan"
description: |-
  Configure driver HA scan for SSE.
---

# fortimanager_object_system_npu_ssehascan
Configure driver HA scan for SSE.

~> This resource is a sub resource for variable `sse_ha_scan` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_system_npu_ssehascan" "trname" {
  gap             = 7000
  max_session_cnt = 100
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `gap` - Scanning message gap(0~32767, default 6000)
* `max_session_cnt` - If the session count(in millions) is larger than this, HA scan will be skipped. (0~0xffff, default 0xffff)
* `min_duration` - Scanning filter for minimum duration of the session. (0~3600, default 0)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuSseHaScan can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_ssehascan.labelname ObjectSystemNpuSseHaScan
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
