---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_acl_move"
description: |-
  Move IPv4 access control list.
---

# fortimanager_packages_firewall_acl_move
Move IPv4 access control list.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_acl" "trname1" {
  pkg       = "default"
  name      = "acl_policy1"
  interface = "port4"
  dstaddr   = ["all"]
  srcaddr   = ["all"]
  service   = ["ALL"]
  policyid  = 4
}

resource "fortimanager_packages_firewall_acl" "trname2" {
  pkg       = "default"
  name      = "acl_policy2"
  interface = "port4"
  dstaddr   = ["all"]
  srcaddr   = ["all"]
  service   = ["ALL"]
  policyid  = 5
}

resource "fortimanager_packages_firewall_acl_move" "trname" {
  pkg        = "default"
  acl        = fortimanager_packages_firewall_acl.trname1.name
  target     = fortimanager_packages_firewall_acl.trname2.name
  option     = "after"
  depends_on = [fortimanager_packages_firewall_acl.trname1, fortimanager_packages_firewall_acl.trname2]
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.
* `acl` - Acl.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the acl changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of acls.
