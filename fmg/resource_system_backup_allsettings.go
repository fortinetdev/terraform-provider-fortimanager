// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Scheduled backup settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemBackupAllSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemBackupAllSettingsUpdate,
		Read:   resourceSystemBackupAllSettingsRead,
		Update: resourceSystemBackupAllSettingsUpdate,
		Delete: resourceSystemBackupAllSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"crptpasswd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"week_days": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemBackupAllSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemBackupAllSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemBackupAllSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemBackupAllSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemBackupAllSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemBackupAllSettings")

	return resourceSystemBackupAllSettingsRead(d, m)
}

func resourceSystemBackupAllSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemBackupAllSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemBackupAllSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemBackupAllSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemBackupAllSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemBackupAllSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemBackupAllSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemBackupAllSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemBackupAllSettingsCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemBackupAllSettingsWeekDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemBackupAllSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cert", flattenSystemBackupAllSettingsCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "SystemBackupAllSettings-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("directory", flattenSystemBackupAllSettingsDirectory(o["directory"], d, "directory")); err != nil {
		if vv, ok := fortiAPIPatch(o["directory"], "SystemBackupAllSettings-Directory"); ok {
			if err = d.Set("directory", vv); err != nil {
				return fmt.Errorf("Error reading directory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading directory: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemBackupAllSettingsProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemBackupAllSettings-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemBackupAllSettingsServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemBackupAllSettings-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemBackupAllSettingsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemBackupAllSettings-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemBackupAllSettingsTime(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "SystemBackupAllSettings-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("user", flattenSystemBackupAllSettingsUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "SystemBackupAllSettings-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("week_days", flattenSystemBackupAllSettingsWeekDays(o["week_days"], d, "week_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["week_days"], "SystemBackupAllSettings-WeekDays"); ok {
			if err = d.Set("week_days", vv); err != nil {
				return fmt.Errorf("Error reading week_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading week_days: %v", err)
		}
	}

	return nil
}

func flattenSystemBackupAllSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemBackupAllSettingsCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsCrptpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemBackupAllSettingsDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemBackupAllSettingsProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemBackupAllSettingsWeekDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemBackupAllSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandSystemBackupAllSettingsCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("crptpasswd"); ok || d.HasChange("crptpasswd") {
		t, err := expandSystemBackupAllSettingsCrptpasswd(d, v, "crptpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crptpasswd"] = t
		}
	}

	if v, ok := d.GetOk("directory"); ok || d.HasChange("directory") {
		t, err := expandSystemBackupAllSettingsDirectory(d, v, "directory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["directory"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok || d.HasChange("passwd") {
		t, err := expandSystemBackupAllSettingsPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemBackupAllSettingsProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemBackupAllSettingsServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemBackupAllSettingsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandSystemBackupAllSettingsTime(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandSystemBackupAllSettingsUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("week_days"); ok || d.HasChange("week_days") {
		t, err := expandSystemBackupAllSettingsWeekDays(d, v, "week_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["week_days"] = t
		}
	}

	return &obj, nil
}
