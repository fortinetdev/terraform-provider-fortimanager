// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for locallog logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocallogSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogSettingUpdate,
		Read:   resourceSystemLocallogSettingRead,
		Update: resourceSystemLocallogSettingUpdate,
		Delete: resourceSystemLocallogSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"log_daemon_crash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_interval_adom_perf_stats": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_interval_dev_no_logging": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_interval_disk_full": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_interval_gbday_exceeded": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLocallogSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogSetting")

	return resourceSystemLocallogSettingRead(d, m)
}

func resourceSystemLocallogSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteSystemLocallogSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLocallogSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogSettingLogDaemonCrash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSettingLogIntervalAdomPerfStats(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSettingLogIntervalDevNoLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSettingLogIntervalDiskFull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSettingLogIntervalGbdayExceeded(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("log_daemon_crash", flattenSystemLocallogSettingLogDaemonCrash(o["log-daemon-crash"], d, "log_daemon_crash")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-daemon-crash"], "SystemLocallogSetting-LogDaemonCrash"); ok {
			if err = d.Set("log_daemon_crash", vv); err != nil {
				return fmt.Errorf("Error reading log_daemon_crash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_daemon_crash: %v", err)
		}
	}

	if err = d.Set("log_interval_adom_perf_stats", flattenSystemLocallogSettingLogIntervalAdomPerfStats(o["log-interval-adom-perf-stats"], d, "log_interval_adom_perf_stats")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-interval-adom-perf-stats"], "SystemLocallogSetting-LogIntervalAdomPerfStats"); ok {
			if err = d.Set("log_interval_adom_perf_stats", vv); err != nil {
				return fmt.Errorf("Error reading log_interval_adom_perf_stats: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_interval_adom_perf_stats: %v", err)
		}
	}

	if err = d.Set("log_interval_dev_no_logging", flattenSystemLocallogSettingLogIntervalDevNoLogging(o["log-interval-dev-no-logging"], d, "log_interval_dev_no_logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-interval-dev-no-logging"], "SystemLocallogSetting-LogIntervalDevNoLogging"); ok {
			if err = d.Set("log_interval_dev_no_logging", vv); err != nil {
				return fmt.Errorf("Error reading log_interval_dev_no_logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_interval_dev_no_logging: %v", err)
		}
	}

	if err = d.Set("log_interval_disk_full", flattenSystemLocallogSettingLogIntervalDiskFull(o["log-interval-disk-full"], d, "log_interval_disk_full")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-interval-disk-full"], "SystemLocallogSetting-LogIntervalDiskFull"); ok {
			if err = d.Set("log_interval_disk_full", vv); err != nil {
				return fmt.Errorf("Error reading log_interval_disk_full: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_interval_disk_full: %v", err)
		}
	}

	if err = d.Set("log_interval_gbday_exceeded", flattenSystemLocallogSettingLogIntervalGbdayExceeded(o["log-interval-gbday-exceeded"], d, "log_interval_gbday_exceeded")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-interval-gbday-exceeded"], "SystemLocallogSetting-LogIntervalGbdayExceeded"); ok {
			if err = d.Set("log_interval_gbday_exceeded", vv); err != nil {
				return fmt.Errorf("Error reading log_interval_gbday_exceeded: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_interval_gbday_exceeded: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogSettingLogDaemonCrash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSettingLogIntervalAdomPerfStats(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSettingLogIntervalDevNoLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSettingLogIntervalDiskFull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSettingLogIntervalGbdayExceeded(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_daemon_crash"); ok || d.HasChange("log_daemon_crash") {
		t, err := expandSystemLocallogSettingLogDaemonCrash(d, v, "log_daemon_crash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-daemon-crash"] = t
		}
	}

	if v, ok := d.GetOk("log_interval_adom_perf_stats"); ok || d.HasChange("log_interval_adom_perf_stats") {
		t, err := expandSystemLocallogSettingLogIntervalAdomPerfStats(d, v, "log_interval_adom_perf_stats")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-interval-adom-perf-stats"] = t
		}
	}

	if v, ok := d.GetOk("log_interval_dev_no_logging"); ok || d.HasChange("log_interval_dev_no_logging") {
		t, err := expandSystemLocallogSettingLogIntervalDevNoLogging(d, v, "log_interval_dev_no_logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-interval-dev-no-logging"] = t
		}
	}

	if v, ok := d.GetOk("log_interval_disk_full"); ok || d.HasChange("log_interval_disk_full") {
		t, err := expandSystemLocallogSettingLogIntervalDiskFull(d, v, "log_interval_disk_full")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-interval-disk-full"] = t
		}
	}

	if v, ok := d.GetOk("log_interval_gbday_exceeded"); ok || d.HasChange("log_interval_gbday_exceeded") {
		t, err := expandSystemLocallogSettingLogIntervalGbdayExceeded(d, v, "log_interval_gbday_exceeded")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-interval-gbday-exceeded"] = t
		}
	}

	return &obj, nil
}
