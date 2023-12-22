---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_gtp_ippolicy"
description: |-
  IP policy.
---

# fortimanager_object_firewall_gtp_ippolicy
IP policy.

~> This resource is a sub resource for variable `ip_policy` of resource `fortimanager_object_firewall_gtp`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `gtp` - Gtp.

* `action` - Action. Valid values: `allow`, `deny`.

* `dstaddr` - Destination address name.
* `dstaddr6` - Destination IPv6 address name.
* `fosid` - ID.
* `srcaddr` - Source address name.
* `srcaddr6` - Source IPv6 address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall GtpIpPolicy can be imported using any of these accepted formats:
```
Set import_options = ["gtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_gtp_ippolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
