// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the timeout value of image upgrade process.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFwmSettingUpgradeTimeout() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFwmSettingUpgradeTimeoutUpdate,
		Read:   resourceFmupdateFwmSettingUpgradeTimeoutRead,
		Update: resourceFmupdateFwmSettingUpgradeTimeoutUpdate,
		Delete: resourceFmupdateFwmSettingUpgradeTimeoutDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
	}
}

func resourceFmupdateFwmSettingUpgradeTimeoutUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFwmSettingUpgradeTimeout(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSettingUpgradeTimeout resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFwmSettingUpgradeTimeout(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSettingUpgradeTimeout resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFwmSettingUpgradeTimeout")

	return resourceFmupdateFwmSettingUpgradeTimeoutRead(d, m)
}

func resourceFmupdateFwmSettingUpgradeTimeoutDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteFmupdateFwmSettingUpgradeTimeout(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFwmSettingUpgradeTimeout resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFwmSettingUpgradeTimeoutRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFwmSettingUpgradeTimeout(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFwmSettingUpgradeTimeout resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFwmSettingUpgradeTimeout(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFwmSettingUpgradeTimeout resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutHaSyncTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutPutImageTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRetrieveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutRpcTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingUpgradeTimeoutTotalTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFwmSettingUpgradeTimeout(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("check_status_timeout", flattenFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeout(o["check-status-timeout"], d, "check_status_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-status-timeout"], "FmupdateFwmSettingUpgradeTimeout-CheckStatusTimeout"); ok {
			if err = d.Set("check_status_timeout", vv); err != nil {
				return fmt.Errorf("Error reading check_status_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_status_timeout: %v", err)
		}
	}

	if err = d.Set("ctrl_check_status_timeout", flattenFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeout(o["ctrl-check-status-timeout"], d, "ctrl_check_status_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ctrl-check-status-timeout"], "FmupdateFwmSettingUpgradeTimeout-CtrlCheckStatusTimeout"); ok {
			if err = d.Set("ctrl_check_status_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ctrl_check_status_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ctrl_check_status_timeout: %v", err)
		}
	}

	if err = d.Set("ctrl_put_image_by_fds_timeout", flattenFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeout(o["ctrl-put-image-by-fds-timeout"], d, "ctrl_put_image_by_fds_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ctrl-put-image-by-fds-timeout"], "FmupdateFwmSettingUpgradeTimeout-CtrlPutImageByFdsTimeout"); ok {
			if err = d.Set("ctrl_put_image_by_fds_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ctrl_put_image_by_fds_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ctrl_put_image_by_fds_timeout: %v", err)
		}
	}

	if err = d.Set("ha_sync_timeout", flattenFmupdateFwmSettingUpgradeTimeoutHaSyncTimeout(o["ha-sync-timeout"], d, "ha_sync_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-sync-timeout"], "FmupdateFwmSettingUpgradeTimeout-HaSyncTimeout"); ok {
			if err = d.Set("ha_sync_timeout", vv); err != nil {
				return fmt.Errorf("Error reading ha_sync_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_sync_timeout: %v", err)
		}
	}

	if err = d.Set("license_check_timeout", flattenFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeout(o["license-check-timeout"], d, "license_check_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["license-check-timeout"], "FmupdateFwmSettingUpgradeTimeout-LicenseCheckTimeout"); ok {
			if err = d.Set("license_check_timeout", vv); err != nil {
				return fmt.Errorf("Error reading license_check_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading license_check_timeout: %v", err)
		}
	}

	if err = d.Set("prepare_image_timeout", flattenFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeout(o["prepare-image-timeout"], d, "prepare_image_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["prepare-image-timeout"], "FmupdateFwmSettingUpgradeTimeout-PrepareImageTimeout"); ok {
			if err = d.Set("prepare_image_timeout", vv); err != nil {
				return fmt.Errorf("Error reading prepare_image_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prepare_image_timeout: %v", err)
		}
	}

	if err = d.Set("put_image_by_fds_timeout", flattenFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeout(o["put-image-by-fds-timeout"], d, "put_image_by_fds_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["put-image-by-fds-timeout"], "FmupdateFwmSettingUpgradeTimeout-PutImageByFdsTimeout"); ok {
			if err = d.Set("put_image_by_fds_timeout", vv); err != nil {
				return fmt.Errorf("Error reading put_image_by_fds_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading put_image_by_fds_timeout: %v", err)
		}
	}

	if err = d.Set("put_image_timeout", flattenFmupdateFwmSettingUpgradeTimeoutPutImageTimeout(o["put-image-timeout"], d, "put_image_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["put-image-timeout"], "FmupdateFwmSettingUpgradeTimeout-PutImageTimeout"); ok {
			if err = d.Set("put_image_timeout", vv); err != nil {
				return fmt.Errorf("Error reading put_image_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading put_image_timeout: %v", err)
		}
	}

	if err = d.Set("reboot_of_fsck_timeout", flattenFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeout(o["reboot-of-fsck-timeout"], d, "reboot_of_fsck_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["reboot-of-fsck-timeout"], "FmupdateFwmSettingUpgradeTimeout-RebootOfFsckTimeout"); ok {
			if err = d.Set("reboot_of_fsck_timeout", vv); err != nil {
				return fmt.Errorf("Error reading reboot_of_fsck_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reboot_of_fsck_timeout: %v", err)
		}
	}

	if err = d.Set("reboot_of_upgrade_timeout", flattenFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeout(o["reboot-of-upgrade-timeout"], d, "reboot_of_upgrade_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["reboot-of-upgrade-timeout"], "FmupdateFwmSettingUpgradeTimeout-RebootOfUpgradeTimeout"); ok {
			if err = d.Set("reboot_of_upgrade_timeout", vv); err != nil {
				return fmt.Errorf("Error reading reboot_of_upgrade_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reboot_of_upgrade_timeout: %v", err)
		}
	}

	if err = d.Set("retrieve_timeout", flattenFmupdateFwmSettingUpgradeTimeoutRetrieveTimeout(o["retrieve-timeout"], d, "retrieve_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["retrieve-timeout"], "FmupdateFwmSettingUpgradeTimeout-RetrieveTimeout"); ok {
			if err = d.Set("retrieve_timeout", vv); err != nil {
				return fmt.Errorf("Error reading retrieve_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retrieve_timeout: %v", err)
		}
	}

	if err = d.Set("rpc_timeout", flattenFmupdateFwmSettingUpgradeTimeoutRpcTimeout(o["rpc-timeout"], d, "rpc_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["rpc-timeout"], "FmupdateFwmSettingUpgradeTimeout-RpcTimeout"); ok {
			if err = d.Set("rpc_timeout", vv); err != nil {
				return fmt.Errorf("Error reading rpc_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rpc_timeout: %v", err)
		}
	}

	if err = d.Set("total_timeout", flattenFmupdateFwmSettingUpgradeTimeoutTotalTimeout(o["total-timeout"], d, "total_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["total-timeout"], "FmupdateFwmSettingUpgradeTimeout-TotalTimeout"); ok {
			if err = d.Set("total_timeout", vv); err != nil {
				return fmt.Errorf("Error reading total_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading total_timeout: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFwmSettingUpgradeTimeoutFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutHaSyncTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutPutImageTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRetrieveTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutRpcTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingUpgradeTimeoutTotalTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFwmSettingUpgradeTimeout(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("check_status_timeout"); ok || d.HasChange("check_status_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutCheckStatusTimeout(d, v, "check_status_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-status-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ctrl_check_status_timeout"); ok || d.HasChange("ctrl_check_status_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutCtrlCheckStatusTimeout(d, v, "ctrl_check_status_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ctrl-check-status-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ctrl_put_image_by_fds_timeout"); ok || d.HasChange("ctrl_put_image_by_fds_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutCtrlPutImageByFdsTimeout(d, v, "ctrl_put_image_by_fds_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ctrl-put-image-by-fds-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ha_sync_timeout"); ok || d.HasChange("ha_sync_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutHaSyncTimeout(d, v, "ha_sync_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-sync-timeout"] = t
		}
	}

	if v, ok := d.GetOk("license_check_timeout"); ok || d.HasChange("license_check_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutLicenseCheckTimeout(d, v, "license_check_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license-check-timeout"] = t
		}
	}

	if v, ok := d.GetOk("prepare_image_timeout"); ok || d.HasChange("prepare_image_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutPrepareImageTimeout(d, v, "prepare_image_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prepare-image-timeout"] = t
		}
	}

	if v, ok := d.GetOk("put_image_by_fds_timeout"); ok || d.HasChange("put_image_by_fds_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutPutImageByFdsTimeout(d, v, "put_image_by_fds_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["put-image-by-fds-timeout"] = t
		}
	}

	if v, ok := d.GetOk("put_image_timeout"); ok || d.HasChange("put_image_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutPutImageTimeout(d, v, "put_image_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["put-image-timeout"] = t
		}
	}

	if v, ok := d.GetOk("reboot_of_fsck_timeout"); ok || d.HasChange("reboot_of_fsck_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutRebootOfFsckTimeout(d, v, "reboot_of_fsck_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reboot-of-fsck-timeout"] = t
		}
	}

	if v, ok := d.GetOk("reboot_of_upgrade_timeout"); ok || d.HasChange("reboot_of_upgrade_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutRebootOfUpgradeTimeout(d, v, "reboot_of_upgrade_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reboot-of-upgrade-timeout"] = t
		}
	}

	if v, ok := d.GetOk("retrieve_timeout"); ok || d.HasChange("retrieve_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutRetrieveTimeout(d, v, "retrieve_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retrieve-timeout"] = t
		}
	}

	if v, ok := d.GetOk("rpc_timeout"); ok || d.HasChange("rpc_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutRpcTimeout(d, v, "rpc_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-timeout"] = t
		}
	}

	if v, ok := d.GetOk("total_timeout"); ok || d.HasChange("total_timeout") {
		t, err := expandFmupdateFwmSettingUpgradeTimeoutTotalTimeout(d, v, "total_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["total-timeout"] = t
		}
	}

	return &obj, nil
}
