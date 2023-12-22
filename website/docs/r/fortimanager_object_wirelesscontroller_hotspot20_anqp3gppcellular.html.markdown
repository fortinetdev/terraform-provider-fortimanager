---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular"
description: |-
  Configure 3GPP public land mobile network (PLMN).
---

# fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular
Configure 3GPP public land mobile network (PLMN).

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`mcc_mnc_list`: `fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular" "trname" {
  mcc_mnc_list {
    id  = 1
    mcc = "mcc"
    mnc = "mnc"
  }
  name = "terr-wictl-hot20-anqp-3gpp-cellular"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `mcc_mnc_list` - Mcc-Mnc-List. The structure of `mcc_mnc_list` block is documented below.
* `name` - 3GPP PLMN name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mcc_mnc_list` block supports:

* `id` - ID.
* `mcc` - Mobile country code.
* `mnc` - Mobile network code.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20Anqp3GppCellular can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_anqp3gppcellular.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
