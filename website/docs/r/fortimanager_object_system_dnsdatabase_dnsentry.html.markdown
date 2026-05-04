---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_dnsdatabase_dnsentry"
description: |-
  DNS entry.
---

# fortimanager_object_system_dnsdatabase_dnsentry
DNS entry.

~> This resource is a sub resource for variable `dns_entry` of resource `fortimanager_object_system_dnsdatabase`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `dns_database` - Dns Database.

* `canonical_name` - Canonical name of the host.
* `hostname` - Name of the host.
* `fosid` - DNS entry ID.
* `ip` - IPv4 address of the host.
* `ipv6` - IPv6 address of the host.
* `preference` - DNS entry preference (0 - 65535, highest preference = 0, default = 10).
* `status` - Enable/disable resource record status. Valid values: `disable`, `enable`.

* `ttl` - Time-to-live for this entry (0 to 2147483647 sec, default = 0).
* `type` - Resource record type. Valid values: `NS`, `MX`, `CNAME`, `A`, `AAAA`, `PTR`, `PTR_V6`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSystem DnsDatabaseDnsEntry can be imported using any of these accepted formats:
```
Set import_options = ["dns_database=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_dnsdatabase_dnsentry.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
