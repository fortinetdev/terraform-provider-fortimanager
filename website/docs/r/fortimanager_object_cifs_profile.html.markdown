---
subcategory: "Object Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_cifs_profile"
description: |-
  Configure CIFS profile.
---

# fortimanager_object_cifs_profile
Configure CIFS profile.

## Example Usage

```hcl
resource "fortimanager_object_cifs_profile" "trname" {
  name = "terr-cifs-profile"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `domain_controller` - Domain for which to decrypt CIFS traffic.
* `name` - Profile name.
* `server_credential_type` - CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.

* `server_keytab` - Server-Keytab. The structure of `server_keytab` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_keytab` block supports:

* `keytab` - Base64 encoded keytab file containing credential of the server.
* `password` - Password for keytab.
* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectCifs Profile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_cifs_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
