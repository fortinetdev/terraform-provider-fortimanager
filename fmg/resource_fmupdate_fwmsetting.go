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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"multiple_steps_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFwmSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFwmSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFwmSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFwmSetting(obj, adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFwmSetting(adomv, mkey, nil)
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

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFwmSetting(adomv, mkey, nil)
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

func flattenFmupdateFwmSettingAutoScanFgtDisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingCheckFgtDisk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingFdsFailoverFmg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingFdsImageTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingImmxSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFwmSettingMultipleStepsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFwmSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_scan_fgt_disk", flattenFmupdateFwmSettingAutoScanFgtDisk(o["auto-scan-fgt-disk"], d, "auto_scan_fgt_disk")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-scan-fgt-disk"], "FmupdateFwmSetting-AutoScanFgtDisk"); ok {
			if err = d.Set("auto_scan_fgt_disk", vv); err != nil {
				return fmt.Errorf("Error reading auto_scan_fgt_disk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_scan_fgt_disk: %v", err)
		}
	}

	if err = d.Set("check_fgt_disk", flattenFmupdateFwmSettingCheckFgtDisk(o["check-fgt-disk"], d, "check_fgt_disk")); err != nil {
		if vv, ok := fortiAPIPatch(o["check-fgt-disk"], "FmupdateFwmSetting-CheckFgtDisk"); ok {
			if err = d.Set("check_fgt_disk", vv); err != nil {
				return fmt.Errorf("Error reading check_fgt_disk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading check_fgt_disk: %v", err)
		}
	}

	if err = d.Set("fds_failover_fmg", flattenFmupdateFwmSettingFdsFailoverFmg(o["fds-failover-fmg"], d, "fds_failover_fmg")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-failover-fmg"], "FmupdateFwmSetting-FdsFailoverFmg"); ok {
			if err = d.Set("fds_failover_fmg", vv); err != nil {
				return fmt.Errorf("Error reading fds_failover_fmg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_failover_fmg: %v", err)
		}
	}

	if err = d.Set("fds_image_timeout", flattenFmupdateFwmSettingFdsImageTimeout(o["fds-image-timeout"], d, "fds_image_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["fds-image-timeout"], "FmupdateFwmSetting-FdsImageTimeout"); ok {
			if err = d.Set("fds_image_timeout", vv); err != nil {
				return fmt.Errorf("Error reading fds_image_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fds_image_timeout: %v", err)
		}
	}

	if err = d.Set("immx_source", flattenFmupdateFwmSettingImmxSource(o["immx-source"], d, "immx_source")); err != nil {
		if vv, ok := fortiAPIPatch(o["immx-source"], "FmupdateFwmSetting-ImmxSource"); ok {
			if err = d.Set("immx_source", vv); err != nil {
				return fmt.Errorf("Error reading immx_source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading immx_source: %v", err)
		}
	}

	if err = d.Set("multiple_steps_interval", flattenFmupdateFwmSettingMultipleStepsInterval(o["multiple-steps-interval"], d, "multiple_steps_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiple-steps-interval"], "FmupdateFwmSetting-MultipleStepsInterval"); ok {
			if err = d.Set("multiple_steps_interval", vv); err != nil {
				return fmt.Errorf("Error reading multiple_steps_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiple_steps_interval: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFwmSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFwmSettingAutoScanFgtDisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingCheckFgtDisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingFdsFailoverFmg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingFdsImageTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingImmxSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFwmSettingMultipleStepsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFwmSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_scan_fgt_disk"); ok {
		t, err := expandFmupdateFwmSettingAutoScanFgtDisk(d, v, "auto_scan_fgt_disk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-scan-fgt-disk"] = t
		}
	}

	if v, ok := d.GetOk("check_fgt_disk"); ok {
		t, err := expandFmupdateFwmSettingCheckFgtDisk(d, v, "check_fgt_disk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-fgt-disk"] = t
		}
	}

	if v, ok := d.GetOk("fds_failover_fmg"); ok {
		t, err := expandFmupdateFwmSettingFdsFailoverFmg(d, v, "fds_failover_fmg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-failover-fmg"] = t
		}
	}

	if v, ok := d.GetOk("fds_image_timeout"); ok {
		t, err := expandFmupdateFwmSettingFdsImageTimeout(d, v, "fds_image_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fds-image-timeout"] = t
		}
	}

	if v, ok := d.GetOk("immx_source"); ok {
		t, err := expandFmupdateFwmSettingImmxSource(d, v, "immx_source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["immx-source"] = t
		}
	}

	if v, ok := d.GetOk("multiple_steps_interval"); ok {
		t, err := expandFmupdateFwmSettingMultipleStepsInterval(d, v, "multiple_steps_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple-steps-interval"] = t
		}
	}

	return &obj, nil
}
