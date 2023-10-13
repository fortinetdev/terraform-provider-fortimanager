// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Rule.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectRouterAccessList6Rule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterAccessList6RuleCreate,
		Read:   resourceObjectRouterAccessList6RuleRead,
		Update: resourceObjectRouterAccessList6RuleUpdate,
		Delete: resourceObjectRouterAccessList6RuleDelete,

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
			"access_list6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exact_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectRouterAccessList6RuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	access_list6 := d.Get("access_list6").(string)
	paradict["access_list6"] = access_list6

	obj, err := getObjectObjectRouterAccessList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterAccessList6Rule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectRouterAccessList6Rule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterAccessList6Rule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterAccessList6RuleRead(d, m)
}

func resourceObjectRouterAccessList6RuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	access_list6 := d.Get("access_list6").(string)
	paradict["access_list6"] = access_list6

	obj, err := getObjectObjectRouterAccessList6Rule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterAccessList6Rule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectRouterAccessList6Rule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterAccessList6Rule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectRouterAccessList6RuleRead(d, m)
}

func resourceObjectRouterAccessList6RuleDelete(d *schema.ResourceData, m interface{}) error {
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

	access_list6 := d.Get("access_list6").(string)
	paradict["access_list6"] = access_list6

	err = c.DeleteObjectRouterAccessList6Rule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterAccessList6Rule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterAccessList6RuleRead(d *schema.ResourceData, m interface{}) error {
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

	access_list6 := d.Get("access_list6").(string)
	if access_list6 == "" {
		access_list6 = importOptionChecking(m.(*FortiClient).Cfg, "access_list6")
		if access_list6 == "" {
			return fmt.Errorf("Parameter access_list6 is missing")
		}
		if err = d.Set("access_list6", access_list6); err != nil {
			return fmt.Errorf("Error set params access_list6: %v", err)
		}
	}
	paradict["access_list6"] = access_list6

	o, err := c.ReadObjectRouterAccessList6Rule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterAccessList6Rule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterAccessList6Rule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterAccessList6Rule resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterAccessList6RuleAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterAccessList6RuleExactMatch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterAccessList6RuleFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterAccessList6RuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterAccessList6RulePrefix62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectRouterAccessList6Rule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenObjectRouterAccessList6RuleAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ObjectRouterAccessList6Rule-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("exact_match", flattenObjectRouterAccessList6RuleExactMatch2edl(o["exact-match"], d, "exact_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["exact-match"], "ObjectRouterAccessList6Rule-ExactMatch"); ok {
			if err = d.Set("exact_match", vv); err != nil {
				return fmt.Errorf("Error reading exact_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exact_match: %v", err)
		}
	}

	if err = d.Set("flags", flattenObjectRouterAccessList6RuleFlags2edl(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "ObjectRouterAccessList6Rule-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectRouterAccessList6RuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectRouterAccessList6Rule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenObjectRouterAccessList6RulePrefix62edl(o["prefix6"], d, "prefix6")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix6"], "ObjectRouterAccessList6Rule-Prefix6"); ok {
			if err = d.Set("prefix6", vv); err != nil {
				return fmt.Errorf("Error reading prefix6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	return nil
}

func flattenObjectRouterAccessList6RuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterAccessList6RuleAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterAccessList6RuleExactMatch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterAccessList6RuleFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterAccessList6RuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterAccessList6RulePrefix62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectRouterAccessList6Rule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandObjectRouterAccessList6RuleAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("exact_match"); ok || d.HasChange("exact_match") {
		t, err := expandObjectRouterAccessList6RuleExactMatch2edl(d, v, "exact_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exact-match"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandObjectRouterAccessList6RuleFlags2edl(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectRouterAccessList6RuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok || d.HasChange("prefix6") {
		t, err := expandObjectRouterAccessList6RulePrefix62edl(d, v, "prefix6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	return &obj, nil
}
