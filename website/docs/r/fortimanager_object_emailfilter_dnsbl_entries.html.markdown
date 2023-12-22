---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_dnsbl_entries"
description: |-
  Spam filter DNSBL and ORBL server.
---

# fortimanager_object_emailfilter_dnsbl_entries
Spam filter DNSBL and ORBL server.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_emailfilter_dnsbl`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_emailfilter_dnsbl_entries" "trname" {
  action     = "spam"
  fosid      = 12
  status     = "disable"
  depends_on = [fortimanager_object_emailfilter_dnsbl.trname]
  dnsbl      = fortimanager_object_emailfilter_dnsbl.trname.fosid
}

resource "fortimanager_object_emailfilter_dnsbl" "trname" {
  comment = "This is a Terraform example"
  fosid   = 1
  name    = "terr-emailfilter-dnsbl"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `dnsbl` - Dnsbl.

* `action` - Reject connection or mark as spam email. Valid values: `spam`, `reject`.

* `fosid` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectEmailfilter DnsblEntries can be imported using any of these accepted formats:
```
Set import_options = ["dnsbl=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_dnsbl_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
