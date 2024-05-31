// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration options for modem 1.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileCellularModem1() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileCellularModem1Update,
		Read:   resourceObjectExtensionControllerExtenderProfileCellularModem1Read,
		Update: resourceObjectExtensionControllerExtenderProfileCellularModem1Update,
		Delete: resourceObjectExtensionControllerExtenderProfileCellularModem1Delete,

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
			"auto_switch": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataplan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disconnect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"disconnect_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"disconnect_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"signal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switch_back": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"switch_back_time": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switch_back_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"conn_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"default_sim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"preferred_carrier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redundant_intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redundant_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sim1_pin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sim1_pin_code": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"sim2_pin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sim2_pin_code": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileCellularModem1Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileCellularModem1(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularModem1 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerExtenderProfileCellularModem1(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellularModem1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileCellularModem1")

	return resourceObjectExtensionControllerExtenderProfileCellularModem1Read(d, m)
}

func resourceObjectExtensionControllerExtenderProfileCellularModem1Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfileCellularModem1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileCellularModem1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileCellularModem1Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileCellularModem1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularModem1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileCellularModem1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellularModem1 resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan3rdl(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect3rdl(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod3rdl(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold3rdl(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal3rdl(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack3rdl(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime3rdl(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer3rdl(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1ConnStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1DefaultSim3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Gps3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1ModemId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileCellularModem1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if isImportTable() {
		if err = d.Set("auto_switch", flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch3rdl(o["auto-switch"], d, "auto_switch")); err != nil {
			if vv, ok := fortiAPIPatch(o["auto-switch"], "ObjectExtensionControllerExtenderProfileCellularModem1-AutoSwitch"); ok {
				if err = d.Set("auto_switch", vv); err != nil {
					return fmt.Errorf("Error reading auto_switch: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auto_switch: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auto_switch"); ok {
			if err = d.Set("auto_switch", flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch3rdl(o["auto-switch"], d, "auto_switch")); err != nil {
				if vv, ok := fortiAPIPatch(o["auto-switch"], "ObjectExtensionControllerExtenderProfileCellularModem1-AutoSwitch"); ok {
					if err = d.Set("auto_switch", vv); err != nil {
						return fmt.Errorf("Error reading auto_switch: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auto_switch: %v", err)
				}
			}
		}
	}

	if err = d.Set("conn_status", flattenObjectExtensionControllerExtenderProfileCellularModem1ConnStatus3rdl(o["conn-status"], d, "conn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-status"], "ObjectExtensionControllerExtenderProfileCellularModem1-ConnStatus"); ok {
			if err = d.Set("conn_status", vv); err != nil {
				return fmt.Errorf("Error reading conn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("default_sim", flattenObjectExtensionControllerExtenderProfileCellularModem1DefaultSim3rdl(o["default-sim"], d, "default_sim")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-sim"], "ObjectExtensionControllerExtenderProfileCellularModem1-DefaultSim"); ok {
			if err = d.Set("default_sim", vv); err != nil {
				return fmt.Errorf("Error reading default_sim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_sim: %v", err)
		}
	}

	if err = d.Set("gps", flattenObjectExtensionControllerExtenderProfileCellularModem1Gps3rdl(o["gps"], d, "gps")); err != nil {
		if vv, ok := fortiAPIPatch(o["gps"], "ObjectExtensionControllerExtenderProfileCellularModem1-Gps"); ok {
			if err = d.Set("gps", vv); err != nil {
				return fmt.Errorf("Error reading gps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gps: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenObjectExtensionControllerExtenderProfileCellularModem1ModemId3rdl(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "ObjectExtensionControllerExtenderProfileCellularModem1-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("preferred_carrier", flattenObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier3rdl(o["preferred-carrier"], d, "preferred_carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-carrier"], "ObjectExtensionControllerExtenderProfileCellularModem1-PreferredCarrier"); ok {
			if err = d.Set("preferred_carrier", vv); err != nil {
				return fmt.Errorf("Error reading preferred_carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_carrier: %v", err)
		}
	}

	if err = d.Set("redundant_intf", flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf3rdl(o["redundant-intf"], d, "redundant_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-intf"], "ObjectExtensionControllerExtenderProfileCellularModem1-RedundantIntf"); ok {
			if err = d.Set("redundant_intf", vv); err != nil {
				return fmt.Errorf("Error reading redundant_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_intf: %v", err)
		}
	}

	if err = d.Set("redundant_mode", flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantMode3rdl(o["redundant-mode"], d, "redundant_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-mode"], "ObjectExtensionControllerExtenderProfileCellularModem1-RedundantMode"); ok {
			if err = d.Set("redundant_mode", vv); err != nil {
				return fmt.Errorf("Error reading redundant_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_mode: %v", err)
		}
	}

	if err = d.Set("sim1_pin", flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin3rdl(o["sim1-pin"], d, "sim1_pin")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim1-pin"], "ObjectExtensionControllerExtenderProfileCellularModem1-Sim1Pin"); ok {
			if err = d.Set("sim1_pin", vv); err != nil {
				return fmt.Errorf("Error reading sim1_pin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim1_pin: %v", err)
		}
	}

	if err = d.Set("sim2_pin", flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin3rdl(o["sim2-pin"], d, "sim2_pin")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim2-pin"], "ObjectExtensionControllerExtenderProfileCellularModem1-Sim2Pin"); ok {
			if err = d.Set("sim2_pin", vv); err != nil {
				return fmt.Errorf("Error reading sim2_pin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim2_pin: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan3rdl(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect3rdl(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod3rdl(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold3rdl(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal3rdl(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack3rdl(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime3rdl(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer3rdl(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1ConnStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1DefaultSim3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Gps3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1ModemId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1RedundantMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectExtensionControllerExtenderProfileCellularModem1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_switch"); ok || d.HasChange("auto_switch") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch3rdl(d, v, "auto_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-switch"] = t
		}
	}

	if v, ok := d.GetOk("conn_status"); ok || d.HasChange("conn_status") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1ConnStatus3rdl(d, v, "conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-status"] = t
		}
	}

	if v, ok := d.GetOk("default_sim"); ok || d.HasChange("default_sim") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1DefaultSim3rdl(d, v, "default_sim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-sim"] = t
		}
	}

	if v, ok := d.GetOk("gps"); ok || d.HasChange("gps") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1Gps3rdl(d, v, "gps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gps"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok || d.HasChange("modem_id") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1ModemId3rdl(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("preferred_carrier"); ok || d.HasChange("preferred_carrier") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier3rdl(d, v, "preferred_carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-carrier"] = t
		}
	}

	if v, ok := d.GetOk("redundant_intf"); ok || d.HasChange("redundant_intf") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf3rdl(d, v, "redundant_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-intf"] = t
		}
	}

	if v, ok := d.GetOk("redundant_mode"); ok || d.HasChange("redundant_mode") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1RedundantMode3rdl(d, v, "redundant_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-mode"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin"); ok || d.HasChange("sim1_pin") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin3rdl(d, v, "sim1_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin_code"); ok || d.HasChange("sim1_pin_code") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode3rdl(d, v, "sim1_pin_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin-code"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin"); ok || d.HasChange("sim2_pin") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin3rdl(d, v, "sim2_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin_code"); ok || d.HasChange("sim2_pin_code") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode3rdl(d, v, "sim2_pin_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin-code"] = t
		}
	}

	return &obj, nil
}
