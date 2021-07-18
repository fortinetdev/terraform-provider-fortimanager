// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IoC settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogIoc() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogIocUpdate,
		Read:   resourceSystemLogIocRead,
		Update: resourceSystemLogIocUpdate,
		Delete: resourceSystemLogIocDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notification_throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rescan_max_runner": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rescan_run_at": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rescan_status": &schema.Schema{
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

func resourceSystemLogIocUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogIoc(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogIoc resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogIoc(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogIoc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogIoc")

	return resourceSystemLogIocRead(d, m)
}

func resourceSystemLogIocDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogIoc(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogIoc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogIocRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogIoc(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogIoc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogIoc(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogIoc resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogIocNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogIocNotificationThrottle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogIocRescanMaxRunner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogIocRescanRunAt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogIocRescanStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogIocStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogIoc(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("notification", flattenSystemLogIocNotification(o["notification"], d, "notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["notification"], "SystemLogIoc-Notification"); ok {
			if err = d.Set("notification", vv); err != nil {
				return fmt.Errorf("Error reading notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notification: %v", err)
		}
	}

	if err = d.Set("notification_throttle", flattenSystemLogIocNotificationThrottle(o["notification-throttle"], d, "notification_throttle")); err != nil {
		if vv, ok := fortiAPIPatch(o["notification-throttle"], "SystemLogIoc-NotificationThrottle"); ok {
			if err = d.Set("notification_throttle", vv); err != nil {
				return fmt.Errorf("Error reading notification_throttle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading notification_throttle: %v", err)
		}
	}

	if err = d.Set("rescan_max_runner", flattenSystemLogIocRescanMaxRunner(o["rescan-max-runner"], d, "rescan_max_runner")); err != nil {
		if vv, ok := fortiAPIPatch(o["rescan-max-runner"], "SystemLogIoc-RescanMaxRunner"); ok {
			if err = d.Set("rescan_max_runner", vv); err != nil {
				return fmt.Errorf("Error reading rescan_max_runner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rescan_max_runner: %v", err)
		}
	}

	if err = d.Set("rescan_run_at", flattenSystemLogIocRescanRunAt(o["rescan-run-at"], d, "rescan_run_at")); err != nil {
		if vv, ok := fortiAPIPatch(o["rescan-run-at"], "SystemLogIoc-RescanRunAt"); ok {
			if err = d.Set("rescan_run_at", vv); err != nil {
				return fmt.Errorf("Error reading rescan_run_at: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rescan_run_at: %v", err)
		}
	}

	if err = d.Set("rescan_status", flattenSystemLogIocRescanStatus(o["rescan-status"], d, "rescan_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["rescan-status"], "SystemLogIoc-RescanStatus"); ok {
			if err = d.Set("rescan_status", vv); err != nil {
				return fmt.Errorf("Error reading rescan_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rescan_status: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLogIocStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLogIoc-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemLogIocFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogIocNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogIocNotificationThrottle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogIocRescanMaxRunner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogIocRescanRunAt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogIocRescanStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogIocStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogIoc(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("notification"); ok {
		t, err := expandSystemLogIocNotification(d, v, "notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification"] = t
		}
	}

	if v, ok := d.GetOk("notification_throttle"); ok {
		t, err := expandSystemLogIocNotificationThrottle(d, v, "notification_throttle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notification-throttle"] = t
		}
	}

	if v, ok := d.GetOk("rescan_max_runner"); ok {
		t, err := expandSystemLogIocRescanMaxRunner(d, v, "rescan_max_runner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rescan-max-runner"] = t
		}
	}

	if v, ok := d.GetOk("rescan_run_at"); ok {
		t, err := expandSystemLogIocRescanRunAt(d, v, "rescan_run_at")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rescan-run-at"] = t
		}
	}

	if v, ok := d.GetOk("rescan_status"); ok {
		t, err := expandSystemLogIocRescanStatus(d, v, "rescan_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rescan-status"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemLogIocStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
