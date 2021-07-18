// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Report estimated browse time settings

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemReportEstBrowseTime() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReportEstBrowseTimeUpdate,
		Read:   resourceSystemReportEstBrowseTimeRead,
		Update: resourceSystemReportEstBrowseTimeUpdate,
		Delete: resourceSystemReportEstBrowseTimeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_read_time": &schema.Schema{
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

func resourceSystemReportEstBrowseTimeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemReportEstBrowseTime(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportEstBrowseTime resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemReportEstBrowseTime(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemReportEstBrowseTime resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemReportEstBrowseTime")

	return resourceSystemReportEstBrowseTimeRead(d, m)
}

func resourceSystemReportEstBrowseTimeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemReportEstBrowseTime(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReportEstBrowseTime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReportEstBrowseTimeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemReportEstBrowseTime(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportEstBrowseTime resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReportEstBrowseTime(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReportEstBrowseTime resource from API: %v", err)
	}
	return nil
}

func flattenSystemReportEstBrowseTimeMaxReadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReportEstBrowseTimeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReportEstBrowseTime(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("max_read_time", flattenSystemReportEstBrowseTimeMaxReadTime(o["max-read-time"], d, "max_read_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-read-time"], "SystemReportEstBrowseTime-MaxReadTime"); ok {
			if err = d.Set("max_read_time", vv); err != nil {
				return fmt.Errorf("Error reading max_read_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_read_time: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemReportEstBrowseTimeStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemReportEstBrowseTime-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemReportEstBrowseTimeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReportEstBrowseTimeMaxReadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReportEstBrowseTimeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReportEstBrowseTime(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_read_time"); ok {
		t, err := expandSystemReportEstBrowseTimeMaxReadTime(d, v, "max_read_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-read-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemReportEstBrowseTimeStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
