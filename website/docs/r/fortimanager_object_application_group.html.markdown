---
subcategory: "ObjectApplication"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_application_group"
description: |-
  Configure firewall application groups.
---

# fortimanager_object_application_group
Configure firewall application groups.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `application` - Application ID list.
* `category` - Application category ID list.
* `comment` - Comment
* `name` - Application group name.
* `type` - Application group type. Valid values: `application`, `category`, `filter`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectApplication Group can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_application_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
