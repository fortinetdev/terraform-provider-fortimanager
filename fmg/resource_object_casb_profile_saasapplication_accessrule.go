// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB profile access rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplicationAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationAccessRuleCreate,
		Read:   resourceObjectCasbProfileSaasApplicationAccessRuleRead,
		Update: resourceObjectCasbProfileSaasApplicationAccessRuleUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationAccessRuleDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bypass": &schema.Schema{
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

func resourceObjectCasbProfileSaasApplicationAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectCasbProfileSaasApplicationAccessRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAccessRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCasbProfileSaasApplicationAccessRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationAccessRuleRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	obj, err := getObjectObjectCasbProfileSaasApplicationAccessRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAccessRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCasbProfileSaasApplicationAccessRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectCasbProfileSaasApplicationAccessRuleRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	err = c.DeleteObjectCasbProfileSaasApplicationAccessRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationAccessRuleRead(d *schema.ResourceData, m interface{}) error {
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
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application

	o, err := c.ReadObjectCasbProfileSaasApplicationAccessRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAccessRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplicationAccessRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAccessRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAction3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleBypass3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAccessRuleName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbProfileSaasApplicationAccessRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectCasbProfileSaasApplicationAccessRuleAction3rdl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectCasbProfileSaasApplicationAccessRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("bypass", flattenObjectCasbProfileSaasApplicationAccessRuleBypass3rdl(o["bypass"], d, "bypass")); err != nil {
		if vv, ok := fortiAPIPatch(o["bypass"], "ObjectCasbProfileSaasApplicationAccessRule-Bypass"); ok {
			if err = d.Set("bypass", vv); err != nil {
				return fmt.Errorf("Error reading bypass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bypass: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectCasbProfileSaasApplicationAccessRuleName3rdl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectCasbProfileSaasApplicationAccessRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationAccessRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationAccessRuleAction3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleBypass3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbProfileSaasApplicationAccessRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRuleAction3rdl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("bypass"); ok || d.HasChange("bypass") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRuleBypass3rdl(d, v, "bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bypass"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRuleName3rdl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
