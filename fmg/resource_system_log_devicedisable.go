// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Disable client device logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogDeviceDisable() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogDeviceDisableCreate,
		Read:   resourceSystemLogDeviceDisableRead,
		Update: resourceSystemLogDeviceDisableUpdate,
		Delete: resourceSystemLogDeviceDisableDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemLogDeviceDisableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogDeviceDisable(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogDeviceDisable resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLogDeviceDisable(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemLogDeviceDisable resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogDeviceDisableRead(d, m)
}

func resourceSystemLogDeviceDisableUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogDeviceDisable(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogDeviceDisable resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogDeviceDisable(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogDeviceDisable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogDeviceDisableRead(d, m)
}

func resourceSystemLogDeviceDisableDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogDeviceDisable(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogDeviceDisable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogDeviceDisableRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogDeviceDisable(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogDeviceDisable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogDeviceDisable(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogDeviceDisable resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogDeviceDisableTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceDisableDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceDisableId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogDeviceDisable(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ttl", flattenSystemLogDeviceDisableTtl(o["TTL"], d, "ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["TTL"], "SystemLogDeviceDisable-Ttl"); ok {
			if err = d.Set("ttl", vv); err != nil {
				return fmt.Errorf("Error reading ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemLogDeviceDisableDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemLogDeviceDisable-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogDeviceDisableId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogDeviceDisable-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenSystemLogDeviceDisableFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogDeviceDisableTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceDisableDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceDisableId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogDeviceDisable(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ttl"); ok || d.HasChange("ttl") {
		t, err := expandSystemLogDeviceDisableTtl(d, v, "ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["TTL"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemLogDeviceDisableDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemLogDeviceDisableId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
