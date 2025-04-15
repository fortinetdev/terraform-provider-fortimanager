// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB user activity tenant extraction.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityMatchTenantExtraction() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityMatchTenantExtractionUpdate,
		Read:   resourceObjectCasbUserActivityMatchTenantExtractionRead,
		Update: resourceObjectCasbUserActivityMatchTenantExtractionUpdate,
		Delete: resourceObjectCasbUserActivityMatchTenantExtractionDelete,

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
							Computed: true,
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectCasbUserActivityMatchTenantExtractionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbUserActivityMatchTenantExtraction(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantExtraction resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityMatchTenantExtraction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchTenantExtraction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectCasbUserActivityMatchTenantExtraction")

	return resourceObjectCasbUserActivityMatchTenantExtractionRead(d, m)
}

func resourceObjectCasbUserActivityMatchTenantExtractionDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbUserActivityMatchTenantExtraction(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityMatchTenantExtraction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityMatchTenantExtractionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCasbUserActivityMatchTenantExtraction(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantExtraction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityMatchTenantExtraction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchTenantExtraction resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityMatchTenantExtractionFilters3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(i["body-type"], d, pre_append)
			tmp["body_type"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-BodyType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersId3rdl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := i["place"]; ok {
			v := flattenObjectCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(i["place"], d, pre_append)
			tmp["place"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityMatchTenantExtraction-Filters-Place")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionJq3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchTenantExtractionType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityMatchTenantExtraction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("filters", flattenObjectCasbUserActivityMatchTenantExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "ObjectCasbUserActivityMatchTenantExtraction-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenObjectCasbUserActivityMatchTenantExtractionFilters3rdl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "ObjectCasbUserActivityMatchTenantExtraction-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	if err = d.Set("jq", flattenObjectCasbUserActivityMatchTenantExtractionJq3rdl(o["jq"], d, "jq")); err != nil {
		if vv, ok := fortiAPIPatch(o["jq"], "ObjectCasbUserActivityMatchTenantExtraction-Jq"); ok {
			if err = d.Set("jq", vv); err != nil {
				return fmt.Errorf("Error reading jq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jq: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectCasbUserActivityMatchTenantExtractionStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectCasbUserActivityMatchTenantExtraction-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCasbUserActivityMatchTenantExtractionType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCasbUserActivityMatchTenantExtraction-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityMatchTenantExtractionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityMatchTenantExtractionFilters3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["body-type"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(d, i["body_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersId3rdl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["place"], _ = expandObjectCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(d, i["place"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersBodyType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersDirection3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionFiltersPlace3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionJq3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchTenantExtractionType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityMatchTenantExtraction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionFilters3rdl(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	if v, ok := d.GetOk("jq"); ok || d.HasChange("jq") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionJq3rdl(d, v, "jq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jq"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCasbUserActivityMatchTenantExtractionType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
