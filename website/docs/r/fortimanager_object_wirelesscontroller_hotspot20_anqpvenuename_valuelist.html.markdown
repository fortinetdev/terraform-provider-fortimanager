---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_anqpvenuename_valuelist"
description: |-
  Name list.
---

# fortimanager_object_wirelesscontroller_hotspot20_anqpvenuename_valuelist
Name list.

~> This resource is a sub resource for variable `value_list` of resource `fortimanager_object_wirelesscontroller_hotspot20_anqpvenuename`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `anqp_venue_name` - Anqp Venue Name.

* `index` - Value index.
* `lang` - Language code.
* `value` - Venue name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

ObjectWirelessController Hotspot20AnqpVenueNameValueList can be imported using any of these accepted formats:
```
Set import_options = ["anqp_venue_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_anqpvenuename_valuelist.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
