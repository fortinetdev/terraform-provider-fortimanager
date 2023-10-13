// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: CASB user activity rules.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectCasbUserActivityMatchRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectCasbUserActivityMatchRulesCreate,
		Read:   resourceObjectCasbUserActivityMatchRulesRead,
		Update: resourceObjectCasbUserActivityMatchRulesUpdate,
		Delete: resourceObjectCasbUserActivityMatchRulesDelete,

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
			"user_activity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"match_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"methods": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectCasbUserActivityMatchRulesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectObjectCasbUserActivityMatchRules(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatchRules resource while getting object: %v", err)
	}

	_, err = c.CreateObjectCasbUserActivityMatchRules(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectCasbUserActivityMatchRules resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchRulesRead(d, m)
}

func resourceObjectCasbUserActivityMatchRulesUpdate(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	obj, err := getObjectObjectCasbUserActivityMatchRules(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchRules resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectCasbUserActivityMatchRules(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectCasbUserActivityMatchRules resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectCasbUserActivityMatchRulesRead(d, m)
}

func resourceObjectCasbUserActivityMatchRulesDelete(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	err = c.DeleteObjectCasbUserActivityMatchRules(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectCasbUserActivityMatchRules resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectCasbUserActivityMatchRulesRead(d *schema.ResourceData, m interface{}) error {
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

	user_activity := d.Get("user_activity").(string)
	match := d.Get("match").(string)
	if user_activity == "" {
		user_activity = importOptionChecking(m.(*FortiClient).Cfg, "user_activity")
		if user_activity == "" {
			return fmt.Errorf("Parameter user_activity is missing")
		}
		if err = d.Set("user_activity", user_activity); err != nil {
			return fmt.Errorf("Error set params user_activity: %v", err)
		}
	}
	if match == "" {
		match = importOptionChecking(m.(*FortiClient).Cfg, "match")
		if match == "" {
			return fmt.Errorf("Parameter match is missing")
		}
		if err = d.Set("match", match); err != nil {
			return fmt.Errorf("Error set params match: %v", err)
		}
	}
	paradict["user_activity"] = user_activity
	paradict["match"] = match

	o, err := c.ReadObjectCasbUserActivityMatchRules(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchRules resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectCasbUserActivityMatchRules(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectCasbUserActivityMatchRules resource from API: %v", err)
	}
	return nil
}

func flattenObjectCasbUserActivityMatchRulesCaseSensitive3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesDomains3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbUserActivityMatchRulesHeaderName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesMatchPattern3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesMatchValue3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesMethods3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectCasbUserActivityMatchRulesNegate3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectCasbUserActivityMatchRulesType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectCasbUserActivityMatchRules(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("case_sensitive", flattenObjectCasbUserActivityMatchRulesCaseSensitive3rdl(o["case-sensitive"], d, "case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["case-sensitive"], "ObjectCasbUserActivityMatchRules-CaseSensitive"); ok {
			if err = d.Set("case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}

	if err = d.Set("domains", flattenObjectCasbUserActivityMatchRulesDomains3rdl(o["domains"], d, "domains")); err != nil {
		if vv, ok := fortiAPIPatch(o["domains"], "ObjectCasbUserActivityMatchRules-Domains"); ok {
			if err = d.Set("domains", vv); err != nil {
				return fmt.Errorf("Error reading domains: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domains: %v", err)
		}
	}

	if err = d.Set("header_name", flattenObjectCasbUserActivityMatchRulesHeaderName3rdl(o["header-name"], d, "header_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-name"], "ObjectCasbUserActivityMatchRules-HeaderName"); ok {
			if err = d.Set("header_name", vv); err != nil {
				return fmt.Errorf("Error reading header_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectCasbUserActivityMatchRulesId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectCasbUserActivityMatchRules-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("match_pattern", flattenObjectCasbUserActivityMatchRulesMatchPattern3rdl(o["match-pattern"], d, "match_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-pattern"], "ObjectCasbUserActivityMatchRules-MatchPattern"); ok {
			if err = d.Set("match_pattern", vv); err != nil {
				return fmt.Errorf("Error reading match_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_pattern: %v", err)
		}
	}

	if err = d.Set("match_value", flattenObjectCasbUserActivityMatchRulesMatchValue3rdl(o["match-value"], d, "match_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["match-value"], "ObjectCasbUserActivityMatchRules-MatchValue"); ok {
			if err = d.Set("match_value", vv); err != nil {
				return fmt.Errorf("Error reading match_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading match_value: %v", err)
		}
	}

	if err = d.Set("methods", flattenObjectCasbUserActivityMatchRulesMethods3rdl(o["methods"], d, "methods")); err != nil {
		if vv, ok := fortiAPIPatch(o["methods"], "ObjectCasbUserActivityMatchRules-Methods"); ok {
			if err = d.Set("methods", vv); err != nil {
				return fmt.Errorf("Error reading methods: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("negate", flattenObjectCasbUserActivityMatchRulesNegate3rdl(o["negate"], d, "negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["negate"], "ObjectCasbUserActivityMatchRules-Negate"); ok {
			if err = d.Set("negate", vv); err != nil {
				return fmt.Errorf("Error reading negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading negate: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectCasbUserActivityMatchRulesType3rdl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectCasbUserActivityMatchRules-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectCasbUserActivityMatchRulesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectCasbUserActivityMatchRulesCaseSensitive3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesDomains3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbUserActivityMatchRulesHeaderName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesMatchPattern3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesMatchValue3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesMethods3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectCasbUserActivityMatchRulesNegate3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectCasbUserActivityMatchRulesType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectCasbUserActivityMatchRules(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("case_sensitive"); ok || d.HasChange("case_sensitive") {
		t, err := expandObjectCasbUserActivityMatchRulesCaseSensitive3rdl(d, v, "case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("domains"); ok || d.HasChange("domains") {
		t, err := expandObjectCasbUserActivityMatchRulesDomains3rdl(d, v, "domains")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domains"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok || d.HasChange("header_name") {
		t, err := expandObjectCasbUserActivityMatchRulesHeaderName3rdl(d, v, "header_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectCasbUserActivityMatchRulesId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("match_pattern"); ok || d.HasChange("match_pattern") {
		t, err := expandObjectCasbUserActivityMatchRulesMatchPattern3rdl(d, v, "match_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-pattern"] = t
		}
	}

	if v, ok := d.GetOk("match_value"); ok || d.HasChange("match_value") {
		t, err := expandObjectCasbUserActivityMatchRulesMatchValue3rdl(d, v, "match_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-value"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok || d.HasChange("methods") {
		t, err := expandObjectCasbUserActivityMatchRulesMethods3rdl(d, v, "methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("negate"); ok || d.HasChange("negate") {
		t, err := expandObjectCasbUserActivityMatchRulesNegate3rdl(d, v, "negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negate"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectCasbUserActivityMatchRulesType3rdl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
