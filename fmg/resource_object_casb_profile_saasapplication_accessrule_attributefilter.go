// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB profile attribute filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterCreate,
		Read:   resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterRead,
		Update: resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterDelete,

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
			"access_rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute_match": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterCreate(d *schema.ResourceData, m interface{}) error {
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
	access_rule := d.Get("access_rule").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	obj, err := getObjectObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_rule := d.Get("access_rule").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	obj, err := getObjectObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterDelete(d *schema.ResourceData, m interface{}) error {
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
	access_rule := d.Get("access_rule").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationAccessRuleAttributeFilterRead(d *schema.ResourceData, m interface{}) error {
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
	access_rule := d.Get("access_rule").(string)
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
	if access_rule == "" {
		access_rule = importOptionChecking(m.(*FortiClient).Cfg, "access_rule")
		if access_rule == "" {
			return fmt.Errorf("Parameter access_rule is missing")
		}
		if err = d.Set("access_rule", access_rule); err != nil {
			return fmt.Errorf("Error set params access_rule: %v", err)
		}
	}
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["access_rule"] = access_rule

	o, err := c.ReadObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("attribute_match", flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(o["attribute-match"], d, "attribute_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-match"], "ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter-AttributeMatch"); ok {
			if err = d.Set("attribute_match", vv); err != nil {
				return fmt.Errorf("Error reading attribute_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_match: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectCasbProfileSaasApplicationAccessRuleAttributeFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationAccessRuleAttributeFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAction4thl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("attribute_match"); ok || d.HasChange("attribute_match") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch4thl(d, v, "attribute_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-match"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectCasbProfileSaasApplicationAccessRuleAttributeFilterId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
