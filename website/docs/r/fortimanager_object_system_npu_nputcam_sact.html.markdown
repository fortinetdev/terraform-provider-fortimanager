---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_nputcam_sact"
description: |-
  Source action of TCAM.
---

# fortimanager_object_system_npu_nputcam_sact
Source action of TCAM.

~> This resource is a sub resource for variable `sact` of resource `fortimanager_object_system_npu_nputcam`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `npu_tcam` - Npu Tcam.

* `act` - tcam sact act.
* `act_v` - Enable to set sact act. Valid values: `disable`, `enable`.

* `bmproc` - tcam sact bmproc.
* `bmproc_v` - Enable to set sact bmproc. Valid values: `disable`, `enable`.

* `df_lif` - tcam sact df-lif.
* `df_lif_v` - Enable to set sact df-lif. Valid values: `disable`, `enable`.

* `dfr` - tcam sact dfr.
* `dfr_v` - Enable to set sact dfr. Valid values: `disable`, `enable`.

* `dmac_skip` - tcam sact dmac-skip.
* `dmac_skip_v` - Enable to set sact dmac-skip. Valid values: `disable`, `enable`.

* `dosen` - tcam sact dosen.
* `dosen_v` - Enable to set sact dosen. Valid values: `disable`, `enable`.

* `espff_proc` - tcam sact espff-proc.
* `espff_proc_v` - Enable to set sact espff-proc. Valid values: `disable`, `enable`.

* `etype_pid` - tcam sact etype-pid.
* `etype_pid_v` - Enable to set sact etype-pid. Valid values: `disable`, `enable`.

* `frag_proc` - tcam sact frag-proc.
* `frag_proc_v` - Enable to set sact frag-proc. Valid values: `disable`, `enable`.

* `fwd` - tcam sact fwd.
* `fwd_lif` - tcam sact fwd-lif.
* `fwd_lif_v` - Enable to set sact fwd-lif. Valid values: `disable`, `enable`.

* `fwd_tvid` - tcam sact fwd-tvid.
* `fwd_tvid_v` - Enable to set sact fwd-vid. Valid values: `disable`, `enable`.

* `fwd_v` - Enable to set sact fwd. Valid values: `disable`, `enable`.

* `icpen` - tcam sact icpen.
* `icpen_v` - Enable to set sact icpen. Valid values: `disable`, `enable`.

* `igmp_mld_snp` - tcam sact igmp-mld-snp.
* `igmp_mld_snp_v` - Enable to set sact igmp-mld-snp. Valid values: `disable`, `enable`.

* `learn` - tcam sact learn.
* `learn_v` - Enable to set sact learn. Valid values: `disable`, `enable`.

* `m_srh_ctrl` - tcam sact m-srh-ctrl.
* `m_srh_ctrl_v` - Enable to set sact m-srh-ctrl. Valid values: `disable`, `enable`.

* `mac_id` - tcam sact mac-id.
* `mac_id_v` - Enable to set sact mac-id. Valid values: `disable`, `enable`.

* `mss` - tcam sact mss.
* `mss_v` - Enable to set sact mss. Valid values: `disable`, `enable`.

* `pleen` - tcam sact pleen.
* `pleen_v` - Enable to set sact pleen. Valid values: `disable`, `enable`.

* `prio_pid` - tcam sact prio-pid.
* `prio_pid_v` - Enable to set sact prio-pid. Valid values: `disable`, `enable`.

* `promis` - tcam sact promis.
* `promis_v` - Enable to set sact promis. Valid values: `disable`, `enable`.

* `rfsh` - tcam sact rfsh.
* `rfsh_v` - Enable to set sact rfsh. Valid values: `disable`, `enable`.

* `smac_skip` - tcam sact smac-skip.
* `smac_skip_v` - Enable to set sact smac-skip. Valid values: `disable`, `enable`.

* `tp_smchk_v` - Enable to set sact tp mode. Valid values: `disable`, `enable`.

* `tp_smchk` - tcam sact tp mode.
* `tpe_id` - tcam sact tpe-id.
* `tpe_id_v` - Enable to set sact tpe-id. Valid values: `disable`, `enable`.

* `vdm` - tcam sact vdm.
* `vdm_v` - Enable to set sact vdm. Valid values: `disable`, `enable`.

* `vdom_id` - tcam sact vdom-id.
* `vdom_id_v` - Enable to set sact vdom-id. Valid values: `disable`, `enable`.

* `x_mode` - tcam sact x-mode.
* `x_mode_v` - Enable to set sact x-mode. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuNpuTcamSact can be imported using any of these accepted formats:
```
Set import_options = ["npu_tcam=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_nputcam_sact.labelname ObjectSystemNpuNpuTcamSact
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
