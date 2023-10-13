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

func resourceObjectWirelessControllerBonjourProfilePolicyList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerBonjourProfilePolicyListCreate,
		Read:   resourceObjectWirelessControllerBonjourProfilePolicyListRead,
		Update: resourceObjectWirelessControllerBonjourProfilePolicyListUpdate,
		Delete: resourceObjectWirelessControllerBonjourProfilePolicyListDelete,

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
			"bonjour_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"to_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectWirelessControllerBonjourProfilePolicyListCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["bonjour_profile"] = bonjour_profile

	obj, err := getObjectObjectWirelessControllerBonjourProfilePolicyList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerBonjourProfilePolicyList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerBonjourProfilePolicyList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policy_id")))

	return resourceObjectWirelessControllerBonjourProfilePolicyListRead(d, m)
}

func resourceObjectWirelessControllerBonjourProfilePolicyListUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["bonjour_profile"] = bonjour_profile

	obj, err := getObjectObjectWirelessControllerBonjourProfilePolicyList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBonjourProfilePolicyList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerBonjourProfilePolicyList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policy_id")))

	return resourceObjectWirelessControllerBonjourProfilePolicyListRead(d, m)
}

func resourceObjectWirelessControllerBonjourProfilePolicyListDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["bonjour_profile"] = bonjour_profile

	err = c.DeleteObjectWirelessControllerBonjourProfilePolicyList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerBonjourProfilePolicyListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectWirelessControllerBonjourProfilePolicyList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfilePolicyList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerBonjourProfilePolicyList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerBonjourProfilePolicyList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerBonjourProfilePolicyListDescription2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyListFromVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyListPolicyId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerBonjourProfilePolicyListServices2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerBonjourProfilePolicyListToVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("description", flattenObjectWirelessControllerBonjourProfilePolicyListDescription2edl(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectWirelessControllerBonjourProfilePolicyList-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("from_vlan", flattenObjectWirelessControllerBonjourProfilePolicyListFromVlan2edl(o["from-vlan"], d, "from_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["from-vlan"], "ObjectWirelessControllerBonjourProfilePolicyList-FromVlan"); ok {
			if err = d.Set("from_vlan", vv); err != nil {
				return fmt.Errorf("Error reading from_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from_vlan: %v", err)
		}
	}

	if err = d.Set("policy_id", flattenObjectWirelessControllerBonjourProfilePolicyListPolicyId2edl(o["policy-id"], d, "policy_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-id"], "ObjectWirelessControllerBonjourProfilePolicyList-PolicyId"); ok {
			if err = d.Set("policy_id", vv); err != nil {
				return fmt.Errorf("Error reading policy_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_id: %v", err)
		}
	}

	if err = d.Set("services", flattenObjectWirelessControllerBonjourProfilePolicyListServices2edl(o["services"], d, "services")); err != nil {
		if vv, ok := fortiAPIPatch(o["services"], "ObjectWirelessControllerBonjourProfilePolicyList-Services"); ok {
			if err = d.Set("services", vv); err != nil {
				return fmt.Errorf("Error reading services: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading services: %v", err)
		}
	}

	if err = d.Set("to_vlan", flattenObjectWirelessControllerBonjourProfilePolicyListToVlan2edl(o["to-vlan"], d, "to_vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["to-vlan"], "ObjectWirelessControllerBonjourProfilePolicyList-ToVlan"); ok {
			if err = d.Set("to_vlan", vv); err != nil {
				return fmt.Errorf("Error reading to_vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading to_vlan: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerBonjourProfilePolicyListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerBonjourProfilePolicyListDescription2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListFromVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListPolicyId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListServices2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerBonjourProfilePolicyListToVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListDescription2edl(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("from_vlan"); ok || d.HasChange("from_vlan") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListFromVlan2edl(d, v, "from_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from-vlan"] = t
		}
	}

	if v, ok := d.GetOk("policy_id"); ok || d.HasChange("policy_id") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListPolicyId2edl(d, v, "policy_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-id"] = t
		}
	}

	if v, ok := d.GetOk("services"); ok || d.HasChange("services") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListServices2edl(d, v, "services")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["services"] = t
		}
	}

	if v, ok := d.GetOk("to_vlan"); ok || d.HasChange("to_vlan") {
		t, err := expandObjectWirelessControllerBonjourProfilePolicyListToVlan2edl(d, v, "to_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["to-vlan"] = t
		}
	}

	return &obj, nil
}
