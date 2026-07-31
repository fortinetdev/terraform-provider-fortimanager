---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_schedule_recurring"
description: |-
  Recurring schedule configuration.
---

# fortimanager_object_firewall_schedule_recurring
Recurring schedule configuration.

## Example Usage

```hcl
resource "fortimanager_object_firewall_schedule_recurring" "trname" {
  color = 1
  day   = ["sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "none"]
  end   = "15:00"
  name  = "terraform-tefv-recurring"
  start = "07:00"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Color of icon on the GUI.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.

* `end` - Time of day to end the schedule, format hh:mm.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `label_day` - Configure a window during the time of day in which the schedule job is executed. Valid values: `none`, `over-night`, `early-morning`, `morning`, `midday`, `afternoon`, `evening`, `night`, `late-night`.

* `global_object` - Global Object.
* `name` - Recurring schedule name.
* `start` - Time of day to start the schedule, format hh:mm.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ScheduleRecurring can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_schedule_recurring.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
