---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_fssopolling_adgrp"
description: |-
  LDAP Group Info.
---

# fortimanager_object_user_fssopolling_adgrp
LDAP Group Info.

~> This resource is a sub resource for variable `adgrp` of resource `fortimanager_object_user_fssopolling`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_user_fssopolling_adgrp" "trname" {
  fsso_polling = fortimanager_object_user_fssopolling.trname.fosid
  name         = "terr-adgrp"
  depends_on   = [fortimanager_object_user_fssopolling.trname]
}

resource "fortimanager_object_user_fssopolling" "trname" {
  fosid = 1
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `fsso_polling` - Fsso Polling.

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser FssoPollingAdgrp can be imported using any of these accepted formats:
```
Set import_options = ["fsso_polling=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_fssopolling_adgrp.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
