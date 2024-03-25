// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Log settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogSettingsUpdate,
		Read:   resourceSystemLogSettingsRead,
		Update: resourceSystemLogSettingsUpdate,
		Delete: resourceSystemLogSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fac_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"faz_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fch_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fct_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fdd_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fgt_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fmg_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fml_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fpx_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsa_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fwb_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"browse_max_logfiles": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_auto_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_resolve_dstip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"download_max_logs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ha_auto_migrate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"import_max_logfiles": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"keep_dev_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_file_archive_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_interval_dev_no_logging": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_upload_interval_dev_no_logging": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rolling_analyzer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
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
				},
			},
			"rolling_local": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
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
				},
			},
			"rolling_regular": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
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
				},
			},
			"sync_search_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"unencrypted_logging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogSettings(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLogSettings")

	return resourceSystemLogSettingsRead(d, m)
}

func resourceSystemLogSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLogSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogSettings(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogSettingsFacCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFazCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFchCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFctCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFddCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFgtCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFmgCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFmlCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFpxCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFsaCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFwbCustomField1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsBrowseMaxLogfiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsDeviceAutoDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsDnsResolveDstip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsDownloadMaxLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsHaAutoMigrate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsImportMaxLogfiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsKeepDevLogs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsLogFileArchiveName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsLogIntervalDevNoLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsLogUploadIntervalDevNoLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := i["days"]; ok {
		result["days"] = flattenSystemLogSettingsRollingAnalyzerDays(i["days"], d, pre_append)
	}

	pre_append = pre + ".0." + "del_files"
	if _, ok := i["del-files"]; ok {
		result["del_files"] = flattenSystemLogSettingsRollingAnalyzerDelFiles(i["del-files"], d, pre_append)
	}

	pre_append = pre + ".0." + "directory"
	if _, ok := i["directory"]; ok {
		result["directory"] = flattenSystemLogSettingsRollingAnalyzerDirectory(i["directory"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_size"
	if _, ok := i["file-size"]; ok {
		result["file_size"] = flattenSystemLogSettingsRollingAnalyzerFileSize(i["file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "gzip_format"
	if _, ok := i["gzip-format"]; ok {
		result["gzip_format"] = flattenSystemLogSettingsRollingAnalyzerGzipFormat(i["gzip-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "hour"
	if _, ok := i["hour"]; ok {
		result["hour"] = flattenSystemLogSettingsRollingAnalyzerHour(i["hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenSystemLogSettingsRollingAnalyzerIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip2"
	if _, ok := i["ip2"]; ok {
		result["ip2"] = flattenSystemLogSettingsRollingAnalyzerIp2(i["ip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip3"
	if _, ok := i["ip3"]; ok {
		result["ip3"] = flattenSystemLogSettingsRollingAnalyzerIp3(i["ip3"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_format"
	if _, ok := i["log-format"]; ok {
		result["log_format"] = flattenSystemLogSettingsRollingAnalyzerLogFormat(i["log-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "min"
	if _, ok := i["min"]; ok {
		result["min"] = flattenSystemLogSettingsRollingAnalyzerMin(i["min"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenSystemLogSettingsRollingAnalyzerPassword(i["password"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password"] = c
		}
	}

	pre_append = pre + ".0." + "password2"
	if _, ok := i["password2"]; ok {
		result["password2"] = flattenSystemLogSettingsRollingAnalyzerPassword2(i["password2"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password2"] = c
		}
	}

	pre_append = pre + ".0." + "password3"
	if _, ok := i["password3"]; ok {
		result["password3"] = flattenSystemLogSettingsRollingAnalyzerPassword3(i["password3"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password3"] = c
		}
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenSystemLogSettingsRollingAnalyzerPort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2"
	if _, ok := i["port2"]; ok {
		result["port2"] = flattenSystemLogSettingsRollingAnalyzerPort2(i["port2"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3"
	if _, ok := i["port3"]; ok {
		result["port3"] = flattenSystemLogSettingsRollingAnalyzerPort3(i["port3"], d, pre_append)
	}

	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := i["rolling-upgrade-status"]; ok {
		result["rolling_upgrade_status"] = flattenSystemLogSettingsRollingAnalyzerRollingUpgradeStatus(i["rolling-upgrade-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "server"
	if _, ok := i["server"]; ok {
		result["server"] = flattenSystemLogSettingsRollingAnalyzerServer(i["server"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenSystemLogSettingsRollingAnalyzerServerType(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "server2"
	if _, ok := i["server2"]; ok {
		result["server2"] = flattenSystemLogSettingsRollingAnalyzerServer2(i["server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server3"
	if _, ok := i["server3"]; ok {
		result["server3"] = flattenSystemLogSettingsRollingAnalyzerServer3(i["server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload"
	if _, ok := i["upload"]; ok {
		result["upload"] = flattenSystemLogSettingsRollingAnalyzerUpload(i["upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_hour"
	if _, ok := i["upload-hour"]; ok {
		result["upload_hour"] = flattenSystemLogSettingsRollingAnalyzerUploadHour(i["upload-hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_mode"
	if _, ok := i["upload-mode"]; ok {
		result["upload_mode"] = flattenSystemLogSettingsRollingAnalyzerUploadMode(i["upload-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := i["upload-trigger"]; ok {
		result["upload_trigger"] = flattenSystemLogSettingsRollingAnalyzerUploadTrigger(i["upload-trigger"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenSystemLogSettingsRollingAnalyzerUsername(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "username2"
	if _, ok := i["username2"]; ok {
		result["username2"] = flattenSystemLogSettingsRollingAnalyzerUsername2(i["username2"], d, pre_append)
	}

	pre_append = pre + ".0." + "username3"
	if _, ok := i["username3"]; ok {
		result["username3"] = flattenSystemLogSettingsRollingAnalyzerUsername3(i["username3"], d, pre_append)
	}

	pre_append = pre + ".0." + "when"
	if _, ok := i["when"]; ok {
		result["when"] = flattenSystemLogSettingsRollingAnalyzerWhen(i["when"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLogSettingsRollingAnalyzerDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerDelFiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerDirectory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerGzipFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerIp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerIp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerLogFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerPassword2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerPassword3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerRollingUpgradeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadTrigger(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUsername2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUsername3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerWhen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := i["days"]; ok {
		result["days"] = flattenSystemLogSettingsRollingLocalDays(i["days"], d, pre_append)
	}

	pre_append = pre + ".0." + "del_files"
	if _, ok := i["del-files"]; ok {
		result["del_files"] = flattenSystemLogSettingsRollingLocalDelFiles(i["del-files"], d, pre_append)
	}

	pre_append = pre + ".0." + "directory"
	if _, ok := i["directory"]; ok {
		result["directory"] = flattenSystemLogSettingsRollingLocalDirectory(i["directory"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_size"
	if _, ok := i["file-size"]; ok {
		result["file_size"] = flattenSystemLogSettingsRollingLocalFileSize(i["file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "gzip_format"
	if _, ok := i["gzip-format"]; ok {
		result["gzip_format"] = flattenSystemLogSettingsRollingLocalGzipFormat(i["gzip-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "hour"
	if _, ok := i["hour"]; ok {
		result["hour"] = flattenSystemLogSettingsRollingLocalHour(i["hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenSystemLogSettingsRollingLocalIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip2"
	if _, ok := i["ip2"]; ok {
		result["ip2"] = flattenSystemLogSettingsRollingLocalIp2(i["ip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip3"
	if _, ok := i["ip3"]; ok {
		result["ip3"] = flattenSystemLogSettingsRollingLocalIp3(i["ip3"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_format"
	if _, ok := i["log-format"]; ok {
		result["log_format"] = flattenSystemLogSettingsRollingLocalLogFormat(i["log-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "min"
	if _, ok := i["min"]; ok {
		result["min"] = flattenSystemLogSettingsRollingLocalMin(i["min"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenSystemLogSettingsRollingLocalPassword(i["password"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password"] = c
		}
	}

	pre_append = pre + ".0." + "password2"
	if _, ok := i["password2"]; ok {
		result["password2"] = flattenSystemLogSettingsRollingLocalPassword2(i["password2"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password2"] = c
		}
	}

	pre_append = pre + ".0." + "password3"
	if _, ok := i["password3"]; ok {
		result["password3"] = flattenSystemLogSettingsRollingLocalPassword3(i["password3"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password3"] = c
		}
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenSystemLogSettingsRollingLocalPort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2"
	if _, ok := i["port2"]; ok {
		result["port2"] = flattenSystemLogSettingsRollingLocalPort2(i["port2"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3"
	if _, ok := i["port3"]; ok {
		result["port3"] = flattenSystemLogSettingsRollingLocalPort3(i["port3"], d, pre_append)
	}

	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := i["rolling-upgrade-status"]; ok {
		result["rolling_upgrade_status"] = flattenSystemLogSettingsRollingLocalRollingUpgradeStatus(i["rolling-upgrade-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "server"
	if _, ok := i["server"]; ok {
		result["server"] = flattenSystemLogSettingsRollingLocalServer(i["server"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenSystemLogSettingsRollingLocalServerType(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "server2"
	if _, ok := i["server2"]; ok {
		result["server2"] = flattenSystemLogSettingsRollingLocalServer2(i["server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server3"
	if _, ok := i["server3"]; ok {
		result["server3"] = flattenSystemLogSettingsRollingLocalServer3(i["server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload"
	if _, ok := i["upload"]; ok {
		result["upload"] = flattenSystemLogSettingsRollingLocalUpload(i["upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_hour"
	if _, ok := i["upload-hour"]; ok {
		result["upload_hour"] = flattenSystemLogSettingsRollingLocalUploadHour(i["upload-hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_mode"
	if _, ok := i["upload-mode"]; ok {
		result["upload_mode"] = flattenSystemLogSettingsRollingLocalUploadMode(i["upload-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := i["upload-trigger"]; ok {
		result["upload_trigger"] = flattenSystemLogSettingsRollingLocalUploadTrigger(i["upload-trigger"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenSystemLogSettingsRollingLocalUsername(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "username2"
	if _, ok := i["username2"]; ok {
		result["username2"] = flattenSystemLogSettingsRollingLocalUsername2(i["username2"], d, pre_append)
	}

	pre_append = pre + ".0." + "username3"
	if _, ok := i["username3"]; ok {
		result["username3"] = flattenSystemLogSettingsRollingLocalUsername3(i["username3"], d, pre_append)
	}

	pre_append = pre + ".0." + "when"
	if _, ok := i["when"]; ok {
		result["when"] = flattenSystemLogSettingsRollingLocalWhen(i["when"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
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

func flattenSystemLogSettingsRollingLocalServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenSystemLogSettingsRollingRegular(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := i["days"]; ok {
		result["days"] = flattenSystemLogSettingsRollingRegularDays(i["days"], d, pre_append)
	}

	pre_append = pre + ".0." + "del_files"
	if _, ok := i["del-files"]; ok {
		result["del_files"] = flattenSystemLogSettingsRollingRegularDelFiles(i["del-files"], d, pre_append)
	}

	pre_append = pre + ".0." + "directory"
	if _, ok := i["directory"]; ok {
		result["directory"] = flattenSystemLogSettingsRollingRegularDirectory(i["directory"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_size"
	if _, ok := i["file-size"]; ok {
		result["file_size"] = flattenSystemLogSettingsRollingRegularFileSize(i["file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "gzip_format"
	if _, ok := i["gzip-format"]; ok {
		result["gzip_format"] = flattenSystemLogSettingsRollingRegularGzipFormat(i["gzip-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "hour"
	if _, ok := i["hour"]; ok {
		result["hour"] = flattenSystemLogSettingsRollingRegularHour(i["hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenSystemLogSettingsRollingRegularIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip2"
	if _, ok := i["ip2"]; ok {
		result["ip2"] = flattenSystemLogSettingsRollingRegularIp2(i["ip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip3"
	if _, ok := i["ip3"]; ok {
		result["ip3"] = flattenSystemLogSettingsRollingRegularIp3(i["ip3"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_format"
	if _, ok := i["log-format"]; ok {
		result["log_format"] = flattenSystemLogSettingsRollingRegularLogFormat(i["log-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "min"
	if _, ok := i["min"]; ok {
		result["min"] = flattenSystemLogSettingsRollingRegularMin(i["min"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenSystemLogSettingsRollingRegularPassword(i["password"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password"] = c
		}
	}

	pre_append = pre + ".0." + "password2"
	if _, ok := i["password2"]; ok {
		result["password2"] = flattenSystemLogSettingsRollingRegularPassword2(i["password2"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password2"] = c
		}
	}

	pre_append = pre + ".0." + "password3"
	if _, ok := i["password3"]; ok {
		result["password3"] = flattenSystemLogSettingsRollingRegularPassword3(i["password3"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password3"] = c
		}
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenSystemLogSettingsRollingRegularPort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2"
	if _, ok := i["port2"]; ok {
		result["port2"] = flattenSystemLogSettingsRollingRegularPort2(i["port2"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3"
	if _, ok := i["port3"]; ok {
		result["port3"] = flattenSystemLogSettingsRollingRegularPort3(i["port3"], d, pre_append)
	}

	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := i["rolling-upgrade-status"]; ok {
		result["rolling_upgrade_status"] = flattenSystemLogSettingsRollingRegularRollingUpgradeStatus(i["rolling-upgrade-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "server"
	if _, ok := i["server"]; ok {
		result["server"] = flattenSystemLogSettingsRollingRegularServer(i["server"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenSystemLogSettingsRollingRegularServerType(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "server2"
	if _, ok := i["server2"]; ok {
		result["server2"] = flattenSystemLogSettingsRollingRegularServer2(i["server2"], d, pre_append)
	}

	pre_append = pre + ".0." + "server3"
	if _, ok := i["server3"]; ok {
		result["server3"] = flattenSystemLogSettingsRollingRegularServer3(i["server3"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload"
	if _, ok := i["upload"]; ok {
		result["upload"] = flattenSystemLogSettingsRollingRegularUpload(i["upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_hour"
	if _, ok := i["upload-hour"]; ok {
		result["upload_hour"] = flattenSystemLogSettingsRollingRegularUploadHour(i["upload-hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_mode"
	if _, ok := i["upload-mode"]; ok {
		result["upload_mode"] = flattenSystemLogSettingsRollingRegularUploadMode(i["upload-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := i["upload-trigger"]; ok {
		result["upload_trigger"] = flattenSystemLogSettingsRollingRegularUploadTrigger(i["upload-trigger"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenSystemLogSettingsRollingRegularUsername(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "username2"
	if _, ok := i["username2"]; ok {
		result["username2"] = flattenSystemLogSettingsRollingRegularUsername2(i["username2"], d, pre_append)
	}

	pre_append = pre + ".0." + "username3"
	if _, ok := i["username3"]; ok {
		result["username3"] = flattenSystemLogSettingsRollingRegularUsername3(i["username3"], d, pre_append)
	}

	pre_append = pre + ".0." + "when"
	if _, ok := i["when"]; ok {
		result["when"] = flattenSystemLogSettingsRollingRegularWhen(i["when"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
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

func flattenSystemLogSettingsSyncSearchTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsUnencryptedLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fac_custom_field1", flattenSystemLogSettingsFacCustomField1(o["FAC-custom-field1"], d, "fac_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FAC-custom-field1"], "SystemLogSettings-FacCustomField1"); ok {
			if err = d.Set("fac_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fac_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fac_custom_field1: %v", err)
		}
	}

	if err = d.Set("faz_custom_field1", flattenSystemLogSettingsFazCustomField1(o["FAZ-custom-field1"], d, "faz_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FAZ-custom-field1"], "SystemLogSettings-FazCustomField1"); ok {
			if err = d.Set("faz_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading faz_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_custom_field1: %v", err)
		}
	}

	if err = d.Set("fch_custom_field1", flattenSystemLogSettingsFchCustomField1(o["FCH-custom-field1"], d, "fch_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FCH-custom-field1"], "SystemLogSettings-FchCustomField1"); ok {
			if err = d.Set("fch_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fch_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fch_custom_field1: %v", err)
		}
	}

	if err = d.Set("fct_custom_field1", flattenSystemLogSettingsFctCustomField1(o["FCT-custom-field1"], d, "fct_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FCT-custom-field1"], "SystemLogSettings-FctCustomField1"); ok {
			if err = d.Set("fct_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fct_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fct_custom_field1: %v", err)
		}
	}

	if err = d.Set("fdd_custom_field1", flattenSystemLogSettingsFddCustomField1(o["FDD-custom-field1"], d, "fdd_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FDD-custom-field1"], "SystemLogSettings-FddCustomField1"); ok {
			if err = d.Set("fdd_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fdd_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fdd_custom_field1: %v", err)
		}
	}

	if err = d.Set("fgt_custom_field1", flattenSystemLogSettingsFgtCustomField1(o["FGT-custom-field1"], d, "fgt_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FGT-custom-field1"], "SystemLogSettings-FgtCustomField1"); ok {
			if err = d.Set("fgt_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fgt_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_custom_field1: %v", err)
		}
	}

	if err = d.Set("fmg_custom_field1", flattenSystemLogSettingsFmgCustomField1(o["FMG-custom-field1"], d, "fmg_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FMG-custom-field1"], "SystemLogSettings-FmgCustomField1"); ok {
			if err = d.Set("fmg_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fmg_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg_custom_field1: %v", err)
		}
	}

	if err = d.Set("fml_custom_field1", flattenSystemLogSettingsFmlCustomField1(o["FML-custom-field1"], d, "fml_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FML-custom-field1"], "SystemLogSettings-FmlCustomField1"); ok {
			if err = d.Set("fml_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fml_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fml_custom_field1: %v", err)
		}
	}

	if err = d.Set("fpx_custom_field1", flattenSystemLogSettingsFpxCustomField1(o["FPX-custom-field1"], d, "fpx_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FPX-custom-field1"], "SystemLogSettings-FpxCustomField1"); ok {
			if err = d.Set("fpx_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fpx_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fpx_custom_field1: %v", err)
		}
	}

	if err = d.Set("fsa_custom_field1", flattenSystemLogSettingsFsaCustomField1(o["FSA-custom-field1"], d, "fsa_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FSA-custom-field1"], "SystemLogSettings-FsaCustomField1"); ok {
			if err = d.Set("fsa_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fsa_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsa_custom_field1: %v", err)
		}
	}

	if err = d.Set("fwb_custom_field1", flattenSystemLogSettingsFwbCustomField1(o["FWB-custom-field1"], d, "fwb_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FWB-custom-field1"], "SystemLogSettings-FwbCustomField1"); ok {
			if err = d.Set("fwb_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fwb_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwb_custom_field1: %v", err)
		}
	}

	if err = d.Set("browse_max_logfiles", flattenSystemLogSettingsBrowseMaxLogfiles(o["browse-max-logfiles"], d, "browse_max_logfiles")); err != nil {
		if vv, ok := fortiAPIPatch(o["browse-max-logfiles"], "SystemLogSettings-BrowseMaxLogfiles"); ok {
			if err = d.Set("browse_max_logfiles", vv); err != nil {
				return fmt.Errorf("Error reading browse_max_logfiles: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading browse_max_logfiles: %v", err)
		}
	}

	if err = d.Set("device_auto_detect", flattenSystemLogSettingsDeviceAutoDetect(o["device-auto-detect"], d, "device_auto_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-auto-detect"], "SystemLogSettings-DeviceAutoDetect"); ok {
			if err = d.Set("device_auto_detect", vv); err != nil {
				return fmt.Errorf("Error reading device_auto_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_auto_detect: %v", err)
		}
	}

	if err = d.Set("dns_resolve_dstip", flattenSystemLogSettingsDnsResolveDstip(o["dns-resolve-dstip"], d, "dns_resolve_dstip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-resolve-dstip"], "SystemLogSettings-DnsResolveDstip"); ok {
			if err = d.Set("dns_resolve_dstip", vv); err != nil {
				return fmt.Errorf("Error reading dns_resolve_dstip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_resolve_dstip: %v", err)
		}
	}

	if err = d.Set("download_max_logs", flattenSystemLogSettingsDownloadMaxLogs(o["download-max-logs"], d, "download_max_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["download-max-logs"], "SystemLogSettings-DownloadMaxLogs"); ok {
			if err = d.Set("download_max_logs", vv); err != nil {
				return fmt.Errorf("Error reading download_max_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading download_max_logs: %v", err)
		}
	}

	if err = d.Set("ha_auto_migrate", flattenSystemLogSettingsHaAutoMigrate(o["ha-auto-migrate"], d, "ha_auto_migrate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-auto-migrate"], "SystemLogSettings-HaAutoMigrate"); ok {
			if err = d.Set("ha_auto_migrate", vv); err != nil {
				return fmt.Errorf("Error reading ha_auto_migrate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_auto_migrate: %v", err)
		}
	}

	if err = d.Set("import_max_logfiles", flattenSystemLogSettingsImportMaxLogfiles(o["import-max-logfiles"], d, "import_max_logfiles")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-max-logfiles"], "SystemLogSettings-ImportMaxLogfiles"); ok {
			if err = d.Set("import_max_logfiles", vv); err != nil {
				return fmt.Errorf("Error reading import_max_logfiles: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_max_logfiles: %v", err)
		}
	}

	if err = d.Set("keep_dev_logs", flattenSystemLogSettingsKeepDevLogs(o["keep-dev-logs"], d, "keep_dev_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-dev-logs"], "SystemLogSettings-KeepDevLogs"); ok {
			if err = d.Set("keep_dev_logs", vv); err != nil {
				return fmt.Errorf("Error reading keep_dev_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_dev_logs: %v", err)
		}
	}

	if err = d.Set("log_file_archive_name", flattenSystemLogSettingsLogFileArchiveName(o["log-file-archive-name"], d, "log_file_archive_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-file-archive-name"], "SystemLogSettings-LogFileArchiveName"); ok {
			if err = d.Set("log_file_archive_name", vv); err != nil {
				return fmt.Errorf("Error reading log_file_archive_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_file_archive_name: %v", err)
		}
	}

	if err = d.Set("log_interval_dev_no_logging", flattenSystemLogSettingsLogIntervalDevNoLogging(o["log-interval-dev-no-logging"], d, "log_interval_dev_no_logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-interval-dev-no-logging"], "SystemLogSettings-LogIntervalDevNoLogging"); ok {
			if err = d.Set("log_interval_dev_no_logging", vv); err != nil {
				return fmt.Errorf("Error reading log_interval_dev_no_logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_interval_dev_no_logging: %v", err)
		}
	}

	if err = d.Set("log_upload_interval_dev_no_logging", flattenSystemLogSettingsLogUploadIntervalDevNoLogging(o["log-upload-interval-dev-no-logging"], d, "log_upload_interval_dev_no_logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-upload-interval-dev-no-logging"], "SystemLogSettings-LogUploadIntervalDevNoLogging"); ok {
			if err = d.Set("log_upload_interval_dev_no_logging", vv); err != nil {
				return fmt.Errorf("Error reading log_upload_interval_dev_no_logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_upload_interval_dev_no_logging: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rolling_analyzer", flattenSystemLogSettingsRollingAnalyzer(o["rolling-analyzer"], d, "rolling_analyzer")); err != nil {
			if vv, ok := fortiAPIPatch(o["rolling-analyzer"], "SystemLogSettings-RollingAnalyzer"); ok {
				if err = d.Set("rolling_analyzer", vv); err != nil {
					return fmt.Errorf("Error reading rolling_analyzer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rolling_analyzer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rolling_analyzer"); ok {
			if err = d.Set("rolling_analyzer", flattenSystemLogSettingsRollingAnalyzer(o["rolling-analyzer"], d, "rolling_analyzer")); err != nil {
				if vv, ok := fortiAPIPatch(o["rolling-analyzer"], "SystemLogSettings-RollingAnalyzer"); ok {
					if err = d.Set("rolling_analyzer", vv); err != nil {
						return fmt.Errorf("Error reading rolling_analyzer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rolling_analyzer: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("rolling_local", flattenSystemLogSettingsRollingLocal(o["rolling-local"], d, "rolling_local")); err != nil {
			if vv, ok := fortiAPIPatch(o["rolling-local"], "SystemLogSettings-RollingLocal"); ok {
				if err = d.Set("rolling_local", vv); err != nil {
					return fmt.Errorf("Error reading rolling_local: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rolling_local: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rolling_local"); ok {
			if err = d.Set("rolling_local", flattenSystemLogSettingsRollingLocal(o["rolling-local"], d, "rolling_local")); err != nil {
				if vv, ok := fortiAPIPatch(o["rolling-local"], "SystemLogSettings-RollingLocal"); ok {
					if err = d.Set("rolling_local", vv); err != nil {
						return fmt.Errorf("Error reading rolling_local: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rolling_local: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("rolling_regular", flattenSystemLogSettingsRollingRegular(o["rolling-regular"], d, "rolling_regular")); err != nil {
			if vv, ok := fortiAPIPatch(o["rolling-regular"], "SystemLogSettings-RollingRegular"); ok {
				if err = d.Set("rolling_regular", vv); err != nil {
					return fmt.Errorf("Error reading rolling_regular: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading rolling_regular: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rolling_regular"); ok {
			if err = d.Set("rolling_regular", flattenSystemLogSettingsRollingRegular(o["rolling-regular"], d, "rolling_regular")); err != nil {
				if vv, ok := fortiAPIPatch(o["rolling-regular"], "SystemLogSettings-RollingRegular"); ok {
					if err = d.Set("rolling_regular", vv); err != nil {
						return fmt.Errorf("Error reading rolling_regular: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading rolling_regular: %v", err)
				}
			}
		}
	}

	if err = d.Set("sync_search_timeout", flattenSystemLogSettingsSyncSearchTimeout(o["sync-search-timeout"], d, "sync_search_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-search-timeout"], "SystemLogSettings-SyncSearchTimeout"); ok {
			if err = d.Set("sync_search_timeout", vv); err != nil {
				return fmt.Errorf("Error reading sync_search_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_search_timeout: %v", err)
		}
	}

	if err = d.Set("unencrypted_logging", flattenSystemLogSettingsUnencryptedLogging(o["unencrypted-logging"], d, "unencrypted_logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["unencrypted-logging"], "SystemLogSettings-UnencryptedLogging"); ok {
			if err = d.Set("unencrypted_logging", vv); err != nil {
				return fmt.Errorf("Error reading unencrypted_logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unencrypted_logging: %v", err)
		}
	}

	return nil
}

func flattenSystemLogSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogSettingsFacCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFazCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFchCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFctCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFddCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFgtCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFmgCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFmlCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFpxCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFsaCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFwbCustomField1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsBrowseMaxLogfiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsDeviceAutoDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsDnsResolveDstip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsDownloadMaxLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsHaAutoMigrate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsImportMaxLogfiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsKeepDevLogs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsLogFileArchiveName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsLogIntervalDevNoLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsLogUploadIntervalDevNoLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days"], _ = expandSystemLogSettingsRollingAnalyzerDays(d, i["days"], pre_append)
	}
	pre_append = pre + ".0." + "del_files"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["del-files"], _ = expandSystemLogSettingsRollingAnalyzerDelFiles(d, i["del_files"], pre_append)
	}
	pre_append = pre + ".0." + "directory"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["directory"], _ = expandSystemLogSettingsRollingAnalyzerDirectory(d, i["directory"], pre_append)
	}
	pre_append = pre + ".0." + "file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-size"], _ = expandSystemLogSettingsRollingAnalyzerFileSize(d, i["file_size"], pre_append)
	}
	pre_append = pre + ".0." + "gzip_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gzip-format"], _ = expandSystemLogSettingsRollingAnalyzerGzipFormat(d, i["gzip_format"], pre_append)
	}
	pre_append = pre + ".0." + "hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hour"], _ = expandSystemLogSettingsRollingAnalyzerHour(d, i["hour"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandSystemLogSettingsRollingAnalyzerIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip2"], _ = expandSystemLogSettingsRollingAnalyzerIp2(d, i["ip2"], pre_append)
	}
	pre_append = pre + ".0." + "ip3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip3"], _ = expandSystemLogSettingsRollingAnalyzerIp3(d, i["ip3"], pre_append)
	}
	pre_append = pre + ".0." + "log_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-format"], _ = expandSystemLogSettingsRollingAnalyzerLogFormat(d, i["log_format"], pre_append)
	}
	pre_append = pre + ".0." + "min"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min"], _ = expandSystemLogSettingsRollingAnalyzerMin(d, i["min"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandSystemLogSettingsRollingAnalyzerPassword(d, i["password"], pre_append)
	}
	pre_append = pre + ".0." + "password2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password2"], _ = expandSystemLogSettingsRollingAnalyzerPassword2(d, i["password2"], pre_append)
	}
	pre_append = pre + ".0." + "password3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password3"], _ = expandSystemLogSettingsRollingAnalyzerPassword3(d, i["password3"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandSystemLogSettingsRollingAnalyzerPort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "port2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2"], _ = expandSystemLogSettingsRollingAnalyzerPort2(d, i["port2"], pre_append)
	}
	pre_append = pre + ".0." + "port3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3"], _ = expandSystemLogSettingsRollingAnalyzerPort3(d, i["port3"], pre_append)
	}
	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rolling-upgrade-status"], _ = expandSystemLogSettingsRollingAnalyzerRollingUpgradeStatus(d, i["rolling_upgrade_status"], pre_append)
	}
	pre_append = pre + ".0." + "server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server"], _ = expandSystemLogSettingsRollingAnalyzerServer(d, i["server"], pre_append)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandSystemLogSettingsRollingAnalyzerServerType(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server2"], _ = expandSystemLogSettingsRollingAnalyzerServer2(d, i["server2"], pre_append)
	}
	pre_append = pre + ".0." + "server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server3"], _ = expandSystemLogSettingsRollingAnalyzerServer3(d, i["server3"], pre_append)
	}
	pre_append = pre + ".0." + "upload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload"], _ = expandSystemLogSettingsRollingAnalyzerUpload(d, i["upload"], pre_append)
	}
	pre_append = pre + ".0." + "upload_hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-hour"], _ = expandSystemLogSettingsRollingAnalyzerUploadHour(d, i["upload_hour"], pre_append)
	}
	pre_append = pre + ".0." + "upload_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-mode"], _ = expandSystemLogSettingsRollingAnalyzerUploadMode(d, i["upload_mode"], pre_append)
	}
	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-trigger"], _ = expandSystemLogSettingsRollingAnalyzerUploadTrigger(d, i["upload_trigger"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandSystemLogSettingsRollingAnalyzerUsername(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "username2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username2"], _ = expandSystemLogSettingsRollingAnalyzerUsername2(d, i["username2"], pre_append)
	}
	pre_append = pre + ".0." + "username3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username3"], _ = expandSystemLogSettingsRollingAnalyzerUsername3(d, i["username3"], pre_append)
	}
	pre_append = pre + ".0." + "when"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["when"], _ = expandSystemLogSettingsRollingAnalyzerWhen(d, i["when"], pre_append)
	}

	return result, nil
}

func expandSystemLogSettingsRollingAnalyzerDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerDelFiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerDirectory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerGzipFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerIp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerIp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerLogFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerRollingUpgradeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadTrigger(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUsername2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUsername3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerWhen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days"], _ = expandSystemLogSettingsRollingLocalDays(d, i["days"], pre_append)
	}
	pre_append = pre + ".0." + "del_files"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["del-files"], _ = expandSystemLogSettingsRollingLocalDelFiles(d, i["del_files"], pre_append)
	}
	pre_append = pre + ".0." + "directory"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["directory"], _ = expandSystemLogSettingsRollingLocalDirectory(d, i["directory"], pre_append)
	}
	pre_append = pre + ".0." + "file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-size"], _ = expandSystemLogSettingsRollingLocalFileSize(d, i["file_size"], pre_append)
	}
	pre_append = pre + ".0." + "gzip_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gzip-format"], _ = expandSystemLogSettingsRollingLocalGzipFormat(d, i["gzip_format"], pre_append)
	}
	pre_append = pre + ".0." + "hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hour"], _ = expandSystemLogSettingsRollingLocalHour(d, i["hour"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandSystemLogSettingsRollingLocalIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip2"], _ = expandSystemLogSettingsRollingLocalIp2(d, i["ip2"], pre_append)
	}
	pre_append = pre + ".0." + "ip3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip3"], _ = expandSystemLogSettingsRollingLocalIp3(d, i["ip3"], pre_append)
	}
	pre_append = pre + ".0." + "log_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-format"], _ = expandSystemLogSettingsRollingLocalLogFormat(d, i["log_format"], pre_append)
	}
	pre_append = pre + ".0." + "min"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min"], _ = expandSystemLogSettingsRollingLocalMin(d, i["min"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandSystemLogSettingsRollingLocalPassword(d, i["password"], pre_append)
	}
	pre_append = pre + ".0." + "password2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password2"], _ = expandSystemLogSettingsRollingLocalPassword2(d, i["password2"], pre_append)
	}
	pre_append = pre + ".0." + "password3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password3"], _ = expandSystemLogSettingsRollingLocalPassword3(d, i["password3"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandSystemLogSettingsRollingLocalPort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "port2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2"], _ = expandSystemLogSettingsRollingLocalPort2(d, i["port2"], pre_append)
	}
	pre_append = pre + ".0." + "port3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3"], _ = expandSystemLogSettingsRollingLocalPort3(d, i["port3"], pre_append)
	}
	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rolling-upgrade-status"], _ = expandSystemLogSettingsRollingLocalRollingUpgradeStatus(d, i["rolling_upgrade_status"], pre_append)
	}
	pre_append = pre + ".0." + "server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server"], _ = expandSystemLogSettingsRollingLocalServer(d, i["server"], pre_append)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandSystemLogSettingsRollingLocalServerType(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server2"], _ = expandSystemLogSettingsRollingLocalServer2(d, i["server2"], pre_append)
	}
	pre_append = pre + ".0." + "server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server3"], _ = expandSystemLogSettingsRollingLocalServer3(d, i["server3"], pre_append)
	}
	pre_append = pre + ".0." + "upload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload"], _ = expandSystemLogSettingsRollingLocalUpload(d, i["upload"], pre_append)
	}
	pre_append = pre + ".0." + "upload_hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-hour"], _ = expandSystemLogSettingsRollingLocalUploadHour(d, i["upload_hour"], pre_append)
	}
	pre_append = pre + ".0." + "upload_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-mode"], _ = expandSystemLogSettingsRollingLocalUploadMode(d, i["upload_mode"], pre_append)
	}
	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-trigger"], _ = expandSystemLogSettingsRollingLocalUploadTrigger(d, i["upload_trigger"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandSystemLogSettingsRollingLocalUsername(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "username2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username2"], _ = expandSystemLogSettingsRollingLocalUsername2(d, i["username2"], pre_append)
	}
	pre_append = pre + ".0." + "username3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username3"], _ = expandSystemLogSettingsRollingLocalUsername3(d, i["username3"], pre_append)
	}
	pre_append = pre + ".0." + "when"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["when"], _ = expandSystemLogSettingsRollingLocalWhen(d, i["when"], pre_append)
	}

	return result, nil
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

func expandSystemLogSettingsRollingLocalServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

func expandSystemLogSettingsRollingRegular(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days"], _ = expandSystemLogSettingsRollingRegularDays(d, i["days"], pre_append)
	}
	pre_append = pre + ".0." + "del_files"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["del-files"], _ = expandSystemLogSettingsRollingRegularDelFiles(d, i["del_files"], pre_append)
	}
	pre_append = pre + ".0." + "directory"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["directory"], _ = expandSystemLogSettingsRollingRegularDirectory(d, i["directory"], pre_append)
	}
	pre_append = pre + ".0." + "file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-size"], _ = expandSystemLogSettingsRollingRegularFileSize(d, i["file_size"], pre_append)
	}
	pre_append = pre + ".0." + "gzip_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gzip-format"], _ = expandSystemLogSettingsRollingRegularGzipFormat(d, i["gzip_format"], pre_append)
	}
	pre_append = pre + ".0." + "hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hour"], _ = expandSystemLogSettingsRollingRegularHour(d, i["hour"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandSystemLogSettingsRollingRegularIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip2"], _ = expandSystemLogSettingsRollingRegularIp2(d, i["ip2"], pre_append)
	}
	pre_append = pre + ".0." + "ip3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip3"], _ = expandSystemLogSettingsRollingRegularIp3(d, i["ip3"], pre_append)
	}
	pre_append = pre + ".0." + "log_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-format"], _ = expandSystemLogSettingsRollingRegularLogFormat(d, i["log_format"], pre_append)
	}
	pre_append = pre + ".0." + "min"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min"], _ = expandSystemLogSettingsRollingRegularMin(d, i["min"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandSystemLogSettingsRollingRegularPassword(d, i["password"], pre_append)
	}
	pre_append = pre + ".0." + "password2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password2"], _ = expandSystemLogSettingsRollingRegularPassword2(d, i["password2"], pre_append)
	}
	pre_append = pre + ".0." + "password3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password3"], _ = expandSystemLogSettingsRollingRegularPassword3(d, i["password3"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandSystemLogSettingsRollingRegularPort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "port2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2"], _ = expandSystemLogSettingsRollingRegularPort2(d, i["port2"], pre_append)
	}
	pre_append = pre + ".0." + "port3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3"], _ = expandSystemLogSettingsRollingRegularPort3(d, i["port3"], pre_append)
	}
	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rolling-upgrade-status"], _ = expandSystemLogSettingsRollingRegularRollingUpgradeStatus(d, i["rolling_upgrade_status"], pre_append)
	}
	pre_append = pre + ".0." + "server"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server"], _ = expandSystemLogSettingsRollingRegularServer(d, i["server"], pre_append)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandSystemLogSettingsRollingRegularServerType(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "server2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server2"], _ = expandSystemLogSettingsRollingRegularServer2(d, i["server2"], pre_append)
	}
	pre_append = pre + ".0." + "server3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server3"], _ = expandSystemLogSettingsRollingRegularServer3(d, i["server3"], pre_append)
	}
	pre_append = pre + ".0." + "upload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload"], _ = expandSystemLogSettingsRollingRegularUpload(d, i["upload"], pre_append)
	}
	pre_append = pre + ".0." + "upload_hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-hour"], _ = expandSystemLogSettingsRollingRegularUploadHour(d, i["upload_hour"], pre_append)
	}
	pre_append = pre + ".0." + "upload_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-mode"], _ = expandSystemLogSettingsRollingRegularUploadMode(d, i["upload_mode"], pre_append)
	}
	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-trigger"], _ = expandSystemLogSettingsRollingRegularUploadTrigger(d, i["upload_trigger"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandSystemLogSettingsRollingRegularUsername(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "username2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username2"], _ = expandSystemLogSettingsRollingRegularUsername2(d, i["username2"], pre_append)
	}
	pre_append = pre + ".0." + "username3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username3"], _ = expandSystemLogSettingsRollingRegularUsername3(d, i["username3"], pre_append)
	}
	pre_append = pre + ".0." + "when"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["when"], _ = expandSystemLogSettingsRollingRegularWhen(d, i["when"], pre_append)
	}

	return result, nil
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

func expandSystemLogSettingsSyncSearchTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsUnencryptedLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fac_custom_field1"); ok || d.HasChange("fac_custom_field1") {
		t, err := expandSystemLogSettingsFacCustomField1(d, v, "fac_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FAC-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("faz_custom_field1"); ok || d.HasChange("faz_custom_field1") {
		t, err := expandSystemLogSettingsFazCustomField1(d, v, "faz_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FAZ-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fch_custom_field1"); ok || d.HasChange("fch_custom_field1") {
		t, err := expandSystemLogSettingsFchCustomField1(d, v, "fch_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FCH-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fct_custom_field1"); ok || d.HasChange("fct_custom_field1") {
		t, err := expandSystemLogSettingsFctCustomField1(d, v, "fct_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FCT-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fdd_custom_field1"); ok || d.HasChange("fdd_custom_field1") {
		t, err := expandSystemLogSettingsFddCustomField1(d, v, "fdd_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FDD-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fgt_custom_field1"); ok || d.HasChange("fgt_custom_field1") {
		t, err := expandSystemLogSettingsFgtCustomField1(d, v, "fgt_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FGT-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fmg_custom_field1"); ok || d.HasChange("fmg_custom_field1") {
		t, err := expandSystemLogSettingsFmgCustomField1(d, v, "fmg_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FMG-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fml_custom_field1"); ok || d.HasChange("fml_custom_field1") {
		t, err := expandSystemLogSettingsFmlCustomField1(d, v, "fml_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FML-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fpx_custom_field1"); ok || d.HasChange("fpx_custom_field1") {
		t, err := expandSystemLogSettingsFpxCustomField1(d, v, "fpx_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FPX-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fsa_custom_field1"); ok || d.HasChange("fsa_custom_field1") {
		t, err := expandSystemLogSettingsFsaCustomField1(d, v, "fsa_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FSA-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fwb_custom_field1"); ok || d.HasChange("fwb_custom_field1") {
		t, err := expandSystemLogSettingsFwbCustomField1(d, v, "fwb_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FWB-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("browse_max_logfiles"); ok || d.HasChange("browse_max_logfiles") {
		t, err := expandSystemLogSettingsBrowseMaxLogfiles(d, v, "browse_max_logfiles")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["browse-max-logfiles"] = t
		}
	}

	if v, ok := d.GetOk("device_auto_detect"); ok || d.HasChange("device_auto_detect") {
		t, err := expandSystemLogSettingsDeviceAutoDetect(d, v, "device_auto_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-auto-detect"] = t
		}
	}

	if v, ok := d.GetOk("dns_resolve_dstip"); ok || d.HasChange("dns_resolve_dstip") {
		t, err := expandSystemLogSettingsDnsResolveDstip(d, v, "dns_resolve_dstip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-resolve-dstip"] = t
		}
	}

	if v, ok := d.GetOk("download_max_logs"); ok || d.HasChange("download_max_logs") {
		t, err := expandSystemLogSettingsDownloadMaxLogs(d, v, "download_max_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["download-max-logs"] = t
		}
	}

	if v, ok := d.GetOk("ha_auto_migrate"); ok || d.HasChange("ha_auto_migrate") {
		t, err := expandSystemLogSettingsHaAutoMigrate(d, v, "ha_auto_migrate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-auto-migrate"] = t
		}
	}

	if v, ok := d.GetOk("import_max_logfiles"); ok || d.HasChange("import_max_logfiles") {
		t, err := expandSystemLogSettingsImportMaxLogfiles(d, v, "import_max_logfiles")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-max-logfiles"] = t
		}
	}

	if v, ok := d.GetOk("keep_dev_logs"); ok || d.HasChange("keep_dev_logs") {
		t, err := expandSystemLogSettingsKeepDevLogs(d, v, "keep_dev_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-dev-logs"] = t
		}
	}

	if v, ok := d.GetOk("log_file_archive_name"); ok || d.HasChange("log_file_archive_name") {
		t, err := expandSystemLogSettingsLogFileArchiveName(d, v, "log_file_archive_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-file-archive-name"] = t
		}
	}

	if v, ok := d.GetOk("log_interval_dev_no_logging"); ok || d.HasChange("log_interval_dev_no_logging") {
		t, err := expandSystemLogSettingsLogIntervalDevNoLogging(d, v, "log_interval_dev_no_logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-interval-dev-no-logging"] = t
		}
	}

	if v, ok := d.GetOk("log_upload_interval_dev_no_logging"); ok || d.HasChange("log_upload_interval_dev_no_logging") {
		t, err := expandSystemLogSettingsLogUploadIntervalDevNoLogging(d, v, "log_upload_interval_dev_no_logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-upload-interval-dev-no-logging"] = t
		}
	}

	if v, ok := d.GetOk("rolling_analyzer"); ok || d.HasChange("rolling_analyzer") {
		t, err := expandSystemLogSettingsRollingAnalyzer(d, v, "rolling_analyzer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-analyzer"] = t
		}
	}

	if v, ok := d.GetOk("rolling_local"); ok || d.HasChange("rolling_local") {
		t, err := expandSystemLogSettingsRollingLocal(d, v, "rolling_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-local"] = t
		}
	}

	if v, ok := d.GetOk("rolling_regular"); ok || d.HasChange("rolling_regular") {
		t, err := expandSystemLogSettingsRollingRegular(d, v, "rolling_regular")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-regular"] = t
		}
	}

	if v, ok := d.GetOk("sync_search_timeout"); ok || d.HasChange("sync_search_timeout") {
		t, err := expandSystemLogSettingsSyncSearchTimeout(d, v, "sync_search_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-search-timeout"] = t
		}
	}

	if v, ok := d.GetOk("unencrypted_logging"); ok || d.HasChange("unencrypted_logging") {
		t, err := expandSystemLogSettingsUnencryptedLogging(d, v, "unencrypted_logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unencrypted-logging"] = t
		}
	}

	return &obj, nil
}
