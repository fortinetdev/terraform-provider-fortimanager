---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_switchcontroller_dsl_policy"
description: |-
  DSL policy.
---

# fortimanager_object_switchcontroller_dsl_policy
DSL policy.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `append_padding` - device pause frame configuration. Valid values: `disable`, `enable`.

* `cpe_aele` - cpe AELE. Valid values: `disable`, `enable`.

* `cpe_aele_mode` - cpe AELE-Mode with given string. Valid values: `ELE_M0`, `ELE_DS`, `ELE_PB`, `ELE_MIN`.

* `cs` - CPE carrier set. Valid values: `A43`, `B43`, `A43C`, `V43`.

* `ds_bitswap` - Enable/disable bitswap. Valid values: `disable`, `enable`.

* `name` - Policy name.
* `pause_frame` - device pause frame configuration. Valid values: `disable`, `enable`.

* `profile` - vdsl CPE profile. Valid values: `auto-30a`, `auto-17a`, `auto-12ab`.

* `type` - Type. Valid values: `Procend`, `Proscend`.

* `us_bitswap` - Enable/disable bitswap. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSwitchController DslPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_switchcontroller_dsl_policy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
