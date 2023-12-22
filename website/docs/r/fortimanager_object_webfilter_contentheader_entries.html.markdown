---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_contentheader_entries"
description: |-
  Configure content types used by web filter.
---

# fortimanager_object_webfilter_contentheader_entries
Configure content types used by web filter.

~> This resource is a sub resource for variable `entries` of resource `fortimanager_object_webfilter_contentheader`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_webfilter_contentheader_entries" "trname" {
  category       = 2
  pattern        = "abc"
  content_header = fortimanager_object_webfilter_contentheader.trname.fosid
  depends_on     = [fortimanager_object_webfilter_contentheader.trname]
}

resource "fortimanager_object_webfilter_contentheader" "trname" {
  comment = "This is a Terraform example"
  fosid   = 4
  name    = "terr-webfilter-content-header3"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `content_header` - Content Header.

* `action` - Action to take for this content type. Valid values: `exempt`, `block`, `allow`.

* `category` - Categories that this content type applies to.
* `pattern` - Content type (regular expression).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pattern}}.

## Import

ObjectWebfilter ContentHeaderEntries can be imported using any of these accepted formats:
```
Set import_options = ["content_header=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_webfilter_contentheader_entries.labelname {{pattern}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
