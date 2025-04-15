// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv4 prefix list rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterPrefixListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterPrefixListRuleCreate,
		Read:   resourceObjectRouterPrefixListRuleRead,
		Update: resourceObjectRouterPrefixListRuleUpdate,
		Delete: resourceObjectRouterPrefixListRuleDelete,

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
			"prefix_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ge": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"le": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectRouterPrefixListRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	prefix_list := d.Get("prefix_list").(string)
	paradict["prefix_list"] = prefix_list

	obj, err := getObjectObjectRouterPrefixListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterPrefixListRule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectRouterPrefixListRule(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterPrefixListRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterPrefixListRuleRead(d, m)
}

func resourceObjectRouterPrefixListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	prefix_list := d.Get("prefix_list").(string)
	paradict["prefix_list"] = prefix_list

	obj, err := getObjectObjectRouterPrefixListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterPrefixListRule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectRouterPrefixListRule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterPrefixListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterPrefixListRuleRead(d, m)
}

func resourceObjectRouterPrefixListRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	prefix_list := d.Get("prefix_list").(string)
	paradict["prefix_list"] = prefix_list

	wsParams["adom"] = adomv

	err = c.DeleteObjectRouterPrefixListRule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterPrefixListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterPrefixListRuleRead(d *schema.ResourceData, m interface{}) error {
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

	prefix_list := d.Get("prefix_list").(string)
	if prefix_list == "" {
		prefix_list = importOptionChecking(m.(*FortiClient).Cfg, "prefix_list")
		if prefix_list == "" {
			return fmt.Errorf("Parameter prefix_list is missing")
		}
		if err = d.Set("prefix_list", prefix_list); err != nil {
			return fmt.Errorf("Error set params prefix_list: %v", err)
		}
	}
	paradict["prefix_list"] = prefix_list

	o, err := c.ReadObjectRouterPrefixListRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterPrefixListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterPrefixListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterPrefixListRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterPrefixListRuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleGe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleLe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRulePrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectRouterPrefixListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectRouterPrefixListRuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectRouterPrefixListRule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("flags", flattenObjectRouterPrefixListRuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "ObjectRouterPrefixListRule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("ge", flattenObjectRouterPrefixListRuleGe2edl(o["ge"], d, "ge")); err != nil {
		if vv, ok := fortiAPIPatch(o["ge"], "ObjectRouterPrefixListRule-Ge"); ok {
			if err = d.Set("ge", vv); err != nil {
				return fmt.Errorf("Error reading ge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectRouterPrefixListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectRouterPrefixListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("le", flattenObjectRouterPrefixListRuleLe2edl(o["le"], d, "le")); err != nil {
		if vv, ok := fortiAPIPatch(o["le"], "ObjectRouterPrefixListRule-Le"); ok {
			if err = d.Set("le", vv); err != nil {
				return fmt.Errorf("Error reading le: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix", flattenObjectRouterPrefixListRulePrefix2edl(o["prefix"], d, "prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix"], "ObjectRouterPrefixListRule-Prefix"); ok {
			if err = d.Set("prefix", vv); err != nil {
				return fmt.Errorf("Error reading prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	return nil
}

func flattenObjectRouterPrefixListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterPrefixListRuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleGe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleLe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRulePrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectObjectRouterPrefixListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectRouterPrefixListRuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandObjectRouterPrefixListRuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("ge"); ok || d.HasChange("ge") {
		t, err := expandObjectRouterPrefixListRuleGe2edl(d, v, "ge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ge"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectRouterPrefixListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("le"); ok || d.HasChange("le") {
		t, err := expandObjectRouterPrefixListRuleLe2edl(d, v, "le")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["le"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok || d.HasChange("prefix") {
		t, err := expandObjectRouterPrefixListRulePrefix2edl(d, v, "prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	return &obj, nil
}
