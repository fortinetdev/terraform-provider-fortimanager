---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_localinpolicy6"
description: |-
  Configure user defined IPv6 local-in policies.
---

# fortimanager_packages_firewall_localinpolicy6
Configure user defined IPv6 local-in policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg` - Package.

* `action` - Action performed on traffic matching the policy (default = deny). Valid values: `deny`, `accept`.

* `comments` - Comment.
* `dstaddr` - Destination address object from available options.
* `intf` - Incoming interface name from available options.
* `policyid` - User defined local in policy ID.
* `schedule` - Schedule object from available options.
* `service` - Service object from available options. Separate names with a space.
* `srcaddr` - Source address object from available options.
* `status` - Enable/disable this local-in policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages FirewallLocalInPolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["pkg=mypkg"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_localinpolicy6.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
