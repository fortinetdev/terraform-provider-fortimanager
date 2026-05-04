---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_dnsdatabase"
description: |-
  Configure DNS databases.
---

# fortimanager_object_system_dnsdatabase
Configure DNS databases.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dns_entry`: `fortimanager_object_system_dnsdatabase_dnsentry`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `allow_transfer` - DNS zone transfer IP address list.
* `authoritative` - Enable/disable authoritative zone. Valid values: `disable`, `enable`.

* `contact` - Email address of the administrator for this zone. You can specify only the username, such as admin or the full email address, such as admin@test.com When using only a username, the domain of the email will be this zone.
* `dns_entry` - Dns-Entry. The structure of `dns_entry` block is documented below.
* `domain` - Domain name.
* `forwarder` - DNS zone forwarder IP address list.
* `ip_master` - IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
* `forwarder6` - Forwarder IPv6 address.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip_primary` - IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
* `name` - Zone name.
* `primary_name` - Domain name of the default DNS server for this zone.
* `rr_max` - Maximum number of resource records (10 - 65536, 0 means infinite).
* `source_ip` - Source IP for forwarding to DNS server.
* `source_ip_interface` - IP address of the specified interface as the source IP address.
* `source_ip6` - IPv6 source IP address for forwarding to DNS server.
* `status` - Enable/disable this DNS zone. Valid values: `disable`, `enable`.

* `ttl` - Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
* `type` - Zone type (primary to manage entries directly, secondary to import entries from other zones). Valid values: `primary`, `secondary`.

* `view` - Zone view (public to serve public clients, shadow to serve internal clients). Valid values: `shadow`, `public`, `shadow-ztna`, `proxy`.

* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `dns_entry` block supports:

* `canonical_name` - Canonical name of the host.
* `hostname` - Name of the host.
* `id` - DNS entry ID.
* `ip` - IPv4 address of the host.
* `ipv6` - IPv6 address of the host.
* `preference` - DNS entry preference (0 - 65535, highest preference = 0, default = 10).
* `status` - Enable/disable resource record status. Valid values: `disable`, `enable`.

* `ttl` - Time-to-live for this entry (0 to 2147483647 sec, default = 0).
* `type` - Resource record type. Valid values: `NS`, `MX`, `CNAME`, `A`, `AAAA`, `PTR`, `PTR_V6`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem DnsDatabase can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_dnsdatabase.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
