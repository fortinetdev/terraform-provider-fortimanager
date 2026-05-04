---
subcategory: "Packages"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_packages_firewall_responseshapingpolicy"
description: |-
  Packages FirewallResponseShapingPolicy
---

# fortimanager_packages_firewall_responseshapingpolicy
Packages FirewallResponseShapingPolicy

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `pkg_folder_path` - Pkg Folder Path.
* `pkg` - Package.

* `class_id` - Class-Id.
* `class_id_reverse` - Class-Id-Reverse.
* `comment` - Comment.
* `diffserv_forward` - Diffserv-Forward. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Diffserv-Reverse. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Diffservcode-Forward.
* `diffservcode_rev` - Diffservcode-Rev.
* `dstaddr` - Dstaddr.
* `dstaddr6` - Dstaddr6.
* `fosid` - Id.
* `ip_version` - Ip-Version. Valid values: `6`, `4`.

* `matched_shaping_policies` - Matched-Shaping-Policies.
* `name` - Name.
* `per_ip_shaper` - Per-Ip-Shaper.
* `schedule` - Schedule.
* `srcaddr` - Srcaddr.
* `status` - Status. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic-Shaper.
* `traffic_shaper_reverse` - Traffic-Shaper-Reverse.
* `uuid` - Uuid.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Packages FirewallResponseShapingPolicy can be imported using any of these accepted formats:
```
Set import_options = ["pkg_folder_path=YOUR_VALUE", "pkg=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_packages_firewall_responseshapingpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
