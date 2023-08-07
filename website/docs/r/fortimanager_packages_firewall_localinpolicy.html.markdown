---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_localinpolicy"
description: |-
  Configure user defined IPv4 local-in policies.
---

# fortimanager_packages_firewall_localinpolicy
Configure user defined IPv4 local-in policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_localinpolicy" "labelname" {
  action            = "deny"
  dstaddr           = ["all"]
  ha_mgmt_intf_only = "disable"
  intf              = ["1-A1"]
  pkg               = "default"
  policyid          = 1
  schedule          = "always"
  service           = ["ALL"]
  srcaddr           = ["all"]
  status            = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Action performed on traffic matching the policy (default = deny). Valid values: `deny`, `accept`.

* `comments` - Comment.
* `dstaddr` - Destination address object from available options.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `ha_mgmt_intf_only` - Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `disable`, `enable`.

* `intf` - Incoming interface name from available options.
* `policyid` - User defined local in policy ID.
* `schedule` - Schedule object from available options.
* `service` - Service object from available options.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `srcaddr` - Source address object from available options.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `status` - Enable/disable this local-in policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `virtual_patch` - Enable/disable virtual patching. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallLocalInPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_localinpolicy.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
