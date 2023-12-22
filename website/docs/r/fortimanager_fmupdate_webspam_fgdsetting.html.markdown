---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_webspam_fgdsetting"
description: |-
  Configure the FortiGuard run parameters.
---

# fortimanager_fmupdate_webspam_fgdsetting
Configure the FortiGuard run parameters.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
`server_override`: `fortimanager_fmupdate_webspam_fgdsetting_serveroverride`



## Example Usage

```hcl
resource "fortimanager_fmupdate_webspam_fgdsetting" "trname" {
  av_cache   = "400"
  av_log     = "all"
  av_preload = "enable"
}
```

## Argument Reference


The following arguments are supported:


* `as_cache` - Antispam service maximum memory usage in megabytes (Maximum = Physical memory-1024, 0: no limit, default = 300).
* `as_log` - Antispam log setting (default = nospam). disable - Disable spam log. nospam - Log non-spam events. all - Log all spam lookups. Valid values: `disable`, `nospam`, `all`.

* `as_preload` - Enable/disable preloading antispam database to memory (default = disable). disable - Disable antispam database preload. enable - Enable antispam database preload. Valid values: `disable`, `enable`.

* `av_cache` - Antivirus service maximum memory usage, in megabytes (100 - 500, default = 300).
* `av_log` - Antivirus log setting (default = novirus). disable - Disable virus log. novirus - Log non-virus events. all - Log all virus lookups. Valid values: `disable`, `novirus`, `all`.

* `av_preload` - Enable/disable preloading antivirus database to memory (default = disable). disable - Disable antivirus database preload. enable - Enable antivirus database preload. Valid values: `disable`, `enable`.

* `av2_cache` - Antispam service maximum memory usage in megabytes (Maximum = Physical memory-1024, 0: no limit, default = 800).
* `av2_log` - Outbreak prevention log setting (default = noav2). disable - Disable av2 log. noav2 - Log non-av2 events. all - Log all av2 lookups. Valid values: `disable`, `noav2`, `all`.

* `av2_preload` - Enable/disable preloading outbreak prevention database to memory (default = disable). disable - Disable outbreak prevention database preload. enable - Enable outbreak prevention database preload. Valid values: `disable`, `enable`.

* `eventlog_query` - Enable/disable record query to event-log besides fgd-log (default = disable). disable - Record query to event-log besides fgd-log. enable - Do not log to event-log. Valid values: `disable`, `enable`.

* `fgd_pull_interval` - Fgd pull interval setting, in minutes (1 - 1440, default = 10).
* `fq_cache` - File query service maximum memory usage, in megabytes (100 - 500, default = 300).
* `fq_log` - File query log setting (default = nofilequery). disable - Disable file query log. nofilequery - Log non-file query events. all - Log all file query events. Valid values: `disable`, `nofilequery`, `all`.

* `fq_preload` - Enable/disable preloading file query database to memory (default = disable). disable - Disable file query db preload. enable - Enable file query db preload. Valid values: `disable`, `enable`.

* `iot_cache` - IoT service maximum memory usage, in megabytes (100 - 500, default = 300).
* `iot_log` - IoT log setting (default = nofilequery). disable - Disable IoT log. nofilequery - Log non-IoT events. all - Log all IoT events. Valid values: `disable`, `nofilequery`, `all`.

* `iot_preload` - Enable/disable preloading IoT database to memory (default = disable). disable - Disable IoT db preload. enable - Enable IoT db preload. Valid values: `disable`, `enable`.

* `iotv_preload` - Enable/disable preloading IoT-Vulnerability database to memory (default = disable). disable - Disable IoT-Vulnerability db preload. enable - Enable IoT-Vulnerability db preload. Valid values: `disable`, `enable`.

* `linkd_log` - Linkd log setting (default = debug). emergency - The unit is unusable. alert - Immediate action is required critical - Functionality is affected. error - Functionality is probably affected. warn - Functionality might be affected. notice - Information about normal events. info - General information. debug - Debug information. disable - Linkd logging is disabled. Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `max_client_worker` - max worker for tcp client connection (0~16: 0 means use cpu number up to 4).
* `max_log_quota` - Maximum log quota setting, in megabytes (100 - 20480, default = 6144).
* `max_unrated_site` - Maximum number of unrated site in memory, in kilobytes(10 - 5120, default = 500).
* `restrict_as1_dbver` - Restrict system update to indicated antispam(1) database version (character limit = 127).
* `restrict_as2_dbver` - Restrict system update to indicated antispam(2) database version (character limit = 127).
* `restrict_as4_dbver` - Restrict system update to indicated antispam(4) database version (character limit = 127).
* `restrict_av_dbver` - Restrict system update to indicated antivirus database version (character limit = 127).
* `restrict_av2_dbver` - Restrict system update to indicated outbreak prevention database version (character limit = 127).
* `restrict_fq_dbver` - Restrict system update to indicated file query database version (character limit = 127).
* `restrict_iots_dbver` - Restrict system update to indicated file query database version (character limit = 127).
* `restrict_wf_dbver` - Restrict system update to indicated web filter database version (character limit = 127).
* `server_override` - Server-Override. The structure of `server_override` block is documented below.
* `stat_log` - stat log setting (default = disable). emergency - The unit is unusable(0). alert - Immediate action is required(1) critical - Functionality is affected(2). error - Functionality is probably affected(3). warn - Functionality might be affected(4). notice - Information about normal events(5). info - General information(6). debug - Debug information(7). disable - Linkd logging is disabled. Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `stat_log_interval` - Statistic log interval setting, in minutes (1 - 1440, default = 60).
* `stat_sync_interval` - Synchronization interval for statistic of unrated site in minutes (1 - 60, default = 60).
* `update_interval` - FortiGuard database update wait time if not enough delta files, in hours (2 - 24, default = 6).
* `update_log` - Enable/disable update log setting (default = enable). disable - Disable update log. enable - Enable update log. Valid values: `disable`, `enable`.

* `wf_cache` - Web filter service maximum memory usage, in megabytes (maximum = Physical memory-1024, 0 = no limit, default = 600).
* `wf_dn_cache_expire_time` - Web filter DN cache expire time, in minutes (1 - 1440, 0 = never, default = 30).
* `wf_dn_cache_max_number` - Maximum number of Web filter DN cache (0 = disable, default = 10000).
* `wf_log` - Web filter log setting (default = nour1) disable - Disable URL log. nourl - Log non-URL events. all - Log all URL lookups. Valid values: `disable`, `nourl`, `all`.

* `wf_preload` - Enable/disable preloading the web filter database into memory (default = disable). disable - Disable web filter database preload. enable - Enable web filter database preload. Valid values: `disable`, `enable`.


The `server_override` block supports:

* `servlist` - Servlist. The structure of `servlist` block is documented below.
* `status` - Override status. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.


The `servlist` block supports:

* `id` - Override server ID (1 - 10).
* `ip` - IPv4 address of the override server.
* `ip6` - IPv6 address of the override server.
* `port` - Port number to use when contacting FortiGuard (1 - 65535, default = 443).
* `service_type` - Override service type. fgd - Server override config for fgd fgc - Server override config for fgc fsa - Server override config for fsa Valid values: `fgd`, `fgc`, `fsa`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate WebSpamFgdSetting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_webspam_fgdsetting.labelname FmupdateWebSpamFgdSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

