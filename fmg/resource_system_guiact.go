// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: System settings through GUI.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGuiact() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGuiactUpdate,
		Read:   resourceSystemGuiactRead,
		Update: resourceSystemGuiactUpdate,
		Delete: resourceSystemGuiactDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"backup_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_conf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eventlog_msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eventlog_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reboot": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reset2default": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"restore_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restore_conf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemGuiactUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemGuiact(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGuiact resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemGuiact(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemGuiact resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemGuiact")

	return resourceSystemGuiactRead(d, m)
}

func resourceSystemGuiactDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemGuiact(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGuiact resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGuiactRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemGuiact(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemGuiact resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGuiact(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGuiact resource from API: %v", err)
	}
	return nil
}

func flattenSystemGuiactBackupAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactBackupConf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactEventlogMsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactEventlogPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactReboot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactReset2Default(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactRestoreAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactRestoreConf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGuiactTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGuiact(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("backup_all", flattenSystemGuiactBackupAll(o["backup_all"], d, "backup_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["backup_all"], "SystemGuiact-BackupAll"); ok {
			if err = d.Set("backup_all", vv); err != nil {
				return fmt.Errorf("Error reading backup_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backup_all: %v", err)
		}
	}

	if err = d.Set("backup_conf", flattenSystemGuiactBackupConf(o["backup_conf"], d, "backup_conf")); err != nil {
		if vv, ok := fortiAPIPatch(o["backup_conf"], "SystemGuiact-BackupConf"); ok {
			if err = d.Set("backup_conf", vv); err != nil {
				return fmt.Errorf("Error reading backup_conf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading backup_conf: %v", err)
		}
	}

	if err = d.Set("eventlog_msg", flattenSystemGuiactEventlogMsg(o["eventlog_msg"], d, "eventlog_msg")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventlog_msg"], "SystemGuiact-EventlogMsg"); ok {
			if err = d.Set("eventlog_msg", vv); err != nil {
				return fmt.Errorf("Error reading eventlog_msg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventlog_msg: %v", err)
		}
	}

	if err = d.Set("eventlog_path", flattenSystemGuiactEventlogPath(o["eventlog_path"], d, "eventlog_path")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventlog_path"], "SystemGuiact-EventlogPath"); ok {
			if err = d.Set("eventlog_path", vv); err != nil {
				return fmt.Errorf("Error reading eventlog_path: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventlog_path: %v", err)
		}
	}

	if err = d.Set("reboot", flattenSystemGuiactReboot(o["reboot"], d, "reboot")); err != nil {
		if vv, ok := fortiAPIPatch(o["reboot"], "SystemGuiact-Reboot"); ok {
			if err = d.Set("reboot", vv); err != nil {
				return fmt.Errorf("Error reading reboot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reboot: %v", err)
		}
	}

	if err = d.Set("reset2default", flattenSystemGuiactReset2Default(o["reset2default"], d, "reset2default")); err != nil {
		if vv, ok := fortiAPIPatch(o["reset2default"], "SystemGuiact-Reset2Default"); ok {
			if err = d.Set("reset2default", vv); err != nil {
				return fmt.Errorf("Error reading reset2default: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reset2default: %v", err)
		}
	}

	if err = d.Set("restore_all", flattenSystemGuiactRestoreAll(o["restore_all"], d, "restore_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["restore_all"], "SystemGuiact-RestoreAll"); ok {
			if err = d.Set("restore_all", vv); err != nil {
				return fmt.Errorf("Error reading restore_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restore_all: %v", err)
		}
	}

	if err = d.Set("restore_conf", flattenSystemGuiactRestoreConf(o["restore_conf"], d, "restore_conf")); err != nil {
		if vv, ok := fortiAPIPatch(o["restore_conf"], "SystemGuiact-RestoreConf"); ok {
			if err = d.Set("restore_conf", vv); err != nil {
				return fmt.Errorf("Error reading restore_conf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading restore_conf: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemGuiactTime(o["time"], d, "time")); err != nil {
		if vv, ok := fortiAPIPatch(o["time"], "SystemGuiact-Time"); ok {
			if err = d.Set("time", vv); err != nil {
				return fmt.Errorf("Error reading time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	return nil
}

func flattenSystemGuiactFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGuiactBackupAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactBackupConf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactEventlogMsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactEventlogPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactReboot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactReset2Default(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactRestoreAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactRestoreConf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGuiactTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGuiact(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("backup_all"); ok || d.HasChange("backup_all") {
		t, err := expandSystemGuiactBackupAll(d, v, "backup_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup_all"] = t
		}
	}

	if v, ok := d.GetOk("backup_conf"); ok || d.HasChange("backup_conf") {
		t, err := expandSystemGuiactBackupConf(d, v, "backup_conf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup_conf"] = t
		}
	}

	if v, ok := d.GetOk("eventlog_msg"); ok || d.HasChange("eventlog_msg") {
		t, err := expandSystemGuiactEventlogMsg(d, v, "eventlog_msg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventlog_msg"] = t
		}
	}

	if v, ok := d.GetOk("eventlog_path"); ok || d.HasChange("eventlog_path") {
		t, err := expandSystemGuiactEventlogPath(d, v, "eventlog_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventlog_path"] = t
		}
	}

	if v, ok := d.GetOk("reboot"); ok || d.HasChange("reboot") {
		t, err := expandSystemGuiactReboot(d, v, "reboot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reboot"] = t
		}
	}

	if v, ok := d.GetOk("reset2default"); ok || d.HasChange("reset2default") {
		t, err := expandSystemGuiactReset2Default(d, v, "reset2default")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reset2default"] = t
		}
	}

	if v, ok := d.GetOk("restore_all"); ok || d.HasChange("restore_all") {
		t, err := expandSystemGuiactRestoreAll(d, v, "restore_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restore_all"] = t
		}
	}

	if v, ok := d.GetOk("restore_conf"); ok || d.HasChange("restore_conf") {
		t, err := expandSystemGuiactRestoreConf(d, v, "restore_conf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restore_conf"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok || d.HasChange("time") {
		t, err := expandSystemGuiactTime(d, v, "time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	return &obj, nil
}
