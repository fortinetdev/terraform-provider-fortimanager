---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_ums_setting"
description: |-
  ObjectUms Setting
---

# fortimanager_object_ums_setting
ObjectUms Setting

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `connector` - Connector.
* `description` - Description.
* `flags` - Flags.
* `name` - Name.
* `type` - Type. Valid values: `aws`, `azure`, `gcp`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUms Setting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_ums_setting.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
