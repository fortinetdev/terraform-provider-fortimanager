---
subcategory: "System Admin"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_user"
description: |-
  Admin user.
---

# fortimanager_system_admin_user
Admin user.

## Example Usage

```hcl
resource "fortimanager_system_admin_user" "trname" {
  password  = ["fortinet"]
  profileid = "Super_User"
  userid    = "tfuser"
}
```

## Argument Reference


The following arguments are supported:


* `fmgadom` - Adom. The structure of `fmgadom` block is documented below.
* `adom_exclude` - Adom-Exclude. The structure of `adom_exclude` block is documented below.
* `app_filter` - App-Filter. The structure of `app_filter` block is documented below.
* `avatar` - Image file for avatar (maximum 4K base64 encoded).
* `ca` - PKI user certificate CA (CA name in local).
* `change_password` - Enable/disable restricted user to change self password. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dashboard` - Dashboard. The structure of `dashboard` block is documented below.
* `dashboard_tabs` - Dashboard-Tabs. The structure of `dashboard_tabs` block is documented below.
* `description` - Description.
* `dev_group` - device group.
* `email_address` - Email address.
* `ext_auth_accprofile_override` - Allow to use the access profile provided by the remote authentication server. disable - Disable access profile override. enable - Enable access profile override. Valid values: `disable`, `enable`.

* `ext_auth_adom_override` - Allow to use the ADOM provided by the remote authentication server. disable - Disable ADOM override. enable - Enable ADOM override. Valid values: `disable`, `enable`.

* `ext_auth_group_match` - Only administrators belonging to this group can login.
* `first_name` - First name.
* `force_password_change` - Enable/disable force password change on next login. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `group` - Group name.
* `hidden` - Hidden administrator.
* `ips_filter` - Ips-Filter. The structure of `ips_filter` block is documented below.
* `ipv6_trusthost1` - Admin user trusted host IPv6, default ::/0 for all.
* `ipv6_trusthost10` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost2` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost3` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost4` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost5` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost6` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost7` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost8` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `ipv6_trusthost9` - Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.
* `last_name` - Last name.
* `ldap_server` - LDAP server name.
* `login_max` - Max login session for this user.
* `meta_data` - Meta-Data. The structure of `meta_data` block is documented below.
* `mobile_number` - Mobile number.
* `pager_number` - Pager number.
* `password` - Password.
* `password_expire` - Password expire time in GMT.
* `phone_number` - Phone number.
* `policy_package` - Policy-Package. The structure of `policy_package` block is documented below.
* `profileid` - Profile ID.
* `radius_server` - RADIUS server name.
* `rpc_permit` - set none/read/read-write rpc-permission. read-write - Read-write permission. none - No permission. read - Read-only permission. Valid values: `read-write`, `none`, `read`.

* `ssh_public_key1` - SSH public key 1.
* `ssh_public_key2` - SSH public key 2.
* `ssh_public_key3` - SSH public key 3.
* `subject` - PKI user certificate name constraints.
* `tacacs_plus_server` - TACACS+ server name.
* `trusthost1` - Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
* `trusthost10` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost2` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost3` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost4` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost5` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost6` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost7` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost8` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost9` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `two_factor_auth` - Enable 2-factor authentication (certificate + password). disable - Disable 2-factor authentication. enable - Enable 2-factor authentication. Valid values: `disable`, `enable`.

* `use_global_theme` - Enable/disble global theme for administration GUI. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `user_theme` - Color scheme to use for the admin user GUI. blue - Blueberry green - Kiwi red - Cherry melongene - Plum spring - Spring summer - Summer autumn - Autumn winter - Winter space - Space calla-lily - Calla Lily binary-tunnel - Binary Tunnel diving - Diving dreamy - Dreamy technology - Technology landscape - Landscape twilight - Twilight canyon - Canyon northern-light - Northern Light astronomy - Astronomy fish - Fish penguin - Penguin mountain - Mountain polar-bear - Polar Bear parrot - Parrot cave - Cave zebra - Zebra contrast-dark - High Contrast Dark Valid values: `blue`, `green`, `red`, `melongene`, `spring`, `summer`, `autumn`, `winter`, `space`, `calla-lily`, `binary-tunnel`, `diving`, `dreamy`, `technology`, `landscape`, `twilight`, `canyon`, `northern-light`, `astronomy`, `fish`, `penguin`, `mountain`, `polar-bear`, `parrot`, `cave`, `zebra`, `contrast-dark`.

* `user_type` - User type. local - Local user. radius - RADIUS user. ldap - LDAP user. tacacs-plus - TACACS+ user. pki-auth - PKI user. group - Group user. sso - SSO user. Valid values: `local`, `radius`, `ldap`, `tacacs-plus`, `pki-auth`, `group`, `sso`.

* `userid` - User name.
* `web_filter` - Web-Filter. The structure of `web_filter` block is documented below.
* `wildcard` - Enable/disable wildcard remote authentication. disable - Disable username wildcard. enable - Enable username wildcard. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `fmgadom` block supports:

* `adom_name` - Admin domain names.

The `adom_exclude` block supports:

* `adom_name` - Admin domain names.

The `app_filter` block supports:

* `app_filter_name` - App filter name.

The `dashboard` block supports:

* `column` - Widget's column ID.
* `diskio_content_type` - Disk I/O Monitor widget's chart type. util - bandwidth utilization. iops - the number of I/O requests. blks - the amount of data of I/O requests. Valid values: `util`, `iops`, `blks`.

* `diskio_period` - Disk I/O Monitor widget's data period. 1hour - 1 hour. 8hour - 8 hour. 24hour - 24 hour. Valid values: `1hour`, `8hour`, `24hour`.

* `log_rate_period` - Log receive monitor widget's data period. 2min  - 2 minutes. 1hour - 1 hour. 6hours - 6 hours. Valid values: `2min `, `1hour`, `6hours`.

* `log_rate_topn` - Log receive monitor widget's number of top items to display. 1 - Top 1. 2 - Top 2. 3 - Top 3. 4 - Top 4. 5 - Top 5. Valid values: `1`, `2`, `3`, `4`, `5`.

* `log_rate_type` - Log receive monitor widget's statistics breakdown options. log - Show log rates for each log type. device - Show log rates for each device. Valid values: `log`, `device`.

* `moduleid` - Widget ID.
* `name` - Widget name.
* `num_entries` - Number of entries.
* `refresh_interval` - Widget's refresh interval.
* `res_cpu_display` - Widget's CPU display type. average  - Average usage of CPU. each - Each usage of CPU. Valid values: `average `, `each`.

* `res_period` - Widget's data period. 10min  - Last 10 minutes. hour - Last hour. day - Last day. Valid values: `10min `, `hour`, `day`.

* `res_view_type` - Widget's data view type. real-time  - Real-time view. history - History view. Valid values: `real-time `, `history`.

* `status` - Widget's opened/closed state. close - Widget closed. open - Widget opened. Valid values: `close`, `open`.

* `tabid` - ID of tab where widget is displayed.
* `time_period` - Log Database Monitor widget's data period. 1hour - 1 hour. 8hour - 8 hour. 24hour - 24 hour. Valid values: `1hour`, `8hour`, `24hour`.

* `widget_type` - Widget type. top-lograte - Log Receive Monitor. sysres - System resources. sysinfo - System Information. licinfo - License Information. jsconsole - CLI Console. sysop - Unit Operation. alert - Alert Message Console. statistics - Statistics. rpteng - Report Engine. raid - Disk Monitor. logrecv - Logs/Data Received. devsummary - Device Summary. logdb-perf - Log Database Performance Monitor. logdb-lag - Log Database Lag Time. disk-io - Disk I/O. log-rcvd-fwd - Log receive and forwarding Monitor. Valid values: `top-lograte`, `sysres`, `sysinfo`, `licinfo`, `jsconsole`, `sysop`, `alert`, `statistics`, `rpteng`, `raid`, `logrecv`, `devsummary`, `logdb-perf`, `logdb-lag`, `disk-io`, `log-rcvd-fwd`.


The `dashboard_tabs` block supports:

* `name` - Tab name.
* `tabid` - Tab ID.

The `ips_filter` block supports:

* `ips_filter_name` - IPS filter name.

The `meta_data` block supports:

* `fieldlength` - Field length.
* `fieldname` - Field name.
* `fieldvalue` - Field value.
* `importance` - Importance. optional - This field is optional. required - This field is required. Valid values: `optional`, `required`.

* `status` - Status. disabled - This field is disabled. enabled - This field is enabled. Valid values: `disabled`, `enabled`.


The `policy_package` block supports:

* `policy_package_name` - Policy package names.

The `web_filter` block supports:

* `web_filter_name` - Web filter name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{userid}}.

## Import

System AdminUser can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_user.labelname {{userid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

