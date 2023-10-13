// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration options for modem 2.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerExtenderProfileCellularModem2() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerExtenderProfileCellularModem2Update,
		Read:   resourceObjectExtenderControllerExtenderProfileCellularModem2Read,
		Update: resourceObjectExtenderControllerExtenderProfileCellularModem2Update,
		Delete: resourceObjectExtenderControllerExtenderProfileCellularModem2Delete,

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
				Computed: true,
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

func resourceObjectExtenderControllerExtenderProfileCellularModem2Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtenderControllerExtenderProfileCellularModem2(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularModem2 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtenderControllerExtenderProfileCellularModem2(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerExtenderProfileCellularModem2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtenderControllerExtenderProfileCellularModem2")

	return resourceObjectExtenderControllerExtenderProfileCellularModem2Read(d, m)
}

func resourceObjectExtenderControllerExtenderProfileCellularModem2Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtenderControllerExtenderProfileCellularModem2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerExtenderProfileCellularModem2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerExtenderProfileCellularModem2Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtenderControllerExtenderProfileCellularModem2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularModem2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerExtenderProfileCellularModem2(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerExtenderProfileCellularModem2 resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch3rdl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan3rdl(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect3rdl(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod3rdl(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold3rdl(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal3rdl(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack3rdl(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime3rdl(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer3rdl(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2ConnStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2DefaultSim3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Gps3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2ModemId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantMode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectObjectExtenderControllerExtenderProfileCellularModem2(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if isImportTable() {
		if err = d.Set("auto_switch", flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch3rdl(o["auto-switch"], d, "auto_switch")); err != nil {
			if vv, ok := fortiAPIPatch(o["auto-switch"], "ObjectExtenderControllerExtenderProfileCellularModem2-AutoSwitch"); ok {
				if err = d.Set("auto_switch", vv); err != nil {
					return fmt.Errorf("Error reading auto_switch: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auto_switch: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auto_switch"); ok {
			if err = d.Set("auto_switch", flattenObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch3rdl(o["auto-switch"], d, "auto_switch")); err != nil {
				if vv, ok := fortiAPIPatch(o["auto-switch"], "ObjectExtenderControllerExtenderProfileCellularModem2-AutoSwitch"); ok {
					if err = d.Set("auto_switch", vv); err != nil {
						return fmt.Errorf("Error reading auto_switch: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auto_switch: %v", err)
				}
			}
		}
	}

	if err = d.Set("conn_status", flattenObjectExtenderControllerExtenderProfileCellularModem2ConnStatus3rdl(o["conn-status"], d, "conn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-status"], "ObjectExtenderControllerExtenderProfileCellularModem2-ConnStatus"); ok {
			if err = d.Set("conn_status", vv); err != nil {
				return fmt.Errorf("Error reading conn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("default_sim", flattenObjectExtenderControllerExtenderProfileCellularModem2DefaultSim3rdl(o["default-sim"], d, "default_sim")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-sim"], "ObjectExtenderControllerExtenderProfileCellularModem2-DefaultSim"); ok {
			if err = d.Set("default_sim", vv); err != nil {
				return fmt.Errorf("Error reading default_sim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_sim: %v", err)
		}
	}

	if err = d.Set("gps", flattenObjectExtenderControllerExtenderProfileCellularModem2Gps3rdl(o["gps"], d, "gps")); err != nil {
		if vv, ok := fortiAPIPatch(o["gps"], "ObjectExtenderControllerExtenderProfileCellularModem2-Gps"); ok {
			if err = d.Set("gps", vv); err != nil {
				return fmt.Errorf("Error reading gps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gps: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenObjectExtenderControllerExtenderProfileCellularModem2ModemId3rdl(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "ObjectExtenderControllerExtenderProfileCellularModem2-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("preferred_carrier", flattenObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier3rdl(o["preferred-carrier"], d, "preferred_carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-carrier"], "ObjectExtenderControllerExtenderProfileCellularModem2-PreferredCarrier"); ok {
			if err = d.Set("preferred_carrier", vv); err != nil {
				return fmt.Errorf("Error reading preferred_carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_carrier: %v", err)
		}
	}

	if err = d.Set("redundant_intf", flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf3rdl(o["redundant-intf"], d, "redundant_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-intf"], "ObjectExtenderControllerExtenderProfileCellularModem2-RedundantIntf"); ok {
			if err = d.Set("redundant_intf", vv); err != nil {
				return fmt.Errorf("Error reading redundant_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_intf: %v", err)
		}
	}

	if err = d.Set("redundant_mode", flattenObjectExtenderControllerExtenderProfileCellularModem2RedundantMode3rdl(o["redundant-mode"], d, "redundant_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-mode"], "ObjectExtenderControllerExtenderProfileCellularModem2-RedundantMode"); ok {
			if err = d.Set("redundant_mode", vv); err != nil {
				return fmt.Errorf("Error reading redundant_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_mode: %v", err)
		}
	}

	if err = d.Set("sim1_pin", flattenObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin3rdl(o["sim1-pin"], d, "sim1_pin")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim1-pin"], "ObjectExtenderControllerExtenderProfileCellularModem2-Sim1Pin"); ok {
			if err = d.Set("sim1_pin", vv); err != nil {
				return fmt.Errorf("Error reading sim1_pin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim1_pin: %v", err)
		}
	}

	if err = d.Set("sim2_pin", flattenObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin3rdl(o["sim2-pin"], d, "sim2_pin")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim2-pin"], "ObjectExtenderControllerExtenderProfileCellularModem2-Sim2Pin"); ok {
			if err = d.Set("sim2_pin", vv); err != nil {
				return fmt.Errorf("Error reading sim2_pin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim2_pin: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerExtenderProfileCellularModem2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan3rdl(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect3rdl(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod3rdl(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold3rdl(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal3rdl(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack3rdl(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime3rdl(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer3rdl(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2ConnStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2DefaultSim3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Gps3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2ModemId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2RedundantMode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectObjectExtenderControllerExtenderProfileCellularModem2(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_switch"); ok || d.HasChange("auto_switch") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2AutoSwitch3rdl(d, v, "auto_switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-switch"] = t
		}
	}

	if v, ok := d.GetOk("conn_status"); ok || d.HasChange("conn_status") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2ConnStatus3rdl(d, v, "conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-status"] = t
		}
	}

	if v, ok := d.GetOk("default_sim"); ok || d.HasChange("default_sim") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2DefaultSim3rdl(d, v, "default_sim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-sim"] = t
		}
	}

	if v, ok := d.GetOk("gps"); ok || d.HasChange("gps") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2Gps3rdl(d, v, "gps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gps"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok || d.HasChange("modem_id") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2ModemId3rdl(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("preferred_carrier"); ok || d.HasChange("preferred_carrier") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2PreferredCarrier3rdl(d, v, "preferred_carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-carrier"] = t
		}
	}

	if v, ok := d.GetOk("redundant_intf"); ok || d.HasChange("redundant_intf") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2RedundantIntf3rdl(d, v, "redundant_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-intf"] = t
		}
	}

	if v, ok := d.GetOk("redundant_mode"); ok || d.HasChange("redundant_mode") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2RedundantMode3rdl(d, v, "redundant_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-mode"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin"); ok || d.HasChange("sim1_pin") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2Sim1Pin3rdl(d, v, "sim1_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin_code"); ok || d.HasChange("sim1_pin_code") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2Sim1PinCode3rdl(d, v, "sim1_pin_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin-code"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin"); ok || d.HasChange("sim2_pin") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2Sim2Pin3rdl(d, v, "sim2_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin_code"); ok || d.HasChange("sim2_pin_code") {
		t, err := expandObjectExtenderControllerExtenderProfileCellularModem2Sim2PinCode3rdl(d, v, "sim2_pin_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin-code"] = t
		}
	}

	return &obj, nil
}
