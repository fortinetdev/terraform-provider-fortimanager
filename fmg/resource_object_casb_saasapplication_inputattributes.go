// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SaaS application input attributes.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbSaasApplicationInputAttributes() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbSaasApplicationInputAttributesCreate,
		Read:   resourceObjectCasbSaasApplicationInputAttributesRead,
		Update: resourceObjectCasbSaasApplicationInputAttributesUpdate,
		Delete: resourceObjectCasbSaasApplicationInputAttributesDelete,

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
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fallback_input": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceObjectCasbSaasApplicationInputAttributesCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbSaasApplicationInputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbSaasApplicationInputAttributes resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbSaasApplicationInputAttributes(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbSaasApplicationInputAttributes resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbSaasApplicationInputAttributesRead(d, m)
}

func resourceObjectCasbSaasApplicationInputAttributesUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectCasbSaasApplicationInputAttributes(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbSaasApplicationInputAttributes resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbSaasApplicationInputAttributes(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbSaasApplicationInputAttributes resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbSaasApplicationInputAttributesRead(d, m)
}

func resourceObjectCasbSaasApplicationInputAttributesDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectCasbSaasApplicationInputAttributes(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbSaasApplicationInputAttributes resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbSaasApplicationInputAttributesRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCasbSaasApplicationInputAttributes(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbSaasApplicationInputAttributes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbSaasApplicationInputAttributes(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbSaasApplicationInputAttributes resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbSaasApplicationInputAttributesAttrType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesFallbackInput2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesRequired2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbSaasApplicationInputAttributesType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbSaasApplicationInputAttributes(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("attr_type", flattenObjectCasbSaasApplicationInputAttributesAttrType2edl(o["attr-type"], d, "attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["attr-type"], "ObjectCasbSaasApplicationInputAttributes-AttrType"); ok {
			if err = d.Set("attr_type", vv); err != nil {
				return fmt.Errorf("Error reading attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attr_type: %v", err)
		}
	}

	if err = d.Set("default", flattenObjectCasbSaasApplicationInputAttributesDefault2edl(o["default"], d, "default")); err != nil {
		if vv, ok := fortiAPIPatch(o["default"], "ObjectCasbSaasApplicationInputAttributes-Default"); ok {
			if err = d.Set("default", vv); err != nil {
				return fmt.Errorf("Error reading default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectCasbSaasApplicationInputAttributesDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectCasbSaasApplicationInputAttributes-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fallback_input", flattenObjectCasbSaasApplicationInputAttributesFallbackInput2edl(o["fallback-input"], d, "fallback_input")); err != nil {
		if vv, ok := fortiAPIPatch(o["fallback-input"], "ObjectCasbSaasApplicationInputAttributes-FallbackInput"); ok {
			if err = d.Set("fallback_input", vv); err != nil {
				return fmt.Errorf("Error reading fallback_input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fallback_input: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbSaasApplicationInputAttributesName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbSaasApplicationInputAttributes-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("required", flattenObjectCasbSaasApplicationInputAttributesRequired2edl(o["required"], d, "required")); err != nil {
		if vv, ok := fortiAPIPatch(o["required"], "ObjectCasbSaasApplicationInputAttributes-Required"); ok {
			if err = d.Set("required", vv); err != nil {
				return fmt.Errorf("Error reading required: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCasbSaasApplicationInputAttributesType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCasbSaasApplicationInputAttributes-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbSaasApplicationInputAttributesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbSaasApplicationInputAttributesAttrType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesFallbackInput2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesRequired2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbSaasApplicationInputAttributesType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbSaasApplicationInputAttributes(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("attr_type"); ok || d.HasChange("attr_type") {
		t, err := expandObjectCasbSaasApplicationInputAttributesAttrType2edl(d, v, "attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attr-type"] = t
		}
	}

	if v, ok := d.GetOk("default"); ok || d.HasChange("default") {
		t, err := expandObjectCasbSaasApplicationInputAttributesDefault2edl(d, v, "default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectCasbSaasApplicationInputAttributesDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fallback_input"); ok || d.HasChange("fallback_input") {
		t, err := expandObjectCasbSaasApplicationInputAttributesFallbackInput2edl(d, v, "fallback_input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-input"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbSaasApplicationInputAttributesName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok || d.HasChange("required") {
		t, err := expandObjectCasbSaasApplicationInputAttributesRequired2edl(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCasbSaasApplicationInputAttributesType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
