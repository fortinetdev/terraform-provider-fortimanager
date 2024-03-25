---
subcategory: "No Category"
layout: "fortimanager"
page_title: "FortiManager: fortimanager_object_system_npu_nputcam_tact"
description: |-
  Target action of TCAM.
---

# fortimanager_object_system_npu_nputcam_tact
Target action of TCAM.

~> This resource is a sub resource for variable `tact` of resource `fortimanager_object_system_npu_nputcam`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `scopetype` - The scope of application of the resource. Valid values: `inherit`, `adom`, `global`. The `inherit` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is `inherit`.
* `adom` - Adom. This value is valid only when the `scopetype` is `adom`, otherwise the value of adom in the provider will be inherited.
* `npu_tcam` - Npu Tcam.

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
* `id` - an identifier for the resource.

## Import

ObjectSystem NpuNpuTcamTact can be imported using any of these accepted formats:
```
Set import_options = ["npu_tcam=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fortimanager_object_system_npu_nputcam_tact.labelname ObjectSystemNpuNpuTcamTact
$ unset "FORTIMANAGER_IMPORT_TABLE"
```
-> **Hint:** The scopetype and adom for import will directly inherit the scopetype and adom configuration of the provider.
