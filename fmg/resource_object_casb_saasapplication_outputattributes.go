// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SaaS application output attributes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbSaasApplicationOutputAttributes() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbSaasApplicationOutputAttributesCreate,
		Read:   resourceObjectCasbSaasApplicationOutputAttributesRead,
		Update: resourceObjectCasbSaasApplicationOutputAttributesUpdate,
		Delete: resourceObjectCasbSaasApplicationOutputAttributesDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCasbSaasApplicationOutputAttributesCreate(d *schema.ResourceData, m interface{}) error {
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

	saas_application := d.Get("saas_application").(string)
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectCasbSaasApplicationOutputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbSaasApplicationOutputAttributes resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbSaasApplicationOutputAttributes(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbSaasApplicationOutputAttributes resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbSaasApplicationOutputAttributesRead(d, m)
}

func resourceObjectCasbSaasApplicationOutputAttributesUpdate(d *schema.ResourceData, m interface{}) error {
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

	saas_application := d.Get("saas_application").(string)
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectCasbSaasApplicationOutputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbSaasApplicationOutputAttributes resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbSaasApplicationOutputAttributes(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbSaasApplicationOutputAttributes resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbSaasApplicationOutputAttributesRead(d, m)
}

func resourceObjectCasbSaasApplicationOutputAttributesDelete(d *schema.ResourceData, m interface{}) error {
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

	saas_application := d.Get("saas_application").(string)
	paradict["saas_application"] = saas_application

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbSaasApplicationOutputAttributes(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbSaasApplicationOutputAttributes resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbSaasApplicationOutputAttributesRead(d *schema.ResourceData, m interface{}) error {
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

	saas_application := d.Get("saas_application").(string)
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	paradict["saas_application"] = saas_application

	o, err := c.ReadObjectCasbSaasApplicationOutputAttributes(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbSaasApplicationOutputAttributes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbSaasApplicationOutputAttributes(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbSaasApplicationOutputAttributes resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbSaasApplicationOutputAttributesAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesRequired2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationOutputAttributesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbSaasApplicationOutputAttributes(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("attr_type", flattenObjectCasbSaasApplicationOutputAttributesAttrType2edl(o["attr-type"], d, "attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["attr-type"], "ObjectCasbSaasApplicationOutputAttributes-AttrType"); ok {
			if err = d.Set("attr_type", vv); err != nil {
				return fmt.Errorf("Error reading attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attr_type: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectCasbSaasApplicationOutputAttributesDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectCasbSaasApplicationOutputAttributes-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbSaasApplicationOutputAttributesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbSaasApplicationOutputAttributes-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("required", flattenObjectCasbSaasApplicationOutputAttributesRequired2edl(o["required"], d, "required")); err != nil {
		if vv, ok := fortiAPIPatch(o["required"], "ObjectCasbSaasApplicationOutputAttributes-Required"); ok {
			if err = d.Set("required", vv); err != nil {
				return fmt.Errorf("Error reading required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCasbSaasApplicationOutputAttributesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCasbSaasApplicationOutputAttributes-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbSaasApplicationOutputAttributesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbSaasApplicationOutputAttributesAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesRequired2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationOutputAttributesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbSaasApplicationOutputAttributes(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attr_type"); ok || d.HasChange("attr_type") {
		t, err := expandObjectCasbSaasApplicationOutputAttributesAttrType2edl(d, v, "attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attr-type"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectCasbSaasApplicationOutputAttributesDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbSaasApplicationOutputAttributesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok || d.HasChange("required") {
		t, err := expandObjectCasbSaasApplicationOutputAttributesRequired2edl(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCasbSaasApplicationOutputAttributesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
