// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the schedule when built-in FortiGuard retrieves antivirus and IPS updates.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFdsSettingUpdateSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingUpdateScheduleUpdate,
		Read:   resourceFmupdateFdsSettingUpdateScheduleRead,
		Update: resourceFmupdateFdsSettingUpdateScheduleUpdate,
		Delete: resourceFmupdateFdsSettingUpdateScheduleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFdsSettingUpdateScheduleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFdsSettingUpdateSchedule(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingUpdateSchedule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFmupdateFdsSettingUpdateSchedule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingUpdateSchedule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFdsSettingUpdateSchedule")

	return resourceFmupdateFdsSettingUpdateScheduleRead(d, m)
}

func resourceFmupdateFdsSettingUpdateScheduleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteFmupdateFdsSettingUpdateSchedule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSettingUpdateSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingUpdateScheduleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFdsSettingUpdateSchedule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingUpdateSchedule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSettingUpdateSchedule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingUpdateSchedule resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingUpdateScheduleDay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleFrequency2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingUpdateScheduleTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFmupdateFdsSettingUpdateSchedule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("day", flattenFmupdateFdsSettingUpdateScheduleDay2edl(o["day"], d, "day")); err != nil {
		if vv, ok := fortiAPIPatch(o["day"], "FmupdateFdsSettingUpdateSchedule-Day"); ok {
			if err = d.Set("day", vv); err != nil {
				return fmt.Errorf("Error reading day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("frequency", flattenFmupdateFdsSettingUpdateScheduleFrequency2edl(o["frequency"], d, "frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["frequency"], "FmupdateFdsSettingUpdateSchedule-Frequency"); ok {
			if err = d.Set("frequency", vv); err != nil {
				return fmt.Errorf("Error reading frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading frequency: %v", err)
		}
	}

	if err = d.Set("status", flattenFmupdateFdsSettingUpdateScheduleStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateFdsSettingUpdateSchedule-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("time", flattenFmupdateFdsSettingUpdateScheduleTime2edl(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "FmupdateFdsSettingUpdateSchedule-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingUpdateScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingUpdateScheduleDay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleFrequency2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingUpdateScheduleTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFmupdateFdsSettingUpdateSchedule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("day"); ok || d.HasChange("day") {
		t, err := expandFmupdateFdsSettingUpdateScheduleDay2edl(d, v, "day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("frequency"); ok || d.HasChange("frequency") {
		t, err := expandFmupdateFdsSettingUpdateScheduleFrequency2edl(d, v, "frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frequency"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateFdsSettingUpdateScheduleStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandFmupdateFdsSettingUpdateScheduleTime2edl(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	return &obj, nil
}
