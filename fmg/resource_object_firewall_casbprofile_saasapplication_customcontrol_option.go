// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFirewall CasbProfileSaasApplicationCustomControlOption

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionCreate,
		Read:   resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionRead,
		Update: resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionUpdate,
		Delete: resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionDelete,

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
			"casb_profile": &schema.Schema{
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

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectObjectFirewallCasbProfileSaasApplicationCustomControlOption(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallCasbProfileSaasApplicationCustomControlOption(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionRead(d, m)
}

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionUpdate(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectObjectFirewallCasbProfileSaasApplicationCustomControlOption(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallCasbProfileSaasApplicationCustomControlOption(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionRead(d, m)
}

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionDelete(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	err = c.DeleteObjectFirewallCasbProfileSaasApplicationCustomControlOption(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallCasbProfileSaasApplicationCustomControlOptionRead(d *schema.ResourceData, m interface{}) error {
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

	casb_profile := d.Get("casb_profile").(string)
	saas_application := d.Get("saas_application").(string)
	custom_control := d.Get("custom_control").(string)
	if casb_profile == "" {
		casb_profile = importOptionChecking(m.(*FortiClient).Cfg, "casb_profile")
		if casb_profile == "" {
			return fmt.Errorf("Parameter casb_profile is missing")
		}
		if err = d.Set("casb_profile", casb_profile); err != nil {
			return fmt.Errorf("Error set params casb_profile: %v", err)
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
	paradict["casb_profile"] = casb_profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	o, err := c.ReadObjectFirewallCasbProfileSaasApplicationCustomControlOption(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallCasbProfileSaasApplicationCustomControlOption(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallCasbProfileSaasApplicationCustomControlOption resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectFirewallCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionName4thl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallCasbProfileSaasApplicationCustomControlOption-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("user_input", flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput4thl(o["user-input"], d, "user_input")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-input"], "ObjectFirewallCasbProfileSaasApplicationCustomControlOption-UserInput"); ok {
			if err = d.Set("user_input", vv); err != nil {
				return fmt.Errorf("Error reading user_input: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_input: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallCasbProfileSaasApplicationCustomControlOptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectFirewallCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionName4thl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("user_input"); ok || d.HasChange("user_input") {
		t, err := expandObjectFirewallCasbProfileSaasApplicationCustomControlOptionUserInput4thl(d, v, "user_input")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-input"] = t
		}
	}

	return &obj, nil
}
