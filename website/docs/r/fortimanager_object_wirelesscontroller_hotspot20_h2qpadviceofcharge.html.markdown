---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge"
description: |-
  Configure advice of charge.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge
Configure advice of charge.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`aoc_list`: `fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge_aoclist`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge" "trname" {
  name = "terr-h2qpadviceofcharge"
  aoc_list {
    name = "terr-h2qpadviceofcharge"
    plan_info {
      name = "terr-h2qpadviceofcharge"
    }
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `aoc_list` - Aoc-List. The structure of `aoc_list` block is documented below.
* `name` - Plan name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `aoc_list` block supports:

* `nai_realm` - NAI realm list name.
* `nai_realm_encoding` - NAI realm encoding.
* `name` - Advice of charge ID.
* `plan_info` - Plan-Info. The structure of `plan_info` block is documented below.
* `type` - Usage charge type. Valid values: `time-based`, `volume-based`, `time-and-volume-based`, `unlimited`.


The `plan_info` block supports:

* `currency` - Currency code.
* `info_file` - Info file.
* `lang` - Language code.
* `name` - Plan name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpAdviceOfCharge can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qpadviceofcharge.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
