// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 prefix list rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterPrefixList6Rule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterPrefixList6RuleCreate,
		Read:   resourceObjectRouterPrefixList6RuleRead,
		Update: resourceObjectRouterPrefixList6RuleUpdate,
		Delete: resourceObjectRouterPrefixList6RuleDelete,

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
			"prefix_list6": &schema.Schema{
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
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectRouterPrefixList6RuleCreate(d *schema.ResourceData, m interface{}) error {
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

	prefix_list6 := d.Get("prefix_list6").(string)
	paradict["prefix_list6"] = prefix_list6

	obj, err := getObjectObjectRouterPrefixList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterPrefixList6Rule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectRouterPrefixList6Rule(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterPrefixList6Rule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterPrefixList6RuleRead(d, m)
}

func resourceObjectRouterPrefixList6RuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	prefix_list6 := d.Get("prefix_list6").(string)
	paradict["prefix_list6"] = prefix_list6

	obj, err := getObjectObjectRouterPrefixList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterPrefixList6Rule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectRouterPrefixList6Rule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterPrefixList6Rule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterPrefixList6RuleRead(d, m)
}

func resourceObjectRouterPrefixList6RuleDelete(d *schema.ResourceData, m interface{}) error {
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

	prefix_list6 := d.Get("prefix_list6").(string)
	paradict["prefix_list6"] = prefix_list6

	wsParams["adom"] = adomv

	err = c.DeleteObjectRouterPrefixList6Rule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterPrefixList6Rule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterPrefixList6RuleRead(d *schema.ResourceData, m interface{}) error {
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

	prefix_list6 := d.Get("prefix_list6").(string)
	if prefix_list6 == "" {
		prefix_list6 = importOptionChecking(m.(*FortiClient).Cfg, "prefix_list6")
		if prefix_list6 == "" {
			return fmt.Errorf("Parameter prefix_list6 is missing")
		}
		if err = d.Set("prefix_list6", prefix_list6); err != nil {
			return fmt.Errorf("Error set params prefix_list6: %v", err)
		}
	}
	paradict["prefix_list6"] = prefix_list6

	o, err := c.ReadObjectRouterPrefixList6Rule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterPrefixList6Rule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterPrefixList6Rule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterPrefixList6Rule resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterPrefixList6RuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixList6RuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixList6RuleGe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixList6RuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixList6RuleLe2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixList6RulePrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectRouterPrefixList6Rule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectRouterPrefixList6RuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectRouterPrefixList6Rule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("flags", flattenObjectRouterPrefixList6RuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "ObjectRouterPrefixList6Rule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("ge", flattenObjectRouterPrefixList6RuleGe2edl(o["ge"], d, "ge")); err != nil {
		if vv, ok := fortiAPIPatch(o["ge"], "ObjectRouterPrefixList6Rule-Ge"); ok {
			if err = d.Set("ge", vv); err != nil {
				return fmt.Errorf("Error reading ge: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ge: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectRouterPrefixList6RuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectRouterPrefixList6Rule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("le", flattenObjectRouterPrefixList6RuleLe2edl(o["le"], d, "le")); err != nil {
		if vv, ok := fortiAPIPatch(o["le"], "ObjectRouterPrefixList6Rule-Le"); ok {
			if err = d.Set("le", vv); err != nil {
				return fmt.Errorf("Error reading le: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading le: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenObjectRouterPrefixList6RulePrefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "ObjectRouterPrefixList6Rule-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenObjectRouterPrefixList6RuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterPrefixList6RuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixList6RuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixList6RuleGe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixList6RuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixList6RuleLe2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixList6RulePrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectRouterPrefixList6Rule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectRouterPrefixList6RuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandObjectRouterPrefixList6RuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("ge"); ok || d.HasChange("ge") {
		t, err := expandObjectRouterPrefixList6RuleGe2edl(d, v, "ge")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ge"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectRouterPrefixList6RuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("le"); ok || d.HasChange("le") {
		t, err := expandObjectRouterPrefixList6RuleLe2edl(d, v, "le")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["le"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandObjectRouterPrefixList6RulePrefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
