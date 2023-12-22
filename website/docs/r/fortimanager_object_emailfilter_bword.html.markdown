---
subcategory: "Object Emailfilter"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_emailfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

# fortimanager_object_emailfilter_bword
Configure AntiSpam banned word list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`entries`: `fortimanager_object_emailfilter_bword_entries`



## Example Usage

```hcl
resource "fortimanager_object_emailfilter_bword" "trname" {
  comment = "This is a Terraform example"
  fosid   = 1
  name    = "terr-emailfilter-bword"
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

* `action` - Mark spam or good. Valid values: `spam`, `clear`.

* `id` - Banned word entry ID.
* `language` - Language for the banned word. Valid values: `western`, `simch`, `trach`, `japanese`, `korean`, `french`, `thai`, `spanish`.

* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.

* `score` - Score value.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.

* `where` - Component of the email to be scanned. Valid values: `all`, `subject`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectEmailfilter Bword can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_emailfilter_bword.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
