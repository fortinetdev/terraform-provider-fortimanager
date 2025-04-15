// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerCustomCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerCustomCommandCreate,
		Read:   resourceObjectSwitchControllerCustomCommandRead,
		Update: resourceObjectSwitchControllerCustomCommandUpdate,
		Delete: resourceObjectSwitchControllerCustomCommandDelete,

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
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerCustomCommandCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerCustomCommand resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectSwitchControllerCustomCommand(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerCustomCommand resource: %v", err)
	}

	d.SetId(getStringKey(d, "command_name"))

	return resourceObjectSwitchControllerCustomCommandRead(d, m)
}

func resourceObjectSwitchControllerCustomCommandUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectSwitchControllerCustomCommand(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerCustomCommand resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerCustomCommand(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerCustomCommand resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "command_name"))

	return resourceObjectSwitchControllerCustomCommandRead(d, m)
}

func resourceObjectSwitchControllerCustomCommandDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectSwitchControllerCustomCommand(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerCustomCommand resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerCustomCommandRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectSwitchControllerCustomCommand(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerCustomCommand resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerCustomCommand(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerCustomCommand resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerCustomCommandCommand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSwitchControllerCustomCommandDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerCustomCommand(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("command", flattenObjectSwitchControllerCustomCommandCommand(o["command"], d, "command")); err != nil {
		if vv, ok := fortiAPIPatch(o["command"], "ObjectSwitchControllerCustomCommand-Command"); ok {
			if err = d.Set("command", vv); err != nil {
				return fmt.Errorf("Error reading command: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command: %v", err)
		}
	}

	if err = d.Set("command_name", flattenObjectSwitchControllerCustomCommandCommandName(o["command-name"], d, "command_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["command-name"], "ObjectSwitchControllerCustomCommand-CommandName"); ok {
			if err = d.Set("command_name", vv); err != nil {
				return fmt.Errorf("Error reading command_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading command_name: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectSwitchControllerCustomCommandDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectSwitchControllerCustomCommand-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerCustomCommandFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerCustomCommandCommand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSwitchControllerCustomCommandDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerCustomCommand(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("command"); ok || d.HasChange("command") {
		t, err := expandObjectSwitchControllerCustomCommandCommand(d, v, "command")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command"] = t
		}
	}

	if v, ok := d.GetOk("command_name"); ok || d.HasChange("command_name") {
		t, err := expandObjectSwitchControllerCustomCommandCommandName(d, v, "command_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectSwitchControllerCustomCommandDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	return &obj, nil
}
