---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetservice_entry"
description: |-
  Entries in the Internet Service database.
---

# fortimanager_object_firewall_internetservice_entry
Entries in the Internet Service database.

~> This resource is a sub resource for variable `entry` of resource `fortimanager_object_firewall_internetservice`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fosid` - Entry ID.
* `ip_number` - Total number of IP addresses.
* `ip_range_number` - Total number of IP ranges.
* `port` - Integer value for the TCP/IP port (0 - 65535).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall InternetServiceEntry can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetservice_entry.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
