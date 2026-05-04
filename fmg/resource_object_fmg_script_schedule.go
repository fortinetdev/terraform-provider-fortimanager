// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFmg ScriptSchedule

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFmgScriptSchedule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFmgScriptScheduleCreate,
		Read:   resourceObjectFmgScriptScheduleRead,
		Update: resourceObjectFmgScriptScheduleUpdate,
		Delete: resourceObjectFmgScriptScheduleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
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
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"datetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"day_of_week": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"run_on_db": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFmgScriptScheduleCreate(d *schema.ResourceData, m interface{}) error {
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

	script := d.Get("script").(string)
	paradict["script"] = script

	obj, err := getObjectObjectFmgScriptSchedule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFmgScriptSchedule resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("device")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectFmgScriptSchedule(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectFmgScriptSchedule(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectFmgScriptSchedule resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectFmgScriptSchedule(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectFmgScriptSchedule resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "device"))

	return resourceObjectFmgScriptScheduleRead(d, m)
}

func resourceObjectFmgScriptScheduleUpdate(d *schema.ResourceData, m interface{}) error {
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

	script := d.Get("script").(string)
	paradict["script"] = script

	obj, err := getObjectObjectFmgScriptSchedule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgScriptSchedule resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFmgScriptSchedule(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFmgScriptSchedule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "device"))

	return resourceObjectFmgScriptScheduleRead(d, m)
}

func resourceObjectFmgScriptScheduleDelete(d *schema.ResourceData, m interface{}) error {
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

	script := d.Get("script").(string)
	paradict["script"] = script

	wsParams["adom"] = adomv

	err = c.DeleteObjectFmgScriptSchedule(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFmgScriptSchedule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFmgScriptScheduleRead(d *schema.ResourceData, m interface{}) error {
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

	script := d.Get("script").(string)
	if script == "" {
		script = importOptionChecking(m.(*FortiClient).Cfg, "script")
		if script == "" {
			return fmt.Errorf("Parameter script is missing")
		}
		if err = d.Set("script", script); err != nil {
			return fmt.Errorf("Error set params script: %v", err)
		}
	}
	paradict["script"] = script

	o, err := c.ReadObjectFmgScriptSchedule(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFmgScriptSchedule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFmgScriptSchedule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFmgScriptSchedule resource from API: %v", err)
	}
	return nil
}

func flattenObjectFmgScriptScheduleDatetime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleDayOfWeek2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleRunOnDb2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleTimestamp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFmgScriptScheduleUser2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFmgScriptSchedule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("datetime", flattenObjectFmgScriptScheduleDatetime2edl(o["datetime"], d, "datetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["datetime"], "ObjectFmgScriptSchedule-Datetime"); ok {
			if err = d.Set("datetime", vv); err != nil {
				return fmt.Errorf("Error reading datetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading datetime: %v", err)
		}
	}

	if err = d.Set("day_of_week", flattenObjectFmgScriptScheduleDayOfWeek2edl(o["day-of-week"], d, "day_of_week")); err != nil {
		if vv, ok := fortiAPIPatch(o["day-of-week"], "ObjectFmgScriptSchedule-DayOfWeek"); ok {
			if err = d.Set("day_of_week", vv); err != nil {
				return fmt.Errorf("Error reading day_of_week: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading day_of_week: %v", err)
		}
	}

	if err = d.Set("device", flattenObjectFmgScriptScheduleDevice2edl(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "ObjectFmgScriptSchedule-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("run_on_db", flattenObjectFmgScriptScheduleRunOnDb2edl(o["run-on-db"], d, "run_on_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["run-on-db"], "ObjectFmgScriptSchedule-RunOnDb"); ok {
			if err = d.Set("run_on_db", vv); err != nil {
				return fmt.Errorf("Error reading run_on_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading run_on_db: %v", err)
		}
	}

	if err = d.Set("timestamp", flattenObjectFmgScriptScheduleTimestamp2edl(o["timestamp"], d, "timestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["timestamp"], "ObjectFmgScriptSchedule-Timestamp"); ok {
			if err = d.Set("timestamp", vv); err != nil {
				return fmt.Errorf("Error reading timestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timestamp: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFmgScriptScheduleType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFmgScriptSchedule-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("user", flattenObjectFmgScriptScheduleUser2edl(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "ObjectFmgScriptSchedule-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	return nil
}

func flattenObjectFmgScriptScheduleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFmgScriptScheduleDatetime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleDayOfWeek2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleRunOnDb2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleTimestamp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFmgScriptScheduleUser2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFmgScriptSchedule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("datetime"); ok || d.HasChange("datetime") {
		t, err := expandObjectFmgScriptScheduleDatetime2edl(d, v, "datetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["datetime"] = t
		}
	}

	if v, ok := d.GetOk("day_of_week"); ok || d.HasChange("day_of_week") {
		t, err := expandObjectFmgScriptScheduleDayOfWeek2edl(d, v, "day_of_week")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day-of-week"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandObjectFmgScriptScheduleDevice2edl(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("run_on_db"); ok || d.HasChange("run_on_db") {
		t, err := expandObjectFmgScriptScheduleRunOnDb2edl(d, v, "run_on_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["run-on-db"] = t
		}
	}

	if v, ok := d.GetOk("timestamp"); ok || d.HasChange("timestamp") {
		t, err := expandObjectFmgScriptScheduleTimestamp2edl(d, v, "timestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timestamp"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFmgScriptScheduleType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandObjectFmgScriptScheduleUser2edl(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
