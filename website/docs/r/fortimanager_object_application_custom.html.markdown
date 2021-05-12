---
subcategory: "ObjectApplication"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_custom"
description: |-
  Configure custom application signatures.
---

# fortimanager_object_application_custom
Configure custom application signatures.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `behavior` - Custom application signature behavior.
* `category` - Custom application category ID (use ? to view available options).
* `comment` - Comment.
* `fosid` - Id.
* `name` - Name.
* `protocol` - Custom application signature protocol.
* `signature` - The text that makes up the actual custom application signature.
* `tag` - Signature tag.
* `technology` - Custom application signature technology.
* `vendor` - Custom application signature vendor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{tag}}.

## Import

ObjectApplication Custom can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_custom.labelname {{tag}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
