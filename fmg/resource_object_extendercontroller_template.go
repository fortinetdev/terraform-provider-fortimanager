// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectExtenderController Template

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectExtenderControllerTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerTemplateCreate,
		Read:   resourceObjectExtenderControllerTemplateRead,
		Update: resourceObjectExtenderControllerTemplateUpdate,
		Delete: resourceObjectExtenderControllerTemplateDelete,

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
			"dataplan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem1_ifname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem1_sim_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem2_ifname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem2_sim_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtenderControllerTemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerTemplate(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerTemplate resource while getting object: %v", err)
	}

	_, err = c.CreateObjectExtenderControllerTemplate(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerTemplate resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerTemplateRead(d, m)
}

func resourceObjectExtenderControllerTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerTemplate(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerTemplate resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtenderControllerTemplate(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerTemplateRead(d, m)
}

func resourceObjectExtenderControllerTemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectExtenderControllerTemplate(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerTemplateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectExtenderControllerTemplate(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerTemplate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerTemplate resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerTemplateDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerTemplateDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerTemplateModem1Ifname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerTemplateModem1SimProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerTemplateModem2Ifname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerTemplateModem2SimProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerTemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerTemplate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dataplan", flattenObjectExtenderControllerTemplateDataplan(o["dataplan"], d, "dataplan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dataplan"], "ObjectExtenderControllerTemplate-Dataplan"); ok {
			if err = d.Set("dataplan", vv); err != nil {
				return fmt.Errorf("Error reading dataplan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dataplan: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectExtenderControllerTemplateDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectExtenderControllerTemplate-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("modem1_ifname", flattenObjectExtenderControllerTemplateModem1Ifname(o["modem1_ifname"], d, "modem1_ifname")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem1_ifname"], "ObjectExtenderControllerTemplate-Modem1Ifname"); ok {
			if err = d.Set("modem1_ifname", vv); err != nil {
				return fmt.Errorf("Error reading modem1_ifname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem1_ifname: %v", err)
		}
	}

	if err = d.Set("modem1_sim_profile", flattenObjectExtenderControllerTemplateModem1SimProfile(o["modem1_sim_profile"], d, "modem1_sim_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem1_sim_profile"], "ObjectExtenderControllerTemplate-Modem1SimProfile"); ok {
			if err = d.Set("modem1_sim_profile", vv); err != nil {
				return fmt.Errorf("Error reading modem1_sim_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem1_sim_profile: %v", err)
		}
	}

	if err = d.Set("modem2_ifname", flattenObjectExtenderControllerTemplateModem2Ifname(o["modem2_ifname"], d, "modem2_ifname")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem2_ifname"], "ObjectExtenderControllerTemplate-Modem2Ifname"); ok {
			if err = d.Set("modem2_ifname", vv); err != nil {
				return fmt.Errorf("Error reading modem2_ifname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem2_ifname: %v", err)
		}
	}

	if err = d.Set("modem2_sim_profile", flattenObjectExtenderControllerTemplateModem2SimProfile(o["modem2_sim_profile"], d, "modem2_sim_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem2_sim_profile"], "ObjectExtenderControllerTemplate-Modem2SimProfile"); ok {
			if err = d.Set("modem2_sim_profile", vv); err != nil {
				return fmt.Errorf("Error reading modem2_sim_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem2_sim_profile: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtenderControllerTemplateName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtenderControllerTemplate-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerTemplateDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerTemplateDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerTemplateModem1Ifname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerTemplateModem1SimProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerTemplateModem2Ifname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerTemplateModem2SimProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerTemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerTemplate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dataplan"); ok || d.HasChange("dataplan") {
		t, err := expandObjectExtenderControllerTemplateDataplan(d, v, "dataplan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dataplan"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectExtenderControllerTemplateDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("modem1_ifname"); ok || d.HasChange("modem1_ifname") {
		t, err := expandObjectExtenderControllerTemplateModem1Ifname(d, v, "modem1_ifname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1_ifname"] = t
		}
	}

	if v, ok := d.GetOk("modem1_sim_profile"); ok || d.HasChange("modem1_sim_profile") {
		t, err := expandObjectExtenderControllerTemplateModem1SimProfile(d, v, "modem1_sim_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1_sim_profile"] = t
		}
	}

	if v, ok := d.GetOk("modem2_ifname"); ok || d.HasChange("modem2_ifname") {
		t, err := expandObjectExtenderControllerTemplateModem2Ifname(d, v, "modem2_ifname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2_ifname"] = t
		}
	}

	if v, ok := d.GetOk("modem2_sim_profile"); ok || d.HasChange("modem2_sim_profile") {
		t, err := expandObjectExtenderControllerTemplateModem2SimProfile(d, v, "modem2_sim_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2_sim_profile"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtenderControllerTemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
