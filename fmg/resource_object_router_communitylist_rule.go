// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Community list rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterCommunityListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterCommunityListRuleCreate,
		Read:   resourceObjectRouterCommunityListRuleRead,
		Update: resourceObjectRouterCommunityListRuleUpdate,
		Delete: resourceObjectRouterCommunityListRuleDelete,

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
			"community_list": &schema.Schema{
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
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"regexp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectRouterCommunityListRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	community_list := d.Get("community_list").(string)
	paradict["community_list"] = community_list

	obj, err := getObjectObjectRouterCommunityListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterCommunityListRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectRouterCommunityListRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterCommunityListRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterCommunityListRuleRead(d, m)
}

func resourceObjectRouterCommunityListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	community_list := d.Get("community_list").(string)
	paradict["community_list"] = community_list

	obj, err := getObjectObjectRouterCommunityListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterCommunityListRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectRouterCommunityListRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterCommunityListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterCommunityListRuleRead(d, m)
}

func resourceObjectRouterCommunityListRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	community_list := d.Get("community_list").(string)
	paradict["community_list"] = community_list

	err = c.DeleteObjectRouterCommunityListRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterCommunityListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterCommunityListRuleRead(d *schema.ResourceData, m interface{}) error {
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

	community_list := d.Get("community_list").(string)
	if community_list == "" {
		community_list = importOptionChecking(m.(*FortiClient).Cfg, "community_list")
		if community_list == "" {
			return fmt.Errorf("Parameter community_list is missing")
		}
		if err = d.Set("community_list", community_list); err != nil {
			return fmt.Errorf("Error set params community_list: %v", err)
		}
	}
	paradict["community_list"] = community_list

	o, err := c.ReadObjectRouterCommunityListRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterCommunityListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterCommunityListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterCommunityListRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterCommunityListRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterCommunityListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterCommunityListRuleMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterCommunityListRuleRegexp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectRouterCommunityListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectRouterCommunityListRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectRouterCommunityListRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectRouterCommunityListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectRouterCommunityListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match", flattenObjectRouterCommunityListRuleMatch2edl(o["match"], d, "match")); err != nil {
		if vv, ok := fortiAPIPatch(o["match"], "ObjectRouterCommunityListRule-Match"); ok {
			if err = d.Set("match", vv); err != nil {
				return fmt.Errorf("Error reading match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match: %v", err)
		}
	}

	if err = d.Set("regexp", flattenObjectRouterCommunityListRuleRegexp2edl(o["regexp"], d, "regexp")); err != nil {
		if vv, ok := fortiAPIPatch(o["regexp"], "ObjectRouterCommunityListRule-Regexp"); ok {
			if err = d.Set("regexp", vv); err != nil {
				return fmt.Errorf("Error reading regexp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading regexp: %v", err)
		}
	}

	return nil
}

func flattenObjectRouterCommunityListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterCommunityListRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterCommunityListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterCommunityListRuleMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterCommunityListRuleRegexp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectRouterCommunityListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectRouterCommunityListRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectRouterCommunityListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandObjectRouterCommunityListRuleMatch2edl(d, v, "match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("regexp"); ok || d.HasChange("regexp") {
		t, err := expandObjectRouterCommunityListRuleRegexp2edl(d, v, "regexp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regexp"] = t
		}
	}

	return &obj, nil
}
