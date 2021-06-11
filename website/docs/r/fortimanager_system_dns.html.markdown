---
subcategory: "System Global"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_dns"
description: |-
  DNS configuration.
---

# fortimanager_system_dns
DNS configuration.

## Example Usage

```hcl
resource "fortimanager_system_dns" "trname" {
  secondary = "8.8.8.8"
}
```

## Argument Reference


The following arguments are supported:


* `ip6_primary` - IPv6 primary DNS IP.
* `ip6_secondary` - IPv6 secondary DNS IP.
* `primary` - Primary DNS IP.
* `secondary` - Secondary DNS IP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_dns.labelname SystemDns
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

