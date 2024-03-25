---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_global_ips_sensor_override"
description: |-
  ObjectGlobal IpsSensorOverride
---

# fortimanager_object_global_ips_sensor_override
ObjectGlobal IpsSensorOverride

~> This resource is a sub resource for variable `override` of resource `fortimanager_object_global_ips_sensor`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `exempt_ip`: `fortimanager_object_global_ips_sensor_override_exemptip`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sensor` - Sensor.

* `action` - Action. Valid values: `pass`, `block`, `reset`.

* `exempt_ip` - Exempt-Ip. The structure of `exempt_ip` block is documented below.
* `log` - Log. Valid values: `disable`, `enable`.

* `log_packet` - Log-Packet. Valid values: `disable`, `enable`.

* `quarantine` - Quarantine. Valid values: `none`, `attacker`, `both`, `interface`.

* `quarantine_expiry` - Quarantine-Expiry.
* `quarantine_log` - Quarantine-Log. Valid values: `disable`, `enable`.

* `rule_id` - Rule-Id.
* `status` - Status. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `exempt_ip` block supports:

* `dst_ip` - Dst-Ip.
* `id` - Id.
* `src_ip` - Src-Ip.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{rule_id}}.

## Import

ObjectGlobal IpsSensorOverride can be imported using any of these accepted formats:
```
Set import_options = ["sensor=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_global_ips_sensor_override.labelname {{rule_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
