// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Log rolling policy for device logs.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogSettingsRollingRegular() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogSettingsRollingRegularUpdate,
		Read:   resourceSystemLogSettingsRollingRegularRead,
		Update: resourceSystemLogSettingsRollingRegularUpdate,
		Delete: resourceSystemLogSettingsRollingRegularDelete,

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
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password2": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password3": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port3": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rolling_upgrade_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"username2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"when": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogSettingsRollingRegularUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogSettingsRollingRegular(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettingsRollingRegular resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogSettingsRollingRegular(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettingsRollingRegular resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogSettingsRollingRegular")

	return resourceSystemLogSettingsRollingRegularRead(d, m)
}

func resourceSystemLogSettingsRollingRegularDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLogSettingsRollingRegular(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogSettingsRollingRegular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogSettingsRollingRegularRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogSettingsRollingRegular(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettingsRollingRegular resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogSettingsRollingRegular(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettingsRollingRegular resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogSettingsRollingRegularDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularDelFiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularGzipFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularIp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularIp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularLogFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularPassword2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularPassword3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularRollingUpgradeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadTrigger(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUsername2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUsername3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularWhen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogSettingsRollingRegular(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("days", flattenSystemLogSettingsRollingRegularDays(o["days"], d, "days")); err != nil {
		if vv, ok := fortiAPIPatch(o["days"], "SystemLogSettingsRollingRegular-Days"); ok {
			if err = d.Set("days", vv); err != nil {
				return fmt.Errorf("Error reading days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading days: %v", err)
		}
	}

	if err = d.Set("del_files", flattenSystemLogSettingsRollingRegularDelFiles(o["del-files"], d, "del_files")); err != nil {
		if vv, ok := fortiAPIPatch(o["del-files"], "SystemLogSettingsRollingRegular-DelFiles"); ok {
			if err = d.Set("del_files", vv); err != nil {
				return fmt.Errorf("Error reading del_files: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading del_files: %v", err)
		}
	}

	if err = d.Set("directory", flattenSystemLogSettingsRollingRegularDirectory(o["directory"], d, "directory")); err != nil {
		if vv, ok := fortiAPIPatch(o["directory"], "SystemLogSettingsRollingRegular-Directory"); ok {
			if err = d.Set("directory", vv); err != nil {
				return fmt.Errorf("Error reading directory: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading directory: %v", err)
		}
	}

	if err = d.Set("file_size", flattenSystemLogSettingsRollingRegularFileSize(o["file-size"], d, "file_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-size"], "SystemLogSettingsRollingRegular-FileSize"); ok {
			if err = d.Set("file_size", vv); err != nil {
				return fmt.Errorf("Error reading file_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_size: %v", err)
		}
	}

	if err = d.Set("gzip_format", flattenSystemLogSettingsRollingRegularGzipFormat(o["gzip-format"], d, "gzip_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["gzip-format"], "SystemLogSettingsRollingRegular-GzipFormat"); ok {
			if err = d.Set("gzip_format", vv); err != nil {
				return fmt.Errorf("Error reading gzip_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gzip_format: %v", err)
		}
	}

	if err = d.Set("hour", flattenSystemLogSettingsRollingRegularHour(o["hour"], d, "hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["hour"], "SystemLogSettingsRollingRegular-Hour"); ok {
			if err = d.Set("hour", vv); err != nil {
				return fmt.Errorf("Error reading hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hour: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemLogSettingsRollingRegularIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SystemLogSettingsRollingRegular-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip2", flattenSystemLogSettingsRollingRegularIp2(o["ip2"], d, "ip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip2"], "SystemLogSettingsRollingRegular-Ip2"); ok {
			if err = d.Set("ip2", vv); err != nil {
				return fmt.Errorf("Error reading ip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip2: %v", err)
		}
	}

	if err = d.Set("ip3", flattenSystemLogSettingsRollingRegularIp3(o["ip3"], d, "ip3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip3"], "SystemLogSettingsRollingRegular-Ip3"); ok {
			if err = d.Set("ip3", vv); err != nil {
				return fmt.Errorf("Error reading ip3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip3: %v", err)
		}
	}

	if err = d.Set("log_format", flattenSystemLogSettingsRollingRegularLogFormat(o["log-format"], d, "log_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-format"], "SystemLogSettingsRollingRegular-LogFormat"); ok {
			if err = d.Set("log_format", vv); err != nil {
				return fmt.Errorf("Error reading log_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_format: %v", err)
		}
	}

	if err = d.Set("min", flattenSystemLogSettingsRollingRegularMin(o["min"], d, "min")); err != nil {
		if vv, ok := fortiAPIPatch(o["min"], "SystemLogSettingsRollingRegular-Min"); ok {
			if err = d.Set("min", vv); err != nil {
				return fmt.Errorf("Error reading min: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading min: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemLogSettingsRollingRegularPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemLogSettingsRollingRegular-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port2", flattenSystemLogSettingsRollingRegularPort2(o["port2"], d, "port2")); err != nil {
		if vv, ok := fortiAPIPatch(o["port2"], "SystemLogSettingsRollingRegular-Port2"); ok {
			if err = d.Set("port2", vv); err != nil {
				return fmt.Errorf("Error reading port2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("port3", flattenSystemLogSettingsRollingRegularPort3(o["port3"], d, "port3")); err != nil {
		if vv, ok := fortiAPIPatch(o["port3"], "SystemLogSettingsRollingRegular-Port3"); ok {
			if err = d.Set("port3", vv); err != nil {
				return fmt.Errorf("Error reading port3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("rolling_upgrade_status", flattenSystemLogSettingsRollingRegularRollingUpgradeStatus(o["rolling-upgrade-status"], d, "rolling_upgrade_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["rolling-upgrade-status"], "SystemLogSettingsRollingRegular-RollingUpgradeStatus"); ok {
			if err = d.Set("rolling_upgrade_status", vv); err != nil {
				return fmt.Errorf("Error reading rolling_upgrade_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rolling_upgrade_status: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemLogSettingsRollingRegularServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemLogSettingsRollingRegular-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemLogSettingsRollingRegularServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemLogSettingsRollingRegular-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("server2", flattenSystemLogSettingsRollingRegularServer2(o["server2"], d, "server2")); err != nil {
		if vv, ok := fortiAPIPatch(o["server2"], "SystemLogSettingsRollingRegular-Server2"); ok {
			if err = d.Set("server2", vv); err != nil {
				return fmt.Errorf("Error reading server2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server2: %v", err)
		}
	}

	if err = d.Set("server3", flattenSystemLogSettingsRollingRegularServer3(o["server3"], d, "server3")); err != nil {
		if vv, ok := fortiAPIPatch(o["server3"], "SystemLogSettingsRollingRegular-Server3"); ok {
			if err = d.Set("server3", vv); err != nil {
				return fmt.Errorf("Error reading server3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server3: %v", err)
		}
	}

	if err = d.Set("upload", flattenSystemLogSettingsRollingRegularUpload(o["upload"], d, "upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload"], "SystemLogSettingsRollingRegular-Upload"); ok {
			if err = d.Set("upload", vv); err != nil {
				return fmt.Errorf("Error reading upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_hour", flattenSystemLogSettingsRollingRegularUploadHour(o["upload-hour"], d, "upload_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-hour"], "SystemLogSettingsRollingRegular-UploadHour"); ok {
			if err = d.Set("upload_hour", vv); err != nil {
				return fmt.Errorf("Error reading upload_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_hour: %v", err)
		}
	}

	if err = d.Set("upload_mode", flattenSystemLogSettingsRollingRegularUploadMode(o["upload-mode"], d, "upload_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-mode"], "SystemLogSettingsRollingRegular-UploadMode"); ok {
			if err = d.Set("upload_mode", vv); err != nil {
				return fmt.Errorf("Error reading upload_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_mode: %v", err)
		}
	}

	if err = d.Set("upload_trigger", flattenSystemLogSettingsRollingRegularUploadTrigger(o["upload-trigger"], d, "upload_trigger")); err != nil {
		if vv, ok := fortiAPIPatch(o["upload-trigger"], "SystemLogSettingsRollingRegular-UploadTrigger"); ok {
			if err = d.Set("upload_trigger", vv); err != nil {
				return fmt.Errorf("Error reading upload_trigger: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading upload_trigger: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemLogSettingsRollingRegularUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemLogSettingsRollingRegular-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("username2", flattenSystemLogSettingsRollingRegularUsername2(o["username2"], d, "username2")); err != nil {
		if vv, ok := fortiAPIPatch(o["username2"], "SystemLogSettingsRollingRegular-Username2"); ok {
			if err = d.Set("username2", vv); err != nil {
				return fmt.Errorf("Error reading username2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username2: %v", err)
		}
	}

	if err = d.Set("username3", flattenSystemLogSettingsRollingRegularUsername3(o["username3"], d, "username3")); err != nil {
		if vv, ok := fortiAPIPatch(o["username3"], "SystemLogSettingsRollingRegular-Username3"); ok {
			if err = d.Set("username3", vv); err != nil {
				return fmt.Errorf("Error reading username3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username3: %v", err)
		}
	}

	if err = d.Set("when", flattenSystemLogSettingsRollingRegularWhen(o["when"], d, "when")); err != nil {
		if vv, ok := fortiAPIPatch(o["when"], "SystemLogSettingsRollingRegular-When"); ok {
			if err = d.Set("when", vv); err != nil {
				return fmt.Errorf("Error reading when: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading when: %v", err)
		}
	}

	return nil
}

func flattenSystemLogSettingsRollingRegularFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogSettingsRollingRegularDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularDelFiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularGzipFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularIp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularIp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularLogFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularRollingUpgradeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadTrigger(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUsername2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUsername3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularWhen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogSettingsRollingRegular(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("days"); ok || d.HasChange("days") {
		t, err := expandSystemLogSettingsRollingRegularDays(d, v, "days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["days"] = t
		}
	}

	if v, ok := d.GetOk("del_files"); ok || d.HasChange("del_files") {
		t, err := expandSystemLogSettingsRollingRegularDelFiles(d, v, "del_files")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["del-files"] = t
		}
	}

	if v, ok := d.GetOk("directory"); ok || d.HasChange("directory") {
		t, err := expandSystemLogSettingsRollingRegularDirectory(d, v, "directory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["directory"] = t
		}
	}

	if v, ok := d.GetOk("file_size"); ok || d.HasChange("file_size") {
		t, err := expandSystemLogSettingsRollingRegularFileSize(d, v, "file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size"] = t
		}
	}

	if v, ok := d.GetOk("gzip_format"); ok || d.HasChange("gzip_format") {
		t, err := expandSystemLogSettingsRollingRegularGzipFormat(d, v, "gzip_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gzip-format"] = t
		}
	}

	if v, ok := d.GetOk("hour"); ok || d.HasChange("hour") {
		t, err := expandSystemLogSettingsRollingRegularHour(d, v, "hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hour"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSystemLogSettingsRollingRegularIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip2"); ok || d.HasChange("ip2") {
		t, err := expandSystemLogSettingsRollingRegularIp2(d, v, "ip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip2"] = t
		}
	}

	if v, ok := d.GetOk("ip3"); ok || d.HasChange("ip3") {
		t, err := expandSystemLogSettingsRollingRegularIp3(d, v, "ip3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip3"] = t
		}
	}

	if v, ok := d.GetOk("log_format"); ok || d.HasChange("log_format") {
		t, err := expandSystemLogSettingsRollingRegularLogFormat(d, v, "log_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-format"] = t
		}
	}

	if v, ok := d.GetOk("min"); ok || d.HasChange("min") {
		t, err := expandSystemLogSettingsRollingRegularMin(d, v, "min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemLogSettingsRollingRegularPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok || d.HasChange("password2") {
		t, err := expandSystemLogSettingsRollingRegularPassword2(d, v, "password2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok || d.HasChange("password3") {
		t, err := expandSystemLogSettingsRollingRegularPassword3(d, v, "password3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemLogSettingsRollingRegularPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("port2"); ok || d.HasChange("port2") {
		t, err := expandSystemLogSettingsRollingRegularPort2(d, v, "port2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("port3"); ok || d.HasChange("port3") {
		t, err := expandSystemLogSettingsRollingRegularPort3(d, v, "port3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("rolling_upgrade_status"); ok || d.HasChange("rolling_upgrade_status") {
		t, err := expandSystemLogSettingsRollingRegularRollingUpgradeStatus(d, v, "rolling_upgrade_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-upgrade-status"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemLogSettingsRollingRegularServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemLogSettingsRollingRegularServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("server2"); ok || d.HasChange("server2") {
		t, err := expandSystemLogSettingsRollingRegularServer2(d, v, "server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server2"] = t
		}
	}

	if v, ok := d.GetOk("server3"); ok || d.HasChange("server3") {
		t, err := expandSystemLogSettingsRollingRegularServer3(d, v, "server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server3"] = t
		}
	}

	if v, ok := d.GetOk("upload"); ok || d.HasChange("upload") {
		t, err := expandSystemLogSettingsRollingRegularUpload(d, v, "upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload"] = t
		}
	}

	if v, ok := d.GetOk("upload_hour"); ok || d.HasChange("upload_hour") {
		t, err := expandSystemLogSettingsRollingRegularUploadHour(d, v, "upload_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-hour"] = t
		}
	}

	if v, ok := d.GetOk("upload_mode"); ok || d.HasChange("upload_mode") {
		t, err := expandSystemLogSettingsRollingRegularUploadMode(d, v, "upload_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-mode"] = t
		}
	}

	if v, ok := d.GetOk("upload_trigger"); ok || d.HasChange("upload_trigger") {
		t, err := expandSystemLogSettingsRollingRegularUploadTrigger(d, v, "upload_trigger")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-trigger"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemLogSettingsRollingRegularUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("username2"); ok || d.HasChange("username2") {
		t, err := expandSystemLogSettingsRollingRegularUsername2(d, v, "username2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username2"] = t
		}
	}

	if v, ok := d.GetOk("username3"); ok || d.HasChange("username3") {
		t, err := expandSystemLogSettingsRollingRegularUsername3(d, v, "username3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username3"] = t
		}
	}

	if v, ok := d.GetOk("when"); ok || d.HasChange("when") {
		t, err := expandSystemLogSettingsRollingRegularWhen(d, v, "when")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["when"] = t
		}
	}

	return &obj, nil
}
