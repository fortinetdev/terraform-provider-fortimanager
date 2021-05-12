---
subcategory: "ObjectGtp"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_gtp_iewhitelist"
description: |-
  IE white list.
---

# fortimanager_object_gtp_iewhitelist
IE white list.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `entries` - Entries. The structure of `entries` block is documented below.
* `name` - IE white list name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `id` - Entry ID.
* `ie` - IE ID (1 - 255).
* `message` - Message ID (1 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectGtp IeWhiteList can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_gtp_iewhitelist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
