---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_domaincontroller"
description: |-
  Configure domain controller entries.
---

# fortimanager_object_user_domaincontroller
Configure domain controller entries.

## Example Usage

```hcl
resource "fortimanager_object_user_domaincontroller" "trname" {
  domain_name = "admin"
  ip_address  = "192.168.1.1"
  name        = "terr-user-domain-controller"
  port        = 445
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `domain_name` - Domain DNS name.
* `extra_server` - Extra-Server. The structure of `extra_server` block is documented below.
* `ip_address` - Domain controller IP address.
* `ldap_server` - LDAP server name.
* `name` - Domain controller entry name.
* `port` - Port to be used for communication with the domain controller (default = 445).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `extra_server` block supports:

* `id` - Server ID.
* `ip_address` - Domain controller IP address.
* `port` - Port to be used for communication with the domain controller (default = 445).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser DomainController can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_domaincontroller.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
