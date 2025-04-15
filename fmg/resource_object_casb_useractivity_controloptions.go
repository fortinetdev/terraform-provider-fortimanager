// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB control options.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityControlOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityControlOptionsCreate,
		Read:   resourceObjectCasbUserActivityControlOptionsRead,
		Update: resourceObjectCasbUserActivityControlOptionsUpdate,
		Delete: resourceObjectCasbUserActivityControlOptionsDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"operations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"search_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"search_pattern": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value_from_input": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"values": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
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

func resourceObjectCasbUserActivityControlOptionsCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["user_activity"] = user_activity

	obj, err := getObjectObjectCasbUserActivityControlOptions(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityControlOptions resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbUserActivityControlOptions(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityControlOptions resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbUserActivityControlOptionsRead(d, m)
}

func resourceObjectCasbUserActivityControlOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["user_activity"] = user_activity

	obj, err := getObjectObjectCasbUserActivityControlOptions(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityControlOptions resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbUserActivityControlOptions(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityControlOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbUserActivityControlOptionsRead(d, m)
}

func resourceObjectCasbUserActivityControlOptionsDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["user_activity"] = user_activity

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbUserActivityControlOptions(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityControlOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityControlOptionsRead(d *schema.ResourceData, m interface{}) error {
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
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	paradict["user_activity"] = user_activity

	o, err := c.ReadObjectCasbUserActivityControlOptions(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityControlOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityControlOptions(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityControlOptions resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityControlOptionsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperations2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectCasbUserActivityControlOptionsOperationsAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := i["case-sensitive"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsCaseSensitive2edl(i["case-sensitive"], d, pre_append)
			tmp["case_sensitive"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-CaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsHeaderName2edl(i["header-name"], d, pre_append)
			tmp["header_name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-HeaderName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_key"
		if _, ok := i["search-key"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsSearchKey2edl(i["search-key"], d, pre_append)
			tmp["search_key"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-SearchKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_pattern"
		if _, ok := i["search-pattern"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsSearchPattern2edl(i["search-pattern"], d, pre_append)
			tmp["search_pattern"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-SearchPattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsTarget2edl(i["target"], d, pre_append)
			tmp["target"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-Target")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_from_input"
		if _, ok := i["value-from-input"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsValueFromInput2edl(i["value-from-input"], d, pre_append)
			tmp["value_from_input"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-ValueFromInput")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := i["values"]; ok {
			v := flattenObjectCasbUserActivityControlOptionsOperationsValues2edl(i["values"], d, pre_append)
			tmp["values"] = fortiAPISubPartPatch(v, "ObjectCasbUserActivityControlOptions-Operations-Values")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCasbUserActivityControlOptionsOperationsAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsCaseSensitive2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsHeaderName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsSearchKey2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsSearchPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsTarget2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsValueFromInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityControlOptionsOperationsValues2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbUserActivityControlOptionsStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityControlOptions(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("name", flattenObjectCasbUserActivityControlOptionsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbUserActivityControlOptions-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("operations", flattenObjectCasbUserActivityControlOptionsOperations2edl(o["operations"], d, "operations")); err != nil {
			if vv, ok := fortiAPIPatch(o["operations"], "ObjectCasbUserActivityControlOptions-Operations"); ok {
				if err = d.Set("operations", vv); err != nil {
					return fmt.Errorf("Error reading operations: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading operations: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("operations"); ok {
			if err = d.Set("operations", flattenObjectCasbUserActivityControlOptionsOperations2edl(o["operations"], d, "operations")); err != nil {
				if vv, ok := fortiAPIPatch(o["operations"], "ObjectCasbUserActivityControlOptions-Operations"); ok {
					if err = d.Set("operations", vv); err != nil {
						return fmt.Errorf("Error reading operations: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading operations: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenObjectCasbUserActivityControlOptionsStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectCasbUserActivityControlOptions-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityControlOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityControlOptionsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandObjectCasbUserActivityControlOptionsOperationsAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["case-sensitive"], _ = expandObjectCasbUserActivityControlOptionsOperationsCaseSensitive2edl(d, i["case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandObjectCasbUserActivityControlOptionsOperationsDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header-name"], _ = expandObjectCasbUserActivityControlOptionsOperationsHeaderName2edl(d, i["header_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCasbUserActivityControlOptionsOperationsName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["search-key"], _ = expandObjectCasbUserActivityControlOptionsOperationsSearchKey2edl(d, i["search_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["search-pattern"], _ = expandObjectCasbUserActivityControlOptionsOperationsSearchPattern2edl(d, i["search_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target"], _ = expandObjectCasbUserActivityControlOptionsOperationsTarget2edl(d, i["target"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_from_input"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value-from-input"], _ = expandObjectCasbUserActivityControlOptionsOperationsValueFromInput2edl(d, i["value_from_input"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["values"], _ = expandObjectCasbUserActivityControlOptionsOperationsValues2edl(d, i["values"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsCaseSensitive2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsHeaderName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsSearchKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsSearchPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsTarget2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsValueFromInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityControlOptionsOperationsValues2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbUserActivityControlOptionsStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityControlOptions(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbUserActivityControlOptionsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("operations"); ok || d.HasChange("operations") {
		t, err := expandObjectCasbUserActivityControlOptionsOperations2edl(d, v, "operations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["operations"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectCasbUserActivityControlOptionsStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
