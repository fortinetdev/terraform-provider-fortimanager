// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit FortiSwitch commands to be pushed to this FortiSwitch device upon rebooting the FortiGate switch controller or the FortiSwitch.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerManagedSwitchCustomCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerManagedSwitchCustomCommandCreate,
		Read:   resourceObjectSwitchControllerManagedSwitchCustomCommandRead,
		Update: resourceObjectSwitchControllerManagedSwitchCustomCommandUpdate,
		Delete: resourceObjectSwitchControllerManagedSwitchCustomCommandDelete,

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
			"command_entry": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"command_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerManagedSwitchCustomCommandCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerManagedSwitchCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchCustomCommand resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerManagedSwitchCustomCommand(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	d.SetId(getStringKey(d, "command_entry"))

	return resourceObjectSwitchControllerManagedSwitchCustomCommandRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchCustomCommandUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerManagedSwitchCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchCustomCommand resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerManagedSwitchCustomCommand(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "command_entry"))

	return resourceObjectSwitchControllerManagedSwitchCustomCommandRead(d, m)
}

func resourceObjectSwitchControllerManagedSwitchCustomCommandDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerManagedSwitchCustomCommand(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerManagedSwitchCustomCommandRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerManagedSwitchCustomCommand(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchCustomCommand resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerManagedSwitchCustomCommand(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerManagedSwitchCustomCommand resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerManagedSwitchCustomCommandCommandName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("command_entry", flattenObjectSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(o["command-entry"], d, "command_entry")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-entry"], "ObjectSwitchControllerManagedSwitchCustomCommand-CommandEntry"); ok {
			if err = d.Set("command_entry", vv); err != nil {
				return fmt.Errorf("Error reading command_entry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_entry: %v", err)
		}
	}

	if err = d.Set("command_name", flattenObjectSwitchControllerManagedSwitchCustomCommandCommandName2edl(o["command-name"], d, "command_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-name"], "ObjectSwitchControllerManagedSwitchCustomCommand-CommandName"); ok {
			if err = d.Set("command_name", vv); err != nil {
				return fmt.Errorf("Error reading command_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerManagedSwitchCustomCommandFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerManagedSwitchCustomCommandCommandName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("command_entry"); ok || d.HasChange("command_entry") {
		t, err := expandObjectSwitchControllerManagedSwitchCustomCommandCommandEntry2edl(d, v, "command_entry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-entry"] = t
		}
	}

	if v, ok := d.GetOk("command_name"); ok || d.HasChange("command_name") {
		t, err := expandObjectSwitchControllerManagedSwitchCustomCommandCommandName2edl(d, v, "command_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-name"] = t
		}
	}

	return &obj, nil
}
