// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VLAN assignment priority.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerManagedSwitchVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerManagedSwitchVlanCreate,
		Read:   resourceObjectSwitchControllerManagedSwitchVlanRead,
		Update: resourceObjectSwitchControllerManagedSwitchVlanUpdate,
		Delete: resourceObjectSwitchControllerManagedSwitchVlanDelete,

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
			"managed_switch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"assignment_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerManagedSwitchVlanCreate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchVlan(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchVlan resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerManagedSwitchVlan(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchVlan resource: %v", err)
	}

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceObjectSwitchControllerManagedSwitchVlanRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchVlanUpdate(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	obj, err := getObjectObjectSwitchControllerManagedSwitchVlan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchVlan resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerManagedSwitchVlan(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchVlan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "vlan_name"))

	return resourceObjectSwitchControllerManagedSwitchVlanRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchVlanDelete(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	paradict["managed_switch"] = managed_switch

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerManagedSwitchVlan(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerManagedSwitchVlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerManagedSwitchVlanRead(d *schema.ResourceData, m interface{}) error {
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

	managed_switch := d.Get("managed_switch").(string)
	if managed_switch == "" {
		managed_switch = importOptionChecking(m.(*FortiClient).Cfg, "managed_switch")
		if managed_switch == "" {
			return fmt.Errorf("Parameter managed_switch is missing")
		}
		if err = d.Set("managed_switch", managed_switch); err != nil {
			return fmt.Errorf("Error set params managed_switch: %v", err)
		}
	}
	paradict["managed_switch"] = managed_switch

	o, err := c.ReadObjectSwitchControllerManagedSwitchVlan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchVlan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerManagedSwitchVlan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchVlan resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerManagedSwitchVlanAssignmentPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchVlanVlanName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerManagedSwitchVlan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("assignment_priority", flattenObjectSwitchControllerManagedSwitchVlanAssignmentPriority2edl(o["assignment-priority"], d, "assignment_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment-priority"], "ObjectSwitchControllerManagedSwitchVlan-AssignmentPriority"); ok {
			if err = d.Set("assignment_priority", vv); err != nil {
				return fmt.Errorf("Error reading assignment_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment_priority: %v", err)
		}
	}

	if err = d.Set("vlan_name", flattenObjectSwitchControllerManagedSwitchVlanVlanName2edl(o["vlan-name"], d, "vlan_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan-name"], "ObjectSwitchControllerManagedSwitchVlan-VlanName"); ok {
			if err = d.Set("vlan_name", vv); err != nil {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan_name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerManagedSwitchVlanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerManagedSwitchVlanAssignmentPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchVlanVlanName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerManagedSwitchVlan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("assignment_priority"); ok || d.HasChange("assignment_priority") {
		t, err := expandObjectSwitchControllerManagedSwitchVlanAssignmentPriority2edl(d, v, "assignment_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment-priority"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandObjectSwitchControllerManagedSwitchVlanVlanName2edl(d, v, "vlan_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	return &obj, nil
}
