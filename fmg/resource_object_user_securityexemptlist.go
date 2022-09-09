// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure security exemption list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectUserSecurityExemptList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserSecurityExemptListCreate,
		Read:   resourceObjectUserSecurityExemptListRead,
		Update: resourceObjectUserSecurityExemptListUpdate,
		Delete: resourceObjectUserSecurityExemptListDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"devices": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceObjectUserSecurityExemptListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserSecurityExemptList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSecurityExemptList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserSecurityExemptList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSecurityExemptList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserSecurityExemptListRead(d, m)
}

func resourceObjectUserSecurityExemptListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectUserSecurityExemptList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSecurityExemptList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserSecurityExemptList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSecurityExemptList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserSecurityExemptListRead(d, m)
}

func resourceObjectUserSecurityExemptListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectUserSecurityExemptList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserSecurityExemptList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserSecurityExemptListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectUserSecurityExemptList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserSecurityExemptList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserSecurityExemptList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserSecurityExemptList resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserSecurityExemptListDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "devices"
		if _, ok := i["devices"]; ok {
			v := flattenObjectUserSecurityExemptListRuleDevices(i["devices"], d, pre_append)
			tmp["devices"] = fortiAPISubPartPatch(v, "ObjectUserSecurityExemptList-Rule-Devices")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenObjectUserSecurityExemptListRuleDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "ObjectUserSecurityExemptList-Rule-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectUserSecurityExemptListRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectUserSecurityExemptList-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenObjectUserSecurityExemptListRuleService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "ObjectUserSecurityExemptList-Rule-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenObjectUserSecurityExemptListRuleSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "ObjectUserSecurityExemptList-Rule-Srcaddr")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectUserSecurityExemptListRuleDevices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserSecurityExemptList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenObjectUserSecurityExemptListDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectUserSecurityExemptList-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectUserSecurityExemptListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserSecurityExemptList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenObjectUserSecurityExemptListRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "ObjectUserSecurityExemptList-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenObjectUserSecurityExemptListRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "ObjectUserSecurityExemptList-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectUserSecurityExemptListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserSecurityExemptListDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "devices"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["devices"], _ = expandObjectUserSecurityExemptListRuleDevices(d, i["devices"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandObjectUserSecurityExemptListRuleDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectUserSecurityExemptListRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandObjectUserSecurityExemptListRuleService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandObjectUserSecurityExemptListRuleSrcaddr(d, i["srcaddr"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectUserSecurityExemptListRuleDevices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserSecurityExemptList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectUserSecurityExemptListDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserSecurityExemptListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectUserSecurityExemptListRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
