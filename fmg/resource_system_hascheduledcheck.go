// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Scheduled HA integrity check.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaScheduledCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaScheduledCheckUpdate,
		Read:   resourceSystemHaScheduledCheckRead,
		Update: resourceSystemHaScheduledCheckUpdate,
		Delete: resourceSystemHaScheduledCheckDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"week_days": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaScheduledCheckUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemHaScheduledCheck(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaScheduledCheck resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaScheduledCheck(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaScheduledCheck resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemHaScheduledCheck")

	return resourceSystemHaScheduledCheckRead(d, m)
}

func resourceSystemHaScheduledCheckDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemHaScheduledCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaScheduledCheck resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaScheduledCheckRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemHaScheduledCheck(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaScheduledCheck resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaScheduledCheck(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaScheduledCheck resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaScheduledCheckStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaScheduledCheckTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaScheduledCheckWeekDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemHaScheduledCheck(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemHaScheduledCheckStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemHaScheduledCheck-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemHaScheduledCheckTime(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "SystemHaScheduledCheck-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("week_days", flattenSystemHaScheduledCheckWeekDays(o["week_days"], d, "week_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["week_days"], "SystemHaScheduledCheck-WeekDays"); ok {
			if err = d.Set("week_days", vv); err != nil {
				return fmt.Errorf("Error reading week_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading week_days: %v", err)
		}
	}

	return nil
}

func flattenSystemHaScheduledCheckFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaScheduledCheckStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaScheduledCheckTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaScheduledCheckWeekDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemHaScheduledCheck(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemHaScheduledCheckStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandSystemHaScheduledCheckTime(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("week_days"); ok || d.HasChange("week_days") {
		t, err := expandSystemHaScheduledCheckWeekDays(d, v, "week_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["week_days"] = t
		}
	}

	return &obj, nil
}
