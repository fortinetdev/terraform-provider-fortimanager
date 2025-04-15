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

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertUpdate,
		Read:   resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertRead,
		Update: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDelete,

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

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert")

	return resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("data_exhausted", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(o["data-exhausted"], d, "data_exhausted")); err != nil {
		if vv, ok := fortiAPIPatch(o["data-exhausted"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-DataExhausted"); ok {
			if err = d.Set("data_exhausted", vv); err != nil {
				return fmt.Errorf("Error reading data_exhausted: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading data_exhausted: %v", err)
		}
	}

	if err = d.Set("fgt_backup_mode_switch", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(o["fgt-backup-mode-switch"], d, "fgt_backup_mode_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgt-backup-mode-switch"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-FgtBackupModeSwitch"); ok {
			if err = d.Set("fgt_backup_mode_switch", vv); err != nil {
				return fmt.Errorf("Error reading fgt_backup_mode_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgt_backup_mode_switch: %v", err)
		}
	}

	if err = d.Set("low_signal_strength", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(o["low-signal-strength"], d, "low_signal_strength")); err != nil {
		if vv, ok := fortiAPIPatch(o["low-signal-strength"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-LowSignalStrength"); ok {
			if err = d.Set("low_signal_strength", vv); err != nil {
				return fmt.Errorf("Error reading low_signal_strength: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading low_signal_strength: %v", err)
		}
	}

	if err = d.Set("mode_switch", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(o["mode-switch"], d, "mode_switch")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode-switch"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-ModeSwitch"); ok {
			if err = d.Set("mode_switch", vv); err != nil {
				return fmt.Errorf("Error reading mode_switch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode_switch: %v", err)
		}
	}

	if err = d.Set("os_image_fallback", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(o["os-image-fallback"], d, "os_image_fallback")); err != nil {
		if vv, ok := fortiAPIPatch(o["os-image-fallback"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-OsImageFallback"); ok {
			if err = d.Set("os_image_fallback", vv); err != nil {
				return fmt.Errorf("Error reading os_image_fallback: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_image_fallback: %v", err)
		}
	}

	if err = d.Set("session_disconnect", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(o["session-disconnect"], d, "session_disconnect")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-disconnect"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-SessionDisconnect"); ok {
			if err = d.Set("session_disconnect", vv); err != nil {
				return fmt.Errorf("Error reading session_disconnect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_disconnect: %v", err)
		}
	}

	if err = d.Set("system_reboot", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(o["system-reboot"], d, "system_reboot")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-reboot"], "ObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert-SystemReboot"); ok {
			if err = d.Set("system_reboot", vv); err != nil {
				return fmt.Errorf("Error reading system_reboot: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_reboot: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("data_exhausted"); ok || d.HasChange("data_exhausted") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted4thl(d, v, "data_exhausted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-exhausted"] = t
		}
	}

	if v, ok := d.GetOk("fgt_backup_mode_switch"); ok || d.HasChange("fgt_backup_mode_switch") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch4thl(d, v, "fgt_backup_mode_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgt-backup-mode-switch"] = t
		}
	}

	if v, ok := d.GetOk("low_signal_strength"); ok || d.HasChange("low_signal_strength") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength4thl(d, v, "low_signal_strength")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low-signal-strength"] = t
		}
	}

	if v, ok := d.GetOk("mode_switch"); ok || d.HasChange("mode_switch") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch4thl(d, v, "mode_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-switch"] = t
		}
	}

	if v, ok := d.GetOk("os_image_fallback"); ok || d.HasChange("os_image_fallback") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback4thl(d, v, "os_image_fallback")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-image-fallback"] = t
		}
	}

	if v, ok := d.GetOk("session_disconnect"); ok || d.HasChange("session_disconnect") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect4thl(d, v, "session_disconnect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-disconnect"] = t
		}
	}

	if v, ok := d.GetOk("system_reboot"); ok || d.HasChange("system_reboot") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot4thl(d, v, "system_reboot")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-reboot"] = t
		}
	}

	return &obj, nil
}
