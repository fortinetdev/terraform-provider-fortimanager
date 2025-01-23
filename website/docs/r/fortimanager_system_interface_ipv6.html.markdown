---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_interface_ipv6"
description: |-
  IPv6 of interface.
---

# fortimanager_system_interface_ipv6
IPv6 of interface.

~> This resource is a sub resource for variable `ipv6` of resource `fortimanager_system_interface`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `interface` - Interface.

* `ip6_address` - IPv6 address/prefix of interface.
* `ip6_allowaccess` - Allow management access to interface. ping - PING access. https - HTTPS access. ssh - SSH access. snmp - SNMP access. http - HTTP access. webservice - Web service access. https-logging - Logging over HTTPS access. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `webservice`, `https-logging`.

* `ip6_autoconf` - Enable/disable address auto config (SLAAC). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System InterfaceIpv6 can be imported using any of these accepted formats:
```
Set import_options = ["interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_interface_ipv6.labelname SystemInterfaceIpv6
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

