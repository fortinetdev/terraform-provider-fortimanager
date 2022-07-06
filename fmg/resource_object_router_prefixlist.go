// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 prefix lists.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectRouterPrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectRouterPrefixListCreate,
		Read:   resourceObjectRouterPrefixListRead,
		Update: resourceObjectRouterPrefixListUpdate,
		Delete: resourceObjectRouterPrefixListDelete,

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
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ge": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"le": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectRouterPrefixListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectRouterPrefixList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterPrefixList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectRouterPrefixList(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectRouterPrefixList resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectRouterPrefixListRead(d, m)
}

func resourceObjectRouterPrefixListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectRouterPrefixList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterPrefixList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectRouterPrefixList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectRouterPrefixList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectRouterPrefixListRead(d, m)
}

func resourceObjectRouterPrefixListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectRouterPrefixList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectRouterPrefixList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectRouterPrefixListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectRouterPrefixList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterPrefixList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectRouterPrefixList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectRouterPrefixList resource from API: %v", err)
	}
	return nil
}

func flattenObjectRouterPrefixListComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectRouterPrefixListRuleAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectRouterPrefixList-Rule-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			v := flattenObjectRouterPrefixListRuleFlags(i["flags"], d, pre_append)
			tmp["flags"] = fortiAPISubPartPatch(v, "ObjectRouterPrefixList-Rule-Flags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := i["ge"]; ok {
			v := flattenObjectRouterPrefixListRuleGe(i["ge"], d, pre_append)
			tmp["ge"] = fortiAPISubPartPatch(v, "ObjectRouterPrefixList-Rule-Ge")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectRouterPrefixListRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectRouterPrefixList-Rule-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := i["le"]; ok {
			v := flattenObjectRouterPrefixListRuleLe(i["le"], d, pre_append)
			tmp["le"] = fortiAPISubPartPatch(v, "ObjectRouterPrefixList-Rule-Le")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenObjectRouterPrefixListRulePrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "ObjectRouterPrefixList-Rule-Prefix")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectRouterPrefixListRuleAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleGe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRuleLe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectRouterPrefixListRulePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectRouterPrefixList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenObjectRouterPrefixListComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectRouterPrefixList-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectRouterPrefixListName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectRouterPrefixList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenObjectRouterPrefixListRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "ObjectRouterPrefixList-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenObjectRouterPrefixListRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "ObjectRouterPrefixList-Rule"); ok {
					if err = d.Set("rule", vv); err != nil {
						return fmt.Errorf("Error reading rule: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectRouterPrefixListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectRouterPrefixListComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandObjectRouterPrefixListRuleAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["flags"], _ = expandObjectRouterPrefixListRuleFlags(d, i["flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ge"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ge"], _ = expandObjectRouterPrefixListRuleGe(d, i["ge"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectRouterPrefixListRuleId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "le"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["le"], _ = expandObjectRouterPrefixListRuleLe(d, i["le"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandObjectRouterPrefixListRulePrefix(d, i["prefix"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectRouterPrefixListRuleAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleGe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRuleLe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectRouterPrefixListRulePrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectObjectRouterPrefixList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectRouterPrefixListComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectRouterPrefixListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectRouterPrefixListRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
