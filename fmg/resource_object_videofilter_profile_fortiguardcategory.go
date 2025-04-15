// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard categories.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVideofilterProfileFortiguardCategory() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVideofilterProfileFortiguardCategoryUpdate,
		Read:   resourceObjectVideofilterProfileFortiguardCategoryRead,
		Update: resourceObjectVideofilterProfileFortiguardCategoryUpdate,
		Delete: resourceObjectVideofilterProfileFortiguardCategoryDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"category_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectVideofilterProfileFortiguardCategoryUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectVideofilterProfileFortiguardCategory(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterProfileFortiguardCategory resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVideofilterProfileFortiguardCategory(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVideofilterProfileFortiguardCategory resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectVideofilterProfileFortiguardCategory")

	return resourceObjectVideofilterProfileFortiguardCategoryRead(d, m)
}

func resourceObjectVideofilterProfileFortiguardCategoryDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectVideofilterProfileFortiguardCategory(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVideofilterProfileFortiguardCategory resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVideofilterProfileFortiguardCategoryRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectVideofilterProfileFortiguardCategory(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterProfileFortiguardCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVideofilterProfileFortiguardCategory(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVideofilterProfileFortiguardCategory resource from API: %v", err)
	}
	return nil
}

func flattenObjectVideofilterProfileFortiguardCategoryFilters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := i["category-id"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersCategoryId2edl(i["category-id"], d, pre_append)
			tmp["category_id"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-CategoryId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenObjectVideofilterProfileFortiguardCategoryFiltersLog2edl(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "ObjectVideofilterProfileFortiguardCategory-Filters-Log")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersCategoryId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVideofilterProfileFortiguardCategoryFiltersLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVideofilterProfileFortiguardCategory(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("filters", flattenObjectVideofilterProfileFortiguardCategoryFilters2edl(o["filters"], d, "filters")); err != nil {
			if vv, ok := fortiAPIPatch(o["filters"], "ObjectVideofilterProfileFortiguardCategory-Filters"); ok {
				if err = d.Set("filters", vv); err != nil {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading filters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filters"); ok {
			if err = d.Set("filters", flattenObjectVideofilterProfileFortiguardCategoryFilters2edl(o["filters"], d, "filters")); err != nil {
				if vv, ok := fortiAPIPatch(o["filters"], "ObjectVideofilterProfileFortiguardCategory-Filters"); ok {
					if err = d.Set("filters", vv); err != nil {
						return fmt.Errorf("Error reading filters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading filters: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectVideofilterProfileFortiguardCategoryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVideofilterProfileFortiguardCategoryFilters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["category-id"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersCategoryId2edl(d, i["category_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandObjectVideofilterProfileFortiguardCategoryFiltersLog2edl(d, i["log"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersCategoryId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVideofilterProfileFortiguardCategoryFiltersLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVideofilterProfileFortiguardCategory(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("filters"); ok || d.HasChange("filters") {
		t, err := expandObjectVideofilterProfileFortiguardCategoryFilters2edl(d, v, "filters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filters"] = t
		}
	}

	return &obj, nil
}
