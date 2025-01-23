---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_fmgcluster"
description: |-
  fmg clsuter.
---

# fortimanager_system_fmgcluster
fmg clsuter.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `peer`: `fortimanager_system_fmgcluster_peer`



## Argument Reference


The following arguments are supported:


* `fqdn` - Local fqdn.
* `ip` - Local address.
* `mode` - Mode. standalone - Standalone. primary - Primary. worker - Worker. Valid values: `standalone`, `primary`, `worker`.

* `peer` - Peer. The structure of `peer` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `peer` block supports:

* `addr` - Address of peer.
* `fqdn` - FQDN of peer.
* `name` - Name of peer.
* `sn` - Serial number of peer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FmgCluster can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_fmgcluster.labelname SystemFmgCluster
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

