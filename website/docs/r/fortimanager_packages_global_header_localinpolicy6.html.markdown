---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_global_header_localinpolicy6"
description: |-
  Configure user defined IPv6 local-in policies.
---

# fortimanager_packages_global_header_localinpolicy6
Configure user defined IPv6 local-in policies.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `_policy_block` - Assigned policy block.  When this attribute is set, the policy represent a policy block, and all other attributes are ignored. This attribute is not available when configuring policy inside a policy block.
* `action` - Action performed on traffic matching the policy (default = deny). Valid values: `deny`, `accept`.

* `comments` - Comment.
* `custom_tags` - Custom tags.
* `dstaddr` - Destination address object from available options.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `disable`, `enable`.

* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this local-in policy.If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service6_src_custom` - Custom IPv6 Internet Service source name.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name.
* `internet_service6_src_fortiguard` - FortiGuard IPv6 Internet Service source name.
* `internet_service6_src_group` - Internet Service6 source group name.
* `internet_service6_src_name` - IPv6 Internet Service source name.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `intf` - Incoming interface name from available options.
* `logtraffic` - Enable/disable local-in traffic logging. Valid values: `disable`, `enable`.

* `policyid` - User defined local in policy ID.
* `schedule` - Schedule object from available options.
* `service` - Service object from available options. Separate names with a space.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `disable`, `enable`.

* `srcaddr` - Source address object from available options.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `disable`, `enable`.

* `status` - Enable/disable this local-in policy. Valid values: `disable`, `enable`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `virtual_patch` - Enable/disable the virtual patching feature. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Packages GlobalHeaderLocalInPolicy6 can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_global_header_localinpolicy6.labelname {{policyid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
