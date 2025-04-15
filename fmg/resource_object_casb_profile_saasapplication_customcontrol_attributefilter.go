// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB attribute filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterCreate,
		Read:   resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterRead,
		Update: resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterUpdate,
		Delete: resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterCreate(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectCasbProfileSaasApplicationCustomControlAttributeFilter(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	obj, err := getObjectObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectCasbProfileSaasApplicationCustomControlAttributeFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterRead(d, m)
}

func resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterDelete(d *schema.ResourceData, m interface{}) error {
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
	custom_control := d.Get("custom_control").(string)
	paradict["profile"] = profile
	paradict["saas_application"] = saas_application
	paradict["custom_control"] = custom_control

	wsParams["adom"] = adomv

	err = c.DeleteObjectCasbProfileSaasApplicationCustomControlAttributeFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbProfileSaasApplicationCustomControlAttributeFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectCasbProfileSaasApplicationCustomControlAttributeFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbProfileSaasApplicationCustomControlAttributeFilter resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction4thl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectCasbProfileSaasApplicationCustomControlAttributeFilter-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("attribute_match", flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch4thl(o["attribute-match"], d, "attribute_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["attribute-match"], "ObjectCasbProfileSaasApplicationCustomControlAttributeFilter-AttributeMatch"); ok {
			if err = d.Set("attribute_match", vv); err != nil {
				return fmt.Errorf("Error reading attribute_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading attribute_match: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectCasbProfileSaasApplicationCustomControlAttributeFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbProfileSaasApplicationCustomControlAttributeFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbProfileSaasApplicationCustomControlAttributeFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAction4thl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("attribute_match"); ok || d.HasChange("attribute_match") {
		t, err := expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch4thl(d, v, "attribute_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-match"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectCasbProfileSaasApplicationCustomControlAttributeFilterId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
