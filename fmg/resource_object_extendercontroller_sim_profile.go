// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectExtenderController SimProfile

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtenderControllerSimProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtenderControllerSimProfileCreate,
		Read:   resourceObjectExtenderControllerSimProfileRead,
		Update: resourceObjectExtenderControllerSimProfileUpdate,
		Delete: resourceObjectExtenderControllerSimProfileDelete,

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
			"auto_switch_profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switch_back": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectExtenderControllerSimProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerSimProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerSimProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectExtenderControllerSimProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectExtenderControllerSimProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerSimProfileRead(d, m)
}

func resourceObjectExtenderControllerSimProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectExtenderControllerSimProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerSimProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtenderControllerSimProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtenderControllerSimProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtenderControllerSimProfileRead(d, m)
}

func resourceObjectExtenderControllerSimProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectExtenderControllerSimProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtenderControllerSimProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtenderControllerSimProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectExtenderControllerSimProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerSimProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtenderControllerSimProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtenderControllerSimProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileDefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileGps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfilePreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileRedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileRedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileSim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileSim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerSimProfileSim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtenderControllerSimProfileSim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtenderControllerSimProfileStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtenderControllerSimProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if isImportTable() {
		if err = d.Set("auto_switch_profile", flattenObjectExtenderControllerSimProfileAutoSwitchProfile(o["auto-switch_profile"], d, "auto_switch_profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["auto-switch_profile"], "ObjectExtenderControllerSimProfile-AutoSwitchProfile"); ok {
				if err = d.Set("auto_switch_profile", vv); err != nil {
					return fmt.Errorf("Error reading auto_switch_profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auto_switch_profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auto_switch_profile"); ok {
			if err = d.Set("auto_switch_profile", flattenObjectExtenderControllerSimProfileAutoSwitchProfile(o["auto-switch_profile"], d, "auto_switch_profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["auto-switch_profile"], "ObjectExtenderControllerSimProfile-AutoSwitchProfile"); ok {
					if err = d.Set("auto_switch_profile", vv); err != nil {
						return fmt.Errorf("Error reading auto_switch_profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auto_switch_profile: %v", err)
				}
			}
		}
	}

	if err = d.Set("conn_status", flattenObjectExtenderControllerSimProfileConnStatus(o["conn-status"], d, "conn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn-status"], "ObjectExtenderControllerSimProfile-ConnStatus"); ok {
			if err = d.Set("conn_status", vv); err != nil {
				return fmt.Errorf("Error reading conn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("default_sim", flattenObjectExtenderControllerSimProfileDefaultSim(o["default-sim"], d, "default_sim")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-sim"], "ObjectExtenderControllerSimProfile-DefaultSim"); ok {
			if err = d.Set("default_sim", vv); err != nil {
				return fmt.Errorf("Error reading default_sim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_sim: %v", err)
		}
	}

	if err = d.Set("description", flattenObjectExtenderControllerSimProfileDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ObjectExtenderControllerSimProfile-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("gps", flattenObjectExtenderControllerSimProfileGps(o["gps"], d, "gps")); err != nil {
		if vv, ok := fortiAPIPatch(o["gps"], "ObjectExtenderControllerSimProfile-Gps"); ok {
			if err = d.Set("gps", vv); err != nil {
				return fmt.Errorf("Error reading gps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gps: %v", err)
		}
	}

	if err = d.Set("modem_id", flattenObjectExtenderControllerSimProfileModemId(o["modem-id"], d, "modem_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["modem-id"], "ObjectExtenderControllerSimProfile-ModemId"); ok {
			if err = d.Set("modem_id", vv); err != nil {
				return fmt.Errorf("Error reading modem_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading modem_id: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtenderControllerSimProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtenderControllerSimProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("preferred_carrier", flattenObjectExtenderControllerSimProfilePreferredCarrier(o["preferred-carrier"], d, "preferred_carrier")); err != nil {
		if vv, ok := fortiAPIPatch(o["preferred-carrier"], "ObjectExtenderControllerSimProfile-PreferredCarrier"); ok {
			if err = d.Set("preferred_carrier", vv); err != nil {
				return fmt.Errorf("Error reading preferred_carrier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preferred_carrier: %v", err)
		}
	}

	if err = d.Set("redundant_intf", flattenObjectExtenderControllerSimProfileRedundantIntf(o["redundant-intf"], d, "redundant_intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-intf"], "ObjectExtenderControllerSimProfile-RedundantIntf"); ok {
			if err = d.Set("redundant_intf", vv); err != nil {
				return fmt.Errorf("Error reading redundant_intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_intf: %v", err)
		}
	}

	if err = d.Set("redundant_mode", flattenObjectExtenderControllerSimProfileRedundantMode(o["redundant-mode"], d, "redundant_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["redundant-mode"], "ObjectExtenderControllerSimProfile-RedundantMode"); ok {
			if err = d.Set("redundant_mode", vv); err != nil {
				return fmt.Errorf("Error reading redundant_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redundant_mode: %v", err)
		}
	}

	if err = d.Set("sim1_pin", flattenObjectExtenderControllerSimProfileSim1Pin(o["sim1-pin"], d, "sim1_pin")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim1-pin"], "ObjectExtenderControllerSimProfile-Sim1Pin"); ok {
			if err = d.Set("sim1_pin", vv); err != nil {
				return fmt.Errorf("Error reading sim1_pin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim1_pin: %v", err)
		}
	}

	if err = d.Set("sim2_pin", flattenObjectExtenderControllerSimProfileSim2Pin(o["sim2-pin"], d, "sim2_pin")); err != nil {
		if vv, ok := fortiAPIPatch(o["sim2-pin"], "ObjectExtenderControllerSimProfile-Sim2Pin"); ok {
			if err = d.Set("sim2_pin", vv); err != nil {
				return fmt.Errorf("Error reading sim2_pin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sim2_pin: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectExtenderControllerSimProfileStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectExtenderControllerSimProfile-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectExtenderControllerSimProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileAutoSwitchProfileSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileDefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileGps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfilePreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileRedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileRedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileSim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileSim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerSimProfileSim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtenderControllerSimProfileSim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtenderControllerSimProfileStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtenderControllerSimProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_switch_profile"); ok || d.HasChange("auto_switch_profile") {
		t, err := expandObjectExtenderControllerSimProfileAutoSwitchProfile(d, v, "auto_switch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-switch_profile"] = t
		}
	}

	if v, ok := d.GetOk("conn_status"); ok || d.HasChange("conn_status") {
		t, err := expandObjectExtenderControllerSimProfileConnStatus(d, v, "conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-status"] = t
		}
	}

	if v, ok := d.GetOk("default_sim"); ok || d.HasChange("default_sim") {
		t, err := expandObjectExtenderControllerSimProfileDefaultSim(d, v, "default_sim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-sim"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandObjectExtenderControllerSimProfileDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("gps"); ok || d.HasChange("gps") {
		t, err := expandObjectExtenderControllerSimProfileGps(d, v, "gps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gps"] = t
		}
	}

	if v, ok := d.GetOk("modem_id"); ok || d.HasChange("modem_id") {
		t, err := expandObjectExtenderControllerSimProfileModemId(d, v, "modem_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtenderControllerSimProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("preferred_carrier"); ok || d.HasChange("preferred_carrier") {
		t, err := expandObjectExtenderControllerSimProfilePreferredCarrier(d, v, "preferred_carrier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-carrier"] = t
		}
	}

	if v, ok := d.GetOk("redundant_intf"); ok || d.HasChange("redundant_intf") {
		t, err := expandObjectExtenderControllerSimProfileRedundantIntf(d, v, "redundant_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-intf"] = t
		}
	}

	if v, ok := d.GetOk("redundant_mode"); ok || d.HasChange("redundant_mode") {
		t, err := expandObjectExtenderControllerSimProfileRedundantMode(d, v, "redundant_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-mode"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin"); ok || d.HasChange("sim1_pin") {
		t, err := expandObjectExtenderControllerSimProfileSim1Pin(d, v, "sim1_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim1_pin_code"); ok || d.HasChange("sim1_pin_code") {
		t, err := expandObjectExtenderControllerSimProfileSim1PinCode(d, v, "sim1_pin_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim1-pin-code"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin"); ok || d.HasChange("sim2_pin") {
		t, err := expandObjectExtenderControllerSimProfileSim2Pin(d, v, "sim2_pin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin"] = t
		}
	}

	if v, ok := d.GetOk("sim2_pin_code"); ok || d.HasChange("sim2_pin_code") {
		t, err := expandObjectExtenderControllerSimProfileSim2PinCode(d, v, "sim2_pin_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim2-pin-code"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectExtenderControllerSimProfileStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
