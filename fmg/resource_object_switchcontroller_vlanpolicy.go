// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerVlanPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerVlanPolicyCreate,
		Read:   resourceObjectSwitchControllerVlanPolicyRead,
		Update: resourceObjectSwitchControllerVlanPolicyUpdate,
		Delete: resourceObjectSwitchControllerVlanPolicyDelete,

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
			"allowed_vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"allowed_vlans_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discard_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"untagged_vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerVlanPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerVlanPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerVlanPolicy resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerVlanPolicy(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerVlanPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerVlanPolicyRead(d, m)
}

func resourceObjectSwitchControllerVlanPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerVlanPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerVlanPolicy resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerVlanPolicy(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerVlanPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerVlanPolicyRead(d, m)
}

func resourceObjectSwitchControllerVlanPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerVlanPolicy(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerVlanPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerVlanPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerVlanPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerVlanPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerVlanPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerVlanPolicy resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerVlanPolicyAllowedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerVlanPolicyAllowedVlansAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerVlanPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerVlanPolicyDiscardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerVlanPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerVlanPolicyUntaggedVlans(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerVlanPolicyVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerVlanPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("allowed_vlans", flattenObjectSwitchControllerVlanPolicyAllowedVlans(o["allowed-vlans"], d, "allowed_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-vlans"], "ObjectSwitchControllerVlanPolicy-AllowedVlans"); ok {
			if err = d.Set("allowed_vlans", vv); err != nil {
				return fmt.Errorf("Error reading allowed_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_vlans: %v", err)
		}
	}

	if err = d.Set("allowed_vlans_all", flattenObjectSwitchControllerVlanPolicyAllowedVlansAll(o["allowed-vlans-all"], d, "allowed_vlans_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowed-vlans-all"], "ObjectSwitchControllerVlanPolicy-AllowedVlansAll"); ok {
			if err = d.Set("allowed_vlans_all", vv); err != nil {
				return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerVlanPolicyDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerVlanPolicy-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("discard_mode", flattenObjectSwitchControllerVlanPolicyDiscardMode(o["discard-mode"], d, "discard_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["discard-mode"], "ObjectSwitchControllerVlanPolicy-DiscardMode"); ok {
			if err = d.Set("discard_mode", vv); err != nil {
				return fmt.Errorf("Error reading discard_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading discard_mode: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerVlanPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerVlanPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("untagged_vlans", flattenObjectSwitchControllerVlanPolicyUntaggedVlans(o["untagged-vlans"], d, "untagged_vlans")); err != nil {
		if vv, ok := fortiAPIPatch(o["untagged-vlans"], "ObjectSwitchControllerVlanPolicy-UntaggedVlans"); ok {
			if err = d.Set("untagged_vlans", vv); err != nil {
				return fmt.Errorf("Error reading untagged_vlans: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untagged_vlans: %v", err)
		}
	}

	if err = d.Set("vlan", flattenObjectSwitchControllerVlanPolicyVlan(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "ObjectSwitchControllerVlanPolicy-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerVlanPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerVlanPolicyAllowedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerVlanPolicyAllowedVlansAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerVlanPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerVlanPolicyDiscardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerVlanPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerVlanPolicyUntaggedVlans(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerVlanPolicyVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerVlanPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowed_vlans"); ok || d.HasChange("allowed_vlans") {
		t, err := expandObjectSwitchControllerVlanPolicyAllowedVlans(d, v, "allowed_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans"] = t
		}
	}

	if v, ok := d.GetOk("allowed_vlans_all"); ok || d.HasChange("allowed_vlans_all") {
		t, err := expandObjectSwitchControllerVlanPolicyAllowedVlansAll(d, v, "allowed_vlans_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans-all"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerVlanPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("discard_mode"); ok || d.HasChange("discard_mode") {
		t, err := expandObjectSwitchControllerVlanPolicyDiscardMode(d, v, "discard_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discard-mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerVlanPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("untagged_vlans"); ok || d.HasChange("untagged_vlans") {
		t, err := expandObjectSwitchControllerVlanPolicyUntaggedVlans(d, v, "untagged_vlans")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandObjectSwitchControllerVlanPolicyVlan(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
