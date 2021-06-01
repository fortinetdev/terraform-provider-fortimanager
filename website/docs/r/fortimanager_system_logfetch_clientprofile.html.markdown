---
subcategory: "System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_logfetch_clientprofile"
description: |-
  Log-fetch client profile settings.
---

# fortimanager_system_logfetch_clientprofile
Log-fetch client profile settings.

## Argument Reference


The following arguments are supported:


* `client_adom` - Log-fetch client side's adom name.
* `data_range` - Data-range for fetched logs. custom - Specify some other date and time range. Valid values: `custom`.

* `data_range_value` - Last n days or hours.
* `device_filter` - Device-Filter. The structure of `device_filter` block is documented below.
* `end_time` - End date and time of the data-range <hh:mm yyyy/mm/dd>.
* `fosid` - Log-fetch client profile ID.
* `index_fetch_logs` - Enable/Disable indexing logs automatically after fetching logs. disable - Disable attribute function. enable - Enable attribute function. Valid values: `disable`, `enable`.

* `log_filter` - Log-Filter. The structure of `log_filter` block is documented below.
* `log_filter_logic` - And/Or logic for log-filters. and - Logic And. or - Logic Or. Valid values: `and`, `or`.

* `log_filter_status` - Enable/Disable log-filter. disable - Disable attribute function. enable - Enable attribute function. Valid values: `disable`, `enable`.

* `name` - Name of log-fetch client profile.
* `password` - Log-fetch server login password.
* `secure_connection` - Enable/Disable protecting log-fetch connection with TLS/SSL. disable - Disable attribute function. enable - Enable attribute function. Valid values: `disable`, `enable`.

* `server_adom` - Log-fetch server side's adom name.
* `server_ip` - Log-fetch server IP address.
* `start_time` - Start date and time of the data-range <hh:mm yyyy/mm/dd>.
* `sync_adom_config` - Enable/Disable sync adom related config. disable - Disable attribute function. enable - Enable attribute function. Valid values: `disable`, `enable`.

* `user` - Log-fetch server login username.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `device_filter` block supports:

* `adom` - Adom name.
* `device` - Device name or Serial number.
* `id` - Add or edit a device filter.
* `vdom` - Vdom filters.

The `log_filter` block supports:

* `field` - Field name.
* `id` - Log filter ID.
* `oper` - Field filter operator. &lt; - =Less than or equal to &gt; - =Greater than or equal to contain - Contain not-contain - Not contain match - Match (expression) Valid values: `=`, `!=`, `<`, `>`, `<=`, `>=`, `contain`, `not-contain`, `match`.

* `value` - Field filter operand or free-text matching expression.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogFetchClientProfile can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_logfetch_clientprofile.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

