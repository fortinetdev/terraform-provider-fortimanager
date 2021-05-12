---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fdssetting_serveroverride"
description: |-
  Server override configure.
---

# fortimanager_fmupdate_fdssetting_serveroverride
Server override configure.

## Argument Reference


The following arguments are supported:


* `servlist` - Servlist. The structure of `servlist` block is documented below.
* `status` - Override status. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `servlist` block supports:

* `id` - Override server ID (1 - 10).
* `ip` - IPv4 address of the override server.
* `ip6` - IPv6 address of the override server.
* `port` - Port number to use when contacting FortiGuard (1 - 65535, default = 443).
* `service_type` - Override service type. fct - Server override config for fct fds - Server override config for fds Valid values: `fct`, `fds`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSettingServerOverride can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fdssetting_serveroverride.labelname FmupdateFdsSettingServerOverride
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

