// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: RAT timeout profile

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectGtpRatTimeoutProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectGtpRatTimeoutProfileCreate,
		Read:   resourceObjectGtpRatTimeoutProfileRead,
		Update: resourceObjectGtpRatTimeoutProfileUpdate,
		Delete: resourceObjectGtpRatTimeoutProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"eutran_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gan_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"geran_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hspa_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ltem_geo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ltem_leo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ltem_meo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ltem_othersat_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ltem_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nbiot_geo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nbiot_leo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nbiot_meo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nbiot_othersat_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nbiot_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nr_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"utran_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wb_eutran_geo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wb_eutran_leo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wb_eutran_meo_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wb_eutran_othersat_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wlan_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectGtpRatTimeoutProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectGtpRatTimeoutProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectGtpRatTimeoutProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectGtpRatTimeoutProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectGtpRatTimeoutProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectGtpRatTimeoutProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectGtpRatTimeoutProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectGtpRatTimeoutProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpRatTimeoutProfileRead(d, m)
}

func resourceObjectGtpRatTimeoutProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectGtpRatTimeoutProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpRatTimeoutProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectGtpRatTimeoutProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectGtpRatTimeoutProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectGtpRatTimeoutProfileRead(d, m)
}

func resourceObjectGtpRatTimeoutProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectGtpRatTimeoutProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectGtpRatTimeoutProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectGtpRatTimeoutProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectGtpRatTimeoutProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectGtpRatTimeoutProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectGtpRatTimeoutProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectGtpRatTimeoutProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectGtpRatTimeoutProfileEutranTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileGanTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileGeranTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileHspaTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileLtemGeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileLtemLeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileLtemMeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileLtemOthersatTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileLtemTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileNbiotGeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileNbiotLeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileNbiotMeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileNbiotOthersatTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileNbiotTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileNrTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileUtranTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileVirtualTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileWbEutranGeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileWbEutranLeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileWbEutranMeoTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileWbEutranOthersatTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectGtpRatTimeoutProfileWlanTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectGtpRatTimeoutProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("eutran_timeout", flattenObjectGtpRatTimeoutProfileEutranTimeout(o["eutran-timeout"], d, "eutran_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["eutran-timeout"], "ObjectGtpRatTimeoutProfile-EutranTimeout"); ok {
			if err = d.Set("eutran_timeout", vv); err != nil {
				return fmt.Errorf("Error reading eutran_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eutran_timeout: %v", err)
		}
	}

	if err = d.Set("gan_timeout", flattenObjectGtpRatTimeoutProfileGanTimeout(o["gan-timeout"], d, "gan_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["gan-timeout"], "ObjectGtpRatTimeoutProfile-GanTimeout"); ok {
			if err = d.Set("gan_timeout", vv); err != nil {
				return fmt.Errorf("Error reading gan_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gan_timeout: %v", err)
		}
	}

	if err = d.Set("geran_timeout", flattenObjectGtpRatTimeoutProfileGeranTimeout(o["geran-timeout"], d, "geran_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["geran-timeout"], "ObjectGtpRatTimeoutProfile-GeranTimeout"); ok {
			if err = d.Set("geran_timeout", vv); err != nil {
				return fmt.Errorf("Error reading geran_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading geran_timeout: %v", err)
		}
	}

	if err = d.Set("hspa_timeout", flattenObjectGtpRatTimeoutProfileHspaTimeout(o["hspa-timeout"], d, "hspa_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["hspa-timeout"], "ObjectGtpRatTimeoutProfile-HspaTimeout"); ok {
			if err = d.Set("hspa_timeout", vv); err != nil {
				return fmt.Errorf("Error reading hspa_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hspa_timeout: %v", err)
		}
	}

	if err = d.Set("ltem_geo_timeout", flattenObjectGtpRatTimeoutProfileLtemGeoTimeout(o["ltem-geo-timeout"], d, "ltem_geo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltem-geo-timeout"], "ObjectGtpRatTimeoutProfile-LtemGeoTimeout"); ok {
			if err = d.Set("ltem_geo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ltem_geo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltem_geo_timeout: %v", err)
		}
	}

	if err = d.Set("ltem_leo_timeout", flattenObjectGtpRatTimeoutProfileLtemLeoTimeout(o["ltem-leo-timeout"], d, "ltem_leo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltem-leo-timeout"], "ObjectGtpRatTimeoutProfile-LtemLeoTimeout"); ok {
			if err = d.Set("ltem_leo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ltem_leo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltem_leo_timeout: %v", err)
		}
	}

	if err = d.Set("ltem_meo_timeout", flattenObjectGtpRatTimeoutProfileLtemMeoTimeout(o["ltem-meo-timeout"], d, "ltem_meo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltem-meo-timeout"], "ObjectGtpRatTimeoutProfile-LtemMeoTimeout"); ok {
			if err = d.Set("ltem_meo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ltem_meo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltem_meo_timeout: %v", err)
		}
	}

	if err = d.Set("ltem_othersat_timeout", flattenObjectGtpRatTimeoutProfileLtemOthersatTimeout(o["ltem-othersat-timeout"], d, "ltem_othersat_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltem-othersat-timeout"], "ObjectGtpRatTimeoutProfile-LtemOthersatTimeout"); ok {
			if err = d.Set("ltem_othersat_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ltem_othersat_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltem_othersat_timeout: %v", err)
		}
	}

	if err = d.Set("ltem_timeout", flattenObjectGtpRatTimeoutProfileLtemTimeout(o["ltem-timeout"], d, "ltem_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ltem-timeout"], "ObjectGtpRatTimeoutProfile-LtemTimeout"); ok {
			if err = d.Set("ltem_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ltem_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ltem_timeout: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectGtpRatTimeoutProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectGtpRatTimeoutProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nbiot_geo_timeout", flattenObjectGtpRatTimeoutProfileNbiotGeoTimeout(o["nbiot-geo-timeout"], d, "nbiot_geo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["nbiot-geo-timeout"], "ObjectGtpRatTimeoutProfile-NbiotGeoTimeout"); ok {
			if err = d.Set("nbiot_geo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading nbiot_geo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nbiot_geo_timeout: %v", err)
		}
	}

	if err = d.Set("nbiot_leo_timeout", flattenObjectGtpRatTimeoutProfileNbiotLeoTimeout(o["nbiot-leo-timeout"], d, "nbiot_leo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["nbiot-leo-timeout"], "ObjectGtpRatTimeoutProfile-NbiotLeoTimeout"); ok {
			if err = d.Set("nbiot_leo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading nbiot_leo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nbiot_leo_timeout: %v", err)
		}
	}

	if err = d.Set("nbiot_meo_timeout", flattenObjectGtpRatTimeoutProfileNbiotMeoTimeout(o["nbiot-meo-timeout"], d, "nbiot_meo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["nbiot-meo-timeout"], "ObjectGtpRatTimeoutProfile-NbiotMeoTimeout"); ok {
			if err = d.Set("nbiot_meo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading nbiot_meo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nbiot_meo_timeout: %v", err)
		}
	}

	if err = d.Set("nbiot_othersat_timeout", flattenObjectGtpRatTimeoutProfileNbiotOthersatTimeout(o["nbiot-othersat-timeout"], d, "nbiot_othersat_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["nbiot-othersat-timeout"], "ObjectGtpRatTimeoutProfile-NbiotOthersatTimeout"); ok {
			if err = d.Set("nbiot_othersat_timeout", vv); err != nil {
				return fmt.Errorf("Error reading nbiot_othersat_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nbiot_othersat_timeout: %v", err)
		}
	}

	if err = d.Set("nbiot_timeout", flattenObjectGtpRatTimeoutProfileNbiotTimeout(o["nbiot-timeout"], d, "nbiot_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["nbiot-timeout"], "ObjectGtpRatTimeoutProfile-NbiotTimeout"); ok {
			if err = d.Set("nbiot_timeout", vv); err != nil {
				return fmt.Errorf("Error reading nbiot_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nbiot_timeout: %v", err)
		}
	}

	if err = d.Set("nr_timeout", flattenObjectGtpRatTimeoutProfileNrTimeout(o["nr-timeout"], d, "nr_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["nr-timeout"], "ObjectGtpRatTimeoutProfile-NrTimeout"); ok {
			if err = d.Set("nr_timeout", vv); err != nil {
				return fmt.Errorf("Error reading nr_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nr_timeout: %v", err)
		}
	}

	if err = d.Set("utran_timeout", flattenObjectGtpRatTimeoutProfileUtranTimeout(o["utran-timeout"], d, "utran_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["utran-timeout"], "ObjectGtpRatTimeoutProfile-UtranTimeout"); ok {
			if err = d.Set("utran_timeout", vv); err != nil {
				return fmt.Errorf("Error reading utran_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utran_timeout: %v", err)
		}
	}

	if err = d.Set("virtual_timeout", flattenObjectGtpRatTimeoutProfileVirtualTimeout(o["virtual-timeout"], d, "virtual_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-timeout"], "ObjectGtpRatTimeoutProfile-VirtualTimeout"); ok {
			if err = d.Set("virtual_timeout", vv); err != nil {
				return fmt.Errorf("Error reading virtual_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_timeout: %v", err)
		}
	}

	if err = d.Set("wb_eutran_geo_timeout", flattenObjectGtpRatTimeoutProfileWbEutranGeoTimeout(o["wb-eutran-geo-timeout"], d, "wb_eutran_geo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["wb-eutran-geo-timeout"], "ObjectGtpRatTimeoutProfile-WbEutranGeoTimeout"); ok {
			if err = d.Set("wb_eutran_geo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading wb_eutran_geo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wb_eutran_geo_timeout: %v", err)
		}
	}

	if err = d.Set("wb_eutran_leo_timeout", flattenObjectGtpRatTimeoutProfileWbEutranLeoTimeout(o["wb-eutran-leo-timeout"], d, "wb_eutran_leo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["wb-eutran-leo-timeout"], "ObjectGtpRatTimeoutProfile-WbEutranLeoTimeout"); ok {
			if err = d.Set("wb_eutran_leo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading wb_eutran_leo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wb_eutran_leo_timeout: %v", err)
		}
	}

	if err = d.Set("wb_eutran_meo_timeout", flattenObjectGtpRatTimeoutProfileWbEutranMeoTimeout(o["wb-eutran-meo-timeout"], d, "wb_eutran_meo_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["wb-eutran-meo-timeout"], "ObjectGtpRatTimeoutProfile-WbEutranMeoTimeout"); ok {
			if err = d.Set("wb_eutran_meo_timeout", vv); err != nil {
				return fmt.Errorf("Error reading wb_eutran_meo_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wb_eutran_meo_timeout: %v", err)
		}
	}

	if err = d.Set("wb_eutran_othersat_timeout", flattenObjectGtpRatTimeoutProfileWbEutranOthersatTimeout(o["wb-eutran-othersat-timeout"], d, "wb_eutran_othersat_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["wb-eutran-othersat-timeout"], "ObjectGtpRatTimeoutProfile-WbEutranOthersatTimeout"); ok {
			if err = d.Set("wb_eutran_othersat_timeout", vv); err != nil {
				return fmt.Errorf("Error reading wb_eutran_othersat_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wb_eutran_othersat_timeout: %v", err)
		}
	}

	if err = d.Set("wlan_timeout", flattenObjectGtpRatTimeoutProfileWlanTimeout(o["wlan-timeout"], d, "wlan_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["wlan-timeout"], "ObjectGtpRatTimeoutProfile-WlanTimeout"); ok {
			if err = d.Set("wlan_timeout", vv); err != nil {
				return fmt.Errorf("Error reading wlan_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wlan_timeout: %v", err)
		}
	}

	return nil
}

func flattenObjectGtpRatTimeoutProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectGtpRatTimeoutProfileEutranTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileGanTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileGeranTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileHspaTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileLtemGeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileLtemLeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileLtemMeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileLtemOthersatTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileLtemTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileNbiotGeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileNbiotLeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileNbiotMeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileNbiotOthersatTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileNbiotTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileNrTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileUtranTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileVirtualTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileWbEutranGeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileWbEutranLeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileWbEutranMeoTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileWbEutranOthersatTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectGtpRatTimeoutProfileWlanTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectGtpRatTimeoutProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("eutran_timeout"); ok || d.HasChange("eutran_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileEutranTimeout(d, v, "eutran_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eutran-timeout"] = t
		}
	}

	if v, ok := d.GetOk("gan_timeout"); ok || d.HasChange("gan_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileGanTimeout(d, v, "gan_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gan-timeout"] = t
		}
	}

	if v, ok := d.GetOk("geran_timeout"); ok || d.HasChange("geran_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileGeranTimeout(d, v, "geran_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geran-timeout"] = t
		}
	}

	if v, ok := d.GetOk("hspa_timeout"); ok || d.HasChange("hspa_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileHspaTimeout(d, v, "hspa_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hspa-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ltem_geo_timeout"); ok || d.HasChange("ltem_geo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileLtemGeoTimeout(d, v, "ltem_geo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltem-geo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ltem_leo_timeout"); ok || d.HasChange("ltem_leo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileLtemLeoTimeout(d, v, "ltem_leo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltem-leo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ltem_meo_timeout"); ok || d.HasChange("ltem_meo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileLtemMeoTimeout(d, v, "ltem_meo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltem-meo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ltem_othersat_timeout"); ok || d.HasChange("ltem_othersat_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileLtemOthersatTimeout(d, v, "ltem_othersat_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltem-othersat-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ltem_timeout"); ok || d.HasChange("ltem_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileLtemTimeout(d, v, "ltem_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ltem-timeout"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectGtpRatTimeoutProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nbiot_geo_timeout"); ok || d.HasChange("nbiot_geo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileNbiotGeoTimeout(d, v, "nbiot_geo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nbiot-geo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("nbiot_leo_timeout"); ok || d.HasChange("nbiot_leo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileNbiotLeoTimeout(d, v, "nbiot_leo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nbiot-leo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("nbiot_meo_timeout"); ok || d.HasChange("nbiot_meo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileNbiotMeoTimeout(d, v, "nbiot_meo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nbiot-meo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("nbiot_othersat_timeout"); ok || d.HasChange("nbiot_othersat_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileNbiotOthersatTimeout(d, v, "nbiot_othersat_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nbiot-othersat-timeout"] = t
		}
	}

	if v, ok := d.GetOk("nbiot_timeout"); ok || d.HasChange("nbiot_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileNbiotTimeout(d, v, "nbiot_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nbiot-timeout"] = t
		}
	}

	if v, ok := d.GetOk("nr_timeout"); ok || d.HasChange("nr_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileNrTimeout(d, v, "nr_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nr-timeout"] = t
		}
	}

	if v, ok := d.GetOk("utran_timeout"); ok || d.HasChange("utran_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileUtranTimeout(d, v, "utran_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utran-timeout"] = t
		}
	}

	if v, ok := d.GetOk("virtual_timeout"); ok || d.HasChange("virtual_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileVirtualTimeout(d, v, "virtual_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-timeout"] = t
		}
	}

	if v, ok := d.GetOk("wb_eutran_geo_timeout"); ok || d.HasChange("wb_eutran_geo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileWbEutranGeoTimeout(d, v, "wb_eutran_geo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wb-eutran-geo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("wb_eutran_leo_timeout"); ok || d.HasChange("wb_eutran_leo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileWbEutranLeoTimeout(d, v, "wb_eutran_leo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wb-eutran-leo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("wb_eutran_meo_timeout"); ok || d.HasChange("wb_eutran_meo_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileWbEutranMeoTimeout(d, v, "wb_eutran_meo_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wb-eutran-meo-timeout"] = t
		}
	}

	if v, ok := d.GetOk("wb_eutran_othersat_timeout"); ok || d.HasChange("wb_eutran_othersat_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileWbEutranOthersatTimeout(d, v, "wb_eutran_othersat_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wb-eutran-othersat-timeout"] = t
		}
	}

	if v, ok := d.GetOk("wlan_timeout"); ok || d.HasChange("wlan_timeout") {
		t, err := expandObjectGtpRatTimeoutProfileWlanTimeout(d, v, "wlan_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wlan-timeout"] = t
		}
	}

	return &obj, nil
}
