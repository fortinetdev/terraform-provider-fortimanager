// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Mask fields of TCAM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcamMask() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamMaskUpdate,
		Read:   resourceObjectSystemNpuNpuTcamMaskRead,
		Update: resourceObjectSystemNpuNpuTcamMaskUpdate,
		Delete: resourceObjectSystemNpuNpuTcamMaskDelete,

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
				Computed: true,
			},
			"dstip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ethertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
				Computed: true,
			},
			"gen_tv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
				Computed: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"src_cfi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_prio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"src_updt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"tcp_cwr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_ece": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_fin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_rst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_urg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tgt_cfi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tgt_prio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tgt_updt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tgt_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceObjectSystemNpuNpuTcamMaskUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpuTcamMask(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamMask resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpuTcamMask(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamMask resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpuTcamMask")

	return resourceObjectSystemNpuNpuTcamMaskRead(d, m)
}

func resourceObjectSystemNpuNpuTcamMaskDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpuTcamMask(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcamMask resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamMaskRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpuTcamMask(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamMask resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcamMask(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamMask resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamMaskDf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstipv63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstmac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskEthertype3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskExtTag3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskFragOff3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenBufCnt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenIv3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenL3Flags3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenL4Flags3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPktCtrl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPri3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPriV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenTv3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIhl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIp4Id3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIp6Fl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIpver3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd103rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd113rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd83rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd93rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskMf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskProtocol3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSlink3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSmacChange3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcCfi3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcPrio3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcUpdt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcipv63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcmac3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcport3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSvid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpAck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpCwr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpEce3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpFin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpPush3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpRst3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpSyn3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpUrg3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtCfi3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtPrio3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtUpdt3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTos3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTtl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTvid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskVdid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcamMask(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("df", flattenObjectSystemNpuNpuTcamMaskDf3rdl(o["df"], d, "df")); err != nil {
		if vv, ok := fortiAPIPatch(o["df"], "ObjectSystemNpuNpuTcamMask-Df"); ok {
			if err = d.Set("df", vv); err != nil {
				return fmt.Errorf("Error reading df: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading df: %v", err)
		}
	}

	if err = d.Set("dstip", flattenObjectSystemNpuNpuTcamMaskDstip3rdl(o["dstip"], d, "dstip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstip"], "ObjectSystemNpuNpuTcamMask-Dstip"); ok {
			if err = d.Set("dstip", vv); err != nil {
				return fmt.Errorf("Error reading dstip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstip: %v", err)
		}
	}

	if err = d.Set("dstipv6", flattenObjectSystemNpuNpuTcamMaskDstipv63rdl(o["dstipv6"], d, "dstipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstipv6"], "ObjectSystemNpuNpuTcamMask-Dstipv6"); ok {
			if err = d.Set("dstipv6", vv); err != nil {
				return fmt.Errorf("Error reading dstipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstipv6: %v", err)
		}
	}

	if err = d.Set("dstmac", flattenObjectSystemNpuNpuTcamMaskDstmac3rdl(o["dstmac"], d, "dstmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstmac"], "ObjectSystemNpuNpuTcamMask-Dstmac"); ok {
			if err = d.Set("dstmac", vv); err != nil {
				return fmt.Errorf("Error reading dstmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstmac: %v", err)
		}
	}

	if err = d.Set("dstport", flattenObjectSystemNpuNpuTcamMaskDstport3rdl(o["dstport"], d, "dstport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport"], "ObjectSystemNpuNpuTcamMask-Dstport"); ok {
			if err = d.Set("dstport", vv); err != nil {
				return fmt.Errorf("Error reading dstport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("ethertype", flattenObjectSystemNpuNpuTcamMaskEthertype3rdl(o["ethertype"], d, "ethertype")); err != nil {
		if vv, ok := fortiAPIPatch(o["ethertype"], "ObjectSystemNpuNpuTcamMask-Ethertype"); ok {
			if err = d.Set("ethertype", vv); err != nil {
				return fmt.Errorf("Error reading ethertype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ethertype: %v", err)
		}
	}

	if err = d.Set("ext_tag", flattenObjectSystemNpuNpuTcamMaskExtTag3rdl(o["ext-tag"], d, "ext_tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-tag"], "ObjectSystemNpuNpuTcamMask-ExtTag"); ok {
			if err = d.Set("ext_tag", vv); err != nil {
				return fmt.Errorf("Error reading ext_tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_tag: %v", err)
		}
	}

	if err = d.Set("frag_off", flattenObjectSystemNpuNpuTcamMaskFragOff3rdl(o["frag-off"], d, "frag_off")); err != nil {
		if vv, ok := fortiAPIPatch(o["frag-off"], "ObjectSystemNpuNpuTcamMask-FragOff"); ok {
			if err = d.Set("frag_off", vv); err != nil {
				return fmt.Errorf("Error reading frag_off: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frag_off: %v", err)
		}
	}

	if err = d.Set("gen_buf_cnt", flattenObjectSystemNpuNpuTcamMaskGenBufCnt3rdl(o["gen-buf-cnt"], d, "gen_buf_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-buf-cnt"], "ObjectSystemNpuNpuTcamMask-GenBufCnt"); ok {
			if err = d.Set("gen_buf_cnt", vv); err != nil {
				return fmt.Errorf("Error reading gen_buf_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_buf_cnt: %v", err)
		}
	}

	if err = d.Set("gen_iv", flattenObjectSystemNpuNpuTcamMaskGenIv3rdl(o["gen-iv"], d, "gen_iv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-iv"], "ObjectSystemNpuNpuTcamMask-GenIv"); ok {
			if err = d.Set("gen_iv", vv); err != nil {
				return fmt.Errorf("Error reading gen_iv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_iv: %v", err)
		}
	}

	if err = d.Set("gen_l3_flags", flattenObjectSystemNpuNpuTcamMaskGenL3Flags3rdl(o["gen-l3-flags"], d, "gen_l3_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-l3-flags"], "ObjectSystemNpuNpuTcamMask-GenL3Flags"); ok {
			if err = d.Set("gen_l3_flags", vv); err != nil {
				return fmt.Errorf("Error reading gen_l3_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_l3_flags: %v", err)
		}
	}

	if err = d.Set("gen_l4_flags", flattenObjectSystemNpuNpuTcamMaskGenL4Flags3rdl(o["gen-l4-flags"], d, "gen_l4_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-l4-flags"], "ObjectSystemNpuNpuTcamMask-GenL4Flags"); ok {
			if err = d.Set("gen_l4_flags", vv); err != nil {
				return fmt.Errorf("Error reading gen_l4_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_l4_flags: %v", err)
		}
	}

	if err = d.Set("gen_pkt_ctrl", flattenObjectSystemNpuNpuTcamMaskGenPktCtrl3rdl(o["gen-pkt-ctrl"], d, "gen_pkt_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-pkt-ctrl"], "ObjectSystemNpuNpuTcamMask-GenPktCtrl"); ok {
			if err = d.Set("gen_pkt_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading gen_pkt_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_pkt_ctrl: %v", err)
		}
	}

	if err = d.Set("gen_pri", flattenObjectSystemNpuNpuTcamMaskGenPri3rdl(o["gen-pri"], d, "gen_pri")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-pri"], "ObjectSystemNpuNpuTcamMask-GenPri"); ok {
			if err = d.Set("gen_pri", vv); err != nil {
				return fmt.Errorf("Error reading gen_pri: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_pri: %v", err)
		}
	}

	if err = d.Set("gen_pri_v", flattenObjectSystemNpuNpuTcamMaskGenPriV3rdl(o["gen-pri-v"], d, "gen_pri_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-pri-v"], "ObjectSystemNpuNpuTcamMask-GenPriV"); ok {
			if err = d.Set("gen_pri_v", vv); err != nil {
				return fmt.Errorf("Error reading gen_pri_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_pri_v: %v", err)
		}
	}

	if err = d.Set("gen_tv", flattenObjectSystemNpuNpuTcamMaskGenTv3rdl(o["gen-tv"], d, "gen_tv")); err != nil {
		if vv, ok := fortiAPIPatch(o["gen-tv"], "ObjectSystemNpuNpuTcamMask-GenTv"); ok {
			if err = d.Set("gen_tv", vv); err != nil {
				return fmt.Errorf("Error reading gen_tv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gen_tv: %v", err)
		}
	}

	if err = d.Set("ihl", flattenObjectSystemNpuNpuTcamMaskIhl3rdl(o["ihl"], d, "ihl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ihl"], "ObjectSystemNpuNpuTcamMask-Ihl"); ok {
			if err = d.Set("ihl", vv); err != nil {
				return fmt.Errorf("Error reading ihl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ihl: %v", err)
		}
	}

	if err = d.Set("ip4_id", flattenObjectSystemNpuNpuTcamMaskIp4Id3rdl(o["ip4-id"], d, "ip4_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip4-id"], "ObjectSystemNpuNpuTcamMask-Ip4Id"); ok {
			if err = d.Set("ip4_id", vv); err != nil {
				return fmt.Errorf("Error reading ip4_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip4_id: %v", err)
		}
	}

	if err = d.Set("ip6_fl", flattenObjectSystemNpuNpuTcamMaskIp6Fl3rdl(o["ip6-fl"], d, "ip6_fl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-fl"], "ObjectSystemNpuNpuTcamMask-Ip6Fl"); ok {
			if err = d.Set("ip6_fl", vv); err != nil {
				return fmt.Errorf("Error reading ip6_fl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_fl: %v", err)
		}
	}

	if err = d.Set("ipver", flattenObjectSystemNpuNpuTcamMaskIpver3rdl(o["ipver"], d, "ipver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipver"], "ObjectSystemNpuNpuTcamMask-Ipver"); ok {
			if err = d.Set("ipver", vv); err != nil {
				return fmt.Errorf("Error reading ipver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipver: %v", err)
		}
	}

	if err = d.Set("l4_wd10", flattenObjectSystemNpuNpuTcamMaskL4Wd103rdl(o["l4-wd10"], d, "l4_wd10")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd10"], "ObjectSystemNpuNpuTcamMask-L4Wd10"); ok {
			if err = d.Set("l4_wd10", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd10: %v", err)
		}
	}

	if err = d.Set("l4_wd11", flattenObjectSystemNpuNpuTcamMaskL4Wd113rdl(o["l4-wd11"], d, "l4_wd11")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd11"], "ObjectSystemNpuNpuTcamMask-L4Wd11"); ok {
			if err = d.Set("l4_wd11", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd11: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd11: %v", err)
		}
	}

	if err = d.Set("l4_wd8", flattenObjectSystemNpuNpuTcamMaskL4Wd83rdl(o["l4-wd8"], d, "l4_wd8")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd8"], "ObjectSystemNpuNpuTcamMask-L4Wd8"); ok {
			if err = d.Set("l4_wd8", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd8: %v", err)
		}
	}

	if err = d.Set("l4_wd9", flattenObjectSystemNpuNpuTcamMaskL4Wd93rdl(o["l4-wd9"], d, "l4_wd9")); err != nil {
		if vv, ok := fortiAPIPatch(o["l4-wd9"], "ObjectSystemNpuNpuTcamMask-L4Wd9"); ok {
			if err = d.Set("l4_wd9", vv); err != nil {
				return fmt.Errorf("Error reading l4_wd9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading l4_wd9: %v", err)
		}
	}

	if err = d.Set("mf", flattenObjectSystemNpuNpuTcamMaskMf3rdl(o["mf"], d, "mf")); err != nil {
		if vv, ok := fortiAPIPatch(o["mf"], "ObjectSystemNpuNpuTcamMask-Mf"); ok {
			if err = d.Set("mf", vv); err != nil {
				return fmt.Errorf("Error reading mf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mf: %v", err)
		}
	}

	if err = d.Set("protocol", flattenObjectSystemNpuNpuTcamMaskProtocol3rdl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "ObjectSystemNpuNpuTcamMask-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("slink", flattenObjectSystemNpuNpuTcamMaskSlink3rdl(o["slink"], d, "slink")); err != nil {
		if vv, ok := fortiAPIPatch(o["slink"], "ObjectSystemNpuNpuTcamMask-Slink"); ok {
			if err = d.Set("slink", vv); err != nil {
				return fmt.Errorf("Error reading slink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slink: %v", err)
		}
	}

	if err = d.Set("smac_change", flattenObjectSystemNpuNpuTcamMaskSmacChange3rdl(o["smac-change"], d, "smac_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["smac-change"], "ObjectSystemNpuNpuTcamMask-SmacChange"); ok {
			if err = d.Set("smac_change", vv); err != nil {
				return fmt.Errorf("Error reading smac_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smac_change: %v", err)
		}
	}

	if err = d.Set("sp", flattenObjectSystemNpuNpuTcamMaskSp3rdl(o["sp"], d, "sp")); err != nil {
		if vv, ok := fortiAPIPatch(o["sp"], "ObjectSystemNpuNpuTcamMask-Sp"); ok {
			if err = d.Set("sp", vv); err != nil {
				return fmt.Errorf("Error reading sp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sp: %v", err)
		}
	}

	if err = d.Set("src_cfi", flattenObjectSystemNpuNpuTcamMaskSrcCfi3rdl(o["src-cfi"], d, "src_cfi")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-cfi"], "ObjectSystemNpuNpuTcamMask-SrcCfi"); ok {
			if err = d.Set("src_cfi", vv); err != nil {
				return fmt.Errorf("Error reading src_cfi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_cfi: %v", err)
		}
	}

	if err = d.Set("src_prio", flattenObjectSystemNpuNpuTcamMaskSrcPrio3rdl(o["src-prio"], d, "src_prio")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-prio"], "ObjectSystemNpuNpuTcamMask-SrcPrio"); ok {
			if err = d.Set("src_prio", vv); err != nil {
				return fmt.Errorf("Error reading src_prio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_prio: %v", err)
		}
	}

	if err = d.Set("src_updt", flattenObjectSystemNpuNpuTcamMaskSrcUpdt3rdl(o["src-updt"], d, "src_updt")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-updt"], "ObjectSystemNpuNpuTcamMask-SrcUpdt"); ok {
			if err = d.Set("src_updt", vv); err != nil {
				return fmt.Errorf("Error reading src_updt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_updt: %v", err)
		}
	}

	if err = d.Set("srcip", flattenObjectSystemNpuNpuTcamMaskSrcip3rdl(o["srcip"], d, "srcip")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcip"], "ObjectSystemNpuNpuTcamMask-Srcip"); ok {
			if err = d.Set("srcip", vv); err != nil {
				return fmt.Errorf("Error reading srcip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcip: %v", err)
		}
	}

	if err = d.Set("srcipv6", flattenObjectSystemNpuNpuTcamMaskSrcipv63rdl(o["srcipv6"], d, "srcipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcipv6"], "ObjectSystemNpuNpuTcamMask-Srcipv6"); ok {
			if err = d.Set("srcipv6", vv); err != nil {
				return fmt.Errorf("Error reading srcipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcipv6: %v", err)
		}
	}

	if err = d.Set("srcmac", flattenObjectSystemNpuNpuTcamMaskSrcmac3rdl(o["srcmac"], d, "srcmac")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcmac"], "ObjectSystemNpuNpuTcamMask-Srcmac"); ok {
			if err = d.Set("srcmac", vv); err != nil {
				return fmt.Errorf("Error reading srcmac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcmac: %v", err)
		}
	}

	if err = d.Set("srcport", flattenObjectSystemNpuNpuTcamMaskSrcport3rdl(o["srcport"], d, "srcport")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcport"], "ObjectSystemNpuNpuTcamMask-Srcport"); ok {
			if err = d.Set("srcport", vv); err != nil {
				return fmt.Errorf("Error reading srcport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcport: %v", err)
		}
	}

	if err = d.Set("svid", flattenObjectSystemNpuNpuTcamMaskSvid3rdl(o["svid"], d, "svid")); err != nil {
		if vv, ok := fortiAPIPatch(o["svid"], "ObjectSystemNpuNpuTcamMask-Svid"); ok {
			if err = d.Set("svid", vv); err != nil {
				return fmt.Errorf("Error reading svid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading svid: %v", err)
		}
	}

	if err = d.Set("tcp_ack", flattenObjectSystemNpuNpuTcamMaskTcpAck3rdl(o["tcp-ack"], d, "tcp_ack")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-ack"], "ObjectSystemNpuNpuTcamMask-TcpAck"); ok {
			if err = d.Set("tcp_ack", vv); err != nil {
				return fmt.Errorf("Error reading tcp_ack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_ack: %v", err)
		}
	}

	if err = d.Set("tcp_cwr", flattenObjectSystemNpuNpuTcamMaskTcpCwr3rdl(o["tcp-cwr"], d, "tcp_cwr")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-cwr"], "ObjectSystemNpuNpuTcamMask-TcpCwr"); ok {
			if err = d.Set("tcp_cwr", vv); err != nil {
				return fmt.Errorf("Error reading tcp_cwr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_cwr: %v", err)
		}
	}

	if err = d.Set("tcp_ece", flattenObjectSystemNpuNpuTcamMaskTcpEce3rdl(o["tcp-ece"], d, "tcp_ece")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-ece"], "ObjectSystemNpuNpuTcamMask-TcpEce"); ok {
			if err = d.Set("tcp_ece", vv); err != nil {
				return fmt.Errorf("Error reading tcp_ece: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_ece: %v", err)
		}
	}

	if err = d.Set("tcp_fin", flattenObjectSystemNpuNpuTcamMaskTcpFin3rdl(o["tcp-fin"], d, "tcp_fin")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-fin"], "ObjectSystemNpuNpuTcamMask-TcpFin"); ok {
			if err = d.Set("tcp_fin", vv); err != nil {
				return fmt.Errorf("Error reading tcp_fin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_fin: %v", err)
		}
	}

	if err = d.Set("tcp_push", flattenObjectSystemNpuNpuTcamMaskTcpPush3rdl(o["tcp-push"], d, "tcp_push")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-push"], "ObjectSystemNpuNpuTcamMask-TcpPush"); ok {
			if err = d.Set("tcp_push", vv); err != nil {
				return fmt.Errorf("Error reading tcp_push: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_push: %v", err)
		}
	}

	if err = d.Set("tcp_rst", flattenObjectSystemNpuNpuTcamMaskTcpRst3rdl(o["tcp-rst"], d, "tcp_rst")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-rst"], "ObjectSystemNpuNpuTcamMask-TcpRst"); ok {
			if err = d.Set("tcp_rst", vv); err != nil {
				return fmt.Errorf("Error reading tcp_rst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_rst: %v", err)
		}
	}

	if err = d.Set("tcp_syn", flattenObjectSystemNpuNpuTcamMaskTcpSyn3rdl(o["tcp-syn"], d, "tcp_syn")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-syn"], "ObjectSystemNpuNpuTcamMask-TcpSyn"); ok {
			if err = d.Set("tcp_syn", vv); err != nil {
				return fmt.Errorf("Error reading tcp_syn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_syn: %v", err)
		}
	}

	if err = d.Set("tcp_urg", flattenObjectSystemNpuNpuTcamMaskTcpUrg3rdl(o["tcp-urg"], d, "tcp_urg")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-urg"], "ObjectSystemNpuNpuTcamMask-TcpUrg"); ok {
			if err = d.Set("tcp_urg", vv); err != nil {
				return fmt.Errorf("Error reading tcp_urg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_urg: %v", err)
		}
	}

	if err = d.Set("tgt_cfi", flattenObjectSystemNpuNpuTcamMaskTgtCfi3rdl(o["tgt-cfi"], d, "tgt_cfi")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-cfi"], "ObjectSystemNpuNpuTcamMask-TgtCfi"); ok {
			if err = d.Set("tgt_cfi", vv); err != nil {
				return fmt.Errorf("Error reading tgt_cfi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_cfi: %v", err)
		}
	}

	if err = d.Set("tgt_prio", flattenObjectSystemNpuNpuTcamMaskTgtPrio3rdl(o["tgt-prio"], d, "tgt_prio")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-prio"], "ObjectSystemNpuNpuTcamMask-TgtPrio"); ok {
			if err = d.Set("tgt_prio", vv); err != nil {
				return fmt.Errorf("Error reading tgt_prio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_prio: %v", err)
		}
	}

	if err = d.Set("tgt_updt", flattenObjectSystemNpuNpuTcamMaskTgtUpdt3rdl(o["tgt-updt"], d, "tgt_updt")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-updt"], "ObjectSystemNpuNpuTcamMask-TgtUpdt"); ok {
			if err = d.Set("tgt_updt", vv); err != nil {
				return fmt.Errorf("Error reading tgt_updt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_updt: %v", err)
		}
	}

	if err = d.Set("tgt_v", flattenObjectSystemNpuNpuTcamMaskTgtV3rdl(o["tgt-v"], d, "tgt_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgt-v"], "ObjectSystemNpuNpuTcamMask-TgtV"); ok {
			if err = d.Set("tgt_v", vv); err != nil {
				return fmt.Errorf("Error reading tgt_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgt_v: %v", err)
		}
	}

	if err = d.Set("tos", flattenObjectSystemNpuNpuTcamMaskTos3rdl(o["tos"], d, "tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["tos"], "ObjectSystemNpuNpuTcamMask-Tos"); ok {
			if err = d.Set("tos", vv); err != nil {
				return fmt.Errorf("Error reading tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tp", flattenObjectSystemNpuNpuTcamMaskTp3rdl(o["tp"], d, "tp")); err != nil {
		if vv, ok := fortiAPIPatch(o["tp"], "ObjectSystemNpuNpuTcamMask-Tp"); ok {
			if err = d.Set("tp", vv); err != nil {
				return fmt.Errorf("Error reading tp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tp: %v", err)
		}
	}

	if err = d.Set("ttl", flattenObjectSystemNpuNpuTcamMaskTtl3rdl(o["ttl"], d, "ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["ttl"], "ObjectSystemNpuNpuTcamMask-Ttl"); ok {
			if err = d.Set("ttl", vv); err != nil {
				return fmt.Errorf("Error reading ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("tvid", flattenObjectSystemNpuNpuTcamMaskTvid3rdl(o["tvid"], d, "tvid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tvid"], "ObjectSystemNpuNpuTcamMask-Tvid"); ok {
			if err = d.Set("tvid", vv); err != nil {
				return fmt.Errorf("Error reading tvid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tvid: %v", err)
		}
	}

	if err = d.Set("vdid", flattenObjectSystemNpuNpuTcamMaskVdid3rdl(o["vdid"], d, "vdid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdid"], "ObjectSystemNpuNpuTcamMask-Vdid"); ok {
			if err = d.Set("vdid", vv); err != nil {
				return fmt.Errorf("Error reading vdid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdid: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamMaskFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamMaskDf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstipv63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstmac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskEthertype3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskExtTag3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskFragOff3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenBufCnt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenIv3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenL3Flags3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenL4Flags3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPktCtrl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPri3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPriV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenTv3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIhl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIp4Id3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIp6Fl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIpver3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd103rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd113rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd83rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd93rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskMf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskProtocol3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSlink3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSmacChange3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcCfi3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcPrio3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcUpdt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcipv63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcmac3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcport3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSvid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpAck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpCwr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpEce3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpFin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpPush3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpRst3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpSyn3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpUrg3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtCfi3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtPrio3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtUpdt3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTos3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTtl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTvid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskVdid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcamMask(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("df"); ok || d.HasChange("df") {
		t, err := expandObjectSystemNpuNpuTcamMaskDf3rdl(d, v, "df")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["df"] = t
		}
	}

	if v, ok := d.GetOk("dstip"); ok || d.HasChange("dstip") {
		t, err := expandObjectSystemNpuNpuTcamMaskDstip3rdl(d, v, "dstip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstip"] = t
		}
	}

	if v, ok := d.GetOk("dstipv6"); ok || d.HasChange("dstipv6") {
		t, err := expandObjectSystemNpuNpuTcamMaskDstipv63rdl(d, v, "dstipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstipv6"] = t
		}
	}

	if v, ok := d.GetOk("dstmac"); ok || d.HasChange("dstmac") {
		t, err := expandObjectSystemNpuNpuTcamMaskDstmac3rdl(d, v, "dstmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstmac"] = t
		}
	}

	if v, ok := d.GetOk("dstport"); ok || d.HasChange("dstport") {
		t, err := expandObjectSystemNpuNpuTcamMaskDstport3rdl(d, v, "dstport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	if v, ok := d.GetOk("ethertype"); ok || d.HasChange("ethertype") {
		t, err := expandObjectSystemNpuNpuTcamMaskEthertype3rdl(d, v, "ethertype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ethertype"] = t
		}
	}

	if v, ok := d.GetOk("ext_tag"); ok || d.HasChange("ext_tag") {
		t, err := expandObjectSystemNpuNpuTcamMaskExtTag3rdl(d, v, "ext_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-tag"] = t
		}
	}

	if v, ok := d.GetOk("frag_off"); ok || d.HasChange("frag_off") {
		t, err := expandObjectSystemNpuNpuTcamMaskFragOff3rdl(d, v, "frag_off")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frag-off"] = t
		}
	}

	if v, ok := d.GetOk("gen_buf_cnt"); ok || d.HasChange("gen_buf_cnt") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenBufCnt3rdl(d, v, "gen_buf_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-buf-cnt"] = t
		}
	}

	if v, ok := d.GetOk("gen_iv"); ok || d.HasChange("gen_iv") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenIv3rdl(d, v, "gen_iv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-iv"] = t
		}
	}

	if v, ok := d.GetOk("gen_l3_flags"); ok || d.HasChange("gen_l3_flags") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenL3Flags3rdl(d, v, "gen_l3_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-l3-flags"] = t
		}
	}

	if v, ok := d.GetOk("gen_l4_flags"); ok || d.HasChange("gen_l4_flags") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenL4Flags3rdl(d, v, "gen_l4_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-l4-flags"] = t
		}
	}

	if v, ok := d.GetOk("gen_pkt_ctrl"); ok || d.HasChange("gen_pkt_ctrl") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenPktCtrl3rdl(d, v, "gen_pkt_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-pkt-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("gen_pri"); ok || d.HasChange("gen_pri") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenPri3rdl(d, v, "gen_pri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-pri"] = t
		}
	}

	if v, ok := d.GetOk("gen_pri_v"); ok || d.HasChange("gen_pri_v") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenPriV3rdl(d, v, "gen_pri_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-pri-v"] = t
		}
	}

	if v, ok := d.GetOk("gen_tv"); ok || d.HasChange("gen_tv") {
		t, err := expandObjectSystemNpuNpuTcamMaskGenTv3rdl(d, v, "gen_tv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gen-tv"] = t
		}
	}

	if v, ok := d.GetOk("ihl"); ok || d.HasChange("ihl") {
		t, err := expandObjectSystemNpuNpuTcamMaskIhl3rdl(d, v, "ihl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ihl"] = t
		}
	}

	if v, ok := d.GetOk("ip4_id"); ok || d.HasChange("ip4_id") {
		t, err := expandObjectSystemNpuNpuTcamMaskIp4Id3rdl(d, v, "ip4_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-id"] = t
		}
	}

	if v, ok := d.GetOk("ip6_fl"); ok || d.HasChange("ip6_fl") {
		t, err := expandObjectSystemNpuNpuTcamMaskIp6Fl3rdl(d, v, "ip6_fl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-fl"] = t
		}
	}

	if v, ok := d.GetOk("ipver"); ok || d.HasChange("ipver") {
		t, err := expandObjectSystemNpuNpuTcamMaskIpver3rdl(d, v, "ipver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipver"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd10"); ok || d.HasChange("l4_wd10") {
		t, err := expandObjectSystemNpuNpuTcamMaskL4Wd103rdl(d, v, "l4_wd10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd10"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd11"); ok || d.HasChange("l4_wd11") {
		t, err := expandObjectSystemNpuNpuTcamMaskL4Wd113rdl(d, v, "l4_wd11")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd11"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd8"); ok || d.HasChange("l4_wd8") {
		t, err := expandObjectSystemNpuNpuTcamMaskL4Wd83rdl(d, v, "l4_wd8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd8"] = t
		}
	}

	if v, ok := d.GetOk("l4_wd9"); ok || d.HasChange("l4_wd9") {
		t, err := expandObjectSystemNpuNpuTcamMaskL4Wd93rdl(d, v, "l4_wd9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l4-wd9"] = t
		}
	}

	if v, ok := d.GetOk("mf"); ok || d.HasChange("mf") {
		t, err := expandObjectSystemNpuNpuTcamMaskMf3rdl(d, v, "mf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mf"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandObjectSystemNpuNpuTcamMaskProtocol3rdl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("slink"); ok || d.HasChange("slink") {
		t, err := expandObjectSystemNpuNpuTcamMaskSlink3rdl(d, v, "slink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slink"] = t
		}
	}

	if v, ok := d.GetOk("smac_change"); ok || d.HasChange("smac_change") {
		t, err := expandObjectSystemNpuNpuTcamMaskSmacChange3rdl(d, v, "smac_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smac-change"] = t
		}
	}

	if v, ok := d.GetOk("sp"); ok || d.HasChange("sp") {
		t, err := expandObjectSystemNpuNpuTcamMaskSp3rdl(d, v, "sp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sp"] = t
		}
	}

	if v, ok := d.GetOk("src_cfi"); ok || d.HasChange("src_cfi") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcCfi3rdl(d, v, "src_cfi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-cfi"] = t
		}
	}

	if v, ok := d.GetOk("src_prio"); ok || d.HasChange("src_prio") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcPrio3rdl(d, v, "src_prio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-prio"] = t
		}
	}

	if v, ok := d.GetOk("src_updt"); ok || d.HasChange("src_updt") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcUpdt3rdl(d, v, "src_updt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-updt"] = t
		}
	}

	if v, ok := d.GetOk("srcip"); ok || d.HasChange("srcip") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcip3rdl(d, v, "srcip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcip"] = t
		}
	}

	if v, ok := d.GetOk("srcipv6"); ok || d.HasChange("srcipv6") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcipv63rdl(d, v, "srcipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcipv6"] = t
		}
	}

	if v, ok := d.GetOk("srcmac"); ok || d.HasChange("srcmac") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcmac3rdl(d, v, "srcmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcmac"] = t
		}
	}

	if v, ok := d.GetOk("srcport"); ok || d.HasChange("srcport") {
		t, err := expandObjectSystemNpuNpuTcamMaskSrcport3rdl(d, v, "srcport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcport"] = t
		}
	}

	if v, ok := d.GetOk("svid"); ok || d.HasChange("svid") {
		t, err := expandObjectSystemNpuNpuTcamMaskSvid3rdl(d, v, "svid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svid"] = t
		}
	}

	if v, ok := d.GetOk("tcp_ack"); ok || d.HasChange("tcp_ack") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpAck3rdl(d, v, "tcp_ack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-ack"] = t
		}
	}

	if v, ok := d.GetOk("tcp_cwr"); ok || d.HasChange("tcp_cwr") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpCwr3rdl(d, v, "tcp_cwr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-cwr"] = t
		}
	}

	if v, ok := d.GetOk("tcp_ece"); ok || d.HasChange("tcp_ece") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpEce3rdl(d, v, "tcp_ece")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-ece"] = t
		}
	}

	if v, ok := d.GetOk("tcp_fin"); ok || d.HasChange("tcp_fin") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpFin3rdl(d, v, "tcp_fin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-fin"] = t
		}
	}

	if v, ok := d.GetOk("tcp_push"); ok || d.HasChange("tcp_push") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpPush3rdl(d, v, "tcp_push")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-push"] = t
		}
	}

	if v, ok := d.GetOk("tcp_rst"); ok || d.HasChange("tcp_rst") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpRst3rdl(d, v, "tcp_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-rst"] = t
		}
	}

	if v, ok := d.GetOk("tcp_syn"); ok || d.HasChange("tcp_syn") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpSyn3rdl(d, v, "tcp_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-syn"] = t
		}
	}

	if v, ok := d.GetOk("tcp_urg"); ok || d.HasChange("tcp_urg") {
		t, err := expandObjectSystemNpuNpuTcamMaskTcpUrg3rdl(d, v, "tcp_urg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-urg"] = t
		}
	}

	if v, ok := d.GetOk("tgt_cfi"); ok || d.HasChange("tgt_cfi") {
		t, err := expandObjectSystemNpuNpuTcamMaskTgtCfi3rdl(d, v, "tgt_cfi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-cfi"] = t
		}
	}

	if v, ok := d.GetOk("tgt_prio"); ok || d.HasChange("tgt_prio") {
		t, err := expandObjectSystemNpuNpuTcamMaskTgtPrio3rdl(d, v, "tgt_prio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-prio"] = t
		}
	}

	if v, ok := d.GetOk("tgt_updt"); ok || d.HasChange("tgt_updt") {
		t, err := expandObjectSystemNpuNpuTcamMaskTgtUpdt3rdl(d, v, "tgt_updt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-updt"] = t
		}
	}

	if v, ok := d.GetOk("tgt_v"); ok || d.HasChange("tgt_v") {
		t, err := expandObjectSystemNpuNpuTcamMaskTgtV3rdl(d, v, "tgt_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgt-v"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok || d.HasChange("tos") {
		t, err := expandObjectSystemNpuNpuTcamMaskTos3rdl(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tp"); ok || d.HasChange("tp") {
		t, err := expandObjectSystemNpuNpuTcamMaskTp3rdl(d, v, "tp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tp"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok || d.HasChange("ttl") {
		t, err := expandObjectSystemNpuNpuTcamMaskTtl3rdl(d, v, "ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("tvid"); ok || d.HasChange("tvid") {
		t, err := expandObjectSystemNpuNpuTcamMaskTvid3rdl(d, v, "tvid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tvid"] = t
		}
	}

	if v, ok := d.GetOk("vdid"); ok || d.HasChange("vdid") {
		t, err := expandObjectSystemNpuNpuTcamMaskVdid3rdl(d, v, "vdid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdid"] = t
		}
	}

	return &obj, nil
}
