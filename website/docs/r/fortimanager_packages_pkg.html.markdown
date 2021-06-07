---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_pkg"
description: |-
  Packages Pkg
---

# fortimanager_packages_pkg
Packages Pkg

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Name.
* `objver` - Obj Ver.
* `oid` - Oid.
* `packagesetting` - Package Setting. The structure of `packagesetting` block is documented below.
* `scopemember` - Scope Member. The structure of `scopemember` block is documented below.
* `subobj` - Subobj. The structure of `subobj` block is documented below.
* `type` - Type. Valid values: `pkg`, `folder`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `packagesetting` block supports:

* `central_nat` - Central-Nat. Valid values: `disable`, `enable`.

* `consolidated_firewall_mode` - Consolidated-Firewall-Mode. Valid values: `disable`, `enable`.

* `fwpolicy_implicit_log` - Fwpolicy-Implicit-Log. Valid values: `disable`, `enable`.

* `fwpolicy6_implicit_log` - Fwpolicy6-Implicit-Log. Valid values: `disable`, `enable`.

* `inspection_mode` - Inspection-Mode. Valid values: `proxy`, `flow`.

* `ngfw_mode` - Ngfw-Mode. Valid values: `profile-based`, `policy-based`.

* `ssl_ssh_profile` - Ssl-Ssh-Profile.

The `scopemember` block supports:

* `name` - Name.
* `vdom` - Vdom.

The `subobj` block supports:

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Packages Pkg can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_pkg.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
