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
		},
	}
}

func resourceObjectCliTemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCliTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCliTemplate resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCliTemplate(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCliTemplate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCliTemplateRead(d, m)
}

func resourceObjectCliTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectCliTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCliTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCliTemplate(obj, adomv, mkey, nil)
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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectCliTemplate(adomv, mkey, nil)
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

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectCliTemplate(adomv, mkey, nil)
	if err != nil {
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

func refreshObjectObjectCliTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
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

	return &obj, nil
}
