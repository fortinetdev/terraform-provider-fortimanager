---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_log_syslogd_filter"
description: |-
  Filters for remote system server.
---

# fortimanager_systemp_log_syslogd_filter
Filters for remote system server.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `exclude_list`: `fortimanager_systemp_log_syslogd_filter_excludelist`
>- `free_style`: `fortimanager_systemp_log_syslogd_filter_freestyle`



## Example Usage

```hcl
resource "fortimanager_systemp_log_syslogd_filter" "trname" {
  devprof = "default"
  cifs    = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `anomaly` - Enable/disable anomaly logging. Valid values: `disable`, `enable`.

* `forti_switch` - Enable/disable Forti-Switch logging. Valid values: `disable`, `enable`.

* `cifs` - Cifs. Valid values: `disable`, `enable`.

* `dns` - Dns. Valid values: `disable`, `enable`.

* `exclude_list` - Exclude-List. The structure of `exclude_list` block is documented below.
* `filter` - Syslog filter.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.

* `forward_traffic` - Enable/disable forward traffic logging. Valid values: `disable`, `enable`.

* `free_style` - Free-Style. The structure of `free_style` block is documented below.
* `gtp` - Enable/disable GTP messages logging. Valid values: `disable`, `enable`.

* `local_traffic` - Enable/disable local in or out traffic logging. Valid values: `disable`, `enable`.

* `multicast_traffic` - Enable/disable multicast traffic logging. Valid values: `disable`, `enable`.

* `netscan_discovery` - Enable/disable netscan discovery event logging. Valid values: `disable`, `enable`.

* `netscan_vulnerability` - Enable/disable netscan vulnerability event logging. Valid values: `disable`, `enable`.

* `severity` - Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `sniffer_traffic` - Enable/disable sniffer traffic logging. Valid values: `disable`, `enable`.

* `ssh` - Ssh. Valid values: `disable`, `enable`.

* `ssl` - Ssl. Valid values: `disable`, `enable`.

* `voip` - Enable/disable VoIP logging. Valid values: `disable`, `enable`.

* `ztna_traffic` - Enable/disable ztna traffic logging. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `exclude_list` block supports:

* `category` - Category. Valid values: `app-ctrl`, `attack`, `dlp`, `event`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `spam`, `anomaly`, `waf`.

* `fields` - Fields. The structure of `fields` block is documented below.
* `id` - Id.

The `fields` block supports:

* `args` - Args.
* `field` - Field.
* `negate` - Negate. Valid values: `disable`, `enable`.


The `free_style` block supports:

* `category` - Log category. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `voip`, `dlp`, `app-ctrl`, `anomaly`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `icap`, `ztna`.

* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter. Valid values: `include`, `exclude`.

* `id` - Entry ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp LogSyslogdFilter can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_log_syslogd_filter.labelname SystempLogSyslogdFilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
