---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_spamfilter_dnsbl"
description: |-
  Configure AntiSpam DNSBL/ORBL.
---

# fortimanager_object_spamfilter_dnsbl
Configure AntiSpam DNSBL/ORBL.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`entries`: `fortimanager_object_spamfilter_dnsbl_entries`



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

* `action` - Reject connection or mark as spam email. Valid values: `spam`, `reject`.

* `id` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectSpamfilter Dnsbl can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_spamfilter_dnsbl.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
