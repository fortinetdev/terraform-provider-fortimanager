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
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
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

func flattenSystemLocallogSyslogd2SettingCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingCsv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingFacility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2SettingSyslogName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogSyslogd2Setting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cert", flattenSystemLocallogSyslogd2SettingCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "SystemLocallogSyslogd2Setting-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

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

	if err = d.Set("reliable", flattenSystemLocallogSyslogd2SettingReliable(o["reliable"], d, "reliable")); err != nil {
		if vv, ok := fortiAPIPatch(o["reliable"], "SystemLocallogSyslogd2Setting-Reliable"); ok {
			if err = d.Set("reliable", vv); err != nil {
				return fmt.Errorf("Error reading reliable: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenSystemLocallogSyslogd2SettingSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "SystemLocallogSyslogd2Setting-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
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

func expandSystemLocallogSyslogd2SettingCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingCsv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingFacility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2SettingSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandSystemLocallogSyslogd2SettingCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("csv"); ok || d.HasChange("csv") {
		t, err := expandSystemLocallogSyslogd2SettingCsv(d, v, "csv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csv"] = t
		}
	}

	if v, ok := d.GetOk("facility"); ok || d.HasChange("facility") {
		t, err := expandSystemLocallogSyslogd2SettingFacility(d, v, "facility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["facility"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok || d.HasChange("reliable") {
		t, err := expandSystemLocallogSyslogd2SettingReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandSystemLocallogSyslogd2SettingSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandSystemLocallogSyslogd2SettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLocallogSyslogd2SettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("syslog_name"); ok || d.HasChange("syslog_name") {
		t, err := expandSystemLocallogSyslogd2SettingSyslogName(d, v, "syslog_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-name"] = t
		}
	}

	return &obj, nil
}
