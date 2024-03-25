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

func resourceObjectFirewallIdentityBasedRouteRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallIdentityBasedRouteRuleCreate,
		Read:   resourceObjectFirewallIdentityBasedRouteRuleRead,
		Update: resourceObjectFirewallIdentityBasedRouteRuleUpdate,
		Delete: resourceObjectFirewallIdentityBasedRouteRuleDelete,

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
			"identity_based_route": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallIdentityBasedRouteRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	identity_based_route := d.Get("identity_based_route").(string)
	paradict["identity_based_route"] = identity_based_route

	obj, err := getObjectObjectFirewallIdentityBasedRouteRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIdentityBasedRouteRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallIdentityBasedRouteRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallIdentityBasedRouteRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallIdentityBasedRouteRuleRead(d, m)
}

func resourceObjectFirewallIdentityBasedRouteRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	identity_based_route := d.Get("identity_based_route").(string)
	paradict["identity_based_route"] = identity_based_route

	obj, err := getObjectObjectFirewallIdentityBasedRouteRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIdentityBasedRouteRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallIdentityBasedRouteRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallIdentityBasedRouteRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectFirewallIdentityBasedRouteRuleRead(d, m)
}

func resourceObjectFirewallIdentityBasedRouteRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	identity_based_route := d.Get("identity_based_route").(string)
	paradict["identity_based_route"] = identity_based_route

	err = c.DeleteObjectFirewallIdentityBasedRouteRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallIdentityBasedRouteRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallIdentityBasedRouteRuleRead(d *schema.ResourceData, m interface{}) error {
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

	identity_based_route := d.Get("identity_based_route").(string)
	if identity_based_route == "" {
		identity_based_route = importOptionChecking(m.(*FortiClient).Cfg, "identity_based_route")
		if identity_based_route == "" {
			return fmt.Errorf("Parameter identity_based_route is missing")
		}
		if err = d.Set("identity_based_route", identity_based_route); err != nil {
			return fmt.Errorf("Error set params identity_based_route: %v", err)
		}
	}
	paradict["identity_based_route"] = identity_based_route

	o, err := c.ReadObjectFirewallIdentityBasedRouteRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIdentityBasedRouteRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallIdentityBasedRouteRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallIdentityBasedRouteRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallIdentityBasedRouteRuleDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteRuleGateway2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallIdentityBasedRouteRuleGroups2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenObjectFirewallIdentityBasedRouteRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallIdentityBasedRouteRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("device", flattenObjectFirewallIdentityBasedRouteRuleDevice2edl(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "ObjectFirewallIdentityBasedRouteRule-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("gateway", flattenObjectFirewallIdentityBasedRouteRuleGateway2edl(o["gateway"], d, "gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway"], "ObjectFirewallIdentityBasedRouteRule-Gateway"); ok {
			if err = d.Set("gateway", vv); err != nil {
				return fmt.Errorf("Error reading gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("groups", flattenObjectFirewallIdentityBasedRouteRuleGroups2edl(o["groups"], d, "groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["groups"], "ObjectFirewallIdentityBasedRouteRule-Groups"); ok {
			if err = d.Set("groups", vv); err != nil {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectFirewallIdentityBasedRouteRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectFirewallIdentityBasedRouteRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallIdentityBasedRouteRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallIdentityBasedRouteRuleDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteRuleGateway2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallIdentityBasedRouteRuleGroups2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandObjectFirewallIdentityBasedRouteRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallIdentityBasedRouteRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandObjectFirewallIdentityBasedRouteRuleDevice2edl(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok || d.HasChange("gateway") {
		t, err := expandObjectFirewallIdentityBasedRouteRuleGateway2edl(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandObjectFirewallIdentityBasedRouteRuleGroups2edl(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectFirewallIdentityBasedRouteRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
