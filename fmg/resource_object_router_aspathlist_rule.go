// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AS path list rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterAspathListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterAspathListRuleCreate,
		Read:   resourceObjectRouterAspathListRuleRead,
		Update: resourceObjectRouterAspathListRuleUpdate,
		Delete: resourceObjectRouterAspathListRuleDelete,

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
			"aspath_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"regexp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectRouterAspathListRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	aspath_list := d.Get("aspath_list").(string)
	paradict["aspath_list"] = aspath_list

	obj, err := getObjectObjectRouterAspathListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterAspathListRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectRouterAspathListRule(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterAspathListRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterAspathListRuleRead(d, m)
}

func resourceObjectRouterAspathListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	aspath_list := d.Get("aspath_list").(string)
	paradict["aspath_list"] = aspath_list

	obj, err := getObjectObjectRouterAspathListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterAspathListRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectRouterAspathListRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterAspathListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterAspathListRuleRead(d, m)
}

func resourceObjectRouterAspathListRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	aspath_list := d.Get("aspath_list").(string)
	paradict["aspath_list"] = aspath_list

	wsParams["adom"] = adomv

	err = c.DeleteObjectRouterAspathListRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterAspathListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterAspathListRuleRead(d *schema.ResourceData, m interface{}) error {
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

	aspath_list := d.Get("aspath_list").(string)
	if aspath_list == "" {
		aspath_list = importOptionChecking(m.(*FortiClient).Cfg, "aspath_list")
		if aspath_list == "" {
			return fmt.Errorf("Parameter aspath_list is missing")
		}
		if err = d.Set("aspath_list", aspath_list); err != nil {
			return fmt.Errorf("Error set params aspath_list: %v", err)
		}
	}
	paradict["aspath_list"] = aspath_list

	o, err := c.ReadObjectRouterAspathListRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterAspathListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterAspathListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterAspathListRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterAspathListRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterAspathListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterAspathListRuleRegexp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectRouterAspathListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectRouterAspathListRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectRouterAspathListRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectRouterAspathListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectRouterAspathListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("regexp", flattenObjectRouterAspathListRuleRegexp2edl(o["regexp"], d, "regexp")); err != nil {
		if vv, ok := fortiAPIPatch(o["regexp"], "ObjectRouterAspathListRule-Regexp"); ok {
			if err = d.Set("regexp", vv); err != nil {
				return fmt.Errorf("Error reading regexp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regexp: %v", err)
		}
	}

	return nil
}

func flattenObjectRouterAspathListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterAspathListRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterAspathListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterAspathListRuleRegexp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectRouterAspathListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectRouterAspathListRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectRouterAspathListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("regexp"); ok || d.HasChange("regexp") {
		t, err := expandObjectRouterAspathListRuleRegexp2edl(d, v, "regexp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regexp"] = t
		}
	}

	return &obj, nil
}
