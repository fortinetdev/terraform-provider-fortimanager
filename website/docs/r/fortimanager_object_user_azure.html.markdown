---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_azure"
description: |-
  ObjectUser Azure
---

# fortimanager_object_user_azure
ObjectUser Azure

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fortimanager_object_user_azure_rule`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `spn` - Spn.
* `alias` - Alias.
* `kbconfig` - Kbconfig.
* `name` - Name.
* `page_size` - Page_Size.
* `password` - Password.
* `proxy_enable` - Proxy_Enable. Valid values: `disable`, `enable`.

* `proxy_host` - Proxy_Host.
* `proxy_passwd` - Proxy_Passwd.
* `proxy_scheme` - Proxy_Scheme. Valid values: `http`, `https`.

* `proxy_user` - Proxy_User.
* `realm` - Realm.
* `region` - Region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.

* `rule` - Rule. The structure of `rule` block is documented below.
* `select_proxy` - Select_Proxy. Valid values: `basic`, `kerberos`.

* `status` - Status. Valid values: `disable`, `enable`.

* `tenantid` - Tenantid.
* `upd_interval` - Upd_Interval.
* `user` - User.
* `verifycert` - Verifycert. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `name` - Name.
* `rule` - Rule.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectUser Azure can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_azure.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
