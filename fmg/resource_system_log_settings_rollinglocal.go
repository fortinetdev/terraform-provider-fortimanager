// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Log rolling policy for local logs.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLogSettingsRollingLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogSettingsRollingLocalUpdate,
		Read:   resourceSystemLogSettingsRollingLocalRead,
		Update: resourceSystemLogSettingsRollingLocalUpdate,
		Delete: resourceSystemLogSettingsRollingLocalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"days": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"del_files": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gzip_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"password3": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rolling_upgrade_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"upload_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"when": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogSettingsRollingLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogSettingsRollingLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettingsRollingLocal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogSettingsRollingLocal(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettingsRollingLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogSettingsRollingLocal")

	return resourceSystemLogSettingsRollingLocalRead(d, m)
}

func resourceSystemLogSettingsRollingLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogSettingsRollingLocal(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogSettingsRollingLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogSettingsRollingLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogSettingsRollingLocal(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettingsRollingLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogSettingsRollingLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettingsRollingLocal resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogSettingsRollingLocalDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalDelFiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalGzipFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalIp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalIp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalLogFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalPassword2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalPassword3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalRollingUpgradeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadTrigger(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUsername2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUsername3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalWhen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogSettingsRollingLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("days", flattenSystemLogSettingsRollingLocalDays(o["days"], d, "days")); err != nil {
		if vv, ok := fortiAPIPatch(o["days"], "SystemLogSettingsRollingLocal-Days"); ok {
			if err = d.Set("days", vv); err != nil {
				return fmt.Errorf("Error reading days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading days: %v", err)
		}
	}

	if err = d.Set("del_files", flattenSystemLogSettingsRollingLocalDelFiles(o["del-files"], d, "del_files")); err != nil {
		if vv, ok := fortiAPIPatch(o["del-files"], "SystemLogSettingsRollingLocal-DelFiles"); ok {
			if err = d.Set("del_files", vv); err != nil {
				return fmt.Errorf("Error reading del_files: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading del_files: %v", err)
		}
	}

	if err = d.Set("directory", flattenSystemLogSettingsRollingLocalDirectory(o["directory"], d, "directory")); err != nil {
		if vv, ok := fortiAPIPatch(o["directory"], "SystemLogSettingsRollingLocal-Directory"); ok {
			if err = d.Set("directory", vv); err != nil {
				return fmt.Errorf("Error reading directory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading directory: %v", err)
		}
	}

	if err = d.Set("file_size", flattenSystemLogSettingsRollingLocalFileSize(o["file-size"], d, "file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-size"], "SystemLogSettingsRollingLocal-FileSize"); ok {
			if err = d.Set("file_size", vv); err != nil {
				return fmt.Errorf("Error reading file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_size: %v", err)
		}
	}

	if err = d.Set("gzip_format", flattenSystemLogSettingsRollingLocalGzipFormat(o["gzip-format"], d, "gzip_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["gzip-format"], "SystemLogSettingsRollingLocal-GzipFormat"); ok {
			if err = d.Set("gzip_format", vv); err != nil {
				return fmt.Errorf("Error reading gzip_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gzip_format: %v", err)
		}
	}

	if err = d.Set("hour", flattenSystemLogSettingsRollingLocalHour(o["hour"], d, "hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["hour"], "SystemLogSettingsRollingLocal-Hour"); ok {
			if err = d.Set("hour", vv); err != nil {
				return fmt.Errorf("Error reading hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hour: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemLogSettingsRollingLocalIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemLogSettingsRollingLocal-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip2", flattenSystemLogSettingsRollingLocalIp2(o["ip2"], d, "ip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip2"], "SystemLogSettingsRollingLocal-Ip2"); ok {
			if err = d.Set("ip2", vv); err != nil {
				return fmt.Errorf("Error reading ip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip2: %v", err)
		}
	}

	if err = d.Set("ip3", flattenSystemLogSettingsRollingLocalIp3(o["ip3"], d, "ip3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip3"], "SystemLogSettingsRollingLocal-Ip3"); ok {
			if err = d.Set("ip3", vv); err != nil {
				return fmt.Errorf("Error reading ip3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip3: %v", err)
		}
	}

	if err = d.Set("log_format", flattenSystemLogSettingsRollingLocalLogFormat(o["log-format"], d, "log_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-format"], "SystemLogSettingsRollingLocal-LogFormat"); ok {
			if err = d.Set("log_format", vv); err != nil {
				return fmt.Errorf("Error reading log_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_format: %v", err)
		}
	}

	if err = d.Set("min", flattenSystemLogSettingsRollingLocalMin(o["min"], d, "min")); err != nil {
		if vv, ok := fortiAPIPatch(o["min"], "SystemLogSettingsRollingLocal-Min"); ok {
			if err = d.Set("min", vv); err != nil {
				return fmt.Errorf("Error reading min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min: %v", err)
		}
	}

	if err = d.Set("password2", flattenSystemLogSettingsRollingLocalPassword2(o["password2"], d, "password2")); err != nil {
		if vv, ok := fortiAPIPatch(o["password2"], "SystemLogSettingsRollingLocal-Password2"); ok {
			if err = d.Set("password2", vv); err != nil {
				return fmt.Errorf("Error reading password2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password2: %v", err)
		}
	}

	if err = d.Set("password3", flattenSystemLogSettingsRollingLocalPassword3(o["password3"], d, "password3")); err != nil {
		if vv, ok := fortiAPIPatch(o["password3"], "SystemLogSettingsRollingLocal-Password3"); ok {
			if err = d.Set("password3", vv); err != nil {
				return fmt.Errorf("Error reading password3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password3: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemLogSettingsRollingLocalPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemLogSettingsRollingLocal-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port2", flattenSystemLogSettingsRollingLocalPort2(o["port2"], d, "port2")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2"], "SystemLogSettingsRollingLocal-Port2"); ok {
			if err = d.Set("port2", vv); err != nil {
				return fmt.Errorf("Error reading port2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("port3", flattenSystemLogSettingsRollingLocalPort3(o["port3"], d, "port3")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3"], "SystemLogSettingsRollingLocal-Port3"); ok {
			if err = d.Set("port3", vv); err != nil {
				return fmt.Errorf("Error reading port3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("rolling_upgrade_status", flattenSystemLogSettingsRollingLocalRollingUpgradeStatus(o["rolling-upgrade-status"], d, "rolling_upgrade_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["rolling-upgrade-status"], "SystemLogSettingsRollingLocal-RollingUpgradeStatus"); ok {
			if err = d.Set("rolling_upgrade_status", vv); err != nil {
				return fmt.Errorf("Error reading rolling_upgrade_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rolling_upgrade_status: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemLogSettingsRollingLocalServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemLogSettingsRollingLocal-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("upload", flattenSystemLogSettingsRollingLocalUpload(o["upload"], d, "upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload"], "SystemLogSettingsRollingLocal-Upload"); ok {
			if err = d.Set("upload", vv); err != nil {
				return fmt.Errorf("Error reading upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_hour", flattenSystemLogSettingsRollingLocalUploadHour(o["upload-hour"], d, "upload_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-hour"], "SystemLogSettingsRollingLocal-UploadHour"); ok {
			if err = d.Set("upload_hour", vv); err != nil {
				return fmt.Errorf("Error reading upload_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_hour: %v", err)
		}
	}

	if err = d.Set("upload_mode", flattenSystemLogSettingsRollingLocalUploadMode(o["upload-mode"], d, "upload_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-mode"], "SystemLogSettingsRollingLocal-UploadMode"); ok {
			if err = d.Set("upload_mode", vv); err != nil {
				return fmt.Errorf("Error reading upload_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_mode: %v", err)
		}
	}

	if err = d.Set("upload_trigger", flattenSystemLogSettingsRollingLocalUploadTrigger(o["upload-trigger"], d, "upload_trigger")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-trigger"], "SystemLogSettingsRollingLocal-UploadTrigger"); ok {
			if err = d.Set("upload_trigger", vv); err != nil {
				return fmt.Errorf("Error reading upload_trigger: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_trigger: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemLogSettingsRollingLocalUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemLogSettingsRollingLocal-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("username2", flattenSystemLogSettingsRollingLocalUsername2(o["username2"], d, "username2")); err != nil {
		if vv, ok := fortiAPIPatch(o["username2"], "SystemLogSettingsRollingLocal-Username2"); ok {
			if err = d.Set("username2", vv); err != nil {
				return fmt.Errorf("Error reading username2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username2: %v", err)
		}
	}

	if err = d.Set("username3", flattenSystemLogSettingsRollingLocalUsername3(o["username3"], d, "username3")); err != nil {
		if vv, ok := fortiAPIPatch(o["username3"], "SystemLogSettingsRollingLocal-Username3"); ok {
			if err = d.Set("username3", vv); err != nil {
				return fmt.Errorf("Error reading username3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username3: %v", err)
		}
	}

	if err = d.Set("when", flattenSystemLogSettingsRollingLocalWhen(o["when"], d, "when")); err != nil {
		if vv, ok := fortiAPIPatch(o["when"], "SystemLogSettingsRollingLocal-When"); ok {
			if err = d.Set("when", vv); err != nil {
				return fmt.Errorf("Error reading when: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading when: %v", err)
		}
	}

	return nil
}

func flattenSystemLogSettingsRollingLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogSettingsRollingLocalDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalDelFiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalGzipFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalIp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalIp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalLogFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalRollingUpgradeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadTrigger(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUsername2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUsername3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalWhen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogSettingsRollingLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("days"); ok || d.HasChange("days") {
		t, err := expandSystemLogSettingsRollingLocalDays(d, v, "days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["days"] = t
		}
	}

	if v, ok := d.GetOk("del_files"); ok || d.HasChange("del_files") {
		t, err := expandSystemLogSettingsRollingLocalDelFiles(d, v, "del_files")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["del-files"] = t
		}
	}

	if v, ok := d.GetOk("directory"); ok || d.HasChange("directory") {
		t, err := expandSystemLogSettingsRollingLocalDirectory(d, v, "directory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["directory"] = t
		}
	}

	if v, ok := d.GetOk("file_size"); ok || d.HasChange("file_size") {
		t, err := expandSystemLogSettingsRollingLocalFileSize(d, v, "file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size"] = t
		}
	}

	if v, ok := d.GetOk("gzip_format"); ok || d.HasChange("gzip_format") {
		t, err := expandSystemLogSettingsRollingLocalGzipFormat(d, v, "gzip_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gzip-format"] = t
		}
	}

	if v, ok := d.GetOk("hour"); ok || d.HasChange("hour") {
		t, err := expandSystemLogSettingsRollingLocalHour(d, v, "hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hour"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemLogSettingsRollingLocalIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip2"); ok || d.HasChange("ip2") {
		t, err := expandSystemLogSettingsRollingLocalIp2(d, v, "ip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip2"] = t
		}
	}

	if v, ok := d.GetOk("ip3"); ok || d.HasChange("ip3") {
		t, err := expandSystemLogSettingsRollingLocalIp3(d, v, "ip3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip3"] = t
		}
	}

	if v, ok := d.GetOk("log_format"); ok || d.HasChange("log_format") {
		t, err := expandSystemLogSettingsRollingLocalLogFormat(d, v, "log_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-format"] = t
		}
	}

	if v, ok := d.GetOk("min"); ok || d.HasChange("min") {
		t, err := expandSystemLogSettingsRollingLocalMin(d, v, "min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemLogSettingsRollingLocalPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok || d.HasChange("password2") {
		t, err := expandSystemLogSettingsRollingLocalPassword2(d, v, "password2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok || d.HasChange("password3") {
		t, err := expandSystemLogSettingsRollingLocalPassword3(d, v, "password3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemLogSettingsRollingLocalPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("port2"); ok || d.HasChange("port2") {
		t, err := expandSystemLogSettingsRollingLocalPort2(d, v, "port2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("port3"); ok || d.HasChange("port3") {
		t, err := expandSystemLogSettingsRollingLocalPort3(d, v, "port3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("rolling_upgrade_status"); ok || d.HasChange("rolling_upgrade_status") {
		t, err := expandSystemLogSettingsRollingLocalRollingUpgradeStatus(d, v, "rolling_upgrade_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-upgrade-status"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemLogSettingsRollingLocalServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("upload"); ok || d.HasChange("upload") {
		t, err := expandSystemLogSettingsRollingLocalUpload(d, v, "upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload"] = t
		}
	}

	if v, ok := d.GetOk("upload_hour"); ok || d.HasChange("upload_hour") {
		t, err := expandSystemLogSettingsRollingLocalUploadHour(d, v, "upload_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-hour"] = t
		}
	}

	if v, ok := d.GetOk("upload_mode"); ok || d.HasChange("upload_mode") {
		t, err := expandSystemLogSettingsRollingLocalUploadMode(d, v, "upload_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-mode"] = t
		}
	}

	if v, ok := d.GetOk("upload_trigger"); ok || d.HasChange("upload_trigger") {
		t, err := expandSystemLogSettingsRollingLocalUploadTrigger(d, v, "upload_trigger")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-trigger"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemLogSettingsRollingLocalUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("username2"); ok || d.HasChange("username2") {
		t, err := expandSystemLogSettingsRollingLocalUsername2(d, v, "username2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username2"] = t
		}
	}

	if v, ok := d.GetOk("username3"); ok || d.HasChange("username3") {
		t, err := expandSystemLogSettingsRollingLocalUsername3(d, v, "username3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username3"] = t
		}
	}

	if v, ok := d.GetOk("when"); ok || d.HasChange("when") {
		t, err := expandSystemLogSettingsRollingLocalWhen(d, v, "when")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["when"] = t
		}
	}

	return &obj, nil
}
