// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCli Template

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCliTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCliTemplateCreate,
		Read:   resourceObjectCliTemplateRead,
		Update: resourceObjectCliTemplateUpdate,
		Delete: resourceObjectCliTemplateDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"modification_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"position": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"provision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"variables": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scopemember": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
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

func resourceObjectCliTemplateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCliTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCliTemplate resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectCliTemplate(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectCliTemplate(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectCliTemplate resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectCliTemplate(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectCliTemplate resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCliTemplateRead(d, m)
}

func resourceObjectCliTemplateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCliTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCliTemplate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCliTemplate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCliTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCliTemplateRead(d, m)
}

func resourceObjectCliTemplateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCliTemplate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCliTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCliTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCliTemplate(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectCliTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCliTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCliTemplate resource from API: %v", err)
	}
	return nil
}

func flattenObjectCliTemplateDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateModificationTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCliTemplatePosition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateProvision(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateVariables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCliTemplateScopeMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectCliTemplateScopeMemberName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectCliTemplate-ScopeMember-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenObjectCliTemplateScopeMemberVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "ObjectCliTemplate-ScopeMember-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectCliTemplateScopeMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCliTemplateScopeMemberVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCliTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenObjectCliTemplateDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectCliTemplate-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("modification_time", flattenObjectCliTemplateModificationTime(o["modification-time"], d, "modification_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["modification-time"], "ObjectCliTemplate-ModificationTime"); ok {
			if err = d.Set("modification_time", vv); err != nil {
				return fmt.Errorf("Error reading modification_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modification_time: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCliTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCliTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("option", flattenObjectCliTemplateOption(o["option"], d, "option")); err != nil {
		if vv, ok := fortiAPIPatch(o["option"], "ObjectCliTemplate-Option"); ok {
			if err = d.Set("option", vv); err != nil {
				return fmt.Errorf("Error reading option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option: %v", err)
		}
	}

	if err = d.Set("position", flattenObjectCliTemplatePosition(o["position"], d, "position")); err != nil {
		if vv, ok := fortiAPIPatch(o["position"], "ObjectCliTemplate-Position"); ok {
			if err = d.Set("position", vv); err != nil {
				return fmt.Errorf("Error reading position: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading position: %v", err)
		}
	}

	if err = d.Set("provision", flattenObjectCliTemplateProvision(o["provision"], d, "provision")); err != nil {
		if vv, ok := fortiAPIPatch(o["provision"], "ObjectCliTemplate-Provision"); ok {
			if err = d.Set("provision", vv); err != nil {
				return fmt.Errorf("Error reading provision: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading provision: %v", err)
		}
	}

	if err = d.Set("script", flattenObjectCliTemplateScript(o["script"], d, "script")); err != nil {
		if vv, ok := fortiAPIPatch(o["script"], "ObjectCliTemplate-Script"); ok {
			if err = d.Set("script", vv); err != nil {
				return fmt.Errorf("Error reading script: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCliTemplateType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCliTemplate-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("variables", flattenObjectCliTemplateVariables(o["variables"], d, "variables")); err != nil {
		if vv, ok := fortiAPIPatch(o["variables"], "ObjectCliTemplate-Variables"); ok {
			if err = d.Set("variables", vv); err != nil {
				return fmt.Errorf("Error reading variables: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading variables: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("scopemember", flattenObjectCliTemplateScopeMember(o["scope member"], d, "scopemember")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope member"], "ObjectCliTemplate-ScopeMember"); ok {
				if err = d.Set("scopemember", vv); err != nil {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scopemember: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scopemember"); ok {
			if err = d.Set("scopemember", flattenObjectCliTemplateScopeMember(o["scope member"], d, "scopemember")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope member"], "ObjectCliTemplate-ScopeMember"); ok {
					if err = d.Set("scopemember", vv); err != nil {
						return fmt.Errorf("Error reading scopemember: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectCliTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCliTemplateDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateModificationTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCliTemplatePosition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateProvision(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateVariables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCliTemplateScopeMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectCliTemplateScopeMemberName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandObjectCliTemplateScopeMemberVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectCliTemplateScopeMemberName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCliTemplateScopeMemberVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCliTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectCliTemplateDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("modification_time"); ok || d.HasChange("modification_time") {
		t, err := expandObjectCliTemplateModificationTime(d, v, "modification_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modification-time"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCliTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandObjectCliTemplateOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	if v, ok := d.GetOk("position"); ok || d.HasChange("position") {
		t, err := expandObjectCliTemplatePosition(d, v, "position")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["position"] = t
		}
	}

	if v, ok := d.GetOk("provision"); ok || d.HasChange("provision") {
		t, err := expandObjectCliTemplateProvision(d, v, "provision")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["provision"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok || d.HasChange("script") {
		t, err := expandObjectCliTemplateScript(d, v, "script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCliTemplateType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("variables"); ok || d.HasChange("variables") {
		t, err := expandObjectCliTemplateVariables(d, v, "variables")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["variables"] = t
		}
	}

	if v, ok := d.GetOk("scopemember"); ok || d.HasChange("scopemember") {
		t, err := expandObjectCliTemplateScopeMember(d, v, "scopemember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope member"] = t
		}
	}

	return &obj, nil
}
