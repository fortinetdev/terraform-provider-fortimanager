---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_localinpolicy"
description: |-
  IPv4 local in policy configuration.
---

# fortimanager_system_localinpolicy
IPv4 local in policy configuration.

## Argument Reference


The following arguments are supported:


* `action` - Action performed on traffic matching this policy. drop - Drop traffic matching this policy (default). reject - Reject traffic matching this policy. accept - Allow traffic matching this policy. Valid values: `drop`, `reject`, `accept`.

* `dport` - Destination port number (0 for all).
* `dst` - Destination IP and mask.
* `fosid` - Entry number.
* `intf` - Incoming interface name.
* `protocol` - Traffic protocol. tcp - TCP only. udp - UDP only. tcp_udp - Both TCP and UDP. Valid values: `tcp`, `udp`, `tcp_udp`.

* `src` - Source IP and mask.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LocalInPolicy can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_localinpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

