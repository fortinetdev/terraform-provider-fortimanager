---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_mheader_entries"
description: |-
  Spam filter mime header content.
---

# fortimanager_object_emailfilter_mheader_entries
Spam filter mime header content.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_emailfilter_mheader`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_emailfilter_mheader_entries" "trname" {
  mheader    = fortimanager_object_emailfilter_mheader.trname.fosid
  action     = "spam"
  fieldbody  = "body"
  fieldname  = "name"
  fosid      = 1
  depends_on = [fortimanager_object_emailfilter_mheader.trname]
}

resource "fortimanager_object_emailfilter_mheader" "trname" {
  comment = "This is a Terraform example"
  fosid   = 1
  name    = "terr-emailfilter-mheader"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `mheader` - Mheader.

* `action` - Mark spam or good. Valid values: `spam`, `clear`.

* `fieldbody` - Pattern for the header field body.
* `fieldname` - Pattern for header field name.
* `fosid` - Mime header entry ID.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.

* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectEmailfilter MheaderEntries can be imported using any of these accepted formats:
```
Set import_options = ["mheader=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_mheader_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
