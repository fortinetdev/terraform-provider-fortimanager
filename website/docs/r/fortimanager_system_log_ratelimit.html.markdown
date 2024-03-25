---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_ratelimit"
description: |-
  Logging rate limit.
---

# fortimanager_system_log_ratelimit
Logging rate limit.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `device`: `fortimanager_system_log_ratelimit_device`
>- `ratelimits`: `fortimanager_system_log_ratelimit_ratelimits`



## Example Usage

```hcl
resource "fortimanager_system_log_ratelimit" "trname" {
  mode = "manual"
  ratelimits {
    id     = 1
    filter = 1
  }
}
```

## Argument Reference


The following arguments are supported:


* `device` - Device. The structure of `device` block is documented below.
* `device_ratelimit_default` - Default maximum device log rate limit.
* `mode` - Logging rate limit mode. disable - Logging rate limit function disabled. manual - System rate limit and device rate limit both configurable, no limit if not configured. Valid values: `disable`, `manual`.

* `ratelimits` - Ratelimits. The structure of `ratelimits` block is documented below.
* `system_ratelimit` - Maximum system log rate limit.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `device` block supports:

* `device` - Device(s) filter according to filter-type setting, wildcard expression supported.
* `filter_type` - Device filter type. devid - Device ID. Valid values: `devid`.

* `id` - Device filter ID.
* `ratelimit` - Maximum device log rate limit.

The `ratelimits` block supports:

* `filter` - Device or ADOM filter according to filter-type setting, wildcard expression supported.
* `filter_type` - Device filter type. devid - Device ID. adom - ADOM name. Valid values: `devid`, `adom`.

* `id` - Filter ID.
* `ratelimit` - Maximum log rate limit.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogRatelimit can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_ratelimit.labelname SystemLogRatelimit
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

