---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai"
description: |-
  Configure online sign up (OSU) provider NAI list.
---

# fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai
Configure online sign up (OSU) provider NAI list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `nai_list`: `fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai_nailist`



## Example Usage

```hcl
resource "fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai" "trname" {
  name = "terr-wictl-hot20"
  nai_list {
    name = "terr-h2qposuprovidernai"
  }
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `nai_list` - Nai-List. The structure of `nai_list` block is documented below.
* `name` - OSU provider NAI ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `nai_list` block supports:

* `name` - OSU NAI ID.
* `osu_nai` - OSU NAI.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectWirelessController Hotspot20H2QpOsuProviderNai can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_wirelesscontroller_hotspot20_h2qposuprovidernai.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
