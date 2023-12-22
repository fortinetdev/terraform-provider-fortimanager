---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_policy6_move"
description: |-
  Move IPv6 policies.
---

# fortimanager_packages_firewall_policy6_move
Move IPv6 policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_policy6_move" "trname" {
  pkg        = "default"
  policy6    = "2"
  target     = "3"
  option     = "before"
  depends_on = [fortimanager_packages_firewall_policy6.trname2, fortimanager_packages_firewall_policy6.trname3]
}

resource "fortimanager_packages_firewall_policy6" "trname2" {
  dstaddr  = "all"
  dstintf  = "any"
  pkg      = "default"
  policyid = 2
  schedule = "always"
  service  = "ALL"
  srcaddr  = "all"
  srcintf  = "any"
}

resource "fortimanager_packages_firewall_policy6" "trname3" {
  dstaddr  = "all"
  dstintf  = "any"
  pkg      = "default"
  policyid = 3
  schedule = "always"
  service  = "ALL"
  srcaddr  = "all"
  srcintf  = "any"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.
* `policy6` - Policy6.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the policy6 changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of policy6s.
