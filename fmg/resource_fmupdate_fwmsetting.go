// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure firmware management settings.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFwmSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFwmSettingUpdate,
		Read:   resourceFmupdateFwmSettingRead,
		Update: resourceFmupdateFwmSettingUpdate,
		Delete: resourceFmupdateFwmSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auto_scan_fgt_disk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_fgt_disk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_failover_fmg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fds_image_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"immx_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multiple_steps_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"upgrade_timeout": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"check_status_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ctrl_check_status_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ctrl_put_image_by_fds_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ha_sync_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"license_check_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prepare_image_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"put_image_by_fds_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"put_image_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"reboot_of_fsck_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"reboot_of_upgrade_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"retrieve_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"rpc_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"total_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFmupdateFwmSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFwmSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFwmSetting(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFwmSetting")

	return resourceFmupdateFwmSettingRead(d, m)
}

func resourceFmupdateFwmSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteFmupdateFwmSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFwmSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFwmSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFwmSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFwmSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFwmSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFwmSetting resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFwmSettingAutoScanFgtDiskFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingCheckFgtDiskFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingFdsFailoverFmgFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingFdsImageTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingImmxSourceFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingLogFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingMultipleStepsIntervalFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "check_status_timeout"
	if _, ok := i["check-status-timeout"]; ok {
		result["check_status_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfa(i["check-status-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ctrl_check_status_timeout"
	if _, ok := i["ctrl-check-status-timeout"]; ok {
		result["ctrl_check_status_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfa(i["ctrl-check-status-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ctrl_put_image_by_fds_timeout"
	if _, ok := i["ctrl-put-image-by-fds-timeout"]; ok {
		result["ctrl_put_image_by_fds_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfa(i["ctrl-put-image-by-fds-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "ha_sync_timeout"
	if _, ok := i["ha-sync-timeout"]; ok {
		result["ha_sync_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfa(i["ha-sync-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "license_check_timeout"
	if _, ok := i["license-check-timeout"]; ok {
		result["license_check_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfa(i["license-check-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "prepare_image_timeout"
	if _, ok := i["prepare-image-timeout"]; ok {
		result["prepare_image_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfa(i["prepare-image-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "put_image_by_fds_timeout"
	if _, ok := i["put-image-by-fds-timeout"]; ok {
		result["put_image_by_fds_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfa(i["put-image-by-fds-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "put_image_timeout"
	if _, ok := i["put-image-timeout"]; ok {
		result["put_image_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfa(i["put-image-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "reboot_of_fsck_timeout"
	if _, ok := i["reboot-of-fsck-timeout"]; ok {
		result["reboot_of_fsck_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfa(i["reboot-of-fsck-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "reboot_of_upgrade_timeout"
	if _, ok := i["reboot-of-upgrade-timeout"]; ok {
		result["reboot_of_upgrade_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfa(i["reboot-of-upgrade-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "retrieve_timeout"
	if _, ok := i["retrieve-timeout"]; ok {
		result["retrieve_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfa(i["retrieve-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "rpc_timeout"
	if _, ok := i["rpc-timeout"]; ok {
		result["rpc_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfa(i["rpc-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "total_timeout"
	if _, ok := i["total-timeout"]; ok {
		result["total_timeout"] = flattenFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfa(i["total-timeout"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFwmSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_scan_fgt_disk", flattenFmupdateFwmSettingAutoScanFgtDiskFfa(o["auto-scan-fgt-disk"], d, "auto_scan_fgt_disk")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-scan-fgt-disk"], "FmupdateFwmSetting-AutoScanFgtDisk"); ok {
			if err = d.Set("auto_scan_fgt_disk", vv); err != nil {
				return fmt.Errorf("Error reading auto_scan_fgt_disk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_scan_fgt_disk: %v", err)
		}
	}

	if err = d.Set("check_fgt_disk", flattenFmupdateFwmSettingCheckFgtDiskFfa(o["check-fgt-disk"], d, "check_fgt_disk")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-fgt-disk"], "FmupdateFwmSetting-CheckFgtDisk"); ok {
			if err = d.Set("check_fgt_disk", vv); err != nil {
				return fmt.Errorf("Error reading check_fgt_disk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_fgt_disk: %v", err)
		}
	}

	if err = d.Set("fds_failover_fmg", flattenFmupdateFwmSettingFdsFailoverFmgFfa(o["fds-failover-fmg"], d, "fds_failover_fmg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-failover-fmg"], "FmupdateFwmSetting-FdsFailoverFmg"); ok {
			if err = d.Set("fds_failover_fmg", vv); err != nil {
				return fmt.Errorf("Error reading fds_failover_fmg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_failover_fmg: %v", err)
		}
	}

	if err = d.Set("fds_image_timeout", flattenFmupdateFwmSettingFdsImageTimeoutFfa(o["fds-image-timeout"], d, "fds_image_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-image-timeout"], "FmupdateFwmSetting-FdsImageTimeout"); ok {
			if err = d.Set("fds_image_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fds_image_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_image_timeout: %v", err)
		}
	}

	if err = d.Set("immx_source", flattenFmupdateFwmSettingImmxSourceFfa(o["immx-source"], d, "immx_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["immx-source"], "FmupdateFwmSetting-ImmxSource"); ok {
			if err = d.Set("immx_source", vv); err != nil {
				return fmt.Errorf("Error reading immx_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading immx_source: %v", err)
		}
	}

	if err = d.Set("log", flattenFmupdateFwmSettingLogFfa(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "FmupdateFwmSetting-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("multiple_steps_interval", flattenFmupdateFwmSettingMultipleStepsIntervalFfa(o["multiple-steps-interval"], d, "multiple_steps_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-steps-interval"], "FmupdateFwmSetting-MultipleStepsInterval"); ok {
			if err = d.Set("multiple_steps_interval", vv); err != nil {
				return fmt.Errorf("Error reading multiple_steps_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_steps_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("upgrade_timeout", flattenFmupdateFwmSettingUpgradeTimeoutFfa(o["upgrade-timeout"], d, "upgrade_timeout")); err != nil {
			if vv, ok := fortiAPIPatch(o["upgrade-timeout"], "FmupdateFwmSetting-UpgradeTimeout"); ok {
				if err = d.Set("upgrade_timeout", vv); err != nil {
					return fmt.Errorf("Error reading upgrade_timeout: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading upgrade_timeout: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("upgrade_timeout"); ok {
			if err = d.Set("upgrade_timeout", flattenFmupdateFwmSettingUpgradeTimeoutFfa(o["upgrade-timeout"], d, "upgrade_timeout")); err != nil {
				if vv, ok := fortiAPIPatch(o["upgrade-timeout"], "FmupdateFwmSetting-UpgradeTimeout"); ok {
					if err = d.Set("upgrade_timeout", vv); err != nil {
						return fmt.Errorf("Error reading upgrade_timeout: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading upgrade_timeout: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFmupdateFwmSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFwmSettingAutoScanFgtDiskFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingCheckFgtDiskFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingFdsFailoverFmgFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingFdsImageTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingImmxSourceFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingLogFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingMultipleStepsIntervalFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "check_status_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["check-status-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfa(d, i["check_status_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ctrl_check_status_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ctrl-check-status-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfa(d, i["ctrl_check_status_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ctrl_put_image_by_fds_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ctrl-put-image-by-fds-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfa(d, i["ctrl_put_image_by_fds_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "ha_sync_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ha-sync-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfa(d, i["ha_sync_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "license_check_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["license-check-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfa(d, i["license_check_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "prepare_image_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["prepare-image-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfa(d, i["prepare_image_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "put_image_by_fds_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["put-image-by-fds-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfa(d, i["put_image_by_fds_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "put_image_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["put-image-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfa(d, i["put_image_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "reboot_of_fsck_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reboot-of-fsck-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfa(d, i["reboot_of_fsck_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "reboot_of_upgrade_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["reboot-of-upgrade-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfa(d, i["reboot_of_upgrade_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "retrieve_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["retrieve-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfa(d, i["retrieve_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "rpc_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["rpc-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfa(d, i["rpc_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "total_timeout"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["total-timeout"], _ = expandFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfa(d, i["total_timeout"], pre_append)
	}

	return result, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutHaSyncTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPutImageTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRetrieveTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRpcTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutTotalTimeoutFfa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFwmSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_scan_fgt_disk"); ok || d.HasChange("auto_scan_fgt_disk") {
		t, err := expandFmupdateFwmSettingAutoScanFgtDiskFfa(d, v, "auto_scan_fgt_disk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-scan-fgt-disk"] = t
		}
	}

	if v, ok := d.GetOk("check_fgt_disk"); ok || d.HasChange("check_fgt_disk") {
		t, err := expandFmupdateFwmSettingCheckFgtDiskFfa(d, v, "check_fgt_disk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-fgt-disk"] = t
		}
	}

	if v, ok := d.GetOk("fds_failover_fmg"); ok || d.HasChange("fds_failover_fmg") {
		t, err := expandFmupdateFwmSettingFdsFailoverFmgFfa(d, v, "fds_failover_fmg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-failover-fmg"] = t
		}
	}

	if v, ok := d.GetOk("fds_image_timeout"); ok || d.HasChange("fds_image_timeout") {
		t, err := expandFmupdateFwmSettingFdsImageTimeoutFfa(d, v, "fds_image_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-image-timeout"] = t
		}
	}

	if v, ok := d.GetOk("immx_source"); ok || d.HasChange("immx_source") {
		t, err := expandFmupdateFwmSettingImmxSourceFfa(d, v, "immx_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["immx-source"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandFmupdateFwmSettingLogFfa(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("multiple_steps_interval"); ok || d.HasChange("multiple_steps_interval") {
		t, err := expandFmupdateFwmSettingMultipleStepsIntervalFfa(d, v, "multiple_steps_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-steps-interval"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_timeout"); ok || d.HasChange("upgrade_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutFfa(d, v, "upgrade_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-timeout"] = t
		}
	}

	return &obj, nil
}
