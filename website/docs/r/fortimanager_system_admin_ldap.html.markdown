---
subcategory: "System Admin"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_ldap"
description: |-
  LDAP server entry configuration.
---

# fortimanager_system_admin_ldap
LDAP server entry configuration.

## Example Usage

```hcl
resource "fortimanager_system_admin_ldap" "trname" {
  name     = "tldap"
  password = ["Fortinet"]
  port     = 390
  server   = "terraform"
  type     = "regular"
  username = "terraform-username"
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Adom. The structure of `fmgadom` block is documented below.
* `adom_attr` - Attribute used to retrieve adom
* `attributes` - Attributes used for group searching.
* `ca_cert` - CA certificate name.
* `cnid` - Common Name Identifier (default = CN).
* `connect_timeout` - LDAP connection timeout (msec).
* `dn` - Distinguished Name.
* `filter` - Filter used for group searching.
* `group` - Full base DN used for group searching.
* `memberof_attr` - Attribute used to retrieve memeberof.
* `name` - LDAP server entry name.
* `password` - Password for initial binding.
* `port` - Port number of LDAP server (default = 389).
* `profile_attr` - Attribute used to retrieve admin profile.
* `secondary_server` - {<name_str|ip_str>} secondary LDAP server domain name or IP.
* `secure` - SSL connection. disable - No SSL. starttls - Use StartTLS. ldaps - Use LDAPS. Valid values: `disable`, `starttls`, `ldaps`.

* `server` - {<name_str|ip_str>} LDAP server domain name or IP.
* `tertiary_server` - {<name_str|ip_str>} tertiary LDAP server domain name or IP.
* `type` - Type of LDAP binding. simple - Simple password authentication without search. anonymous - Bind using anonymous user search. regular - Bind using username/password and then search. Valid values: `simple`, `anonymous`, `regular`.

* `username` - Username (full DN) for initial binding.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fmgadom` block supports:

* `adom_name` - Admin domain names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminLdap can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_ldap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

