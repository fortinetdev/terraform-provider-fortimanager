---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_dnsfilter_domainfilter_entries_move"
description: |-
  DNS domain filter entries.
---

# fortimanager_object_dnsfilter_domainfilter_entries_move
DNS domain filter entries.

## Example Usage

```hcl
resource "fortimanager_object_dnsfilter_domainfilter_entries_move" "trname" {
  domain_filter = fortimanager_object_dnsfilter_domainfilter.trname.fosid
  entries       = 1
  target        = 2
  option        = "after"
  depends_on    = [fortimanager_object_dnsfilter_domainfilter_entries.trname2, fortimanager_object_dnsfilter_domainfilter_entries.trname]
}

resource "fortimanager_object_dnsfilter_domainfilter_entries" "trname2" {
  domain_filter = fortimanager_object_dnsfilter_domainfilter.trname.fosid
  fosid         = 2
  domain        = "domain"
  depends_on    = [fortimanager_object_dnsfilter_domainfilter.trname]
}

resource "fortimanager_object_dnsfilter_domainfilter_entries" "trname" {
  domain_filter = fortimanager_object_dnsfilter_domainfilter.trname.fosid
  fosid         = 1
  domain        = "domain"
  depends_on    = [fortimanager_object_dnsfilter_domainfilter.trname]
}

resource "fortimanager_object_dnsfilter_domainfilter" "trname" {
  name    = "F11"
  comment = "This is a Terraform example"
  fosid   = 25
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `domain_filter` - Domain Filter.
* `entries` - Entries.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the entries changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of entry.
