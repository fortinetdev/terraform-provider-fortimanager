---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_qosmap"
description: |-
  Configure QoS map set.
---

# fortimanager_object_wirelesscontroller_hotspot20_qosmap
Configure QoS map set.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`dscp_except`: `fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscpexcept`
`dscp_range`: `fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscprange`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_qosmap" "labelname" {
  name = "ss"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `dscp_except` - Dscp-Except. The structure of `dscp_except` block is documented below.
* `dscp_range` - Dscp-Range. The structure of `dscp_range` block is documented below.
* `name` - QOS-MAP name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dscp_except` block supports:

* `dscp` - DSCP value.
* `index` - DSCP exception index.
* `up` - User priority.

The `dscp_range` block supports:

* `high` - DSCP high value.
* `index` - DSCP range index.
* `low` - DSCP low value.
* `up` - User priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20QosMap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_qosmap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
