// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Wan Template

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanTemplateCreate,
		Read:   resourceWanTemplateRead,
		Update: resourceWanTemplateUpdate,
		Delete: resourceWanTemplateDelete,

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
				Required: true,
			},
			"oid": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceWanTemplateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating WanTemplate resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateWanTemplate(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WanTemplate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceWanTemplateRead(d, m)
}

func resourceWanTemplateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating WanTemplate resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWanTemplate(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWanTemplateRead(d, m)
}

func resourceWanTemplateDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanTemplate(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanTemplate(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanTemplate resource from API: %v", err)
	}
	return nil
}

func flattenWanTemplateDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanTemplateOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanTemplateScopeMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenWanTemplateScopeMemberName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "WanTemplate-ScopeMember-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenWanTemplateScopeMemberVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "WanTemplate-ScopeMember-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanTemplateScopeMemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanTemplateScopeMemberVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanTemplateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenWanTemplateDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "WanTemplate-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenWanTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WanTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("oid", flattenWanTemplateOid(o["oid"], d, "oid")); err != nil {
		if vv, ok := fortiAPIPatch(o["oid"], "WanTemplate-Oid"); ok {
			if err = d.Set("oid", vv); err != nil {
				return fmt.Errorf("Error reading oid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("scopemember", flattenWanTemplateScopeMember(o["scope member"], d, "scopemember")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope member"], "WanTemplate-ScopeMember"); ok {
				if err = d.Set("scopemember", vv); err != nil {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scopemember: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scopemember"); ok {
			if err = d.Set("scopemember", flattenWanTemplateScopeMember(o["scope member"], d, "scopemember")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope member"], "WanTemplate-ScopeMember"); ok {
					if err = d.Set("scopemember", vv); err != nil {
						return fmt.Errorf("Error reading scopemember: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scopemember: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenWanTemplateType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "WanTemplate-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenWanTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanTemplateDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanTemplateOid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanTemplateScopeMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWanTemplateScopeMemberName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandWanTemplateScopeMemberVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanTemplateScopeMemberName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanTemplateScopeMemberVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanTemplateType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandWanTemplateDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWanTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok || d.HasChange("oid") {
		t, err := expandWanTemplateOid(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("scopemember"); ok || d.HasChange("scopemember") {
		t, err := expandWanTemplateScopeMember(d, v, "scopemember")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope member"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandWanTemplateType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
