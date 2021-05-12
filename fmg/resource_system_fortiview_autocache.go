// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiView auto-cache settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemFortiviewAutoCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortiviewAutoCacheUpdate,
		Read:   resourceSystemFortiviewAutoCacheRead,
		Update: resourceSystemFortiviewAutoCacheUpdate,
		Delete: resourceSystemFortiviewAutoCacheDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aggressive_fortiview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFortiviewAutoCacheUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemFortiviewAutoCache(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiviewAutoCache resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFortiviewAutoCache(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortiviewAutoCache resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemFortiviewAutoCache")

	return resourceSystemFortiviewAutoCacheRead(d, m)
}

func resourceSystemFortiviewAutoCacheDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemFortiviewAutoCache(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFortiviewAutoCache resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortiviewAutoCacheRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemFortiviewAutoCache(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiviewAutoCache resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortiviewAutoCache(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortiviewAutoCache resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortiviewAutoCacheAggressiveFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemFortiviewAutoCacheInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortiviewAutoCacheStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemFortiviewAutoCache(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aggressive_fortiview", flattenSystemFortiviewAutoCacheAggressiveFortiview(o["aggressive-fortiview"], d, "aggressive_fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggressive-fortiview"], "SystemFortiviewAutoCache-AggressiveFortiview"); ok {
			if err = d.Set("aggressive_fortiview", vv); err != nil {
				return fmt.Errorf("Error reading aggressive_fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggressive_fortiview: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemFortiviewAutoCacheInterval(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "SystemFortiviewAutoCache-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFortiviewAutoCacheStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemFortiviewAutoCache-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemFortiviewAutoCacheFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFortiviewAutoCacheAggressiveFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiviewAutoCacheInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortiviewAutoCacheStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortiviewAutoCache(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggressive_fortiview"); ok {
		t, err := expandSystemFortiviewAutoCacheAggressiveFortiview(d, v, "aggressive_fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggressive-fortiview"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		t, err := expandSystemFortiviewAutoCacheInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemFortiviewAutoCacheStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}