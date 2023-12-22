---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_connector"
description: |-
  ObjectUser Connector
---

# fortimanager_object_user_connector
ObjectUser Connector

## Example Usage

```hcl
resource "fortimanager_object_user_connector" "trname" {
  name = "terr-connector"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `name` - Name.
* `status` - Status. Valid values: `disable`, `enable`.

* `type` - Type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Connector can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_connector.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
