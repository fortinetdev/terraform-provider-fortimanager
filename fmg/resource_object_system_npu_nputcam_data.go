// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Data fields of TCAM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcamData() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamDataUpdate,
		Read:   resourceObjectSystemNpuNpuTcamDataRead,
		Update: resourceObjectSystemNpuNpuTcamDataUpdate,
		Delete: resourceObjectSystemNpuNpuTcamDataDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"npu_tcam": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"df": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ethertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ext_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"frag_off": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gen_buf_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gen_iv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gen_l3_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gen_l4_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gen_pkt_ctrl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gen_pri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gen_pri_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gen_tv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ihl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip4_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_fl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"l4_wd10": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"l4_wd11": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"l4_wd8": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"l4_wd9": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"slink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"smac_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"src_cfi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_prio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"src_updt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"svid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_ack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_cwr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_ece": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_fin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_rst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_urg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tgt_cfi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tgt_prio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tgt_updt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tgt_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tvid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vdid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpuTcamDataUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	npu_tcam := d.Get("npu_tcam").(string)
	paradict["npu_tcam"] = npu_tcam

	obj, err := getObjectObjectSystemNpuNpuTcamData(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamData resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpuTcamData(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamData resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpuTcamData")

	return resourceObjectSystemNpuNpuTcamDataRead(d, m)
}

func resourceObjectSystemNpuNpuTcamDataDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	npu_tcam := d.Get("npu_tcam").(string)
	paradict["npu_tcam"] = npu_tcam

	err = c.DeleteObjectSystemNpuNpuTcamData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcamData resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamDataRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	npu_tcam := d.Get("npu_tcam").(string)
	if npu_tcam == "" {
		npu_tcam = importOptionChecking(m.(*FortiClient).Cfg, "npu_tcam")
		if npu_tcam == "" {
			return fmt.Errorf("Parameter npu_tcam is missing")
		}
		if err = d.Set("npu_tcam", npu_tcam); err != nil {
			return fmt.Errorf("Error set params npu_tcam: %v", err)
		}
	}
	paradict["npu_tcam"] = npu_tcam

	o, err := c.ReadObjectSystemNpuNpuTcamData(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamData resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcamData(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamData resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamDataDf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstipv63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstmac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataEthertype3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataExtTag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataFragOff3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenBufCnt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenIv3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenL3Flags3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenL4Flags3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPktCtrl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPri3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPriV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenTv3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIhl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIp4Id3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIp6Fl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIpver3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd103rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd113rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd83rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd93rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataMf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataProtocol3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSlink3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSmacChange3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcCfi3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcPrio3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcUpdt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcipv63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcmac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSvid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpAck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpCwr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpEce3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpFin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpPush3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpRst3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpSyn3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpUrg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtCfi3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtPrio3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtUpdt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTos3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTtl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTvid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataVdid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcamData(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("df", flattenObjectSystemNpuNpuTcamDataDf3rdl(o["df"], d, "df")); err != nil {
		if vv, ok := fortiAPIPatch(o["df"], "ObjectSystemNpuNpuTcamData-Df"); ok {
			if err = d.Set("df", vv); err != nil {
				return fmt.Errorf("Error reading df: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading df: %v", err)
		}
	}

	if err = d.Set("dstip", flattenObjectSystemNpuNpuTcamDataDstip3rdl(o["dstip"], d, "dstip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstip"], "ObjectSystemNpuNpuTcamData-Dstip"); ok {
			if err = d.Set("dstip", vv); err != nil {
				return fmt.Errorf("Error reading dstip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstip: %v", err)
		}
	}

	if err = d.Set("dstipv6", flattenObjectSystemNpuNpuTcamDataDstipv63rdl(o["dstipv6"], d, "dstipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstipv6"], "ObjectSystemNpuNpuTcamData-Dstipv6"); ok {
			if err = d.Set("dstipv6", vv); err != nil {
				return fmt.Errorf("Error reading dstipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstipv6: %v", err)
		}
	}

	if err = d.Set("dstmac", flattenObjectSystemNpuNpuTcamDataDstmac3rdl(o["dstmac"], d, "dstmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstmac"], "ObjectSystemNpuNpuTcamData-Dstmac"); ok {
			if err = d.Set("dstmac", vv); err != nil {
				return fmt.Errorf("Error reading dstmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstmac: %v", err)
		}
	}

	if err = d.Set("dstport", flattenObjectSystemNpuNpuTcamDataDstport3rdl(o["dstport"], d, "dstport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport"], "ObjectSystemNpuNpuTcamData-Dstport"); ok {
			if err = d.Set("dstport", vv); err != nil {
				return fmt.Errorf("Error reading dstport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("ethertype", flattenObjectSystemNpuNpuTcamDataEthertype3rdl(o["ethertype"], d, "ethertype")); err != nil {
		if vv, ok := fortiAPIPatch(o["ethertype"], "ObjectSystemNpuNpuTcamData-Ethertype"); ok {
			if err = d.Set("ethertype", vv); err != nil {
				return fmt.Errorf("Error reading ethertype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ethertype: %v", err)
		}
	}

	if err = d.Set("ext_tag", flattenObjectSystemNpuNpuTcamDataExtTag3rdl(o["ext-tag"], d, "ext_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-tag"], "ObjectSystemNpuNpuTcamData-ExtTag"); ok {
			if err = d.Set("ext_tag", vv); err != nil {
				return fmt.Errorf("Error reading ext_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_tag: %v", err)
		}
	}

	if err = d.Set("frag_off", flattenObjectSystemNpuNpuTcamDataFragOff3rdl(o["frag-off"], d, "frag_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["frag-off"], "ObjectSystemNpuNpuTcamData-FragOff"); ok {
			if err = d.Set("frag_off", vv); err != nil {
				return fmt.Errorf("Error reading frag_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frag_off: %v", err)
		}
	}

	if err = d.Set("gen_buf_cnt", flattenObjectSystemNpuNpuTcamDataGenBufCnt3rdl(o["gen-buf-cnt"], d, "gen_buf_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-buf-cnt"], "ObjectSystemNpuNpuTcamData-GenBufCnt"); ok {
			if err = d.Set("gen_buf_cnt", vv); err != nil {
				return fmt.Errorf("Error reading gen_buf_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_buf_cnt: %v", err)
		}
	}

	if err = d.Set("gen_iv", flattenObjectSystemNpuNpuTcamDataGenIv3rdl(o["gen-iv"], d, "gen_iv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-iv"], "ObjectSystemNpuNpuTcamData-GenIv"); ok {
			if err = d.Set("gen_iv", vv); err != nil {
				return fmt.Errorf("Error reading gen_iv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_iv: %v", err)
		}
	}

	if err = d.Set("gen_l3_flags", flattenObjectSystemNpuNpuTcamDataGenL3Flags3rdl(o["gen-l3-flags"], d, "gen_l3_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-l3-flags"], "ObjectSystemNpuNpuTcamData-GenL3Flags"); ok {
			if err = d.Set("gen_l3_flags", vv); err != nil {
				return fmt.Errorf("Error reading gen_l3_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_l3_flags: %v", err)
		}
	}

	if err = d.Set("gen_l4_flags", flattenObjectSystemNpuNpuTcamDataGenL4Flags3rdl(o["gen-l4-flags"], d, "gen_l4_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-l4-flags"], "ObjectSystemNpuNpuTcamData-GenL4Flags"); ok {
			if err = d.Set("gen_l4_flags", vv); err != nil {
				return fmt.Errorf("Error reading gen_l4_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_l4_flags: %v", err)
		}
	}

	if err = d.Set("gen_pkt_ctrl", flattenObjectSystemNpuNpuTcamDataGenPktCtrl3rdl(o["gen-pkt-ctrl"], d, "gen_pkt_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-pkt-ctrl"], "ObjectSystemNpuNpuTcamData-GenPktCtrl"); ok {
			if err = d.Set("gen_pkt_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading gen_pkt_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_pkt_ctrl: %v", err)
		}
	}

	if err = d.Set("gen_pri", flattenObjectSystemNpuNpuTcamDataGenPri3rdl(o["gen-pri"], d, "gen_pri")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-pri"], "ObjectSystemNpuNpuTcamData-GenPri"); ok {
			if err = d.Set("gen_pri", vv); err != nil {
				return fmt.Errorf("Error reading gen_pri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_pri: %v", err)
		}
	}

	if err = d.Set("gen_pri_v", flattenObjectSystemNpuNpuTcamDataGenPriV3rdl(o["gen-pri-v"], d, "gen_pri_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-pri-v"], "ObjectSystemNpuNpuTcamData-GenPriV"); ok {
			if err = d.Set("gen_pri_v", vv); err != nil {
				return fmt.Errorf("Error reading gen_pri_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_pri_v: %v", err)
		}
	}

	if err = d.Set("gen_tv", flattenObjectSystemNpuNpuTcamDataGenTv3rdl(o["gen-tv"], d, "gen_tv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-tv"], "ObjectSystemNpuNpuTcamData-GenTv"); ok {
			if err = d.Set("gen_tv", vv); err != nil {
				return fmt.Errorf("Error reading gen_tv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_tv: %v", err)
		}
	}

	if err = d.Set("ihl", flattenObjectSystemNpuNpuTcamDataIhl3rdl(o["ihl"], d, "ihl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ihl"], "ObjectSystemNpuNpuTcamData-Ihl"); ok {
			if err = d.Set("ihl", vv); err != nil {
				return fmt.Errorf("Error reading ihl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ihl: %v", err)
		}
	}

	if err = d.Set("ip4_id", flattenObjectSystemNpuNpuTcamDataIp4Id3rdl(o["ip4-id"], d, "ip4_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-id"], "ObjectSystemNpuNpuTcamData-Ip4Id"); ok {
			if err = d.Set("ip4_id", vv); err != nil {
				return fmt.Errorf("Error reading ip4_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_id: %v", err)
		}
	}

	if err = d.Set("ip6_fl", flattenObjectSystemNpuNpuTcamDataIp6Fl3rdl(o["ip6-fl"], d, "ip6_fl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-fl"], "ObjectSystemNpuNpuTcamData-Ip6Fl"); ok {
			if err = d.Set("ip6_fl", vv); err != nil {
				return fmt.Errorf("Error reading ip6_fl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_fl: %v", err)
		}
	}

	if err = d.Set("ipver", flattenObjectSystemNpuNpuTcamDataIpver3rdl(o["ipver"], d, "ipver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipver"], "ObjectSystemNpuNpuTcamData-Ipver"); ok {
			if err = d.Set("ipver", vv); err != nil {
				return fmt.Errorf("Error reading ipver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipver: %v", err)
		}
	}

	if err = d.Set("l4_wd10", flattenObjectSystemNpuNpuTcamDataL4Wd103rdl(o["l4-wd10"], d, "l4_wd10")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd10"], "ObjectSystemNpuNpuTcamData-L4Wd10"); ok {
			if err = d.Set("l4_wd10", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd10: %v", err)
		}
	}

	if err = d.Set("l4_wd11", flattenObjectSystemNpuNpuTcamDataL4Wd113rdl(o["l4-wd11"], d, "l4_wd11")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd11"], "ObjectSystemNpuNpuTcamData-L4Wd11"); ok {
			if err = d.Set("l4_wd11", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd11: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd11: %v", err)
		}
	}

	if err = d.Set("l4_wd8", flattenObjectSystemNpuNpuTcamDataL4Wd83rdl(o["l4-wd8"], d, "l4_wd8")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd8"], "ObjectSystemNpuNpuTcamData-L4Wd8"); ok {
			if err = d.Set("l4_wd8", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd8: %v", err)
		}
	}

	if err = d.Set("l4_wd9", flattenObjectSystemNpuNpuTcamDataL4Wd93rdl(o["l4-wd9"], d, "l4_wd9")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd9"], "ObjectSystemNpuNpuTcamData-L4Wd9"); ok {
			if err = d.Set("l4_wd9", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd9: %v", err)
		}
	}

	if err = d.Set("mf", flattenObjectSystemNpuNpuTcamDataMf3rdl(o["mf"], d, "mf")); err != nil {
		if vv, ok := fortiAPIPatch(o["mf"], "ObjectSystemNpuNpuTcamData-Mf"); ok {
			if err = d.Set("mf", vv); err != nil {
				return fmt.Errorf("Error reading mf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mf: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectSystemNpuNpuTcamDataProtocol3rdl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectSystemNpuNpuTcamData-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("slink", flattenObjectSystemNpuNpuTcamDataSlink3rdl(o["slink"], d, "slink")); err != nil {
		if vv, ok := fortiAPIPatch(o["slink"], "ObjectSystemNpuNpuTcamData-Slink"); ok {
			if err = d.Set("slink", vv); err != nil {
				return fmt.Errorf("Error reading slink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slink: %v", err)
		}
	}

	if err = d.Set("smac_change", flattenObjectSystemNpuNpuTcamDataSmacChange3rdl(o["smac-change"], d, "smac_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["smac-change"], "ObjectSystemNpuNpuTcamData-SmacChange"); ok {
			if err = d.Set("smac_change", vv); err != nil {
				return fmt.Errorf("Error reading smac_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smac_change: %v", err)
		}
	}

	if err = d.Set("sp", flattenObjectSystemNpuNpuTcamDataSp3rdl(o["sp"], d, "sp")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp"], "ObjectSystemNpuNpuTcamData-Sp"); ok {
			if err = d.Set("sp", vv); err != nil {
				return fmt.Errorf("Error reading sp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp: %v", err)
		}
	}

	if err = d.Set("src_cfi", flattenObjectSystemNpuNpuTcamDataSrcCfi3rdl(o["src-cfi"], d, "src_cfi")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-cfi"], "ObjectSystemNpuNpuTcamData-SrcCfi"); ok {
			if err = d.Set("src_cfi", vv); err != nil {
				return fmt.Errorf("Error reading src_cfi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_cfi: %v", err)
		}
	}

	if err = d.Set("src_prio", flattenObjectSystemNpuNpuTcamDataSrcPrio3rdl(o["src-prio"], d, "src_prio")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-prio"], "ObjectSystemNpuNpuTcamData-SrcPrio"); ok {
			if err = d.Set("src_prio", vv); err != nil {
				return fmt.Errorf("Error reading src_prio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_prio: %v", err)
		}
	}

	if err = d.Set("src_updt", flattenObjectSystemNpuNpuTcamDataSrcUpdt3rdl(o["src-updt"], d, "src_updt")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-updt"], "ObjectSystemNpuNpuTcamData-SrcUpdt"); ok {
			if err = d.Set("src_updt", vv); err != nil {
				return fmt.Errorf("Error reading src_updt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_updt: %v", err)
		}
	}

	if err = d.Set("srcip", flattenObjectSystemNpuNpuTcamDataSrcip3rdl(o["srcip"], d, "srcip")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcip"], "ObjectSystemNpuNpuTcamData-Srcip"); ok {
			if err = d.Set("srcip", vv); err != nil {
				return fmt.Errorf("Error reading srcip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcip: %v", err)
		}
	}

	if err = d.Set("srcipv6", flattenObjectSystemNpuNpuTcamDataSrcipv63rdl(o["srcipv6"], d, "srcipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcipv6"], "ObjectSystemNpuNpuTcamData-Srcipv6"); ok {
			if err = d.Set("srcipv6", vv); err != nil {
				return fmt.Errorf("Error reading srcipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcipv6: %v", err)
		}
	}

	if err = d.Set("srcmac", flattenObjectSystemNpuNpuTcamDataSrcmac3rdl(o["srcmac"], d, "srcmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcmac"], "ObjectSystemNpuNpuTcamData-Srcmac"); ok {
			if err = d.Set("srcmac", vv); err != nil {
				return fmt.Errorf("Error reading srcmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcmac: %v", err)
		}
	}

	if err = d.Set("srcport", flattenObjectSystemNpuNpuTcamDataSrcport3rdl(o["srcport"], d, "srcport")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcport"], "ObjectSystemNpuNpuTcamData-Srcport"); ok {
			if err = d.Set("srcport", vv); err != nil {
				return fmt.Errorf("Error reading srcport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcport: %v", err)
		}
	}

	if err = d.Set("svid", flattenObjectSystemNpuNpuTcamDataSvid3rdl(o["svid"], d, "svid")); err != nil {
		if vv, ok := fortiAPIPatch(o["svid"], "ObjectSystemNpuNpuTcamData-Svid"); ok {
			if err = d.Set("svid", vv); err != nil {
				return fmt.Errorf("Error reading svid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svid: %v", err)
		}
	}

	if err = d.Set("tcp_ack", flattenObjectSystemNpuNpuTcamDataTcpAck3rdl(o["tcp-ack"], d, "tcp_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-ack"], "ObjectSystemNpuNpuTcamData-TcpAck"); ok {
			if err = d.Set("tcp_ack", vv); err != nil {
				return fmt.Errorf("Error reading tcp_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_ack: %v", err)
		}
	}

	if err = d.Set("tcp_cwr", flattenObjectSystemNpuNpuTcamDataTcpCwr3rdl(o["tcp-cwr"], d, "tcp_cwr")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-cwr"], "ObjectSystemNpuNpuTcamData-TcpCwr"); ok {
			if err = d.Set("tcp_cwr", vv); err != nil {
				return fmt.Errorf("Error reading tcp_cwr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_cwr: %v", err)
		}
	}

	if err = d.Set("tcp_ece", flattenObjectSystemNpuNpuTcamDataTcpEce3rdl(o["tcp-ece"], d, "tcp_ece")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-ece"], "ObjectSystemNpuNpuTcamData-TcpEce"); ok {
			if err = d.Set("tcp_ece", vv); err != nil {
				return fmt.Errorf("Error reading tcp_ece: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_ece: %v", err)
		}
	}

	if err = d.Set("tcp_fin", flattenObjectSystemNpuNpuTcamDataTcpFin3rdl(o["tcp-fin"], d, "tcp_fin")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-fin"], "ObjectSystemNpuNpuTcamData-TcpFin"); ok {
			if err = d.Set("tcp_fin", vv); err != nil {
				return fmt.Errorf("Error reading tcp_fin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_fin: %v", err)
		}
	}

	if err = d.Set("tcp_push", flattenObjectSystemNpuNpuTcamDataTcpPush3rdl(o["tcp-push"], d, "tcp_push")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-push"], "ObjectSystemNpuNpuTcamData-TcpPush"); ok {
			if err = d.Set("tcp_push", vv); err != nil {
				return fmt.Errorf("Error reading tcp_push: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_push: %v", err)
		}
	}

	if err = d.Set("tcp_rst", flattenObjectSystemNpuNpuTcamDataTcpRst3rdl(o["tcp-rst"], d, "tcp_rst")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst"], "ObjectSystemNpuNpuTcamData-TcpRst"); ok {
			if err = d.Set("tcp_rst", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst: %v", err)
		}
	}

	if err = d.Set("tcp_syn", flattenObjectSystemNpuNpuTcamDataTcpSyn3rdl(o["tcp-syn"], d, "tcp_syn")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-syn"], "ObjectSystemNpuNpuTcamData-TcpSyn"); ok {
			if err = d.Set("tcp_syn", vv); err != nil {
				return fmt.Errorf("Error reading tcp_syn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_syn: %v", err)
		}
	}

	if err = d.Set("tcp_urg", flattenObjectSystemNpuNpuTcamDataTcpUrg3rdl(o["tcp-urg"], d, "tcp_urg")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-urg"], "ObjectSystemNpuNpuTcamData-TcpUrg"); ok {
			if err = d.Set("tcp_urg", vv); err != nil {
				return fmt.Errorf("Error reading tcp_urg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_urg: %v", err)
		}
	}

	if err = d.Set("tgt_cfi", flattenObjectSystemNpuNpuTcamDataTgtCfi3rdl(o["tgt-cfi"], d, "tgt_cfi")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-cfi"], "ObjectSystemNpuNpuTcamData-TgtCfi"); ok {
			if err = d.Set("tgt_cfi", vv); err != nil {
				return fmt.Errorf("Error reading tgt_cfi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_cfi: %v", err)
		}
	}

	if err = d.Set("tgt_prio", flattenObjectSystemNpuNpuTcamDataTgtPrio3rdl(o["tgt-prio"], d, "tgt_prio")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-prio"], "ObjectSystemNpuNpuTcamData-TgtPrio"); ok {
			if err = d.Set("tgt_prio", vv); err != nil {
				return fmt.Errorf("Error reading tgt_prio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_prio: %v", err)
		}
	}

	if err = d.Set("tgt_updt", flattenObjectSystemNpuNpuTcamDataTgtUpdt3rdl(o["tgt-updt"], d, "tgt_updt")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-updt"], "ObjectSystemNpuNpuTcamData-TgtUpdt"); ok {
			if err = d.Set("tgt_updt", vv); err != nil {
				return fmt.Errorf("Error reading tgt_updt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_updt: %v", err)
		}
	}

	if err = d.Set("tgt_v", flattenObjectSystemNpuNpuTcamDataTgtV3rdl(o["tgt-v"], d, "tgt_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-v"], "ObjectSystemNpuNpuTcamData-TgtV"); ok {
			if err = d.Set("tgt_v", vv); err != nil {
				return fmt.Errorf("Error reading tgt_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_v: %v", err)
		}
	}

	if err = d.Set("tos", flattenObjectSystemNpuNpuTcamDataTos3rdl(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "ObjectSystemNpuNpuTcamData-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tp", flattenObjectSystemNpuNpuTcamDataTp3rdl(o["tp"], d, "tp")); err != nil {
		if vv, ok := fortiAPIPatch(o["tp"], "ObjectSystemNpuNpuTcamData-Tp"); ok {
			if err = d.Set("tp", vv); err != nil {
				return fmt.Errorf("Error reading tp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tp: %v", err)
		}
	}

	if err = d.Set("ttl", flattenObjectSystemNpuNpuTcamDataTtl3rdl(o["ttl"], d, "ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl"], "ObjectSystemNpuNpuTcamData-Ttl"); ok {
			if err = d.Set("ttl", vv); err != nil {
				return fmt.Errorf("Error reading ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("tvid", flattenObjectSystemNpuNpuTcamDataTvid3rdl(o["tvid"], d, "tvid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tvid"], "ObjectSystemNpuNpuTcamData-Tvid"); ok {
			if err = d.Set("tvid", vv); err != nil {
				return fmt.Errorf("Error reading tvid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tvid: %v", err)
		}
	}

	if err = d.Set("vdid", flattenObjectSystemNpuNpuTcamDataVdid3rdl(o["vdid"], d, "vdid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdid"], "ObjectSystemNpuNpuTcamData-Vdid"); ok {
			if err = d.Set("vdid", vv); err != nil {
				return fmt.Errorf("Error reading vdid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdid: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamDataFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamDataDf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstipv63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstmac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataEthertype3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataExtTag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataFragOff3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenBufCnt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenIv3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenL3Flags3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenL4Flags3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPktCtrl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPri3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPriV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenTv3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIhl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIp4Id3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIp6Fl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIpver3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd103rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd113rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd83rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd93rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataMf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataProtocol3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSlink3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSmacChange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcCfi3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcPrio3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcUpdt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcipv63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcmac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSvid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpAck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpCwr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpEce3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpFin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpPush3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpRst3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpSyn3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpUrg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtCfi3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtPrio3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtUpdt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTos3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTtl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTvid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataVdid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcamData(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("df"); ok || d.HasChange("df") {
		t, err := expandObjectSystemNpuNpuTcamDataDf3rdl(d, v, "df")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["df"] = t
		}
	}

	if v, ok := d.GetOk("dstip"); ok || d.HasChange("dstip") {
		t, err := expandObjectSystemNpuNpuTcamDataDstip3rdl(d, v, "dstip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstip"] = t
		}
	}

	if v, ok := d.GetOk("dstipv6"); ok || d.HasChange("dstipv6") {
		t, err := expandObjectSystemNpuNpuTcamDataDstipv63rdl(d, v, "dstipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstipv6"] = t
		}
	}

	if v, ok := d.GetOk("dstmac"); ok || d.HasChange("dstmac") {
		t, err := expandObjectSystemNpuNpuTcamDataDstmac3rdl(d, v, "dstmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstmac"] = t
		}
	}

	if v, ok := d.GetOk("dstport"); ok || d.HasChange("dstport") {
		t, err := expandObjectSystemNpuNpuTcamDataDstport3rdl(d, v, "dstport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	if v, ok := d.GetOk("ethertype"); ok || d.HasChange("ethertype") {
		t, err := expandObjectSystemNpuNpuTcamDataEthertype3rdl(d, v, "ethertype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ethertype"] = t
		}
	}

	if v, ok := d.GetOk("ext_tag"); ok || d.HasChange("ext_tag") {
		t, err := expandObjectSystemNpuNpuTcamDataExtTag3rdl(d, v, "ext_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-tag"] = t
		}
	}

	if v, ok := d.GetOk("frag_off"); ok || d.HasChange("frag_off") {
		t, err := expandObjectSystemNpuNpuTcamDataFragOff3rdl(d, v, "frag_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frag-off"] = t
		}
	}

	if v, ok := d.GetOk("gen_buf_cnt"); ok || d.HasChange("gen_buf_cnt") {
		t, err := expandObjectSystemNpuNpuTcamDataGenBufCnt3rdl(d, v, "gen_buf_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-buf-cnt"] = t
		}
	}

	if v, ok := d.GetOk("gen_iv"); ok || d.HasChange("gen_iv") {
		t, err := expandObjectSystemNpuNpuTcamDataGenIv3rdl(d, v, "gen_iv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-iv"] = t
		}
	}

	if v, ok := d.GetOk("gen_l3_flags"); ok || d.HasChange("gen_l3_flags") {
		t, err := expandObjectSystemNpuNpuTcamDataGenL3Flags3rdl(d, v, "gen_l3_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-l3-flags"] = t
		}
	}

	if v, ok := d.GetOk("gen_l4_flags"); ok || d.HasChange("gen_l4_flags") {
		t, err := expandObjectSystemNpuNpuTcamDataGenL4Flags3rdl(d, v, "gen_l4_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-l4-flags"] = t
		}
	}

	if v, ok := d.GetOk("gen_pkt_ctrl"); ok || d.HasChange("gen_pkt_ctrl") {
		t, err := expandObjectSystemNpuNpuTcamDataGenPktCtrl3rdl(d, v, "gen_pkt_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-pkt-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("gen_pri"); ok || d.HasChange("gen_pri") {
		t, err := expandObjectSystemNpuNpuTcamDataGenPri3rdl(d, v, "gen_pri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-pri"] = t
		}
	}

	if v, ok := d.GetOk("gen_pri_v"); ok || d.HasChange("gen_pri_v") {
		t, err := expandObjectSystemNpuNpuTcamDataGenPriV3rdl(d, v, "gen_pri_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-pri-v"] = t
		}
	}

	if v, ok := d.GetOk("gen_tv"); ok || d.HasChange("gen_tv") {
		t, err := expandObjectSystemNpuNpuTcamDataGenTv3rdl(d, v, "gen_tv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-tv"] = t
		}
	}

	if v, ok := d.GetOk("ihl"); ok || d.HasChange("ihl") {
		t, err := expandObjectSystemNpuNpuTcamDataIhl3rdl(d, v, "ihl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ihl"] = t
		}
	}

	if v, ok := d.GetOk("ip4_id"); ok || d.HasChange("ip4_id") {
		t, err := expandObjectSystemNpuNpuTcamDataIp4Id3rdl(d, v, "ip4_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-id"] = t
		}
	}

	if v, ok := d.GetOk("ip6_fl"); ok || d.HasChange("ip6_fl") {
		t, err := expandObjectSystemNpuNpuTcamDataIp6Fl3rdl(d, v, "ip6_fl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-fl"] = t
		}
	}

	if v, ok := d.GetOk("ipver"); ok || d.HasChange("ipver") {
		t, err := expandObjectSystemNpuNpuTcamDataIpver3rdl(d, v, "ipver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipver"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd10"); ok || d.HasChange("l4_wd10") {
		t, err := expandObjectSystemNpuNpuTcamDataL4Wd103rdl(d, v, "l4_wd10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd10"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd11"); ok || d.HasChange("l4_wd11") {
		t, err := expandObjectSystemNpuNpuTcamDataL4Wd113rdl(d, v, "l4_wd11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd11"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd8"); ok || d.HasChange("l4_wd8") {
		t, err := expandObjectSystemNpuNpuTcamDataL4Wd83rdl(d, v, "l4_wd8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd8"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd9"); ok || d.HasChange("l4_wd9") {
		t, err := expandObjectSystemNpuNpuTcamDataL4Wd93rdl(d, v, "l4_wd9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd9"] = t
		}
	}

	if v, ok := d.GetOk("mf"); ok || d.HasChange("mf") {
		t, err := expandObjectSystemNpuNpuTcamDataMf3rdl(d, v, "mf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mf"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectSystemNpuNpuTcamDataProtocol3rdl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("slink"); ok || d.HasChange("slink") {
		t, err := expandObjectSystemNpuNpuTcamDataSlink3rdl(d, v, "slink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slink"] = t
		}
	}

	if v, ok := d.GetOk("smac_change"); ok || d.HasChange("smac_change") {
		t, err := expandObjectSystemNpuNpuTcamDataSmacChange3rdl(d, v, "smac_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smac-change"] = t
		}
	}

	if v, ok := d.GetOk("sp"); ok || d.HasChange("sp") {
		t, err := expandObjectSystemNpuNpuTcamDataSp3rdl(d, v, "sp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp"] = t
		}
	}

	if v, ok := d.GetOk("src_cfi"); ok || d.HasChange("src_cfi") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcCfi3rdl(d, v, "src_cfi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-cfi"] = t
		}
	}

	if v, ok := d.GetOk("src_prio"); ok || d.HasChange("src_prio") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcPrio3rdl(d, v, "src_prio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-prio"] = t
		}
	}

	if v, ok := d.GetOk("src_updt"); ok || d.HasChange("src_updt") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcUpdt3rdl(d, v, "src_updt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-updt"] = t
		}
	}

	if v, ok := d.GetOk("srcip"); ok || d.HasChange("srcip") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcip3rdl(d, v, "srcip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcip"] = t
		}
	}

	if v, ok := d.GetOk("srcipv6"); ok || d.HasChange("srcipv6") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcipv63rdl(d, v, "srcipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcipv6"] = t
		}
	}

	if v, ok := d.GetOk("srcmac"); ok || d.HasChange("srcmac") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcmac3rdl(d, v, "srcmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcmac"] = t
		}
	}

	if v, ok := d.GetOk("srcport"); ok || d.HasChange("srcport") {
		t, err := expandObjectSystemNpuNpuTcamDataSrcport3rdl(d, v, "srcport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcport"] = t
		}
	}

	if v, ok := d.GetOk("svid"); ok || d.HasChange("svid") {
		t, err := expandObjectSystemNpuNpuTcamDataSvid3rdl(d, v, "svid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svid"] = t
		}
	}

	if v, ok := d.GetOk("tcp_ack"); ok || d.HasChange("tcp_ack") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpAck3rdl(d, v, "tcp_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-ack"] = t
		}
	}

	if v, ok := d.GetOk("tcp_cwr"); ok || d.HasChange("tcp_cwr") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpCwr3rdl(d, v, "tcp_cwr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-cwr"] = t
		}
	}

	if v, ok := d.GetOk("tcp_ece"); ok || d.HasChange("tcp_ece") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpEce3rdl(d, v, "tcp_ece")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-ece"] = t
		}
	}

	if v, ok := d.GetOk("tcp_fin"); ok || d.HasChange("tcp_fin") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpFin3rdl(d, v, "tcp_fin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-fin"] = t
		}
	}

	if v, ok := d.GetOk("tcp_push"); ok || d.HasChange("tcp_push") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpPush3rdl(d, v, "tcp_push")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-push"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst"); ok || d.HasChange("tcp_rst") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpRst3rdl(d, v, "tcp_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst"] = t
		}
	}

	if v, ok := d.GetOk("tcp_syn"); ok || d.HasChange("tcp_syn") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpSyn3rdl(d, v, "tcp_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-syn"] = t
		}
	}

	if v, ok := d.GetOk("tcp_urg"); ok || d.HasChange("tcp_urg") {
		t, err := expandObjectSystemNpuNpuTcamDataTcpUrg3rdl(d, v, "tcp_urg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-urg"] = t
		}
	}

	if v, ok := d.GetOk("tgt_cfi"); ok || d.HasChange("tgt_cfi") {
		t, err := expandObjectSystemNpuNpuTcamDataTgtCfi3rdl(d, v, "tgt_cfi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-cfi"] = t
		}
	}

	if v, ok := d.GetOk("tgt_prio"); ok || d.HasChange("tgt_prio") {
		t, err := expandObjectSystemNpuNpuTcamDataTgtPrio3rdl(d, v, "tgt_prio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-prio"] = t
		}
	}

	if v, ok := d.GetOk("tgt_updt"); ok || d.HasChange("tgt_updt") {
		t, err := expandObjectSystemNpuNpuTcamDataTgtUpdt3rdl(d, v, "tgt_updt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-updt"] = t
		}
	}

	if v, ok := d.GetOk("tgt_v"); ok || d.HasChange("tgt_v") {
		t, err := expandObjectSystemNpuNpuTcamDataTgtV3rdl(d, v, "tgt_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-v"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandObjectSystemNpuNpuTcamDataTos3rdl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tp"); ok || d.HasChange("tp") {
		t, err := expandObjectSystemNpuNpuTcamDataTp3rdl(d, v, "tp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tp"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok || d.HasChange("ttl") {
		t, err := expandObjectSystemNpuNpuTcamDataTtl3rdl(d, v, "ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("tvid"); ok || d.HasChange("tvid") {
		t, err := expandObjectSystemNpuNpuTcamDataTvid3rdl(d, v, "tvid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tvid"] = t
		}
	}

	if v, ok := d.GetOk("vdid"); ok || d.HasChange("vdid") {
		t, err := expandObjectSystemNpuNpuTcamDataVdid3rdl(d, v, "vdid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdid"] = t
		}
	}

	return &obj, nil
}
