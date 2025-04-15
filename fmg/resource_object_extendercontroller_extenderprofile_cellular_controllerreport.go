// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender controller report configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfileCellularControllerReport() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileCellularControllerReportUpdate,
		Read:   resourceObjectExtenderControllerExtenderProfileCellularControllerReportRead,
		Update: resourceObjectExtenderControllerExtenderProfileCellularControllerReportUpdate,
		Delete: resourceObjectExtenderControllerExtenderProfileCellularControllerReportDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signal_threshold": &schema.Schema{
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

func resourceObjectExtenderControllerExtenderProfileCellularControllerReportUpdate(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtenderControllerExtenderProfileCellularControllerReport(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularControllerReport resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtenderControllerExtenderProfileCellularControllerReport(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularControllerReport resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtenderControllerExtenderProfileCellularControllerReport")

	return resourceObjectExtenderControllerExtenderProfileCellularControllerReportRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileCellularControllerReportDelete(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectExtenderControllerExtenderProfileCellularControllerReport(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfileCellularControllerReport resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileCellularControllerReportRead(d *schema.ResourceData, m interface{}) error {
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

	extender_profile := d.Get("extender_profile").(string)
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadObjectExtenderControllerExtenderProfileCellularControllerReport(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularControllerReport resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfileCellularControllerReport(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularControllerReport resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("interval", flattenObjectExtenderControllerExtenderProfileCellularControllerReportInterval3rdl(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "ObjectExtenderControllerExtenderProfileCellularControllerReport-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("signal_threshold", flattenObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(o["signal-threshold"], d, "signal_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["signal-threshold"], "ObjectExtenderControllerExtenderProfileCellularControllerReport-SignalThreshold"); ok {
			if err = d.Set("signal_threshold", vv); err != nil {
				return fmt.Errorf("Error reading signal_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signal_threshold: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtenderControllerExtenderProfileCellularControllerReportStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtenderControllerExtenderProfileCellularControllerReport-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularControllerReportFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReportInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularControllerReportStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerExtenderProfileCellularControllerReport(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularControllerReportInterval3rdl(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("signal_threshold"); ok || d.HasChange("signal_threshold") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularControllerReportSignalThreshold3rdl(d, v, "signal_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signal-threshold"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularControllerReportStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
