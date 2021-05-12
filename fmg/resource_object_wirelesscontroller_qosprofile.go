// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WiFi quality of service (QoS) profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerQosProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerQosProfileCreate,
		Read:   resourceObjectWirelessControllerQosProfileRead,
		Update: resourceObjectWirelessControllerQosProfileUpdate,
		Delete: resourceObjectWirelessControllerQosProfileDelete,

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
			"bandwidth_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"burst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_admission_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"call_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downlink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downlink_sta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_be": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_bk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_mapping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_vi": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"dscp_wmm_vo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"uplink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uplink_sta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wmm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_be_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wmm_bk_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wmm_dscp_marking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_uapsd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wmm_vi_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wmm_vo_dscp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerQosProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerQosProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerQosProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerQosProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerQosProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerQosProfileRead(d, m)
}

func resourceObjectWirelessControllerQosProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerQosProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerQosProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerQosProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerQosProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerQosProfileRead(d, m)
}

func resourceObjectWirelessControllerQosProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerQosProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerQosProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerQosProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerQosProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerQosProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerQosProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerQosProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerQosProfileBandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileBandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileCallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileCallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileDownlinkSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileDscpWmmBe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectWirelessControllerQosProfileDscpWmmBk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectWirelessControllerQosProfileDscpWmmMapping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileDscpWmmVi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectWirelessControllerQosProfileDscpWmmVo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectWirelessControllerQosProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileUplinkSta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileWmm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileWmmBeDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileWmmBkDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileWmmDscpMarking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileWmmUapsd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerQosProfileWmmViDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerQosProfileWmmVoDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerQosProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bandwidth_admission_control", flattenObjectWirelessControllerQosProfileBandwidthAdmissionControl(o["bandwidth-admission-control"], d, "bandwidth_admission_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-admission-control"], "ObjectWirelessControllerQosProfile-BandwidthAdmissionControl"); ok {
			if err = d.Set("bandwidth_admission_control", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_admission_control: %v", err)
		}
	}

	if err = d.Set("bandwidth_capacity", flattenObjectWirelessControllerQosProfileBandwidthCapacity(o["bandwidth-capacity"], d, "bandwidth_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-capacity"], "ObjectWirelessControllerQosProfile-BandwidthCapacity"); ok {
			if err = d.Set("bandwidth_capacity", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_capacity: %v", err)
		}
	}

	if err = d.Set("burst", flattenObjectWirelessControllerQosProfileBurst(o["burst"], d, "burst")); err != nil {
		if vv, ok := fortiAPIPatch(o["burst"], "ObjectWirelessControllerQosProfile-Burst"); ok {
			if err = d.Set("burst", vv); err != nil {
				return fmt.Errorf("Error reading burst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading burst: %v", err)
		}
	}

	if err = d.Set("call_admission_control", flattenObjectWirelessControllerQosProfileCallAdmissionControl(o["call-admission-control"], d, "call_admission_control")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-admission-control"], "ObjectWirelessControllerQosProfile-CallAdmissionControl"); ok {
			if err = d.Set("call_admission_control", vv); err != nil {
				return fmt.Errorf("Error reading call_admission_control: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_admission_control: %v", err)
		}
	}

	if err = d.Set("call_capacity", flattenObjectWirelessControllerQosProfileCallCapacity(o["call-capacity"], d, "call_capacity")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-capacity"], "ObjectWirelessControllerQosProfile-CallCapacity"); ok {
			if err = d.Set("call_capacity", vv); err != nil {
				return fmt.Errorf("Error reading call_capacity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_capacity: %v", err)
		}
	}

	if err = d.Set("comment", flattenObjectWirelessControllerQosProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerQosProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("downlink", flattenObjectWirelessControllerQosProfileDownlink(o["downlink"], d, "downlink")); err != nil {
		if vv, ok := fortiAPIPatch(o["downlink"], "ObjectWirelessControllerQosProfile-Downlink"); ok {
			if err = d.Set("downlink", vv); err != nil {
				return fmt.Errorf("Error reading downlink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downlink: %v", err)
		}
	}

	if err = d.Set("downlink_sta", flattenObjectWirelessControllerQosProfileDownlinkSta(o["downlink-sta"], d, "downlink_sta")); err != nil {
		if vv, ok := fortiAPIPatch(o["downlink-sta"], "ObjectWirelessControllerQosProfile-DownlinkSta"); ok {
			if err = d.Set("downlink_sta", vv); err != nil {
				return fmt.Errorf("Error reading downlink_sta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading downlink_sta: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_be", flattenObjectWirelessControllerQosProfileDscpWmmBe(o["dscp-wmm-be"], d, "dscp_wmm_be")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-be"], "ObjectWirelessControllerQosProfile-DscpWmmBe"); ok {
			if err = d.Set("dscp_wmm_be", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_be: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_bk", flattenObjectWirelessControllerQosProfileDscpWmmBk(o["dscp-wmm-bk"], d, "dscp_wmm_bk")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-bk"], "ObjectWirelessControllerQosProfile-DscpWmmBk"); ok {
			if err = d.Set("dscp_wmm_bk", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_bk: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_mapping", flattenObjectWirelessControllerQosProfileDscpWmmMapping(o["dscp-wmm-mapping"], d, "dscp_wmm_mapping")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-mapping"], "ObjectWirelessControllerQosProfile-DscpWmmMapping"); ok {
			if err = d.Set("dscp_wmm_mapping", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_mapping: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_mapping: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_vi", flattenObjectWirelessControllerQosProfileDscpWmmVi(o["dscp-wmm-vi"], d, "dscp_wmm_vi")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-vi"], "ObjectWirelessControllerQosProfile-DscpWmmVi"); ok {
			if err = d.Set("dscp_wmm_vi", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_vi: %v", err)
		}
	}

	if err = d.Set("dscp_wmm_vo", flattenObjectWirelessControllerQosProfileDscpWmmVo(o["dscp-wmm-vo"], d, "dscp_wmm_vo")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-wmm-vo"], "ObjectWirelessControllerQosProfile-DscpWmmVo"); ok {
			if err = d.Set("dscp_wmm_vo", vv); err != nil {
				return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_wmm_vo: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerQosProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerQosProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uplink", flattenObjectWirelessControllerQosProfileUplink(o["uplink"], d, "uplink")); err != nil {
		if vv, ok := fortiAPIPatch(o["uplink"], "ObjectWirelessControllerQosProfile-Uplink"); ok {
			if err = d.Set("uplink", vv); err != nil {
				return fmt.Errorf("Error reading uplink: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uplink: %v", err)
		}
	}

	if err = d.Set("uplink_sta", flattenObjectWirelessControllerQosProfileUplinkSta(o["uplink-sta"], d, "uplink_sta")); err != nil {
		if vv, ok := fortiAPIPatch(o["uplink-sta"], "ObjectWirelessControllerQosProfile-UplinkSta"); ok {
			if err = d.Set("uplink_sta", vv); err != nil {
				return fmt.Errorf("Error reading uplink_sta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uplink_sta: %v", err)
		}
	}

	if err = d.Set("wmm", flattenObjectWirelessControllerQosProfileWmm(o["wmm"], d, "wmm")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm"], "ObjectWirelessControllerQosProfile-Wmm"); ok {
			if err = d.Set("wmm", vv); err != nil {
				return fmt.Errorf("Error reading wmm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm: %v", err)
		}
	}

	if err = d.Set("wmm_be_dscp", flattenObjectWirelessControllerQosProfileWmmBeDscp(o["wmm-be-dscp"], d, "wmm_be_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-be-dscp"], "ObjectWirelessControllerQosProfile-WmmBeDscp"); ok {
			if err = d.Set("wmm_be_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_be_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_be_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_bk_dscp", flattenObjectWirelessControllerQosProfileWmmBkDscp(o["wmm-bk-dscp"], d, "wmm_bk_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-bk-dscp"], "ObjectWirelessControllerQosProfile-WmmBkDscp"); ok {
			if err = d.Set("wmm_bk_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_bk_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_bk_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_dscp_marking", flattenObjectWirelessControllerQosProfileWmmDscpMarking(o["wmm-dscp-marking"], d, "wmm_dscp_marking")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-dscp-marking"], "ObjectWirelessControllerQosProfile-WmmDscpMarking"); ok {
			if err = d.Set("wmm_dscp_marking", vv); err != nil {
				return fmt.Errorf("Error reading wmm_dscp_marking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_dscp_marking: %v", err)
		}
	}

	if err = d.Set("wmm_uapsd", flattenObjectWirelessControllerQosProfileWmmUapsd(o["wmm-uapsd"], d, "wmm_uapsd")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-uapsd"], "ObjectWirelessControllerQosProfile-WmmUapsd"); ok {
			if err = d.Set("wmm_uapsd", vv); err != nil {
				return fmt.Errorf("Error reading wmm_uapsd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_uapsd: %v", err)
		}
	}

	if err = d.Set("wmm_vi_dscp", flattenObjectWirelessControllerQosProfileWmmViDscp(o["wmm-vi-dscp"], d, "wmm_vi_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-vi-dscp"], "ObjectWirelessControllerQosProfile-WmmViDscp"); ok {
			if err = d.Set("wmm_vi_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_vi_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_vi_dscp: %v", err)
		}
	}

	if err = d.Set("wmm_vo_dscp", flattenObjectWirelessControllerQosProfileWmmVoDscp(o["wmm-vo-dscp"], d, "wmm_vo_dscp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wmm-vo-dscp"], "ObjectWirelessControllerQosProfile-WmmVoDscp"); ok {
			if err = d.Set("wmm_vo_dscp", vv); err != nil {
				return fmt.Errorf("Error reading wmm_vo_dscp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wmm_vo_dscp: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerQosProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerQosProfileBandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileBandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileCallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileCallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileDownlinkSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileDscpWmmBe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerQosProfileDscpWmmBk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerQosProfileDscpWmmMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileDscpWmmVi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerQosProfileDscpWmmVo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerQosProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileUplinkSta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmmBeDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmmBkDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmmDscpMarking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmmUapsd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmmViDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerQosProfileWmmVoDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerQosProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_admission_control"); ok {
		t, err := expandObjectWirelessControllerQosProfileBandwidthAdmissionControl(d, v, "bandwidth_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_capacity"); ok {
		t, err := expandObjectWirelessControllerQosProfileBandwidthCapacity(d, v, "bandwidth_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-capacity"] = t
		}
	}

	if v, ok := d.GetOk("burst"); ok {
		t, err := expandObjectWirelessControllerQosProfileBurst(d, v, "burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["burst"] = t
		}
	}

	if v, ok := d.GetOk("call_admission_control"); ok {
		t, err := expandObjectWirelessControllerQosProfileCallAdmissionControl(d, v, "call_admission_control")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-admission-control"] = t
		}
	}

	if v, ok := d.GetOk("call_capacity"); ok {
		t, err := expandObjectWirelessControllerQosProfileCallCapacity(d, v, "call_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-capacity"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectWirelessControllerQosProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("downlink"); ok {
		t, err := expandObjectWirelessControllerQosProfileDownlink(d, v, "downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink"] = t
		}
	}

	if v, ok := d.GetOk("downlink_sta"); ok {
		t, err := expandObjectWirelessControllerQosProfileDownlinkSta(d, v, "downlink_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-sta"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_be"); ok {
		t, err := expandObjectWirelessControllerQosProfileDscpWmmBe(d, v, "dscp_wmm_be")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-be"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_bk"); ok {
		t, err := expandObjectWirelessControllerQosProfileDscpWmmBk(d, v, "dscp_wmm_bk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-bk"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_mapping"); ok {
		t, err := expandObjectWirelessControllerQosProfileDscpWmmMapping(d, v, "dscp_wmm_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-mapping"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vi"); ok {
		t, err := expandObjectWirelessControllerQosProfileDscpWmmVi(d, v, "dscp_wmm_vi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vi"] = t
		}
	}

	if v, ok := d.GetOk("dscp_wmm_vo"); ok {
		t, err := expandObjectWirelessControllerQosProfileDscpWmmVo(d, v, "dscp_wmm_vo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-wmm-vo"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerQosProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uplink"); ok {
		t, err := expandObjectWirelessControllerQosProfileUplink(d, v, "uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink"] = t
		}
	}

	if v, ok := d.GetOk("uplink_sta"); ok {
		t, err := expandObjectWirelessControllerQosProfileUplinkSta(d, v, "uplink_sta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-sta"] = t
		}
	}

	if v, ok := d.GetOk("wmm"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmm(d, v, "wmm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm"] = t
		}
	}

	if v, ok := d.GetOk("wmm_be_dscp"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmmBeDscp(d, v, "wmm_be_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-be-dscp"] = t
		}
	}

	if v, ok := d.GetOk("wmm_bk_dscp"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmmBkDscp(d, v, "wmm_bk_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-bk-dscp"] = t
		}
	}

	if v, ok := d.GetOk("wmm_dscp_marking"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmmDscpMarking(d, v, "wmm_dscp_marking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOk("wmm_uapsd"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmmUapsd(d, v, "wmm_uapsd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-uapsd"] = t
		}
	}

	if v, ok := d.GetOk("wmm_vi_dscp"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmmViDscp(d, v, "wmm_vi_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-vi-dscp"] = t
		}
	}

	if v, ok := d.GetOk("wmm_vo_dscp"); ok {
		t, err := expandObjectWirelessControllerQosProfileWmmVoDscp(d, v, "wmm_vo_dscp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wmm-vo-dscp"] = t
		}
	}

	return &obj, nil
}
