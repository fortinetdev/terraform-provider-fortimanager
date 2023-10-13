---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_vpn_ssl_web_portal_macaddrcheckrule"
description: |-
  Client MAC address check rule.
---

# fortimanager_object_vpn_ssl_web_portal_macaddrcheckrule
Client MAC address check rule.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `portal` - Portal.

* `mac_addr_list` - Client MAC address list.
* `mac_addr_mask` - Client MAC address mask.
* `name` - Client MAC address check rule name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectVpn SslWebPortalMacAddrCheckRule can be imported using any of these accepted formats:
```
Set import_options = ["portal=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_vpn_ssl_web_portal_macaddrcheckrule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
