---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_pkg"
description: |-
  Packages Pkg
---

# fortimanager_packages_pkg
Packages Pkg

## Example Usage

```hcl
resource "fortimanager_packages_pkg" "trname" {
  name = "terr-pkg"
  type = "pkg"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.

* `name` - Name.
* `objver` - Obj Ver.
* `oid` - Oid.
* `packagesettings` - Package Settings. The structure of `packagesettings` block is documented below.
* `packagesetting` - Package Setting. The structure of `packagesetting` block is documented below.
* `scopemember` - Scope Member. The structure of `scopemember` block is documented below.
* `type` - Type. Valid values: `pkg`, `folder`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `packagesettings` block supports:

* `central_nat` - Central-Nat. Valid values: `disable`, `enable`.

* `consolidated_firewall_mode` - Consolidated-Firewall-Mode. Valid values: `disable`, `enable`.

* `fwpolicy_implicit_log` - Fwpolicy-Implicit-Log. Valid values: `disable`, `enable`.

* `fwpolicy6_implicit_log` - Fwpolicy6-Implicit-Log. Valid values: `disable`, `enable`.

* `inspection_mode` - Inspection-Mode. Valid values: `proxy`, `flow`.

* `ngfw_mode` - Ngfw-Mode. Valid values: `profile-based`, `policy-based`.

* `policy_offload_level` - Policy-Offload-Level. Valid values: `disable`, `default`, `dos-offload`, `full-offload`.

* `ssl_ssh_profile` - Ssl-Ssh-Profile.

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.
* `output_folder_path` - Folder path of the package..
* `output_pkg_name` - Package name. The value will be empty string if the type of the resource is `folder`..

## Import

Packages Pkg can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_pkg.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
