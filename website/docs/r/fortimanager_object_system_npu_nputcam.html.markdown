---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_nputcam"
description: |-
  Configure NPU TCAM policies.
---

# fortimanager_object_system_npu_nputcam
Configure NPU TCAM policies.

~> This resource is a sub resource for variable `npu_tcam` of resource `fortimanager_object_system_npu`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `data`: `fortimanager_object_system_npu_nputcam_data`
>- `mask`: `fortimanager_object_system_npu_nputcam_mask`
>- `mir_act`: `fortimanager_object_system_npu_nputcam_miract`
>- `pri_act`: `fortimanager_object_system_npu_nputcam_priact`
>- `sact`: `fortimanager_object_system_npu_nputcam_sact`
>- `tact`: `fortimanager_object_system_npu_nputcam_tact`



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.

* `data` - Data. The structure of `data` block is documented below.
* `dbg_dump` - Debug driver dump data/mask pdq.
* `mask` - Mask. The structure of `mask` block is documented below.
* `mir_act` - Mir-Act. The structure of `mir_act` block is documented below.
* `name` - NPU TCAM policies name.
* `oid` - NPU TCAM OID.
* `pri_act` - Pri-Act. The structure of `pri_act` block is documented below.
* `sact` - Sact. The structure of `sact` block is documented below.
* `tact` - Tact. The structure of `tact` block is documented below.
* `type` - TCAM policy type. Valid values: `L2_src_tc`, `L2_tgt_tc`, `L2_src_mir`, `L2_tgt_mir`, `L2_src_act`, `L2_tgt_act`, `IPv4_src_tc`, `IPv4_tgt_tc`, `IPv4_src_mir`, `IPv4_tgt_mir`, `IPv4_src_act`, `IPv4_tgt_act`, `IPv6_src_tc`, `IPv6_tgt_tc`, `IPv6_src_mir`, `IPv6_tgt_mir`, `IPv6_src_act`, `IPv6_tgt_act`.

* `vid` - NPU TCAM VID.

The `data` block supports:

* `df` - tcam data ip flag df. Valid values: `disable`, `enable`.

* `dstip` - tcam data dst ipv4 address.
* `dstipv6` - tcam data dst ipv6 address.
* `dstmac` - tcam data dst macaddr.
* `dstport` - tcam data L4 dst port.
* `ethertype` - tcam data ethertype.
* `ext_tag` - tcam data extension tag. Valid values: `disable`, `enable`.

* `frag_off` - tcam data ip flag fragment offset.
* `gen_buf_cnt` - tcam data gen info buffer count.
* `gen_iv` - tcam data gen info iv. Valid values: `invalid`, `valid`.

* `gen_l3_flags` - tcam data gen info L3 flags.
* `gen_l4_flags` - tcam data gen info L4 flags.
* `gen_pkt_ctrl` - tcam data gen info packet control.
* `gen_pri` - tcam data gen info priority.
* `gen_pri_v` - tcam data gen info priority valid. Valid values: `invalid`, `valid`.

* `gen_tv` - tcam data gen info tv. Valid values: `invalid`, `valid`.

* `ihl` - tcam data ipv4 IHL.
* `ip4_id` - tcam data ipv4 id.
* `ip6_fl` - tcam data ipv6 flow label.
* `ipver` - tcam data ip header version.
* `l4_wd10` - tcam data L4 word10.
* `l4_wd11` - tcam data L4 word11.
* `l4_wd8` - tcam data L4 word8.
* `l4_wd9` - tcam data L4 word9.
* `mf` - tcam data ip flag mf. Valid values: `disable`, `enable`.

* `protocol` - tcam data ip protocol.
* `slink` - tcam data sublink.
* `smac_change` - tcam data source MAC change. Valid values: `disable`, `enable`.

* `sp` - tcam data source port.
* `src_cfi` - tcam data source cfi. Valid values: `disable`, `enable`.

* `src_prio` - tcam data source priority.
* `src_updt` - tcam data source update. Valid values: `disable`, `enable`.

* `srcip` - tcam data src ipv4 address.
* `srcipv6` - tcam data src ipv6 address.
* `srcmac` - tcam data src macaddr.
* `srcport` - tcam data L4 src port.
* `svid` - tcam data source vid.
* `tcp_ack` - tcam data tcp flag ack. Valid values: `disable`, `enable`.

* `tcp_cwr` - tcam data tcp flag cwr. Valid values: `disable`, `enable`.

* `tcp_ece` - tcam data tcp flag ece. Valid values: `disable`, `enable`.

* `tcp_fin` - tcam data tcp flag fin. Valid values: `disable`, `enable`.

* `tcp_push` - tcam data tcp flag push. Valid values: `disable`, `enable`.

* `tcp_rst` - tcam data tcp flag rst. Valid values: `disable`, `enable`.

* `tcp_syn` - tcam data tcp flag syn. Valid values: `disable`, `enable`.

* `tcp_urg` - tcam data tcp flag urg. Valid values: `disable`, `enable`.

* `tgt_cfi` - tcam data target cfi. Valid values: `disable`, `enable`.

* `tgt_prio` - tcam data target priority.
* `tgt_updt` - tcam data target port update. Valid values: `disable`, `enable`.

* `tgt_v` - tcam data target valid. Valid values: `invalid`, `valid`.

* `tos` - tcam data ip tos.
* `tp` - tcam data target port.
* `ttl` - tcam data ip ttl.
* `tvid` - tcam data target vid.
* `vdid` - tcam data vdom id.

The `mask` block supports:

* `df` - tcam mask ip flag df. Valid values: `disable`, `enable`.

* `dstip` - tcam mask dst ipv4 address.
* `dstipv6` - tcam mask dst ipv6 address.
* `dstmac` - tcam mask dst macaddr.
* `dstport` - tcam mask L4 dst port.
* `ethertype` - tcam mask ethertype.
* `ext_tag` - tcam mask extension tag. Valid values: `disable`, `enable`.

* `frag_off` - tcam data ip flag fragment offset.
* `gen_buf_cnt` - tcam mask gen info buffer count.
* `gen_iv` - tcam mask gen info iv. Valid values: `invalid`, `valid`.

* `gen_l3_flags` - tcam mask gen info L3 flags.
* `gen_l4_flags` - tcam mask gen info L4 flags.
* `gen_pkt_ctrl` - tcam mask gen info packet control.
* `gen_pri` - tcam mask gen info priority.
* `gen_pri_v` - tcam mask gen info priority valid. Valid values: `invalid`, `valid`.

* `gen_tv` - tcam mask gen info tv. Valid values: `invalid`, `valid`.

* `ihl` - tcam mask ipv4 IHL.
* `ip4_id` - tcam mask ipv4 id.
* `ip6_fl` - tcam mask ipv6 flow label.
* `ipver` - tcam mask ip header version.
* `l4_wd10` - tcam mask L4 word10.
* `l4_wd11` - tcam mask L4 word11.
* `l4_wd8` - tcam mask L4 word8.
* `l4_wd9` - tcam mask L4 word9.
* `mf` - tcam mask ip flag mf. Valid values: `disable`, `enable`.

* `protocol` - tcam mask ip protocol.
* `slink` - tcam mask sublink.
* `smac_change` - tcam mask source MAC change. Valid values: `disable`, `enable`.

* `sp` - tcam mask source port.
* `src_cfi` - tcam mask source cfi. Valid values: `disable`, `enable`.

* `src_prio` - tcam mask source priority.
* `src_updt` - tcam mask source update. Valid values: `disable`, `enable`.

* `srcip` - tcam mask src ipv4 address.
* `srcipv6` - tcam mask src ipv6 address.
* `srcmac` - tcam mask src macaddr.
* `srcport` - tcam mask L4 src port.
* `svid` - tcam mask source vid.
* `tcp_ack` - tcam mask tcp flag ack. Valid values: `disable`, `enable`.

* `tcp_cwr` - tcam mask tcp flag cwr. Valid values: `disable`, `enable`.

* `tcp_ece` - tcam mask tcp flag ece. Valid values: `disable`, `enable`.

* `tcp_fin` - tcam mask tcp flag fin. Valid values: `disable`, `enable`.

* `tcp_push` - tcam mask tcp flag push. Valid values: `disable`, `enable`.

* `tcp_rst` - tcam mask tcp flag rst. Valid values: `disable`, `enable`.

* `tcp_syn` - tcam mask tcp flag syn. Valid values: `disable`, `enable`.

* `tcp_urg` - tcam mask tcp flag urg. Valid values: `disable`, `enable`.

* `tgt_cfi` - tcam mask target cfi. Valid values: `disable`, `enable`.

* `tgt_prio` - tcam mask target priority.
* `tgt_updt` - tcam mask target port update. Valid values: `disable`, `enable`.

* `tgt_v` - tcam mask target valid. Valid values: `invalid`, `valid`.

* `tos` - tcam mask ip tos.
* `tp` - tcam mask target port.
* `ttl` - tcam mask ip ttl.
* `tvid` - tcam mask target vid.
* `vdid` - tcam mask vdom id.

The `mir_act` block supports:

* `vlif` - tcam mirror action vlif.

The `pri_act` block supports:

* `priority` - tcam priority action priority.
* `weight` - tcam priority action weight.

The `sact` block supports:

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


The `tact` block supports:

* `act` - tcam tact act.
* `act_v` - Enable to set tact act. Valid values: `disable`, `enable`.

* `fmtuv4_s` - tcam tact fmtuv4-s.
* `fmtuv4_s_v` - Enable to set tact fmtuv4-s. Valid values: `disable`, `enable`.

* `fmtuv6_s` - tcam tact fmtuv6-s.
* `fmtuv6_s_v` - Enable to set tact fmtuv6-s. Valid values: `disable`, `enable`.

* `lnkid` - tcam tact lnkid.
* `lnkid_v` - Enable to set tact lnkid. Valid values: `disable`, `enable`.

* `mac_id` - tcam tact mac-id.
* `mac_id_v` - Enable to set tact mac-id. Valid values: `disable`, `enable`.

* `mss_t` - tcam tact mss.
* `mss_t_v` - Enable to set tact mss. Valid values: `disable`, `enable`.

* `mtuv4` - tcam tact mtuv4.
* `mtuv4_v` - Enable to set tact mtuv4. Valid values: `disable`, `enable`.

* `mtuv6` - tcam tact mtuv6.
* `mtuv6_v` - Enable to set tact mtuv6. Valid values: `disable`, `enable`.

* `slif_act` - tcam tact slif-act.
* `slif_act_v` - Enable to set tact slif-act. Valid values: `disable`, `enable`.

* `sublnkid` - tcam tact sublnkid.
* `sublnkid_v` - Enable to set tact sublnkid. Valid values: `disable`, `enable`.

* `tgtv_act` - tcam tact tgtv-act.
* `tgtv_act_v` - Enable to set tact tgtv-act. Valid values: `disable`, `enable`.

* `tlif_act` - tcam tact tlif-act.
* `tlif_act_v` - Enable to set tact tlif-act. Valid values: `disable`, `enable`.

* `tpeid` - tcam tact tpeid.
* `tpeid_v` - Enable to set tact tpeid. Valid values: `disable`, `enable`.

* `v6fe` - tcam tact v6fe.
* `v6fe_v` - Enable to set tact v6fe. Valid values: `disable`, `enable`.

* `vep_en_v` - Enable to set tact vep-en. Valid values: `disable`, `enable`.

* `vep_slid` - tcam tact vep_slid.
* `vep_slid_v` - Enable to set tact vep-slid. Valid values: `disable`, `enable`.

* `vep_en` - tcam tact vep_en.
* `xlt_lif` - tcam tact xlt-lif.
* `xlt_lif_v` - Enable to set tact xlt-lif. Valid values: `disable`, `enable`.

* `xlt_vid` - tcam tact xlt-vid.
* `xlt_vid_v` - Enable to set tact xlt-vid. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ObjectSystem NpuNpuTcam can be imported using any of these accepted formats:
```

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_nputcam.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
