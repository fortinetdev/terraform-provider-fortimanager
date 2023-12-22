---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fwmsetting_upgradetimeout"
description: |-
  Configure the timeout value of image upgrade process.
---

# fortimanager_fmupdate_fwmsetting_upgradetimeout
Configure the timeout value of image upgrade process.

~> This resource is a sub resource for variable `upgrade_timeout` of resource `fortimanager_fmupdate_fwmsetting`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fortimanager_fmupdate_fwmsetting_upgradetimeout" "trname" {
  check_status_timeout      = 600
  ctrl_check_status_timeout = 1200
}
```

## Argument Reference


The following arguments are supported:


* `check_status_timeout` - timeout for checking status after tunnnel is up.(1-6000s, default=600)
* `ctrl_check_status_timeout` - timeout for checking fap/fsw/fext status after request upgrade.(1-12000s, default=1200)
* `ctrl_put_image_by_fds_timeout` - timeout for waiting device get fap/fsw/fext image from fortiguard.(1-9000ss, default=900)
* `ha_sync_timeout` - timeout for waiting HA sync.(1-18000s, default=1800)
* `license_check_timeout` - timeout for waiting fortigate check license.(1-6000s, default=600)
* `prepare_image_timeout` - timeout for preparing image.(1-6000s, default=600)
* `put_image_by_fds_timeout` - timeout for waiting device get image from fortiguard.(1-18000s, default=1800)
* `put_image_timeout` - timeout for waiting send image over tunnel.(1-18000s, default=1800)
* `reboot_of_fsck_timeout` - timeout for waiting fortigate reboot.(1-18000s, default=1800)
* `reboot_of_upgrade_timeout` - timeout for waiting fortigate reboot after image upgrade.(1-12000s, default=1200)
* `retrieve_timeout` - timeout for waiting retrieve.(1-18000s, default=1800)
* `rpc_timeout` - timeout for waiting fortigate rpc response.(1-1800s, default=180)
* `total_timeout` - timeout for the whole fortigate upgrade(1-86400s, default=3600)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FwmSettingUpgradeTimeout can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fwmsetting_upgradetimeout.labelname FmupdateFwmSettingUpgradeTimeout
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

