---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_webfilter_contentheader_entries_move"
description: |-
  Move content types used by web filter.
---

# fortimanager_object_webfilter_contentheader_entries_move
Move content types used by web filter.

## Example Usage

```hcl
resource "fortimanager_object_webfilter_contentheader_entries_move" "trname" {
  content_header = fortimanager_object_webfilter_contentheader.trname2.fosid
  entries        = fortimanager_object_webfilter_contentheader_entries.trname2.pattern
  target         = fortimanager_object_webfilter_contentheader_entries.trname.pattern
  option         = "after"
  depends_on     = [fortimanager_object_webfilter_contentheader_entries.trname2, fortimanager_object_webfilter_contentheader_entries.trname]
}

resource "fortimanager_object_webfilter_contentheader_entries" "trname2" {
  category       = 3
  pattern        = "abc"
  content_header = fortimanager_object_webfilter_contentheader.trname2.fosid
  depends_on     = [fortimanager_object_webfilter_contentheader.trname2]
}

resource "fortimanager_object_webfilter_contentheader_entries" "trname" {
  category       = 2
  pattern        = "abc"
  content_header = fortimanager_object_webfilter_contentheader.trname2.fosid
  depends_on     = [fortimanager_object_webfilter_contentheader.trname2]
}

resource "fortimanager_object_webfilter_contentheader" "trname2" {
  comment = "This is a Terraform example"
  fosid   = 2
  name    = "terr-webfilter-content-header2"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `content_header` - Content Header.
* `entries` - Entries.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{pattern}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the entries changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of entry.
