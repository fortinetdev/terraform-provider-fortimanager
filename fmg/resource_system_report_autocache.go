// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report auto-cache settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemReportAutoCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReportAutoCacheUpdate,
		Read:   resourceSystemReportAutoCacheRead,
		Update: resourceSystemReportAutoCacheUpdate,
		Delete: resourceSystemReportAutoCacheDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aggressive_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemReportAutoCacheUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemReportAutoCache(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportAutoCache resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemReportAutoCache(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportAutoCache resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemReportAutoCache")

	return resourceSystemReportAutoCacheRead(d, m)
}

func resourceSystemReportAutoCacheDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemReportAutoCache(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReportAutoCache resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReportAutoCacheRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemReportAutoCache(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportAutoCache resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReportAutoCache(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportAutoCache resource from API: %v", err)
	}
	return nil
}

func flattenSystemReportAutoCacheAggressiveSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenSystemReportAutoCacheOrder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "oldest-first",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemReportAutoCacheStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectSystemReportAutoCache(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aggressive_schedule", flattenSystemReportAutoCacheAggressiveSchedule(o["aggressive-schedule"], d, "aggressive_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["aggressive-schedule"], "SystemReportAutoCache-AggressiveSchedule"); ok {
			if err = d.Set("aggressive_schedule", vv); err != nil {
				return fmt.Errorf("Error reading aggressive_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aggressive_schedule: %v", err)
		}
	}

	if err = d.Set("order", flattenSystemReportAutoCacheOrder(o["order"], d, "order")); err != nil {
		if vv, ok := fortiAPIPatch(o["order"], "SystemReportAutoCache-Order"); ok {
			if err = d.Set("order", vv); err != nil {
				return fmt.Errorf("Error reading order: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading order: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemReportAutoCacheStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemReportAutoCache-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemReportAutoCacheFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReportAutoCacheAggressiveSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportAutoCacheOrder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportAutoCacheStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReportAutoCache(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggressive_schedule"); ok {
		t, err := expandSystemReportAutoCacheAggressiveSchedule(d, v, "aggressive_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggressive-schedule"] = t
		}
	}

	if v, ok := d.GetOk("order"); ok {
		t, err := expandSystemReportAutoCacheOrder(d, v, "order")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["order"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemReportAutoCacheStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
