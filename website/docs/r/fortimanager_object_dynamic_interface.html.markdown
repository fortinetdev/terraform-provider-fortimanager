---
subcategory: "Object Dynamic"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dynamic_interface"
description: |-
  ObjectDynamic Interface
---

# fortimanager_object_dynamic_interface
ObjectDynamic Interface

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `color` - Color.
* `default_mapping` - Default-Mapping. Valid values: `disable`, `enable`.

* `defmap_intf` - Defmap-Intf.
* `defmap_intrazone_deny` - Defmap-Intrazone-Deny. Valid values: `disable`, `enable`.

* `defmap_zonemember` - Defmap-Zonemember.
* `description` - Description.
* `dynamic_mapping` - Dynamic_Mapping. The structure of `dynamic_mapping` block is documented below.
* `egress_shaping_profile` - Egress-Shaping-Profile.
* `ingress_shaping_profile` - Ingress-Shaping-Profile.
* `name` - Name.
* `platform_mapping` - Platform_Mapping. The structure of `platform_mapping` block is documented below.
* `single_intf` - Single-Intf. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_mapping` block supports:

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `egress_shaping_profile` - Egress-Shaping-Profile.
* `ingress_shaping_profile` - Ingress-Shaping-Profile.
* `intrazone_deny` - Intrazone-Deny. Valid values: `disable`, `enable`.

* `local_intf` - Local-Intf.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `platform_mapping` block supports:

* `egress_shaping_profile` - Egress-Shaping-Profile.
* `ingress_shaping_profile` - Ingress-Shaping-Profile.
* `intf_zone` - Intf-Zone.
* `intrazone_deny` - Intrazone-Deny. Valid values: `disable`, `enable`.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDynamic Interface can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dynamic_interface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
