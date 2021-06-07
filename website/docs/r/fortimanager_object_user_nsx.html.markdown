---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_nsx"
description: |-
  ObjectUser Nsx
---

# fortimanager_object_user_nsx
ObjectUser Nsx

## Example Usage

```hcl
resource "fortimanager_object_user_nsx" "labelname" {
  fmgip      = "1.1.1.1"
  fmgpasswd  = ["tesssssss"]
  fmguser    = "dfa"
  name       = "ewwe"
  password   = ["tesssssss"]
  server     = "3.3.3.3"
  service_id = []
  status     = "disable"
  user       = "sgic"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `fmgip` - Fmgip.
* `fmgpasswd` - Fmgpasswd.
* `fmguser` - Fmguser.
* `name` - Name.
* `password` - Password.
* `server` - Server.
* `service_id` - Service-Id.
* `status` - Status. Valid values: `disable`, `enable`.

* `user` - User.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Nsx can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_nsx.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
