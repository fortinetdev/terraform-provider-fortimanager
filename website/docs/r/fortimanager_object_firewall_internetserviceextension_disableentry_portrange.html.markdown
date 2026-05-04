---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetserviceextension_disableentry_portrange"
description: |-
  Port ranges in the disable entry.
---

# fortimanager_object_firewall_internetserviceextension_disableentry_portrange
Port ranges in the disable entry.

~> This resource is a sub resource for variable `port_range` of resource `fortimanager_object_firewall_internetserviceextension_disableentry`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `internet_service_extension` - Internet Service Extension.
* `disable_entry` - Disable Entry.

* `end_port` - Ending TCP/UDP/SCTP destination port (0 to 65535).
* `fosid` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (0 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall InternetServiceExtensionDisableEntryPortRange can be imported using any of these accepted formats:
```
Set import_options = ["internet_service_extension=YOUR_VALUE", "disable_entry=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetserviceextension_disableentry_portrange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
