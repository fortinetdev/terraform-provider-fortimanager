---
subcategory: "Object Dnsfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dnsfilter_domainfilter"
description: |-
  Configure DNS domain filters.
---

# fortimanager_object_dnsfilter_domainfilter
Configure DNS domain filters.

## Example Usage

```hcl
resource "fortimanager_object_dnsfilter_domainfilter" "trname" {
  name    = "F11"
  comment = "FDS332911"
  fosid   = 25
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.

* `domain` - Domain entries to be filtered.
* `id` - Id.
* `status` - Enable/disable this domain filter. Valid values: `disable`, `enable`.

* `type` - DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDnsfilter DomainFilter can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dnsfilter_domainfilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
