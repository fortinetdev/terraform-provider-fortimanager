---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo"
description: |-
  Plan info.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo
Plan info.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `h2qp_advice_of_charge` - H2Qp Advice Of Charge.
* `aoc_list` - Aoc List.

* `currency` - Currency code.
* `info_file` - Info file.
* `lang` - Language code.
* `name` - Plan name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpAdviceOfChargeAocListPlanInfo can be imported using any of these accepted formats:
```
Set import_options = ["h2qp_advice_of_charge=YOUR_VALUE", "aoc_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist_planinfo.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
