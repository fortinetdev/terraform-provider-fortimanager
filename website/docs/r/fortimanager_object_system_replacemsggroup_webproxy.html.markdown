---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_replacemsggroup_webproxy"
description: |-
  Replacement message table entries.
---

# fortimanager_object_system_replacemsggroup_webproxy
Replacement message table entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `replacemsg_group` - Replacemsg Group.

* `buffer` - Message string.
* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

ObjectSystem ReplacemsgGroupWebproxy can be imported using any of these accepted formats:
```
Set import_options = ["replacemsg_group=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_replacemsggroup_webproxy.labelname {{msg_type}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
