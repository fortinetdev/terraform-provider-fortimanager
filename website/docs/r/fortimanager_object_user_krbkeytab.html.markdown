---
subcategory: "ObjectUser"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_krbkeytab"
description: |-
  Configure Kerberos keytab entries.
---

# fortimanager_object_user_krbkeytab
Configure Kerberos keytab entries.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `keytab` - base64 coded keytab file containing a pre-shared key.
* `ldap_server` - LDAP server name.
* `name` - Kerberos keytab entry name.
* `pac_data` - Enable/disable parsing PAC data in the ticket. Valid values: `disable`, `enable`.

* `password` - Password for keytab.
* `principal` - Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser KrbKeytab can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_krbkeytab.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
