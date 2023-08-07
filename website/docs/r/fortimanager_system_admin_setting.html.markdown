---
subcategory: "System Admin"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_setting"
description: |-
  Admin setting.
---

# fortimanager_system_admin_setting
Admin setting.

## Example Usage

```hcl
resource "fortimanager_system_admin_setting" "trname" {
  idle_timeout = "400"
}
```

## Argument Reference


The following arguments are supported:


* `access_banner` - Enable/disable access banner. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `admin_https_redirect` - Enable/disable redirection of HTTP admin traffic to HTTPS. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `admin_login_max` - Maximum number admin users logged in at one time (1 - 256).
* `admin_server_cert` - HTTPS & Web Service server certificate.
* `allow_register` - Enable/disable allowance of register an unregistered device. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `auth_addr` - IP which is used by FGT to authorize FMG.
* `auth_port` - Port which is used by FGT to authorize FMG.
* `auto_update` - Enable/disable FortiGate automatic update. disable - Disable device automatic update. enable - Enable device automatic update. Valid values: `disable`, `enable`.

* `banner_message` - Banner message.
* `central_ftgd_local_cat_id` - Central FortiGuard local category id management, and do not auto assign id during installation. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `chassis_mgmt` - Enable or disable chassis management. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `chassis_update_interval` - Chassis background update interval (4 - 1440 mins).
* `device_sync_status` - Enable/disable device synchronization status indication. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `firmware_upgrade_check` - Enable/disable firmware upgrade check. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fsw_ignore_platform_check` - Enable/disable FortiSwitch Manager switch platform support check. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `gui_theme` - Color scheme to use for the administration GUI. blue - Blueberry green - Kiwi red - Cherry melongene - Plum spring - Spring summer - Summer autumn - Autumn winter - Winter space - Space calla-lily - Calla Lily binary-tunnel - Binary Tunnel diving - Diving dreamy - Dreamy technology - Technology landscape - Landscape twilight - Twilight canyon - Canyon northern-light - Northern Light astronomy - Astronomy fish - Fish penguin - Penguin mountain - Mountain polar-bear - Polar Bear parrot - Parrot cave - Cave zebra - Zebra contrast-dark - High Contrast Dark Valid values: `blue`, `green`, `red`, `melongene`, `spring`, `summer`, `autumn`, `winter`, `space`, `calla-lily`, `binary-tunnel`, `diving`, `dreamy`, `technology`, `landscape`, `twilight`, `canyon`, `northern-light`, `astronomy`, `fish`, `penguin`, `mountain`, `polar-bear`, `parrot`, `cave`, `zebra`, `contrast-dark`.

* `http_port` - HTTP port.
* `https_port` - HTTPS port.
* `idle_timeout` - Idle timeout (1 - 480 min).
* `idle_timeout_api` - Idle timeout for API sessions (1 - 28800 sec).
* `idle_timeout_gui` - Idle timeout for GUI sessions (60 - 28800 sec).
* `idle_timeout_sso` - Idle timeout for SSO sessions (60 - 28800 sec).
* `install_ifpolicy_only` - Allow install interface policy only. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `mgmt_addr` - IP of FortiManager used by FGFM.
* `mgmt_fqdn` - FQDN of FortiManager used by FGFM.
* `objects_force_deletion` - Enable/disable used objects force deletion. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `offline_mode` - Enable/disable offline mode. disable - Disable offline mode. enable - Enable offline mode. Valid values: `disable`, `enable`.

* `preferred_fgfm_intf` - Preferred interface for FGFM connection.
* `register_passwd` - Password for register a device.
* `rtm_max_monitor_by_days` - Maximum rtm monitor (sdwan, traffic shaping, etc) history by days (1 - 180).
* `rtm_temp_file_limit` - Set rtm monitor temp file limit by hours. Lower value will reduce disk usage, but may cause data loss (1 - 120).
* `sdwan_monitor_history` - Enable/disable sdwan-monitor-history. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `sdwan_skip_unmapped_input_device` - Skip unmapped interface for sdwan/rule/input-device instead of report mapping error. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `shell_access` - Enable/disable shell access. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `shell_password` - Password for shell access.
* `show_add_multiple` - Show add multiple button. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_adom_devman` - Show ADOM device manager tools on GUI. disable - Hide device manager tools on GUI. enable - Show device manager tools on GUI. Valid values: `disable`, `enable`.

* `show_checkbox_in_table` - Show checkboxs in tables on GUI. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_device_import_export` - Enable/disable import/export of ADOM, device, and group lists. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_fct_manager` - Enable/disable FCT manager. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_hostname` - Enable/disable hostname display in the GUI login page. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `show_automatic_script` - Enable/disable automatic script. disable - Disable script option. enable - Enable script option. Valid values: `disable`, `enable`.

* `show_grouping_script` - Enable/disable grouping script. disable - Disable script option. enable - Enable script option. Valid values: `disable`, `enable`.

* `show_schedule_script` - Enable or disable schedule script. disable - Disable script option. enable - Enable script option. Valid values: `disable`, `enable`.

* `show_tcl_script` - Enable/disable TCL script. disable - Disable script option. enable - Enable script option. Valid values: `disable`, `enable`.

* `traffic_shaping_history` - Enable/disable traffic-shaping-history. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `unreg_dev_opt` - Action to take when unregistered device connects to FortiManager. add_no_service - Add unregistered devices but deny service requests. ignore - Ignore unregistered devices. add_allow_service - Add unregistered devices and allow service requests. Valid values: `add_no_service`, `ignore`, `add_allow_service`.

* `webadmin_language` - Web admin language. auto_detect - Automatically detect language. english - English. simplified_chinese - Simplified Chinese. traditional_chinese - Traditional Chinese. japanese - Japanese. korean - Korean. spanish - Spanish. Valid values: `auto_detect`, `english`, `simplified_chinese`, `traditional_chinese`, `japanese`, `korean`, `spanish`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AdminSetting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_setting.labelname SystemAdminSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

