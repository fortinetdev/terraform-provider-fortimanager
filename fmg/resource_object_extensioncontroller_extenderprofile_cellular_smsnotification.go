// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender cellular SMS notification configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotification() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationUpdate,
		Read:   resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationRead,
		Update: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationDelete,

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
			"alert": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"receiver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"phone_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["extender_profile"] = extender_profile

	obj, err := getObjectObjectExtensionControllerExtenderProfileCellularSmsNotification(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularSmsNotification resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerExtenderProfileCellularSmsNotification(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularSmsNotification resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileCellularSmsNotification")

	return resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["extender_profile"] = extender_profile

	err = c.DeleteObjectExtensionControllerExtenderProfileCellularSmsNotification(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileCellularSmsNotification resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileCellularSmsNotificationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileCellularSmsNotification(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularSmsNotification resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileCellularSmsNotification(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularSmsNotification resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("alert", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(o["alert"], d, "alert")); err != nil {
			if vv, ok := fortiAPIPatch(o["alert"], "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Alert"); ok {
				if err = d.Set("alert", vv); err != nil {
					return fmt.Errorf("Error reading alert: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading alert: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alert"); ok {
			if err = d.Set("alert", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(o["alert"], d, "alert")); err != nil {
				if vv, ok := fortiAPIPatch(o["alert"], "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Alert"); ok {
					if err = d.Set("alert", vv); err != nil {
						return fmt.Errorf("Error reading alert: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading alert: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("receiver", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(o["receiver"], d, "receiver")); err != nil {
			if vv, ok := fortiAPIPatch(o["receiver"], "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver"); ok {
				if err = d.Set("receiver", vv); err != nil {
					return fmt.Errorf("Error reading receiver: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading receiver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("receiver"); ok {
			if err = d.Set("receiver", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(o["receiver"], d, "receiver")); err != nil {
				if vv, ok := fortiAPIPatch(o["receiver"], "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver"); ok {
					if err = d.Set("receiver", vv); err != nil {
						return fmt.Errorf("Error reading receiver: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading receiver: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alert"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alert"); ok || d.HasChange("alert") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert3rdl(d, v, "alert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert"] = t
		}
	}

	if v, ok := d.GetOk("receiver"); ok || d.HasChange("receiver") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver3rdl(d, v, "receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["receiver"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
