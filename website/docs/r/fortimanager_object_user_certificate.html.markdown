---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_certificate"
description: |-
  Configure certificate users.
---

# fortimanager_object_user_certificate
Configure certificate users.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `common_name` - Certificate common name.
* `fosid` - Id.
* `issuer` - CA certificate used for client certificate verification.
* `name` - User name.
* `status` - Enable/disable allowing the certificate user to authenticate with the FortiGate unit. Valid values: `disable`, `enable`.

* `type` - Type of certificate authentication method. Valid values: `single-certificate`, `trusted-issuer`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Certificate can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_certificate.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
