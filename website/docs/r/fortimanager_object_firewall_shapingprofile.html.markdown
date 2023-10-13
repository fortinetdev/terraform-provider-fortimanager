---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_shapingprofile"
description: |-
  Configure shaping profiles.
---

# fortimanager_object_firewall_shapingprofile
Configure shaping profiles.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Comment.
* `default_class_id` - Default class ID to handle unclassified packets (including all local traffic).
* `profile_name` - Shaping profile name.
* `shaping_entries` - Shaping-Entries. The structure of `shaping_entries` block is documented below.
* `type` - Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `shaping_entries` block supports:

* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `class_id` - Class ID.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwith in percentage.
* `id` - ID number.
* `limit` - Hard limit on the real queue size in packets.
* `max` - Average queue size in packets at which RED drop probability is maximal.
* `maximum_bandwidth_percentage` - Maximum bandwith in percentage.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `priority` - Priority. Valid values: `low`, `medium`, `high`, `critical`, `top`.

* `red_probability` - Maximum probability (in percentage) for RED marking.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_name}}.

## Import

ObjectFirewall ShapingProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_shapingprofile.labelname {{profile_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
