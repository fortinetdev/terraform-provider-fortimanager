---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_vip_gslbpublicips"
description: |-
  Publicly accessible IP addresses for the FortiGSLB service.
---

# fortimanager_object_firewall_vip_gslbpublicips
Publicly accessible IP addresses for the FortiGSLB service.

~> This resource is a sub resource for variable `gslb_public_ips` of resource `fortimanager_object_firewall_vip`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `vip` - Vip.

* `index` - Index of this public IP setting.
* `ip` - The publicly accessible IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

ObjectFirewall VipGslbPublicIps can be imported using any of these accepted formats:
```
Set import_options = ["vip=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_vip_gslbpublicips.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
