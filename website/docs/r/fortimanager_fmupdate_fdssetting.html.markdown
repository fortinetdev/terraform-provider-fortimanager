---
subcategory: "Fmupdate"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_fmupdate_fdssetting"
description: |-
  Configure FortiGuard settings.
---

# fortimanager_fmupdate_fdssetting
Configure FortiGuard settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `push_override`: `fortimanager_fmupdate_fdssetting_pushoverride`
>- `push_override_to_client`: `fortimanager_fmupdate_fdssetting_pushoverridetoclient`
>- `server_override`: `fortimanager_fmupdate_fdssetting_serveroverride`
>- `update_schedule`: `fortimanager_fmupdate_fdssetting_updateschedule`



## Example Usage

```hcl
resource "fortimanager_fmupdate_fdssetting" "trname" {
  fds_clt_ssl_protocol = "sslv3"
  fds_ssl_protocol     = "sslv3"
  fmtr_log             = "emergency"
}
```

## Argument Reference


The following arguments are supported:


* `user_agent` - Configure the user agent string.
* `fds_clt_ssl_protocol` - The SSL protocols version for connecting fds server (default = tlsv1.2). sslv3 - set SSLv3 as the client version. tlsv1.0 - set TLSv1.0 as the client version. tlsv1.1 - set TLSv1.1 as the client version. tlsv1.2 - set TLSv1.2 as the client version (default). tlsv1.3 - set TLSv1.3 as the client version. Valid values: `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.

* `fds_ssl_protocol` - The SSL protocols version for receiving fgt connection (default = tlsv1.2). sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version (default). tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: `sslv3`, `tlsv1.0`, `tlsv1.1`, `tlsv1.2`, `tlsv1.3`.

* `fmtr_log` - fmtr log level emergency - Log level - emergency alert - Log level - alert critical - Log level - critical error - Log level - error warn - Log level - warn notice - Log level - notice info - Log level - info debug - Log level - debug disable - Disable linkd log Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `fortiguard_anycast` - Enable/disable use of FortiGuard's anycast network disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet fortinet - Use Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. aws - Use Fortinet's AWS servers to provide FortiGuard services in FortiGuard's anycast network. Valid values: `fortinet`, `aws`.

* `linkd_log` - The linkd log level (default = info). emergency - Log level - emergency alert - Log level - alert critical - Log level - critical error - Log level - error warn - Log level - warn notice - Log level - notice info - Log level - info debug - Log level - debug disable - Disable linkd log Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `max_av_ips_version` - The maximum number of downloadable, full version AV/IPS packages (1 - 1000, default = 20).
* `max_work` - The maximum number of worker processing download requests (1 - 32, default = 1).
* `push_override` - Push-Override. The structure of `push_override` block is documented below.
* `push_override_to_client` - Push-Override-To-Client. The structure of `push_override_to_client` block is documented below.
* `send_report` - send report/fssi to fds server. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `send_setup` - forward setup to fds server. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `server_override` - Server-Override. The structure of `server_override` block is documented below.
* `system_support_fai` - Supported FortiNDR versions. 7.x - Support version 7.x Valid values: `7.x`.

* `system_support_faz` - Supported FortiAnalyzer versions. 6.x - Support version 6.x 7.x - Support version 7.x Valid values: `6.x`, `7.x`.

* `system_support_fct` - Supported FortiClient versions. 4.x - Support version 4.x 5.0 - Support version 5.0 5.2 - Support version 5.2 5.4 - Support version 5.4 5.6 - Support version 5.6 6.0 - Support version 6.0 6.2 - Support version 6.2 6.4 - Support version 6.4 Valid values: `4.x`, `5.0`, `5.2`, `5.4`, `5.6`, `6.0`, `6.2`, `6.4`.

* `system_support_fdc` - Supported FortiDeceptor versions. 3.x - Support version 3.x Valid values: `3.x`.

* `system_support_fgt` - Supported FortiOS versions. 5.4 - Support version 5.4 5.6 - Support version 5.6 6.0 - Support version 6.0 6.2 - Support version 6.2 6.4 - Support version 6.4 Valid values: `5.4`, `5.6`, `6.0`, `6.2`, `6.4`.

* `system_support_fis` - Supported FortiIsolator versions. 1.x - Support version 1.x 2.x - Support version 2.x Valid values: `1.x`, `2.x`.

* `system_support_fml` - Supported FortiMail versions. 4.x - Support version 4.x 5.x - Support version 5.x 6.x - Support version 6.x Valid values: `4.x`, `5.x`, `6.x`.

* `system_support_fsa` - Supported FortiSandbox versions. 1.x - Support version 1.x 2.x - Support version 2.x 3.x - Support version 3.x Valid values: `1.x`, `2.x`, `3.x`.

* `system_support_fts` - Supported FortiTester versions. 4.x - Support version 4.x Valid values: `4.x`.

* `system_support_fsw` - Supported FortiSwitch versions. 4.x - Support version 4.x 5.0 - Support version 5.0 5.2 - Support version 5.2 5.4 - Support version 5.4 5.6 - Support version 5.6 6.0 - Support version 6.0 6.2 - Support version 6.2 6.4 - Support version 6.4 Valid values: `4.x`, `5.0`, `5.2`, `5.4`, `5.6`, `6.0`, `6.2`, `6.4`.

* `umsvc_log` - The um_service log level (default = info). emergency - Log level - emergency alert - Log level - alert critical - Log level - critical error - Log level - error warn - Log level - warn notice - Log level - notice info - Log level - info debug - Log level - debug disable - Disable linkd log Valid values: `emergency`, `alert`, `critical`, `error`, `warn`, `notice`, `info`, `debug`, `disable`.

* `unreg_dev_option` - set the option for unregister devices ignore - Ignore all unregistered devices. svc-only - Allow update requests without adding the device. add-service - Add unregistered devices and allow update request. Valid values: `ignore`, `svc-only`, `add-service`.

* `update_schedule` - Update-Schedule. The structure of `update_schedule` block is documented below.
* `wanip_query_mode` - public ip query mode disable - Do not query public ip ipify - Get public IP through https://api.ipify.org Valid values: `disable`, `ipify`.


The `push_override` block supports:

* `ip` - External or virtual IP address of the NAT device that will forward push messages to the FortiManager unit.
* `port` - Receiving port number on the NAT device (1 - 65535, default = 9443).
* `status` - Enable/disable push updates for clients (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.


The `push_override_to_client` block supports:

* `announce_ip` - Announce-Ip. The structure of `announce_ip` block is documented below.
* `status` - Enable/disable push updates (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.


The `announce_ip` block supports:

* `id` - ID of the announce IP address (1 - 10).
* `ip` - Announce IPv4 address.
* `port` - Announce IP port (1 - 65535, default = 8890).

The `server_override` block supports:

* `servlist` - Servlist. The structure of `servlist` block is documented below.
* `status` - Override status. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.


The `servlist` block supports:

* `id` - Override server ID (1 - 10).
* `ip` - IPv4 address of the override server.
* `ip6` - IPv6 address of the override server.
* `port` - Port number to use when contacting FortiGuard (1 - 65535, default = 443).
* `service_type` - Override service type. fct - Server override config for fct fds - Server override config for fds Valid values: `fct`, `fds`.


The `update_schedule` block supports:

* `day` - Configure the day the update will occur, if the freqnecy is weekly (Sunday - Saturday, default = Monday). Sunday - Update every Sunday. Monday - Update every Monday. Tuesday - Update every Tuesday. Wednesday - Update every Wednesday. Thursday - Update every Thursday. Friday - Update every Friday. Saturday - Update every Saturday. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.

* `frequency` - Configure update frequency: every - time interval, daily - once a day, weekly - once a week (default = every). every - Time interval. daily - Every day. weekly - Every week. Valid values: `every`, `daily`, `weekly`.

* `status` - Enable/disable scheduled updates. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `time` - Time interval between updates, or the hour and minute when the update occurs (hh: 0 - 23, mm: 0 - 59 or 60 = random, default = 00:10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate FdsSetting can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_fmupdate_fdssetting.labelname FmupdateFdsSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

