---
subcategory: "Packages Policy"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_shapingpolicy"
description: |-
  Configure shaping policies.
---

# fortimanager_packages_firewall_shapingpolicy
Configure shaping policies.

## Example Usage

```hcl
resource "fortimanager_packages_firewall_shapingpolicy" "labelname" {
  diffserv_forward     = "disable"
  diffserv_reverse     = "disable"
  dstaddr              = ["all"]
  dstintf              = ["any"]
  fosid                = 1
  internet_service     = "disable"
  internet_service_src = "disable"
  ip_version           = "4"
  pkg                  = "default"
  service              = ["ALL"]
  srcaddr              = ["all"]
  status               = "enable"
  tos                  = "0x00"
  tos_mask             = "0x00"
  tos_negate           = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `app_category` - IDs of one or more application categories that this shaper applies application control traffic shaping to.
* `app_group` - One or more application group names.
* `application` - IDs of one or more applications that this shaper applies application control traffic shaping to.
* `class_id` - Traffic class ID.
* `comment` - Comments.
* `cos` - VLAN CoS bit pattern.
* `cos_mask` - VLAN CoS evaluated bits.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dstaddr` - IPv4 destination address and address group names.
* `dstaddr6` - IPv6 destination address and address group names.
* `dstintf` - One or more outgoing (egress) interfaces.
* `groups` - Apply this traffic shaping policy to user groups that have authenticated with the FortiGate.
* `fosid` - Shaping policy ID (0 - 4294967295).
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.
* `internet_service_name` - Internet Service ID.
* `internet_service_id` - Internet Service ID.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `disable`, `enable`.

* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_name` - Internet Service source name.
* `internet_service_src_id` - Internet Service source ID.
* `ip_version` - Apply this traffic shaping policy to IPv4 or IPv6 traffic. Valid values: `4`, `6`.

* `name` - Shaping policy name.
* `per_ip_shaper` - Per-IP traffic shaper to apply with this policy.
* `schedule` - Schedule name.
* `service` - Service and service group names.
* `srcaddr` - IPv4 source address and address group names.
* `srcaddr6` - IPv6 source address and address group names.
* `srcintf` - One or more incoming (ingress) interfaces.
* `status` - Enable/disable this traffic shaping policy. Valid values: `disable`, `enable`.

* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the firewall policy.
* `traffic_shaper_reverse` - Traffic shaper to apply to response traffic received by the firewall policy.
* `traffic_type` - Traffic type. Valid values: `forwarding`, `local-in`, `local-out`.

* `url_category` - IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to.
* `users` - Apply this traffic shaping policy to individual users that have authenticated with the FortiGate.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Packages FirewallShapingPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_shapingpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
