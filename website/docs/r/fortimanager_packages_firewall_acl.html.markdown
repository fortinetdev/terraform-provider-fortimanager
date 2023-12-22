---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_acl"
description: |-
  Configure IPv4 access control list.
---

# fortimanager_packages_firewall_acl
Configure IPv4 access control list.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_acl" "trname" {
  pkg       = "default"
  name      = "acl_policy"
  interface = "port4"
  dstaddr   = ["all"]
  srcaddr   = ["all"]
  service   = ["ALL"]
  policyid  = 1
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `comments` - Comment.
* `dstaddr` - Destination address name.
* `interface` - Interface name.
* `name` - Policy name.
* `policyid` - Policy ID.
* `service` - Service name.
* `srcaddr` - Source address name.
* `status` - Enable/disable access control list status. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallAcl can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_acl.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
