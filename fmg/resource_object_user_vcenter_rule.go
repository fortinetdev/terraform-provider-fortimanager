// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectUser VcenterRule

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserVcenterRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserVcenterRuleCreate,
		Read:   resourceObjectUserVcenterRuleRead,
		Update: resourceObjectUserVcenterRuleUpdate,
		Delete: resourceObjectUserVcenterRuleDelete,

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
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserVcenterRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vcenter := d.Get("vcenter").(string)
	paradict["vcenter"] = vcenter

	obj, err := getObjectObjectUserVcenterRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserVcenterRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserVcenterRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserVcenterRule resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserVcenterRuleRead(d, m)
}

func resourceObjectUserVcenterRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	vcenter := d.Get("vcenter").(string)
	paradict["vcenter"] = vcenter

	obj, err := getObjectObjectUserVcenterRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserVcenterRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserVcenterRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserVcenterRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectUserVcenterRuleRead(d, m)
}

func resourceObjectUserVcenterRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	vcenter := d.Get("vcenter").(string)
	paradict["vcenter"] = vcenter

	err = c.DeleteObjectUserVcenterRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserVcenterRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserVcenterRuleRead(d *schema.ResourceData, m interface{}) error {
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

	vcenter := d.Get("vcenter").(string)
	if vcenter == "" {
		vcenter = importOptionChecking(m.(*FortiClient).Cfg, "vcenter")
		if vcenter == "" {
			return fmt.Errorf("Parameter vcenter is missing")
		}
		if err = d.Set("vcenter", vcenter); err != nil {
			return fmt.Errorf("Error set params vcenter: %v", err)
		}
	}
	paradict["vcenter"] = vcenter

	o, err := c.ReadObjectUserVcenterRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserVcenterRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserVcenterRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserVcenterRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserVcenterRuleName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserVcenterRuleRule2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserVcenterRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectUserVcenterRuleName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectUserVcenterRule-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("rule", flattenObjectUserVcenterRuleRule2edl(o["rule"], d, "rule")); err != nil {
		if vv, ok := fortiAPIPatch(o["rule"], "ObjectUserVcenterRule-Rule"); ok {
			if err = d.Set("rule", vv); err != nil {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}

	return nil
}

func flattenObjectUserVcenterRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserVcenterRuleName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserVcenterRuleRule2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserVcenterRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectUserVcenterRuleName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectUserVcenterRuleRule2edl(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
