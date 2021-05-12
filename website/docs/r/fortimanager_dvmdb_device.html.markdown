---
subcategory: "Dvmdb"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_dvmdb_device"
description: |-
  Device table, most attributes are read-only and can only be changed internally. Refer to Device Manager Command module for API to add, delete, and manage devices.
---

# fortimanager_dvmdb_device
Device table, most attributes are read-only and can only be changed internally. Refer to Device Manager Command module for API to add, delete, and manage devices.

## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `adm_pass` - Adm_Pass.
* `adm_usr` - Adm_Usr.
* `app_ver` - App_Ver.
* `av_ver` - Av_Ver.
* `beta` - Beta.
* `branch_pt` - Branch_Pt.
* `build` - Build.
* `checksum` - Checksum.
* `conf_status` - Conf_Status. Valid values: `unknown`, `insync`, `outofsync`.

* `conn_mode` - Conn_Mode. Valid values: `active`, `passive`.

* `conn_status` - Conn_Status. Valid values: `UNKNOWN`, `up`, `down`.

* `db_status` - Db_Status. Valid values: `unknown`, `nomod`, `mod`.

* `desc` - Desc.
* `dev_status` - Dev_Status. Valid values: `none`, `unknown`, `checkedin`, `inprogress`, `installed`, `aborted`, `sched`, `retry`, `canceled`, `pending`, `retrieved`, `changed_conf`, `sync_fail`, `timeout`, `rev_revert`, `auto_updated`.

* `fap_cnt` - Fap_Cnt.
* `fazfull_act` - Faz.Full_Act.
* `fazperm` - Faz.Perm.
* `fazquota` - Faz.Quota.
* `fazused` - Faz.Used.
* `fex_cnt` - Fex_Cnt.
* `flags` - Flags. Valid values: `has_hdd`, `vdom_enabled`, `discover`, `reload`, `interim_build`, `offline_mode`, `is_model`, `fips_mode`, `linked_to_model`, `ip-conflict`, `faz-autosync`.

* `foslic_cpu` - VM Meter vCPU count.
* `foslic_dr_site` - VM Meter DR Site status. Valid values: `disable`, `enable`.

* `foslic_inst_time` - VM Meter first deployment time (in UNIX timestamp).
* `foslic_last_sync` - VM Meter last synchronized time (in UNIX timestamp).
* `foslic_ram` - VM Meter device RAM size (in MB).
* `foslic_type` - VM Meter license type. Valid values: `temporary`, `trial`, `regular`, `trial_expired`.

* `foslic_utm` - VM Meter services fw - Firewall av - Anti-virus ips - IPS app - App control url - Web filter utm - Full UTM fwb - FortiWeb Valid values: `fw`, `av`, `ips`, `app`, `url`, `utm`, `fwb`.

* `fsw_cnt` - Fsw_Cnt.
* `ha_group_id` - Ha_Group_Id.
* `ha_group_name` - Ha_Group_Name.
* `ha_mode` - Ha_Mode. Valid values: `standalone`, `AP`, `AA`, `ELBC`, `DUAL`, `fmg-enabled`, `autoscale`, `unknown`.

* `hdisk_size` - Hdisk_Size.
* `hostname` - Hostname.
* `hw_rev_major` - Hw_Rev_Major.
* `hw_rev_minor` - Hw_Rev_Minor.
* `hyperscale` - Hyperscale.
* `ip` - Ip.
* `ips_ext` - Ips_Ext.
* `ips_ver` - Ips_Ver.
* `last_checked` - Last_Checked.
* `last_resync` - Last_Resync.
* `latitude` - Latitude.
* `lic_flags` - Lic_Flags.
* `lic_region` - Lic_Region.
* `location_from` - Location_From.
* `logdisk_size` - Logdisk_Size.
* `longitude` - Longitude.
* `maxvdom` - Maxvdom.
* `metafields` - Default metafields: "City", "Province/State", "Country", "Company/Organization", "Contact".
* `mgmt_id` - Mgmt_Id.
* `mgmt_if` - Mgmt_If.
* `mgmt_mode` - Mgmt_Mode. Valid values: `unreg`, `fmg`, `faz`, `fmgfaz`.

* `mgt_vdom` - Mgt_Vdom.
* `module_sn` - Module_Sn.
* `mr` - Mr.
* `name` - Unique name for the device.
* `nsxt_service_name` - Nsxt_Service_Name.
* `os_type` - Os_Type. Valid values: `unknown`, `fos`, `fsw`, `foc`, `fml`, `faz`, `fwb`, `fch`, `fct`, `log`, `fmg`, `fsa`, `fdd`, `fac`, `fpx`, `fna`, `ffw`, `fsr`, `fad`, `fdc`.

* `os_ver` - Os_Ver. Valid values: `unknown`, `0.0`, `1.0`, `2.0`, `3.0`, `4.0`, `5.0`, `6.0`, `7.0`, `8.0`.

* `patch` - Patch.
* `platform_str` - Platform_Str.
* `prefer_img_ver` - Prefer_Img_Ver.
* `prio` - Prio.
* `private_key` - Private_Key.
* `private_key_status` - Private_Key_Status.
* `psk` - Psk.
* `role` - Role. Valid values: `master`, `ha-slave`, `autoscale-slave`.

* `sn` - Unique value for each device.
* `vdom` - Vdom. The structure of `vdom` block is documented below.
* `version` - Version.
* `vm_cpu` - Vm_Cpu.
* `vm_cpu_limit` - Vm_Cpu_Limit.
* `vm_lic_expire` - Vm_Lic_Expire.
* `vm_mem` - Vm_Mem.
* `vm_mem_limit` - Vm_Mem_Limit.
* `vm_status` - Vm_Status.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `vdom` block supports:

* `comments` - Comments.
* `metafields` - Meta Fields.
* `name` - Name.
* `opmode` - Opmode. Valid values: `nat`, `transparent`.

* `rtm_prof_id` - Rtm_Prof_Id.
* `status` - Status.
* `vpn_id` - Vpn_Id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dvmdb Device can be imported using any of these accepted formats:
```
$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_dvmdb_device.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom here will directly inherit the scopetype and adom configuration of the provider.
