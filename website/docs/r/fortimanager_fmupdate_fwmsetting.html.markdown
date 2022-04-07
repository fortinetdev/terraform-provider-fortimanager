---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fwmsetting"
description: |-
  Configure firmware management settings.
---

# fortimanager_fmupdate_fwmsetting
Configure firmware management settings.

## Example Usage

```hcl
resource "fortimanager_fmupdate_fwmsetting" "trname" {
  auto_scan_fgt_disk = "enable"
  check_fgt_disk     = "enable"
  fds_failover_fmg   = "enable"
  immx_source        = "fmg"
}
```

## Argument Reference


The following arguments are supported:


* `auto_scan_fgt_disk` - auto scan fgt disk if needed. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `check_fgt_disk` - check fgt disk before upgrade image. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fds_failover_fmg` - using fmg local image file is download from fds fails. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fds_image_timeout` - timer for fgt download image from fortiguard (300-3600s default=1800)
* `immx_source` - Configure which of IMMX file to be used for choosing upgrade pach. Default is file for FortiManager fmg - Use IMMX file for FortiManager fgt - Use IMMX file for FortiGate cloud - Use IMMX file for FortiCloud Valid values: `fmg`, `fgt`, `cloud`.

* `log` - Configure log setting for fwm daemon fwm - FWM daemon log fwm_dm - FWM and Deployment service log fwm_dm_json - FWM and Deployment service log with JSON data between FMG-FGT Valid values: `fwm`, `fwm_dm`, `fwm_dm_json`.

* `multiple_steps_interval` - waiting time between multiple steps upgrade (30-180s, default=60)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FwmSetting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fwmsetting.labelname FmupdateFwmSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

