---
subcategory: "Device Manager"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_device"
description: |-
  Manage a device to the Device Manager database.
---

# fortimanager_device
Manage a device to the Device Manager database.

## Example Usage

```hcl
resource "fortimanager_device" "trname" {
  adom = "root"
  ip        = "<FGT_IP>"
  mgmt_mode = "fmgfaz"
  name      = "terraform-test"
  adm_usr   = "admin"
  adm_pass  = "Fortinet12345##"
  sn = "FGVMSLTM000000000"
}
```

## Argument Reference

* `adm_pass` (String, Sensitive) <i>add real and promote device</i>. Device admin password.
* `adm_usr` (String) <i>add real and promote device</i>. Device admin username.
* `authorizationtemplate` (String) <i>add model device only</i>. Fabric Authorization Template to auto genreate for the new model device upon creation.
* `desc` (String) Description. 
* `deviceaction` (String) Specify add device operations, or leave blank to add real device:<ul><li>"add_model" - add a model device.<li>"promote_unreg" - promote an unregistered device to be managed by FortiManager using information from database.</ul>
* `deviceblueprint` (String) <i>add model device only</i>. Device blueprint to apply to the new model device.
* `fazquota` (Number) <i>available for all operations</i>.
* `adom` (String) Name or ID of the ADOM where the command is to be executed on.
* `groups` (Block List) Groups. (see [below for nested schema](#nestedblock--groups))
* `ip` (String) <i>add real device only</i>. Add device will probe with this IP using the log in credential specified.
* `metafields` (Map of String) <i>add real and model device.</i>.
* `mgmt_mode` (String) <i>add real and model device</i>. Valid values: `unreg`, `fmg`, `faz`, `fmgfaz`.
* `mr` (Number) <i>add model device only</i>.
* `name` (String) <i>required for all operations</i>. Unique name for the device.
* `os_type` (String) <i>add model device only</i>. Valid values: `unknown`, `fos`, `fsw`, `foc`, `fml`, `faz`, `fwb`, `fch`, `fct`, `log`, `fmg`, `fsa`, `fdd`, `fac`, `fpx`, `fna`.
* `os_ver` (String) <i>add model device only</i>. Valid values: `unknown`, `0.0`, `1.0`, `2.0`, `3.0`, `4.0`, `5.0`, `6.0`, `7.0`, `8.0`.
* `patch` (Number) <i>add model device only</i>.
* `platform_str` (String) <i>add model device only</i>. Required for determine the platform for VM platforms.
* `sn` (String) <i>add model device only</i>. This attribute will be used to determine the device platform, except for VM platforms, where <i>platform_str</i> is also required.

<a id="nestedblock--groups"></a>
### Nested Schema for `groups`

Optional:

* `name` (String) Name.
* `vdom` (String) Vdom.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` (String) an identifier for the resource with format {{name}}.

* `app_ver` (String) App_Ver.
* `av_ver` (String) Av_Ver.
* `beta` (Number) Beta.
* `branch_pt` (Number) Branch_Pt.
* `build` (Number) Build.
* `checksum` (String) Checksum.
* `cluster_worker` (Set of String) Cluster_Worker.
* `conf_status` (String) Conf_Status.
* `conn_mode` (String) Conn_Mode.
* `conn_status` (String) Conn_Status.
* `db_status` (String) Db_Status.
* `dev_status` (String) Dev_Status.
* `eip` (String) Eip.
* `fap_cnt` (Number) Fap_Cnt.
* `fazfull_act` (Number) Faz.Full_Act.
* `fazperm` (Number) Faz.Perm.
* `fazused` (Number) Faz.Used.
* `fex_cnt` (Number) Fex_Cnt.
* `first_tunnel_up` (Number) First_Tunnel_Up.
* `flags` (Set of String) Flags.
* `foslic_cpu` (Number) VM Meter vCPU count.
* `foslic_dr_site` (String) VM Meter DR Site status.
* `foslic_inst_time` (Number) VM Meter first deployment time (in UNIX timestamp).
* `foslic_last_sync` (Number) VM Meter last synchronized time (in UNIX timestamp).
* `foslic_ram` (Number) VM Meter device RAM size (in MB).
* `foslic_type` (String) VM Meter license type. 
* `foslic_utm` (String) VM Meter services fw - Firewall av - Anti-virus ips - IPS app - App control url - Web filter utm - Full UTM fwb - FortiWeb
* `fsw_cnt` (Number) Fsw_Cnt.
* `ha_group_id` (Number) Ha_Group_Id.
* `ha_group_name` (String) Ha_Group_Name.
* `ha_mode` (String) Ha_Mode.
* `ha_upgrade_mode` (Number) Ha_Upgrade_Mode.
* `havsn` (String) Ha.Vsn.
* `hdisk_size` (Number) Hdisk_Size.
* `hostname` (String) Hostname.
* `hw_generation` (Number) Hw_Generation.
* `hw_rev_major` (Number) Hw_Rev_Major.
* `hw_rev_minor` (Number) Hw_Rev_Minor.
* `hyperscale` (Number) Hyperscale.
* `ips_ext` (Number) Ips_Ext.
* `ips_ver` (String) Ips_Ver.
* `last_checked` (Number) Last_Checked.
* `last_resync` (Number) Last_Resync.
* `latitude` (String) Latitude.
* `lic_flags` (Number) Lic_Flags.
* `lic_region` (String) Lic_Region.
* `location_from` (String) Location_From.
* `logdisk_size` (Number) Logdisk_Size.
* `longitude` (String) Longitude.
* `maxvdom` (Number) Maxvdom.
* `mgmt_if` (String) Mgmt_If.
* `mgmt_uuid` (String) Mgmt_Uuid.
* `mgt_vdom` (String) Mgt_Vdom.
* `module_sn` (String) Module_Sn.
* `nsxt_service_name` (String) Nsxt_Service_Name.
* `prefer_img_ver` (String) Prefer_Img_Ver.
* `prio` (Number) Prio.
* `private_key` (String) Private_Key.
* `private_key_status` (Number) Private_Key_Status.
* `psk` (String) Psk.
* `relver_info` (String) Relver_Info.
* `role` (String) Role. 
* `sov_sase_license` (String) Sov_Sase_License.
* `tunnel_sn` (String) Tunnel_Sn.
* `vdom` (Block List) Vdom. (see [below for nested schema](#nestedblock--vdom))
* `version` (Number) Version.
* `vm_cpu` (Number) Vm_Cpu.
* `vm_cpu_limit` (Number) Vm_Cpu_Limit.
* `vm_lic_expire` (Number) Vm_Lic_Expire.
* `vm_lic_overdue_since` (Number) Vm_Lic_Overdue_Since.
* `vm_mem` (Number) Vm_Mem.
* `vm_mem_limit` (Number) Vm_Mem_Limit.
* `vm_payg_status` (Number) Vm_Payg_Status.
* `vm_status` (String) Vm_Status.

<a id="nestedblock--vdom"></a>
### Nested Schema for `vdom`

Optional:

* `comments` (String) Comments.
* `metafields` (Map of String) Meta Fields.
* `name` (String) Name.
* `opmode` (String) Opmode.
* `rtm_prof_id` (Number) Rtm_Prof_Id.
* `status` (String) Status.
* `vdom_type` (String) Vdom_Type.
* `vpn_id` (Number) Vpn_Id.

## Import

Device can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_device.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The adom is not required for import operation. But delete operation need it. Please make sure the adom is configured in the resource or in the provider configuration.
