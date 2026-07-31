// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB user activity tenant session extraction.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityMatchTenantSessionExtraction() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityMatchTenantSessionExtractionUpdate,
		Read:   resourceObjectCasbUserActivityMatchTenantSessionExtractionRead,
		Update: resourceObjectCasbUserActivityMatchTenantSessionExtractionUpdate,
		Delete: resourceObjectCasbUserActivityMatchTenantSessionExtractionDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cookie_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"place": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"jq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
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

func resourceObjectCasbUserActivityMatchTenantSessionExtractionUpdate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectObjectCasbUserActivityMatchTenantSessionExtraction(d, false)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantSessionExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityMatchTenantSessionExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantSessionExtraction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectCasbUserActivityMatchTenantSessionExtraction")

	return resourceObjectCasbUserActivityMatchTenantSessionExtractionRead(d, m)
}

func resourceObjectCasbUserActivityMatchTenantSessionExtractionDelete(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectObjectCasbUserActivityMatchTenantSessionExtraction(d, true)

	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantSessionExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityMatchTenantSessionExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing ObjectCasbUserActivityMatchTenantSessionExtraction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityMatchTenantSessionExtractionRead(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	if match == "" {
		match = importOptionChecking(m.(*FortiClient).Cfg, "match")
		if match == "" {
			return fmt.Errorf("Parameter match is missing")
		}
		if err = d.Set("match", match); err != nil {
			return fmt.Errorf("Error set params match: %v", err)
		}
	}
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	o, err := c.ReadObjectCasbUserActivityMatchTenantSessionExtraction(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantSessionExtraction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityMatchTenantSessionExtraction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantSessionExtraction resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFilters3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := i["body-type"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cookie_name"
		if _, ok := i["cookie-name"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(i["cookie-name"], d, pre_append)
			tmp["cookie_name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters-CookieName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionJq3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityMatchTenantSessionExtraction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("filters", flattenObjectCasbUserActivityMatchTenantSessionExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenObjectCasbUserActivityMatchTenantSessionExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "ObjectCasbUserActivityMatchTenantSessionExtraction-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("jq", flattenObjectCasbUserActivityMatchTenantSessionExtractionJq3rdl(o["jq"], d, "jq")); err != nil {
		if vv, ok := fortiAPIPatch(o["jq"], "ObjectCasbUserActivityMatchTenantSessionExtraction-Jq"); ok {
			if err = d.Set("jq", vv); err != nil {
				return fmt.Errorf("Error reading jq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jq: %v", err)
		}
	}

	if err = d.Set("session_match", flattenObjectCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(o["session-match"], d, "session_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-match"], "ObjectCasbUserActivityMatchTenantSessionExtraction-SessionMatch"); ok {
			if err = d.Set("session_match", vv); err != nil {
				return fmt.Errorf("Error reading session_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_match: %v", err)
		}
	}

	if err = d.Set("session_source", flattenObjectCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(o["session-source"], d, "session_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-source"], "ObjectCasbUserActivityMatchTenantSessionExtraction-SessionSource"); ok {
			if err = d.Set("session_source", vv); err != nil {
				return fmt.Errorf("Error reading session_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_source: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectCasbUserActivityMatchTenantSessionExtractionStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectCasbUserActivityMatchTenantSessionExtraction-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityMatchTenantSessionExtractionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFilters3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["body-type"], _ = expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cookie_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cookie-name"], _ = expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(d, i["cookie_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersBodyType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersCookieName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionFiltersPlace3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionJq3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantSessionExtractionStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityMatchTenantSessionExtraction(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if bemptysontable {
		obj["filters"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
			t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionFilters3rdl(d, v, "filters")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filters"] = t
			}
		}
	}

	if v, ok := d.GetOk("jq"); ok || d.HasChange("jq") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionJq3rdl(d, v, "jq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jq"] = t
		}
	}

	if v, ok := d.GetOk("session_match"); ok || d.HasChange("session_match") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionSessionMatch3rdl(d, v, "session_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-match"] = t
		}
	}

	if v, ok := d.GetOk("session_source"); ok || d.HasChange("session_source") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionSessionSource3rdl(d, v, "session_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-source"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectCasbUserActivityMatchTenantSessionExtractionStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
