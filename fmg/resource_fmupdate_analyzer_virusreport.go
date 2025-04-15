// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Send virus detection notification to FortiGuard.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateAnalyzerVirusreport(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAnalyzerVirusreport resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFmupdateAnalyzerVirusreport(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteFmupdateAnalyzerVirusreport(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateAnalyzerVirusreport(mkey, paradict)
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

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateAnalyzerVirusreportStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
