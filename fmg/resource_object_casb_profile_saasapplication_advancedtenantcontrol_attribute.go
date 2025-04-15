// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB advanced tenant control attribute.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeCreate,
		Read:   resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeRead,
		Update: resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeDelete,

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
			"saas_application": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advanced_tenant_control": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"input": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeCreate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	advanced_tenant_control := d.Get("advanced_tenant_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["advanced_tenant_control"] = advanced_tenant_control

	obj, err := getObjectObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeUpdate(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	advanced_tenant_control := d.Get("advanced_tenant_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["advanced_tenant_control"] = advanced_tenant_control

	obj, err := getObjectObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeDelete(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	advanced_tenant_control := d.Get("advanced_tenant_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["advanced_tenant_control"] = advanced_tenant_control

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeRead(d *schema.ResourceData, m interface{}) error {
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
	saas_application := d.Get("saas_application").(string)
	advanced_tenant_control := d.Get("advanced_tenant_control").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	if saas_application == "" {
		saas_application = importOptionChecking(m.(*FortiClient).Cfg, "saas_application")
		if saas_application == "" {
			return fmt.Errorf("Parameter saas_application is missing")
		}
		if err = d.Set("saas_application", saas_application); err != nil {
			return fmt.Errorf("Error set params saas_application: %v", err)
		}
	}
	if advanced_tenant_control == "" {
		advanced_tenant_control = importOptionChecking(m.(*FortiClient).Cfg, "advanced_tenant_control")
		if advanced_tenant_control == "" {
			return fmt.Errorf("Parameter advanced_tenant_control is missing")
		}
		if err = d.Set("advanced_tenant_control", advanced_tenant_control); err != nil {
			return fmt.Errorf("Error set params advanced_tenant_control: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["advanced_tenant_control"] = advanced_tenant_control

	o, err := c.ReadObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("input", flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput4thl(o["input"], d, "input")); err != nil {
		if vv, ok := fortiAPIPatch(o["input"], "ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute-Input"); ok {
			if err = d.Set("input", vv); err != nil {
				return fmt.Errorf("Error reading input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading input: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbProfileSaasApplicationAdvancedTenantControlAttribute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("input"); ok || d.HasChange("input") {
		t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeInput4thl(d, v, "input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbProfileSaasApplicationAdvancedTenantControlAttributeName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
