// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Monitored interfaces.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemHaMonitoredInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaMonitoredInterfacesCreate,
		Read:   resourceSystemHaMonitoredInterfacesRead,
		Update: resourceSystemHaMonitoredInterfacesUpdate,
		Delete: resourceSystemHaMonitoredInterfacesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"interface_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaMonitoredInterfacesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaMonitoredInterfaces(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaMonitoredInterfaces resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaMonitoredInterfaces(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemHaMonitoredInterfaces resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemHaMonitoredInterfacesRead(d, m)
}

func resourceSystemHaMonitoredInterfacesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemHaMonitoredInterfaces(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitoredInterfaces resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaMonitoredInterfaces(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaMonitoredInterfaces resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemHaMonitoredInterfacesRead(d, m)
}

func resourceSystemHaMonitoredInterfacesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemHaMonitoredInterfaces(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaMonitoredInterfaces resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaMonitoredInterfacesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemHaMonitoredInterfaces(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitoredInterfaces resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaMonitoredInterfaces(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaMonitoredInterfaces resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaMonitoredInterfacesInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaMonitoredInterfaces(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("interface_name", flattenSystemHaMonitoredInterfacesInterfaceName(o["interface-name"], d, "interface_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-name"], "SystemHaMonitoredInterfaces-InterfaceName"); ok {
			if err = d.Set("interface_name", vv); err != nil {
				return fmt.Errorf("Error reading interface_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_name: %v", err)
		}
	}

	return nil
}

func flattenSystemHaMonitoredInterfacesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaMonitoredInterfacesInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaMonitoredInterfaces(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interface_name"); ok || d.HasChange("interface_name") {
		t, err := expandSystemHaMonitoredInterfacesInterfaceName(d, v, "interface_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-name"] = t
		}
	}

	return &obj, nil
}
