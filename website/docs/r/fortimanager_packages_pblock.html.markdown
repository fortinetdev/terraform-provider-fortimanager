---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_pblock"
description: |-
  Packages Pblock
---

# fortimanager_packages_pblock
Packages Pblock

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`firewall_consolidated_policy`: `fortimanager_packages_pblock_firewall_consolidated_policy`
`firewall_policy`: `fortimanager_packages_pblock_firewall_policy`
`firewall_policy6`: `fortimanager_packages_pblock_firewall_policy6`
`firewall_security_policy`: `fortimanager_packages_pblock_firewall_securitypolicy`



## Example Usage

```hcl
resource "fortimanager_packages_pblock" "trname" {
  name = "terr-pblock"
  type = "pblock"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Name.
* `oid` - Oid.
* `packagesettings` - Package Settings. The structure of `packagesettings` block is documented below.
* `type` - Type. Valid values: `pblock`.


The `packagesettings` block supports:

* `central_nat` - Central-Nat. Valid values: `disable`, `enable`.

* `consolidated_firewall_mode` - Consolidated-Firewall-Mode. Valid values: `disable`, `enable`.

* `fwpolicy_implicit_log` - Fwpolicy-Implicit-Log. Valid values: `disable`, `enable`.

* `fwpolicy6_implicit_log` - Fwpolicy6-Implicit-Log. Valid values: `disable`, `enable`.

* `inspection_mode` - Inspection-Mode. Valid values: `proxy`, `flow`.

* `ngfw_mode` - Ngfw-Mode. Valid values: `profile-based`, `policy-based`.

* `policy_offload_level` - Policy-Offload-Level. Valid values: `disable`, `default`, `dos-offload`, `full-offload`.

* `ssl_ssh_profile` - Ssl-Ssh-Profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Packages Pblock can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_pblock.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
