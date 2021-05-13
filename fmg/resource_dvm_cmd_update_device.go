// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Refresh the FGFM connection and system information of a device.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmCmdUpdateDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdUpdateDeviceUpdate,
		Read:   resourceDvmCmdUpdateDeviceRead,
		Update: resourceDvmCmdUpdateDeviceUpdate,
		Delete: resourceDvmCmdUpdateDeviceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDvmCmdUpdateDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectDvmCmdUpdateDevice(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdUpdateDevice resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmCmdUpdateDevice(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdUpdateDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdUpdateDevice")

	return resourceDvmCmdUpdateDeviceRead(d, m)
}

func resourceDvmCmdUpdateDeviceDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdUpdateDeviceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdUpdateDeviceAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdUpdateDeviceDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdUpdateDeviceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "create_task",
			2: "nonblocking",
			4: "log_dev",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func refreshObjectDvmCmdUpdateDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenDvmCmdUpdateDeviceAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmCmdUpdateDevice-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("device", flattenDvmCmdUpdateDeviceDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "DvmCmdUpdateDevice-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("flags", flattenDvmCmdUpdateDeviceFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmCmdUpdateDevice-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	return nil
}

func flattenDvmCmdUpdateDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdUpdateDeviceAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdUpdateDeviceDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdUpdateDeviceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDvmCmdUpdateDevice(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok {
		t, err := expandDvmCmdUpdateDeviceAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandDvmCmdUpdateDeviceDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok {
		t, err := expandDvmCmdUpdateDeviceFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	return &obj, nil
}
