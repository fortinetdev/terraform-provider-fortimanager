// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure switch object tags.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerSwitchInterfaceTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerSwitchInterfaceTagCreate,
		Read:   resourceObjectSwitchControllerSwitchInterfaceTagRead,
		Update: resourceObjectSwitchControllerSwitchInterfaceTagUpdate,
		Delete: resourceObjectSwitchControllerSwitchInterfaceTagDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerSwitchInterfaceTagCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerSwitchInterfaceTag(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerSwitchInterfaceTag resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSwitchControllerSwitchInterfaceTag(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSwitchInterfaceTagRead(d, m)
}

func resourceObjectSwitchControllerSwitchInterfaceTagUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSwitchControllerSwitchInterfaceTag(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSwitchInterfaceTag resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSwitchControllerSwitchInterfaceTag(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSwitchInterfaceTagRead(d, m)
}

func resourceObjectSwitchControllerSwitchInterfaceTagDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSwitchControllerSwitchInterfaceTag(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerSwitchInterfaceTagRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSwitchControllerSwitchInterfaceTag(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerSwitchInterfaceTag resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerSwitchInterfaceTag(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerSwitchInterfaceTag resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerSwitchInterfaceTagName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerSwitchInterfaceTag(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("name", flattenObjectSwitchControllerSwitchInterfaceTagName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerSwitchInterfaceTag-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerSwitchInterfaceTagFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerSwitchInterfaceTagName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerSwitchInterfaceTag(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerSwitchInterfaceTagName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
