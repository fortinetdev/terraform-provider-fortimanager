---
subcategory: "Object Dynamic"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dynamic_address"
description: |-
  ObjectDynamic Address
---

# fortimanager_object_dynamic_address
ObjectDynamic Address

## Example Usage

```hcl
resource "fortimanager_object_dynamic_address" "trname" {
  default     = "all"
  description = "This is a Terraform example"
  name        = "terr-dynamic-address"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `default` - Default.
* `description` - Description.
* `dynamic_addr_mapping` - Dynamic_Addr_Mapping. The structure of `dynamic_addr_mapping` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dynamic_addr_mapping` block supports:

* `addr` - Addr.
* `id` - Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectDynamic Address can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dynamic_address.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
