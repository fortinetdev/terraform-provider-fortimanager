---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_nputcam_mask"
description: |-
  Mask fields of TCAM.
---

# fortimanager_object_system_npu_nputcam_mask
Mask fields of TCAM.

~> This resource is a sub resource for variable `mask` of resource `fortimanager_object_system_npu_nputcam`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `npu_tcam` - Npu Tcam.

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuNpuTcamMask can be imported using any of these accepted formats:
```
Set import_options = ["npu_tcam=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_nputcam_mask.labelname ObjectSystemNpuNpuTcamMask
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
