---
subcategory: "Object Firewall"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_firewall_internetservice"
description: |-
  Show Internet Service application.
---

# fortimanager_object_firewall_internetservice
Show Internet Service application.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`entry`: `fortimanager_object_firewall_internetservice_entry`



## Example Usage

```hcl
resource "fortimanager_object_firewall_internetservice" "trname" {
  name  = "Google-DNS"
  fosid = 65539
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `city` - City sequence number list.
* `city6` - IPv6 City sequence number list.
* `country` - Country sequence number list.
* `country6` - IPv6 Country sequence number list.
* `database` - Database. Valid values: `isdb`, `irdb`.

* `direction` - Direction. Valid values: `src`, `dst`, `both`.

* `entry` - Entry. The structure of `entry` block is documented below.
* `extra_ip_range_number` - Extra-Ip-Range-Number.
* `extra_ip6_range_number` - Extra-Ip6-Range-Number.
* `icon_id` - Icon-Id.
* `fosid` - Id.
* `ip_number` - Ip-Number.
* `ip_range_number` - Ip-Range-Number.
* `ip6_range_number` - Ip6-Range-Number.
* `jitter_threshold` - Jitter-Threshold.
* `latency_threshold` - Latency-Threshold.
* `name` - Name.
* `offset` - Offset.
* `obsolete` - Obsolete.
* `region` - Region sequence number list.
* `region6` - IPv6 Region sequence number list.
* `packetloss_threshold` - Packetloss-Threshold.
* `reputation` - Reputation.
* `singularity` - Singularity.
* `sld_id` - Sld-Id.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entry` block supports:

* `id` - Entry ID.
* `ip_number` - Total number of IP addresses.
* `ip_range_number` - Total number of IP ranges.
* `port` - Integer value for the TCP/IP port (0 - 65535).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectFirewall InternetService can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_firewall_internetservice.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
