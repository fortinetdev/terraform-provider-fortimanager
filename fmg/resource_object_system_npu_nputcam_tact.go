// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Target action of TCAM.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuNpuTcamTact() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuNpuTcamTactUpdate,
		Read:   resourceObjectSystemNpuNpuTcamTactRead,
		Update: resourceObjectSystemNpuNpuTcamTactUpdate,
		Delete: resourceObjectSystemNpuNpuTcamTactDelete,

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
	}
}

func resourceObjectSystemNpuNpuTcamTactUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSystemNpuNpuTcamTact(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamTact resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuNpuTcamTact(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuNpuTcamTact resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuNpuTcamTact")

	return resourceObjectSystemNpuNpuTcamTactRead(d, m)
}

func resourceObjectSystemNpuNpuTcamTactDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSystemNpuNpuTcamTact(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuNpuTcamTact resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuNpuTcamTactRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSystemNpuNpuTcamTact(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamTact resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuNpuTcamTact(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuNpuTcamTact resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuNpuTcamTactAct3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactActV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv4S3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv4SV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv6S3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactFmtuv6SV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactLnkid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactLnkidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMacId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMacIdV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMssT3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMssTV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv43rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv4V3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactMtuv6V3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSlifAct3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSlifActV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSublnkid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactSublnkidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTgtvAct3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTgtvActV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTlifAct3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTlifActV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTpeid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactTpeidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactV6Fe3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactV6FeV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepEnV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepSlid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepSlidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactVepEn3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltLif3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltLifV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltVid3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuNpuTcamTactXltVidV3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuNpuTcamTact(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("act", flattenObjectSystemNpuNpuTcamTactAct3rdl(o["act"], d, "act")); err != nil {
		if vv, ok := fortiAPIPatch(o["act"], "ObjectSystemNpuNpuTcamTact-Act"); ok {
			if err = d.Set("act", vv); err != nil {
				return fmt.Errorf("Error reading act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading act: %v", err)
		}
	}

	if err = d.Set("act_v", flattenObjectSystemNpuNpuTcamTactActV3rdl(o["act-v"], d, "act_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["act-v"], "ObjectSystemNpuNpuTcamTact-ActV"); ok {
			if err = d.Set("act_v", vv); err != nil {
				return fmt.Errorf("Error reading act_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading act_v: %v", err)
		}
	}

	if err = d.Set("fmtuv4_s", flattenObjectSystemNpuNpuTcamTactFmtuv4S3rdl(o["fmtuv4-s"], d, "fmtuv4_s")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmtuv4-s"], "ObjectSystemNpuNpuTcamTact-Fmtuv4S"); ok {
			if err = d.Set("fmtuv4_s", vv); err != nil {
				return fmt.Errorf("Error reading fmtuv4_s: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmtuv4_s: %v", err)
		}
	}

	if err = d.Set("fmtuv4_s_v", flattenObjectSystemNpuNpuTcamTactFmtuv4SV3rdl(o["fmtuv4-s-v"], d, "fmtuv4_s_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmtuv4-s-v"], "ObjectSystemNpuNpuTcamTact-Fmtuv4SV"); ok {
			if err = d.Set("fmtuv4_s_v", vv); err != nil {
				return fmt.Errorf("Error reading fmtuv4_s_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmtuv4_s_v: %v", err)
		}
	}

	if err = d.Set("fmtuv6_s", flattenObjectSystemNpuNpuTcamTactFmtuv6S3rdl(o["fmtuv6-s"], d, "fmtuv6_s")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmtuv6-s"], "ObjectSystemNpuNpuTcamTact-Fmtuv6S"); ok {
			if err = d.Set("fmtuv6_s", vv); err != nil {
				return fmt.Errorf("Error reading fmtuv6_s: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmtuv6_s: %v", err)
		}
	}

	if err = d.Set("fmtuv6_s_v", flattenObjectSystemNpuNpuTcamTactFmtuv6SV3rdl(o["fmtuv6-s-v"], d, "fmtuv6_s_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmtuv6-s-v"], "ObjectSystemNpuNpuTcamTact-Fmtuv6SV"); ok {
			if err = d.Set("fmtuv6_s_v", vv); err != nil {
				return fmt.Errorf("Error reading fmtuv6_s_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmtuv6_s_v: %v", err)
		}
	}

	if err = d.Set("lnkid", flattenObjectSystemNpuNpuTcamTactLnkid3rdl(o["lnkid"], d, "lnkid")); err != nil {
		if vv, ok := fortiAPIPatch(o["lnkid"], "ObjectSystemNpuNpuTcamTact-Lnkid"); ok {
			if err = d.Set("lnkid", vv); err != nil {
				return fmt.Errorf("Error reading lnkid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lnkid: %v", err)
		}
	}

	if err = d.Set("lnkid_v", flattenObjectSystemNpuNpuTcamTactLnkidV3rdl(o["lnkid-v"], d, "lnkid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["lnkid-v"], "ObjectSystemNpuNpuTcamTact-LnkidV"); ok {
			if err = d.Set("lnkid_v", vv); err != nil {
				return fmt.Errorf("Error reading lnkid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lnkid_v: %v", err)
		}
	}

	if err = d.Set("mac_id", flattenObjectSystemNpuNpuTcamTactMacId3rdl(o["mac-id"], d, "mac_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-id"], "ObjectSystemNpuNpuTcamTact-MacId"); ok {
			if err = d.Set("mac_id", vv); err != nil {
				return fmt.Errorf("Error reading mac_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_id: %v", err)
		}
	}

	if err = d.Set("mac_id_v", flattenObjectSystemNpuNpuTcamTactMacIdV3rdl(o["mac-id-v"], d, "mac_id_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-id-v"], "ObjectSystemNpuNpuTcamTact-MacIdV"); ok {
			if err = d.Set("mac_id_v", vv); err != nil {
				return fmt.Errorf("Error reading mac_id_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_id_v: %v", err)
		}
	}

	if err = d.Set("mss_t", flattenObjectSystemNpuNpuTcamTactMssT3rdl(o["mss-t"], d, "mss_t")); err != nil {
		if vv, ok := fortiAPIPatch(o["mss-t"], "ObjectSystemNpuNpuTcamTact-MssT"); ok {
			if err = d.Set("mss_t", vv); err != nil {
				return fmt.Errorf("Error reading mss_t: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mss_t: %v", err)
		}
	}

	if err = d.Set("mss_t_v", flattenObjectSystemNpuNpuTcamTactMssTV3rdl(o["mss-t-v"], d, "mss_t_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["mss-t-v"], "ObjectSystemNpuNpuTcamTact-MssTV"); ok {
			if err = d.Set("mss_t_v", vv); err != nil {
				return fmt.Errorf("Error reading mss_t_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mss_t_v: %v", err)
		}
	}

	if err = d.Set("mtuv4", flattenObjectSystemNpuNpuTcamTactMtuv43rdl(o["mtuv4"], d, "mtuv4")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtuv4"], "ObjectSystemNpuNpuTcamTact-Mtuv4"); ok {
			if err = d.Set("mtuv4", vv); err != nil {
				return fmt.Errorf("Error reading mtuv4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtuv4: %v", err)
		}
	}

	if err = d.Set("mtuv4_v", flattenObjectSystemNpuNpuTcamTactMtuv4V3rdl(o["mtuv4-v"], d, "mtuv4_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtuv4-v"], "ObjectSystemNpuNpuTcamTact-Mtuv4V"); ok {
			if err = d.Set("mtuv4_v", vv); err != nil {
				return fmt.Errorf("Error reading mtuv4_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtuv4_v: %v", err)
		}
	}

	if err = d.Set("mtuv6", flattenObjectSystemNpuNpuTcamTactMtuv63rdl(o["mtuv6"], d, "mtuv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtuv6"], "ObjectSystemNpuNpuTcamTact-Mtuv6"); ok {
			if err = d.Set("mtuv6", vv); err != nil {
				return fmt.Errorf("Error reading mtuv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtuv6: %v", err)
		}
	}

	if err = d.Set("mtuv6_v", flattenObjectSystemNpuNpuTcamTactMtuv6V3rdl(o["mtuv6-v"], d, "mtuv6_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtuv6-v"], "ObjectSystemNpuNpuTcamTact-Mtuv6V"); ok {
			if err = d.Set("mtuv6_v", vv); err != nil {
				return fmt.Errorf("Error reading mtuv6_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtuv6_v: %v", err)
		}
	}

	if err = d.Set("slif_act", flattenObjectSystemNpuNpuTcamTactSlifAct3rdl(o["slif-act"], d, "slif_act")); err != nil {
		if vv, ok := fortiAPIPatch(o["slif-act"], "ObjectSystemNpuNpuTcamTact-SlifAct"); ok {
			if err = d.Set("slif_act", vv); err != nil {
				return fmt.Errorf("Error reading slif_act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slif_act: %v", err)
		}
	}

	if err = d.Set("slif_act_v", flattenObjectSystemNpuNpuTcamTactSlifActV3rdl(o["slif-act-v"], d, "slif_act_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["slif-act-v"], "ObjectSystemNpuNpuTcamTact-SlifActV"); ok {
			if err = d.Set("slif_act_v", vv); err != nil {
				return fmt.Errorf("Error reading slif_act_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading slif_act_v: %v", err)
		}
	}

	if err = d.Set("sublnkid", flattenObjectSystemNpuNpuTcamTactSublnkid3rdl(o["sublnkid"], d, "sublnkid")); err != nil {
		if vv, ok := fortiAPIPatch(o["sublnkid"], "ObjectSystemNpuNpuTcamTact-Sublnkid"); ok {
			if err = d.Set("sublnkid", vv); err != nil {
				return fmt.Errorf("Error reading sublnkid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sublnkid: %v", err)
		}
	}

	if err = d.Set("sublnkid_v", flattenObjectSystemNpuNpuTcamTactSublnkidV3rdl(o["sublnkid-v"], d, "sublnkid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["sublnkid-v"], "ObjectSystemNpuNpuTcamTact-SublnkidV"); ok {
			if err = d.Set("sublnkid_v", vv); err != nil {
				return fmt.Errorf("Error reading sublnkid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sublnkid_v: %v", err)
		}
	}

	if err = d.Set("tgtv_act", flattenObjectSystemNpuNpuTcamTactTgtvAct3rdl(o["tgtv-act"], d, "tgtv_act")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgtv-act"], "ObjectSystemNpuNpuTcamTact-TgtvAct"); ok {
			if err = d.Set("tgtv_act", vv); err != nil {
				return fmt.Errorf("Error reading tgtv_act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgtv_act: %v", err)
		}
	}

	if err = d.Set("tgtv_act_v", flattenObjectSystemNpuNpuTcamTactTgtvActV3rdl(o["tgtv-act-v"], d, "tgtv_act_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgtv-act-v"], "ObjectSystemNpuNpuTcamTact-TgtvActV"); ok {
			if err = d.Set("tgtv_act_v", vv); err != nil {
				return fmt.Errorf("Error reading tgtv_act_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgtv_act_v: %v", err)
		}
	}

	if err = d.Set("tlif_act", flattenObjectSystemNpuNpuTcamTactTlifAct3rdl(o["tlif-act"], d, "tlif_act")); err != nil {
		if vv, ok := fortiAPIPatch(o["tlif-act"], "ObjectSystemNpuNpuTcamTact-TlifAct"); ok {
			if err = d.Set("tlif_act", vv); err != nil {
				return fmt.Errorf("Error reading tlif_act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tlif_act: %v", err)
		}
	}

	if err = d.Set("tlif_act_v", flattenObjectSystemNpuNpuTcamTactTlifActV3rdl(o["tlif-act-v"], d, "tlif_act_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tlif-act-v"], "ObjectSystemNpuNpuTcamTact-TlifActV"); ok {
			if err = d.Set("tlif_act_v", vv); err != nil {
				return fmt.Errorf("Error reading tlif_act_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tlif_act_v: %v", err)
		}
	}

	if err = d.Set("tpeid", flattenObjectSystemNpuNpuTcamTactTpeid3rdl(o["tpeid"], d, "tpeid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tpeid"], "ObjectSystemNpuNpuTcamTact-Tpeid"); ok {
			if err = d.Set("tpeid", vv); err != nil {
				return fmt.Errorf("Error reading tpeid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tpeid: %v", err)
		}
	}

	if err = d.Set("tpeid_v", flattenObjectSystemNpuNpuTcamTactTpeidV3rdl(o["tpeid-v"], d, "tpeid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["tpeid-v"], "ObjectSystemNpuNpuTcamTact-TpeidV"); ok {
			if err = d.Set("tpeid_v", vv); err != nil {
				return fmt.Errorf("Error reading tpeid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tpeid_v: %v", err)
		}
	}

	if err = d.Set("v6fe", flattenObjectSystemNpuNpuTcamTactV6Fe3rdl(o["v6fe"], d, "v6fe")); err != nil {
		if vv, ok := fortiAPIPatch(o["v6fe"], "ObjectSystemNpuNpuTcamTact-V6Fe"); ok {
			if err = d.Set("v6fe", vv); err != nil {
				return fmt.Errorf("Error reading v6fe: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading v6fe: %v", err)
		}
	}

	if err = d.Set("v6fe_v", flattenObjectSystemNpuNpuTcamTactV6FeV3rdl(o["v6fe-v"], d, "v6fe_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["v6fe-v"], "ObjectSystemNpuNpuTcamTact-V6FeV"); ok {
			if err = d.Set("v6fe_v", vv); err != nil {
				return fmt.Errorf("Error reading v6fe_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading v6fe_v: %v", err)
		}
	}

	if err = d.Set("vep_en_v", flattenObjectSystemNpuNpuTcamTactVepEnV3rdl(o["vep-en-v"], d, "vep_en_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["vep-en-v"], "ObjectSystemNpuNpuTcamTact-VepEnV"); ok {
			if err = d.Set("vep_en_v", vv); err != nil {
				return fmt.Errorf("Error reading vep_en_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vep_en_v: %v", err)
		}
	}

	if err = d.Set("vep_slid", flattenObjectSystemNpuNpuTcamTactVepSlid3rdl(o["vep-slid"], d, "vep_slid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vep-slid"], "ObjectSystemNpuNpuTcamTact-VepSlid"); ok {
			if err = d.Set("vep_slid", vv); err != nil {
				return fmt.Errorf("Error reading vep_slid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vep_slid: %v", err)
		}
	}

	if err = d.Set("vep_slid_v", flattenObjectSystemNpuNpuTcamTactVepSlidV3rdl(o["vep-slid-v"], d, "vep_slid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["vep-slid-v"], "ObjectSystemNpuNpuTcamTact-VepSlidV"); ok {
			if err = d.Set("vep_slid_v", vv); err != nil {
				return fmt.Errorf("Error reading vep_slid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vep_slid_v: %v", err)
		}
	}

	if err = d.Set("vep_en", flattenObjectSystemNpuNpuTcamTactVepEn3rdl(o["vep_en"], d, "vep_en")); err != nil {
		if vv, ok := fortiAPIPatch(o["vep_en"], "ObjectSystemNpuNpuTcamTact-VepEn"); ok {
			if err = d.Set("vep_en", vv); err != nil {
				return fmt.Errorf("Error reading vep_en: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vep_en: %v", err)
		}
	}

	if err = d.Set("xlt_lif", flattenObjectSystemNpuNpuTcamTactXltLif3rdl(o["xlt-lif"], d, "xlt_lif")); err != nil {
		if vv, ok := fortiAPIPatch(o["xlt-lif"], "ObjectSystemNpuNpuTcamTact-XltLif"); ok {
			if err = d.Set("xlt_lif", vv); err != nil {
				return fmt.Errorf("Error reading xlt_lif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xlt_lif: %v", err)
		}
	}

	if err = d.Set("xlt_lif_v", flattenObjectSystemNpuNpuTcamTactXltLifV3rdl(o["xlt-lif-v"], d, "xlt_lif_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["xlt-lif-v"], "ObjectSystemNpuNpuTcamTact-XltLifV"); ok {
			if err = d.Set("xlt_lif_v", vv); err != nil {
				return fmt.Errorf("Error reading xlt_lif_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xlt_lif_v: %v", err)
		}
	}

	if err = d.Set("xlt_vid", flattenObjectSystemNpuNpuTcamTactXltVid3rdl(o["xlt-vid"], d, "xlt_vid")); err != nil {
		if vv, ok := fortiAPIPatch(o["xlt-vid"], "ObjectSystemNpuNpuTcamTact-XltVid"); ok {
			if err = d.Set("xlt_vid", vv); err != nil {
				return fmt.Errorf("Error reading xlt_vid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xlt_vid: %v", err)
		}
	}

	if err = d.Set("xlt_vid_v", flattenObjectSystemNpuNpuTcamTactXltVidV3rdl(o["xlt-vid-v"], d, "xlt_vid_v")); err != nil {
		if vv, ok := fortiAPIPatch(o["xlt-vid-v"], "ObjectSystemNpuNpuTcamTact-XltVidV"); ok {
			if err = d.Set("xlt_vid_v", vv); err != nil {
				return fmt.Errorf("Error reading xlt_vid_v: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading xlt_vid_v: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuNpuTcamTactFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuNpuTcamTactAct3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactActV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv4S3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv4SV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv6S3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactFmtuv6SV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactLnkid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactLnkidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMacId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMacIdV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMssT3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMssTV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv43rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv4V3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactMtuv6V3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSlifAct3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSlifActV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSublnkid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactSublnkidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTgtvAct3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTgtvActV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTlifAct3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTlifActV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTpeid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactTpeidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactV6Fe3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactV6FeV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepEnV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepSlid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepSlidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactVepEn3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltLif3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltLifV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltVid3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuNpuTcamTactXltVidV3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuNpuTcamTact(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("act"); ok || d.HasChange("act") {
		t, err := expandObjectSystemNpuNpuTcamTactAct3rdl(d, v, "act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["act"] = t
		}
	}

	if v, ok := d.GetOk("act_v"); ok || d.HasChange("act_v") {
		t, err := expandObjectSystemNpuNpuTcamTactActV3rdl(d, v, "act_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["act-v"] = t
		}
	}

	if v, ok := d.GetOk("fmtuv4_s"); ok || d.HasChange("fmtuv4_s") {
		t, err := expandObjectSystemNpuNpuTcamTactFmtuv4S3rdl(d, v, "fmtuv4_s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmtuv4-s"] = t
		}
	}

	if v, ok := d.GetOk("fmtuv4_s_v"); ok || d.HasChange("fmtuv4_s_v") {
		t, err := expandObjectSystemNpuNpuTcamTactFmtuv4SV3rdl(d, v, "fmtuv4_s_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmtuv4-s-v"] = t
		}
	}

	if v, ok := d.GetOk("fmtuv6_s"); ok || d.HasChange("fmtuv6_s") {
		t, err := expandObjectSystemNpuNpuTcamTactFmtuv6S3rdl(d, v, "fmtuv6_s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmtuv6-s"] = t
		}
	}

	if v, ok := d.GetOk("fmtuv6_s_v"); ok || d.HasChange("fmtuv6_s_v") {
		t, err := expandObjectSystemNpuNpuTcamTactFmtuv6SV3rdl(d, v, "fmtuv6_s_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmtuv6-s-v"] = t
		}
	}

	if v, ok := d.GetOk("lnkid"); ok || d.HasChange("lnkid") {
		t, err := expandObjectSystemNpuNpuTcamTactLnkid3rdl(d, v, "lnkid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lnkid"] = t
		}
	}

	if v, ok := d.GetOk("lnkid_v"); ok || d.HasChange("lnkid_v") {
		t, err := expandObjectSystemNpuNpuTcamTactLnkidV3rdl(d, v, "lnkid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lnkid-v"] = t
		}
	}

	if v, ok := d.GetOk("mac_id"); ok || d.HasChange("mac_id") {
		t, err := expandObjectSystemNpuNpuTcamTactMacId3rdl(d, v, "mac_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-id"] = t
		}
	}

	if v, ok := d.GetOk("mac_id_v"); ok || d.HasChange("mac_id_v") {
		t, err := expandObjectSystemNpuNpuTcamTactMacIdV3rdl(d, v, "mac_id_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-id-v"] = t
		}
	}

	if v, ok := d.GetOk("mss_t"); ok || d.HasChange("mss_t") {
		t, err := expandObjectSystemNpuNpuTcamTactMssT3rdl(d, v, "mss_t")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mss-t"] = t
		}
	}

	if v, ok := d.GetOk("mss_t_v"); ok || d.HasChange("mss_t_v") {
		t, err := expandObjectSystemNpuNpuTcamTactMssTV3rdl(d, v, "mss_t_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mss-t-v"] = t
		}
	}

	if v, ok := d.GetOk("mtuv4"); ok || d.HasChange("mtuv4") {
		t, err := expandObjectSystemNpuNpuTcamTactMtuv43rdl(d, v, "mtuv4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtuv4"] = t
		}
	}

	if v, ok := d.GetOk("mtuv4_v"); ok || d.HasChange("mtuv4_v") {
		t, err := expandObjectSystemNpuNpuTcamTactMtuv4V3rdl(d, v, "mtuv4_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtuv4-v"] = t
		}
	}

	if v, ok := d.GetOk("mtuv6"); ok || d.HasChange("mtuv6") {
		t, err := expandObjectSystemNpuNpuTcamTactMtuv63rdl(d, v, "mtuv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtuv6"] = t
		}
	}

	if v, ok := d.GetOk("mtuv6_v"); ok || d.HasChange("mtuv6_v") {
		t, err := expandObjectSystemNpuNpuTcamTactMtuv6V3rdl(d, v, "mtuv6_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtuv6-v"] = t
		}
	}

	if v, ok := d.GetOk("slif_act"); ok || d.HasChange("slif_act") {
		t, err := expandObjectSystemNpuNpuTcamTactSlifAct3rdl(d, v, "slif_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slif-act"] = t
		}
	}

	if v, ok := d.GetOk("slif_act_v"); ok || d.HasChange("slif_act_v") {
		t, err := expandObjectSystemNpuNpuTcamTactSlifActV3rdl(d, v, "slif_act_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slif-act-v"] = t
		}
	}

	if v, ok := d.GetOk("sublnkid"); ok || d.HasChange("sublnkid") {
		t, err := expandObjectSystemNpuNpuTcamTactSublnkid3rdl(d, v, "sublnkid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sublnkid"] = t
		}
	}

	if v, ok := d.GetOk("sublnkid_v"); ok || d.HasChange("sublnkid_v") {
		t, err := expandObjectSystemNpuNpuTcamTactSublnkidV3rdl(d, v, "sublnkid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sublnkid-v"] = t
		}
	}

	if v, ok := d.GetOk("tgtv_act"); ok || d.HasChange("tgtv_act") {
		t, err := expandObjectSystemNpuNpuTcamTactTgtvAct3rdl(d, v, "tgtv_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgtv-act"] = t
		}
	}

	if v, ok := d.GetOk("tgtv_act_v"); ok || d.HasChange("tgtv_act_v") {
		t, err := expandObjectSystemNpuNpuTcamTactTgtvActV3rdl(d, v, "tgtv_act_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgtv-act-v"] = t
		}
	}

	if v, ok := d.GetOk("tlif_act"); ok || d.HasChange("tlif_act") {
		t, err := expandObjectSystemNpuNpuTcamTactTlifAct3rdl(d, v, "tlif_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tlif-act"] = t
		}
	}

	if v, ok := d.GetOk("tlif_act_v"); ok || d.HasChange("tlif_act_v") {
		t, err := expandObjectSystemNpuNpuTcamTactTlifActV3rdl(d, v, "tlif_act_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tlif-act-v"] = t
		}
	}

	if v, ok := d.GetOk("tpeid"); ok || d.HasChange("tpeid") {
		t, err := expandObjectSystemNpuNpuTcamTactTpeid3rdl(d, v, "tpeid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tpeid"] = t
		}
	}

	if v, ok := d.GetOk("tpeid_v"); ok || d.HasChange("tpeid_v") {
		t, err := expandObjectSystemNpuNpuTcamTactTpeidV3rdl(d, v, "tpeid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tpeid-v"] = t
		}
	}

	if v, ok := d.GetOk("v6fe"); ok || d.HasChange("v6fe") {
		t, err := expandObjectSystemNpuNpuTcamTactV6Fe3rdl(d, v, "v6fe")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["v6fe"] = t
		}
	}

	if v, ok := d.GetOk("v6fe_v"); ok || d.HasChange("v6fe_v") {
		t, err := expandObjectSystemNpuNpuTcamTactV6FeV3rdl(d, v, "v6fe_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["v6fe-v"] = t
		}
	}

	if v, ok := d.GetOk("vep_en_v"); ok || d.HasChange("vep_en_v") {
		t, err := expandObjectSystemNpuNpuTcamTactVepEnV3rdl(d, v, "vep_en_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vep-en-v"] = t
		}
	}

	if v, ok := d.GetOk("vep_slid"); ok || d.HasChange("vep_slid") {
		t, err := expandObjectSystemNpuNpuTcamTactVepSlid3rdl(d, v, "vep_slid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vep-slid"] = t
		}
	}

	if v, ok := d.GetOk("vep_slid_v"); ok || d.HasChange("vep_slid_v") {
		t, err := expandObjectSystemNpuNpuTcamTactVepSlidV3rdl(d, v, "vep_slid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vep-slid-v"] = t
		}
	}

	if v, ok := d.GetOk("vep_en"); ok || d.HasChange("vep_en") {
		t, err := expandObjectSystemNpuNpuTcamTactVepEn3rdl(d, v, "vep_en")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vep_en"] = t
		}
	}

	if v, ok := d.GetOk("xlt_lif"); ok || d.HasChange("xlt_lif") {
		t, err := expandObjectSystemNpuNpuTcamTactXltLif3rdl(d, v, "xlt_lif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xlt-lif"] = t
		}
	}

	if v, ok := d.GetOk("xlt_lif_v"); ok || d.HasChange("xlt_lif_v") {
		t, err := expandObjectSystemNpuNpuTcamTactXltLifV3rdl(d, v, "xlt_lif_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xlt-lif-v"] = t
		}
	}

	if v, ok := d.GetOk("xlt_vid"); ok || d.HasChange("xlt_vid") {
		t, err := expandObjectSystemNpuNpuTcamTactXltVid3rdl(d, v, "xlt_vid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xlt-vid"] = t
		}
	}

	if v, ok := d.GetOk("xlt_vid_v"); ok || d.HasChange("xlt_vid_v") {
		t, err := expandObjectSystemNpuNpuTcamTactXltVidV3rdl(d, v, "xlt_vid_v")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xlt-vid-v"] = t
		}
	}

	return &obj, nil
}
