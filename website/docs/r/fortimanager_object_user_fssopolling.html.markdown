---
subcategory: "Object User"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_user_fssopolling"
description: |-
  Configure FSSO active directory servers for polling mode.
---

# fortimanager_object_user_fssopolling
Configure FSSO active directory servers for polling mode.

## Example Usage

```hcl
resource "fortimanager_object_user_fssopolling" "trname" {
  default_domain    = "terr-user-fsso-polling"
  fosid             = 1
  logon_history     = 8
  password          = ["fortinet"]
  polling_frequency = 10
  server            = "terraform-server"
  smb_ntlmv1_auth   = "disable"
  smbv1             = "disable"
  status            = "enable"
  user              = "admin"
}
```

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `_gui_meta` - _Gui_Meta.
* `adgrp` - Adgrp. The structure of `adgrp` block is documented below.
* `default_domain` - Default domain managed by this Active Directory server.
* `fosid` - Active Directory server ID.
* `ldap_server` - LDAP server name used in LDAP connection strings.
* `logon_history` - Number of hours of logon history to keep, 0 means keep all history.
* `password` - Password required to log into this Active Directory server
* `polling_frequency` - Polling frequency (every 1 to 30 seconds).
* `port` - Port to communicate with this Active Directory server.
* `server` - Host name or IP address of the Active Directory server.
* `smb_ntlmv1_auth` - Enable/disable support of NTLMv1 for Samba authentication. Valid values: `disable`, `enable`.

* `smbv1` - Enable/disable support of SMBv1 for Samba. Valid values: `disable`, `enable`.

* `status` - Enable/disable polling for the status of this Active Directory server. Valid values: `disable`, `enable`.

* `user` - User name required to log into this Active Directory server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `adgrp` block supports:

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ObjectUser FssoPolling can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_user_fssopolling.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
