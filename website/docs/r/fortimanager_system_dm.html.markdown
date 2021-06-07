---
subcategory: "System Others"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_dm"
description: |-
  Configure dm.
---

# fortimanager_system_dm
Configure dm.

## Argument Reference


The following arguments are supported:


* `concurrent_install_image_limit` - Maximum number of concurrent install image (1 - 1000).
* `concurrent_install_limit` - Maximum number of concurrent installs (5 - 2000).
* `concurrent_install_script_limit` - Maximum number of concurrent install scripts (5 - 2000).
* `conf_merge_after_script` - Merge config after run script on remote device, instead of full retrieve. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `discover_timeout` - Check connection timeout when discover device (3 - 15).
* `dpm_logsize` - Maximum dpm log size per device (1 - 100000K).
* `fgfm_install_refresh_count` - Maximum FGFM install refresh attempt.
* `fgfm_sock_timeout` - Maximum FGFM socket idle time (90 - 1800 sec).
* `fgfm_keepalive_itvl` - FGFM protocol keep alive interval (30 - 600 sec).
* `force_remote_diff` - Always use remote diff when installing. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `fortiap_refresh_cnt` - Max auto refresh FortiAP number each time (1 - 10000).
* `fortiap_refresh_itvl` - Auto refresh FortiAP status interval (0 - 1440) minutes, set to 0 will disable auto refresh.
* `fortiext_refresh_cnt` - Max device number for FortiExtender auto refresh (1 - 10000).
* `install_image_timeout` - Maximum waiting time for image transfer and device upgrade (10*60 - 24*60*60 seconds).
* `install_tunnel_retry_itvl` - Time to re-establish tunnel during install (10 - 60 sec).
* `max_revs` - Maximum number of revisions saved (1 - 250).
* `nr_retry` - Number of retries.
* `retry` - Enable/disable configuration install retry. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `retry_intvl` - Retry interval.
* `rollback_allow_reboot` - Enable/disable FortiGate reboot to rollback when installing script/config. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `script_logsize` - Maximum script log size per device (1 - 10000K).
* `skip_scep_check` - Enable/disable installing scep related objects even if scep url is configured. disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `skip_tunnel_fcp_req` - Enable/disable skip the fcp request sent from fgfm tunnel disable - Disable. enable - Enable. Valid values: `disable`, `enable`.

* `verify_install` - Verify install against remote configuration. disable - Disable. optimal - Verify installation for command errors. enable - Always verify installation. Valid values: `disable`, `optimal`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dm can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_dm.labelname SystemDm
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

