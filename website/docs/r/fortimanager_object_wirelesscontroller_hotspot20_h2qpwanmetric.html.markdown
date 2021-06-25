---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qpwanmetric"
description: |-
  Configure WAN metrics.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qpwanmetric
Configure WAN metrics.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_h2qpwanmetric" "trname" {
  downlink_load             = 1
  downlink_speed            = 2400
  link_at_capacity          = "disable"
  link_status               = "down"
  load_measurement_duration = 10
  name                      = "terr-wictl-hot20-heqp-wan-metric"
  symmetric_wan_link        = "asymmetric"
  uplink_speed              = 2400
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `downlink_load` - Downlink load.
* `downlink_speed` - Downlink speed (in kilobits/s).
* `link_at_capacity` - Link at capacity. Valid values: `disable`, `enable`.

* `link_status` - Link status. Valid values: `down`, `up`, `in-test`.

* `load_measurement_duration` - Load measurement duration (in tenths of a second).
* `name` - WAN metric name.
* `symmetric_wan_link` - WAN link symmetry. Valid values: `asymmetric`, `symmetric`.

* `uplink_load` - Uplink load.
* `uplink_speed` - Uplink speed (in kilobits/s).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpWanMetric can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qpwanmetric.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
