---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscprange"
description: |-
  Differentiated Services Code Point (DSCP) ranges.
---

# fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscprange
Differentiated Services Code Point (DSCP) ranges.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `qos_map` - Qos Map.

* `high` - DSCP high value.
* `index` - DSCP range index.
* `low` - DSCP low value.
* `up` - User priority.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

ObjectWirelessController Hotspot20QosMapDscpRange can be imported using any of these accepted formats:
```
Set import_options = ["qos_map=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_qosmap_dscprange.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
