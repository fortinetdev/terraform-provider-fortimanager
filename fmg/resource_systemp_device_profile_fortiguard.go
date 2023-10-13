// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Systemp DeviceProfileFortiguard

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystempDeviceProfileFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystempDeviceProfileFortiguardUpdate,
		Read:   resourceSystempDeviceProfileFortiguardRead,
		Update: resourceSystempDeviceProfileFortiguardUpdate,
		Delete: resourceSystempDeviceProfileFortiguardDelete,

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
			"devprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_day": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_end_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_firmware_upgrade_start_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystempDeviceProfileFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	obj, err := getObjectSystempDeviceProfileFortiguard(d)
	if err != nil {
		return fmt.Errorf("Error updating SystempDeviceProfileFortiguard resource while getting object: %v", err)
	}

	_, err = c.UpdateSystempDeviceProfileFortiguard(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystempDeviceProfileFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystempDeviceProfileFortiguard")

	return resourceSystempDeviceProfileFortiguardRead(d, m)
}

func resourceSystempDeviceProfileFortiguardDelete(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	paradict["devprof"] = devprof

	err = c.DeleteSystempDeviceProfileFortiguard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystempDeviceProfileFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystempDeviceProfileFortiguardRead(d *schema.ResourceData, m interface{}) error {
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

	devprof := d.Get("devprof").(string)
	if devprof == "" {
		devprof = importOptionChecking(m.(*FortiClient).Cfg, "devprof")
		if devprof == "" {
			return fmt.Errorf("Parameter devprof is missing")
		}
		if err = d.Set("devprof", devprof); err != nil {
			return fmt.Errorf("Error set params devprof: %v", err)
		}
	}
	paradict["devprof"] = devprof

	o, err := c.ReadSystempDeviceProfileFortiguard(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystempDeviceProfileFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystempDeviceProfileFortiguard(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystempDeviceProfileFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystempDeviceProfileFortiguardAutoFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeEndHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeStartHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortiguardTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystempDeviceProfileFortiguardTargetIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystempDeviceProfileFortiguard(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("auto_firmware_upgrade", flattenSystempDeviceProfileFortiguardAutoFirmwareUpgrade(o["auto-firmware-upgrade"], d, "auto_firmware_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade"], "SystempDeviceProfileFortiguard-AutoFirmwareUpgrade"); ok {
			if err = d.Set("auto_firmware_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_day", flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeDay(o["auto-firmware-upgrade-day"], d, "auto_firmware_upgrade_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-day"], "SystempDeviceProfileFortiguard-AutoFirmwareUpgradeDay"); ok {
			if err = d.Set("auto_firmware_upgrade_day", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_day: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_delay", flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeDelay(o["auto-firmware-upgrade-delay"], d, "auto_firmware_upgrade_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-delay"], "SystempDeviceProfileFortiguard-AutoFirmwareUpgradeDelay"); ok {
			if err = d.Set("auto_firmware_upgrade_delay", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_delay: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_end_hour", flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeEndHour(o["auto-firmware-upgrade-end-hour"], d, "auto_firmware_upgrade_end_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-end-hour"], "SystempDeviceProfileFortiguard-AutoFirmwareUpgradeEndHour"); ok {
			if err = d.Set("auto_firmware_upgrade_end_hour", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_end_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_end_hour: %v", err)
		}
	}

	if err = d.Set("auto_firmware_upgrade_start_hour", flattenSystempDeviceProfileFortiguardAutoFirmwareUpgradeStartHour(o["auto-firmware-upgrade-start-hour"], d, "auto_firmware_upgrade_start_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-firmware-upgrade-start-hour"], "SystempDeviceProfileFortiguard-AutoFirmwareUpgradeStartHour"); ok {
			if err = d.Set("auto_firmware_upgrade_start_hour", vv); err != nil {
				return fmt.Errorf("Error reading auto_firmware_upgrade_start_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_firmware_upgrade_start_hour: %v", err)
		}
	}

	if err = d.Set("target", flattenSystempDeviceProfileFortiguardTarget(o["target"], d, "target")); err != nil {
		if vv, ok := fortiAPIPatch(o["target"], "SystempDeviceProfileFortiguard-Target"); ok {
			if err = d.Set("target", vv); err != nil {
				return fmt.Errorf("Error reading target: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target: %v", err)
		}
	}

	if err = d.Set("target_ip", flattenSystempDeviceProfileFortiguardTargetIp(o["target-ip"], d, "target_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["target-ip"], "SystempDeviceProfileFortiguard-TargetIp"); ok {
			if err = d.Set("target_ip", vv); err != nil {
				return fmt.Errorf("Error reading target_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading target_ip: %v", err)
		}
	}

	return nil
}

func flattenSystempDeviceProfileFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystempDeviceProfileFortiguardAutoFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeEndHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeStartHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortiguardTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystempDeviceProfileFortiguardTargetIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystempDeviceProfileFortiguard(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_firmware_upgrade"); ok || d.HasChange("auto_firmware_upgrade") {
		t, err := expandSystempDeviceProfileFortiguardAutoFirmwareUpgrade(d, v, "auto_firmware_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_day"); ok || d.HasChange("auto_firmware_upgrade_day") {
		t, err := expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeDay(d, v, "auto_firmware_upgrade_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-day"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_delay"); ok || d.HasChange("auto_firmware_upgrade_delay") {
		t, err := expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeDelay(d, v, "auto_firmware_upgrade_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-delay"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_end_hour"); ok || d.HasChange("auto_firmware_upgrade_end_hour") {
		t, err := expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeEndHour(d, v, "auto_firmware_upgrade_end_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-end-hour"] = t
		}
	}

	if v, ok := d.GetOk("auto_firmware_upgrade_start_hour"); ok || d.HasChange("auto_firmware_upgrade_start_hour") {
		t, err := expandSystempDeviceProfileFortiguardAutoFirmwareUpgradeStartHour(d, v, "auto_firmware_upgrade_start_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-firmware-upgrade-start-hour"] = t
		}
	}

	if v, ok := d.GetOk("target"); ok || d.HasChange("target") {
		t, err := expandSystempDeviceProfileFortiguardTarget(d, v, "target")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target"] = t
		}
	}

	if v, ok := d.GetOk("target_ip"); ok || d.HasChange("target_ip") {
		t, err := expandSystempDeviceProfileFortiguardTargetIp(d, v, "target_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-ip"] = t
		}
	}

	return &obj, nil
}
