---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_nputcam_data"
description: |-
  Data fields of TCAM.
---

# fortimanager_object_system_npu_nputcam_data
Data fields of TCAM.

~> This resource is a sub resource for variable `data` of resource `fortimanager_object_system_npu_nputcam`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `npu_tcam` - Npu Tcam.

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuNpuTcamData can be imported using any of these accepted formats:
```
Set import_options = ["npu_tcam=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_nputcam_data.labelname ObjectSystemNpuNpuTcamData
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
