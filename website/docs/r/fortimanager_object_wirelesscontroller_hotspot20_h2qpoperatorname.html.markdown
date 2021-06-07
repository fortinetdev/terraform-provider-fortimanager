---
subcategory: "Object Wireless-Controller"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qpoperatorname"
description: |-
  Configure operator friendly name.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qpoperatorname
Configure operator friendly name.

## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_h2qpoperatorname" "labelname" {
  name = "ss"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Friendly name ID.
* `value_list` - Value-List. The structure of `value_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `value_list` block supports:

* `index` - Value index.
* `lang` - Language code.
* `value` - Friendly name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpOperatorName can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qpoperatorname.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
