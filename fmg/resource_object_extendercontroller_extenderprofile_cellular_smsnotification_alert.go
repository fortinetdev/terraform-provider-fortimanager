// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SMS alert list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertUpdate,
		Read:   resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertRead,
		Update: resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertUpdate,
		Delete: resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"extender_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_exhausted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgt_backup_mode_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"low_signal_strength": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_image_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_disconnect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_reboot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert")

	return resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertRead(d, m)
}

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	extender_profile := d.Get("extender_profile").(string)
	paradict["extender_profile"] = extender_profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	extender_profile := d.Get("extender_profile").(string)
	if extender_profile == "" {
		extender_profile = importOptionChecking(m.(*FortiClient).Cfg, "extender_profile")
		if extender_profile == "" {
			return fmt.Errorf("Parameter extender_profile is missing")
		}
		if err = d.Set("extender_profile", extender_profile); err != nil {
			return fmt.Errorf("Error set params extender_profile: %v", err)
		}
	}
	paradict["extender_profile"] = extender_profile

	o, err := c.ReadObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("data_exhausted", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(o["data-exhausted"], d, "data_exhausted")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-exhausted"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-DataExhausted"); ok {
			if err = d.Set("data_exhausted", vv); err != nil {
				return fmt.Errorf("Error reading data_exhausted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_exhausted: %v", err)
		}
	}

	if err = d.Set("fgt_backup_mode_switch", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(o["fgt-backup-mode-switch"], d, "fgt_backup_mode_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-backup-mode-switch"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-FgtBackupModeSwitch"); ok {
			if err = d.Set("fgt_backup_mode_switch", vv); err != nil {
				return fmt.Errorf("Error reading fgt_backup_mode_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_backup_mode_switch: %v", err)
		}
	}

	if err = d.Set("low_signal_strength", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(o["low-signal-strength"], d, "low_signal_strength")); err != nil {
		if vv, ok := fortiAPIPatch(o["low-signal-strength"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-LowSignalStrength"); ok {
			if err = d.Set("low_signal_strength", vv); err != nil {
				return fmt.Errorf("Error reading low_signal_strength: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low_signal_strength: %v", err)
		}
	}

	if err = d.Set("mode_switch", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(o["mode-switch"], d, "mode_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-switch"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-ModeSwitch"); ok {
			if err = d.Set("mode_switch", vv); err != nil {
				return fmt.Errorf("Error reading mode_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_switch: %v", err)
		}
	}

	if err = d.Set("os_image_fallback", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(o["os-image-fallback"], d, "os_image_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-image-fallback"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-OsImageFallback"); ok {
			if err = d.Set("os_image_fallback", vv); err != nil {
				return fmt.Errorf("Error reading os_image_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_image_fallback: %v", err)
		}
	}

	if err = d.Set("session_disconnect", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(o["session-disconnect"], d, "session_disconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-disconnect"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-SessionDisconnect"); ok {
			if err = d.Set("session_disconnect", vv); err != nil {
				return fmt.Errorf("Error reading session_disconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_disconnect: %v", err)
		}
	}

	if err = d.Set("system_reboot", flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(o["system-reboot"], d, "system_reboot")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-reboot"], "ObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert-SystemReboot"); ok {
			if err = d.Set("system_reboot", vv); err != nil {
				return fmt.Errorf("Error reading system_reboot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_reboot: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("data_exhausted"); ok || d.HasChange("data_exhausted") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(d, v, "data_exhausted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-exhausted"] = t
		}
	}

	if v, ok := d.GetOk("fgt_backup_mode_switch"); ok || d.HasChange("fgt_backup_mode_switch") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(d, v, "fgt_backup_mode_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-backup-mode-switch"] = t
		}
	}

	if v, ok := d.GetOk("low_signal_strength"); ok || d.HasChange("low_signal_strength") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(d, v, "low_signal_strength")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low-signal-strength"] = t
		}
	}

	if v, ok := d.GetOk("mode_switch"); ok || d.HasChange("mode_switch") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(d, v, "mode_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-switch"] = t
		}
	}

	if v, ok := d.GetOk("os_image_fallback"); ok || d.HasChange("os_image_fallback") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(d, v, "os_image_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-image-fallback"] = t
		}
	}

	if v, ok := d.GetOk("session_disconnect"); ok || d.HasChange("session_disconnect") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(d, v, "session_disconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-disconnect"] = t
		}
	}

	if v, ok := d.GetOk("system_reboot"); ok || d.HasChange("system_reboot") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(d, v, "system_reboot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-reboot"] = t
		}
	}

	return &obj, nil
}
