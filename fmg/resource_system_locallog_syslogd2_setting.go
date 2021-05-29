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

func resourceSystemLocallogSyslogd2Setting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogSyslogd2SettingUpdate,
		Read:   resourceSystemLocallogSyslogd2SettingRead,
		Update: resourceSystemLocallogSyslogd2SettingUpdate,
		Delete: resourceSystemLocallogSyslogd2SettingDelete,

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

func resourceSystemLocallogSyslogd2SettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogSyslogd2Setting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogd2Setting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogSyslogd2Setting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogd2Setting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogSyslogd2Setting")

	return resourceSystemLocallogSyslogd2SettingRead(d, m)
}

func resourceSystemLocallogSyslogd2SettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogSyslogd2Setting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogSyslogd2Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogSyslogd2SettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogSyslogd2Setting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogd2Setting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogSyslogd2Setting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogd2Setting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogSyslogd2SettingCsv(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenSystemLocallogSyslogd2SettingFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "kernel",
			8:   "user",
			12:  "ntp",
			13:  "audit",
			14:  "alert",
			15:  "clock",
			16:  "mail",
			24:  "daemon",
			32:  "auth",
			40:  "syslog",
			48:  "lpr",
			56:  "news",
			64:  "uucp",
			72:  "cron",
			80:  "authpriv",
			88:  "ftp",
			128: "local0",
			136: "local1",
			144: "local2",
			152: "local3",
			160: "local4",
			168: "local5",
			176: "local6",
			184: "local7",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd2SettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "emergency",
			1: "alert",
			2: "critical",
			3: "error",
			4: "warning",
			5: "notification",
			6: "information",
			7: "debug",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd2SettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenSystemLocallogSyslogd2SettingSyslogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogSyslogd2Setting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("csv", flattenSystemLocallogSyslogd2SettingCsv(o["csv"], d, "csv")); err != nil {
		if vv, ok := fortiAPIPatch(o["csv"], "SystemLocallogSyslogd2Setting-Csv"); ok {
			if err = d.Set("csv", vv); err != nil {
				return fmt.Errorf("Error reading csv: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csv: %v", err)
		}
	}

	if err = d.Set("facility", flattenSystemLocallogSyslogd2SettingFacility(o["facility"], d, "facility")); err != nil {
		if vv, ok := fortiAPIPatch(o["facility"], "SystemLocallogSyslogd2Setting-Facility"); ok {
			if err = d.Set("facility", vv); err != nil {
				return fmt.Errorf("Error reading facility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading facility: %v", err)
		}
	}

	if err = d.Set("severity", flattenSystemLocallogSyslogd2SettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SystemLocallogSyslogd2Setting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLocallogSyslogd2SettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogSyslogd2Setting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("syslog_name", flattenSystemLocallogSyslogd2SettingSyslogName(o["syslog-name"], d, "syslog_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["syslog-name"], "SystemLocallogSyslogd2Setting-SyslogName"); ok {
			if err = d.Set("syslog_name", vv); err != nil {
				return fmt.Errorf("Error reading syslog_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading syslog_name: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogSyslogd2SettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogSyslogd2SettingCsv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingSyslogName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogSyslogd2Setting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("csv"); ok {
		t, err := expandSystemLocallogSyslogd2SettingCsv(d, v, "csv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csv"] = t
		}
	}

	if v, ok := d.GetOk("facility"); ok {
		t, err := expandSystemLocallogSyslogd2SettingFacility(d, v, "facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["facility"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandSystemLocallogSyslogd2SettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemLocallogSyslogd2SettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("syslog_name"); ok {
		t, err := expandSystemLocallogSyslogd2SettingSyslogName(d, v, "syslog_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-name"] = t
		}
	}

	return &obj, nil
}
