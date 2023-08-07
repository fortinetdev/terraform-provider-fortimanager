---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_swtrhash"
description: |-
  Configure switch traditional hashing.
---

# fortimanager_object_system_npu_swtrhash
Configure switch traditional hashing.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `draco15` - Enable/disable DRACO15 hashing. Valid values: `disable`, `enable`.

* `tcp_udp_port` - Include/exclude TCP/UDP source and destination port for unicast trunk traffic. Valid values: `include`, `exclude`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuSwTrHash can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_swtrhash.labelname ObjectSystemNpuSwTrHash
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
