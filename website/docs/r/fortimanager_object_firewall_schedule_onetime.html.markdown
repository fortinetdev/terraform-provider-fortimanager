---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_schedule_onetime"
description: |-
  Onetime schedule configuration.
---

# fortimanager_object_firewall_schedule_onetime
Onetime schedule configuration.

## Example Usage

```hcl
resource "fortimanager_object_firewall_schedule_onetime" "trname" {
  color = 1
  end   = "15:00 2020/08/17"
  name  = "terraform-tefv-onetime"
  start = "07:00 2020/08/17"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Color of icon on the GUI.
* `end` - Schedule end date and time, format hh:mm yyyy/mm/dd.
* `expiration_days` - Write an event log message this many days before the schedule expires.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `global_object` - Global Object.
* `name` - Onetime schedule name.
* `start` - Schedule start date and time, format hh:mm yyyy/mm/dd.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectFirewall ScheduleOnetime can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_schedule_onetime.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
