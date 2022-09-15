// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure consolidated IPv4/IPv6 policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesFirewallConsolidatedPolicyMove() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallConsolidatedPolicyMoveUpdate,
		Read:   resourcePackagesFirewallConsolidatedPolicyMoveRead,
		Update: resourcePackagesFirewallConsolidatedPolicyMoveUpdate,
		Delete: resourcePackagesFirewallConsolidatedPolicyMoveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"state_pos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
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
			"pkg": &schema.Schema{
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

func resourcePackagesFirewallConsolidatedPolicyMoveUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	policy := d.Get("policy").(string)
	var paralist []string
	paralist = append(paralist, pkg)
	paralist = append(paralist, policy)

	target := d.Get("target").(string)
	obj, err := getObjectPackagesFirewallConsolidatedPolicyMove(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallConsolidatedPolicyMove resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallConsolidatedPolicyMove(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallConsolidatedPolicyMove resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("PackagesFirewallConsolidatedPolicyMove" + "_" + pkg + "_" + policy + "_" + target)

	return resourcePackagesFirewallConsolidatedPolicyMoveRead(d, m)
}

func resourcePackagesFirewallConsolidatedPolicyMoveDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourcePackagesFirewallConsolidatedPolicyMoveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	sid, err := strconv.Atoi(d.Get("policy").(string))
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallConsolidatedPolicyMove resource: %v", err)
	}
	did, err := strconv.Atoi(d.Get("target").(string))
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallConsolidatedPolicyMove resource: %v", err)
	}
	action := d.Get("option").(string)

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesFirewallConsolidatedPolicyMove(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallConsolidatedPolicyMove resource: %v", err)
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
				if _, ok := v["policyid"]; !ok {
					return fmt.Errorf("Error reading PackagesFirewallConsolidatedPolicyMove resource: policyid doesn't exist.")
				}

				idn := -1
				if vidn, ok := v["policyid"].(float64); !ok {
					return fmt.Errorf("Error reading PackagesFirewallConsolidatedPolicyMove resource: wrong policyid.")
				} else {
					idn = int(vidn)
				}

				if idn == sid {
					now_sid = i
				}

				if idn == did {
					now_did = i
				}
			} else {
				return fmt.Errorf("Error reading PackagesFirewallConsolidatedPolicyMove resource: no valid map string.")
			}

			i += 1
		}

		state_pos := ""

		if now_sid == -1 || now_did == -1 {
			if now_sid == -1 && now_did == -1 {
				state_pos = "policyid(" + strconv.Itoa(sid) + ") and target(" + strconv.Itoa(did) + ") were deleted"
			} else if now_sid == -1 {
				state_pos = "policyid(" + strconv.Itoa(sid) + ") was deleted"
			} else if now_did == -1 {
				state_pos = "target(" + strconv.Itoa(did) + ") was deleted"
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
					state_pos = "policyid(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(relative_pos) + " behind target(" + strconv.Itoa(did) + ")"
				} else {
					state_pos = "policyid(" + strconv.Itoa(sid) + ") is " + strconv.Itoa(-relative_pos) + " ahead of target(" + strconv.Itoa(did) + ")"
				}
			}
		}

		d.Set("state_pos", state_pos)
	}

	return nil
}

func flattenPackagesFirewallConsolidatedPolicyMoveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallConsolidatedPolicyMoveTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallConsolidatedPolicyMoveOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallConsolidatedPolicyMove(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandPackagesFirewallConsolidatedPolicyMoveTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandPackagesFirewallConsolidatedPolicyMoveOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	return &obj, nil
}
