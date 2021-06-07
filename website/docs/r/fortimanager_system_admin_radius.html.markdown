---
subcategory: "System Admin"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_radius"
description: |-
  Configure radius.
---

# fortimanager_system_admin_radius
Configure radius.

## Argument Reference


The following arguments are supported:


* `auth_type` - Authentication protocol. any - Use any supported authentication protocol. pap - PAP. chap - CHAP. mschap2 - MSCHAPv2. Valid values: `any`, `pap`, `chap`, `mschap2`.

* `name` - Name.
* `nas_ip` - NAS IP address and called station ID.
* `port` - Server port.
* `secondary_secret` - Secondary server secret.
* `secondary_server` - Secondary server name/IP.
* `secret` - Server secret.
* `server` - Server name/IP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminRadius can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_radius.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

