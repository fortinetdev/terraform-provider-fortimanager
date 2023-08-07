---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_socfabric"
description: |-
  SOC Fabric.
---

# fortimanager_system_socfabric
SOC Fabric.

## Argument Reference


The following arguments are supported:


* `name` - Fabric name.
* `port` - communication port (1 - 65535).
* `psk` - Fabric auth pwd.
* `role` - Enable or Disable SOC Fabric. member - SOC Fabric member. supervisor - SOC Fabric supervisor. Valid values: `member`, `supervisor`.

* `secure_connection` - Enable or Disable SSL/TLS. disable - Disable SSL/TLS. enable - Enable SSL/TLS. Valid values: `disable`, `enable`.

* `status` - Enable or Disable SOC Fabric. disable - Disable SOC Fabric. enable - Enable SOC Fabric. Valid values: `disable`, `enable`.

* `supervisor` - IP/FQDN of supervisor.
* `trusted_list` - Trusted-List. The structure of `trusted_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `trusted_list` block supports:

* `id` - Trusted list ID.
* `serial` - FAZ serial number(support wildcard).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SocFabric can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_socfabric.labelname SystemSocFabric
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

