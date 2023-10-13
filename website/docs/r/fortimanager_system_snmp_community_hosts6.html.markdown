---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_snmp_community_hosts6"
description: |-
  Allow hosts configuration for IPv6.
---

# fortimanager_system_snmp_community_hosts6
Allow hosts configuration for IPv6.

## Argument Reference


The following arguments are supported:

* `community` - Community.

* `fosid` - Host entry ID.
* `interface` - Allow interface name.
* `ip` - Allow host IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SnmpCommunityHosts6 can be imported using any of these accepted formats:
```
Set import_options = ["community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_snmp_community_hosts6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

