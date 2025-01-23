---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_systemp_system_replacemsg_mms"
description: |-
  Replacement messages.
---

# fortimanager_systemp_system_replacemsg_mms
Replacement messages.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `devprof` - Devprof.

* `buffer` - Message string.
* `charset` - character encoding used for replacement message Valid values: `us-ascii`, `utf-8`.

* `format` - Format flag. Valid values: `none`, `text`, `html`, `wml`.

* `header` - Header flag. Valid values: `none`, `http`, `8bit`.

* `image` - Message string.
* `msg_type` - Message type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Systemp SystemReplacemsgMms can be imported using any of these accepted formats:
```
Set import_options = ["devprof=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_systemp_system_replacemsg_mms.labelname SystempSystemReplacemsgMms
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
