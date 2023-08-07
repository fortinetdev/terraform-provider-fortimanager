---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cloud_orchestration"
description: |-
  ObjectCloud Orchestration
---

# fortimanager_object_cloud_orchestration
ObjectCloud Orchestration

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `configuration` - Configuration.
* `description` - Description.
* `name` - Name.
* `type` - Type. Valid values: `aws`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCloud Orchestration can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cloud_orchestration.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
