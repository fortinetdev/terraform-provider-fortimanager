---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dynamic_address_dynamic_addr_mapping"
description: |-
  ObjectDynamic AddressDynamicAddrMapping
---

# fortimanager_object_dynamic_address_dynamic_addr_mapping
ObjectDynamic AddressDynamicAddrMapping

~> This resource is a sub resource for variable `dynamic_addr_mapping` of resource `fortimanager_object_dynamic_address`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_dynamic_address_dynamic_addr_mapping" "trname" {
  address    = fortimanager_object_dynamic_address.trname.name
  addr       = "34.56.1.5"
  fosid      = 1
  depends_on = [fortimanager_object_dynamic_address.trname]
}

resource "fortimanager_object_dynamic_address" "trname" {
  name = "terr-dynamic-address"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `address` - Address.

* `addr` - Addr.
* `fosid` - Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDynamic AddressDynamicAddrMapping can be imported using any of these accepted formats:
```
Set import_options = ["address=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dynamic_address_dynamic_addr_mapping.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
