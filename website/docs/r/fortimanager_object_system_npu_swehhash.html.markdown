---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_swehhash"
description: |-
  Configure switch enhanced hashing.
---

# fortimanager_object_system_npu_swehhash
Configure switch enhanced hashing.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `computation` - Set hashing computation. Valid values: `xor16`, `xor8`, `xor4`, `crc16`.

* `destination_ip_lower_16` - Include/exclude destination IP address lower 16 bits. Valid values: `include`, `exclude`.

* `destination_ip_upper_16` - Include/exclude destination IP address upper 16 bits. Valid values: `include`, `exclude`.

* `destination_port` - Include/exclude destination port if TCP/UDP. Valid values: `include`, `exclude`.

* `ip_protocol` - Include/exclude IP protocol. Valid values: `include`, `exclude`.

* `netmask_length` - Network mask length.
* `source_ip_lower_16` - Include/exclude source IP address lower 16 bits. Valid values: `include`, `exclude`.

* `source_ip_upper_16` - Include/exclude source IP address upper 16 bits. Valid values: `include`, `exclude`.

* `source_port` - Include/exclude source port if TCP/UDP. Valid values: `include`, `exclude`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuSwEhHash can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_swehhash.labelname ObjectSystemNpuSwEhHash
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
