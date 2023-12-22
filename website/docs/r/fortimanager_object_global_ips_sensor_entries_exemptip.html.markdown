---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_global_ips_sensor_entries_exemptip"
description: |-
  Traffic from selected source or destination IP addresses is exempt from this signature.
---

# fortimanager_object_global_ips_sensor_entries_exemptip
Traffic from selected source or destination IP addresses is exempt from this signature.

~> This resource is a sub resource for variable `exempt_ip` of resource `fortimanager_object_global_ips_sensor_entries`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sensor` - Sensor.
* `entries` - Entries.

* `dst_ip` - Destination IP address and netmask (applies to packet matching the signature).
* `fosid` - Exempt IP ID.
* `src_ip` - Source IP address and netmask (applies to packet matching the signature).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectGlobal IpsSensorEntriesExemptIp can be imported using any of these accepted formats:
```
Set import_options = ["sensor=YOUR_VALUE", "entries=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_global_ips_sensor_entries_exemptip.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
