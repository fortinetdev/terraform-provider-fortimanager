---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_backup_allsettings"
description: |-
  Scheduled backup settings.
---

# fortimanager_system_backup_allsettings
Scheduled backup settings.

## Example Usage

```hcl
resource "fortimanager_system_backup_allsettings" "trname" {
  crptpasswd = ["fortinet"]
  directory  = "terraform"
  passwd     = ["fortinet"]
  protocol   = "ftp"
  server     = "192.168.1.1"
  status     = "enable"
  time       = "1:00:00"
  user       = "admin"
  week_days  = ["monday"]
}
```

## Argument Reference


The following arguments are supported:


* `cert` - SSH certificate for authentication.
* `crptpasswd` - Optional password to protect backup content.
* `directory` - Directory in which file will be stored on backup server.
* `passwd` - Backup server login user password.
* `protocol` - Protocol used to backup. sftp - SFTP. ftp - FTP. scp - SCP. Valid values: `sftp`, `ftp`, `scp`.

* `server` - Backup server name/IP and port.
* `status` - Enable/disable schedule backup. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `time` - Time to backup.
* `user` - Backup server login user.
* `week_days` - Week days to backup. monday - Monday. tuesday - Tuesday. wednesday - Wednesday. thursday - Thursday. friday - Friday. saturday - Saturday. sunday - Sunday. Valid values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System BackupAllSettings can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_backup_allsettings.labelname SystemBackupAllSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

