---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetserviceextension_disableentry_ip6range"
description: |-
  IPv6 ranges in the disable entry.
---

# fortimanager_object_firewall_internetserviceextension_disableentry_ip6range
IPv6 ranges in the disable entry.

~> This resource is a sub resource for variable `ip6_range` of resource `fortimanager_object_firewall_internetserviceextension_disableentry`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `internet_service_extension` - Internet Service Extension.
* `disable_entry` - Disable Entry.

* `end_ip6` - End IPv6 address.
* `fosid` - Disable entry range ID.
* `start_ip6` - Start IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall InternetServiceExtensionDisableEntryIp6Range can be imported using any of these accepted formats:
```
Set import_options = ["internet_service_extension=YOUR_VALUE", "disable_entry=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetserviceextension_disableentry_ip6range.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
