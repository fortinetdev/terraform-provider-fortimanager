---
subcategory: "System LocalLog"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_locallog_memory_filter"
description: |-
  Filter for memory logging.
---

# fortimanager_system_locallog_memory_filter
Filter for memory logging.

## Argument Reference


The following arguments are supported:


* `aid` - Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `devcfg` - Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `devops` - Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `diskquota` - Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dm` - Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `docker` - Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dvm` - Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `ediscovery` - Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `epmgr` - Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `event` - Log event messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `eventmgmt` - Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `faz` - Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fazha` - Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fazsys` - Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fgd` - Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fgfm` - Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fips` - Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fmgws` - Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fmlmgr` - Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fmwmgr` - Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiview` - Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `glbcfg` - Log global database message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `ha` - Log HA message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `hcache` - Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `incident` - Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `iolog` - Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `logd` - Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `logdb` - Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `logdev` - Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `logfile` - Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: `enable`, `disable`.

* `logging` - Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `lrmgr` - Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `objcfg` - Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `report` - Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `rev` - Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `rtmon` - Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `scfw` - Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `scply` - Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `scrmgr` - Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `scvpn` - Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `system` - Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `webport` - Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogMemoryFilter can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_locallog_memory_filter.labelname SystemLocallogMemoryFilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

