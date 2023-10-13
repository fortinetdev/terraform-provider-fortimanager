// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB custom control option.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplicationCustomControlOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationCustomControlOptionCreate,
		Read:   resourceObjectCasbProfileSaasApplicationCustomControlOptionRead,
		Update: resourceObjectCasbProfileSaasApplicationCustomControlOptionUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationCustomControlOptionDelete,

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
			"custom_control": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"user_input": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCasbProfileSaasApplicationCustomControlOptionCreate(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectObjectCasbProfileSaasApplicationCustomControlOption(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationCustomControlOption resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCasbProfileSaasApplicationCustomControlOption(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationCustomControlOptionRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationCustomControlOptionUpdate(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectObjectCasbProfileSaasApplicationCustomControlOption(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationCustomControlOption resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCasbProfileSaasApplicationCustomControlOption(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationCustomControlOptionRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationCustomControlOptionDelete(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	err = c.DeleteObjectCasbProfileSaasApplicationCustomControlOption(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationCustomControlOptionRead(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
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
	if custom_control == "" {
		custom_control = importOptionChecking(m.(*FortiClient).Cfg, "custom_control")
		if custom_control == "" {
			return fmt.Errorf("Parameter custom_control is missing")
		}
		if err = d.Set("custom_control", custom_control); err != nil {
			return fmt.Errorf("Error set params custom_control: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	o, err := c.ReadObjectCasbProfileSaasApplicationCustomControlOption(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplicationCustomControlOption(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationCustomControlOption resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionUserInput4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectCasbProfileSaasApplicationCustomControlOptionName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbProfileSaasApplicationCustomControlOption-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("user_input", flattenObjectCasbProfileSaasApplicationCustomControlOptionUserInput4thl(o["user-input"], d, "user_input")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-input"], "ObjectCasbProfileSaasApplicationCustomControlOption-UserInput"); ok {
			if err = d.Set("user_input", vv); err != nil {
				return fmt.Errorf("Error reading user_input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_input: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationCustomControlOptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationCustomControlOptionName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlOptionUserInput4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbProfileSaasApplicationCustomControlOptionName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("user_input"); ok || d.HasChange("user_input") {
		t, err := expandObjectCasbProfileSaasApplicationCustomControlOptionUserInput4thl(d, v, "user_input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-input"] = t
		}
	}

	return &obj, nil
}
