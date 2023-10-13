// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Bonjour policy list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerBonjourProfilePolicyListMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerBonjourProfilePolicyListMoveUpdate,
		Read:   resourceObjectWirelessControllerBonjourProfilePolicyListMoveRead,
		Update: resourceObjectWirelessControllerBonjourProfilePolicyListMoveUpdate,
		Delete: resourceObjectWirelessControllerBonjourProfilePolicyListMoveDelete,

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
			"bonjour_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_list": &schema.Schema{
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

func resourceObjectWirelessControllerBonjourProfilePolicyListMoveUpdate(d *schema.ResourceData, m interface{}) error {
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

	bonjour_profile := d.Get("bonjour_profile").(string)
	policy_list := d.Get("policy_list").(string)
	paradict["bonjour_profile"] = bonjour_profile
	paradict["policy_list"] = policy_list

	target := d.Get("target").(string)
	obj, err := getObjectObjectWirelessControllerBonjourProfilePolicyListMove(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBonjourProfilePolicyListMove resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerBonjourProfilePolicyListMove(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBonjourProfilePolicyListMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerBonjourProfilePolicyListMove" + "_" + bonjour_profile + "_" + policy_list + "_" + target)

	return resourceObjectWirelessControllerBonjourProfilePolicyListMoveRead(d, m)
}

func resourceObjectWirelessControllerBonjourProfilePolicyListMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerBonjourProfilePolicyListMoveRead(d *schema.ResourceData, m interface{}) error {
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

	sid := d.Get("policy_list").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	bonjour_profile := d.Get("bonjour_profile").(string)
	if bonjour_profile == "" {
		bonjour_profile = importOptionChecking(m.(*FortiClient).Cfg, "bonjour_profile")
		if bonjour_profile == "" {
			return fmt.Errorf("Parameter bonjour_profile is missing")
		}
		if err = d.Set("bonjour_profile", bonjour_profile); err != nil {
			return fmt.Errorf("Error set params bonjour_profile: %v", err)
		}
	}
	paradict["bonjour_profile"] = bonjour_profile

	o, err := c.ReadObjectWirelessControllerBonjourProfilePolicyListMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfilePolicyListMove resource: %v", err)
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
				if _, ok := v["policy-id"]; !ok {
					return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfilePolicyListMove resource: policy_id doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["policy-id"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfilePolicyListMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "policy_id(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "policy_id(" + sid + ") was deleted"
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
					state_pos = "policy_id(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "policy_id(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenObjectWirelessControllerBonjourProfilePolicyListMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerBonjourProfilePolicyListMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerBonjourProfilePolicyListMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
