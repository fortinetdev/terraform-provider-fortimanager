// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for remote syslog server.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogSyslogdSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogSyslogdSettingUpdate,
		Read:   resourceSystemLocallogSyslogdSettingRead,
		Update: resourceSystemLocallogSyslogdSettingUpdate,
		Delete: resourceSystemLocallogSyslogdSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"csv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syslog_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogSyslogdSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogSyslogdSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogdSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogSyslogdSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogdSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogSyslogdSetting")

	return resourceSystemLocallogSyslogdSettingRead(d, m)
}

func resourceSystemLocallogSyslogdSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogSyslogdSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogSyslogdSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogSyslogdSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogSyslogdSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogdSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogSyslogdSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogdSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogSyslogdSettingCsv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogdSettingFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogdSettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogdSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogdSettingSyslogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogSyslogdSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("csv", flattenSystemLocallogSyslogdSettingCsv(o["csv"], d, "csv")); err != nil {
		if vv, ok := fortiAPIPatch(o["csv"], "SystemLocallogSyslogdSetting-Csv"); ok {
			if err = d.Set("csv", vv); err != nil {
				return fmt.Errorf("Error reading csv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csv: %v", err)
		}
	}

	if err = d.Set("facility", flattenSystemLocallogSyslogdSettingFacility(o["facility"], d, "facility")); err != nil {
		if vv, ok := fortiAPIPatch(o["facility"], "SystemLocallogSyslogdSetting-Facility"); ok {
			if err = d.Set("facility", vv); err != nil {
				return fmt.Errorf("Error reading facility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading facility: %v", err)
		}
	}

	if err = d.Set("severity", flattenSystemLocallogSyslogdSettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SystemLocallogSyslogdSetting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLocallogSyslogdSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogSyslogdSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("syslog_name", flattenSystemLocallogSyslogdSettingSyslogName(o["syslog-name"], d, "syslog_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-name"], "SystemLocallogSyslogdSetting-SyslogName"); ok {
			if err = d.Set("syslog_name", vv); err != nil {
				return fmt.Errorf("Error reading syslog_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_name: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogSyslogdSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogSyslogdSettingCsv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogdSettingFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogdSettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogdSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogdSettingSyslogName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogSyslogdSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("csv"); ok {
		t, err := expandSystemLocallogSyslogdSettingCsv(d, v, "csv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csv"] = t
		}
	}

	if v, ok := d.GetOk("facility"); ok {
		t, err := expandSystemLocallogSyslogdSettingFacility(d, v, "facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["facility"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandSystemLocallogSyslogdSettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemLocallogSyslogdSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("syslog_name"); ok {
		t, err := expandSystemLocallogSyslogdSettingSyslogName(d, v, "syslog_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-name"] = t
		}
	}

	return &obj, nil
}
