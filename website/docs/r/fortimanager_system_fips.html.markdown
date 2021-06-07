---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_fips"
description: |-
  Settings for FIPS-CC mode.
---

# fortimanager_system_fips
Settings for FIPS-CC mode.

## Example Usage

```hcl
resource "fortimanager_system_fips" "trname" {
  entropy_token    = "enable"
  re_seed_interval = 1440
  status           = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `entropy_token` - Enable/disable entropy token when switching to FIPS mode. enable - Enable entropy token. disable - Disable entropy token. dynamic - Dynamically detect entropy token during bootup. Valid values: `enable`, `disable`, `dynamic`.

* `re_seed_interval` - Kernel FIPS-compliant PRNG re-seed interval (0 to 1440 minutes)
* `status` - Enable/disable FIPS-CC mode. disable - Disable FIPS-CC mode. enable - Enable FIPS-CC mode. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fips can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_fips.labelname SystemFips
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

