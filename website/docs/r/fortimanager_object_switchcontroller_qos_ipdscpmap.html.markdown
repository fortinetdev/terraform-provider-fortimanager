---
subcategory: "Object Switch-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_qos_ipdscpmap"
description: |-
  Configure FortiSwitch QoS IP precedence/DSCP.
---

# fortimanager_object_switchcontroller_qos_ipdscpmap
Configure FortiSwitch QoS IP precedence/DSCP.

## Example Usage

```hcl
resource "fortimanager_object_switchcontroller_qos_ipdscpmap" "trname" {
  name = "vp1"

  map {
    cos_queue = 1
    name      = "1"
    value     = "46"
  }
  map {
    cos_queue = 2
    name      = "2"
    value     = "24,26,48,56"
  }
  map {
    cos_queue = 3
    name      = "5"
    value     = "34"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `description` - Description of the ip-dscp map name.
* `map` - Map. The structure of `map` block is documented below.
* `name` - Dscp map name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `map` block supports:

* `cos_queue` - COS queue number.
* `diffserv` - Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.

* `ip_precedence` - IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.

* `name` - Dscp mapping entry name.
* `value` - Raw values of DSCP (0 - 63).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController QosIpDscpMap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_qos_ipdscpmap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
