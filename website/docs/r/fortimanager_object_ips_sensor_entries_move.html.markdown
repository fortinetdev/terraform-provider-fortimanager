---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ips_sensor_entries_move"
description: |-
  IPS sensor filter.
---

# fortimanager_object_ips_sensor_entries_move
IPS sensor filter.

## Example Usage

```hcl
resource "fortimanager_object_ips_sensor_entries_move" "trname" {
  sensor     = fortimanager_object_ips_sensor.trname.name
  entries    = fortimanager_object_ips_sensor_entries.trname.fosid
  target     = fortimanager_object_ips_sensor_entries.trname2.fosid
  option     = "after"
  depends_on = [fortimanager_object_ips_sensor_entries.trname, fortimanager_object_ips_sensor_entries.trname2]
}

resource "fortimanager_object_ips_sensor_entries" "trname" {
  sensor     = fortimanager_object_ips_sensor.trname.name
  fosid      = 23
  depends_on = [fortimanager_object_ips_sensor.trname]
}

resource "fortimanager_object_ips_sensor_entries" "trname2" {
  sensor     = fortimanager_object_ips_sensor.trname.name
  fosid      = 24
  depends_on = [fortimanager_object_ips_sensor.trname]
}

resource "fortimanager_object_ips_sensor" "trname" {
  name = "terr-sensor"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `sensor` - Sensor.
* `entries` - Entries.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the entries changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of entry.
