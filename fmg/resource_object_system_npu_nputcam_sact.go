// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Source action of TCAM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcamSact() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamSactUpdate,
		Read:   resourceObjectSystemNpuNpuTcamSactRead,
		Update: resourceObjectSystemNpuNpuTcamSactUpdate,
		Delete: resourceObjectSystemNpuNpuTcamSactDelete,

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
			"act": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"act_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bmproc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bmproc_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"df_lif": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"df_lif_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dfr_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dmac_skip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dmac_skip_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dosen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dosen_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"espff_proc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"espff_proc_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"etype_pid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"etype_pid_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frag_proc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"frag_proc_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"fwd_tvid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fwd_tvid_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwd_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icpen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icpen_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_mld_snp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"igmp_mld_snp_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"learn_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"m_srh_ctrl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"m_srh_ctrl_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mac_id_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mss_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pleen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pleen_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prio_pid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prio_pid_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"promis": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"promis_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rfsh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rfsh_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smac_skip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"smac_skip_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tp_smchk_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
			},
			"vdm": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vdm_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vdom_id_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"x_mode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"x_mode_v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectSystemNpuNpuTcamSactUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	npu_tcam := d.Get("npu_tcam").(string)
	paradict["npu_tcam"] = npu_tcam

	obj, err := getObjectObjectSystemNpuNpuTcamSact(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamSact resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuNpuTcamSact(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamSact resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpuTcamSact")

	return resourceObjectSystemNpuNpuTcamSactRead(d, m)
}

func resourceObjectSystemNpuNpuTcamSactDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	npu_tcam := d.Get("npu_tcam").(string)
	paradict["npu_tcam"] = npu_tcam

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemNpuNpuTcamSact(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcamSact resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamSactRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpuTcamSact(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamSact resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcamSact(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamSact resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamSactAct3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactActV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactBmproc3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactBmprocV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfLif3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfLifV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfr3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDfrV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDmacSkip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDmacSkipV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDosen3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactDosenV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEspffProc3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEspffProcV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEtypePid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactEtypePidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFragProc3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFragProcV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwd3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdLif3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdLifV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdTvid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdTvidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactFwdV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIcpen3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIcpenV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIgmpMldSnp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactIgmpMldSnpV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactLearn3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactLearnV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMSrhCtrl3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMSrhCtrlV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMacId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMacIdV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMss3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactMssV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPleen3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPleenV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPrioPid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPrioPidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPromis3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactPromisV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactRfsh3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactRfshV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactSmacSkip3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactSmacSkipV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpSmchkV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpSmchk3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactTpeIdV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdm3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdmV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdomId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactVdomIdV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactXMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamSactXModeV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcamSact(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("act", flattenObjectSystemNpuNpuTcamSactAct3rdl(o["act"], d, "act")); err != nil {
		if vv, ok := fortiAPIPatch(o["act"], "ObjectSystemNpuNpuTcamSact-Act"); ok {
			if err = d.Set("act", vv); err != nil {
				return fmt.Errorf("Error reading act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading act: %v", err)
		}
	}

	if err = d.Set("act_v", flattenObjectSystemNpuNpuTcamSactActV3rdl(o["act-v"], d, "act_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["act-v"], "ObjectSystemNpuNpuTcamSact-ActV"); ok {
			if err = d.Set("act_v", vv); err != nil {
				return fmt.Errorf("Error reading act_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading act_v: %v", err)
		}
	}

	if err = d.Set("bmproc", flattenObjectSystemNpuNpuTcamSactBmproc3rdl(o["bmproc"], d, "bmproc")); err != nil {
		if vv, ok := fortiAPIPatch(o["bmproc"], "ObjectSystemNpuNpuTcamSact-Bmproc"); ok {
			if err = d.Set("bmproc", vv); err != nil {
				return fmt.Errorf("Error reading bmproc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bmproc: %v", err)
		}
	}

	if err = d.Set("bmproc_v", flattenObjectSystemNpuNpuTcamSactBmprocV3rdl(o["bmproc-v"], d, "bmproc_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["bmproc-v"], "ObjectSystemNpuNpuTcamSact-BmprocV"); ok {
			if err = d.Set("bmproc_v", vv); err != nil {
				return fmt.Errorf("Error reading bmproc_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bmproc_v: %v", err)
		}
	}

	if err = d.Set("df_lif", flattenObjectSystemNpuNpuTcamSactDfLif3rdl(o["df-lif"], d, "df_lif")); err != nil {
		if vv, ok := fortiAPIPatch(o["df-lif"], "ObjectSystemNpuNpuTcamSact-DfLif"); ok {
			if err = d.Set("df_lif", vv); err != nil {
				return fmt.Errorf("Error reading df_lif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading df_lif: %v", err)
		}
	}

	if err = d.Set("df_lif_v", flattenObjectSystemNpuNpuTcamSactDfLifV3rdl(o["df-lif-v"], d, "df_lif_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["df-lif-v"], "ObjectSystemNpuNpuTcamSact-DfLifV"); ok {
			if err = d.Set("df_lif_v", vv); err != nil {
				return fmt.Errorf("Error reading df_lif_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading df_lif_v: %v", err)
		}
	}

	if err = d.Set("dfr", flattenObjectSystemNpuNpuTcamSactDfr3rdl(o["dfr"], d, "dfr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dfr"], "ObjectSystemNpuNpuTcamSact-Dfr"); ok {
			if err = d.Set("dfr", vv); err != nil {
				return fmt.Errorf("Error reading dfr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dfr: %v", err)
		}
	}

	if err = d.Set("dfr_v", flattenObjectSystemNpuNpuTcamSactDfrV3rdl(o["dfr-v"], d, "dfr_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["dfr-v"], "ObjectSystemNpuNpuTcamSact-DfrV"); ok {
			if err = d.Set("dfr_v", vv); err != nil {
				return fmt.Errorf("Error reading dfr_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dfr_v: %v", err)
		}
	}

	if err = d.Set("dmac_skip", flattenObjectSystemNpuNpuTcamSactDmacSkip3rdl(o["dmac-skip"], d, "dmac_skip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dmac-skip"], "ObjectSystemNpuNpuTcamSact-DmacSkip"); ok {
			if err = d.Set("dmac_skip", vv); err != nil {
				return fmt.Errorf("Error reading dmac_skip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dmac_skip: %v", err)
		}
	}

	if err = d.Set("dmac_skip_v", flattenObjectSystemNpuNpuTcamSactDmacSkipV3rdl(o["dmac-skip-v"], d, "dmac_skip_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["dmac-skip-v"], "ObjectSystemNpuNpuTcamSact-DmacSkipV"); ok {
			if err = d.Set("dmac_skip_v", vv); err != nil {
				return fmt.Errorf("Error reading dmac_skip_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dmac_skip_v: %v", err)
		}
	}

	if err = d.Set("dosen", flattenObjectSystemNpuNpuTcamSactDosen3rdl(o["dosen"], d, "dosen")); err != nil {
		if vv, ok := fortiAPIPatch(o["dosen"], "ObjectSystemNpuNpuTcamSact-Dosen"); ok {
			if err = d.Set("dosen", vv); err != nil {
				return fmt.Errorf("Error reading dosen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dosen: %v", err)
		}
	}

	if err = d.Set("dosen_v", flattenObjectSystemNpuNpuTcamSactDosenV3rdl(o["dosen-v"], d, "dosen_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["dosen-v"], "ObjectSystemNpuNpuTcamSact-DosenV"); ok {
			if err = d.Set("dosen_v", vv); err != nil {
				return fmt.Errorf("Error reading dosen_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dosen_v: %v", err)
		}
	}

	if err = d.Set("espff_proc", flattenObjectSystemNpuNpuTcamSactEspffProc3rdl(o["espff-proc"], d, "espff_proc")); err != nil {
		if vv, ok := fortiAPIPatch(o["espff-proc"], "ObjectSystemNpuNpuTcamSact-EspffProc"); ok {
			if err = d.Set("espff_proc", vv); err != nil {
				return fmt.Errorf("Error reading espff_proc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading espff_proc: %v", err)
		}
	}

	if err = d.Set("espff_proc_v", flattenObjectSystemNpuNpuTcamSactEspffProcV3rdl(o["espff-proc-v"], d, "espff_proc_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["espff-proc-v"], "ObjectSystemNpuNpuTcamSact-EspffProcV"); ok {
			if err = d.Set("espff_proc_v", vv); err != nil {
				return fmt.Errorf("Error reading espff_proc_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading espff_proc_v: %v", err)
		}
	}

	if err = d.Set("etype_pid", flattenObjectSystemNpuNpuTcamSactEtypePid3rdl(o["etype-pid"], d, "etype_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["etype-pid"], "ObjectSystemNpuNpuTcamSact-EtypePid"); ok {
			if err = d.Set("etype_pid", vv); err != nil {
				return fmt.Errorf("Error reading etype_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading etype_pid: %v", err)
		}
	}

	if err = d.Set("etype_pid_v", flattenObjectSystemNpuNpuTcamSactEtypePidV3rdl(o["etype-pid-v"], d, "etype_pid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["etype-pid-v"], "ObjectSystemNpuNpuTcamSact-EtypePidV"); ok {
			if err = d.Set("etype_pid_v", vv); err != nil {
				return fmt.Errorf("Error reading etype_pid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading etype_pid_v: %v", err)
		}
	}

	if err = d.Set("frag_proc", flattenObjectSystemNpuNpuTcamSactFragProc3rdl(o["frag-proc"], d, "frag_proc")); err != nil {
		if vv, ok := fortiAPIPatch(o["frag-proc"], "ObjectSystemNpuNpuTcamSact-FragProc"); ok {
			if err = d.Set("frag_proc", vv); err != nil {
				return fmt.Errorf("Error reading frag_proc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frag_proc: %v", err)
		}
	}

	if err = d.Set("frag_proc_v", flattenObjectSystemNpuNpuTcamSactFragProcV3rdl(o["frag-proc-v"], d, "frag_proc_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["frag-proc-v"], "ObjectSystemNpuNpuTcamSact-FragProcV"); ok {
			if err = d.Set("frag_proc_v", vv); err != nil {
				return fmt.Errorf("Error reading frag_proc_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frag_proc_v: %v", err)
		}
	}

	if err = d.Set("fwd", flattenObjectSystemNpuNpuTcamSactFwd3rdl(o["fwd"], d, "fwd")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd"], "ObjectSystemNpuNpuTcamSact-Fwd"); ok {
			if err = d.Set("fwd", vv); err != nil {
				return fmt.Errorf("Error reading fwd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd: %v", err)
		}
	}

	if err = d.Set("fwd_lif", flattenObjectSystemNpuNpuTcamSactFwdLif3rdl(o["fwd-lif"], d, "fwd_lif")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-lif"], "ObjectSystemNpuNpuTcamSact-FwdLif"); ok {
			if err = d.Set("fwd_lif", vv); err != nil {
				return fmt.Errorf("Error reading fwd_lif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_lif: %v", err)
		}
	}

	if err = d.Set("fwd_lif_v", flattenObjectSystemNpuNpuTcamSactFwdLifV3rdl(o["fwd-lif-v"], d, "fwd_lif_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-lif-v"], "ObjectSystemNpuNpuTcamSact-FwdLifV"); ok {
			if err = d.Set("fwd_lif_v", vv); err != nil {
				return fmt.Errorf("Error reading fwd_lif_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_lif_v: %v", err)
		}
	}

	if err = d.Set("fwd_tvid", flattenObjectSystemNpuNpuTcamSactFwdTvid3rdl(o["fwd-tvid"], d, "fwd_tvid")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-tvid"], "ObjectSystemNpuNpuTcamSact-FwdTvid"); ok {
			if err = d.Set("fwd_tvid", vv); err != nil {
				return fmt.Errorf("Error reading fwd_tvid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_tvid: %v", err)
		}
	}

	if err = d.Set("fwd_tvid_v", flattenObjectSystemNpuNpuTcamSactFwdTvidV3rdl(o["fwd-tvid-v"], d, "fwd_tvid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-tvid-v"], "ObjectSystemNpuNpuTcamSact-FwdTvidV"); ok {
			if err = d.Set("fwd_tvid_v", vv); err != nil {
				return fmt.Errorf("Error reading fwd_tvid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_tvid_v: %v", err)
		}
	}

	if err = d.Set("fwd_v", flattenObjectSystemNpuNpuTcamSactFwdV3rdl(o["fwd-v"], d, "fwd_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwd-v"], "ObjectSystemNpuNpuTcamSact-FwdV"); ok {
			if err = d.Set("fwd_v", vv); err != nil {
				return fmt.Errorf("Error reading fwd_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwd_v: %v", err)
		}
	}

	if err = d.Set("icpen", flattenObjectSystemNpuNpuTcamSactIcpen3rdl(o["icpen"], d, "icpen")); err != nil {
		if vv, ok := fortiAPIPatch(o["icpen"], "ObjectSystemNpuNpuTcamSact-Icpen"); ok {
			if err = d.Set("icpen", vv); err != nil {
				return fmt.Errorf("Error reading icpen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icpen: %v", err)
		}
	}

	if err = d.Set("icpen_v", flattenObjectSystemNpuNpuTcamSactIcpenV3rdl(o["icpen-v"], d, "icpen_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["icpen-v"], "ObjectSystemNpuNpuTcamSact-IcpenV"); ok {
			if err = d.Set("icpen_v", vv); err != nil {
				return fmt.Errorf("Error reading icpen_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icpen_v: %v", err)
		}
	}

	if err = d.Set("igmp_mld_snp", flattenObjectSystemNpuNpuTcamSactIgmpMldSnp3rdl(o["igmp-mld-snp"], d, "igmp_mld_snp")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-mld-snp"], "ObjectSystemNpuNpuTcamSact-IgmpMldSnp"); ok {
			if err = d.Set("igmp_mld_snp", vv); err != nil {
				return fmt.Errorf("Error reading igmp_mld_snp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_mld_snp: %v", err)
		}
	}

	if err = d.Set("igmp_mld_snp_v", flattenObjectSystemNpuNpuTcamSactIgmpMldSnpV3rdl(o["igmp-mld-snp-v"], d, "igmp_mld_snp_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["igmp-mld-snp-v"], "ObjectSystemNpuNpuTcamSact-IgmpMldSnpV"); ok {
			if err = d.Set("igmp_mld_snp_v", vv); err != nil {
				return fmt.Errorf("Error reading igmp_mld_snp_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading igmp_mld_snp_v: %v", err)
		}
	}

	if err = d.Set("learn", flattenObjectSystemNpuNpuTcamSactLearn3rdl(o["learn"], d, "learn")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn"], "ObjectSystemNpuNpuTcamSact-Learn"); ok {
			if err = d.Set("learn", vv); err != nil {
				return fmt.Errorf("Error reading learn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn: %v", err)
		}
	}

	if err = d.Set("learn_v", flattenObjectSystemNpuNpuTcamSactLearnV3rdl(o["learn-v"], d, "learn_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-v"], "ObjectSystemNpuNpuTcamSact-LearnV"); ok {
			if err = d.Set("learn_v", vv); err != nil {
				return fmt.Errorf("Error reading learn_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_v: %v", err)
		}
	}

	if err = d.Set("m_srh_ctrl", flattenObjectSystemNpuNpuTcamSactMSrhCtrl3rdl(o["m-srh-ctrl"], d, "m_srh_ctrl")); err != nil {
		if vv, ok := fortiAPIPatch(o["m-srh-ctrl"], "ObjectSystemNpuNpuTcamSact-MSrhCtrl"); ok {
			if err = d.Set("m_srh_ctrl", vv); err != nil {
				return fmt.Errorf("Error reading m_srh_ctrl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading m_srh_ctrl: %v", err)
		}
	}

	if err = d.Set("m_srh_ctrl_v", flattenObjectSystemNpuNpuTcamSactMSrhCtrlV3rdl(o["m-srh-ctrl-v"], d, "m_srh_ctrl_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["m-srh-ctrl-v"], "ObjectSystemNpuNpuTcamSact-MSrhCtrlV"); ok {
			if err = d.Set("m_srh_ctrl_v", vv); err != nil {
				return fmt.Errorf("Error reading m_srh_ctrl_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading m_srh_ctrl_v: %v", err)
		}
	}

	if err = d.Set("mac_id", flattenObjectSystemNpuNpuTcamSactMacId3rdl(o["mac-id"], d, "mac_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-id"], "ObjectSystemNpuNpuTcamSact-MacId"); ok {
			if err = d.Set("mac_id", vv); err != nil {
				return fmt.Errorf("Error reading mac_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_id: %v", err)
		}
	}

	if err = d.Set("mac_id_v", flattenObjectSystemNpuNpuTcamSactMacIdV3rdl(o["mac-id-v"], d, "mac_id_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-id-v"], "ObjectSystemNpuNpuTcamSact-MacIdV"); ok {
			if err = d.Set("mac_id_v", vv); err != nil {
				return fmt.Errorf("Error reading mac_id_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_id_v: %v", err)
		}
	}

	if err = d.Set("mss", flattenObjectSystemNpuNpuTcamSactMss3rdl(o["mss"], d, "mss")); err != nil {
		if vv, ok := fortiAPIPatch(o["mss"], "ObjectSystemNpuNpuTcamSact-Mss"); ok {
			if err = d.Set("mss", vv); err != nil {
				return fmt.Errorf("Error reading mss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mss: %v", err)
		}
	}

	if err = d.Set("mss_v", flattenObjectSystemNpuNpuTcamSactMssV3rdl(o["mss-v"], d, "mss_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["mss-v"], "ObjectSystemNpuNpuTcamSact-MssV"); ok {
			if err = d.Set("mss_v", vv); err != nil {
				return fmt.Errorf("Error reading mss_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mss_v: %v", err)
		}
	}

	if err = d.Set("pleen", flattenObjectSystemNpuNpuTcamSactPleen3rdl(o["pleen"], d, "pleen")); err != nil {
		if vv, ok := fortiAPIPatch(o["pleen"], "ObjectSystemNpuNpuTcamSact-Pleen"); ok {
			if err = d.Set("pleen", vv); err != nil {
				return fmt.Errorf("Error reading pleen: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pleen: %v", err)
		}
	}

	if err = d.Set("pleen_v", flattenObjectSystemNpuNpuTcamSactPleenV3rdl(o["pleen-v"], d, "pleen_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["pleen-v"], "ObjectSystemNpuNpuTcamSact-PleenV"); ok {
			if err = d.Set("pleen_v", vv); err != nil {
				return fmt.Errorf("Error reading pleen_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pleen_v: %v", err)
		}
	}

	if err = d.Set("prio_pid", flattenObjectSystemNpuNpuTcamSactPrioPid3rdl(o["prio-pid"], d, "prio_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["prio-pid"], "ObjectSystemNpuNpuTcamSact-PrioPid"); ok {
			if err = d.Set("prio_pid", vv); err != nil {
				return fmt.Errorf("Error reading prio_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prio_pid: %v", err)
		}
	}

	if err = d.Set("prio_pid_v", flattenObjectSystemNpuNpuTcamSactPrioPidV3rdl(o["prio-pid-v"], d, "prio_pid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["prio-pid-v"], "ObjectSystemNpuNpuTcamSact-PrioPidV"); ok {
			if err = d.Set("prio_pid_v", vv); err != nil {
				return fmt.Errorf("Error reading prio_pid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prio_pid_v: %v", err)
		}
	}

	if err = d.Set("promis", flattenObjectSystemNpuNpuTcamSactPromis3rdl(o["promis"], d, "promis")); err != nil {
		if vv, ok := fortiAPIPatch(o["promis"], "ObjectSystemNpuNpuTcamSact-Promis"); ok {
			if err = d.Set("promis", vv); err != nil {
				return fmt.Errorf("Error reading promis: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading promis: %v", err)
		}
	}

	if err = d.Set("promis_v", flattenObjectSystemNpuNpuTcamSactPromisV3rdl(o["promis-v"], d, "promis_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["promis-v"], "ObjectSystemNpuNpuTcamSact-PromisV"); ok {
			if err = d.Set("promis_v", vv); err != nil {
				return fmt.Errorf("Error reading promis_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading promis_v: %v", err)
		}
	}

	if err = d.Set("rfsh", flattenObjectSystemNpuNpuTcamSactRfsh3rdl(o["rfsh"], d, "rfsh")); err != nil {
		if vv, ok := fortiAPIPatch(o["rfsh"], "ObjectSystemNpuNpuTcamSact-Rfsh"); ok {
			if err = d.Set("rfsh", vv); err != nil {
				return fmt.Errorf("Error reading rfsh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rfsh: %v", err)
		}
	}

	if err = d.Set("rfsh_v", flattenObjectSystemNpuNpuTcamSactRfshV3rdl(o["rfsh-v"], d, "rfsh_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["rfsh-v"], "ObjectSystemNpuNpuTcamSact-RfshV"); ok {
			if err = d.Set("rfsh_v", vv); err != nil {
				return fmt.Errorf("Error reading rfsh_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rfsh_v: %v", err)
		}
	}

	if err = d.Set("smac_skip", flattenObjectSystemNpuNpuTcamSactSmacSkip3rdl(o["smac-skip"], d, "smac_skip")); err != nil {
		if vv, ok := fortiAPIPatch(o["smac-skip"], "ObjectSystemNpuNpuTcamSact-SmacSkip"); ok {
			if err = d.Set("smac_skip", vv); err != nil {
				return fmt.Errorf("Error reading smac_skip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smac_skip: %v", err)
		}
	}

	if err = d.Set("smac_skip_v", flattenObjectSystemNpuNpuTcamSactSmacSkipV3rdl(o["smac-skip-v"], d, "smac_skip_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["smac-skip-v"], "ObjectSystemNpuNpuTcamSact-SmacSkipV"); ok {
			if err = d.Set("smac_skip_v", vv); err != nil {
				return fmt.Errorf("Error reading smac_skip_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading smac_skip_v: %v", err)
		}
	}

	if err = d.Set("tp_smchk_v", flattenObjectSystemNpuNpuTcamSactTpSmchkV3rdl(o["tp-smchk-v"], d, "tp_smchk_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tp-smchk-v"], "ObjectSystemNpuNpuTcamSact-TpSmchkV"); ok {
			if err = d.Set("tp_smchk_v", vv); err != nil {
				return fmt.Errorf("Error reading tp_smchk_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tp_smchk_v: %v", err)
		}
	}

	if err = d.Set("tp_smchk", flattenObjectSystemNpuNpuTcamSactTpSmchk3rdl(o["tp_smchk"], d, "tp_smchk")); err != nil {
		if vv, ok := fortiAPIPatch(o["tp_smchk"], "ObjectSystemNpuNpuTcamSact-TpSmchk"); ok {
			if err = d.Set("tp_smchk", vv); err != nil {
				return fmt.Errorf("Error reading tp_smchk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tp_smchk: %v", err)
		}
	}

	if err = d.Set("tpe_id", flattenObjectSystemNpuNpuTcamSactTpeId3rdl(o["tpe-id"], d, "tpe_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tpe-id"], "ObjectSystemNpuNpuTcamSact-TpeId"); ok {
			if err = d.Set("tpe_id", vv); err != nil {
				return fmt.Errorf("Error reading tpe_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tpe_id: %v", err)
		}
	}

	if err = d.Set("tpe_id_v", flattenObjectSystemNpuNpuTcamSactTpeIdV3rdl(o["tpe-id-v"], d, "tpe_id_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tpe-id-v"], "ObjectSystemNpuNpuTcamSact-TpeIdV"); ok {
			if err = d.Set("tpe_id_v", vv); err != nil {
				return fmt.Errorf("Error reading tpe_id_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tpe_id_v: %v", err)
		}
	}

	if err = d.Set("vdm", flattenObjectSystemNpuNpuTcamSactVdm3rdl(o["vdm"], d, "vdm")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdm"], "ObjectSystemNpuNpuTcamSact-Vdm"); ok {
			if err = d.Set("vdm", vv); err != nil {
				return fmt.Errorf("Error reading vdm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdm: %v", err)
		}
	}

	if err = d.Set("vdm_v", flattenObjectSystemNpuNpuTcamSactVdmV3rdl(o["vdm-v"], d, "vdm_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdm-v"], "ObjectSystemNpuNpuTcamSact-VdmV"); ok {
			if err = d.Set("vdm_v", vv); err != nil {
				return fmt.Errorf("Error reading vdm_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdm_v: %v", err)
		}
	}

	if err = d.Set("vdom_id", flattenObjectSystemNpuNpuTcamSactVdomId3rdl(o["vdom-id"], d, "vdom_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-id"], "ObjectSystemNpuNpuTcamSact-VdomId"); ok {
			if err = d.Set("vdom_id", vv); err != nil {
				return fmt.Errorf("Error reading vdom_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_id: %v", err)
		}
	}

	if err = d.Set("vdom_id_v", flattenObjectSystemNpuNpuTcamSactVdomIdV3rdl(o["vdom-id-v"], d, "vdom_id_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-id-v"], "ObjectSystemNpuNpuTcamSact-VdomIdV"); ok {
			if err = d.Set("vdom_id_v", vv); err != nil {
				return fmt.Errorf("Error reading vdom_id_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_id_v: %v", err)
		}
	}

	if err = d.Set("x_mode", flattenObjectSystemNpuNpuTcamSactXMode3rdl(o["x-mode"], d, "x_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["x-mode"], "ObjectSystemNpuNpuTcamSact-XMode"); ok {
			if err = d.Set("x_mode", vv); err != nil {
				return fmt.Errorf("Error reading x_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading x_mode: %v", err)
		}
	}

	if err = d.Set("x_mode_v", flattenObjectSystemNpuNpuTcamSactXModeV3rdl(o["x-mode-v"], d, "x_mode_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["x-mode-v"], "ObjectSystemNpuNpuTcamSact-XModeV"); ok {
			if err = d.Set("x_mode_v", vv); err != nil {
				return fmt.Errorf("Error reading x_mode_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading x_mode_v: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamSactFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamSactAct3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactActV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactBmproc3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactBmprocV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfLif3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfLifV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfr3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDfrV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDmacSkip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDmacSkipV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDosen3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactDosenV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEspffProc3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEspffProcV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEtypePid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactEtypePidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFragProc3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFragProcV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwd3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdLif3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdLifV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdTvid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdTvidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactFwdV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIcpen3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIcpenV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIgmpMldSnp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactIgmpMldSnpV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactLearn3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactLearnV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMSrhCtrl3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMSrhCtrlV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMacId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMacIdV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMss3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactMssV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPleen3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPleenV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPrioPid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPrioPidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPromis3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactPromisV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactRfsh3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactRfshV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactSmacSkip3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactSmacSkipV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpSmchkV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpSmchk3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactTpeIdV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdm3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdmV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdomId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactVdomIdV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactXMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamSactXModeV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcamSact(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("act"); ok || d.HasChange("act") {
		t, err := expandObjectSystemNpuNpuTcamSactAct3rdl(d, v, "act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["act"] = t
		}
	}

	if v, ok := d.GetOk("act_v"); ok || d.HasChange("act_v") {
		t, err := expandObjectSystemNpuNpuTcamSactActV3rdl(d, v, "act_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["act-v"] = t
		}
	}

	if v, ok := d.GetOk("bmproc"); ok || d.HasChange("bmproc") {
		t, err := expandObjectSystemNpuNpuTcamSactBmproc3rdl(d, v, "bmproc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bmproc"] = t
		}
	}

	if v, ok := d.GetOk("bmproc_v"); ok || d.HasChange("bmproc_v") {
		t, err := expandObjectSystemNpuNpuTcamSactBmprocV3rdl(d, v, "bmproc_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bmproc-v"] = t
		}
	}

	if v, ok := d.GetOk("df_lif"); ok || d.HasChange("df_lif") {
		t, err := expandObjectSystemNpuNpuTcamSactDfLif3rdl(d, v, "df_lif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["df-lif"] = t
		}
	}

	if v, ok := d.GetOk("df_lif_v"); ok || d.HasChange("df_lif_v") {
		t, err := expandObjectSystemNpuNpuTcamSactDfLifV3rdl(d, v, "df_lif_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["df-lif-v"] = t
		}
	}

	if v, ok := d.GetOk("dfr"); ok || d.HasChange("dfr") {
		t, err := expandObjectSystemNpuNpuTcamSactDfr3rdl(d, v, "dfr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dfr"] = t
		}
	}

	if v, ok := d.GetOk("dfr_v"); ok || d.HasChange("dfr_v") {
		t, err := expandObjectSystemNpuNpuTcamSactDfrV3rdl(d, v, "dfr_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dfr-v"] = t
		}
	}

	if v, ok := d.GetOk("dmac_skip"); ok || d.HasChange("dmac_skip") {
		t, err := expandObjectSystemNpuNpuTcamSactDmacSkip3rdl(d, v, "dmac_skip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dmac-skip"] = t
		}
	}

	if v, ok := d.GetOk("dmac_skip_v"); ok || d.HasChange("dmac_skip_v") {
		t, err := expandObjectSystemNpuNpuTcamSactDmacSkipV3rdl(d, v, "dmac_skip_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dmac-skip-v"] = t
		}
	}

	if v, ok := d.GetOk("dosen"); ok || d.HasChange("dosen") {
		t, err := expandObjectSystemNpuNpuTcamSactDosen3rdl(d, v, "dosen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dosen"] = t
		}
	}

	if v, ok := d.GetOk("dosen_v"); ok || d.HasChange("dosen_v") {
		t, err := expandObjectSystemNpuNpuTcamSactDosenV3rdl(d, v, "dosen_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dosen-v"] = t
		}
	}

	if v, ok := d.GetOk("espff_proc"); ok || d.HasChange("espff_proc") {
		t, err := expandObjectSystemNpuNpuTcamSactEspffProc3rdl(d, v, "espff_proc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["espff-proc"] = t
		}
	}

	if v, ok := d.GetOk("espff_proc_v"); ok || d.HasChange("espff_proc_v") {
		t, err := expandObjectSystemNpuNpuTcamSactEspffProcV3rdl(d, v, "espff_proc_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["espff-proc-v"] = t
		}
	}

	if v, ok := d.GetOk("etype_pid"); ok || d.HasChange("etype_pid") {
		t, err := expandObjectSystemNpuNpuTcamSactEtypePid3rdl(d, v, "etype_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["etype-pid"] = t
		}
	}

	if v, ok := d.GetOk("etype_pid_v"); ok || d.HasChange("etype_pid_v") {
		t, err := expandObjectSystemNpuNpuTcamSactEtypePidV3rdl(d, v, "etype_pid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["etype-pid-v"] = t
		}
	}

	if v, ok := d.GetOk("frag_proc"); ok || d.HasChange("frag_proc") {
		t, err := expandObjectSystemNpuNpuTcamSactFragProc3rdl(d, v, "frag_proc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frag-proc"] = t
		}
	}

	if v, ok := d.GetOk("frag_proc_v"); ok || d.HasChange("frag_proc_v") {
		t, err := expandObjectSystemNpuNpuTcamSactFragProcV3rdl(d, v, "frag_proc_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frag-proc-v"] = t
		}
	}

	if v, ok := d.GetOk("fwd"); ok || d.HasChange("fwd") {
		t, err := expandObjectSystemNpuNpuTcamSactFwd3rdl(d, v, "fwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd"] = t
		}
	}

	if v, ok := d.GetOk("fwd_lif"); ok || d.HasChange("fwd_lif") {
		t, err := expandObjectSystemNpuNpuTcamSactFwdLif3rdl(d, v, "fwd_lif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-lif"] = t
		}
	}

	if v, ok := d.GetOk("fwd_lif_v"); ok || d.HasChange("fwd_lif_v") {
		t, err := expandObjectSystemNpuNpuTcamSactFwdLifV3rdl(d, v, "fwd_lif_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-lif-v"] = t
		}
	}

	if v, ok := d.GetOk("fwd_tvid"); ok || d.HasChange("fwd_tvid") {
		t, err := expandObjectSystemNpuNpuTcamSactFwdTvid3rdl(d, v, "fwd_tvid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-tvid"] = t
		}
	}

	if v, ok := d.GetOk("fwd_tvid_v"); ok || d.HasChange("fwd_tvid_v") {
		t, err := expandObjectSystemNpuNpuTcamSactFwdTvidV3rdl(d, v, "fwd_tvid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-tvid-v"] = t
		}
	}

	if v, ok := d.GetOk("fwd_v"); ok || d.HasChange("fwd_v") {
		t, err := expandObjectSystemNpuNpuTcamSactFwdV3rdl(d, v, "fwd_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwd-v"] = t
		}
	}

	if v, ok := d.GetOk("icpen"); ok || d.HasChange("icpen") {
		t, err := expandObjectSystemNpuNpuTcamSactIcpen3rdl(d, v, "icpen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icpen"] = t
		}
	}

	if v, ok := d.GetOk("icpen_v"); ok || d.HasChange("icpen_v") {
		t, err := expandObjectSystemNpuNpuTcamSactIcpenV3rdl(d, v, "icpen_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icpen-v"] = t
		}
	}

	if v, ok := d.GetOk("igmp_mld_snp"); ok || d.HasChange("igmp_mld_snp") {
		t, err := expandObjectSystemNpuNpuTcamSactIgmpMldSnp3rdl(d, v, "igmp_mld_snp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-mld-snp"] = t
		}
	}

	if v, ok := d.GetOk("igmp_mld_snp_v"); ok || d.HasChange("igmp_mld_snp_v") {
		t, err := expandObjectSystemNpuNpuTcamSactIgmpMldSnpV3rdl(d, v, "igmp_mld_snp_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-mld-snp-v"] = t
		}
	}

	if v, ok := d.GetOk("learn"); ok || d.HasChange("learn") {
		t, err := expandObjectSystemNpuNpuTcamSactLearn3rdl(d, v, "learn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn"] = t
		}
	}

	if v, ok := d.GetOk("learn_v"); ok || d.HasChange("learn_v") {
		t, err := expandObjectSystemNpuNpuTcamSactLearnV3rdl(d, v, "learn_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-v"] = t
		}
	}

	if v, ok := d.GetOk("m_srh_ctrl"); ok || d.HasChange("m_srh_ctrl") {
		t, err := expandObjectSystemNpuNpuTcamSactMSrhCtrl3rdl(d, v, "m_srh_ctrl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["m-srh-ctrl"] = t
		}
	}

	if v, ok := d.GetOk("m_srh_ctrl_v"); ok || d.HasChange("m_srh_ctrl_v") {
		t, err := expandObjectSystemNpuNpuTcamSactMSrhCtrlV3rdl(d, v, "m_srh_ctrl_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["m-srh-ctrl-v"] = t
		}
	}

	if v, ok := d.GetOk("mac_id"); ok || d.HasChange("mac_id") {
		t, err := expandObjectSystemNpuNpuTcamSactMacId3rdl(d, v, "mac_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-id"] = t
		}
	}

	if v, ok := d.GetOk("mac_id_v"); ok || d.HasChange("mac_id_v") {
		t, err := expandObjectSystemNpuNpuTcamSactMacIdV3rdl(d, v, "mac_id_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-id-v"] = t
		}
	}

	if v, ok := d.GetOk("mss"); ok || d.HasChange("mss") {
		t, err := expandObjectSystemNpuNpuTcamSactMss3rdl(d, v, "mss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mss"] = t
		}
	}

	if v, ok := d.GetOk("mss_v"); ok || d.HasChange("mss_v") {
		t, err := expandObjectSystemNpuNpuTcamSactMssV3rdl(d, v, "mss_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mss-v"] = t
		}
	}

	if v, ok := d.GetOk("pleen"); ok || d.HasChange("pleen") {
		t, err := expandObjectSystemNpuNpuTcamSactPleen3rdl(d, v, "pleen")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pleen"] = t
		}
	}

	if v, ok := d.GetOk("pleen_v"); ok || d.HasChange("pleen_v") {
		t, err := expandObjectSystemNpuNpuTcamSactPleenV3rdl(d, v, "pleen_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pleen-v"] = t
		}
	}

	if v, ok := d.GetOk("prio_pid"); ok || d.HasChange("prio_pid") {
		t, err := expandObjectSystemNpuNpuTcamSactPrioPid3rdl(d, v, "prio_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prio-pid"] = t
		}
	}

	if v, ok := d.GetOk("prio_pid_v"); ok || d.HasChange("prio_pid_v") {
		t, err := expandObjectSystemNpuNpuTcamSactPrioPidV3rdl(d, v, "prio_pid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prio-pid-v"] = t
		}
	}

	if v, ok := d.GetOk("promis"); ok || d.HasChange("promis") {
		t, err := expandObjectSystemNpuNpuTcamSactPromis3rdl(d, v, "promis")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["promis"] = t
		}
	}

	if v, ok := d.GetOk("promis_v"); ok || d.HasChange("promis_v") {
		t, err := expandObjectSystemNpuNpuTcamSactPromisV3rdl(d, v, "promis_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["promis-v"] = t
		}
	}

	if v, ok := d.GetOk("rfsh"); ok || d.HasChange("rfsh") {
		t, err := expandObjectSystemNpuNpuTcamSactRfsh3rdl(d, v, "rfsh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfsh"] = t
		}
	}

	if v, ok := d.GetOk("rfsh_v"); ok || d.HasChange("rfsh_v") {
		t, err := expandObjectSystemNpuNpuTcamSactRfshV3rdl(d, v, "rfsh_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfsh-v"] = t
		}
	}

	if v, ok := d.GetOk("smac_skip"); ok || d.HasChange("smac_skip") {
		t, err := expandObjectSystemNpuNpuTcamSactSmacSkip3rdl(d, v, "smac_skip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smac-skip"] = t
		}
	}

	if v, ok := d.GetOk("smac_skip_v"); ok || d.HasChange("smac_skip_v") {
		t, err := expandObjectSystemNpuNpuTcamSactSmacSkipV3rdl(d, v, "smac_skip_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smac-skip-v"] = t
		}
	}

	if v, ok := d.GetOk("tp_smchk_v"); ok || d.HasChange("tp_smchk_v") {
		t, err := expandObjectSystemNpuNpuTcamSactTpSmchkV3rdl(d, v, "tp_smchk_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tp-smchk-v"] = t
		}
	}

	if v, ok := d.GetOk("tp_smchk"); ok || d.HasChange("tp_smchk") {
		t, err := expandObjectSystemNpuNpuTcamSactTpSmchk3rdl(d, v, "tp_smchk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tp_smchk"] = t
		}
	}

	if v, ok := d.GetOk("tpe_id"); ok || d.HasChange("tpe_id") {
		t, err := expandObjectSystemNpuNpuTcamSactTpeId3rdl(d, v, "tpe_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tpe-id"] = t
		}
	}

	if v, ok := d.GetOk("tpe_id_v"); ok || d.HasChange("tpe_id_v") {
		t, err := expandObjectSystemNpuNpuTcamSactTpeIdV3rdl(d, v, "tpe_id_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tpe-id-v"] = t
		}
	}

	if v, ok := d.GetOk("vdm"); ok || d.HasChange("vdm") {
		t, err := expandObjectSystemNpuNpuTcamSactVdm3rdl(d, v, "vdm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdm"] = t
		}
	}

	if v, ok := d.GetOk("vdm_v"); ok || d.HasChange("vdm_v") {
		t, err := expandObjectSystemNpuNpuTcamSactVdmV3rdl(d, v, "vdm_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdm-v"] = t
		}
	}

	if v, ok := d.GetOk("vdom_id"); ok || d.HasChange("vdom_id") {
		t, err := expandObjectSystemNpuNpuTcamSactVdomId3rdl(d, v, "vdom_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-id"] = t
		}
	}

	if v, ok := d.GetOk("vdom_id_v"); ok || d.HasChange("vdom_id_v") {
		t, err := expandObjectSystemNpuNpuTcamSactVdomIdV3rdl(d, v, "vdom_id_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-id-v"] = t
		}
	}

	if v, ok := d.GetOk("x_mode"); ok || d.HasChange("x_mode") {
		t, err := expandObjectSystemNpuNpuTcamSactXMode3rdl(d, v, "x_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["x-mode"] = t
		}
	}

	if v, ok := d.GetOk("x_mode_v"); ok || d.HasChange("x_mode_v") {
		t, err := expandObjectSystemNpuNpuTcamSactXModeV3rdl(d, v, "x_mode_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["x-mode-v"] = t
		}
	}

	return &obj, nil
}
