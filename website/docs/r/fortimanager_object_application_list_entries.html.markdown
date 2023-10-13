---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_list_entries"
description: |-
  Application list entries.
---

# fortimanager_object_application_list_entries
Application list entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `list` - List.

* `action` - Pass or block traffic, or reset connection for traffic from this application. Valid values: `pass`, `block`, `reset`.

* `application` - ID of allowed applications.
* `behavior` - Application behavior filter.
* `category` - Category ID list.
* `exclusion` - ID of excluded applications.
* `fosid` - Entry ID.
* `log` - Enable/disable logging for this application list. Valid values: `disable`, `enable`.

* `log_packet` - Enable/disable packet logging. Valid values: `disable`, `enable`.

* `parameters` - Parameters. The structure of `parameters` block is documented below.
* `per_ip_shaper` - Per-IP traffic shaper.
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.

* `protocols` - Application protocol filter.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`.

* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.

* `rate_count` - Count of the rate.
* `rate_duration` - Duration (sec) of the rate.
* `rate_mode` - Rate limit mode. Valid values: `periodical`, `continuous`.

* `rate_track` - Track the packet protocol field. Valid values: `none`, `src-ip`, `dest-ip`, `dhcp-client-mac`, `dns-domain`.

* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
* `session_ttl` - Session TTL (0 = default).
* `shaper` - Traffic shaper.
* `shaper_reverse` - Reverse traffic shaper.
* `sub_category` - Application Sub-category ID list.
* `technology` - Application technology filter.
* `vendor` - Application vendor filter.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `parameters` block supports:

* `id` - Parameter ID.
* `members` - Members. The structure of `members` block is documented below.
* `value` - Parameter value.

The `members` block supports:

* `id` - Parameter.
* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectApplication ListEntries can be imported using any of these accepted formats:
```
Set import_options = ["list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_list_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
