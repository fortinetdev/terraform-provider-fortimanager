// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AOC list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListDelete,

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
			"h2qp_advice_of_charge": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nai_realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nai_realm_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"plan_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"currency": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"info_file": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lang": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListCreate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListUpdate(d *schema.ResourceData, m interface{}) error {
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

	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListDelete(d *schema.ResourceData, m interface{}) error {
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

	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge

	wsParams["adom"] = adomv

	err = c.DeleteObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListRead(d *schema.ResourceData, m interface{}) error {
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

	h2qp_advice_of_charge := d.Get("h2qp_advice_of_charge").(string)
	if h2qp_advice_of_charge == "" {
		h2qp_advice_of_charge = importOptionChecking(m.(*FortiClient).Cfg, "h2qp_advice_of_charge")
		if h2qp_advice_of_charge == "" {
			return fmt.Errorf("Parameter h2qp_advice_of_charge is missing")
		}
		if err = d.Set("h2qp_advice_of_charge", h2qp_advice_of_charge); err != nil {
			return fmt.Errorf("Error set params h2qp_advice_of_charge: %v", err)
		}
	}
	paradict["h2qp_advice_of_charge"] = h2qp_advice_of_charge

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "currency"
		if _, ok := i["currency"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency2edl(i["currency"], d, pre_append)
			tmp["currency"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-Currency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "info_file"
		if _, ok := i["info-file"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile2edl(i["info-file"], d, pre_append)
			tmp["info_file"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-InfoFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang2edl(i["lang"], d, pre_append)
			tmp["lang"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-Lang")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("nai_realm", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm2edl(o["nai-realm"], d, "nai_realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["nai-realm"], "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-NaiRealm"); ok {
			if err = d.Set("nai_realm", vv); err != nil {
				return fmt.Errorf("Error reading nai_realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nai_realm: %v", err)
		}
	}

	if err = d.Set("nai_realm_encoding", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding2edl(o["nai-realm-encoding"], d, "nai_realm_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["nai-realm-encoding"], "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-NaiRealmEncoding"); ok {
			if err = d.Set("nai_realm_encoding", vv); err != nil {
				return fmt.Errorf("Error reading nai_realm_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nai_realm_encoding: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("plan_info", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo2edl(o["plan-info"], d, "plan_info")); err != nil {
			if vv, ok := fortiAPIPatch(o["plan-info"], "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo"); ok {
				if err = d.Set("plan_info", vv); err != nil {
					return fmt.Errorf("Error reading plan_info: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading plan_info: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("plan_info"); ok {
			if err = d.Set("plan_info", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo2edl(o["plan-info"], d, "plan_info")); err != nil {
				if vv, ok := fortiAPIPatch(o["plan-info"], "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo"); ok {
					if err = d.Set("plan_info", vv); err != nil {
						return fmt.Errorf("Error reading plan_info: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading plan_info: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "currency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["currency"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency2edl(d, i["currency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "info_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["info-file"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile2edl(d, i["info_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lang"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang2edl(d, i["lang"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("nai_realm"); ok || d.HasChange("nai_realm") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm2edl(d, v, "nai_realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm"] = t
		}
	}

	if v, ok := d.GetOk("nai_realm_encoding"); ok || d.HasChange("nai_realm_encoding") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding2edl(d, v, "nai_realm_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm-encoding"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("plan_info"); ok || d.HasChange("plan_info") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo2edl(d, v, "plan_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["plan-info"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
