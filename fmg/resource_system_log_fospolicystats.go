// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiOS policy statistics settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogFosPolicyStats() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogFosPolicyStatsUpdate,
		Read:   resourceSystemLogFosPolicyStatsRead,
		Update: resourceSystemLogFosPolicyStatsUpdate,
		Delete: resourceSystemLogFosPolicyStatsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"retention_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sampling_interval": &schema.Schema{
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

func resourceSystemLogFosPolicyStatsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogFosPolicyStats(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFosPolicyStats resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogFosPolicyStats(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFosPolicyStats resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogFosPolicyStats")

	return resourceSystemLogFosPolicyStatsRead(d, m)
}

func resourceSystemLogFosPolicyStatsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLogFosPolicyStats(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogFosPolicyStats resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogFosPolicyStatsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogFosPolicyStats(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFosPolicyStats resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogFosPolicyStats(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFosPolicyStats resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogFosPolicyStatsRetentionDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFosPolicyStatsSamplingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFosPolicyStatsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogFosPolicyStats(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("retention_days", flattenSystemLogFosPolicyStatsRetentionDays(o["retention-days"], d, "retention_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention-days"], "SystemLogFosPolicyStats-RetentionDays"); ok {
			if err = d.Set("retention_days", vv); err != nil {
				return fmt.Errorf("Error reading retention_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention_days: %v", err)
		}
	}

	if err = d.Set("sampling_interval", flattenSystemLogFosPolicyStatsSamplingInterval(o["sampling-interval"], d, "sampling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sampling-interval"], "SystemLogFosPolicyStats-SamplingInterval"); ok {
			if err = d.Set("sampling_interval", vv); err != nil {
				return fmt.Errorf("Error reading sampling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sampling_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLogFosPolicyStatsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLogFosPolicyStats-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemLogFosPolicyStatsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogFosPolicyStatsRetentionDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFosPolicyStatsSamplingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFosPolicyStatsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogFosPolicyStats(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("retention_days"); ok || d.HasChange("retention_days") {
		t, err := expandSystemLogFosPolicyStatsRetentionDays(d, v, "retention_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention-days"] = t
		}
	}

	if v, ok := d.GetOk("sampling_interval"); ok || d.HasChange("sampling_interval") {
		t, err := expandSystemLogFosPolicyStatsSamplingInterval(d, v, "sampling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sampling-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLogFosPolicyStatsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
