// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure advice of charge.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfCharge() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeCreate,
		Read:   resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeRead,
		Update: resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeUpdate,
		Delete: resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeDelete,

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
			"aoc_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nai_realm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nai_realm_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"plan_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currency": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"info_file": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"lang": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20H2QpAdviceOfCharge(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20H2QpAdviceOfCharge(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeRead(d, m)
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20H2QpAdviceOfCharge(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20H2QpAdviceOfChargeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20H2QpAdviceOfCharge(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20H2QpAdviceOfCharge resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := i["nai-realm"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(i["nai-realm"], d, pre_append)
			tmp["nai_realm"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList-NaiRealm")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm_encoding"
		if _, ok := i["nai-realm-encoding"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(i["nai-realm-encoding"], d, pre_append)
			tmp["nai_realm_encoding"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList-NaiRealmEncoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "plan_info"
		if _, ok := i["plan-info"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(i["plan-info"], d, pre_append)
			tmp["plan_info"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList-PlanInfo")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList-Type")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(i["currency"], d, pre_append)
			tmp["currency"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-Currency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "info_file"
		if _, ok := i["info-file"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(i["info-file"], d, pre_append)
			tmp["info_file"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-InfoFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(i["lang"], d, pre_append)
			tmp["lang"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-Lang")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList-PlanInfo-Name")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("aoc_list", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(o["aoc-list"], d, "aoc_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["aoc-list"], "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList"); ok {
				if err = d.Set("aoc_list", vv); err != nil {
					return fmt.Errorf("Error reading aoc_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading aoc_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aoc_list"); ok {
			if err = d.Set("aoc_list", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(o["aoc-list"], d, "aoc_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["aoc-list"], "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-AocList"); ok {
					if err = d.Set("aoc_list", vv); err != nil {
						return fmt.Errorf("Error reading aoc_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading aoc_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20H2QpAdviceOfCharge-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20H2QpAdviceOfChargeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nai-realm"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(d, i["nai_realm"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm_encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nai-realm-encoding"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(d, i["nai_realm_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "plan_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["plan-info"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d, i["plan_info"], pre_append)
		} else {
			tmp["plan-info"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(d, i["type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "currency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["currency"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(d, i["currency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "info_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["info-file"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(d, i["info_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lang"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(d, i["lang"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aoc_list"); ok || d.HasChange("aoc_list") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d, v, "aoc_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aoc-list"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerHotspot20H2QpAdviceOfChargeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
