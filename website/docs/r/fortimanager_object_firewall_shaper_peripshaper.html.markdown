---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_shaper_peripshaper"
description: |-
  Configure per-IP traffic shaper.
---

# fortimanager_object_firewall_shaper_peripshaper
Configure per-IP traffic shaper.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `bandwidth_unit` - Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.

* `diffserv_forward` - Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
* `diffservcode_rev` - Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
* `max_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `max_concurrent_session` - Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `name` - Traffic shaper name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ShaperPerIpShaper can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_shaper_peripshaper.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
