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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed: true,
			},
			"faz_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fch_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fct_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fdd_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgt_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fml_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fpx_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsa_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwb_custom_field1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"browse_max_logfiles": &schema.Schema{
				Type:     schema.TypeInt,
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
			"rolling_analyzer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
				},
			},
			"rolling_local": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
				},
			},
			"rolling_regular": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
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
				},
			},
			"sync_search_timeout": &schema.Schema{
				Type:     schema.TypeInt,
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

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLogSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogSettings(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLogSettings(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLogSettings(adomv, mkey, nil)
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

func flattenSystemLogSettingsFacCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFazCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFchCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFctCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFddCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFgtCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFmgCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFmlCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFpxCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFsaCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsFwbCustomField1Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsBrowseMaxLogfilesSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsDnsResolveDstipSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsDownloadMaxLogsSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsHaAutoMigrateSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsImportMaxLogfilesSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsKeepDevLogsSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsLogFileArchiveNameSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerSlsa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := i["days"]; ok {
		result["days"] = flattenSystemLogSettingsRollingAnalyzerDaysSlsa(i["days"], d, pre_append)
	}

	pre_append = pre + ".0." + "del_files"
	if _, ok := i["del-files"]; ok {
		result["del_files"] = flattenSystemLogSettingsRollingAnalyzerDelFilesSlsa(i["del-files"], d, pre_append)
	}

	pre_append = pre + ".0." + "directory"
	if _, ok := i["directory"]; ok {
		result["directory"] = flattenSystemLogSettingsRollingAnalyzerDirectorySlsa(i["directory"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_size"
	if _, ok := i["file-size"]; ok {
		result["file_size"] = flattenSystemLogSettingsRollingAnalyzerFileSizeSlsa(i["file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "gzip_format"
	if _, ok := i["gzip-format"]; ok {
		result["gzip_format"] = flattenSystemLogSettingsRollingAnalyzerGzipFormatSlsa(i["gzip-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "hour"
	if _, ok := i["hour"]; ok {
		result["hour"] = flattenSystemLogSettingsRollingAnalyzerHourSlsa(i["hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenSystemLogSettingsRollingAnalyzerIpSlsa(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip2"
	if _, ok := i["ip2"]; ok {
		result["ip2"] = flattenSystemLogSettingsRollingAnalyzerIp2Slsa(i["ip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip3"
	if _, ok := i["ip3"]; ok {
		result["ip3"] = flattenSystemLogSettingsRollingAnalyzerIp3Slsa(i["ip3"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_format"
	if _, ok := i["log-format"]; ok {
		result["log_format"] = flattenSystemLogSettingsRollingAnalyzerLogFormatSlsa(i["log-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "min"
	if _, ok := i["min"]; ok {
		result["min"] = flattenSystemLogSettingsRollingAnalyzerMinSlsa(i["min"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenSystemLogSettingsRollingAnalyzerPasswordSlsa(i["password"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password"] = c
		}
	}

	pre_append = pre + ".0." + "password2"
	if _, ok := i["password2"]; ok {
		result["password2"] = flattenSystemLogSettingsRollingAnalyzerPassword2Slsa(i["password2"], d, pre_append)
	}

	pre_append = pre + ".0." + "password3"
	if _, ok := i["password3"]; ok {
		result["password3"] = flattenSystemLogSettingsRollingAnalyzerPassword3Slsa(i["password3"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenSystemLogSettingsRollingAnalyzerPortSlsa(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2"
	if _, ok := i["port2"]; ok {
		result["port2"] = flattenSystemLogSettingsRollingAnalyzerPort2Slsa(i["port2"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3"
	if _, ok := i["port3"]; ok {
		result["port3"] = flattenSystemLogSettingsRollingAnalyzerPort3Slsa(i["port3"], d, pre_append)
	}

	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := i["rolling-upgrade-status"]; ok {
		result["rolling_upgrade_status"] = flattenSystemLogSettingsRollingAnalyzerRollingUpgradeStatusSlsa(i["rolling-upgrade-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenSystemLogSettingsRollingAnalyzerServerTypeSlsa(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload"
	if _, ok := i["upload"]; ok {
		result["upload"] = flattenSystemLogSettingsRollingAnalyzerUploadSlsa(i["upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_hour"
	if _, ok := i["upload-hour"]; ok {
		result["upload_hour"] = flattenSystemLogSettingsRollingAnalyzerUploadHourSlsa(i["upload-hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_mode"
	if _, ok := i["upload-mode"]; ok {
		result["upload_mode"] = flattenSystemLogSettingsRollingAnalyzerUploadModeSlsa(i["upload-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := i["upload-trigger"]; ok {
		result["upload_trigger"] = flattenSystemLogSettingsRollingAnalyzerUploadTriggerSlsa(i["upload-trigger"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenSystemLogSettingsRollingAnalyzerUsernameSlsa(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "username2"
	if _, ok := i["username2"]; ok {
		result["username2"] = flattenSystemLogSettingsRollingAnalyzerUsername2Slsa(i["username2"], d, pre_append)
	}

	pre_append = pre + ".0." + "username3"
	if _, ok := i["username3"]; ok {
		result["username3"] = flattenSystemLogSettingsRollingAnalyzerUsername3Slsa(i["username3"], d, pre_append)
	}

	pre_append = pre + ".0." + "when"
	if _, ok := i["when"]; ok {
		result["when"] = flattenSystemLogSettingsRollingAnalyzerWhenSlsa(i["when"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLogSettingsRollingAnalyzerDaysSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerDelFilesSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerDirectorySlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerFileSizeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerGzipFormatSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerHourSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerIpSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerIp2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerIp3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerLogFormatSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerMinSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerPasswordSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerPassword2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerPassword3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingAnalyzerPortSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerPort2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerPort3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerRollingUpgradeStatusSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerServerTypeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadHourSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadModeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUploadTriggerSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUsernameSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUsername2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerUsername3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingAnalyzerWhenSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalSlsa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := i["days"]; ok {
		result["days"] = flattenSystemLogSettingsRollingLocalDaysSlsa(i["days"], d, pre_append)
	}

	pre_append = pre + ".0." + "del_files"
	if _, ok := i["del-files"]; ok {
		result["del_files"] = flattenSystemLogSettingsRollingLocalDelFilesSlsa(i["del-files"], d, pre_append)
	}

	pre_append = pre + ".0." + "directory"
	if _, ok := i["directory"]; ok {
		result["directory"] = flattenSystemLogSettingsRollingLocalDirectorySlsa(i["directory"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_size"
	if _, ok := i["file-size"]; ok {
		result["file_size"] = flattenSystemLogSettingsRollingLocalFileSizeSlsa(i["file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "gzip_format"
	if _, ok := i["gzip-format"]; ok {
		result["gzip_format"] = flattenSystemLogSettingsRollingLocalGzipFormatSlsa(i["gzip-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "hour"
	if _, ok := i["hour"]; ok {
		result["hour"] = flattenSystemLogSettingsRollingLocalHourSlsa(i["hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenSystemLogSettingsRollingLocalIpSlsa(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip2"
	if _, ok := i["ip2"]; ok {
		result["ip2"] = flattenSystemLogSettingsRollingLocalIp2Slsa(i["ip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip3"
	if _, ok := i["ip3"]; ok {
		result["ip3"] = flattenSystemLogSettingsRollingLocalIp3Slsa(i["ip3"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_format"
	if _, ok := i["log-format"]; ok {
		result["log_format"] = flattenSystemLogSettingsRollingLocalLogFormatSlsa(i["log-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "min"
	if _, ok := i["min"]; ok {
		result["min"] = flattenSystemLogSettingsRollingLocalMinSlsa(i["min"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenSystemLogSettingsRollingLocalPasswordSlsa(i["password"], d, pre_append)
		c := d.Get(pre_append).(*schema.Set)
		if c.Len() > 0 {
			result["password"] = c
		}
	}

	pre_append = pre + ".0." + "password2"
	if _, ok := i["password2"]; ok {
		result["password2"] = flattenSystemLogSettingsRollingLocalPassword2Slsa(i["password2"], d, pre_append)
	}

	pre_append = pre + ".0." + "password3"
	if _, ok := i["password3"]; ok {
		result["password3"] = flattenSystemLogSettingsRollingLocalPassword3Slsa(i["password3"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenSystemLogSettingsRollingLocalPortSlsa(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2"
	if _, ok := i["port2"]; ok {
		result["port2"] = flattenSystemLogSettingsRollingLocalPort2Slsa(i["port2"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3"
	if _, ok := i["port3"]; ok {
		result["port3"] = flattenSystemLogSettingsRollingLocalPort3Slsa(i["port3"], d, pre_append)
	}

	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := i["rolling-upgrade-status"]; ok {
		result["rolling_upgrade_status"] = flattenSystemLogSettingsRollingLocalRollingUpgradeStatusSlsa(i["rolling-upgrade-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenSystemLogSettingsRollingLocalServerTypeSlsa(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload"
	if _, ok := i["upload"]; ok {
		result["upload"] = flattenSystemLogSettingsRollingLocalUploadSlsa(i["upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_hour"
	if _, ok := i["upload-hour"]; ok {
		result["upload_hour"] = flattenSystemLogSettingsRollingLocalUploadHourSlsa(i["upload-hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_mode"
	if _, ok := i["upload-mode"]; ok {
		result["upload_mode"] = flattenSystemLogSettingsRollingLocalUploadModeSlsa(i["upload-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := i["upload-trigger"]; ok {
		result["upload_trigger"] = flattenSystemLogSettingsRollingLocalUploadTriggerSlsa(i["upload-trigger"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenSystemLogSettingsRollingLocalUsernameSlsa(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "username2"
	if _, ok := i["username2"]; ok {
		result["username2"] = flattenSystemLogSettingsRollingLocalUsername2Slsa(i["username2"], d, pre_append)
	}

	pre_append = pre + ".0." + "username3"
	if _, ok := i["username3"]; ok {
		result["username3"] = flattenSystemLogSettingsRollingLocalUsername3Slsa(i["username3"], d, pre_append)
	}

	pre_append = pre + ".0." + "when"
	if _, ok := i["when"]; ok {
		result["when"] = flattenSystemLogSettingsRollingLocalWhenSlsa(i["when"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLogSettingsRollingLocalDaysSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalDelFilesSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalDirectorySlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalFileSizeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalGzipFormatSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalHourSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalIpSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalIp2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalIp3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalLogFormatSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalMinSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalPasswordSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalPassword2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalPassword3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingLocalPortSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalPort2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalPort3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalRollingUpgradeStatusSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalServerTypeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadHourSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadModeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUploadTriggerSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUsernameSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUsername2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalUsername3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingLocalWhenSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularSlsa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := i["days"]; ok {
		result["days"] = flattenSystemLogSettingsRollingRegularDaysSlsa(i["days"], d, pre_append)
	}

	pre_append = pre + ".0." + "del_files"
	if _, ok := i["del-files"]; ok {
		result["del_files"] = flattenSystemLogSettingsRollingRegularDelFilesSlsa(i["del-files"], d, pre_append)
	}

	pre_append = pre + ".0." + "directory"
	if _, ok := i["directory"]; ok {
		result["directory"] = flattenSystemLogSettingsRollingRegularDirectorySlsa(i["directory"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_size"
	if _, ok := i["file-size"]; ok {
		result["file_size"] = flattenSystemLogSettingsRollingRegularFileSizeSlsa(i["file-size"], d, pre_append)
	}

	pre_append = pre + ".0." + "gzip_format"
	if _, ok := i["gzip-format"]; ok {
		result["gzip_format"] = flattenSystemLogSettingsRollingRegularGzipFormatSlsa(i["gzip-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "hour"
	if _, ok := i["hour"]; ok {
		result["hour"] = flattenSystemLogSettingsRollingRegularHourSlsa(i["hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenSystemLogSettingsRollingRegularIpSlsa(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip2"
	if _, ok := i["ip2"]; ok {
		result["ip2"] = flattenSystemLogSettingsRollingRegularIp2Slsa(i["ip2"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip3"
	if _, ok := i["ip3"]; ok {
		result["ip3"] = flattenSystemLogSettingsRollingRegularIp3Slsa(i["ip3"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_format"
	if _, ok := i["log-format"]; ok {
		result["log_format"] = flattenSystemLogSettingsRollingRegularLogFormatSlsa(i["log-format"], d, pre_append)
	}

	pre_append = pre + ".0." + "min"
	if _, ok := i["min"]; ok {
		result["min"] = flattenSystemLogSettingsRollingRegularMinSlsa(i["min"], d, pre_append)
	}

	pre_append = pre + ".0." + "password"
	if _, ok := i["password"]; ok {
		result["password"] = flattenSystemLogSettingsRollingRegularPasswordSlsa(i["password"], d, pre_append)
	}

	pre_append = pre + ".0." + "password2"
	if _, ok := i["password2"]; ok {
		result["password2"] = flattenSystemLogSettingsRollingRegularPassword2Slsa(i["password2"], d, pre_append)
	}

	pre_append = pre + ".0." + "password3"
	if _, ok := i["password3"]; ok {
		result["password3"] = flattenSystemLogSettingsRollingRegularPassword3Slsa(i["password3"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenSystemLogSettingsRollingRegularPortSlsa(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2"
	if _, ok := i["port2"]; ok {
		result["port2"] = flattenSystemLogSettingsRollingRegularPort2Slsa(i["port2"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3"
	if _, ok := i["port3"]; ok {
		result["port3"] = flattenSystemLogSettingsRollingRegularPort3Slsa(i["port3"], d, pre_append)
	}

	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := i["rolling-upgrade-status"]; ok {
		result["rolling_upgrade_status"] = flattenSystemLogSettingsRollingRegularRollingUpgradeStatusSlsa(i["rolling-upgrade-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "server_type"
	if _, ok := i["server-type"]; ok {
		result["server_type"] = flattenSystemLogSettingsRollingRegularServerTypeSlsa(i["server-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload"
	if _, ok := i["upload"]; ok {
		result["upload"] = flattenSystemLogSettingsRollingRegularUploadSlsa(i["upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_hour"
	if _, ok := i["upload-hour"]; ok {
		result["upload_hour"] = flattenSystemLogSettingsRollingRegularUploadHourSlsa(i["upload-hour"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_mode"
	if _, ok := i["upload-mode"]; ok {
		result["upload_mode"] = flattenSystemLogSettingsRollingRegularUploadModeSlsa(i["upload-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := i["upload-trigger"]; ok {
		result["upload_trigger"] = flattenSystemLogSettingsRollingRegularUploadTriggerSlsa(i["upload-trigger"], d, pre_append)
	}

	pre_append = pre + ".0." + "username"
	if _, ok := i["username"]; ok {
		result["username"] = flattenSystemLogSettingsRollingRegularUsernameSlsa(i["username"], d, pre_append)
	}

	pre_append = pre + ".0." + "username2"
	if _, ok := i["username2"]; ok {
		result["username2"] = flattenSystemLogSettingsRollingRegularUsername2Slsa(i["username2"], d, pre_append)
	}

	pre_append = pre + ".0." + "username3"
	if _, ok := i["username3"]; ok {
		result["username3"] = flattenSystemLogSettingsRollingRegularUsername3Slsa(i["username3"], d, pre_append)
	}

	pre_append = pre + ".0." + "when"
	if _, ok := i["when"]; ok {
		result["when"] = flattenSystemLogSettingsRollingRegularWhenSlsa(i["when"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLogSettingsRollingRegularDaysSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularDelFilesSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularDirectorySlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularFileSizeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularGzipFormatSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularHourSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularIpSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularIp2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularIp3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularLogFormatSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularMinSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularPasswordSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularPassword2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularPassword3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLogSettingsRollingRegularPortSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularPort2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularPort3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularRollingUpgradeStatusSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularServerTypeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadHourSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadModeSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUploadTriggerSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUsernameSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUsername2Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularUsername3Slsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsRollingRegularWhenSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogSettingsSyncSearchTimeoutSlsa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fac_custom_field1", flattenSystemLogSettingsFacCustomField1Slsa(o["FAC-custom-field1"], d, "fac_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FAC-custom-field1"], "SystemLogSettings-FacCustomField1"); ok {
			if err = d.Set("fac_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fac_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fac_custom_field1: %v", err)
		}
	}

	if err = d.Set("faz_custom_field1", flattenSystemLogSettingsFazCustomField1Slsa(o["FAZ-custom-field1"], d, "faz_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FAZ-custom-field1"], "SystemLogSettings-FazCustomField1"); ok {
			if err = d.Set("faz_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading faz_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_custom_field1: %v", err)
		}
	}

	if err = d.Set("fch_custom_field1", flattenSystemLogSettingsFchCustomField1Slsa(o["FCH-custom-field1"], d, "fch_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FCH-custom-field1"], "SystemLogSettings-FchCustomField1"); ok {
			if err = d.Set("fch_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fch_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fch_custom_field1: %v", err)
		}
	}

	if err = d.Set("fct_custom_field1", flattenSystemLogSettingsFctCustomField1Slsa(o["FCT-custom-field1"], d, "fct_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FCT-custom-field1"], "SystemLogSettings-FctCustomField1"); ok {
			if err = d.Set("fct_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fct_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fct_custom_field1: %v", err)
		}
	}

	if err = d.Set("fdd_custom_field1", flattenSystemLogSettingsFddCustomField1Slsa(o["FDD-custom-field1"], d, "fdd_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FDD-custom-field1"], "SystemLogSettings-FddCustomField1"); ok {
			if err = d.Set("fdd_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fdd_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fdd_custom_field1: %v", err)
		}
	}

	if err = d.Set("fgt_custom_field1", flattenSystemLogSettingsFgtCustomField1Slsa(o["FGT-custom-field1"], d, "fgt_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FGT-custom-field1"], "SystemLogSettings-FgtCustomField1"); ok {
			if err = d.Set("fgt_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fgt_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_custom_field1: %v", err)
		}
	}

	if err = d.Set("fmg_custom_field1", flattenSystemLogSettingsFmgCustomField1Slsa(o["FMG-custom-field1"], d, "fmg_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FMG-custom-field1"], "SystemLogSettings-FmgCustomField1"); ok {
			if err = d.Set("fmg_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fmg_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmg_custom_field1: %v", err)
		}
	}

	if err = d.Set("fml_custom_field1", flattenSystemLogSettingsFmlCustomField1Slsa(o["FML-custom-field1"], d, "fml_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FML-custom-field1"], "SystemLogSettings-FmlCustomField1"); ok {
			if err = d.Set("fml_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fml_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fml_custom_field1: %v", err)
		}
	}

	if err = d.Set("fpx_custom_field1", flattenSystemLogSettingsFpxCustomField1Slsa(o["FPX-custom-field1"], d, "fpx_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FPX-custom-field1"], "SystemLogSettings-FpxCustomField1"); ok {
			if err = d.Set("fpx_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fpx_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fpx_custom_field1: %v", err)
		}
	}

	if err = d.Set("fsa_custom_field1", flattenSystemLogSettingsFsaCustomField1Slsa(o["FSA-custom-field1"], d, "fsa_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FSA-custom-field1"], "SystemLogSettings-FsaCustomField1"); ok {
			if err = d.Set("fsa_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fsa_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsa_custom_field1: %v", err)
		}
	}

	if err = d.Set("fwb_custom_field1", flattenSystemLogSettingsFwbCustomField1Slsa(o["FWB-custom-field1"], d, "fwb_custom_field1")); err != nil {
		if vv, ok := fortiAPIPatch(o["FWB-custom-field1"], "SystemLogSettings-FwbCustomField1"); ok {
			if err = d.Set("fwb_custom_field1", vv); err != nil {
				return fmt.Errorf("Error reading fwb_custom_field1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwb_custom_field1: %v", err)
		}
	}

	if err = d.Set("browse_max_logfiles", flattenSystemLogSettingsBrowseMaxLogfilesSlsa(o["browse-max-logfiles"], d, "browse_max_logfiles")); err != nil {
		if vv, ok := fortiAPIPatch(o["browse-max-logfiles"], "SystemLogSettings-BrowseMaxLogfiles"); ok {
			if err = d.Set("browse_max_logfiles", vv); err != nil {
				return fmt.Errorf("Error reading browse_max_logfiles: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading browse_max_logfiles: %v", err)
		}
	}

	if err = d.Set("dns_resolve_dstip", flattenSystemLogSettingsDnsResolveDstipSlsa(o["dns-resolve-dstip"], d, "dns_resolve_dstip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-resolve-dstip"], "SystemLogSettings-DnsResolveDstip"); ok {
			if err = d.Set("dns_resolve_dstip", vv); err != nil {
				return fmt.Errorf("Error reading dns_resolve_dstip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_resolve_dstip: %v", err)
		}
	}

	if err = d.Set("download_max_logs", flattenSystemLogSettingsDownloadMaxLogsSlsa(o["download-max-logs"], d, "download_max_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["download-max-logs"], "SystemLogSettings-DownloadMaxLogs"); ok {
			if err = d.Set("download_max_logs", vv); err != nil {
				return fmt.Errorf("Error reading download_max_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading download_max_logs: %v", err)
		}
	}

	if err = d.Set("ha_auto_migrate", flattenSystemLogSettingsHaAutoMigrateSlsa(o["ha-auto-migrate"], d, "ha_auto_migrate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-auto-migrate"], "SystemLogSettings-HaAutoMigrate"); ok {
			if err = d.Set("ha_auto_migrate", vv); err != nil {
				return fmt.Errorf("Error reading ha_auto_migrate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_auto_migrate: %v", err)
		}
	}

	if err = d.Set("import_max_logfiles", flattenSystemLogSettingsImportMaxLogfilesSlsa(o["import-max-logfiles"], d, "import_max_logfiles")); err != nil {
		if vv, ok := fortiAPIPatch(o["import-max-logfiles"], "SystemLogSettings-ImportMaxLogfiles"); ok {
			if err = d.Set("import_max_logfiles", vv); err != nil {
				return fmt.Errorf("Error reading import_max_logfiles: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading import_max_logfiles: %v", err)
		}
	}

	if err = d.Set("keep_dev_logs", flattenSystemLogSettingsKeepDevLogsSlsa(o["keep-dev-logs"], d, "keep_dev_logs")); err != nil {
		if vv, ok := fortiAPIPatch(o["keep-dev-logs"], "SystemLogSettings-KeepDevLogs"); ok {
			if err = d.Set("keep_dev_logs", vv); err != nil {
				return fmt.Errorf("Error reading keep_dev_logs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keep_dev_logs: %v", err)
		}
	}

	if err = d.Set("log_file_archive_name", flattenSystemLogSettingsLogFileArchiveNameSlsa(o["log-file-archive-name"], d, "log_file_archive_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-file-archive-name"], "SystemLogSettings-LogFileArchiveName"); ok {
			if err = d.Set("log_file_archive_name", vv); err != nil {
				return fmt.Errorf("Error reading log_file_archive_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_file_archive_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rolling_analyzer", flattenSystemLogSettingsRollingAnalyzerSlsa(o["rolling-analyzer"], d, "rolling_analyzer")); err != nil {
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
			if err = d.Set("rolling_analyzer", flattenSystemLogSettingsRollingAnalyzerSlsa(o["rolling-analyzer"], d, "rolling_analyzer")); err != nil {
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
		if err = d.Set("rolling_local", flattenSystemLogSettingsRollingLocalSlsa(o["rolling-local"], d, "rolling_local")); err != nil {
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
			if err = d.Set("rolling_local", flattenSystemLogSettingsRollingLocalSlsa(o["rolling-local"], d, "rolling_local")); err != nil {
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
		if err = d.Set("rolling_regular", flattenSystemLogSettingsRollingRegularSlsa(o["rolling-regular"], d, "rolling_regular")); err != nil {
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
			if err = d.Set("rolling_regular", flattenSystemLogSettingsRollingRegularSlsa(o["rolling-regular"], d, "rolling_regular")); err != nil {
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

	if err = d.Set("sync_search_timeout", flattenSystemLogSettingsSyncSearchTimeoutSlsa(o["sync-search-timeout"], d, "sync_search_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["sync-search-timeout"], "SystemLogSettings-SyncSearchTimeout"); ok {
			if err = d.Set("sync_search_timeout", vv); err != nil {
				return fmt.Errorf("Error reading sync_search_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sync_search_timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemLogSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogSettingsFacCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFazCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFchCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFctCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFddCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFgtCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFmgCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFmlCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFpxCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFsaCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsFwbCustomField1Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsBrowseMaxLogfilesSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsDnsResolveDstipSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsDownloadMaxLogsSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsHaAutoMigrateSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsImportMaxLogfilesSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsKeepDevLogsSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsLogFileArchiveNameSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days"], _ = expandSystemLogSettingsRollingAnalyzerDaysSlsa(d, i["days"], pre_append)
	} else {
		result["days"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "del_files"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["del-files"], _ = expandSystemLogSettingsRollingAnalyzerDelFilesSlsa(d, i["del_files"], pre_append)
	}
	pre_append = pre + ".0." + "directory"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["directory"], _ = expandSystemLogSettingsRollingAnalyzerDirectorySlsa(d, i["directory"], pre_append)
	}
	pre_append = pre + ".0." + "file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-size"], _ = expandSystemLogSettingsRollingAnalyzerFileSizeSlsa(d, i["file_size"], pre_append)
	}
	pre_append = pre + ".0." + "gzip_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gzip-format"], _ = expandSystemLogSettingsRollingAnalyzerGzipFormatSlsa(d, i["gzip_format"], pre_append)
	}
	pre_append = pre + ".0." + "hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hour"], _ = expandSystemLogSettingsRollingAnalyzerHourSlsa(d, i["hour"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandSystemLogSettingsRollingAnalyzerIpSlsa(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip2"], _ = expandSystemLogSettingsRollingAnalyzerIp2Slsa(d, i["ip2"], pre_append)
	}
	pre_append = pre + ".0." + "ip3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip3"], _ = expandSystemLogSettingsRollingAnalyzerIp3Slsa(d, i["ip3"], pre_append)
	}
	pre_append = pre + ".0." + "log_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-format"], _ = expandSystemLogSettingsRollingAnalyzerLogFormatSlsa(d, i["log_format"], pre_append)
	}
	pre_append = pre + ".0." + "min"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min"], _ = expandSystemLogSettingsRollingAnalyzerMinSlsa(d, i["min"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandSystemLogSettingsRollingAnalyzerPasswordSlsa(d, i["password"], pre_append)
	} else {
		result["password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "password2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password2"], _ = expandSystemLogSettingsRollingAnalyzerPassword2Slsa(d, i["password2"], pre_append)
	} else {
		result["password2"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "password3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password3"], _ = expandSystemLogSettingsRollingAnalyzerPassword3Slsa(d, i["password3"], pre_append)
	} else {
		result["password3"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandSystemLogSettingsRollingAnalyzerPortSlsa(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "port2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2"], _ = expandSystemLogSettingsRollingAnalyzerPort2Slsa(d, i["port2"], pre_append)
	}
	pre_append = pre + ".0." + "port3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3"], _ = expandSystemLogSettingsRollingAnalyzerPort3Slsa(d, i["port3"], pre_append)
	}
	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rolling-upgrade-status"], _ = expandSystemLogSettingsRollingAnalyzerRollingUpgradeStatusSlsa(d, i["rolling_upgrade_status"], pre_append)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandSystemLogSettingsRollingAnalyzerServerTypeSlsa(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "upload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload"], _ = expandSystemLogSettingsRollingAnalyzerUploadSlsa(d, i["upload"], pre_append)
	}
	pre_append = pre + ".0." + "upload_hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-hour"], _ = expandSystemLogSettingsRollingAnalyzerUploadHourSlsa(d, i["upload_hour"], pre_append)
	}
	pre_append = pre + ".0." + "upload_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-mode"], _ = expandSystemLogSettingsRollingAnalyzerUploadModeSlsa(d, i["upload_mode"], pre_append)
	}
	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-trigger"], _ = expandSystemLogSettingsRollingAnalyzerUploadTriggerSlsa(d, i["upload_trigger"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandSystemLogSettingsRollingAnalyzerUsernameSlsa(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "username2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username2"], _ = expandSystemLogSettingsRollingAnalyzerUsername2Slsa(d, i["username2"], pre_append)
	}
	pre_append = pre + ".0." + "username3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username3"], _ = expandSystemLogSettingsRollingAnalyzerUsername3Slsa(d, i["username3"], pre_append)
	}
	pre_append = pre + ".0." + "when"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["when"], _ = expandSystemLogSettingsRollingAnalyzerWhenSlsa(d, i["when"], pre_append)
	}

	return result, nil
}

func expandSystemLogSettingsRollingAnalyzerDaysSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerDelFilesSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerDirectorySlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerFileSizeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerGzipFormatSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerHourSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerIpSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerIp2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerIp3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerLogFormatSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerMinSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerPasswordSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerPassword2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerPassword3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingAnalyzerPortSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerPort2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerPort3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerRollingUpgradeStatusSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerServerTypeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadHourSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadModeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUploadTriggerSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUsernameSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUsername2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerUsername3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingAnalyzerWhenSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days"], _ = expandSystemLogSettingsRollingLocalDaysSlsa(d, i["days"], pre_append)
	} else {
		result["days"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "del_files"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["del-files"], _ = expandSystemLogSettingsRollingLocalDelFilesSlsa(d, i["del_files"], pre_append)
	}
	pre_append = pre + ".0." + "directory"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["directory"], _ = expandSystemLogSettingsRollingLocalDirectorySlsa(d, i["directory"], pre_append)
	}
	pre_append = pre + ".0." + "file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-size"], _ = expandSystemLogSettingsRollingLocalFileSizeSlsa(d, i["file_size"], pre_append)
	}
	pre_append = pre + ".0." + "gzip_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gzip-format"], _ = expandSystemLogSettingsRollingLocalGzipFormatSlsa(d, i["gzip_format"], pre_append)
	}
	pre_append = pre + ".0." + "hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hour"], _ = expandSystemLogSettingsRollingLocalHourSlsa(d, i["hour"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandSystemLogSettingsRollingLocalIpSlsa(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip2"], _ = expandSystemLogSettingsRollingLocalIp2Slsa(d, i["ip2"], pre_append)
	}
	pre_append = pre + ".0." + "ip3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip3"], _ = expandSystemLogSettingsRollingLocalIp3Slsa(d, i["ip3"], pre_append)
	}
	pre_append = pre + ".0." + "log_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-format"], _ = expandSystemLogSettingsRollingLocalLogFormatSlsa(d, i["log_format"], pre_append)
	}
	pre_append = pre + ".0." + "min"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min"], _ = expandSystemLogSettingsRollingLocalMinSlsa(d, i["min"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandSystemLogSettingsRollingLocalPasswordSlsa(d, i["password"], pre_append)
	} else {
		result["password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "password2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password2"], _ = expandSystemLogSettingsRollingLocalPassword2Slsa(d, i["password2"], pre_append)
	} else {
		result["password2"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "password3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password3"], _ = expandSystemLogSettingsRollingLocalPassword3Slsa(d, i["password3"], pre_append)
	} else {
		result["password3"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandSystemLogSettingsRollingLocalPortSlsa(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "port2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2"], _ = expandSystemLogSettingsRollingLocalPort2Slsa(d, i["port2"], pre_append)
	}
	pre_append = pre + ".0." + "port3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3"], _ = expandSystemLogSettingsRollingLocalPort3Slsa(d, i["port3"], pre_append)
	}
	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rolling-upgrade-status"], _ = expandSystemLogSettingsRollingLocalRollingUpgradeStatusSlsa(d, i["rolling_upgrade_status"], pre_append)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandSystemLogSettingsRollingLocalServerTypeSlsa(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "upload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload"], _ = expandSystemLogSettingsRollingLocalUploadSlsa(d, i["upload"], pre_append)
	}
	pre_append = pre + ".0." + "upload_hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-hour"], _ = expandSystemLogSettingsRollingLocalUploadHourSlsa(d, i["upload_hour"], pre_append)
	}
	pre_append = pre + ".0." + "upload_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-mode"], _ = expandSystemLogSettingsRollingLocalUploadModeSlsa(d, i["upload_mode"], pre_append)
	}
	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-trigger"], _ = expandSystemLogSettingsRollingLocalUploadTriggerSlsa(d, i["upload_trigger"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandSystemLogSettingsRollingLocalUsernameSlsa(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "username2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username2"], _ = expandSystemLogSettingsRollingLocalUsername2Slsa(d, i["username2"], pre_append)
	}
	pre_append = pre + ".0." + "username3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username3"], _ = expandSystemLogSettingsRollingLocalUsername3Slsa(d, i["username3"], pre_append)
	}
	pre_append = pre + ".0." + "when"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["when"], _ = expandSystemLogSettingsRollingLocalWhenSlsa(d, i["when"], pre_append)
	}

	return result, nil
}

func expandSystemLogSettingsRollingLocalDaysSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalDelFilesSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalDirectorySlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalFileSizeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalGzipFormatSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalHourSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalIpSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalIp2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalIp3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalLogFormatSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalMinSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalPasswordSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalPassword2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalPassword3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingLocalPortSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalPort2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalPort3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalRollingUpgradeStatusSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalServerTypeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadHourSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadModeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUploadTriggerSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUsernameSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUsername2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalUsername3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingLocalWhenSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "days"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["days"], _ = expandSystemLogSettingsRollingRegularDaysSlsa(d, i["days"], pre_append)
	} else {
		result["days"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "del_files"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["del-files"], _ = expandSystemLogSettingsRollingRegularDelFilesSlsa(d, i["del_files"], pre_append)
	}
	pre_append = pre + ".0." + "directory"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["directory"], _ = expandSystemLogSettingsRollingRegularDirectorySlsa(d, i["directory"], pre_append)
	}
	pre_append = pre + ".0." + "file_size"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-size"], _ = expandSystemLogSettingsRollingRegularFileSizeSlsa(d, i["file_size"], pre_append)
	}
	pre_append = pre + ".0." + "gzip_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gzip-format"], _ = expandSystemLogSettingsRollingRegularGzipFormatSlsa(d, i["gzip_format"], pre_append)
	}
	pre_append = pre + ".0." + "hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["hour"], _ = expandSystemLogSettingsRollingRegularHourSlsa(d, i["hour"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandSystemLogSettingsRollingRegularIpSlsa(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "ip2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip2"], _ = expandSystemLogSettingsRollingRegularIp2Slsa(d, i["ip2"], pre_append)
	}
	pre_append = pre + ".0." + "ip3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip3"], _ = expandSystemLogSettingsRollingRegularIp3Slsa(d, i["ip3"], pre_append)
	}
	pre_append = pre + ".0." + "log_format"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log-format"], _ = expandSystemLogSettingsRollingRegularLogFormatSlsa(d, i["log_format"], pre_append)
	}
	pre_append = pre + ".0." + "min"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["min"], _ = expandSystemLogSettingsRollingRegularMinSlsa(d, i["min"], pre_append)
	}
	pre_append = pre + ".0." + "password"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password"], _ = expandSystemLogSettingsRollingRegularPasswordSlsa(d, i["password"], pre_append)
	} else {
		result["password"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "password2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password2"], _ = expandSystemLogSettingsRollingRegularPassword2Slsa(d, i["password2"], pre_append)
	} else {
		result["password2"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "password3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["password3"], _ = expandSystemLogSettingsRollingRegularPassword3Slsa(d, i["password3"], pre_append)
	} else {
		result["password3"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port"], _ = expandSystemLogSettingsRollingRegularPortSlsa(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "port2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port2"], _ = expandSystemLogSettingsRollingRegularPort2Slsa(d, i["port2"], pre_append)
	}
	pre_append = pre + ".0." + "port3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["port3"], _ = expandSystemLogSettingsRollingRegularPort3Slsa(d, i["port3"], pre_append)
	}
	pre_append = pre + ".0." + "rolling_upgrade_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rolling-upgrade-status"], _ = expandSystemLogSettingsRollingRegularRollingUpgradeStatusSlsa(d, i["rolling_upgrade_status"], pre_append)
	}
	pre_append = pre + ".0." + "server_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["server-type"], _ = expandSystemLogSettingsRollingRegularServerTypeSlsa(d, i["server_type"], pre_append)
	}
	pre_append = pre + ".0." + "upload"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload"], _ = expandSystemLogSettingsRollingRegularUploadSlsa(d, i["upload"], pre_append)
	}
	pre_append = pre + ".0." + "upload_hour"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-hour"], _ = expandSystemLogSettingsRollingRegularUploadHourSlsa(d, i["upload_hour"], pre_append)
	}
	pre_append = pre + ".0." + "upload_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-mode"], _ = expandSystemLogSettingsRollingRegularUploadModeSlsa(d, i["upload_mode"], pre_append)
	}
	pre_append = pre + ".0." + "upload_trigger"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upload-trigger"], _ = expandSystemLogSettingsRollingRegularUploadTriggerSlsa(d, i["upload_trigger"], pre_append)
	}
	pre_append = pre + ".0." + "username"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username"], _ = expandSystemLogSettingsRollingRegularUsernameSlsa(d, i["username"], pre_append)
	}
	pre_append = pre + ".0." + "username2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username2"], _ = expandSystemLogSettingsRollingRegularUsername2Slsa(d, i["username2"], pre_append)
	}
	pre_append = pre + ".0." + "username3"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["username3"], _ = expandSystemLogSettingsRollingRegularUsername3Slsa(d, i["username3"], pre_append)
	}
	pre_append = pre + ".0." + "when"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["when"], _ = expandSystemLogSettingsRollingRegularWhenSlsa(d, i["when"], pre_append)
	}

	return result, nil
}

func expandSystemLogSettingsRollingRegularDaysSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularDelFilesSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularDirectorySlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularFileSizeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularGzipFormatSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularHourSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularIpSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularIp2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularIp3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularLogFormatSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularMinSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularPasswordSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularPassword2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularPassword3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLogSettingsRollingRegularPortSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularPort2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularPort3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularRollingUpgradeStatusSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularServerTypeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadHourSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadModeSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUploadTriggerSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUsernameSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUsername2Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularUsername3Slsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsRollingRegularWhenSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogSettingsSyncSearchTimeoutSlsa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fac_custom_field1"); ok || d.HasChange("fac_custom_field1") {
		t, err := expandSystemLogSettingsFacCustomField1Slsa(d, v, "fac_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FAC-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("faz_custom_field1"); ok || d.HasChange("faz_custom_field1") {
		t, err := expandSystemLogSettingsFazCustomField1Slsa(d, v, "faz_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FAZ-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fch_custom_field1"); ok || d.HasChange("fch_custom_field1") {
		t, err := expandSystemLogSettingsFchCustomField1Slsa(d, v, "fch_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FCH-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fct_custom_field1"); ok || d.HasChange("fct_custom_field1") {
		t, err := expandSystemLogSettingsFctCustomField1Slsa(d, v, "fct_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FCT-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fdd_custom_field1"); ok || d.HasChange("fdd_custom_field1") {
		t, err := expandSystemLogSettingsFddCustomField1Slsa(d, v, "fdd_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FDD-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fgt_custom_field1"); ok || d.HasChange("fgt_custom_field1") {
		t, err := expandSystemLogSettingsFgtCustomField1Slsa(d, v, "fgt_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FGT-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fmg_custom_field1"); ok || d.HasChange("fmg_custom_field1") {
		t, err := expandSystemLogSettingsFmgCustomField1Slsa(d, v, "fmg_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FMG-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fml_custom_field1"); ok || d.HasChange("fml_custom_field1") {
		t, err := expandSystemLogSettingsFmlCustomField1Slsa(d, v, "fml_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FML-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fpx_custom_field1"); ok || d.HasChange("fpx_custom_field1") {
		t, err := expandSystemLogSettingsFpxCustomField1Slsa(d, v, "fpx_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FPX-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fsa_custom_field1"); ok || d.HasChange("fsa_custom_field1") {
		t, err := expandSystemLogSettingsFsaCustomField1Slsa(d, v, "fsa_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FSA-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("fwb_custom_field1"); ok || d.HasChange("fwb_custom_field1") {
		t, err := expandSystemLogSettingsFwbCustomField1Slsa(d, v, "fwb_custom_field1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["FWB-custom-field1"] = t
		}
	}

	if v, ok := d.GetOk("browse_max_logfiles"); ok || d.HasChange("browse_max_logfiles") {
		t, err := expandSystemLogSettingsBrowseMaxLogfilesSlsa(d, v, "browse_max_logfiles")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["browse-max-logfiles"] = t
		}
	}

	if v, ok := d.GetOk("dns_resolve_dstip"); ok || d.HasChange("dns_resolve_dstip") {
		t, err := expandSystemLogSettingsDnsResolveDstipSlsa(d, v, "dns_resolve_dstip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-resolve-dstip"] = t
		}
	}

	if v, ok := d.GetOk("download_max_logs"); ok || d.HasChange("download_max_logs") {
		t, err := expandSystemLogSettingsDownloadMaxLogsSlsa(d, v, "download_max_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["download-max-logs"] = t
		}
	}

	if v, ok := d.GetOk("ha_auto_migrate"); ok || d.HasChange("ha_auto_migrate") {
		t, err := expandSystemLogSettingsHaAutoMigrateSlsa(d, v, "ha_auto_migrate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-auto-migrate"] = t
		}
	}

	if v, ok := d.GetOk("import_max_logfiles"); ok || d.HasChange("import_max_logfiles") {
		t, err := expandSystemLogSettingsImportMaxLogfilesSlsa(d, v, "import_max_logfiles")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["import-max-logfiles"] = t
		}
	}

	if v, ok := d.GetOk("keep_dev_logs"); ok || d.HasChange("keep_dev_logs") {
		t, err := expandSystemLogSettingsKeepDevLogsSlsa(d, v, "keep_dev_logs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-dev-logs"] = t
		}
	}

	if v, ok := d.GetOk("log_file_archive_name"); ok || d.HasChange("log_file_archive_name") {
		t, err := expandSystemLogSettingsLogFileArchiveNameSlsa(d, v, "log_file_archive_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-file-archive-name"] = t
		}
	}

	if v, ok := d.GetOk("rolling_analyzer"); ok || d.HasChange("rolling_analyzer") {
		t, err := expandSystemLogSettingsRollingAnalyzerSlsa(d, v, "rolling_analyzer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-analyzer"] = t
		}
	}

	if v, ok := d.GetOk("rolling_local"); ok || d.HasChange("rolling_local") {
		t, err := expandSystemLogSettingsRollingLocalSlsa(d, v, "rolling_local")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-local"] = t
		}
	}

	if v, ok := d.GetOk("rolling_regular"); ok || d.HasChange("rolling_regular") {
		t, err := expandSystemLogSettingsRollingRegularSlsa(d, v, "rolling_regular")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rolling-regular"] = t
		}
	}

	if v, ok := d.GetOk("sync_search_timeout"); ok || d.HasChange("sync_search_timeout") {
		t, err := expandSystemLogSettingsSyncSearchTimeoutSlsa(d, v, "sync_search_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-search-timeout"] = t
		}
	}

	return &obj, nil
}
