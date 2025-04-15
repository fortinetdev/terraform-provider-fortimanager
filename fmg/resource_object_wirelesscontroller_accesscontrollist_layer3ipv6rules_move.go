// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AP ACL layer3 ipv6 rule list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveUpdate,
		Read:   resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveRead,
		Update: resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveUpdate,
		Delete: resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveDelete,

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
			"access_control_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"layer3_ipv6_rules": &schema.Schema{
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

func resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveUpdate(d *schema.ResourceData, m interface{}) error {
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

	access_control_list := d.Get("access_control_list").(string)
	layer3_ipv6_rules := d.Get("layer3_ipv6_rules").(string)
	paradict["access_control_list"] = access_control_list
	paradict["layer3_ipv6_rules"] = layer3_ipv6_rules

	target := d.Get("target").(string)
	obj, err := getObjectObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove" + "_" + access_control_list + "_" + layer3_ipv6_rules + "_" + target)

	return resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveRead(d, m)
}

func resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveRead(d *schema.ResourceData, m interface{}) error {
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

	sid := d.Get("layer3_ipv6_rules").(string)
	did := d.Get("target").(string)
	action := d.Get("option").(string)

	access_control_list := d.Get("access_control_list").(string)
	if access_control_list == "" {
		access_control_list = importOptionChecking(m.(*FortiClient).Cfg, "access_control_list")
		if access_control_list == "" {
			return fmt.Errorf("Parameter access_control_list is missing")
		}
		if err = d.Set("access_control_list", access_control_list); err != nil {
			return fmt.Errorf("Error set params access_control_list: %v", err)
		}
	}
	paradict["access_control_list"] = access_control_list

	o, err := c.ReadObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove resource: %v", err)
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
				if _, ok := v["rule-id"]; !ok {
					return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove resource: rule_id doesn't exist.")
				}

				vid := fmt.Sprintf("%v", v["rule-id"])

				if vid == sid {
					now_sid = i
				}

				if vid == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading ObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "rule_id(" + sid + ") and target(" + did + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "rule_id(" + sid + ") was deleted"
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
					state_pos = "rule_id(" + sid + ") is " + strconv.Itoa(relative_pos) + " behind target(" + did + ")"
				} else {
					state_pos = "rule_id(" + sid + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + did + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerAccessControlListLayer3Ipv6RulesMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandObjectWirelessControllerAccessControlListLayer3Ipv6RulesMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
