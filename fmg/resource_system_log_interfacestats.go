// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Interface statistics settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogInterfaceStats() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogInterfaceStatsUpdate,
		Read:   resourceSystemLogInterfaceStatsRead,
		Update: resourceSystemLogInterfaceStatsUpdate,
		Delete: resourceSystemLogInterfaceStatsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"billing_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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

func resourceSystemLogInterfaceStatsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogInterfaceStats(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogInterfaceStats resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogInterfaceStats(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogInterfaceStats resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogInterfaceStats")

	return resourceSystemLogInterfaceStatsRead(d, m)
}

func resourceSystemLogInterfaceStatsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLogInterfaceStats(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogInterfaceStats resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogInterfaceStatsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogInterfaceStats(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogInterfaceStats resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogInterfaceStats(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogInterfaceStats resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogInterfaceStatsBillingReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogInterfaceStatsRetentionDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogInterfaceStatsSamplingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogInterfaceStatsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogInterfaceStats(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("billing_report", flattenSystemLogInterfaceStatsBillingReport(o["billing-report"], d, "billing_report")); err != nil {
		if vv, ok := fortiAPIPatch(o["billing-report"], "SystemLogInterfaceStats-BillingReport"); ok {
			if err = d.Set("billing_report", vv); err != nil {
				return fmt.Errorf("Error reading billing_report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading billing_report: %v", err)
		}
	}

	if err = d.Set("retention_days", flattenSystemLogInterfaceStatsRetentionDays(o["retention-days"], d, "retention_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["retention-days"], "SystemLogInterfaceStats-RetentionDays"); ok {
			if err = d.Set("retention_days", vv); err != nil {
				return fmt.Errorf("Error reading retention_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retention_days: %v", err)
		}
	}

	if err = d.Set("sampling_interval", flattenSystemLogInterfaceStatsSamplingInterval(o["sampling-interval"], d, "sampling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["sampling-interval"], "SystemLogInterfaceStats-SamplingInterval"); ok {
			if err = d.Set("sampling_interval", vv); err != nil {
				return fmt.Errorf("Error reading sampling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sampling_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLogInterfaceStatsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLogInterfaceStats-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemLogInterfaceStatsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogInterfaceStatsBillingReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogInterfaceStatsRetentionDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogInterfaceStatsSamplingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogInterfaceStatsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogInterfaceStats(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("billing_report"); ok || d.HasChange("billing_report") {
		t, err := expandSystemLogInterfaceStatsBillingReport(d, v, "billing_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-report"] = t
		}
	}

	if v, ok := d.GetOk("retention_days"); ok || d.HasChange("retention_days") {
		t, err := expandSystemLogInterfaceStatsRetentionDays(d, v, "retention_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retention-days"] = t
		}
	}

	if v, ok := d.GetOk("sampling_interval"); ok || d.HasChange("sampling_interval") {
		t, err := expandSystemLogInterfaceStatsSamplingInterval(d, v, "sampling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sampling-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLogInterfaceStatsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
