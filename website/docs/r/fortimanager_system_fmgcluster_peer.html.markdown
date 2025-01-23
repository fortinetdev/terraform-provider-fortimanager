---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_fmgcluster_peer"
description: |-
  Peer.
---

# fortimanager_system_fmgcluster_peer
Peer.

~> This resource is a sub resource for variable `peer` of resource `fortimanager_system_fmgcluster`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:


* `addr` - Address of peer.
* `fqdn` - FQDN of peer.
* `name` - Name of peer.
* `sn` - Serial number of peer.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{sn}}.

## Import

System FmgClusterPeer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_fmgcluster_peer.labelname {{sn}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

