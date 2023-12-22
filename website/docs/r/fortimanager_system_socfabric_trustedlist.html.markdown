---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_socfabric_trustedlist"
description: |-
  Pre-authorized security fabric nodes
---

# fortimanager_system_socfabric_trustedlist
Pre-authorized security fabric nodes

~> This resource is a sub resource for variable `trusted_list` of resource `fortimanager_system_socfabric`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:


* `fosid` - Trusted list ID.
* `serial` - FAZ serial number(support wildcard).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SocFabricTrustedList can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_socfabric_trustedlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

