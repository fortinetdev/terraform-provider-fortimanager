---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_sys_restore"
description: |-
  Restore FortiManager configuration backup.
---

# fortimanager_sys_restore
Restore FortiManager configuration backup.

## Argument Reference


The following arguments are supported:


* `filename` - Path and file name for the configuration on remote server.
* `passwd` - Password for password-protected configuration backup file.
* `port` - Remote server port. Default port is used if not specified.
* `server` - Remote server address.
* `service` - Protocol to restore configuration. Valid values: `ftp`, `scp`, `sftp`, `tftp`.

* `username` - User name to log in to the remote server.
* `userpasswd` - Password to log in to the remote server.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Sys Restore can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_sys_restore.labelname SysRestore
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
