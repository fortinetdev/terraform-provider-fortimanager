// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender cellular configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfileCellular() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileCellularUpdate,
		Read:   resourceObjectExtensionControllerExtenderProfileCellularRead,
		Update: resourceObjectExtensionControllerExtenderProfileCellularUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileCellularDelete,

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
			"controller_report": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"signal_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dataplan": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"modem1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"modem2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"sms_notification": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileCellularUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfileCellular(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellular resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerExtenderProfileCellular(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfileCellular resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectExtensionControllerExtenderProfileCellular")

	return resourceObjectExtensionControllerExtenderProfileCellularRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileCellularDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfileCellular(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfileCellular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileCellularRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfileCellular(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellular resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfileCellular(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfileCellular resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReport2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReportInterval2edl(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(i["signal-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReportStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReportInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReportStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem12edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenObjectExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenObjectExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Gps2edl(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenObjectExtensionControllerExtenderProfileCellularModem1ModemId2edl(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(i["sim2-pin"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Gps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1ModemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem22edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenObjectExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenObjectExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Gps2edl(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenObjectExtensionControllerExtenderProfileCellularModem2ModemId2edl(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(i["sim2-pin"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Gps2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2ModemId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotification2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(i["alert"], d, pre_append)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(i["receiver"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfileCellular(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if isImportTable() {
		if err = d.Set("controller_report", flattenObjectExtensionControllerExtenderProfileCellularControllerReport2edl(o["controller-report"], d, "controller_report")); err != nil {
			if vv, ok := fortiAPIPatch(o["controller-report"], "ObjectExtensionControllerExtenderProfileCellular-ControllerReport"); ok {
				if err = d.Set("controller_report", vv); err != nil {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading controller_report: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("controller_report"); ok {
			if err = d.Set("controller_report", flattenObjectExtensionControllerExtenderProfileCellularControllerReport2edl(o["controller-report"], d, "controller_report")); err != nil {
				if vv, ok := fortiAPIPatch(o["controller-report"], "ObjectExtensionControllerExtenderProfileCellular-ControllerReport"); ok {
					if err = d.Set("controller_report", vv); err != nil {
						return fmt.Errorf("Error reading controller_report: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			}
		}
	}

	if err = d.Set("dataplan", flattenObjectExtensionControllerExtenderProfileCellularDataplan2edl(o["dataplan"], d, "dataplan")); err != nil {
		if vv, ok := fortiAPIPatch(o["dataplan"], "ObjectExtensionControllerExtenderProfileCellular-Dataplan"); ok {
			if err = d.Set("dataplan", vv); err != nil {
				return fmt.Errorf("Error reading dataplan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dataplan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("modem1", flattenObjectExtensionControllerExtenderProfileCellularModem12edl(o["modem1"], d, "modem1")); err != nil {
			if vv, ok := fortiAPIPatch(o["modem1"], "ObjectExtensionControllerExtenderProfileCellular-Modem1"); ok {
				if err = d.Set("modem1", vv); err != nil {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading modem1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem1"); ok {
			if err = d.Set("modem1", flattenObjectExtensionControllerExtenderProfileCellularModem12edl(o["modem1"], d, "modem1")); err != nil {
				if vv, ok := fortiAPIPatch(o["modem1"], "ObjectExtensionControllerExtenderProfileCellular-Modem1"); ok {
					if err = d.Set("modem1", vv); err != nil {
						return fmt.Errorf("Error reading modem1: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("modem2", flattenObjectExtensionControllerExtenderProfileCellularModem22edl(o["modem2"], d, "modem2")); err != nil {
			if vv, ok := fortiAPIPatch(o["modem2"], "ObjectExtensionControllerExtenderProfileCellular-Modem2"); ok {
				if err = d.Set("modem2", vv); err != nil {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading modem2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem2"); ok {
			if err = d.Set("modem2", flattenObjectExtensionControllerExtenderProfileCellularModem22edl(o["modem2"], d, "modem2")); err != nil {
				if vv, ok := fortiAPIPatch(o["modem2"], "ObjectExtensionControllerExtenderProfileCellular-Modem2"); ok {
					if err = d.Set("modem2", vv); err != nil {
						return fmt.Errorf("Error reading modem2: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sms_notification", flattenObjectExtensionControllerExtenderProfileCellularSmsNotification2edl(o["sms-notification"], d, "sms_notification")); err != nil {
			if vv, ok := fortiAPIPatch(o["sms-notification"], "ObjectExtensionControllerExtenderProfileCellular-SmsNotification"); ok {
				if err = d.Set("sms_notification", vv); err != nil {
					return fmt.Errorf("Error reading sms_notification: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sms_notification: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sms_notification"); ok {
			if err = d.Set("sms_notification", flattenObjectExtensionControllerExtenderProfileCellularSmsNotification2edl(o["sms-notification"], d, "sms_notification")); err != nil {
				if vv, ok := fortiAPIPatch(o["sms-notification"], "ObjectExtensionControllerExtenderProfileCellular-SmsNotification"); ok {
					if err = d.Set("sms_notification", vv); err != nil {
						return fmt.Errorf("Error reading sms_notification: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sms_notification: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileCellularFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interval"], _ = expandObjectExtensionControllerExtenderProfileCellularControllerReportInterval2edl(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(d, i["signal_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtensionControllerExtenderProfileCellularControllerReportStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReportInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReportStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Gps2edl(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1ModemId2edl(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode2edl(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode2edl(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1ConnStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1DefaultSim2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Gps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1ModemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1RedundantMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Gps2edl(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2ModemId2edl(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim1PinCode2edl(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim2PinCode2edl(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2ConnStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2DefaultSim2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Gps2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2ModemId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2RedundantMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim1PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim2PinCode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotification2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(d, i["alert"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["alert"] = t
		}
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(d, i["receiver"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["receiver"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["alert"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(d, i["status"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfileCellular(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("controller_report"); ok || d.HasChange("controller_report") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularControllerReport2edl(d, v, "controller_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["controller-report"] = t
		}
	}

	if v, ok := d.GetOk("dataplan"); ok || d.HasChange("dataplan") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularDataplan2edl(d, v, "dataplan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dataplan"] = t
		}
	}

	if v, ok := d.GetOk("modem1"); ok || d.HasChange("modem1") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem12edl(d, v, "modem1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1"] = t
		}
	}

	if v, ok := d.GetOk("modem2"); ok || d.HasChange("modem2") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem22edl(d, v, "modem2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2"] = t
		}
	}

	if v, ok := d.GetOk("sms_notification"); ok || d.HasChange("sms_notification") {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotification2edl(d, v, "sms_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-notification"] = t
		}
	}

	return &obj, nil
}
