---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dynamic_certificate_local_dynamic_mapping"
description: |-
  ObjectDynamic CertificateLocalDynamicMapping
---

# fortimanager_object_dynamic_certificate_local_dynamic_mapping
ObjectDynamic CertificateLocalDynamicMapping

~> This resource is a sub resource for variable `dynamic_mapping` of resource `fortimanager_object_dynamic_certificate_local`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `local` - Local.

* `_scope` - _Scope. The structure of `_scope` block is documented below.
* `local_cert` - Local-Cert.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `_scope` block supports:

* `name` - Name.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format "{{_scope.name}} {{_scope.vdom}}".

## Import

ObjectDynamic CertificateLocalDynamicMapping can be imported using any of these accepted formats:
```
Set import_options = ["local=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dynamic_certificate_local_dynamic_mapping.labelname {{_scope.name}}.{{_scope.vdom}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
