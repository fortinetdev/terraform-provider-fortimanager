---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dnsfilter_domainfilter_entries"
description: |-
  DNS domain filter entries.
---

# fortimanager_object_dnsfilter_domainfilter_entries
DNS domain filter entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `domain_filter` - Domain Filter.

* `action` - Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.

* `domain` - Domain entries to be filtered.
* `fosid` - Id.
* `status` - Enable/disable this domain filter. Valid values: `disable`, `enable`.

* `type` - DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectDnsfilter DomainFilterEntries can be imported using any of these accepted formats:
```
Set import_options = ["domain_filter=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_dnsfilter_domainfilter_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
