// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure NPU TCAM policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcam() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamCreate,
		Read:   resourceObjectSystemNpuNpuTcamRead,
		Update: resourceObjectSystemNpuNpuTcamUpdate,
		Delete: resourceObjectSystemNpuNpuTcamDelete,

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
			"data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"dbg_dump": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"mir_act": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlif": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"oid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pri_act": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"sact": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"act": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"act_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bmproc": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"bmproc_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"df_lif": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"df_lif_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dfr": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dfr_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dmac_skip": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dmac_skip_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dosen": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dosen_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"espff_proc": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"espff_proc_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"etype_pid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"etype_pid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"frag_proc": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"frag_proc_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fwd": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fwd_lif": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fwd_lif_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fwd_tvid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fwd_tvid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fwd_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"icpen": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"icpen_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"igmp_mld_snp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"igmp_mld_snp_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"learn": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"learn_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"m_srh_ctrl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"m_srh_ctrl_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac_id_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mss": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mss_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pleen": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pleen_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prio_pid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prio_pid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"promis": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"promis_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rfsh": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rfsh_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"smac_skip": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"smac_skip_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tp_smchk_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tp_smchk": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tpe_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tpe_id_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdm": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vdm_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vdom_id_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"x_mode": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"x_mode_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tact": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"act": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"act_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fmtuv4_s": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fmtuv4_s_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fmtuv6_s": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fmtuv6_s_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lnkid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lnkid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac_id_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mss_t": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mss_t_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mtuv4": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mtuv4_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mtuv6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mtuv6_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"slif_act": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"slif_act_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sublnkid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sublnkid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tgtv_act": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tgtv_act_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tlif_act": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tlif_act_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tpeid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tpeid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"v6fe": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"v6fe_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vep_en_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vep_slid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"vep_slid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vep_en": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"xlt_lif": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"xlt_lif_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"xlt_vid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"xlt_vid_v": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpuTcamCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuNpuTcam(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpuTcam resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuNpuTcam(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuNpuTcam resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuNpuTcamRead(d, m)
}

func resourceObjectSystemNpuNpuTcamUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpuTcam(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcam resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpuTcam(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSystemNpuNpuTcamRead(d, m)
}

func resourceObjectSystemNpuNpuTcamDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpuTcam(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpuTcam(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcam(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcam resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamData2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := i["df"]; ok {
		result["df"] = flattenObjectSystemNpuNpuTcamDataDf2edl(i["df"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstip"
	if _, ok := i["dstip"]; ok {
		result["dstip"] = flattenObjectSystemNpuNpuTcamDataDstip2edl(i["dstip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstipv6"
	if _, ok := i["dstipv6"]; ok {
		result["dstipv6"] = flattenObjectSystemNpuNpuTcamDataDstipv62edl(i["dstipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstmac"
	if _, ok := i["dstmac"]; ok {
		result["dstmac"] = flattenObjectSystemNpuNpuTcamDataDstmac2edl(i["dstmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstport"
	if _, ok := i["dstport"]; ok {
		result["dstport"] = flattenObjectSystemNpuNpuTcamDataDstport2edl(i["dstport"], d, pre_append)
	}

	pre_append = pre + ".0." + "ethertype"
	if _, ok := i["ethertype"]; ok {
		result["ethertype"] = flattenObjectSystemNpuNpuTcamDataEthertype2edl(i["ethertype"], d, pre_append)
	}

	pre_append = pre + ".0." + "ext_tag"
	if _, ok := i["ext-tag"]; ok {
		result["ext_tag"] = flattenObjectSystemNpuNpuTcamDataExtTag2edl(i["ext-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_off"
	if _, ok := i["frag-off"]; ok {
		result["frag_off"] = flattenObjectSystemNpuNpuTcamDataFragOff2edl(i["frag-off"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := i["gen-buf-cnt"]; ok {
		result["gen_buf_cnt"] = flattenObjectSystemNpuNpuTcamDataGenBufCnt2edl(i["gen-buf-cnt"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_iv"
	if _, ok := i["gen-iv"]; ok {
		result["gen_iv"] = flattenObjectSystemNpuNpuTcamDataGenIv2edl(i["gen-iv"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := i["gen-l3-flags"]; ok {
		result["gen_l3_flags"] = flattenObjectSystemNpuNpuTcamDataGenL3Flags2edl(i["gen-l3-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := i["gen-l4-flags"]; ok {
		result["gen_l4_flags"] = flattenObjectSystemNpuNpuTcamDataGenL4Flags2edl(i["gen-l4-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := i["gen-pkt-ctrl"]; ok {
		result["gen_pkt_ctrl"] = flattenObjectSystemNpuNpuTcamDataGenPktCtrl2edl(i["gen-pkt-ctrl"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri"
	if _, ok := i["gen-pri"]; ok {
		result["gen_pri"] = flattenObjectSystemNpuNpuTcamDataGenPri2edl(i["gen-pri"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := i["gen-pri-v"]; ok {
		result["gen_pri_v"] = flattenObjectSystemNpuNpuTcamDataGenPriV2edl(i["gen-pri-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_tv"
	if _, ok := i["gen-tv"]; ok {
		result["gen_tv"] = flattenObjectSystemNpuNpuTcamDataGenTv2edl(i["gen-tv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ihl"
	if _, ok := i["ihl"]; ok {
		result["ihl"] = flattenObjectSystemNpuNpuTcamDataIhl2edl(i["ihl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip4_id"
	if _, ok := i["ip4-id"]; ok {
		result["ip4_id"] = flattenObjectSystemNpuNpuTcamDataIp4Id2edl(i["ip4-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := i["ip6-fl"]; ok {
		result["ip6_fl"] = flattenObjectSystemNpuNpuTcamDataIp6Fl2edl(i["ip6-fl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipver"
	if _, ok := i["ipver"]; ok {
		result["ipver"] = flattenObjectSystemNpuNpuTcamDataIpver2edl(i["ipver"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := i["l4-wd10"]; ok {
		result["l4_wd10"] = flattenObjectSystemNpuNpuTcamDataL4Wd102edl(i["l4-wd10"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := i["l4-wd11"]; ok {
		result["l4_wd11"] = flattenObjectSystemNpuNpuTcamDataL4Wd112edl(i["l4-wd11"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := i["l4-wd8"]; ok {
		result["l4_wd8"] = flattenObjectSystemNpuNpuTcamDataL4Wd82edl(i["l4-wd8"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := i["l4-wd9"]; ok {
		result["l4_wd9"] = flattenObjectSystemNpuNpuTcamDataL4Wd92edl(i["l4-wd9"], d, pre_append)
	}

	pre_append = pre + ".0." + "mf"
	if _, ok := i["mf"]; ok {
		result["mf"] = flattenObjectSystemNpuNpuTcamDataMf2edl(i["mf"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectSystemNpuNpuTcamDataProtocol2edl(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "slink"
	if _, ok := i["slink"]; ok {
		result["slink"] = flattenObjectSystemNpuNpuTcamDataSlink2edl(i["slink"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_change"
	if _, ok := i["smac-change"]; ok {
		result["smac_change"] = flattenObjectSystemNpuNpuTcamDataSmacChange2edl(i["smac-change"], d, pre_append)
	}

	pre_append = pre + ".0." + "sp"
	if _, ok := i["sp"]; ok {
		result["sp"] = flattenObjectSystemNpuNpuTcamDataSp2edl(i["sp"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_cfi"
	if _, ok := i["src-cfi"]; ok {
		result["src_cfi"] = flattenObjectSystemNpuNpuTcamDataSrcCfi2edl(i["src-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_prio"
	if _, ok := i["src-prio"]; ok {
		result["src_prio"] = flattenObjectSystemNpuNpuTcamDataSrcPrio2edl(i["src-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_updt"
	if _, ok := i["src-updt"]; ok {
		result["src_updt"] = flattenObjectSystemNpuNpuTcamDataSrcUpdt2edl(i["src-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcip"
	if _, ok := i["srcip"]; ok {
		result["srcip"] = flattenObjectSystemNpuNpuTcamDataSrcip2edl(i["srcip"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcipv6"
	if _, ok := i["srcipv6"]; ok {
		result["srcipv6"] = flattenObjectSystemNpuNpuTcamDataSrcipv62edl(i["srcipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcmac"
	if _, ok := i["srcmac"]; ok {
		result["srcmac"] = flattenObjectSystemNpuNpuTcamDataSrcmac2edl(i["srcmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcport"
	if _, ok := i["srcport"]; ok {
		result["srcport"] = flattenObjectSystemNpuNpuTcamDataSrcport2edl(i["srcport"], d, pre_append)
	}

	pre_append = pre + ".0." + "svid"
	if _, ok := i["svid"]; ok {
		result["svid"] = flattenObjectSystemNpuNpuTcamDataSvid2edl(i["svid"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := i["tcp-ack"]; ok {
		result["tcp_ack"] = flattenObjectSystemNpuNpuTcamDataTcpAck2edl(i["tcp-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := i["tcp-cwr"]; ok {
		result["tcp_cwr"] = flattenObjectSystemNpuNpuTcamDataTcpCwr2edl(i["tcp-cwr"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := i["tcp-ece"]; ok {
		result["tcp_ece"] = flattenObjectSystemNpuNpuTcamDataTcpEce2edl(i["tcp-ece"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := i["tcp-fin"]; ok {
		result["tcp_fin"] = flattenObjectSystemNpuNpuTcamDataTcpFin2edl(i["tcp-fin"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_push"
	if _, ok := i["tcp-push"]; ok {
		result["tcp_push"] = flattenObjectSystemNpuNpuTcamDataTcpPush2edl(i["tcp-push"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := i["tcp-rst"]; ok {
		result["tcp_rst"] = flattenObjectSystemNpuNpuTcamDataTcpRst2edl(i["tcp-rst"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := i["tcp-syn"]; ok {
		result["tcp_syn"] = flattenObjectSystemNpuNpuTcamDataTcpSyn2edl(i["tcp-syn"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := i["tcp-urg"]; ok {
		result["tcp_urg"] = flattenObjectSystemNpuNpuTcamDataTcpUrg2edl(i["tcp-urg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := i["tgt-cfi"]; ok {
		result["tgt_cfi"] = flattenObjectSystemNpuNpuTcamDataTgtCfi2edl(i["tgt-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := i["tgt-prio"]; ok {
		result["tgt_prio"] = flattenObjectSystemNpuNpuTcamDataTgtPrio2edl(i["tgt-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := i["tgt-updt"]; ok {
		result["tgt_updt"] = flattenObjectSystemNpuNpuTcamDataTgtUpdt2edl(i["tgt-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_v"
	if _, ok := i["tgt-v"]; ok {
		result["tgt_v"] = flattenObjectSystemNpuNpuTcamDataTgtV2edl(i["tgt-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tos"
	if _, ok := i["tos"]; ok {
		result["tos"] = flattenObjectSystemNpuNpuTcamDataTos2edl(i["tos"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp"
	if _, ok := i["tp"]; ok {
		result["tp"] = flattenObjectSystemNpuNpuTcamDataTp2edl(i["tp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ttl"
	if _, ok := i["ttl"]; ok {
		result["ttl"] = flattenObjectSystemNpuNpuTcamDataTtl2edl(i["ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "tvid"
	if _, ok := i["tvid"]; ok {
		result["tvid"] = flattenObjectSystemNpuNpuTcamDataTvid2edl(i["tvid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdid"
	if _, ok := i["vdid"]; ok {
		result["vdid"] = flattenObjectSystemNpuNpuTcamDataVdid2edl(i["vdid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamDataDf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstipv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstmac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataDstport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataEthertype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataExtTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataFragOff2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenBufCnt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenIv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenL3Flags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenL4Flags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPktCtrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPri2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenPriV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataGenTv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIhl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIp4Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIp6Fl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataIpver2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd102edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd112edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd82edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataL4Wd92edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataMf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSlink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSmacChange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcCfi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcPrio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcUpdt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcipv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcmac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSrcport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataSvid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpCwr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpEce2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpFin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpPush2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpRst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpSyn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTcpUrg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtCfi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtPrio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtUpdt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTgtV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataTvid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDataVdid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamDbgDump2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMask2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := i["df"]; ok {
		result["df"] = flattenObjectSystemNpuNpuTcamMaskDf2edl(i["df"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstip"
	if _, ok := i["dstip"]; ok {
		result["dstip"] = flattenObjectSystemNpuNpuTcamMaskDstip2edl(i["dstip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstipv6"
	if _, ok := i["dstipv6"]; ok {
		result["dstipv6"] = flattenObjectSystemNpuNpuTcamMaskDstipv62edl(i["dstipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstmac"
	if _, ok := i["dstmac"]; ok {
		result["dstmac"] = flattenObjectSystemNpuNpuTcamMaskDstmac2edl(i["dstmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstport"
	if _, ok := i["dstport"]; ok {
		result["dstport"] = flattenObjectSystemNpuNpuTcamMaskDstport2edl(i["dstport"], d, pre_append)
	}

	pre_append = pre + ".0." + "ethertype"
	if _, ok := i["ethertype"]; ok {
		result["ethertype"] = flattenObjectSystemNpuNpuTcamMaskEthertype2edl(i["ethertype"], d, pre_append)
	}

	pre_append = pre + ".0." + "ext_tag"
	if _, ok := i["ext-tag"]; ok {
		result["ext_tag"] = flattenObjectSystemNpuNpuTcamMaskExtTag2edl(i["ext-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_off"
	if _, ok := i["frag-off"]; ok {
		result["frag_off"] = flattenObjectSystemNpuNpuTcamMaskFragOff2edl(i["frag-off"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := i["gen-buf-cnt"]; ok {
		result["gen_buf_cnt"] = flattenObjectSystemNpuNpuTcamMaskGenBufCnt2edl(i["gen-buf-cnt"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_iv"
	if _, ok := i["gen-iv"]; ok {
		result["gen_iv"] = flattenObjectSystemNpuNpuTcamMaskGenIv2edl(i["gen-iv"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := i["gen-l3-flags"]; ok {
		result["gen_l3_flags"] = flattenObjectSystemNpuNpuTcamMaskGenL3Flags2edl(i["gen-l3-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := i["gen-l4-flags"]; ok {
		result["gen_l4_flags"] = flattenObjectSystemNpuNpuTcamMaskGenL4Flags2edl(i["gen-l4-flags"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := i["gen-pkt-ctrl"]; ok {
		result["gen_pkt_ctrl"] = flattenObjectSystemNpuNpuTcamMaskGenPktCtrl2edl(i["gen-pkt-ctrl"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri"
	if _, ok := i["gen-pri"]; ok {
		result["gen_pri"] = flattenObjectSystemNpuNpuTcamMaskGenPri2edl(i["gen-pri"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := i["gen-pri-v"]; ok {
		result["gen_pri_v"] = flattenObjectSystemNpuNpuTcamMaskGenPriV2edl(i["gen-pri-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "gen_tv"
	if _, ok := i["gen-tv"]; ok {
		result["gen_tv"] = flattenObjectSystemNpuNpuTcamMaskGenTv2edl(i["gen-tv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ihl"
	if _, ok := i["ihl"]; ok {
		result["ihl"] = flattenObjectSystemNpuNpuTcamMaskIhl2edl(i["ihl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip4_id"
	if _, ok := i["ip4-id"]; ok {
		result["ip4_id"] = flattenObjectSystemNpuNpuTcamMaskIp4Id2edl(i["ip4-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := i["ip6-fl"]; ok {
		result["ip6_fl"] = flattenObjectSystemNpuNpuTcamMaskIp6Fl2edl(i["ip6-fl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipver"
	if _, ok := i["ipver"]; ok {
		result["ipver"] = flattenObjectSystemNpuNpuTcamMaskIpver2edl(i["ipver"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := i["l4-wd10"]; ok {
		result["l4_wd10"] = flattenObjectSystemNpuNpuTcamMaskL4Wd102edl(i["l4-wd10"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := i["l4-wd11"]; ok {
		result["l4_wd11"] = flattenObjectSystemNpuNpuTcamMaskL4Wd112edl(i["l4-wd11"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := i["l4-wd8"]; ok {
		result["l4_wd8"] = flattenObjectSystemNpuNpuTcamMaskL4Wd82edl(i["l4-wd8"], d, pre_append)
	}

	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := i["l4-wd9"]; ok {
		result["l4_wd9"] = flattenObjectSystemNpuNpuTcamMaskL4Wd92edl(i["l4-wd9"], d, pre_append)
	}

	pre_append = pre + ".0." + "mf"
	if _, ok := i["mf"]; ok {
		result["mf"] = flattenObjectSystemNpuNpuTcamMaskMf2edl(i["mf"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenObjectSystemNpuNpuTcamMaskProtocol2edl(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "slink"
	if _, ok := i["slink"]; ok {
		result["slink"] = flattenObjectSystemNpuNpuTcamMaskSlink2edl(i["slink"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_change"
	if _, ok := i["smac-change"]; ok {
		result["smac_change"] = flattenObjectSystemNpuNpuTcamMaskSmacChange2edl(i["smac-change"], d, pre_append)
	}

	pre_append = pre + ".0." + "sp"
	if _, ok := i["sp"]; ok {
		result["sp"] = flattenObjectSystemNpuNpuTcamMaskSp2edl(i["sp"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_cfi"
	if _, ok := i["src-cfi"]; ok {
		result["src_cfi"] = flattenObjectSystemNpuNpuTcamMaskSrcCfi2edl(i["src-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_prio"
	if _, ok := i["src-prio"]; ok {
		result["src_prio"] = flattenObjectSystemNpuNpuTcamMaskSrcPrio2edl(i["src-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_updt"
	if _, ok := i["src-updt"]; ok {
		result["src_updt"] = flattenObjectSystemNpuNpuTcamMaskSrcUpdt2edl(i["src-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcip"
	if _, ok := i["srcip"]; ok {
		result["srcip"] = flattenObjectSystemNpuNpuTcamMaskSrcip2edl(i["srcip"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcipv6"
	if _, ok := i["srcipv6"]; ok {
		result["srcipv6"] = flattenObjectSystemNpuNpuTcamMaskSrcipv62edl(i["srcipv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcmac"
	if _, ok := i["srcmac"]; ok {
		result["srcmac"] = flattenObjectSystemNpuNpuTcamMaskSrcmac2edl(i["srcmac"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcport"
	if _, ok := i["srcport"]; ok {
		result["srcport"] = flattenObjectSystemNpuNpuTcamMaskSrcport2edl(i["srcport"], d, pre_append)
	}

	pre_append = pre + ".0." + "svid"
	if _, ok := i["svid"]; ok {
		result["svid"] = flattenObjectSystemNpuNpuTcamMaskSvid2edl(i["svid"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := i["tcp-ack"]; ok {
		result["tcp_ack"] = flattenObjectSystemNpuNpuTcamMaskTcpAck2edl(i["tcp-ack"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := i["tcp-cwr"]; ok {
		result["tcp_cwr"] = flattenObjectSystemNpuNpuTcamMaskTcpCwr2edl(i["tcp-cwr"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := i["tcp-ece"]; ok {
		result["tcp_ece"] = flattenObjectSystemNpuNpuTcamMaskTcpEce2edl(i["tcp-ece"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := i["tcp-fin"]; ok {
		result["tcp_fin"] = flattenObjectSystemNpuNpuTcamMaskTcpFin2edl(i["tcp-fin"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_push"
	if _, ok := i["tcp-push"]; ok {
		result["tcp_push"] = flattenObjectSystemNpuNpuTcamMaskTcpPush2edl(i["tcp-push"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := i["tcp-rst"]; ok {
		result["tcp_rst"] = flattenObjectSystemNpuNpuTcamMaskTcpRst2edl(i["tcp-rst"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := i["tcp-syn"]; ok {
		result["tcp_syn"] = flattenObjectSystemNpuNpuTcamMaskTcpSyn2edl(i["tcp-syn"], d, pre_append)
	}

	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := i["tcp-urg"]; ok {
		result["tcp_urg"] = flattenObjectSystemNpuNpuTcamMaskTcpUrg2edl(i["tcp-urg"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := i["tgt-cfi"]; ok {
		result["tgt_cfi"] = flattenObjectSystemNpuNpuTcamMaskTgtCfi2edl(i["tgt-cfi"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := i["tgt-prio"]; ok {
		result["tgt_prio"] = flattenObjectSystemNpuNpuTcamMaskTgtPrio2edl(i["tgt-prio"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := i["tgt-updt"]; ok {
		result["tgt_updt"] = flattenObjectSystemNpuNpuTcamMaskTgtUpdt2edl(i["tgt-updt"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgt_v"
	if _, ok := i["tgt-v"]; ok {
		result["tgt_v"] = flattenObjectSystemNpuNpuTcamMaskTgtV2edl(i["tgt-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tos"
	if _, ok := i["tos"]; ok {
		result["tos"] = flattenObjectSystemNpuNpuTcamMaskTos2edl(i["tos"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp"
	if _, ok := i["tp"]; ok {
		result["tp"] = flattenObjectSystemNpuNpuTcamMaskTp2edl(i["tp"], d, pre_append)
	}

	pre_append = pre + ".0." + "ttl"
	if _, ok := i["ttl"]; ok {
		result["ttl"] = flattenObjectSystemNpuNpuTcamMaskTtl2edl(i["ttl"], d, pre_append)
	}

	pre_append = pre + ".0." + "tvid"
	if _, ok := i["tvid"]; ok {
		result["tvid"] = flattenObjectSystemNpuNpuTcamMaskTvid2edl(i["tvid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdid"
	if _, ok := i["vdid"]; ok {
		result["vdid"] = flattenObjectSystemNpuNpuTcamMaskVdid2edl(i["vdid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamMaskDf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstipv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstmac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskDstport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskEthertype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskExtTag2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskFragOff2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenBufCnt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenIv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenL3Flags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenL4Flags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPktCtrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPri2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenPriV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskGenTv2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIhl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIp4Id2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIp6Fl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskIpver2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd102edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd112edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd82edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskL4Wd92edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskMf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSlink2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSmacChange2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcCfi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcPrio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcUpdt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcipv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcmac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSrcport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskSvid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpAck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpCwr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpEce2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpFin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpPush2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpRst2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpSyn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTcpUrg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtCfi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtPrio2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtUpdt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTgtV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskTvid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMaskVdid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamMirAct2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlif"
	if _, ok := i["vlif"]; ok {
		result["vlif"] = flattenObjectSystemNpuNpuTcamMirActVlif2edl(i["vlif"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamMirActVlif2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamOid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamPriAct2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenObjectSystemNpuNpuTcamPriActPriority2edl(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "weight"
	if _, ok := i["weight"]; ok {
		result["weight"] = flattenObjectSystemNpuNpuTcamPriActWeight2edl(i["weight"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamPriActPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamPriActWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSact2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := i["act"]; ok {
		result["act"] = flattenObjectSystemNpuNpuTcamSactAct2edl(i["act"], d, pre_append)
	}

	pre_append = pre + ".0." + "act_v"
	if _, ok := i["act-v"]; ok {
		result["act_v"] = flattenObjectSystemNpuNpuTcamSactActV2edl(i["act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "bmproc"
	if _, ok := i["bmproc"]; ok {
		result["bmproc"] = flattenObjectSystemNpuNpuTcamSactBmproc2edl(i["bmproc"], d, pre_append)
	}

	pre_append = pre + ".0." + "bmproc_v"
	if _, ok := i["bmproc-v"]; ok {
		result["bmproc_v"] = flattenObjectSystemNpuNpuTcamSactBmprocV2edl(i["bmproc-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "df_lif"
	if _, ok := i["df-lif"]; ok {
		result["df_lif"] = flattenObjectSystemNpuNpuTcamSactDfLif2edl(i["df-lif"], d, pre_append)
	}

	pre_append = pre + ".0." + "df_lif_v"
	if _, ok := i["df-lif-v"]; ok {
		result["df_lif_v"] = flattenObjectSystemNpuNpuTcamSactDfLifV2edl(i["df-lif-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "dfr"
	if _, ok := i["dfr"]; ok {
		result["dfr"] = flattenObjectSystemNpuNpuTcamSactDfr2edl(i["dfr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dfr_v"
	if _, ok := i["dfr-v"]; ok {
		result["dfr_v"] = flattenObjectSystemNpuNpuTcamSactDfrV2edl(i["dfr-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "dmac_skip"
	if _, ok := i["dmac-skip"]; ok {
		result["dmac_skip"] = flattenObjectSystemNpuNpuTcamSactDmacSkip2edl(i["dmac-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dmac_skip_v"
	if _, ok := i["dmac-skip-v"]; ok {
		result["dmac_skip_v"] = flattenObjectSystemNpuNpuTcamSactDmacSkipV2edl(i["dmac-skip-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "dosen"
	if _, ok := i["dosen"]; ok {
		result["dosen"] = flattenObjectSystemNpuNpuTcamSactDosen2edl(i["dosen"], d, pre_append)
	}

	pre_append = pre + ".0." + "dosen_v"
	if _, ok := i["dosen-v"]; ok {
		result["dosen_v"] = flattenObjectSystemNpuNpuTcamSactDosenV2edl(i["dosen-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "espff_proc"
	if _, ok := i["espff-proc"]; ok {
		result["espff_proc"] = flattenObjectSystemNpuNpuTcamSactEspffProc2edl(i["espff-proc"], d, pre_append)
	}

	pre_append = pre + ".0." + "espff_proc_v"
	if _, ok := i["espff-proc-v"]; ok {
		result["espff_proc_v"] = flattenObjectSystemNpuNpuTcamSactEspffProcV2edl(i["espff-proc-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "etype_pid"
	if _, ok := i["etype-pid"]; ok {
		result["etype_pid"] = flattenObjectSystemNpuNpuTcamSactEtypePid2edl(i["etype-pid"], d, pre_append)
	}

	pre_append = pre + ".0." + "etype_pid_v"
	if _, ok := i["etype-pid-v"]; ok {
		result["etype_pid_v"] = flattenObjectSystemNpuNpuTcamSactEtypePidV2edl(i["etype-pid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_proc"
	if _, ok := i["frag-proc"]; ok {
		result["frag_proc"] = flattenObjectSystemNpuNpuTcamSactFragProc2edl(i["frag-proc"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_proc_v"
	if _, ok := i["frag-proc-v"]; ok {
		result["frag_proc_v"] = flattenObjectSystemNpuNpuTcamSactFragProcV2edl(i["frag-proc-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd"
	if _, ok := i["fwd"]; ok {
		result["fwd"] = flattenObjectSystemNpuNpuTcamSactFwd2edl(i["fwd"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_lif"
	if _, ok := i["fwd-lif"]; ok {
		result["fwd_lif"] = flattenObjectSystemNpuNpuTcamSactFwdLif2edl(i["fwd-lif"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_lif_v"
	if _, ok := i["fwd-lif-v"]; ok {
		result["fwd_lif_v"] = flattenObjectSystemNpuNpuTcamSactFwdLifV2edl(i["fwd-lif-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_tvid"
	if _, ok := i["fwd-tvid"]; ok {
		result["fwd_tvid"] = flattenObjectSystemNpuNpuTcamSactFwdTvid2edl(i["fwd-tvid"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_tvid_v"
	if _, ok := i["fwd-tvid-v"]; ok {
		result["fwd_tvid_v"] = flattenObjectSystemNpuNpuTcamSactFwdTvidV2edl(i["fwd-tvid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fwd_v"
	if _, ok := i["fwd-v"]; ok {
		result["fwd_v"] = flattenObjectSystemNpuNpuTcamSactFwdV2edl(i["fwd-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "icpen"
	if _, ok := i["icpen"]; ok {
		result["icpen"] = flattenObjectSystemNpuNpuTcamSactIcpen2edl(i["icpen"], d, pre_append)
	}

	pre_append = pre + ".0." + "icpen_v"
	if _, ok := i["icpen-v"]; ok {
		result["icpen_v"] = flattenObjectSystemNpuNpuTcamSactIcpenV2edl(i["icpen-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "igmp_mld_snp"
	if _, ok := i["igmp-mld-snp"]; ok {
		result["igmp_mld_snp"] = flattenObjectSystemNpuNpuTcamSactIgmpMldSnp2edl(i["igmp-mld-snp"], d, pre_append)
	}

	pre_append = pre + ".0." + "igmp_mld_snp_v"
	if _, ok := i["igmp-mld-snp-v"]; ok {
		result["igmp_mld_snp_v"] = flattenObjectSystemNpuNpuTcamSactIgmpMldSnpV2edl(i["igmp-mld-snp-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "learn"
	if _, ok := i["learn"]; ok {
		result["learn"] = flattenObjectSystemNpuNpuTcamSactLearn2edl(i["learn"], d, pre_append)
	}

	pre_append = pre + ".0." + "learn_v"
	if _, ok := i["learn-v"]; ok {
		result["learn_v"] = flattenObjectSystemNpuNpuTcamSactLearnV2edl(i["learn-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "m_srh_ctrl"
	if _, ok := i["m-srh-ctrl"]; ok {
		result["m_srh_ctrl"] = flattenObjectSystemNpuNpuTcamSactMSrhCtrl2edl(i["m-srh-ctrl"], d, pre_append)
	}

	pre_append = pre + ".0." + "m_srh_ctrl_v"
	if _, ok := i["m-srh-ctrl-v"]; ok {
		result["m_srh_ctrl_v"] = flattenObjectSystemNpuNpuTcamSactMSrhCtrlV2edl(i["m-srh-ctrl-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id"
	if _, ok := i["mac-id"]; ok {
		result["mac_id"] = flattenObjectSystemNpuNpuTcamSactMacId2edl(i["mac-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := i["mac-id-v"]; ok {
		result["mac_id_v"] = flattenObjectSystemNpuNpuTcamSactMacIdV2edl(i["mac-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss"
	if _, ok := i["mss"]; ok {
		result["mss"] = flattenObjectSystemNpuNpuTcamSactMss2edl(i["mss"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss_v"
	if _, ok := i["mss-v"]; ok {
		result["mss_v"] = flattenObjectSystemNpuNpuTcamSactMssV2edl(i["mss-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "pleen"
	if _, ok := i["pleen"]; ok {
		result["pleen"] = flattenObjectSystemNpuNpuTcamSactPleen2edl(i["pleen"], d, pre_append)
	}

	pre_append = pre + ".0." + "pleen_v"
	if _, ok := i["pleen-v"]; ok {
		result["pleen_v"] = flattenObjectSystemNpuNpuTcamSactPleenV2edl(i["pleen-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "prio_pid"
	if _, ok := i["prio-pid"]; ok {
		result["prio_pid"] = flattenObjectSystemNpuNpuTcamSactPrioPid2edl(i["prio-pid"], d, pre_append)
	}

	pre_append = pre + ".0." + "prio_pid_v"
	if _, ok := i["prio-pid-v"]; ok {
		result["prio_pid_v"] = flattenObjectSystemNpuNpuTcamSactPrioPidV2edl(i["prio-pid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "promis"
	if _, ok := i["promis"]; ok {
		result["promis"] = flattenObjectSystemNpuNpuTcamSactPromis2edl(i["promis"], d, pre_append)
	}

	pre_append = pre + ".0." + "promis_v"
	if _, ok := i["promis-v"]; ok {
		result["promis_v"] = flattenObjectSystemNpuNpuTcamSactPromisV2edl(i["promis-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "rfsh"
	if _, ok := i["rfsh"]; ok {
		result["rfsh"] = flattenObjectSystemNpuNpuTcamSactRfsh2edl(i["rfsh"], d, pre_append)
	}

	pre_append = pre + ".0." + "rfsh_v"
	if _, ok := i["rfsh-v"]; ok {
		result["rfsh_v"] = flattenObjectSystemNpuNpuTcamSactRfshV2edl(i["rfsh-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_skip"
	if _, ok := i["smac-skip"]; ok {
		result["smac_skip"] = flattenObjectSystemNpuNpuTcamSactSmacSkip2edl(i["smac-skip"], d, pre_append)
	}

	pre_append = pre + ".0." + "smac_skip_v"
	if _, ok := i["smac-skip-v"]; ok {
		result["smac_skip_v"] = flattenObjectSystemNpuNpuTcamSactSmacSkipV2edl(i["smac-skip-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp_smchk_v"
	if _, ok := i["tp-smchk-v"]; ok {
		result["tp_smchk_v"] = flattenObjectSystemNpuNpuTcamSactTpSmchkV2edl(i["tp-smchk-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tp_smchk"
	if _, ok := i["tp_smchk"]; ok {
		result["tp_smchk"] = flattenObjectSystemNpuNpuTcamSactTpSmchk2edl(i["tp_smchk"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpe_id"
	if _, ok := i["tpe-id"]; ok {
		result["tpe_id"] = flattenObjectSystemNpuNpuTcamSactTpeId2edl(i["tpe-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpe_id_v"
	if _, ok := i["tpe-id-v"]; ok {
		result["tpe_id_v"] = flattenObjectSystemNpuNpuTcamSactTpeIdV2edl(i["tpe-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdm"
	if _, ok := i["vdm"]; ok {
		result["vdm"] = flattenObjectSystemNpuNpuTcamSactVdm2edl(i["vdm"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdm_v"
	if _, ok := i["vdm-v"]; ok {
		result["vdm_v"] = flattenObjectSystemNpuNpuTcamSactVdmV2edl(i["vdm-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom_id"
	if _, ok := i["vdom-id"]; ok {
		result["vdom_id"] = flattenObjectSystemNpuNpuTcamSactVdomId2edl(i["vdom-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom_id_v"
	if _, ok := i["vdom-id-v"]; ok {
		result["vdom_id_v"] = flattenObjectSystemNpuNpuTcamSactVdomIdV2edl(i["vdom-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "x_mode"
	if _, ok := i["x-mode"]; ok {
		result["x_mode"] = flattenObjectSystemNpuNpuTcamSactXMode2edl(i["x-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "x_mode_v"
	if _, ok := i["x-mode-v"]; ok {
		result["x_mode_v"] = flattenObjectSystemNpuNpuTcamSactXModeV2edl(i["x-mode-v"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamSactAct2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactActV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactBmproc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactBmprocV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfLif2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfLifV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfrV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDmacSkip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDmacSkipV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDosen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDosenV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEspffProc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEspffProcV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEtypePid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEtypePidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFragProc2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFragProcV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdLif2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdLifV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdTvid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdTvidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIcpen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIcpenV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIgmpMldSnp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIgmpMldSnpV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactLearn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactLearnV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMSrhCtrl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMSrhCtrlV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMacId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMacIdV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMssV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPleen2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPleenV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPrioPid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPrioPidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPromis2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPromisV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactRfsh2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactRfshV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactSmacSkip2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactSmacSkipV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpSmchkV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpSmchk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpeId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpeIdV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdmV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdomId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdomIdV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactXMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactXModeV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTact2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := i["act"]; ok {
		result["act"] = flattenObjectSystemNpuNpuTcamTactAct2edl(i["act"], d, pre_append)
	}

	pre_append = pre + ".0." + "act_v"
	if _, ok := i["act-v"]; ok {
		result["act_v"] = flattenObjectSystemNpuNpuTcamTactActV2edl(i["act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv4_s"
	if _, ok := i["fmtuv4-s"]; ok {
		result["fmtuv4_s"] = flattenObjectSystemNpuNpuTcamTactFmtuv4S2edl(i["fmtuv4-s"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv4_s_v"
	if _, ok := i["fmtuv4-s-v"]; ok {
		result["fmtuv4_s_v"] = flattenObjectSystemNpuNpuTcamTactFmtuv4SV2edl(i["fmtuv4-s-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv6_s"
	if _, ok := i["fmtuv6-s"]; ok {
		result["fmtuv6_s"] = flattenObjectSystemNpuNpuTcamTactFmtuv6S2edl(i["fmtuv6-s"], d, pre_append)
	}

	pre_append = pre + ".0." + "fmtuv6_s_v"
	if _, ok := i["fmtuv6-s-v"]; ok {
		result["fmtuv6_s_v"] = flattenObjectSystemNpuNpuTcamTactFmtuv6SV2edl(i["fmtuv6-s-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "lnkid"
	if _, ok := i["lnkid"]; ok {
		result["lnkid"] = flattenObjectSystemNpuNpuTcamTactLnkid2edl(i["lnkid"], d, pre_append)
	}

	pre_append = pre + ".0." + "lnkid_v"
	if _, ok := i["lnkid-v"]; ok {
		result["lnkid_v"] = flattenObjectSystemNpuNpuTcamTactLnkidV2edl(i["lnkid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id"
	if _, ok := i["mac-id"]; ok {
		result["mac_id"] = flattenObjectSystemNpuNpuTcamTactMacId2edl(i["mac-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := i["mac-id-v"]; ok {
		result["mac_id_v"] = flattenObjectSystemNpuNpuTcamTactMacIdV2edl(i["mac-id-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss_t"
	if _, ok := i["mss-t"]; ok {
		result["mss_t"] = flattenObjectSystemNpuNpuTcamTactMssT2edl(i["mss-t"], d, pre_append)
	}

	pre_append = pre + ".0." + "mss_t_v"
	if _, ok := i["mss-t-v"]; ok {
		result["mss_t_v"] = flattenObjectSystemNpuNpuTcamTactMssTV2edl(i["mss-t-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv4"
	if _, ok := i["mtuv4"]; ok {
		result["mtuv4"] = flattenObjectSystemNpuNpuTcamTactMtuv42edl(i["mtuv4"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv4_v"
	if _, ok := i["mtuv4-v"]; ok {
		result["mtuv4_v"] = flattenObjectSystemNpuNpuTcamTactMtuv4V2edl(i["mtuv4-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv6"
	if _, ok := i["mtuv6"]; ok {
		result["mtuv6"] = flattenObjectSystemNpuNpuTcamTactMtuv62edl(i["mtuv6"], d, pre_append)
	}

	pre_append = pre + ".0." + "mtuv6_v"
	if _, ok := i["mtuv6-v"]; ok {
		result["mtuv6_v"] = flattenObjectSystemNpuNpuTcamTactMtuv6V2edl(i["mtuv6-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "slif_act"
	if _, ok := i["slif-act"]; ok {
		result["slif_act"] = flattenObjectSystemNpuNpuTcamTactSlifAct2edl(i["slif-act"], d, pre_append)
	}

	pre_append = pre + ".0." + "slif_act_v"
	if _, ok := i["slif-act-v"]; ok {
		result["slif_act_v"] = flattenObjectSystemNpuNpuTcamTactSlifActV2edl(i["slif-act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "sublnkid"
	if _, ok := i["sublnkid"]; ok {
		result["sublnkid"] = flattenObjectSystemNpuNpuTcamTactSublnkid2edl(i["sublnkid"], d, pre_append)
	}

	pre_append = pre + ".0." + "sublnkid_v"
	if _, ok := i["sublnkid-v"]; ok {
		result["sublnkid_v"] = flattenObjectSystemNpuNpuTcamTactSublnkidV2edl(i["sublnkid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgtv_act"
	if _, ok := i["tgtv-act"]; ok {
		result["tgtv_act"] = flattenObjectSystemNpuNpuTcamTactTgtvAct2edl(i["tgtv-act"], d, pre_append)
	}

	pre_append = pre + ".0." + "tgtv_act_v"
	if _, ok := i["tgtv-act-v"]; ok {
		result["tgtv_act_v"] = flattenObjectSystemNpuNpuTcamTactTgtvActV2edl(i["tgtv-act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tlif_act"
	if _, ok := i["tlif-act"]; ok {
		result["tlif_act"] = flattenObjectSystemNpuNpuTcamTactTlifAct2edl(i["tlif-act"], d, pre_append)
	}

	pre_append = pre + ".0." + "tlif_act_v"
	if _, ok := i["tlif-act-v"]; ok {
		result["tlif_act_v"] = flattenObjectSystemNpuNpuTcamTactTlifActV2edl(i["tlif-act-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpeid"
	if _, ok := i["tpeid"]; ok {
		result["tpeid"] = flattenObjectSystemNpuNpuTcamTactTpeid2edl(i["tpeid"], d, pre_append)
	}

	pre_append = pre + ".0." + "tpeid_v"
	if _, ok := i["tpeid-v"]; ok {
		result["tpeid_v"] = flattenObjectSystemNpuNpuTcamTactTpeidV2edl(i["tpeid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "v6fe"
	if _, ok := i["v6fe"]; ok {
		result["v6fe"] = flattenObjectSystemNpuNpuTcamTactV6Fe2edl(i["v6fe"], d, pre_append)
	}

	pre_append = pre + ".0." + "v6fe_v"
	if _, ok := i["v6fe-v"]; ok {
		result["v6fe_v"] = flattenObjectSystemNpuNpuTcamTactV6FeV2edl(i["v6fe-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_en_v"
	if _, ok := i["vep-en-v"]; ok {
		result["vep_en_v"] = flattenObjectSystemNpuNpuTcamTactVepEnV2edl(i["vep-en-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_slid"
	if _, ok := i["vep-slid"]; ok {
		result["vep_slid"] = flattenObjectSystemNpuNpuTcamTactVepSlid2edl(i["vep-slid"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_slid_v"
	if _, ok := i["vep-slid-v"]; ok {
		result["vep_slid_v"] = flattenObjectSystemNpuNpuTcamTactVepSlidV2edl(i["vep-slid-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "vep_en"
	if _, ok := i["vep_en"]; ok {
		result["vep_en"] = flattenObjectSystemNpuNpuTcamTactVepEn2edl(i["vep_en"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_lif"
	if _, ok := i["xlt-lif"]; ok {
		result["xlt_lif"] = flattenObjectSystemNpuNpuTcamTactXltLif2edl(i["xlt-lif"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_lif_v"
	if _, ok := i["xlt-lif-v"]; ok {
		result["xlt_lif_v"] = flattenObjectSystemNpuNpuTcamTactXltLifV2edl(i["xlt-lif-v"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_vid"
	if _, ok := i["xlt-vid"]; ok {
		result["xlt_vid"] = flattenObjectSystemNpuNpuTcamTactXltVid2edl(i["xlt-vid"], d, pre_append)
	}

	pre_append = pre + ".0." + "xlt_vid_v"
	if _, ok := i["xlt-vid-v"]; ok {
		result["xlt_vid_v"] = flattenObjectSystemNpuNpuTcamTactXltVidV2edl(i["xlt-vid-v"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectSystemNpuNpuTcamTactAct2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactActV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv4S2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv4SV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv6S2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv6SV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactLnkid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactLnkidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMacId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMacIdV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMssT2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMssTV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv4V2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv6V2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSlifAct2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSlifActV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSublnkid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSublnkidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTgtvAct2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTgtvActV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTlifAct2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTlifActV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTpeid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTpeidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactV6Fe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactV6FeV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepEnV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepSlid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepSlidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepEn2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltLif2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltLifV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltVid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltVidV2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamVid2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcam(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if isImportTable() {
		if err = d.Set("data", flattenObjectSystemNpuNpuTcamData2edl(o["data"], d, "data")); err != nil {
			if vv, ok := fortiAPIPatch(o["data"], "ObjectSystemNpuNpuTcam-Data"); ok {
				if err = d.Set("data", vv); err != nil {
					return fmt.Errorf("Error reading data: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("data"); ok {
			if err = d.Set("data", flattenObjectSystemNpuNpuTcamData2edl(o["data"], d, "data")); err != nil {
				if vv, ok := fortiAPIPatch(o["data"], "ObjectSystemNpuNpuTcam-Data"); ok {
					if err = d.Set("data", vv); err != nil {
						return fmt.Errorf("Error reading data: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading data: %v", err)
				}
			}
		}
	}

	if err = d.Set("dbg_dump", flattenObjectSystemNpuNpuTcamDbgDump2edl(o["dbg-dump"], d, "dbg_dump")); err != nil {
		if vv, ok := fortiAPIPatch(o["dbg-dump"], "ObjectSystemNpuNpuTcam-DbgDump"); ok {
			if err = d.Set("dbg_dump", vv); err != nil {
				return fmt.Errorf("Error reading dbg_dump: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dbg_dump: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mask", flattenObjectSystemNpuNpuTcamMask2edl(o["mask"], d, "mask")); err != nil {
			if vv, ok := fortiAPIPatch(o["mask"], "ObjectSystemNpuNpuTcam-Mask"); ok {
				if err = d.Set("mask", vv); err != nil {
					return fmt.Errorf("Error reading mask: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mask: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mask"); ok {
			if err = d.Set("mask", flattenObjectSystemNpuNpuTcamMask2edl(o["mask"], d, "mask")); err != nil {
				if vv, ok := fortiAPIPatch(o["mask"], "ObjectSystemNpuNpuTcam-Mask"); ok {
					if err = d.Set("mask", vv); err != nil {
						return fmt.Errorf("Error reading mask: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mask: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mir_act", flattenObjectSystemNpuNpuTcamMirAct2edl(o["mir-act"], d, "mir_act")); err != nil {
			if vv, ok := fortiAPIPatch(o["mir-act"], "ObjectSystemNpuNpuTcam-MirAct"); ok {
				if err = d.Set("mir_act", vv); err != nil {
					return fmt.Errorf("Error reading mir_act: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mir_act: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mir_act"); ok {
			if err = d.Set("mir_act", flattenObjectSystemNpuNpuTcamMirAct2edl(o["mir-act"], d, "mir_act")); err != nil {
				if vv, ok := fortiAPIPatch(o["mir-act"], "ObjectSystemNpuNpuTcam-MirAct"); ok {
					if err = d.Set("mir_act", vv); err != nil {
						return fmt.Errorf("Error reading mir_act: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mir_act: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectSystemNpuNpuTcamName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSystemNpuNpuTcam-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("oid", flattenObjectSystemNpuNpuTcamOid2edl(o["oid"], d, "oid")); err != nil {
		if vv, ok := fortiAPIPatch(o["oid"], "ObjectSystemNpuNpuTcam-Oid"); ok {
			if err = d.Set("oid", vv); err != nil {
				return fmt.Errorf("Error reading oid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pri_act", flattenObjectSystemNpuNpuTcamPriAct2edl(o["pri-act"], d, "pri_act")); err != nil {
			if vv, ok := fortiAPIPatch(o["pri-act"], "ObjectSystemNpuNpuTcam-PriAct"); ok {
				if err = d.Set("pri_act", vv); err != nil {
					return fmt.Errorf("Error reading pri_act: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pri_act: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pri_act"); ok {
			if err = d.Set("pri_act", flattenObjectSystemNpuNpuTcamPriAct2edl(o["pri-act"], d, "pri_act")); err != nil {
				if vv, ok := fortiAPIPatch(o["pri-act"], "ObjectSystemNpuNpuTcam-PriAct"); ok {
					if err = d.Set("pri_act", vv); err != nil {
						return fmt.Errorf("Error reading pri_act: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pri_act: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sact", flattenObjectSystemNpuNpuTcamSact2edl(o["sact"], d, "sact")); err != nil {
			if vv, ok := fortiAPIPatch(o["sact"], "ObjectSystemNpuNpuTcam-Sact"); ok {
				if err = d.Set("sact", vv); err != nil {
					return fmt.Errorf("Error reading sact: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sact: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sact"); ok {
			if err = d.Set("sact", flattenObjectSystemNpuNpuTcamSact2edl(o["sact"], d, "sact")); err != nil {
				if vv, ok := fortiAPIPatch(o["sact"], "ObjectSystemNpuNpuTcam-Sact"); ok {
					if err = d.Set("sact", vv); err != nil {
						return fmt.Errorf("Error reading sact: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sact: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("tact", flattenObjectSystemNpuNpuTcamTact2edl(o["tact"], d, "tact")); err != nil {
			if vv, ok := fortiAPIPatch(o["tact"], "ObjectSystemNpuNpuTcam-Tact"); ok {
				if err = d.Set("tact", vv); err != nil {
					return fmt.Errorf("Error reading tact: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading tact: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tact"); ok {
			if err = d.Set("tact", flattenObjectSystemNpuNpuTcamTact2edl(o["tact"], d, "tact")); err != nil {
				if vv, ok := fortiAPIPatch(o["tact"], "ObjectSystemNpuNpuTcam-Tact"); ok {
					if err = d.Set("tact", vv); err != nil {
						return fmt.Errorf("Error reading tact: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading tact: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenObjectSystemNpuNpuTcamType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectSystemNpuNpuTcam-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vid", flattenObjectSystemNpuNpuTcamVid2edl(o["vid"], d, "vid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vid"], "ObjectSystemNpuNpuTcam-Vid"); ok {
			if err = d.Set("vid", vv); err != nil {
				return fmt.Errorf("Error reading vid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vid: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamData2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df"], _ = expandObjectSystemNpuNpuTcamDataDf2edl(d, i["df"], pre_append)
	}
	pre_append = pre + ".0." + "dstip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstip"], _ = expandObjectSystemNpuNpuTcamDataDstip2edl(d, i["dstip"], pre_append)
	}
	pre_append = pre + ".0." + "dstipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstipv6"], _ = expandObjectSystemNpuNpuTcamDataDstipv62edl(d, i["dstipv6"], pre_append)
	}
	pre_append = pre + ".0." + "dstmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstmac"], _ = expandObjectSystemNpuNpuTcamDataDstmac2edl(d, i["dstmac"], pre_append)
	}
	pre_append = pre + ".0." + "dstport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstport"], _ = expandObjectSystemNpuNpuTcamDataDstport2edl(d, i["dstport"], pre_append)
	}
	pre_append = pre + ".0." + "ethertype"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ethertype"], _ = expandObjectSystemNpuNpuTcamDataEthertype2edl(d, i["ethertype"], pre_append)
	}
	pre_append = pre + ".0." + "ext_tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ext-tag"], _ = expandObjectSystemNpuNpuTcamDataExtTag2edl(d, i["ext_tag"], pre_append)
	}
	pre_append = pre + ".0." + "frag_off"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-off"], _ = expandObjectSystemNpuNpuTcamDataFragOff2edl(d, i["frag_off"], pre_append)
	}
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-buf-cnt"], _ = expandObjectSystemNpuNpuTcamDataGenBufCnt2edl(d, i["gen_buf_cnt"], pre_append)
	}
	pre_append = pre + ".0." + "gen_iv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-iv"], _ = expandObjectSystemNpuNpuTcamDataGenIv2edl(d, i["gen_iv"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l3-flags"], _ = expandObjectSystemNpuNpuTcamDataGenL3Flags2edl(d, i["gen_l3_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l4-flags"], _ = expandObjectSystemNpuNpuTcamDataGenL4Flags2edl(d, i["gen_l4_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pkt-ctrl"], _ = expandObjectSystemNpuNpuTcamDataGenPktCtrl2edl(d, i["gen_pkt_ctrl"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri"], _ = expandObjectSystemNpuNpuTcamDataGenPri2edl(d, i["gen_pri"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri-v"], _ = expandObjectSystemNpuNpuTcamDataGenPriV2edl(d, i["gen_pri_v"], pre_append)
	}
	pre_append = pre + ".0." + "gen_tv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-tv"], _ = expandObjectSystemNpuNpuTcamDataGenTv2edl(d, i["gen_tv"], pre_append)
	}
	pre_append = pre + ".0." + "ihl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ihl"], _ = expandObjectSystemNpuNpuTcamDataIhl2edl(d, i["ihl"], pre_append)
	}
	pre_append = pre + ".0." + "ip4_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip4-id"], _ = expandObjectSystemNpuNpuTcamDataIp4Id2edl(d, i["ip4_id"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-fl"], _ = expandObjectSystemNpuNpuTcamDataIp6Fl2edl(d, i["ip6_fl"], pre_append)
	}
	pre_append = pre + ".0." + "ipver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipver"], _ = expandObjectSystemNpuNpuTcamDataIpver2edl(d, i["ipver"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd10"], _ = expandObjectSystemNpuNpuTcamDataL4Wd102edl(d, i["l4_wd10"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd11"], _ = expandObjectSystemNpuNpuTcamDataL4Wd112edl(d, i["l4_wd11"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd8"], _ = expandObjectSystemNpuNpuTcamDataL4Wd82edl(d, i["l4_wd8"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd9"], _ = expandObjectSystemNpuNpuTcamDataL4Wd92edl(d, i["l4_wd9"], pre_append)
	}
	pre_append = pre + ".0." + "mf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mf"], _ = expandObjectSystemNpuNpuTcamDataMf2edl(d, i["mf"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectSystemNpuNpuTcamDataProtocol2edl(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "slink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slink"], _ = expandObjectSystemNpuNpuTcamDataSlink2edl(d, i["slink"], pre_append)
	}
	pre_append = pre + ".0." + "smac_change"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-change"], _ = expandObjectSystemNpuNpuTcamDataSmacChange2edl(d, i["smac_change"], pre_append)
	}
	pre_append = pre + ".0." + "sp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sp"], _ = expandObjectSystemNpuNpuTcamDataSp2edl(d, i["sp"], pre_append)
	}
	pre_append = pre + ".0." + "src_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-cfi"], _ = expandObjectSystemNpuNpuTcamDataSrcCfi2edl(d, i["src_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "src_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-prio"], _ = expandObjectSystemNpuNpuTcamDataSrcPrio2edl(d, i["src_prio"], pre_append)
	}
	pre_append = pre + ".0." + "src_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-updt"], _ = expandObjectSystemNpuNpuTcamDataSrcUpdt2edl(d, i["src_updt"], pre_append)
	}
	pre_append = pre + ".0." + "srcip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcip"], _ = expandObjectSystemNpuNpuTcamDataSrcip2edl(d, i["srcip"], pre_append)
	}
	pre_append = pre + ".0." + "srcipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcipv6"], _ = expandObjectSystemNpuNpuTcamDataSrcipv62edl(d, i["srcipv6"], pre_append)
	}
	pre_append = pre + ".0." + "srcmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcmac"], _ = expandObjectSystemNpuNpuTcamDataSrcmac2edl(d, i["srcmac"], pre_append)
	}
	pre_append = pre + ".0." + "srcport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcport"], _ = expandObjectSystemNpuNpuTcamDataSrcport2edl(d, i["srcport"], pre_append)
	}
	pre_append = pre + ".0." + "svid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["svid"], _ = expandObjectSystemNpuNpuTcamDataSvid2edl(d, i["svid"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ack"], _ = expandObjectSystemNpuNpuTcamDataTcpAck2edl(d, i["tcp_ack"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-cwr"], _ = expandObjectSystemNpuNpuTcamDataTcpCwr2edl(d, i["tcp_cwr"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ece"], _ = expandObjectSystemNpuNpuTcamDataTcpEce2edl(d, i["tcp_ece"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-fin"], _ = expandObjectSystemNpuNpuTcamDataTcpFin2edl(d, i["tcp_fin"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_push"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-push"], _ = expandObjectSystemNpuNpuTcamDataTcpPush2edl(d, i["tcp_push"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-rst"], _ = expandObjectSystemNpuNpuTcamDataTcpRst2edl(d, i["tcp_rst"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-syn"], _ = expandObjectSystemNpuNpuTcamDataTcpSyn2edl(d, i["tcp_syn"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-urg"], _ = expandObjectSystemNpuNpuTcamDataTcpUrg2edl(d, i["tcp_urg"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-cfi"], _ = expandObjectSystemNpuNpuTcamDataTgtCfi2edl(d, i["tgt_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-prio"], _ = expandObjectSystemNpuNpuTcamDataTgtPrio2edl(d, i["tgt_prio"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-updt"], _ = expandObjectSystemNpuNpuTcamDataTgtUpdt2edl(d, i["tgt_updt"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-v"], _ = expandObjectSystemNpuNpuTcamDataTgtV2edl(d, i["tgt_v"], pre_append)
	}
	pre_append = pre + ".0." + "tos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tos"], _ = expandObjectSystemNpuNpuTcamDataTos2edl(d, i["tos"], pre_append)
	}
	pre_append = pre + ".0." + "tp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp"], _ = expandObjectSystemNpuNpuTcamDataTp2edl(d, i["tp"], pre_append)
	}
	pre_append = pre + ".0." + "ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ttl"], _ = expandObjectSystemNpuNpuTcamDataTtl2edl(d, i["ttl"], pre_append)
	}
	pre_append = pre + ".0." + "tvid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tvid"], _ = expandObjectSystemNpuNpuTcamDataTvid2edl(d, i["tvid"], pre_append)
	}
	pre_append = pre + ".0." + "vdid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdid"], _ = expandObjectSystemNpuNpuTcamDataVdid2edl(d, i["vdid"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamDataDf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstipv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstmac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataDstport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataEthertype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataExtTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataFragOff2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenBufCnt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenIv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenL3Flags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenL4Flags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPktCtrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPri2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenPriV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataGenTv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIhl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIp4Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIp6Fl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataIpver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd102edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd112edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataL4Wd92edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataMf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSlink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSmacChange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcCfi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcPrio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcUpdt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcipv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcmac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSrcport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataSvid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpCwr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpEce2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpFin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpPush2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpRst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpSyn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTcpUrg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtCfi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtPrio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtUpdt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTgtV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataTvid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDataVdid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamDbgDump2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMask2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "df"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df"], _ = expandObjectSystemNpuNpuTcamMaskDf2edl(d, i["df"], pre_append)
	}
	pre_append = pre + ".0." + "dstip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstip"], _ = expandObjectSystemNpuNpuTcamMaskDstip2edl(d, i["dstip"], pre_append)
	}
	pre_append = pre + ".0." + "dstipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstipv6"], _ = expandObjectSystemNpuNpuTcamMaskDstipv62edl(d, i["dstipv6"], pre_append)
	}
	pre_append = pre + ".0." + "dstmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstmac"], _ = expandObjectSystemNpuNpuTcamMaskDstmac2edl(d, i["dstmac"], pre_append)
	}
	pre_append = pre + ".0." + "dstport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dstport"], _ = expandObjectSystemNpuNpuTcamMaskDstport2edl(d, i["dstport"], pre_append)
	}
	pre_append = pre + ".0." + "ethertype"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ethertype"], _ = expandObjectSystemNpuNpuTcamMaskEthertype2edl(d, i["ethertype"], pre_append)
	}
	pre_append = pre + ".0." + "ext_tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ext-tag"], _ = expandObjectSystemNpuNpuTcamMaskExtTag2edl(d, i["ext_tag"], pre_append)
	}
	pre_append = pre + ".0." + "frag_off"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-off"], _ = expandObjectSystemNpuNpuTcamMaskFragOff2edl(d, i["frag_off"], pre_append)
	}
	pre_append = pre + ".0." + "gen_buf_cnt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-buf-cnt"], _ = expandObjectSystemNpuNpuTcamMaskGenBufCnt2edl(d, i["gen_buf_cnt"], pre_append)
	}
	pre_append = pre + ".0." + "gen_iv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-iv"], _ = expandObjectSystemNpuNpuTcamMaskGenIv2edl(d, i["gen_iv"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l3_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l3-flags"], _ = expandObjectSystemNpuNpuTcamMaskGenL3Flags2edl(d, i["gen_l3_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_l4_flags"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-l4-flags"], _ = expandObjectSystemNpuNpuTcamMaskGenL4Flags2edl(d, i["gen_l4_flags"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pkt_ctrl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pkt-ctrl"], _ = expandObjectSystemNpuNpuTcamMaskGenPktCtrl2edl(d, i["gen_pkt_ctrl"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri"], _ = expandObjectSystemNpuNpuTcamMaskGenPri2edl(d, i["gen_pri"], pre_append)
	}
	pre_append = pre + ".0." + "gen_pri_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-pri-v"], _ = expandObjectSystemNpuNpuTcamMaskGenPriV2edl(d, i["gen_pri_v"], pre_append)
	}
	pre_append = pre + ".0." + "gen_tv"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gen-tv"], _ = expandObjectSystemNpuNpuTcamMaskGenTv2edl(d, i["gen_tv"], pre_append)
	}
	pre_append = pre + ".0." + "ihl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ihl"], _ = expandObjectSystemNpuNpuTcamMaskIhl2edl(d, i["ihl"], pre_append)
	}
	pre_append = pre + ".0." + "ip4_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip4-id"], _ = expandObjectSystemNpuNpuTcamMaskIp4Id2edl(d, i["ip4_id"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_fl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip6-fl"], _ = expandObjectSystemNpuNpuTcamMaskIp6Fl2edl(d, i["ip6_fl"], pre_append)
	}
	pre_append = pre + ".0." + "ipver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipver"], _ = expandObjectSystemNpuNpuTcamMaskIpver2edl(d, i["ipver"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd10"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd10"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd102edl(d, i["l4_wd10"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd11"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd11"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd112edl(d, i["l4_wd11"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd8"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd8"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd82edl(d, i["l4_wd8"], pre_append)
	}
	pre_append = pre + ".0." + "l4_wd9"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["l4-wd9"], _ = expandObjectSystemNpuNpuTcamMaskL4Wd92edl(d, i["l4_wd9"], pre_append)
	}
	pre_append = pre + ".0." + "mf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mf"], _ = expandObjectSystemNpuNpuTcamMaskMf2edl(d, i["mf"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandObjectSystemNpuNpuTcamMaskProtocol2edl(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "slink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slink"], _ = expandObjectSystemNpuNpuTcamMaskSlink2edl(d, i["slink"], pre_append)
	}
	pre_append = pre + ".0." + "smac_change"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-change"], _ = expandObjectSystemNpuNpuTcamMaskSmacChange2edl(d, i["smac_change"], pre_append)
	}
	pre_append = pre + ".0." + "sp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sp"], _ = expandObjectSystemNpuNpuTcamMaskSp2edl(d, i["sp"], pre_append)
	}
	pre_append = pre + ".0." + "src_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-cfi"], _ = expandObjectSystemNpuNpuTcamMaskSrcCfi2edl(d, i["src_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "src_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-prio"], _ = expandObjectSystemNpuNpuTcamMaskSrcPrio2edl(d, i["src_prio"], pre_append)
	}
	pre_append = pre + ".0." + "src_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-updt"], _ = expandObjectSystemNpuNpuTcamMaskSrcUpdt2edl(d, i["src_updt"], pre_append)
	}
	pre_append = pre + ".0." + "srcip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcip"], _ = expandObjectSystemNpuNpuTcamMaskSrcip2edl(d, i["srcip"], pre_append)
	}
	pre_append = pre + ".0." + "srcipv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcipv6"], _ = expandObjectSystemNpuNpuTcamMaskSrcipv62edl(d, i["srcipv6"], pre_append)
	}
	pre_append = pre + ".0." + "srcmac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcmac"], _ = expandObjectSystemNpuNpuTcamMaskSrcmac2edl(d, i["srcmac"], pre_append)
	}
	pre_append = pre + ".0." + "srcport"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["srcport"], _ = expandObjectSystemNpuNpuTcamMaskSrcport2edl(d, i["srcport"], pre_append)
	}
	pre_append = pre + ".0." + "svid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["svid"], _ = expandObjectSystemNpuNpuTcamMaskSvid2edl(d, i["svid"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ack"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ack"], _ = expandObjectSystemNpuNpuTcamMaskTcpAck2edl(d, i["tcp_ack"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_cwr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-cwr"], _ = expandObjectSystemNpuNpuTcamMaskTcpCwr2edl(d, i["tcp_cwr"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_ece"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-ece"], _ = expandObjectSystemNpuNpuTcamMaskTcpEce2edl(d, i["tcp_ece"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_fin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-fin"], _ = expandObjectSystemNpuNpuTcamMaskTcpFin2edl(d, i["tcp_fin"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_push"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-push"], _ = expandObjectSystemNpuNpuTcamMaskTcpPush2edl(d, i["tcp_push"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_rst"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-rst"], _ = expandObjectSystemNpuNpuTcamMaskTcpRst2edl(d, i["tcp_rst"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_syn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-syn"], _ = expandObjectSystemNpuNpuTcamMaskTcpSyn2edl(d, i["tcp_syn"], pre_append)
	}
	pre_append = pre + ".0." + "tcp_urg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tcp-urg"], _ = expandObjectSystemNpuNpuTcamMaskTcpUrg2edl(d, i["tcp_urg"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_cfi"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-cfi"], _ = expandObjectSystemNpuNpuTcamMaskTgtCfi2edl(d, i["tgt_cfi"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_prio"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-prio"], _ = expandObjectSystemNpuNpuTcamMaskTgtPrio2edl(d, i["tgt_prio"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_updt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-updt"], _ = expandObjectSystemNpuNpuTcamMaskTgtUpdt2edl(d, i["tgt_updt"], pre_append)
	}
	pre_append = pre + ".0." + "tgt_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgt-v"], _ = expandObjectSystemNpuNpuTcamMaskTgtV2edl(d, i["tgt_v"], pre_append)
	}
	pre_append = pre + ".0." + "tos"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tos"], _ = expandObjectSystemNpuNpuTcamMaskTos2edl(d, i["tos"], pre_append)
	}
	pre_append = pre + ".0." + "tp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp"], _ = expandObjectSystemNpuNpuTcamMaskTp2edl(d, i["tp"], pre_append)
	}
	pre_append = pre + ".0." + "ttl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ttl"], _ = expandObjectSystemNpuNpuTcamMaskTtl2edl(d, i["ttl"], pre_append)
	}
	pre_append = pre + ".0." + "tvid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tvid"], _ = expandObjectSystemNpuNpuTcamMaskTvid2edl(d, i["tvid"], pre_append)
	}
	pre_append = pre + ".0." + "vdid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdid"], _ = expandObjectSystemNpuNpuTcamMaskVdid2edl(d, i["vdid"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamMaskDf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstipv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstmac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskDstport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskEthertype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskExtTag2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskFragOff2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenBufCnt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenIv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenL3Flags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenL4Flags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPktCtrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPri2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenPriV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskGenTv2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIhl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIp4Id2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIp6Fl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskIpver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd102edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd112edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskL4Wd92edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskMf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSlink2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSmacChange2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcCfi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcPrio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcUpdt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcipv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcmac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSrcport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskSvid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpAck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpCwr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpEce2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpFin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpPush2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpRst2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpSyn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTcpUrg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtCfi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtPrio2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtUpdt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTgtV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskTvid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMaskVdid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamMirAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vlif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlif"], _ = expandObjectSystemNpuNpuTcamMirActVlif2edl(d, i["vlif"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamMirActVlif2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamOid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamPriAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandObjectSystemNpuNpuTcamPriActPriority2edl(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "weight"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["weight"], _ = expandObjectSystemNpuNpuTcamPriActWeight2edl(d, i["weight"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamPriActPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamPriActWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act"], _ = expandObjectSystemNpuNpuTcamSactAct2edl(d, i["act"], pre_append)
	}
	pre_append = pre + ".0." + "act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act-v"], _ = expandObjectSystemNpuNpuTcamSactActV2edl(d, i["act_v"], pre_append)
	}
	pre_append = pre + ".0." + "bmproc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bmproc"], _ = expandObjectSystemNpuNpuTcamSactBmproc2edl(d, i["bmproc"], pre_append)
	}
	pre_append = pre + ".0." + "bmproc_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["bmproc-v"], _ = expandObjectSystemNpuNpuTcamSactBmprocV2edl(d, i["bmproc_v"], pre_append)
	}
	pre_append = pre + ".0." + "df_lif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df-lif"], _ = expandObjectSystemNpuNpuTcamSactDfLif2edl(d, i["df_lif"], pre_append)
	}
	pre_append = pre + ".0." + "df_lif_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["df-lif-v"], _ = expandObjectSystemNpuNpuTcamSactDfLifV2edl(d, i["df_lif_v"], pre_append)
	}
	pre_append = pre + ".0." + "dfr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dfr"], _ = expandObjectSystemNpuNpuTcamSactDfr2edl(d, i["dfr"], pre_append)
	}
	pre_append = pre + ".0." + "dfr_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dfr-v"], _ = expandObjectSystemNpuNpuTcamSactDfrV2edl(d, i["dfr_v"], pre_append)
	}
	pre_append = pre + ".0." + "dmac_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dmac-skip"], _ = expandObjectSystemNpuNpuTcamSactDmacSkip2edl(d, i["dmac_skip"], pre_append)
	}
	pre_append = pre + ".0." + "dmac_skip_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dmac-skip-v"], _ = expandObjectSystemNpuNpuTcamSactDmacSkipV2edl(d, i["dmac_skip_v"], pre_append)
	}
	pre_append = pre + ".0." + "dosen"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dosen"], _ = expandObjectSystemNpuNpuTcamSactDosen2edl(d, i["dosen"], pre_append)
	}
	pre_append = pre + ".0." + "dosen_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dosen-v"], _ = expandObjectSystemNpuNpuTcamSactDosenV2edl(d, i["dosen_v"], pre_append)
	}
	pre_append = pre + ".0." + "espff_proc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["espff-proc"], _ = expandObjectSystemNpuNpuTcamSactEspffProc2edl(d, i["espff_proc"], pre_append)
	}
	pre_append = pre + ".0." + "espff_proc_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["espff-proc-v"], _ = expandObjectSystemNpuNpuTcamSactEspffProcV2edl(d, i["espff_proc_v"], pre_append)
	}
	pre_append = pre + ".0." + "etype_pid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["etype-pid"], _ = expandObjectSystemNpuNpuTcamSactEtypePid2edl(d, i["etype_pid"], pre_append)
	}
	pre_append = pre + ".0." + "etype_pid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["etype-pid-v"], _ = expandObjectSystemNpuNpuTcamSactEtypePidV2edl(d, i["etype_pid_v"], pre_append)
	}
	pre_append = pre + ".0." + "frag_proc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-proc"], _ = expandObjectSystemNpuNpuTcamSactFragProc2edl(d, i["frag_proc"], pre_append)
	}
	pre_append = pre + ".0." + "frag_proc_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["frag-proc-v"], _ = expandObjectSystemNpuNpuTcamSactFragProcV2edl(d, i["frag_proc_v"], pre_append)
	}
	pre_append = pre + ".0." + "fwd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd"], _ = expandObjectSystemNpuNpuTcamSactFwd2edl(d, i["fwd"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_lif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-lif"], _ = expandObjectSystemNpuNpuTcamSactFwdLif2edl(d, i["fwd_lif"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_lif_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-lif-v"], _ = expandObjectSystemNpuNpuTcamSactFwdLifV2edl(d, i["fwd_lif_v"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_tvid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-tvid"], _ = expandObjectSystemNpuNpuTcamSactFwdTvid2edl(d, i["fwd_tvid"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_tvid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-tvid-v"], _ = expandObjectSystemNpuNpuTcamSactFwdTvidV2edl(d, i["fwd_tvid_v"], pre_append)
	}
	pre_append = pre + ".0." + "fwd_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fwd-v"], _ = expandObjectSystemNpuNpuTcamSactFwdV2edl(d, i["fwd_v"], pre_append)
	}
	pre_append = pre + ".0." + "icpen"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icpen"], _ = expandObjectSystemNpuNpuTcamSactIcpen2edl(d, i["icpen"], pre_append)
	}
	pre_append = pre + ".0." + "icpen_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icpen-v"], _ = expandObjectSystemNpuNpuTcamSactIcpenV2edl(d, i["icpen_v"], pre_append)
	}
	pre_append = pre + ".0." + "igmp_mld_snp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["igmp-mld-snp"], _ = expandObjectSystemNpuNpuTcamSactIgmpMldSnp2edl(d, i["igmp_mld_snp"], pre_append)
	}
	pre_append = pre + ".0." + "igmp_mld_snp_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["igmp-mld-snp-v"], _ = expandObjectSystemNpuNpuTcamSactIgmpMldSnpV2edl(d, i["igmp_mld_snp_v"], pre_append)
	}
	pre_append = pre + ".0." + "learn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["learn"], _ = expandObjectSystemNpuNpuTcamSactLearn2edl(d, i["learn"], pre_append)
	}
	pre_append = pre + ".0." + "learn_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["learn-v"], _ = expandObjectSystemNpuNpuTcamSactLearnV2edl(d, i["learn_v"], pre_append)
	}
	pre_append = pre + ".0." + "m_srh_ctrl"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["m-srh-ctrl"], _ = expandObjectSystemNpuNpuTcamSactMSrhCtrl2edl(d, i["m_srh_ctrl"], pre_append)
	}
	pre_append = pre + ".0." + "m_srh_ctrl_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["m-srh-ctrl-v"], _ = expandObjectSystemNpuNpuTcamSactMSrhCtrlV2edl(d, i["m_srh_ctrl_v"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id"], _ = expandObjectSystemNpuNpuTcamSactMacId2edl(d, i["mac_id"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id-v"], _ = expandObjectSystemNpuNpuTcamSactMacIdV2edl(d, i["mac_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "mss"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss"], _ = expandObjectSystemNpuNpuTcamSactMss2edl(d, i["mss"], pre_append)
	}
	pre_append = pre + ".0." + "mss_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss-v"], _ = expandObjectSystemNpuNpuTcamSactMssV2edl(d, i["mss_v"], pre_append)
	}
	pre_append = pre + ".0." + "pleen"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pleen"], _ = expandObjectSystemNpuNpuTcamSactPleen2edl(d, i["pleen"], pre_append)
	}
	pre_append = pre + ".0." + "pleen_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pleen-v"], _ = expandObjectSystemNpuNpuTcamSactPleenV2edl(d, i["pleen_v"], pre_append)
	}
	pre_append = pre + ".0." + "prio_pid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["prio-pid"], _ = expandObjectSystemNpuNpuTcamSactPrioPid2edl(d, i["prio_pid"], pre_append)
	}
	pre_append = pre + ".0." + "prio_pid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["prio-pid-v"], _ = expandObjectSystemNpuNpuTcamSactPrioPidV2edl(d, i["prio_pid_v"], pre_append)
	}
	pre_append = pre + ".0." + "promis"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["promis"], _ = expandObjectSystemNpuNpuTcamSactPromis2edl(d, i["promis"], pre_append)
	}
	pre_append = pre + ".0." + "promis_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["promis-v"], _ = expandObjectSystemNpuNpuTcamSactPromisV2edl(d, i["promis_v"], pre_append)
	}
	pre_append = pre + ".0." + "rfsh"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rfsh"], _ = expandObjectSystemNpuNpuTcamSactRfsh2edl(d, i["rfsh"], pre_append)
	}
	pre_append = pre + ".0." + "rfsh_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rfsh-v"], _ = expandObjectSystemNpuNpuTcamSactRfshV2edl(d, i["rfsh_v"], pre_append)
	}
	pre_append = pre + ".0." + "smac_skip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-skip"], _ = expandObjectSystemNpuNpuTcamSactSmacSkip2edl(d, i["smac_skip"], pre_append)
	}
	pre_append = pre + ".0." + "smac_skip_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["smac-skip-v"], _ = expandObjectSystemNpuNpuTcamSactSmacSkipV2edl(d, i["smac_skip_v"], pre_append)
	}
	pre_append = pre + ".0." + "tp_smchk_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp-smchk-v"], _ = expandObjectSystemNpuNpuTcamSactTpSmchkV2edl(d, i["tp_smchk_v"], pre_append)
	}
	pre_append = pre + ".0." + "tp_smchk"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tp_smchk"], _ = expandObjectSystemNpuNpuTcamSactTpSmchk2edl(d, i["tp_smchk"], pre_append)
	}
	pre_append = pre + ".0." + "tpe_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpe-id"], _ = expandObjectSystemNpuNpuTcamSactTpeId2edl(d, i["tpe_id"], pre_append)
	}
	pre_append = pre + ".0." + "tpe_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpe-id-v"], _ = expandObjectSystemNpuNpuTcamSactTpeIdV2edl(d, i["tpe_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "vdm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdm"], _ = expandObjectSystemNpuNpuTcamSactVdm2edl(d, i["vdm"], pre_append)
	}
	pre_append = pre + ".0." + "vdm_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdm-v"], _ = expandObjectSystemNpuNpuTcamSactVdmV2edl(d, i["vdm_v"], pre_append)
	}
	pre_append = pre + ".0." + "vdom_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdom-id"], _ = expandObjectSystemNpuNpuTcamSactVdomId2edl(d, i["vdom_id"], pre_append)
	}
	pre_append = pre + ".0." + "vdom_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vdom-id-v"], _ = expandObjectSystemNpuNpuTcamSactVdomIdV2edl(d, i["vdom_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "x_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["x-mode"], _ = expandObjectSystemNpuNpuTcamSactXMode2edl(d, i["x_mode"], pre_append)
	}
	pre_append = pre + ".0." + "x_mode_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["x-mode-v"], _ = expandObjectSystemNpuNpuTcamSactXModeV2edl(d, i["x_mode_v"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamSactAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactActV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactBmproc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactBmprocV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfLif2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfLifV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfrV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDmacSkip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDmacSkipV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDosen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDosenV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEspffProc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEspffProcV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEtypePid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEtypePidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFragProc2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFragProcV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdLif2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdLifV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdTvid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdTvidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIcpen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIcpenV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIgmpMldSnp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIgmpMldSnpV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactLearn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactLearnV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMSrhCtrl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMSrhCtrlV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMacId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMacIdV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMssV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPleen2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPleenV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPrioPid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPrioPidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPromis2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPromisV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactRfsh2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactRfshV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactSmacSkip2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactSmacSkipV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpSmchkV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpSmchk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpeId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpeIdV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdmV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdomId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdomIdV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactXMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactXModeV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTact2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act"], _ = expandObjectSystemNpuNpuTcamTactAct2edl(d, i["act"], pre_append)
	}
	pre_append = pre + ".0." + "act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["act-v"], _ = expandObjectSystemNpuNpuTcamTactActV2edl(d, i["act_v"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv4_s"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv4-s"], _ = expandObjectSystemNpuNpuTcamTactFmtuv4S2edl(d, i["fmtuv4_s"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv4_s_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv4-s-v"], _ = expandObjectSystemNpuNpuTcamTactFmtuv4SV2edl(d, i["fmtuv4_s_v"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv6_s"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv6-s"], _ = expandObjectSystemNpuNpuTcamTactFmtuv6S2edl(d, i["fmtuv6_s"], pre_append)
	}
	pre_append = pre + ".0." + "fmtuv6_s_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fmtuv6-s-v"], _ = expandObjectSystemNpuNpuTcamTactFmtuv6SV2edl(d, i["fmtuv6_s_v"], pre_append)
	}
	pre_append = pre + ".0." + "lnkid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lnkid"], _ = expandObjectSystemNpuNpuTcamTactLnkid2edl(d, i["lnkid"], pre_append)
	}
	pre_append = pre + ".0." + "lnkid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["lnkid-v"], _ = expandObjectSystemNpuNpuTcamTactLnkidV2edl(d, i["lnkid_v"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id"], _ = expandObjectSystemNpuNpuTcamTactMacId2edl(d, i["mac_id"], pre_append)
	}
	pre_append = pre + ".0." + "mac_id_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mac-id-v"], _ = expandObjectSystemNpuNpuTcamTactMacIdV2edl(d, i["mac_id_v"], pre_append)
	}
	pre_append = pre + ".0." + "mss_t"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss-t"], _ = expandObjectSystemNpuNpuTcamTactMssT2edl(d, i["mss_t"], pre_append)
	}
	pre_append = pre + ".0." + "mss_t_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mss-t-v"], _ = expandObjectSystemNpuNpuTcamTactMssTV2edl(d, i["mss_t_v"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv4"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv4"], _ = expandObjectSystemNpuNpuTcamTactMtuv42edl(d, i["mtuv4"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv4_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv4-v"], _ = expandObjectSystemNpuNpuTcamTactMtuv4V2edl(d, i["mtuv4_v"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv6"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv6"], _ = expandObjectSystemNpuNpuTcamTactMtuv62edl(d, i["mtuv6"], pre_append)
	}
	pre_append = pre + ".0." + "mtuv6_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mtuv6-v"], _ = expandObjectSystemNpuNpuTcamTactMtuv6V2edl(d, i["mtuv6_v"], pre_append)
	}
	pre_append = pre + ".0." + "slif_act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slif-act"], _ = expandObjectSystemNpuNpuTcamTactSlifAct2edl(d, i["slif_act"], pre_append)
	}
	pre_append = pre + ".0." + "slif_act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["slif-act-v"], _ = expandObjectSystemNpuNpuTcamTactSlifActV2edl(d, i["slif_act_v"], pre_append)
	}
	pre_append = pre + ".0." + "sublnkid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sublnkid"], _ = expandObjectSystemNpuNpuTcamTactSublnkid2edl(d, i["sublnkid"], pre_append)
	}
	pre_append = pre + ".0." + "sublnkid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sublnkid-v"], _ = expandObjectSystemNpuNpuTcamTactSublnkidV2edl(d, i["sublnkid_v"], pre_append)
	}
	pre_append = pre + ".0." + "tgtv_act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgtv-act"], _ = expandObjectSystemNpuNpuTcamTactTgtvAct2edl(d, i["tgtv_act"], pre_append)
	}
	pre_append = pre + ".0." + "tgtv_act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tgtv-act-v"], _ = expandObjectSystemNpuNpuTcamTactTgtvActV2edl(d, i["tgtv_act_v"], pre_append)
	}
	pre_append = pre + ".0." + "tlif_act"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tlif-act"], _ = expandObjectSystemNpuNpuTcamTactTlifAct2edl(d, i["tlif_act"], pre_append)
	}
	pre_append = pre + ".0." + "tlif_act_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tlif-act-v"], _ = expandObjectSystemNpuNpuTcamTactTlifActV2edl(d, i["tlif_act_v"], pre_append)
	}
	pre_append = pre + ".0." + "tpeid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpeid"], _ = expandObjectSystemNpuNpuTcamTactTpeid2edl(d, i["tpeid"], pre_append)
	}
	pre_append = pre + ".0." + "tpeid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tpeid-v"], _ = expandObjectSystemNpuNpuTcamTactTpeidV2edl(d, i["tpeid_v"], pre_append)
	}
	pre_append = pre + ".0." + "v6fe"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["v6fe"], _ = expandObjectSystemNpuNpuTcamTactV6Fe2edl(d, i["v6fe"], pre_append)
	}
	pre_append = pre + ".0." + "v6fe_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["v6fe-v"], _ = expandObjectSystemNpuNpuTcamTactV6FeV2edl(d, i["v6fe_v"], pre_append)
	}
	pre_append = pre + ".0." + "vep_en_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep-en-v"], _ = expandObjectSystemNpuNpuTcamTactVepEnV2edl(d, i["vep_en_v"], pre_append)
	}
	pre_append = pre + ".0." + "vep_slid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep-slid"], _ = expandObjectSystemNpuNpuTcamTactVepSlid2edl(d, i["vep_slid"], pre_append)
	}
	pre_append = pre + ".0." + "vep_slid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep-slid-v"], _ = expandObjectSystemNpuNpuTcamTactVepSlidV2edl(d, i["vep_slid_v"], pre_append)
	}
	pre_append = pre + ".0." + "vep_en"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vep_en"], _ = expandObjectSystemNpuNpuTcamTactVepEn2edl(d, i["vep_en"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_lif"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-lif"], _ = expandObjectSystemNpuNpuTcamTactXltLif2edl(d, i["xlt_lif"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_lif_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-lif-v"], _ = expandObjectSystemNpuNpuTcamTactXltLifV2edl(d, i["xlt_lif_v"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_vid"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-vid"], _ = expandObjectSystemNpuNpuTcamTactXltVid2edl(d, i["xlt_vid"], pre_append)
	}
	pre_append = pre + ".0." + "xlt_vid_v"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["xlt-vid-v"], _ = expandObjectSystemNpuNpuTcamTactXltVidV2edl(d, i["xlt_vid_v"], pre_append)
	}

	return result, nil
}

func expandObjectSystemNpuNpuTcamTactAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactActV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv4S2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv4SV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv6S2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv6SV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactLnkid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactLnkidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMacId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMacIdV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMssT2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMssTV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv4V2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv6V2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSlifAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSlifActV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSublnkid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSublnkidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTgtvAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTgtvActV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTlifAct2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTlifActV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTpeid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTpeidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactV6Fe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactV6FeV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepEnV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepSlid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepSlidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepEn2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltLif2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltLifV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltVid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltVidV2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamVid2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcam(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("data"); ok || d.HasChange("data") {
		t, err := expandObjectSystemNpuNpuTcamData2edl(d, v, "data")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data"] = t
		}
	}

	if v, ok := d.GetOk("dbg_dump"); ok || d.HasChange("dbg_dump") {
		t, err := expandObjectSystemNpuNpuTcamDbgDump2edl(d, v, "dbg_dump")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dbg-dump"] = t
		}
	}

	if v, ok := d.GetOk("mask"); ok || d.HasChange("mask") {
		t, err := expandObjectSystemNpuNpuTcamMask2edl(d, v, "mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mask"] = t
		}
	}

	if v, ok := d.GetOk("mir_act"); ok || d.HasChange("mir_act") {
		t, err := expandObjectSystemNpuNpuTcamMirAct2edl(d, v, "mir_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mir-act"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSystemNpuNpuTcamName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok || d.HasChange("oid") {
		t, err := expandObjectSystemNpuNpuTcamOid2edl(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("pri_act"); ok || d.HasChange("pri_act") {
		t, err := expandObjectSystemNpuNpuTcamPriAct2edl(d, v, "pri_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pri-act"] = t
		}
	}

	if v, ok := d.GetOk("sact"); ok || d.HasChange("sact") {
		t, err := expandObjectSystemNpuNpuTcamSact2edl(d, v, "sact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sact"] = t
		}
	}

	if v, ok := d.GetOk("tact"); ok || d.HasChange("tact") {
		t, err := expandObjectSystemNpuNpuTcamTact2edl(d, v, "tact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tact"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectSystemNpuNpuTcamType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vid"); ok || d.HasChange("vid") {
		t, err := expandObjectSystemNpuNpuTcamVid2edl(d, v, "vid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vid"] = t
		}
	}

	return &obj, nil
}
