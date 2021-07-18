---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_adgrp"
description: |-
  Configure FSSO groups.
---

# fortimanager_object_user_adgrp
Configure FSSO groups.

## Example Usage

```hcl
resource "fortimanager_object_user_adgrp" "trname" {
  connector_source = "terr-tefv"
  fosid            = 1
  name             = "terr-user-adgrp"
  server_name      = "FortiManager"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `connector_source` - FSSO connector source.
* `fosid` - Id.
* `name` - Name.
* `server_name` - FSSO agent name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Adgrp can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_adgrp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
