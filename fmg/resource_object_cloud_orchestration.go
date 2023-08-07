// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectCloud Orchestration

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCloudOrchestration() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCloudOrchestrationCreate,
		Read:   resourceObjectCloudOrchestrationRead,
		Update: resourceObjectCloudOrchestrationUpdate,
		Delete: resourceObjectCloudOrchestrationDelete,

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
			"configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCloudOrchestrationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectCloudOrchestration(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestration resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCloudOrchestration(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCloudOrchestration resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestrationRead(d, m)
}

func resourceObjectCloudOrchestrationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCloudOrchestration(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestration resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCloudOrchestration(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCloudOrchestration resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCloudOrchestrationRead(d, m)
}

func resourceObjectCloudOrchestrationDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCloudOrchestration(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCloudOrchestration resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCloudOrchestrationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCloudOrchestration(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestration resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCloudOrchestration(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCloudOrchestration resource from API: %v", err)
	}
	return nil
}

func flattenObjectCloudOrchestrationConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestrationDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestrationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCloudOrchestrationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCloudOrchestration(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("configuration", flattenObjectCloudOrchestrationConfiguration(o["configuration"], d, "configuration")); err != nil {
		if vv, ok := fortiAPIPatch(o["configuration"], "ObjectCloudOrchestration-Configuration"); ok {
			if err = d.Set("configuration", vv); err != nil {
				return fmt.Errorf("Error reading configuration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading configuration: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectCloudOrchestrationDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectCloudOrchestration-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCloudOrchestrationName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCloudOrchestration-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCloudOrchestrationType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCloudOrchestration-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectCloudOrchestrationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCloudOrchestrationConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestrationDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestrationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCloudOrchestrationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCloudOrchestration(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("configuration"); ok || d.HasChange("configuration") {
		t, err := expandObjectCloudOrchestrationConfiguration(d, v, "configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectCloudOrchestrationDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCloudOrchestrationName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCloudOrchestrationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
