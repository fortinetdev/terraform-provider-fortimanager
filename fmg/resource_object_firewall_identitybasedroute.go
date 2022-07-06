// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure identity based routing.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallIdentityBasedRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallIdentityBasedRouteCreate,
		Read:   resourceObjectFirewallIdentityBasedRouteRead,
		Update: resourceObjectFirewallIdentityBasedRouteUpdate,
		Delete: resourceObjectFirewallIdentityBasedRouteDelete,

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
						"device": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"groups": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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

func resourceObjectFirewallIdentityBasedRouteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallIdentityBasedRoute(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIdentityBasedRoute resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallIdentityBasedRoute(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIdentityBasedRoute resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallIdentityBasedRouteRead(d, m)
}

func resourceObjectFirewallIdentityBasedRouteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallIdentityBasedRoute(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIdentityBasedRoute resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallIdentityBasedRoute(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIdentityBasedRoute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallIdentityBasedRouteRead(d, m)
}

func resourceObjectFirewallIdentityBasedRouteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallIdentityBasedRoute(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallIdentityBasedRoute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallIdentityBasedRouteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallIdentityBasedRoute(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIdentityBasedRoute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallIdentityBasedRoute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIdentityBasedRoute resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallIdentityBasedRouteComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			v := flattenObjectFirewallIdentityBasedRouteRuleDevice(i["device"], d, pre_append)
			tmp["device"] = fortiAPISubPartPatch(v, "ObjectFirewallIdentityBasedRoute-Rule-Device")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenObjectFirewallIdentityBasedRouteRuleGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "ObjectFirewallIdentityBasedRoute-Rule-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenObjectFirewallIdentityBasedRouteRuleGroups(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "ObjectFirewallIdentityBasedRoute-Rule-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectFirewallIdentityBasedRouteRuleId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectFirewallIdentityBasedRoute-Rule-Id")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectFirewallIdentityBasedRouteRuleDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteRuleGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteRuleGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallIdentityBasedRoute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("comments", flattenObjectFirewallIdentityBasedRouteComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectFirewallIdentityBasedRoute-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallIdentityBasedRouteName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallIdentityBasedRoute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rule", flattenObjectFirewallIdentityBasedRouteRule(o["rule"], d, "rule")); err != nil {
			if vv, ok := fortiAPIPatch(o["rule"], "ObjectFirewallIdentityBasedRoute-Rule"); ok {
				if err = d.Set("rule", vv); err != nil {
					return fmt.Errorf("Error reading rule: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rule"); ok {
			if err = d.Set("rule", flattenObjectFirewallIdentityBasedRouteRule(o["rule"], d, "rule")); err != nil {
				if vv, ok := fortiAPIPatch(o["rule"], "ObjectFirewallIdentityBasedRoute-Rule"); ok {
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

func flattenObjectFirewallIdentityBasedRouteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallIdentityBasedRouteComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device"], _ = expandObjectFirewallIdentityBasedRouteRuleDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandObjectFirewallIdentityBasedRouteRuleGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandObjectFirewallIdentityBasedRouteRuleGroups(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandObjectFirewallIdentityBasedRouteRuleId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectFirewallIdentityBasedRouteRuleDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteRuleGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteRuleGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteRuleId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallIdentityBasedRoute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandObjectFirewallIdentityBasedRouteComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallIdentityBasedRouteName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("rule"); ok || d.HasChange("rule") {
		t, err := expandObjectFirewallIdentityBasedRouteRule(d, v, "rule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}

	return &obj, nil
}
