---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_log_ratelimit_ratelimits"
description: |-
  Per device or ADOM log rate limits.
---

# fortimanager_system_log_ratelimit_ratelimits
Per device or ADOM log rate limits.

~> This resource is a sub resource for variable `ratelimits` of resource `fortimanager_system_log_ratelimit`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_system_log_ratelimit_ratelimits" "trname" {
  fosid  = 2
  filter = 1
}
```

## Argument Reference


The following arguments are supported:


* `filter` - Device or ADOM filter according to filter-type setting, wildcard expression supported.
* `filter_type` - Device filter type. devid - Device ID. adom - ADOM name. Valid values: `devid`, `adom`.

* `fosid` - Filter ID.
* `ratelimit` - Maximum log rate limit.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System LogRatelimitRatelimits can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_log_ratelimit_ratelimits.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

