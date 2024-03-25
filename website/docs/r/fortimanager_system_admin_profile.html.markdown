---
subcategory: "System Admin"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_system_admin_profile"
description: |-
  Admin profile.
---

# fortimanager_system_admin_profile
Admin profile.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `datamask_custom_fields`: `fortimanager_system_admin_profile_datamaskcustomfields`
>- `write_passwd_profiles`: `fortimanager_system_admin_profile_writepasswdprofiles`
>- `write_passwd_user_list`: `fortimanager_system_admin_profile_writepasswduserlist`



## Example Usage

```hcl
resource "fortimanager_system_admin_profile" "trname" {
  description = "terraform-tefv-description"
  profileid   = "terraform-tefv-profile2"
  scope       = "adom"
  type        = "system"
}
```

## Argument Reference


The following arguments are supported:


* `adom_lock` - ADOM locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `adom_policy_packages` - ADOM policy packages. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `adom_switch` - Administrator domain. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `allow_to_install` - Enable/disable the restricted user to install objects to the devices. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `app_filter` - App filter. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `assignment` - Assignment permission. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `change_password` - Enable/disable the user to change self password. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `config_retrieve` - Configuration retrieve. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `config_revert` - Revert Configuration from Revision History none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `consistency_check` - Consistency check. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `datamask` - Enable/disable data masking. disable - Disable data masking. enable - Enable data masking. Valid values: `disable`, `enable`.

* `datamask_custom_fields` - Datamask-Custom-Fields. The structure of `datamask_custom_fields` block is documented below.
* `datamask_custom_priority` - Prioritize custom fields. disable - Disable custom field search priority. enable - Enable custom field search priority. Valid values: `disable`, `enable`.

* `datamask_fields` - Data masking fields. user - User name. srcip - Source IP. srcname - Source name. srcmac - Source MAC. dstip - Destination IP. dstname - Dst name. email - Email. message - Message. domain - Domain. Valid values: `user`, `srcip`, `srcname`, `srcmac`, `dstip`, `dstname`, `email`, `message`, `domain`.

* `datamask_key` - Data masking encryption key.
* `datamask_unmasked_time` - Time in days without data masking.
* `deploy_management` - Install to devices. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `description` - Description.
* `device_ap` - Manage AP. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_config` - Manage device configurations. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_forticlient` - Manage FortiClient. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_fortiextender` - Manage FortiExtender. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_fortiswitch` - Manage FortiSwitch. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_manager` - Device manager. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_op` - Device add/delete/edit. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_policy_package_lock` - Device/Policy Package locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_profile` - Device profile permission. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_revision_deletion` - Delete device revision. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `device_wan_link_load_balance` - Manage WAN link load balance. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `event_management` - Event management. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `extension_access` - Manage extension access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fabric_viewer` - Fabric viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fgd_center_advanced` - FortiGuard Center Advanced. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fgd_center_fmw_mgmt` - FortiGuard Center Firmware Management. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fgd_center_licensing` - FortiGuard Center Licensing. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fgd_center` - FortiGuard Center. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `fgt_gui_proxy` - FortiGate GUI proxy. disable - No permission. enable - With permission. Valid values: `disable`, `enable`.

* `global_policy_packages` - Global policy packages. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `import_policy_packages` - Import Policy Package. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `intf_mapping` - Interface Mapping none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `ips_baseline_cfg` - Ips baseline sensor configration. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `ips_baseline_ovrd` - Enable/disable override baseline ips sensor. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `ips_filter` - IPS filter. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `ips_lock` - IPS locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `ips_objects` - Ips objects configuration. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

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
* `log_viewer` - Log viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `policy_ips_attrs` - Policy ips attributes configuration. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `policy_objects` - Policy objects permission. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `profileid` - Profile ID.
* `read_passwd` - View password in clear text. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `realtime_monitor` - Realtime monitor. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `report_viewer` - Report viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `rpc_permit` - Set none/read/read-write rpc-permission read-write - Read-write permission. none - No permission. read - Read-only permission. Valid values: `read-write`, `none`, `read`.

* `run_report` - Run reports. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `scope` - Scope. global - Global scope. adom - ADOM scope. Valid values: `global`, `adom`.

* `script_access` - Script access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `set_install_targets` - Edit installation targets. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `super_user_profile` - Enable/disable super user profile disable - Disable super user profile enable - Enable super user profile Valid values: `disable`, `enable`.

* `system_setting` - System setting. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `term_access` - Terminal access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `triage_events` - Triage events. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

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
* `type` - profile type. system - System admin. restricted - Restricted admin. Valid values: `system`, `restricted`.

* `update_incidents` - Create/update incidents. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `vpn_manager` - VPN manager. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: `none`, `read`, `read-write`.

* `web_filter` - Web filter. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `write_passwd_access` - set all/specify-by-user/specify-by-profile write password access mode. all - All except super users. specify-by-user - Specify by user. specify-by-profile - Specify by profile. Valid values: `all`, `specify-by-user`, `specify-by-profile`.

* `write_passwd_profiles` - Write-Passwd-Profiles. The structure of `write_passwd_profiles` block is documented below.
* `write_passwd_user_list` - Write-Passwd-User-List. The structure of `write_passwd_user_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `datamask_custom_fields` block supports:

* `field_category` - Field categories. log - Log. fortiview - FortiView. alert - Event management. ueba - UEBA. all - All. Valid values: `log`, `fortiview`, `alert`, `ueba`, `all`.

* `field_name` - Field name.
* `field_status` - Field status. disable - Disable field. enable - Enable field. Valid values: `disable`, `enable`.

* `field_type` - Field type. string - String. ip - IP. mac - MAC address. email - Email address. unknown - Unknown. Valid values: `string`, `ip`, `mac`, `email`, `unknown`.


The `write_passwd_profiles` block supports:

* `profileid` - Profile ID.

The `write_passwd_user_list` block supports:

* `userid` - User ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profileid}}.

## Import

System AdminProfile can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_system_admin_profile.labelname {{profileid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

