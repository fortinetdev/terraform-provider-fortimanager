---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_central_dnat"
description: |-
  Packages CentralDnat
---

# fortimanager_packages_central_dnat
Packages CentralDnat

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Packages CentralDnat can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_central_dnat.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
