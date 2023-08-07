---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_shaper_trafficshaper"
description: |-
  Configure shared traffic shaper.
---

# fortimanager_object_firewall_shaper_trafficshaper
Configure shared traffic shaper.

## Example Usage

```hcl
resource "fortimanager_object_firewall_shaper_trafficshaper" "trname" {
  bandwidth_unit = "mbps"
  diffserv       = "disable"
  name           = "terraform"
  per_policy     = "disable"
  priority       = "medium"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `bandwidth_unit` - Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.

* `cos` - VLAN CoS mark.
* `cos_marking` - Enable/disable VLAN CoS marking. Valid values: `disable`, `enable`.

* `cos_marking_method` - Select VLAN CoS marking method. Valid values: `multi-stage`, `static`.

* `diffserv` - Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `disable`, `enable`.

* `diffservcode` - DiffServ setting to be applied to traffic accepted by this shaper.
* `dscp_marking_method` - Select DSCP marking method. Valid values: `multi-stage`, `static`.

* `exceed_bandwidth` - Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
* `exceed_class_id` - Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].
* `exceed_cos` - VLAN CoS mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
* `exceed_dscp` - DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
* `guaranteed_bandwidth` - Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
* `maximum_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `maximum_cos` - VLAN CoS mark for traffic in [exceed-bandwidth, maximum-bandwidth].
* `maximum_dscp` - DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
* `name` - Traffic shaper name.
* `overhead` - Per-packet size overhead used in rate computations.
* `per_policy` - Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.

* `priority` - Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `high`, `medium`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ShaperTrafficShaper can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_shaper_trafficshaper.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
