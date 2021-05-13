// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Send virus detection notification to FortiGuard.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateAnalyzerVirusreport() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateAnalyzerVirusreportUpdate,
		Read:   resourceFmupdateAnalyzerVirusreportRead,
		Update: resourceFmupdateAnalyzerVirusreportUpdate,
		Delete: resourceFmupdateAnalyzerVirusreportDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateAnalyzerVirusreportUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateAnalyzerVirusreport(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAnalyzerVirusreport resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateAnalyzerVirusreport(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAnalyzerVirusreport resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateAnalyzerVirusreport")

	return resourceFmupdateAnalyzerVirusreportRead(d, m)
}

func resourceFmupdateAnalyzerVirusreportDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateAnalyzerVirusreport(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateAnalyzerVirusreport resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateAnalyzerVirusreportRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateAnalyzerVirusreport(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateAnalyzerVirusreport resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateAnalyzerVirusreport(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateAnalyzerVirusreport resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateAnalyzerVirusreportStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectFmupdateAnalyzerVirusreport(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenFmupdateAnalyzerVirusreportStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateAnalyzerVirusreport-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdateAnalyzerVirusreportFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateAnalyzerVirusreportStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateAnalyzerVirusreport(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFmupdateAnalyzerVirusreportStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
