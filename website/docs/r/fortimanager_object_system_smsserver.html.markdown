---
subcategory: "Object System"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_smsserver"
description: |-
  Configure SMS server for sending SMS messages to support user authentication.
---

# fortimanager_object_system_smsserver
Configure SMS server for sending SMS messages to support user authentication.

## Example Usage

```hcl
resource "fortimanager_object_system_smsserver" "trname" {
  mail_server = "terraform-tefv"
  name        = "terraform-tefv-smsserver"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `mail_server` - Email-to-SMS server domain name.
* `name` - Name of SMS server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem SmsServer can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_smsserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
