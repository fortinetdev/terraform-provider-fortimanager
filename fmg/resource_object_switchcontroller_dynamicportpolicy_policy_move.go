// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Port policies with matching criteria and actions.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerDynamicPortPolicyPolicyMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveUpdate,
		Read:   resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveRead,
		Update: resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveUpdate,
		Delete: resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"dynamic_port_policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveUpdate(d *schema.ResourceData, m interface{}) error {
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

	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	policy := d.Get("policy").(string)
	paradict["dynamic_port_policy"] = dynamic_port_policy
	paradict["policy"] = policy

	target := d.Get("target").(string)
	obj, err := getObjectObjectSwitchControllerDynamicPortPolicyPolicyMove(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDynamicPortPolicyPolicyMove resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerDynamicPortPolicyPolicyMove(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerDynamicPortPolicyPolicyMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSwitchControllerDynamicPortPolicyPolicyMove" + "_" + dynamic_port_policy + "_" + policy + "_" + target)

	return resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveRead(d, m)
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerDynamicPortPolicyPolicyMoveRead(d *schema.ResourceData, m interface{}) error {
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

	sid := d.Get("policy").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	dynamic_port_policy := d.Get("dynamic_port_policy").(string)
	if dynamic_port_policy == "" {
		dynamic_port_policy = importOptionChecking(m.(*FortiClient).Cfg, "dynamic_port_policy")
		if dynamic_port_policy == "" {
			return fmt.Errorf("Parameter dynamic_port_policy is missing")
		}
		if err = d.Set("dynamic_port_policy", dynamic_port_policy); err != nil {
			return fmt.Errorf("Error set params dynamic_port_policy: %v", err)
		}
	}
	paradict["dynamic_port_policy"] = dynamic_port_policy

	o, err := c.ReadObjectSwitchControllerDynamicPortPolicyPolicyMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicyPolicyMove resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if o != nil {
		i := 1
		now_sid := -1
		now_did := -1

		for _, z := range o {
			if v, ok := z.(map[string]interface{}); ok {
				if _, ok := v["name"]; !ok {
					return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicyPolicyMove resource: name doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["name"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading ObjectSwitchControllerDynamicPortPolicyPolicyMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "name(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "name(" + sid + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + did + ") was deleted"
			}
		} else {
			bconsistent := true
			if action == "before" {
				if now_sid != now_did-1 {
					bconsistent = false
				}
			}

			if action == "after" {
				if now_sid != now_did+1 {
					bconsistent = false
				}
			}

			if bconsistent == false {
				relative_pos := now_sid - now_did

				if relative_pos > 0 {
					state_pos = "name(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "name(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenObjectSwitchControllerDynamicPortPolicyPolicyMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerDynamicPortPolicyPolicyMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerDynamicPortPolicyPolicyMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandObjectSwitchControllerDynamicPortPolicyPolicyMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
