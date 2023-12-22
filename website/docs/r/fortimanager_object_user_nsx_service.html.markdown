---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_nsx_service"
description: |-
  ObjectUser NsxService
---

# fortimanager_object_user_nsx_service
ObjectUser NsxService

~> This resource is a sub resource for variable `service` of resource `fortimanager_object_user_nsx`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_object_user_nsx_service" "trname" {
  nsx         = fortimanager_object_user_nsx.trname.name
  fosid       = 1
  integration = "east-west"
  name        = "terr-service"
  depends_on  = [fortimanager_object_user_nsx.trname]
}

resource "fortimanager_object_user_nsx" "trname" {
  name      = "terr-nsx"
  user      = "admin"
  server    = "3.3.3.3"
  fmgip     = "1.1.1.1"
  fmgpasswd = ["psw"]
  fmguser   = "admin"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `nsx` - Nsx.

* `fosid` - Id.
* `integration` - Integration. Valid values: `east-west`, `north-south`.

* `name` - Name.
* `ref_id` - Ref-Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectUser NsxService can be imported using any of these accepted formats:
```
Set import_options = ["nsx=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_nsx_service.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
