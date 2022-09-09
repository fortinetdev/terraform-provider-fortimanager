// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for local disk logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogDiskSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogDiskSettingUpdate,
		Read:   resourceSystemLocallogDiskSettingRead,
		Update: resourceSystemLocallogDiskSettingUpdate,
		Delete: resourceSystemLocallogDiskSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_disk_full_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_disk_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_log_file_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_log_file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"roll_day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"roll_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roll_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_type": &schema.Schema{
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
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_delete_files": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uploaddir": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uploadip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadpass": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"uploadport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uploadsched": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadtype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"uploaduser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uploadzip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogDiskSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogDiskSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogDiskSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogDiskSetting(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogDiskSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogDiskSetting")

	return resourceSystemLocallogDiskSettingRead(d, m)
}

func resourceSystemLocallogDiskSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogDiskSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogDiskSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogDiskSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogDiskSetting(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogDiskSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogDiskSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogDiskSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogDiskSettingDiskfull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingLogDiskFullPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingLogDiskQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingMaxLogFileNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingMaxLogFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingRollDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLocallogDiskSettingRollSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingRollTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadDeleteFiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploaddir(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadpass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLocallogDiskSettingUploadport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadsched(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLocallogDiskSettingUploaduser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskSettingUploadzip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogDiskSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("diskfull", flattenSystemLocallogDiskSettingDiskfull(o["diskfull"], d, "diskfull")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskfull"], "SystemLocallogDiskSetting-Diskfull"); ok {
			if err = d.Set("diskfull", vv); err != nil {
				return fmt.Errorf("Error reading diskfull: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("log_disk_full_percentage", flattenSystemLocallogDiskSettingLogDiskFullPercentage(o["log-disk-full-percentage"], d, "log_disk_full_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-disk-full-percentage"], "SystemLocallogDiskSetting-LogDiskFullPercentage"); ok {
			if err = d.Set("log_disk_full_percentage", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_full_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_full_percentage: %v", err)
		}
	}

	if err = d.Set("log_disk_quota", flattenSystemLocallogDiskSettingLogDiskQuota(o["log-disk-quota"], d, "log_disk_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-disk-quota"], "SystemLocallogDiskSetting-LogDiskQuota"); ok {
			if err = d.Set("log_disk_quota", vv); err != nil {
				return fmt.Errorf("Error reading log_disk_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_disk_quota: %v", err)
		}
	}

	if err = d.Set("max_log_file_num", flattenSystemLocallogDiskSettingMaxLogFileNum(o["max-log-file-num"], d, "max_log_file_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-file-num"], "SystemLocallogDiskSetting-MaxLogFileNum"); ok {
			if err = d.Set("max_log_file_num", vv); err != nil {
				return fmt.Errorf("Error reading max_log_file_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_file_num: %v", err)
		}
	}

	if err = d.Set("max_log_file_size", flattenSystemLocallogDiskSettingMaxLogFileSize(o["max-log-file-size"], d, "max_log_file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-log-file-size"], "SystemLocallogDiskSetting-MaxLogFileSize"); ok {
			if err = d.Set("max_log_file_size", vv); err != nil {
				return fmt.Errorf("Error reading max_log_file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_log_file_size: %v", err)
		}
	}

	if err = d.Set("roll_day", flattenSystemLocallogDiskSettingRollDay(o["roll-day"], d, "roll_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["roll-day"], "SystemLocallogDiskSetting-RollDay"); ok {
			if err = d.Set("roll_day", vv); err != nil {
				return fmt.Errorf("Error reading roll_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roll_day: %v", err)
		}
	}

	if err = d.Set("roll_schedule", flattenSystemLocallogDiskSettingRollSchedule(o["roll-schedule"], d, "roll_schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["roll-schedule"], "SystemLocallogDiskSetting-RollSchedule"); ok {
			if err = d.Set("roll_schedule", vv); err != nil {
				return fmt.Errorf("Error reading roll_schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roll_schedule: %v", err)
		}
	}

	if err = d.Set("roll_time", flattenSystemLocallogDiskSettingRollTime(o["roll-time"], d, "roll_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["roll-time"], "SystemLocallogDiskSetting-RollTime"); ok {
			if err = d.Set("roll_time", vv); err != nil {
				return fmt.Errorf("Error reading roll_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading roll_time: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemLocallogDiskSettingServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemLocallogDiskSetting-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("severity", flattenSystemLocallogDiskSettingSeverity(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SystemLocallogDiskSetting-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLocallogDiskSettingStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLocallogDiskSetting-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload", flattenSystemLocallogDiskSettingUpload(o["upload"], d, "upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload"], "SystemLocallogDiskSetting-Upload"); ok {
			if err = d.Set("upload", vv); err != nil {
				return fmt.Errorf("Error reading upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_delete_files", flattenSystemLocallogDiskSettingUploadDeleteFiles(o["upload-delete-files"], d, "upload_delete_files")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-delete-files"], "SystemLocallogDiskSetting-UploadDeleteFiles"); ok {
			if err = d.Set("upload_delete_files", vv); err != nil {
				return fmt.Errorf("Error reading upload_delete_files: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_delete_files: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenSystemLocallogDiskSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-time"], "SystemLocallogDiskSetting-UploadTime"); ok {
			if err = d.Set("upload_time", vv); err != nil {
				return fmt.Errorf("Error reading upload_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("uploaddir", flattenSystemLocallogDiskSettingUploaddir(o["uploaddir"], d, "uploaddir")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploaddir"], "SystemLocallogDiskSetting-Uploaddir"); ok {
			if err = d.Set("uploaddir", vv); err != nil {
				return fmt.Errorf("Error reading uploaddir: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploaddir: %v", err)
		}
	}

	if err = d.Set("uploadip", flattenSystemLocallogDiskSettingUploadip(o["uploadip"], d, "uploadip")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadip"], "SystemLocallogDiskSetting-Uploadip"); ok {
			if err = d.Set("uploadip", vv); err != nil {
				return fmt.Errorf("Error reading uploadip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadip: %v", err)
		}
	}

	if err = d.Set("uploadpass", flattenSystemLocallogDiskSettingUploadpass(o["uploadpass"], d, "uploadpass")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadpass"], "SystemLocallogDiskSetting-Uploadpass"); ok {
			if err = d.Set("uploadpass", vv); err != nil {
				return fmt.Errorf("Error reading uploadpass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadpass: %v", err)
		}
	}

	if err = d.Set("uploadport", flattenSystemLocallogDiskSettingUploadport(o["uploadport"], d, "uploadport")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadport"], "SystemLocallogDiskSetting-Uploadport"); ok {
			if err = d.Set("uploadport", vv); err != nil {
				return fmt.Errorf("Error reading uploadport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadport: %v", err)
		}
	}

	if err = d.Set("uploadsched", flattenSystemLocallogDiskSettingUploadsched(o["uploadsched"], d, "uploadsched")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadsched"], "SystemLocallogDiskSetting-Uploadsched"); ok {
			if err = d.Set("uploadsched", vv); err != nil {
				return fmt.Errorf("Error reading uploadsched: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadsched: %v", err)
		}
	}

	if err = d.Set("uploadtype", flattenSystemLocallogDiskSettingUploadtype(o["uploadtype"], d, "uploadtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadtype"], "SystemLocallogDiskSetting-Uploadtype"); ok {
			if err = d.Set("uploadtype", vv); err != nil {
				return fmt.Errorf("Error reading uploadtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadtype: %v", err)
		}
	}

	if err = d.Set("uploaduser", flattenSystemLocallogDiskSettingUploaduser(o["uploaduser"], d, "uploaduser")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploaduser"], "SystemLocallogDiskSetting-Uploaduser"); ok {
			if err = d.Set("uploaduser", vv); err != nil {
				return fmt.Errorf("Error reading uploaduser: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploaduser: %v", err)
		}
	}

	if err = d.Set("uploadzip", flattenSystemLocallogDiskSettingUploadzip(o["uploadzip"], d, "uploadzip")); err != nil {
		if vv, ok := fortiAPIPatch(o["uploadzip"], "SystemLocallogDiskSetting-Uploadzip"); ok {
			if err = d.Set("uploadzip", vv); err != nil {
				return fmt.Errorf("Error reading uploadzip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uploadzip: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogDiskSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogDiskSettingDiskfull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingLogDiskFullPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingLogDiskQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingMaxLogFileNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingMaxLogFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingRollDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLocallogDiskSettingRollSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingRollTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadDeleteFiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploaddir(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadpass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLocallogDiskSettingUploadport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadsched(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLocallogDiskSettingUploaduser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskSettingUploadzip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogDiskSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("diskfull"); ok || d.HasChange("diskfull") {
		t, err := expandSystemLocallogDiskSettingDiskfull(d, v, "diskfull")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_full_percentage"); ok || d.HasChange("log_disk_full_percentage") {
		t, err := expandSystemLocallogDiskSettingLogDiskFullPercentage(d, v, "log_disk_full_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-full-percentage"] = t
		}
	}

	if v, ok := d.GetOk("log_disk_quota"); ok || d.HasChange("log_disk_quota") {
		t, err := expandSystemLocallogDiskSettingLogDiskQuota(d, v, "log_disk_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-disk-quota"] = t
		}
	}

	if v, ok := d.GetOk("max_log_file_num"); ok || d.HasChange("max_log_file_num") {
		t, err := expandSystemLocallogDiskSettingMaxLogFileNum(d, v, "max_log_file_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-file-num"] = t
		}
	}

	if v, ok := d.GetOk("max_log_file_size"); ok || d.HasChange("max_log_file_size") {
		t, err := expandSystemLocallogDiskSettingMaxLogFileSize(d, v, "max_log_file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-file-size"] = t
		}
	}

	if v, ok := d.GetOk("roll_day"); ok || d.HasChange("roll_day") {
		t, err := expandSystemLocallogDiskSettingRollDay(d, v, "roll_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-day"] = t
		}
	}

	if v, ok := d.GetOk("roll_schedule"); ok || d.HasChange("roll_schedule") {
		t, err := expandSystemLocallogDiskSettingRollSchedule(d, v, "roll_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-schedule"] = t
		}
	}

	if v, ok := d.GetOk("roll_time"); ok || d.HasChange("roll_time") {
		t, err := expandSystemLocallogDiskSettingRollTime(d, v, "roll_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-time"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemLocallogDiskSettingServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandSystemLocallogDiskSettingSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLocallogDiskSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload"); ok || d.HasChange("upload") {
		t, err := expandSystemLocallogDiskSettingUpload(d, v, "upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload"] = t
		}
	}

	if v, ok := d.GetOk("upload_delete_files"); ok || d.HasChange("upload_delete_files") {
		t, err := expandSystemLocallogDiskSettingUploadDeleteFiles(d, v, "upload_delete_files")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-delete-files"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok || d.HasChange("upload_time") {
		t, err := expandSystemLocallogDiskSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("uploaddir"); ok || d.HasChange("uploaddir") {
		t, err := expandSystemLocallogDiskSettingUploaddir(d, v, "uploaddir")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploaddir"] = t
		}
	}

	if v, ok := d.GetOk("uploadip"); ok || d.HasChange("uploadip") {
		t, err := expandSystemLocallogDiskSettingUploadip(d, v, "uploadip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadip"] = t
		}
	}

	if v, ok := d.GetOk("uploadpass"); ok || d.HasChange("uploadpass") {
		t, err := expandSystemLocallogDiskSettingUploadpass(d, v, "uploadpass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadpass"] = t
		}
	}

	if v, ok := d.GetOk("uploadport"); ok || d.HasChange("uploadport") {
		t, err := expandSystemLocallogDiskSettingUploadport(d, v, "uploadport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadport"] = t
		}
	}

	if v, ok := d.GetOk("uploadsched"); ok || d.HasChange("uploadsched") {
		t, err := expandSystemLocallogDiskSettingUploadsched(d, v, "uploadsched")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadsched"] = t
		}
	}

	if v, ok := d.GetOk("uploadtype"); ok || d.HasChange("uploadtype") {
		t, err := expandSystemLocallogDiskSettingUploadtype(d, v, "uploadtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadtype"] = t
		}
	}

	if v, ok := d.GetOk("uploaduser"); ok || d.HasChange("uploaduser") {
		t, err := expandSystemLocallogDiskSettingUploaduser(d, v, "uploaduser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploaduser"] = t
		}
	}

	if v, ok := d.GetOk("uploadzip"); ok || d.HasChange("uploadzip") {
		t, err := expandSystemLocallogDiskSettingUploadzip(d, v, "uploadzip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadzip"] = t
		}
	}

	return &obj, nil
}
