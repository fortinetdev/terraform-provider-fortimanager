---
subcategory: "Object ICAP"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_icap_servergroup"
description: |-
  Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing.
---

# fortimanager_object_icap_servergroup
Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_list`: `fortimanager_object_icap_servergroup_serverlist`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `ldb_method` - Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.

* `name` - Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `name` - ICAP server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectIcap ServerGroup can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_icap_servergroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
