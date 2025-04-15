// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report auto-cache settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"sche_rpt_only": &schema.Schema{
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemReportAutoCache(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportAutoCache resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemReportAutoCache(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemReportAutoCache(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemReportAutoCache(mkey, paradict)
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
	return v
}

func flattenSystemReportAutoCacheOrder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportAutoCacheScheRptOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportAutoCacheStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

	if err = d.Set("sche_rpt_only", flattenSystemReportAutoCacheScheRptOnly(o["sche-rpt-only"], d, "sche_rpt_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["sche-rpt-only"], "SystemReportAutoCache-ScheRptOnly"); ok {
			if err = d.Set("sche_rpt_only", vv); err != nil {
				return fmt.Errorf("Error reading sche_rpt_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sche_rpt_only: %v", err)
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

func expandSystemReportAutoCacheScheRptOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportAutoCacheStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReportAutoCache(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggressive_schedule"); ok || d.HasChange("aggressive_schedule") {
		t, err := expandSystemReportAutoCacheAggressiveSchedule(d, v, "aggressive_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggressive-schedule"] = t
		}
	}

	if v, ok := d.GetOk("order"); ok || d.HasChange("order") {
		t, err := expandSystemReportAutoCacheOrder(d, v, "order")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["order"] = t
		}
	}

	if v, ok := d.GetOk("sche_rpt_only"); ok || d.HasChange("sche_rpt_only") {
		t, err := expandSystemReportAutoCacheScheRptOnly(d, v, "sche_rpt_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sche-rpt-only"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemReportAutoCacheStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
