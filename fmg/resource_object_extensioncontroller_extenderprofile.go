// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiExtender extender profile configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectExtensionControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectExtensionControllerExtenderProfileCreate,
		Read:   resourceObjectExtensionControllerExtenderProfileRead,
		Update: resourceObjectExtensionControllerExtenderProfileUpdate,
		Delete: resourceObjectExtensionControllerExtenderProfileDelete,

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
			"_is_factory_setting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cellular": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
												},
												"disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"disconnect_period": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"disconnect_threshold": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"signal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
												},
												"switch_back_timer": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
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
									},
									"gps": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"sim1_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sim1_pin_code": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"sim2_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sim2_pin_code": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
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
												},
												"disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"disconnect_period": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"disconnect_threshold": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"signal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
												},
												"switch_back_timer": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
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
									},
									"gps": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
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
									},
									"sim1_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sim1_pin_code": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"sim2_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"sim2_pin_code": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
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
												},
												"fgt_backup_mode_switch": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"low_signal_strength": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"mode_switch": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"os_image_fallback": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"session_disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"system_reboot": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
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
												},
											},
										},
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backhaul": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"weight": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"backhaul_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"backhaul_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link_loadbalance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"login_password": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectExtensionControllerExtenderProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectExtensionControllerExtenderProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectExtensionControllerExtenderProfile(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectExtensionControllerExtenderProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectExtensionControllerExtenderProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectExtensionControllerExtenderProfile(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectExtensionControllerExtenderProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectExtensionControllerExtenderProfileRead(d, m)
}

func resourceObjectExtensionControllerExtenderProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectExtensionControllerExtenderProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectExtensionControllerExtenderProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectExtensionControllerExtenderProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectExtensionControllerExtenderProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectExtensionControllerExtenderProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectExtensionControllerExtenderProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectExtensionControllerExtenderProfileIsFactorySetting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileBandwidthLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellular(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "controller_report"
	if _, ok := i["controller-report"]; ok {
		result["controller_report"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReport(i["controller-report"], d, pre_append)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtensionControllerExtenderProfileCellularDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1"
	if _, ok := i["modem1"]; ok {
		result["modem1"] = flattenObjectExtensionControllerExtenderProfileCellularModem1(i["modem1"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2"
	if _, ok := i["modem2"]; ok {
		result["modem2"] = flattenObjectExtensionControllerExtenderProfileCellularModem2(i["modem2"], d, pre_append)
	}

	pre_append = pre + ".0." + "sms_notification"
	if _, ok := i["sms-notification"]; ok {
		result["sms_notification"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotification(i["sms-notification"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReport(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReportInterval(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(i["signal-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtensionControllerExtenderProfileCellularControllerReportStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReportInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularControllerReportStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenObjectExtensionControllerExtenderProfileCellularModem1ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenObjectExtensionControllerExtenderProfileCellularModem1DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenObjectExtensionControllerExtenderProfileCellularModem1ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode(i["sim1-pin-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode(i["sim2-pin-code"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch(i["auto-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenObjectExtensionControllerExtenderProfileCellularModem2ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenObjectExtensionControllerExtenderProfileCellularModem2DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem_id"
	if _, ok := i["modem-id"]; ok {
		result["modem_id"] = flattenObjectExtensionControllerExtenderProfileCellularModem2ModemId(i["modem-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Sim1PinCode(i["sim1-pin-code"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenObjectExtensionControllerExtenderProfileCellularModem2Sim2PinCode(i["sim2-pin-code"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2ModemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Sim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularModem2Sim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotification(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(i["alert"], d, pre_append)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(i["receiver"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(i["data-exhausted"], d, pre_append)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(i["fgt-backup-mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(i["low-signal-strength"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(i["mode-switch"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(i["os-image-fallback"], d, pre_append)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(i["session-disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(i["system-reboot"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(i["phone-number"], d, pre_append)
			tmp["phone_number"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-PhoneNumber")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileCellularSmsNotification-Receiver-Status")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileExtension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul"
	if _, ok := i["backhaul"]; ok {
		result["backhaul"] = flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaul(i["backhaul"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {
		result["backhaul_interface"] = flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {
		result["backhaul_ip"] = flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {
		result["ipsec_tunnel"] = flattenObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := i["link-loadbalance"]; ok {
		result["link_loadbalance"] = flattenObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(i["link-loadbalance"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaul(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "ObjectExtensionControllerExtenderProfileLanExtension-Backhaul-Weight")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileLoginPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectExtensionControllerExtenderProfileLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileModel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectExtensionControllerExtenderProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectExtensionControllerExtenderProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("_is_factory_setting", flattenObjectExtensionControllerExtenderProfileIsFactorySetting(o["_is_factory_setting"], d, "_is_factory_setting")); err != nil {
		if vv, ok := fortiAPIPatch(o["_is_factory_setting"], "ObjectExtensionControllerExtenderProfile-IsFactorySetting"); ok {
			if err = d.Set("_is_factory_setting", vv); err != nil {
				return fmt.Errorf("Error reading _is_factory_setting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading _is_factory_setting: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenObjectExtensionControllerExtenderProfileAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ObjectExtensionControllerExtenderProfile-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenObjectExtensionControllerExtenderProfileBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-limit"], "ObjectExtensionControllerExtenderProfile-BandwidthLimit"); ok {
			if err = d.Set("bandwidth_limit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cellular", flattenObjectExtensionControllerExtenderProfileCellular(o["cellular"], d, "cellular")); err != nil {
			if vv, ok := fortiAPIPatch(o["cellular"], "ObjectExtensionControllerExtenderProfile-Cellular"); ok {
				if err = d.Set("cellular", vv); err != nil {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cellular: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cellular"); ok {
			if err = d.Set("cellular", flattenObjectExtensionControllerExtenderProfileCellular(o["cellular"], d, "cellular")); err != nil {
				if vv, ok := fortiAPIPatch(o["cellular"], "ObjectExtensionControllerExtenderProfile-Cellular"); ok {
					if err = d.Set("cellular", vv); err != nil {
						return fmt.Errorf("Error reading cellular: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			}
		}
	}

	if err = d.Set("enforce_bandwidth", flattenObjectExtensionControllerExtenderProfileEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-bandwidth"], "ObjectExtensionControllerExtenderProfile-EnforceBandwidth"); ok {
			if err = d.Set("enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("extension", flattenObjectExtensionControllerExtenderProfileExtension(o["extension"], d, "extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension"], "ObjectExtensionControllerExtenderProfile-Extension"); ok {
			if err = d.Set("extension", vv); err != nil {
				return fmt.Errorf("Error reading extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectExtensionControllerExtenderProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectExtensionControllerExtenderProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan_extension", flattenObjectExtensionControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["lan-extension"], "ObjectExtensionControllerExtenderProfile-LanExtension"); ok {
				if err = d.Set("lan_extension", vv); err != nil {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenObjectExtensionControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["lan-extension"], "ObjectExtensionControllerExtenderProfile-LanExtension"); ok {
					if err = d.Set("lan_extension", vv); err != nil {
						return fmt.Errorf("Error reading lan_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("login_password", flattenObjectExtensionControllerExtenderProfileLoginPassword(o["login-password"], d, "login_password")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-password"], "ObjectExtensionControllerExtenderProfile-LoginPassword"); ok {
			if err = d.Set("login_password", vv); err != nil {
				return fmt.Errorf("Error reading login_password: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_password: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenObjectExtensionControllerExtenderProfileLoginPasswordChange(o["login-password-change"], d, "login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-password-change"], "ObjectExtensionControllerExtenderProfile-LoginPasswordChange"); ok {
			if err = d.Set("login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("model", flattenObjectExtensionControllerExtenderProfileModel(o["model"], d, "model")); err != nil {
		if vv, ok := fortiAPIPatch(o["model"], "ObjectExtensionControllerExtenderProfile-Model"); ok {
			if err = d.Set("model", vv); err != nil {
				return fmt.Errorf("Error reading model: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectExtensionControllerExtenderProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectExtensionControllerExtenderProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectExtensionControllerExtenderProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectExtensionControllerExtenderProfileIsFactorySetting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileBandwidthLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellular(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "controller_report"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularControllerReport(d, i["controller_report"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["controller-report"] = t
		}
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtensionControllerExtenderProfileCellularDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "modem1"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1(d, i["modem1"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["modem1"] = t
		}
	}
	pre_append = pre + ".0." + "modem2"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem2(d, i["modem2"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["modem2"] = t
		}
	}
	pre_append = pre + ".0." + "sms_notification"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotification(d, i["sms_notification"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["sms-notification"] = t
		}
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["interval"], _ = expandObjectExtensionControllerExtenderProfileCellularControllerReportInterval(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(d, i["signal_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtensionControllerExtenderProfileCellularControllerReportStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReportInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularControllerReportStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch(d, i["auto_switch"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["auto-switch"] = t
		}
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["conn-status"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["default-sim"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["gps"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "modem_id"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem-id"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2ModemId(d, i["modem_id"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["preferred-carrier"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-intf"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["redundant-mode"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim1-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sim2-pin-code"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dataplan"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-period"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["disconnect-threshold"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["signal"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-time"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["switch-back-timer"], _ = expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2ModemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d, i["alert"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["alert"] = t
		}
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d, i["receiver"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["receiver"] = t
		}
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-exhausted"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d, i["data_exhausted"], pre_append)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fgt-backup-mode-switch"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d, i["fgt_backup_mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["low-signal-strength"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d, i["low_signal_strength"], pre_append)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mode-switch"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d, i["mode_switch"], pre_append)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os-image-fallback"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d, i["os_image_fallback"], pre_append)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["session-disconnect"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d, i["session_disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["system-reboot"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d, i["system_reboot"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["alert"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["phone-number"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d, i["phone_number"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileCellularSmsNotificationStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d, i["backhaul"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["backhaul"] = t
		}
	}
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-interface"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-ip"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-tunnel"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["link-loadbalance"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(d, i["link_loadbalance"], pre_append)
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileLoginPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectExtensionControllerExtenderProfileLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileModel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectExtensionControllerExtenderProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectExtensionControllerExtenderProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("_is_factory_setting"); ok || d.HasChange("_is_factory_setting") {
		t, err := expandObjectExtensionControllerExtenderProfileIsFactorySetting(d, v, "_is_factory_setting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["_is_factory_setting"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandObjectExtensionControllerExtenderProfileAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok || d.HasChange("bandwidth_limit") {
		t, err := expandObjectExtensionControllerExtenderProfileBandwidthLimit(d, v, "bandwidth_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("cellular"); ok || d.HasChange("cellular") {
		t, err := expandObjectExtensionControllerExtenderProfileCellular(d, v, "cellular")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cellular"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok || d.HasChange("enforce_bandwidth") {
		t, err := expandObjectExtensionControllerExtenderProfileEnforceBandwidth(d, v, "enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("extension"); ok || d.HasChange("extension") {
		t, err := expandObjectExtensionControllerExtenderProfileExtension(d, v, "extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectExtensionControllerExtenderProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok || d.HasChange("lan_extension") {
		t, err := expandObjectExtensionControllerExtenderProfileLanExtension(d, v, "lan_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok || d.HasChange("login_password") {
		t, err := expandObjectExtensionControllerExtenderProfileLoginPassword(d, v, "login_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok || d.HasChange("login_password_change") {
		t, err := expandObjectExtensionControllerExtenderProfileLoginPasswordChange(d, v, "login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("model"); ok || d.HasChange("model") {
		t, err := expandObjectExtensionControllerExtenderProfileModel(d, v, "model")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectExtensionControllerExtenderProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
